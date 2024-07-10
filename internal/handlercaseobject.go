package internal

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/google/uuid"

	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/av-belyakov/shaper_stix_2.1/databaseapi/mongodbapi"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal/createrstixobject"
	listhandlerjson "github.com/av-belyakov/shaper_stix_2.1/internal/listhandlerjson"
	do "github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix/domainobjects"
	"github.com/av-belyakov/shaper_stix_2.1/ruleinteraction"
	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

// NewHandlerCaseObject формирует новый обработчик объекта типа Case
func NewHandlerCaseObject(
	input <-chan datamodels.ChanOutputDecodeJSON,
	listRules ruleinteraction.ListRule,
	mdbModule *mongodbapi.MongoDBModule,
	counting chan<- datamodels.DataCounterSettings,
	logging chan<- datamodels.MessageLogging) {

	var (
		caseId float64
		rootId string

		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		//формируем новый объект 'report' в обёртке
		reportWrap *do.WrapperReport = do.NewWrapperReportDomainObjectsSTIX()
		//формируем объект для хранения значения свойства 'source'
		//identitySource = methodstixobjects.NewIdentityDomainObjectsSTIX()
		identitySource = do.NewWrapperIdentityDomainObjectsSTIX()

		//формируем объект для хранения значения свойства 'event.organization'
		//identityOrganization = methodstixobjects.NewIdentityDomainObjectsSTIX()
		identityOrganization = do.NewWrapperIdentityDomainObjectsSTIX()
		//формируем объект для хранения значения свойства 'event.object.owner'
		//identityOwner = methodstixobjects.NewIdentityDomainObjectsSTIX()
		identityOwner = do.NewWrapperIdentityDomainObjectsSTIX()

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

	reportWrap.SetValueID(fmt.Sprintf("report--%s", uuid.NewString()))
	identityOwner.SetAnyID(fmt.Sprintf("identity--%s", uuid.NewString()))
	identitySource.SetAnyID(fmt.Sprintf("identity--%s", uuid.NewString()))
	identityOrganization.SetAnyID(fmt.Sprintf("identity--%s", uuid.NewString()))

	for data := range input {
		var handlerIsExist bool

		//*************** Обработка правил ***************
		//обработка правил REPLACEMENT (замена)
		newValue, _, err := listRules.ReplacementRuleHandler(data.ValueType, data.FieldName, data.Value)
		if err != nil {
			_, f, l, _ := runtime.Caller(0)

			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'search value \"%s\" from rule of section \"REPLACE\" is not fulfilled' %s:%d", data.Value, f, l-1),
				MsgType: "warning",
			}
		}
		//обработка правил PASS (пропуск)
		listRules.PassRuleHandler(data.FieldBranch, newValue)
		//**********************************************

		//ищем id события
		if cid, ok := searchCaseId(data); ok {
			caseId = cid
		}

		//добавляем информацию об источнике из свойства 'source'
		if data.FieldBranch == "source" {
			identitySource.SetAnyName(newValue)
		}

		//добавляем информацию об организации из свойства 'event.organization'
		if data.FieldBranch == "event.organisation" {
			identityOrganization.SetAnyName(newValue)
		}

		//добавляем информацию о лице создавшем объект из свойства 'event.object.owner'
		if data.FieldBranch == "event.object.owner" {
			identityOwner.SetAnyName(newValue)
		}

		if data.FieldBranch == "event.object.updatedAt" {
			identitySource.SetAnyModified(newValue)
			identityOrganization.SetAnyModified(newValue)
			identityOwner.SetAnyModified(newValue)
		}

		if data.FieldBranch == "event.objectId" {
			identitySource.SetAnyElementId(newValue)
			identityOrganization.SetAnyElementId(newValue)
			identityOwner.SetAnyElementId(newValue)
		}

		if data.FieldBranch == "event.object.tlp" {
			identitySource.SetAnyTlp(newValue)
			identityOrganization.SetAnyTlp(newValue)
			identityOwner.SetAnyTlp(newValue)
		}

		//******************************************************************************
		//******** формирование обьекта относящегося к Report Domain Object STIX *******
		if reports, ok := listWrapReport[data.FieldBranch]; ok {
			for _, f := range reports {
				f(newValue)
			}
		}

		//*********************************************************************
		//******* формирование вспомогательного объекта содержащего ***********
		//***** данные из свойства 'observables' объекта 'case' TheHive  ******
		// для всех полей входящих в observables, кроме содержимого поля reports
		if lf, ok := listHandlerObservables[data.FieldBranch]; ok {
			handlerIsExist = true

			for _, f := range lf {
				r := reflect.TypeOf(newValue)
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := newValue.([]interface{}); ok {
						for _, value := range s {
							f(value)
						}
					}
				default:
					f(newValue)

				}
			}

			continue
		}
		//для всех полей входящих в состав observables.reports
		if strings.Contains(data.FieldBranch, "observables.reports.") {
			handlerIsExist = true
			so.HandlerReportValue(data.FieldBranch, newValue)
		}

		//****************************************************************************
		//****************** обработка и логирование списка полей ********************
		//*********** которые не были обработаны сформированными обработчиками *******
		if !handlerIsExist {
			// записываем в лог-файл поля, которые не были обработаны
			listRawFields[data.FieldBranch] = fmt.Sprint(newValue)
		}

		// отправляем список полей которые не смогли обработать
		if len(listRawFields) > 0 {
			logging <- datamodels.MessageLogging{
				MsgData: supportingfunctions.JoinRawFieldsToString(listRawFields, "rootId", rootId),
				MsgType: "case_raw_fields",
			}
		}
	}

	var isAllowed bool
	//проверяем что бы хотя бы одно правило разрешало пропуск кейса
	if listRules.GetRulePassany() || listRules.SomePassRuleIsTrue() {
		isAllowed = true

		//сетчик кейсов соответствующих или не соответствующих правилам
		counting <- datamodels.DataCounterSettings{
			DataType: "update events meet rules",
			Count:    1,
		}
	}

	if !isAllowed {
		// ***********************************
		// Это логирование только для теста!!!
		// ***********************************
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("TEST_INFO func 'NewMispFormat', case with id '%d' does not comply with the rules", int(caseId)),
			MsgType: "testing",
		}
		//
		//

		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'the message with case id %d was not sent to MISP because it does not comply with the rules' %s:%d", int(caseId), f, l-1),
			MsgType: "warning",
		}

		return
	}

	// ***********************************
	// Это логирование только для теста!!!
	// ***********************************
	logging <- datamodels.MessageLogging{
		MsgData: fmt.Sprintf("TEST_INFO func 'NewMispFormat', case with id '%d' equal rules, send data to MISP module", int(caseId)),
		MsgType: "testing",
	}
	//
	//

	//*********************************************************************************
	//******* Вспомогательный, временный объект где хранятся уже сформированные *******
	//***** и обработанные объекты типа 'observables', соответствующие объектам *******
	//*** получаемым с TheHive. Объекты хранятся в виде списка объектов			*******
	observables := datamodels.NewObservablesMessageTheHive()
	observables.SetObservables(so.GetObservables())

	var newObject datamodels.HandlerSTIXObject

	listObjectSTIX := []datamodels.HandlerSTIXObject(nil)
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

		//************************************************************************
		//
		// Все же решил создавать объект relationship в модуле взаимодействия с БД
		//
		//************************************************************************

		if newObject != nil {
			listRefObjectId = append(listRefObjectId, stixhelpers.IdentifierTypeSTIX(newObject.GetID()))

			//добавляем обрабатываемый объект STIX в хранилище объектов
			listObjectSTIX = append(listObjectSTIX, newObject)

			newObject = nil
		}
	}

	//*****************************************************************************************
	//********* добавляем id объекта содержащего информацию о создателе объекта case **********
	listRefObjectId = append(listRefObjectId, stixhelpers.IdentifierTypeSTIX(identityOwner.GetID()))

	//******************************************************************************
	//*** добавляем id объекта содержащего информацию о источнике объекта case ****
	listRefObjectId = append(listRefObjectId, stixhelpers.IdentifierTypeSTIX(identitySource.GetID()))

	//***************************************************************************************
	//**** добавляем id объекта содержащего информацию об организации относящейся к case ****
	listRefObjectId = append(listRefObjectId, stixhelpers.IdentifierTypeSTIX(identityOrganization.GetID()))

	//добавляем id обрабатываемого объекта в Report Domain Object STIX
	reportWrap.SetValueObjectRefs(listRefObjectId)

	//добавляем в список вновь созданых STIX объектов, объекты, которые были созданы
	//при анализе всех других объектов кроме объекта 'observables'
	listObjectSTIX = append(listObjectSTIX, reportWrap)
	listObjectSTIX = append(listObjectSTIX, identityOwner, identitySource, identityOrganization)

	//передача данных в MongoDB
	mdbModule.SendingDataToModule(mongodbapi.ChanInput{
		CommonChan: mongodbapi.CommonChan{
			Section: "data insert",
			Command: "insert",
		},
		Data: listObjectSTIX,
	})
}
