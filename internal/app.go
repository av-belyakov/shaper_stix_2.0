package internal

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/av-belyakov/simplelogger"

	"shaper_stix/confighandler"
	"shaper_stix/databaseapi/mongodbapi"
	"shaper_stix/datamodels"
	"shaper_stix/memorytemporarystorage"
	"shaper_stix/natsapi"
	"shaper_stix/zabbixapi"
)

// NewApp инициализирует приложение
func NewApp(confApp confighandler.ConfigApp, sl simplelogger.SimpleLoggerSettings) error {
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

	//sl.WriteLoggingData
	if natsModule, err := natsapi.NewClientNATS(*confApp.GetAppNATS(), logging, counting); err != nil {
		return fmt.Errorf("error module 'natsapi': %w", err)
	}

	if mdbModule, err := mongodbapi.NewClientMongoDB(*confApp.GetAppMongoDB(), logging, counting); err != nil {
		return fmt.Errorf("error module 'mongodbapi': %w", err)
	}

	//вывод сообщения о запуске приложения с указанием версии приложения
	if msg, err := writeLaunchMessage(); err != nil {
		sl.WriteLoggingData(fmt.Sprint(err), "error")
		log.Println(err)
	} else {
		log.Println(msg)
	}

	//
	//
	//тут уже непосредственно блокирующий роутер с обработчиками
	//
	//

	return nil
}
