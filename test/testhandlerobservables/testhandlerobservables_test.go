package testhandlerobservables_test

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal/decodejson"
	"github.com/av-belyakov/shaper_stix_2.1/internal/listhandlerjson"
	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

func handlerCaseToTmpObservables(
	input <-chan datamodels.ChanOutputDecodeJSON,
	listHandlerObservables map[string][]func(interface{}),
	logging chan<- datamodels.MessageLogging) {

	for data := range input {
		//************************************************************************
		//********** Сбор всех объектов относящихся к полю Observables  **********
		// для всех полей входящих в observables, кроме содержимого
		//поля reports
		if lf, ok := listHandlerObservables[data.FieldBranch]; ok {
			//fmt.Println("func 'handlerCaseToTmpObservables', fieldBranch: ", data.FieldBranch, ", value:", data.Value)

			for _, f := range lf {
				r := reflect.TypeOf(data.Value)
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := data.Value.([]interface{}); ok {
						for _, value := range s {
							f(value)
						}
					}
				default:
					f(data.Value)

				}
			}

			continue
		}
	}

	logging <- datamodels.MessageLogging{
		MsgData: "",
		MsgType: "STOP TEST",
	}
}

var _ = Describe("Testhandlerobservables", Ordered, func() {
	var (
		fileByte []byte
		fileErr  error

		chanStoppedCounting  chan struct{}
		chanOutputDecodeJson chan datamodels.ChanOutputDecodeJSON
		logging              chan datamodels.MessageLogging
		counting             chan datamodels.DataCounterSettings

		//******************* Вспомогательный объект для 'observables' **********************
		so                     = datamodels.NewSupportiveObservables()
		listHandlerObservables = listhandlerjson.NewListHandlerObservablesElement(so)
	)

	BeforeAll(func() {
		chanStoppedCounting = make(chan struct{})
		chanOutputDecodeJson = make(chan datamodels.ChanOutputDecodeJSON)
		logging = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)

		fileByte, fileErr = supportingfunctions.ReadFileJson("test/filestest", "event_1.json")

		go func() {
			fmt.Println("___ Logging START")
			defer fmt.Println("___ Logging STOP")

			for log := range logging {
				if log.MsgType == "error" {
					fmt.Println("----", log, "----")
				}

				if log.MsgType == "STOP TEST" {
					chanStoppedCounting <- struct{}{}

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
		go handlerCaseToTmpObservables(chanOutputDecodeJson, listHandlerObservables, logging)
	})

	Context("Тест 1. Проверка успешности чтения Json файла", func() {
		It("При чтении файла не должно быть ошибок", func() {
			Expect(fileErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Формирование вспомогательного списка 'observables'", func() {
		It("Вспомогательный объект должен быть успешно заполнен", func() {
			<-chanStoppedCounting

			observables := datamodels.NewObservablesMessageTheHive()
			observables.SetObservables(so.GetObservables())

			//тут нужно инициировать списки ниже перечисленных объектов STIX
			//а потом заполнять их основаваясь на dataType

			//fmt.Println(observables.ToStringBeautiful(0))
			for _, v := range observables.GetObservables() {
				fmt.Println("dataType:", v.GetDataType(), " data:", v.GetData())

				switch v.DataType {
				case "whois":
				//LocationDomainObjectsSTIX

				case "domain":
				//DomainNameCyberObservableObjectSTIX

				case "url":
				//URLCyberObservableObjectSTIX

				case "snort_sid":
				//IndicatorDomainObjectsSTIX где есть поля Pattern, PatternVersion, PatternType

				case "file":
				//FileCyberObservableObjectSTIX

				case "filename":
				//FileCyberObservableObjectSTIX

				case "email":
				//EmailAddressCyberObservableObjectSTIX or EmailMessageCyberObservableObjectSTIX

				case "method":
				//http

				case "community":

				case "yara":
				//yara rules
				//IndicatorDomainObjectsSTIX где есть поля Pattern, PatternVersion, PatternType

				case "ip":
				//IPv4AddressCyberObservableObjectSTIX

				case "ip_home":
				//IPv4AddressCyberObservableObjectSTIX

				case "phone-number":
					//IdentityDomainObjectsSTIX

					//case "hash":
				}
			}

			//теперь надо распределить данные по типам STIX объектов
			//основываясь на dataType observables, причем объектов одного и
			//того же типа может быть много, по этому типы STIX объектов
			//должны быть в виде списка

			Expect(len(observables.GetObservables())).Should(Equal(9))
		})
	})
})
