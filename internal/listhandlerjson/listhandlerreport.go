package listhandlerjson

import wrappers "github.com/av-belyakov/shaper_stix_2.1/internal/wrappersObjectSTIX"

func NewHandlerReportDomainObjectSTIX(elem *wrappers.WrapperReport) map[string][]func(interface{}) {
	return map[string][]func(interface{}){
		//основные свойства объекта 'report' STIX
		"event.rootId":             {elem.SetAnyID},
		"event.object.createdAt":   {elem.SetAnyCreated},
		"event.object.updatedAt":   {elem.SetAnyModified},
		"event.object.title":       {elem.SetAnyName},
		"event.object.description": {elem.SetAnyDescription},
		//расширеные свойства невходящие в общую спецификацию объекта 'report' STIX
		"event.objectType":              {elem.ReportOutsideSpecification.SetAnyObjectType},
		"event.object.caseId":           {elem.ReportOutsideSpecification.SetAnyCaseId},
		"event.object.startDate":        {elem.ReportOutsideSpecification.SetAnyStartDate},
		"event.object.endDate":          {elem.ReportOutsideSpecification.SetAnyEndDate},
		"event.object.impactStatus":     {elem.ReportOutsideSpecification.SetAnyImpactStatus},
		"event.object.resolutionStatus": {elem.ReportOutsideSpecification.SetAnyResolutionStatus},
	}
}
