package internal

import (
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	listhandlerjson "github.com/av-belyakov/shaper_stix_2.1/internal/listhandlerjson"
	wrappers "github.com/av-belyakov/shaper_stix_2.1/internal/wrappersObjectSTIX"
)

func NewVerifiedObject(
	input chan datamodels.ChanOutputDecodeJSON,
	procRules *ProcessingRules,
	couting chan<- datamodels.DataCounterSettings,
	logging chan<- datamodels.MessageLogging) {

	var (
		//формируем новый объект 'report' в обёртке
		reportWrap *wrappers.WrapperReport = wrappers.NewWrapperReportDomainObjectsSTIX()
	)

	//обработчик формирующий объект 'report' в обёртке
	listWrapReport := listhandlerjson.NewHandlerReportDomainObjectSTIX(reportWrap)

	/*
		Необходимо как минимум сохранять следующие свойства объекта TheHiveCase:

		эти свойства нужно складывать в другие объекты:
		 source
		 event.organization
		 event.organizationId
		 event.object.owner

		теги связанные с geoip

		еще надо решить вопрос с группировкой путем формирования объекта 'grouping'
	*/

	for data := range input {
		/*
			тут надо формировать новые объекты STIX
			и проверять принимаемые данные на соответствие
			правилам
		*/
	}
}
