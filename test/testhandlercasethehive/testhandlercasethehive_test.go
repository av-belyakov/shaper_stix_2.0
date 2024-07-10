package testhandlercasethehive_test

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/av-belyakov/shaper_stix_2.1/confighandler"
	"github.com/av-belyakov/shaper_stix_2.1/databaseapi/mongodbapi"
	wrap "github.com/av-belyakov/shaper_stix_2.1/databaseapi/mongodbapi/wrappers"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal"
	"github.com/av-belyakov/shaper_stix_2.1/internal/decodejson"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix/domainobjects"
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

		confApp    confighandler.ConfigApp
		errConfApp error

		connMdb    *mongo.Client
		errConnMdb error
	)

	BeforeAll(func() {
		//chanStoppedCounting = make(chan struct{})
		chanOutputDecodeJson = make(chan datamodels.ChanOutputDecodeJSON)
		logging = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)

		confApp, errConfApp = confighandler.NewConfig("shaper_stix_2.1")

		mongoDBModule = &mongodbapi.MongoDBModule{
			ChanInputToModule:    make(chan mongodbapi.ChanInput),
			ChanOutputFromModule: make(chan mongodbapi.ChanOutput),
		}

		ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
		connMdb, errConnMdb = mongodbapi.NewConnection(ctx, *confApp.GetAppMongoDB())

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

			Expect(errConfApp).ShouldNot(HaveOccurred())
			Expect(errConnMdb).ShouldNot(HaveOccurred())
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

			//
			// загрузка STIX объектов в MongoDB
			//
			qp := wrap.QueryParameters{
				ConnectDB:      connMdb,
				NameDB:         confApp.NameDB,
				CollectionName: "stix_object_collection",
			}

			fmt.Println(qp)

			type listObject map[string][]interface{}

			lo := make(listObject, 0)

			//преобразуем в новый список объектов STIX по их типам
			if list, ok := objects.Data.([]datamodels.HandlerSTIXObject); ok {
				for k, v := range list {
					fmt.Printf("%d. \n", k)
					fmt.Println(v.ToStringBeautiful(1))

					if _, ok := lo[v.GetType()]; !ok {
						lo[v.GetType()] = []interface{}(nil)
					}

					lo[v.GetType()] = append(lo[v.GetType()], v.GetObject())
				}
			}

			//Получаем список case_id по которым будем искать объекты Report
			reportCases := []string(nil)
			reportInt, ok := lo["report"]
			Expect(ok).Should(BeTrue())
			//для reports
			for _, v := range reportInt {
				report, ok := v.(*domainobjects.WrapperReport)
				caseId := report.GetReportOutsideSpecification().CaseId
				reportCases = append(reportCases, caseId)
				fmt.Println("REPORT:", report)
				Expect(ok).Should(BeTrue())
			}

			//выполняем поиск объектов типа 'report' с подходящими case_id
			cur, err := qp.Find((bson.D{{Key: "outside_specification.case_id", Value: bson.D{{Key: "$in", Value: reportCases}}}}))
			Expect(err).ShouldNot(HaveOccurred())

			//fmt.Println("==== List id:")
			//sort.Slice(listId, func(i, j int) bool {
			//	a := strings.Split(listId[i], "-")
			//	b := strings.Split(listId[j], "-")

			//	return a[0] < b[0]
			//})
			//for _, v := range listId {
			//	fmt.Println(v)
			//}

			//*******************************************************
			// Insert работает, но пока его закоментируем для других
			//тестов
			//-------------------------------------------------------
			//_, err := qp.InsertData(listInterface, []mongo.IndexModel{
			//	{
			//		Keys: bson.D{
			//			{Key: "commonpropertiesobjectstix.type", Value: 1},
			//			{Key: "commonpropertiesobjectstix.id", Value: 1},
			//			{Key: "outside_specification.case_id", Value: 1},
			//		},
			//		Options: &options.IndexOptions{},
			//	},
			//})
			//Expect(err).ShouldNot(HaveOccurred())

			//1. Получить список STIX объектов по case id (для report), и по
			//common_outside_specification.element_id для всех остальных
			//
			//2. Выполнить сравнение с имеющимися STIX объектами
			//
			//3. Внести изменения в существующие на основе новых и загрузить в БД

			Expect(true).Should(BeTrue())
		})
	})
})
