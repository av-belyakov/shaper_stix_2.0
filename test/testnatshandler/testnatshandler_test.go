package testnatshandler_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"shaper_stix/confighandler"
	"shaper_stix/datamodels"
	"shaper_stix/natsapi"
	"shaper_stix/supportingfunctions"
)

const DIR_ROOT = "shaper_stix_2.1"

var _ = Describe("Testnatshandler", Ordered, func() {
	var (
		errConf, errConnect error
		conf                confighandler.ConfigApp

		logging  = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)
		done     = make(chan struct{})
		mnats    *natsapi.ModuleNATS
	)

	BeforeAll(func() {
		conf, errConf = confighandler.NewConfig(DIR_ROOT)

		go func() {
			fmt.Println("___ Logging START")
			defer fmt.Println("___ Logging STOP")

			select {
			case log := <-logging:
				if log.MsgType == "error" {
					fmt.Println("----", log, "----")
				}

				if log.MsgType == "STOP TEST" {
					return
				}

			case <-done:
				return
			}
		}()

		mnats, errConnect = natsapi.NewClientNATS(*conf.GetAppNATS(), logging, counting)
	})

	Context("Тест 0. Чтение конфигурвционного файла", func() {
		It("При чтении конфигурационного файла не должно быть ошибок", func() {
			Expect(errConf).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 1. Подключение к NATS", func() {
		It("Должно быть успешно установлено подключение к NATS", func() {
			Expect(errConnect).ShouldNot(HaveOccurred())
		})

		It("Должно быть получено сообщение типа 'alertupdate'", func() {
			var str string
			var err error

			for data := range mnats.GetDataReceptionChannel() {
				if data.SubjectType == "subject_alert" {
					done <- struct{}{}

					str, err = supportingfunctions.NewReadReflectJSONSprint(data.Data)

					break
				}
			}

			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println(str)
			Expect(len(str)).ShouldNot(Equal(0))
		})
	})
})
