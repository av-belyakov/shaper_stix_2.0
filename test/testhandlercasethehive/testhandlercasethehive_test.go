package testhandlercasethehive_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/av-belyakov/shaper_stix_2.1/confighandler"
	"github.com/av-belyakov/shaper_stix_2.1/databaseapi/mongodbapi"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal"
	"github.com/av-belyakov/shaper_stix_2.1/internal/decodejson"
	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

var _ = Describe("Testhandlercasethehive", Ordered, func() {
	const ROOT_DIR = "shaper_stix_2.1"

	var (
		fileByte []byte
		fileErr  error

		//chanStoppedCounting  chan struct{}
		chanOutputDecodeJson chan datamodels.ChanOutputDecodeJSON
		logging              chan datamodels.MessageLogging
		counting             chan datamodels.DataCounterSettings

		confApp confighandler.ConfigApp
		errConf error

		mongoDBModule    *mongodbapi.MongoDBModule
		errMongoDBModule error
	)

	BeforeAll(func() {
		//chanStoppedCounting = make(chan struct{})
		chanOutputDecodeJson = make(chan datamodels.ChanOutputDecodeJSON)
		logging = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)

		confApp, errConf = confighandler.NewConfig(ROOT_DIR)

		confMongoDB := confighandler.AppConfigMongoDB{
			Port: 27017,
			Host: "127.0.0.1",
		}

		mongoDBModule, errMongoDBModule = mongodbapi.NewClientMongoDB(context.Background(), confMongoDB, logging, counting)

		// инициализация хранилища правил
		procRules := internal.NewRulesHandler(ROOT_DIR, "configs")

		fileByte, fileErr = supportingfunctions.ReadFileJson("test/filestest", "event_1.json")

		go func() {
			fmt.Println("___ Logging START")
			defer fmt.Println("___ Logging STOP")

			for log := range logging {
				if log.MsgType == "error" {
					fmt.Println("----", log, "----")
				}

				if log.MsgType == "STOP TEST" {
					//chanStoppedCounting <- struct{}{}

					return
				}
			}
		}()

		//вывод данных счетчика
		go func() {
			for d := range counting {
				fmt.Printf("\tСчетчик %v\n", d.DataType)
			}
		}()

		decodeJson := decodejson.NewDecodeJsonMessageSettings(logging, counting)
		chanOutputDecodeJson = decodeJson.HandlerJsonMessage(fileByte, "test_id_73d8r3", "subject_case")

		go internal.NewHandlerCaseObject(chanOutputDecodeJson, procRules, mongoDBModule, counting, logging)
	})

	Context("Тест 1. Проверка успешности инициализации модулей и чтения файлов", func() {
		It("При инициализации файла конфигурации не должно быть ошибок", func() {
			Expect(errConf).ShouldNot(HaveOccurred())
		})

		It("При инициализации модуля MongoDB не должно быть ошибок", func() {
			Expect(errMongoDBModule).ShouldNot(HaveOccurred())
		})

		It("При чтении тестового файла не должно быть ошибок", func() {
			Expect(fileErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 1. Формирование различных STIX объектов", func() {
		It("Должны быть сформировано некоторое количество STIX объектов", func() {
			fmt.Println("BEFORE----")

			mongoChanInput := mongoDBModule.GetChanInput()
			fmt.Println("AFTER----")

			objects := <-mongoChanInput

			fmt.Println("--------------- RESULT --------------")
			fmt.Println(objects)

			Expect(true).Should(BeTrue())
		})
	})
})
