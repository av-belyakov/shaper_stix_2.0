package internal

import (
	"github.com/google/uuid"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	listhandlerjson "github.com/av-belyakov/shaper_stix_2.1/internal/listhandlerjson"
	wrappers "github.com/av-belyakov/shaper_stix_2.1/internal/wrappersObjectSTIX"
)

// NewHandlerCaseObject формирует новый обработчик объекта типа Case
func NewHandlerCaseObject(
	input chan datamodels.ChanOutputDecodeJSON,
	procRules *ProcessingRules,
	couting chan<- datamodels.DataCounterSettings,
	logging chan<- datamodels.MessageLogging) {

	var (
		//формируем новый объект 'report' в обёртке
		reportWrap *wrappers.WrapperReport = wrappers.NewWrapperReportDomainObjectsSTIX()
		//формируем объект для хранения значения свойства 'source'
		identitySource = methodstixobjects.NewIdentityDomainObjectsSTIX()
		//формируем объект для хранения значения свойства 'event.organization'
		identityOrganization = methodstixobjects.NewIdentityDomainObjectsSTIX()
		//формируем объект для хранения значения свойства 'event.object.owner'
		identityOwner = methodstixobjects.NewIdentityDomainObjectsSTIX()
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

		//формируем объекты на основе разного типа из метода 'observables.dataType'
		switch data.FieldBranch {
		case "whois":

		case "url":
		}

	}
}
