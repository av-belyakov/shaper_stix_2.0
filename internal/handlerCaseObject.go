package internal

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	listhandlerjson "github.com/av-belyakov/shaper_stix_2.1/internal/listhandlerjson"
	wrappers "github.com/av-belyakov/shaper_stix_2.1/internal/wrappersObjectSTIX"
	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

// NewHandlerCaseObject формирует новый обработчик объекта типа Case
func NewHandlerCaseObject(
	input chan datamodels.ChanOutputDecodeJSON,
	procRules *ProcessingRules,
	couting chan<- datamodels.DataCounterSettings,
	logging chan<- datamodels.MessageLogging) {

	var (
		rootId string

		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		//формируем новый объект 'report' в обёртке
		reportWrap *wrappers.WrapperReport = wrappers.NewWrapperReportDomainObjectsSTIX()
		//формируем объект для хранения значения свойства 'source'
		identitySource = methodstixobjects.NewIdentityDomainObjectsSTIX()
		//формируем объект для хранения значения свойства 'event.organization'
		identityOrganization = methodstixobjects.NewIdentityDomainObjectsSTIX()
		//формируем объект для хранения значения свойства 'event.object.owner'
		identityOwner = methodstixobjects.NewIdentityDomainObjectsSTIX()

		//******************* Вспомогательный объект для Observables **********************
		so                     = datamodels.NewSupportiveObservables()
		listHandlerObservables = listhandlerjson.NewListHandlerObservablesElement(so)
	)

	reportWrap.SetValueID(uuid.New().String())
	identityOwner.SetAnyName(uuid.New().String())
	identitySource.SetAnyID(uuid.New().String())
	identityOrganization.SetAnyID(uuid.New().String())

	//обработчик формирующий объект 'report' в обёртке
	listWrapReport := listhandlerjson.NewHandlerReportDomainObjectSTIX(reportWrap)

	/*
		Необходимо как минимум сохранять следующие свойства объекта TheHiveCase:

		эти свойства нужно складывать в другие объекты:
		 event.object.owner

		теги связанные с geoip

		еще надо решить вопрос с группировкой путем формирования объекта 'grouping'
	*/

	for data := range input {
		var handlerIsExist bool

		//добавляем информацию об источнике из свойства 'source'
		if data.FieldBranch == "source" {
			identitySource.SetAnyName(data.Value)
		}

		//добавляем информацию об организации из свойства 'event.organization'
		if data.FieldBranch == "event.organization" {
			identityOrganization.SetAnyName(data.Value)
		}

		//добавляем информацию об организации из свойства 'event.object.owner'
		if data.FieldBranch == "event.object.owner" {
			identityOwner.SetAnyName(data.Value)
		}
		//добавляем информацию об организации из свойства 'event.object.updateAt'
		if data.FieldBranch == "event.object.updateAt" {
			identityOwner.SetAnyModified(data.Value)
		}

		//******************************************************************************
		//******** формирование обьекта относящегося к Report Domain Object STIX *******
		if reports, ok := listWrapReport[data.FieldBranch]; ok {
			for _, f := range reports {
				f(data.Value)
			}
		}

		//************************************************************************
		//********** Сбор всех объектов относящихся к полю Observables  **********
		// для всех полей входящих в observables, кроме содержимого
		//поля reports
		if lf, ok := listHandlerObservables[data.FieldBranch]; ok {
			handlerIsExist = true

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

		//для всех полей входящих в состав observables.reports
		if strings.Contains(data.FieldBranch, "observables.reports.") {
			handlerIsExist = true
			so.HandlerReportValue(data.FieldBranch, data.Value)
		}

		//формируем объекты на основе разного типа из метода 'observables.dataType'
		switch data.FieldBranch {
		// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		// тут надо сделать список временных объектов в которые будут укладыватся
		// все поля объекта observables в независимости от типа, а затем перебирать
		// этот временный объект и на основе 'observables.dataType' подбирать для
		// этих данных необходимый объект STIX
		//
		// может быть несколько объектов одного типа как например объект domain
		// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		case "whois":
			//LocationDomainObjectsSTIX

		case "domain":
			//DomainNameCyberObservableObjectSTIX
			/*
							"domain" : [
				              {
				                "ioc" : true,
				                "tlp" : 2,
				                "_createdAt" : "2024-03-20T14:14:10+03:00",
				                "_updatedAt" : "2024-03-20T14:19:26+03:00",
				                "startDate" : "2024-03-20T14:14:10+03:00",
				                "_createdBy" : "i.monahov@cloud.gcm",
				                "_updatedBy" : "i.monahov@cloud.gcm",
				                "_id" : "~87845617776",
				                "_type" : "Observable",
				                "data" : "collect[.]clickandanalytics[.]com",
				                "dataType" : "domain",
				                "message" : ".",
				                "tags" : {
				                  "misp:network activity" : [
				                    "\"domain\""
				                  ]
				                },
				                "tagsAll" : [
				                  "misp:Network activity=\"domain\""
				                ],
				                "attachment" : { }
				              },
				              {
				                "ioc" : true,
				                "tlp" : 2,
				                "_createdAt" : "2024-03-20T14:14:10+03:00",
				                "_updatedAt" : "2024-03-20T14:19:26+03:00",
				                "startDate" : "2024-03-20T14:14:10+03:00",
				                "_createdBy" : "i.monahov@cloud.gcm",
				                "_updatedBy" : "i.monahov@cloud.gcm",
				                "_id" : "~87845625968",
				                "_type" : "Observable",
				                "data" : "gate[.]getmygateway[.]com",
				                "dataType" : "domain",
				                "message" : ".",
				                "tags" : {
				                  "misp:network activity" : [
				                    "\"domain\""
				                  ]
				                },
				                "tagsAll" : [
				                  "misp:Network activity=\"domain\""
				                ],
				                "attachment" : { }
				              },
			*/

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

		case "hash":

		}

		//****************************************************************************
		//****************** обработка и логирование списка полей ********************
		//*********** которые не были обработаны сформированными обработчиками *******
		if !handlerIsExist {
			// записываем в лог-файл поля, которые не были обработаны
			listRawFields[data.FieldBranch] = fmt.Sprint(data.Value)
		}

		// отправляем список полей которые не смогли обработать
		if len(listRawFields) > 0 {
			logging <- datamodels.MessageLogging{
				MsgData: supportingfunctions.JoinRawFieldsToString(listRawFields, "rootId", rootId),
				MsgType: "case_raw_fields",
			}
		}
	}
}
