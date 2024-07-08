package listhandlerjson

import do "github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix/domainobjects"

func NewHandlerReportDomainObjectSTIX(elem *do.WrapperReport) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//основные свойства объекта 'report' STIX
		"event.object.createdAt":   {elem.SetAnyCreated},
		"event.object.updatedAt":   {elem.SetAnyModified},
		"event.object.title":       {elem.SetAnyName},
		"event.object.description": {elem.SetAnyDescription},
		//расширеные свойства невходящие в общую спецификацию объекта 'report' STIX
		"event.objectType":              {elem.ReportOutsideSpecification.SetAnyObjectType},
		"event.rootId":                  {elem.ReportOutsideSpecification.SetAnyRootId},
		"event.objectId":                {elem.ReportOutsideSpecification.SetAnyObjectId},
		"event.object.caseId":           {elem.ReportOutsideSpecification.SetAnyCaseId},
		"event.object.startDate":        {elem.ReportOutsideSpecification.SetAnyStartDate},
		"event.object.endDate":          {elem.ReportOutsideSpecification.SetAnyEndDate},
		"event.object.impactStatus":     {elem.ReportOutsideSpecification.SetAnyImpactStatus},
		"event.object.resolutionStatus": {elem.ReportOutsideSpecification.SetAnyResolutionStatus},
	}
}
