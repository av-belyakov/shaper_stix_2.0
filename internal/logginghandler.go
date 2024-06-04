package internal

import (
	"fmt"

	"github.com/av-belyakov/simplelogger"

	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/zabbixapi"
)

// NewLoggingHandler обработчик логов выполняет обработку логов и запись их
// в файл, а также передачу логов модулю Zabbix
func NewLoggingHandler(
	channelZabbix chan<- zabbixapi.MessageSettings,
	sl simplelogger.SimpleLoggerSettings,
	logging <-chan datamodels.MessageLogging) {
	for msg := range logging {
		_ = sl.WriteLoggingData(msg.MsgData, msg.MsgType)

		if msg.MsgType == "error" || msg.MsgType == "warning" {
			channelZabbix <- zabbixapi.MessageSettings{
				EventType: "error",
				Message:   fmt.Sprintf("%s: %s", msg.MsgType, msg.MsgData),
			}
		}

		if msg.MsgType == "info" {
			channelZabbix <- zabbixapi.MessageSettings{
				EventType: "info",
				Message:   msg.MsgData,
			}
		}
	}
}
