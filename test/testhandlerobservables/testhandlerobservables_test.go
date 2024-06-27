package testhandlerobservables_test

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/cyberobservableobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/domainobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
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

			type GetterID interface {
				GetID() string
			}

			type GetterType interface {
				GetType() string
			}

			type GetterCommonPropertiesObjectSTIX interface {
				GetterID
				GetterType
			}

			//createLocationDomainObjectsSTIX формирует объект 'location'
			createLocationDomainObjectsSTIX := func(observable datamodels.ObservableMessage) *domainobjectsstix.LocationDomainObjectsSTIX {
				location := methodstixobjects.NewLocationDomainObjectsSTIX()
				location.SetValueID(uuid.NewString())
				location.SetValueCountry(observable.Data)
				if len(observable.Tags) > 0 {
					location.SetValueName(observable.Tags[0])
				}
				if observable.Message != "" {
					location.SetValueDescription(observable.Message)
				}

				return location
			}

			//createDomainNameCyberObservableObjectSTIX формирует объект 'domain-name'
			createDomainNameCyberObservableObjectSTIX := func(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.DomainNameCyberObservableObjectSTIX {
				domainName := methodstixobjects.NewDomainNameCyberObservableObjectSTIX()
				domainName.SetValueID(fmt.Sprintf("domain-name-%s", uuid.NewString()))
				domainName.SetValueValue(observable.Data)

				return domainName
			}

			//createURLCyberObservableObjectSTIX формирует объект 'url'
			createURLCyberObservableObjectSTIX := func(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.URLCyberObservableObjectSTIX {
				url := methodstixobjects.NewURLCyberObservableObjectSTIX()
				url.SetValueID(fmt.Sprintf("url-%s", uuid.NewString()))
				url.SetValueValue(observable.Data)

				return url
			}

			//createIndicatorSnortIdDomainObjectsSTIX формирует объект 'indicator'
			createIndicatorSnortIdDomainObjectsSTIX := func(observable datamodels.ObservableMessage) *domainobjectsstix.IndicatorDomainObjectsSTIX {
				indicatior := methodstixobjects.NewIndicatorDomainObjectsSTIX()
				indicatior.SetValueID(fmt.Sprintf("indicator-%s", uuid.NewString()))
				indicatior.SetValueName("snort_sid")
				indicatior.SetValueDescription("list of signatures of the Snort computer attack detection tool")
				indicatior.SetValuePattern(observable.Data)
				indicatior.SetValuePatternType("list of numbers")

				if observable.Message != "" {
					indicatior.SetValueDescription(observable.Message)
				}

				return indicatior
			}

			//createIndicatorYaraDomainObjectsSTIX формирует объект 'indicator'
			createIndicatorYaraDomainObjectsSTIX := func(observable datamodels.ObservableMessage) *domainobjectsstix.IndicatorDomainObjectsSTIX {
				indicatior := methodstixobjects.NewIndicatorDomainObjectsSTIX()
				indicatior.SetValueID(fmt.Sprintf("indicator-%s", uuid.NewString()))
				indicatior.SetValueName("yara")
				indicatior.SetValueDescription("yara rule")
				indicatior.SetValuePattern(observable.Data)
				indicatior.SetValuePatternType("string")

				if observable.Message != "" {
					indicatior.SetValueDescription(observable.Message)
				}

				return indicatior
			}

			//createFileCyberObservableObjectSTIX формирует объект 'file'
			createFileCyberObservableObjectSTIX := func(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.FileCyberObservableObjectSTIX {
				file := methodstixobjects.NewFileCyberObservableObjectSTIX()
				file.SetValueID(fmt.Sprintf("file-%s", uuid.NewString()))
				file.SetValueName(observable.Data)

				if observable.Attachment.Name != "" {
					file.SetValueName(observable.Attachment.Name)
				}
				if observable.Attachment.Size > 0 {
					file.SetValueSize(observable.Attachment.Size)
				}
				if len(observable.Attachment.Hashes) > 0 {
					hashes := stixhelpers.HashesTypeSTIX{}

					for k, v := range observable.Attachment.Hashes {
						hashes[fmt.Sprintf("hash_%d", k)] = v
					}

					file.SetValueHashes(hashes)
				}

				return file
			}

			// createEmailAddressCyberObservableObjectSTIX формирует объект 'email-addr'
			createEmailAddressCyberObservableObjectSTIX := func(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.EmailAddressCyberObservableObjectSTIX {
				email := methodstixobjects.NewEmailAddressCyberObservableObjectSTIX()
				email.SetValueID(fmt.Sprintf("email-addr-%s", uuid.NewString()))
				email.SetValueValue(observable.Data)

				return email
			}

			// createIPv4AddressCyberObservableObjectSTIX формирует объект 'ipv4-addr'
			createIPv4AddressCyberObservableObjectSTIX := func(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.IPv4AddressCyberObservableObjectSTIX {
				ipv4 := methodstixobjects.NewIPv4AddressCyberObservableObjectSTIX()
				ipv4.SetValueID(fmt.Sprintf("ipv4-addr-%s", uuid.NewString()))
				ipv4.SetValueValue(observable.Data)

				return ipv4
			}

			//createIdentityDomainObjectsSTIX формирует объект 'identity'
			createIdentityDomainObjectsSTIX := func(observable datamodels.ObservableMessage) *domainobjectsstix.IdentityDomainObjectsSTIX {
				identity := methodstixobjects.NewIdentityDomainObjectsSTIX()
				identity.SetValueID(fmt.Sprintf("identity-%s", uuid.NewString()))
				identity.SetValueName("phone-number")
				identity.SetValueDescription("personal phone number")
				identity.SetValueContactInformation(observable.Data)

				if observable.Message != "" {
					identity.SetValueDescription(observable.Message)
				}

				return identity
			}

			var newobject GetterCommonPropertiesObjectSTIX
			listObjectSTIX := []GetterCommonPropertiesObjectSTIX(nil)

			//fmt.Println(observables.ToStringBeautiful(0))
			for _, v := range observables.GetObservables() {
				fmt.Println("dataType:", v.GetDataType(), " data:", v.GetData())

				switch v.DataType {
				case "whois":
					//LocationDomainObjectsSTIX
					newobject = createLocationDomainObjectsSTIX(v)

				case "domain":
					//DomainNameCyberObservableObjectSTIX
					newobject = createDomainNameCyberObservableObjectSTIX(v)

				case "url":
					//URLCyberObservableObjectSTIX
					newobject = createURLCyberObservableObjectSTIX(v)

				case "snort_sid":
					//IndicatorDomainObjectsSTIX где есть поля Pattern, PatternVersion, PatternType
					newobject = createIndicatorSnortIdDomainObjectsSTIX(v)

				case "yara":
					//IndicatorDomainObjectsSTIX где есть поля Pattern, PatternVersion, PatternType
					newobject = createIndicatorYaraDomainObjectsSTIX(v)

				case "file":
					//FileCyberObservableObjectSTIX
					newobject = createFileCyberObservableObjectSTIX(v)

				case "filename":
					//FileCyberObservableObjectSTIX
					newobject = createFileCyberObservableObjectSTIX(v)

				case "email":
					//EmailAddressCyberObservableObjectSTIX
					newobject = createEmailAddressCyberObservableObjectSTIX(v)

				case "ip":
					//IPv4AddressCyberObservableObjectSTIX
					newobject = createIPv4AddressCyberObservableObjectSTIX(v)

				case "ip_home":
					//IPv4AddressCyberObservableObjectSTIX
					newobject = createIPv4AddressCyberObservableObjectSTIX(v)

				case "phone-number":
					//IdentityDomainObjectsSTIX
					newobject = createIdentityDomainObjectsSTIX(v)

				case "hash":
					//ObservedDataDomainObjectsSTIX
					// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
					// для hash надо подыскать какой то другой объект
					// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

				}

				if newobject != nil {
					listObjectSTIX = append(listObjectSTIX, newobject)
					newobject = nil
				}
			}

			fmt.Println("COUNT:", len(listObjectSTIX))

			Expect(listObjectSTIX).ToNot(BeNil())

			for k, v := range listObjectSTIX {
				//fmt.Println(v)
				fmt.Printf("%d. type:'%s', id:'%s'\n", k, v.GetType(), v.GetID())
			}

			Expect(len(listObjectSTIX)).ShouldNot(Equal(0))
			Expect(len(observables.GetObservables())).Should(Equal(9))
		})
	})
})
