package listhandlerjson

import "github.com/av-belyakov/methodstixobjects/datamodels/domainobjectsstix"

func NewHandlerReportDomainObjectSTIX(elem *domainobjectsstix.ReportDomainObjectsSTIX) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		"event.rootId":             {elem.SetAnyID},
		"event.object.createdAt":   {elem.SetAnyCreated},
		"event.object.updatedAt":   {elem.SetAnyModified},
		"event.object.title":       {elem.SetAnyName},
		"event.object.description": {elem.SetAnyDescription},
	}
}

/*
Необходимо как минимум сохранять следующие свойства объекта TheHiveCase:

source
event.organization
event.organizationId
event.objectType
event.object.caseId
event.object.startDate
event.object.endDate
event.object.impactStatus
event.object.resolutionStatus
event.object.owner

теги связанные с geoip
*/
