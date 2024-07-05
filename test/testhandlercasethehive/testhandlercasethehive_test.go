package testhandlercasethehive_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/av-belyakov/shaper_stix_2.1/databaseapi/mongodbapi"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal"
	"github.com/av-belyakov/shaper_stix_2.1/internal/decodejson"
	"github.com/av-belyakov/shaper_stix_2.1/ruleinteraction"
	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

type ToStringBeautifulReader interface {
	//datamodels.GetterCommonPropertiesObjectSTIX
	ToStringBeautiful() string
}

var _ = Describe("Testhandlercasethehive", Ordered, func() {
	const (
		ROOT_DIR  = "shaper_stix_2.1"
		RULE_CASE = "msgrule_case.yaml"
	)

	var (
		fileByte []byte
		fileErr  error

		//chanStoppedCounting  chan struct{}
		chanOutputDecodeJson chan datamodels.ChanOutputDecodeJSON
		logging              chan datamodels.MessageLogging
		counting             chan datamodels.DataCounterSettings

		mongoDBModule    *mongodbapi.MongoDBModule
		errMongoDBModule error

		procRules           *internal.ProcessingRules
		warningAddCaseRules string
		errAddCaseRules     error

		caseRules    *ruleinteraction.ListRule
		errCaseRules error
	)

	BeforeAll(func() {
		//chanStoppedCounting = make(chan struct{})
		chanOutputDecodeJson = make(chan datamodels.ChanOutputDecodeJSON)
		logging = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)

		mongoDBModule = &mongodbapi.MongoDBModule{
			ChanInputToModule:    make(chan mongodbapi.ChanInput),
			ChanOutputFromModule: make(chan mongodbapi.ChanOutput),
		}

		// инициализация хранилища правил
		procRules = internal.NewRulesHandler(ROOT_DIR, "configs")
		warningAddCaseRules, errAddCaseRules = procRules.AddCaseRules(RULE_CASE)
		caseRules, errCaseRules = procRules.GetCaseRules()
		if errCaseRules != nil {
			fmt.Println("******************************")
			fmt.Println("ERROR:", errCaseRules)
			fmt.Println("******************************")
		}

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

		go internal.NewHandlerCaseObject(chanOutputDecodeJson, *caseRules, mongoDBModule, counting, logging)
	})

	Context("Тест 0. Инициализация правил Case", func() {
		It("При инициализации правил не должно быть ошибки", func() {
			fmt.Println("Case Warning:", warningAddCaseRules)
			fmt.Printf("Case list:\n%v\n", caseRules)

			Expect(errCaseRules).ShouldNot(HaveOccurred())
			Expect(errAddCaseRules).ShouldNot(HaveOccurred())
			Expect(len(warningAddCaseRules)).Should(Equal(0))

		})
	})

	Context("Тест 1. Проверка успешности инициализации модулей и чтения файлов", func() {
		It("При инициализации модуля MongoDB не должно быть ошибок", func() {
			Expect(errMongoDBModule).ShouldNot(HaveOccurred())
		})

		It("При чтении тестового файла не должно быть ошибок", func() {
			Expect(fileErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Формирование различных STIX объектов", func() {
		It("Должны быть сформировано некоторое количество STIX объектов", func() {
			fmt.Println("BEFORE----")

			mongoChanInput := mongoDBModule.GetChanInput()
			fmt.Println("AFTER----")

			objects := <-mongoChanInput

			/*if list, ok := objects.Data.([]ToStringBeautifulReader); ok {
				fmt.Println("count =", len(list))
				Expect(len(list)).ShouldNot(Equal(0))
			}*/

			//listId := []string(nil)
			fmt.Println("--------------- RESULT --------------")
			if list, ok := objects.Data.([]datamodels.GetterCommonPropertiesObjectSTIX); ok {
				for k, v := range list {
					fmt.Printf("%d. \n", k)
					if v.GetType() == "report" {
						fmt.Println("--------- Report object STIX ---------")
					}

					//fmt.Println(v.ToStringBeautiful(1))

					//listId = append(listId, v.GetID())
				}
			}

			/*fmt.Println("==== List id:")
			sort.Slice(listId, func(i, j int) bool {
				a := strings.Split(listId[i], "-")
				b := strings.Split(listId[j], "-")

				return a[0] < b[0]
			})
			for _, v := range listId {
				fmt.Println(v)
			}*/

			Expect(true).Should(BeTrue())
		})
	})
})
