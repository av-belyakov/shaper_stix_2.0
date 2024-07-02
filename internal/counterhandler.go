package internal

import (
	"fmt"
	"log"
	"time"

	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/memorytemporarystorage"
	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
	"github.com/av-belyakov/shaper_stix_2.1/zabbixapi"
)

// NewCounterHandler обработчик счетчиков
func NewCounterHandler(
	channelZabbix chan<- zabbixapi.MessageSettings,
	storageApp *memorytemporarystorage.CommonStorageTemporary,
	counting <-chan datamodels.DataCounterSettings) {

	for data := range counting {
		d, h, m, s := supportingfunctions.GetDifference(storageApp.GetStartTimeDataCounter(), time.Now())
		patternTime := fmt.Sprintf("со старта приложения: дней %d, часов %d, минут %d, секунд %d", d, h, m, s)
		var msg string

		switch data.DataType {
		case "update accepted events":
			storageApp.IncrementAcceptedEvents()
			msg = fmt.Sprintf("принято: %d, %s", storageApp.GetAcceptedEvents(), patternTime)

		case "update processed events":
			storageApp.IncrementProcessedEvents()
			msg = fmt.Sprintf("обработано: %d, %s", storageApp.GetProcessedEvents(), patternTime)

		case "update count insert MongoDB":
			if data.DataMsg == "subject_case" {
				storageApp.IncrementCaseInsertMongoDB()
				msg = fmt.Sprintf("подписка-'subject_case', добавлено в MongoDB: %d, %s", storageApp.GetCaseInsertMongoDB(), patternTime)
			}

			if data.DataMsg == "subject_alert" {
				storageApp.IncrementAlertInsertMongoDB()
				msg = fmt.Sprintf("подписка-'subject_alert', добавлено в MongoDB: %d, %s", storageApp.GetAlertInsertMongoDB(), patternTime)
			}
		}

		log.Printf("\t%s\n", msg)
		channelZabbix <- zabbixapi.MessageSettings{
			EventType: "info",
			Message:   msg,
		}
	}
}
