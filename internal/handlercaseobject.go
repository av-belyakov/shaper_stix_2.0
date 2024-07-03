package internal

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/av-belyakov/shaper_stix_2.1/databaseapi/mongodbapi"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal/createrstixobject"
	listhandlerjson "github.com/av-belyakov/shaper_stix_2.1/internal/listhandlerjson"
	wrappers "github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix"
	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

// NewHandlerCaseObject формирует новый обработчик объекта типа Case
func NewHandlerCaseObject(
	input <-chan datamodels.ChanOutputDecodeJSON,
	procRules *ProcessingRules,
	mdbModule *mongodbapi.MongoDBModule,
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

		//*************** Обработчик формирующий объект 'report' в обёртке ****************
		listWrapReport = listhandlerjson.NewHandlerReportDomainObjectSTIX(reportWrap)

		//**************** Вспомогательный, временный объект для 'observables' **********************
		//*** В этом объекте свойства хранят как список уже собранных 'observables', так и
		//*** объект 'observableTmp' в который собираются поступающие а NewHandlerCaseObject
		//данные. После заполнения объекта 'observableTmp' данные из него переносятся в основной
		//список объектов 'observable'
		so                     = datamodels.NewSupportiveObservables()
		listHandlerObservables = listhandlerjson.NewListHandlerObservablesElement(so)
	)

	reportWrap.SetValueID(uuid.New().String())
	identityOwner.SetAnyName(uuid.New().String())
	identitySource.SetAnyID(uuid.New().String())
	identityOrganization.SetAnyID(uuid.New().String())

	//***************************************
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// 			забыл сделать обработку правил которые содержатся в procRules
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//***************************************
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

		//*********************************************************************
		//******* формирование вспомогательного объекта содержащего ***********
		//***** данные из свойства 'observables' объекта 'case' TheHive  ******
		// для всех полей входящих в observables, кроме содержимого поля reports
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

	//*********************************************************************************
	//******* Вспомогательный, временный объект где хранятся уже сформированные *******
	//***** и обработанные объекты типа 'observables', соответствующие объектам *******
	//*** получаемым с TheHive. Объекты хранятся в виде списка объектов			*******
	observables := datamodels.NewObservablesMessageTheHive()
	observables.SetObservables(so.GetObservables())

	var newObject datamodels.GetterCommonPropertiesObjectSTIX

	listObjectSTIX := []datamodels.GetterCommonPropertiesObjectSTIX(nil)
	listRefObjectId := []stixhelpers.IdentifierTypeSTIX(nil)

	for _, v := range observables.GetObservables() {
		fmt.Printf("func 'NewHandlerCaseObject' dataType: '%s' - data: '%s'\n", v.GetDataType(), v.GetData())

		switch v.DataType {
		case "whois":
			newObject = createrstixobject.CreateLocationDomainObjectsSTIX(v)

		case "domain":
			newObject = createrstixobject.CreateDomainNameCyberObservableObjectSTIX(v)

		case "url", "url_pcap":
			newObject = createrstixobject.CreateURLCyberObservableObjectSTIX(v)

		case "snort_sid":
			newObject = createrstixobject.CreateIndicatorSnortIdDomainObjectsSTIX(v)

		case "yara":
			newObject = createrstixobject.CreateIndicatorYaraDomainObjectsSTIX(v)

		case "file", "filename":
			newObject = createrstixobject.CreateFileCyberObservableObjectSTIX(v)

		case "mail", "email":
			newObject = createrstixobject.CreateEmailAddressCyberObservableObjectSTIX(v)

		case "ip":
			newObject = createrstixobject.CreateIPv4AddressCyberObservableObjectSTIX(v)

		case "ip_home":
			newObject = createrstixobject.CreateIPv4AddressCyberObservableObjectSTIX(v)

		case "phone-number":
			newObject = createrstixobject.CreateIdentityDomainObjectsSTIX(v)

		case "hash":
			newObject = createrstixobject.CreateIndicatorHashDomainObjectsSTIX(v)

		case "user-agent":
			newObject = createrstixobject.CreateIndicatorUserAgentDomainObjectsSTIX(v)

		}

		if newObject != nil {
			listRefObjectId = append(listRefObjectId, stixhelpers.IdentifierTypeSTIX(newObject.GetID()))

			//создаем объект Relationship для установки обратной связи между
			//объектом Report и обрабатываемым объектом
			relationship := methodstixobjects.NewRelationshipObjectSTIX()
			relationship.SetValueID(fmt.Sprintf("relationship-%s", uuid.NewString()))
			//исходный объект, то есть обрабатываемый в настоящее время
			relationship.SetValueSourceRef(stixhelpers.IdentifierTypeSTIX(newObject.GetID()))
			//целевой объект, то есть объект Report
			relationship.SetValueTargetRef(stixhelpers.IdentifierTypeSTIX(reportWrap.GetID()))

			//добавляем обрабатываемый объект STIX и объект Relationship в хранилище объектов
			listObjectSTIX = append(listObjectSTIX, newObject, relationship)

			newObject = nil
		}
	}

	fmt.Println("func 'NewHandlerCaseObject' добавляем id обрабатываемого объекта в Report Domain Object STIX")

	//добавляем id обрабатываемого объекта в Report Domain Object STIX
	reportWrap.SetValueObjectRefs(listRefObjectId)

	fmt.Println("func 'NewHandlerCaseObject' передача данных в MongoDB")

	//передача данных в MongoDB
	mdbModule.SendingDataToModule(mongodbapi.ChanInput{
		CommonChan: mongodbapi.CommonChan{
			Section: "data insert",
			Command: "insert",
		},
		Data: listObjectSTIX,
	})
}
