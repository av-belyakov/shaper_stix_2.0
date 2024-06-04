package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/av-belyakov/simplelogger"

	"github.com/av-belyakov/shaper_stix_2.1/confighandler"
	"github.com/av-belyakov/shaper_stix_2.1/databaseapi/mongodbapi"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal/decodejson"
	"github.com/av-belyakov/shaper_stix_2.1/memorytemporarystorage"
	"github.com/av-belyakov/shaper_stix_2.1/natsapi"
	"github.com/av-belyakov/shaper_stix_2.1/zabbixapi"
)

type shortEventSettings struct {
	Source string         `json:"source"`
	Event  nameObjectType `json:"event"`
}

type nameObjectType struct {
	ObjectType string `json:"objectType"`
}

// NewApp инициализирует приложение
func NewApp(ctx context.Context, confApp confighandler.ConfigApp, sl simplelogger.SimpleLoggerSettings) error {
	//инициализируем модуль временного хранения информации
	storageApp := memorytemporarystorage.NewTemporaryStorage()
	//добавляем время инициализации счетчика хранения
	storageApp.SetStartTimeDataCounter(time.Now())

	// инициализация модуля для взаимодействия с Zabbix
	channelZabbix := make(chan zabbixapi.MessageSettings)
	ctxz, ctxCancelZ := context.WithCancel(context.Background())
	defer func() {
		ctxCancelZ()
		close(channelZabbix)
	}()
	if err := zabbixapi.InteractionZabbix(ctxz, confApp, sl, channelZabbix); err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err.Error(), f, l-3), "error")

		log.Fatalln(err.Error())
	}

	// инициализация обработчика счетчика
	counting := make(chan datamodels.DataCounterSettings)
	go NewCounterHandler(channelZabbix, storageApp, counting)

	// инициализация обработчика для логирования данных
	logging := make(chan datamodels.MessageLogging)
	go NewLoggingHandler(channelZabbix, sl, logging)

	defer func() {
		close(counting)
		close(logging)
	}()

	natsModule, err := natsapi.NewClientNATS(*confApp.GetAppNATS(), logging, counting)
	if err != nil {
		return fmt.Errorf("error module 'natsapi': %w", err)
	}

	ctxMdb, ctxCancelMdb := context.WithCancel(context.Background())
	defer ctxCancelMdb()

	mdbModule, err := mongodbapi.NewClientMongoDB(ctxMdb, *confApp.GetAppMongoDB(), logging, counting)
	if err != nil {
		return fmt.Errorf("error module 'mongodbapi': %w", err)
	}

	//вывод сообщения о запуске приложения с указанием версии приложения
	if msg, err := writeLaunchMessage(); err != nil {
		sl.WriteLoggingData(fmt.Sprint(err), "error")
		log.Println(err)
	} else {
		log.Println(msg)
	}

	decodeJson := decodejson.NewDecodeJsonMessageSettings(logging, counting)

	for {
		select {
		case <-ctx.Done():
			return nil

		case data := <-natsModule.GetDataReceptionChannel():
			fmt.Println("func 'NewApp' прием данных из канала для взаимодействия с NATS")
			fmt.Println("DATA:", data)

			eventSettings := shortEventSettings{}
			if err := json.Unmarshal(data.Data, &eventSettings); err != nil {
				_, f, l, _ := runtime.Caller(0)
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l+1),
					MsgType: "error",
				}

				continue
			}

			switch eventSettings.Event.ObjectType {
			case "case":
				chanOutputDecodeJson, chanDecodeJsonDone := decodeJson.HandlerJsonMessage(data.Data, data.MsgId, "subject_case")

				//				chansOut := supportingfunctions.CreateChannelDuplication[datamodels.ChanOutputDecodeJSON](chanOutputDecodeJson, 2)
				//				chansDone := supportingfunctions.CreateChannelDuplication[bool](chanDecodeJsonDone, 2)

				//используется для хранения в MongoDB
				//				go NewVerifiedTheHiveFormatCase(chansOut[0], chansDone[0], mdbModule, settings.logging)
				//используется для хранения в Elasticsearch
				//				go NewVerifiedElasticsearchFormatCase(chansOut[1], chansDone[1], esModule, settings.logging)

			case "alert":
			//				chanOutputDecodeJson, chanDecodeJsonDone := decodeJsonAlert.HandlerJsonMessage(data.Data, data.MsgId, "subject_alert")

			//				chansOut := supportingfunctions.CreateChannelDuplication[datamodels.ChanOutputDecodeJSON](chanOutputDecodeJson, 2)
			//				chansDone := supportingfunctions.CreateChannelDuplication[bool](chanDecodeJsonDone, 2)

			//используется для хранения в MongoDB
			//				go NewVerifiedTheHiveFormatAlert(chansOut[0], chansDone[0], mdbModule, settings.logging)
			//используется для хранения в Elasticsearch
			//				go NewVerifiedElasticsearchFormatAlert(chansOut[1], chansDone[1], esModule, settings.logging)

			default:
				_, f, l, _ := runtime.Caller(0)
				logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'undefined type objectType' %s:%d", f, l+1),
					MsgType: "error",
				}
			}

		case data := <-mdbModule.GetChanOutput():
			fmt.Println("func 'NewApp' прием данных из канала для взаимодействия с СУБД MongoDB")
			fmt.Println("DATA:", data)

		}
	}
}
