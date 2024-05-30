package isemsmrsicimportmodule

import (
	"ISEMS-MRSICT_importmodule/confighandler"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/av-belyakov/simplelogger"
)

const ROOT_DIR = "placeholder_elasticsearch"

func getLoggerSettings(cls []confighandler.LogSet) []simplelogger.MessageTypeSettings {
	loggerConf := make([]simplelogger.MessageTypeSettings, 0, len(cls))

	for _, v := range cls {
		loggerConf = append(loggerConf, simplelogger.MessageTypeSettings{
			MsgTypeName:   v.MsgTypeName,
			WritingFile:   v.WritingFile,
			PathDirectory: v.PathDirectory,
			WritingStdout: v.WritingStdout,
			MaxFileSize:   v.MaxFileSize,
		})
	}

	return loggerConf
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	//инициализируем модуль чтения конфигурационного файла
	confApp, err := confighandler.NewConfig(ROOT_DIR)
	if err != nil {
		log.Fatalf("error module 'confighandler': %v", err)
	}

	//инициализируем модуль логирования
	sl, err := simplelogger.NewSimpleLogger(ROOT_DIR, getLoggerSettings(confApp.GetListLogs()))
	if err != nil {
		log.Fatalf("error module 'simplelogger': %v", err)
	}

	go func() {
		osCall := <-sigChan
		msg := fmt.Sprintf("stop 'main' function, %s", osCall.String())
		_ = sl.WriteLoggingData(msg, "info")

		/*
			close(counting)
			close(logging)
			close(channelZabbix)

			ctxCancelZ()
			ctxCancelCore()
		*/
	}()
}
