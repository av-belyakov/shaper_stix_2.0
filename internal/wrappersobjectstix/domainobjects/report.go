package domainobjects

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/domainobjectsstix"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix"
)

// NewWrapperReportDomainObjectsSTIX формирует новый объект 'report' с расширеными
// свойствами выходящими за пределы спецификации STIX 2.1.
func NewWrapperReportDomainObjectsSTIX() *WrapperReport {
	return &WrapperReport{
		methodstixobjects.NewReportDomainObjectsSTIX(),
		ReportOutsideSpecification{},
	}
}

func (e *WrapperReport) Get() *WrapperReport {
	return e
}

func (e *WrapperReport) GetObject() interface{} {
	return e
}

// GetReportDomainObject возвращает значения ReportDomainObjectsSTIX
func (e *WrapperReport) GetReportDomainObject() *domainobjectsstix.ReportDomainObjectsSTIX {
	return e.ReportDomainObjectsSTIX
}

// GetReportDomainObject возвращает значения ReportOutsideSpecification
func (e *WrapperReport) GetReportOutsideSpecification() *ReportOutsideSpecification {
	return &e.ReportOutsideSpecification
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e WrapperReport) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	str.WriteString(e.ReportDomainObjectsSTIX.ToStringBeautiful(num))
	str.WriteString(e.ReportOutsideSpecification.ToStringBeautiful(num))

	return str.String()
}

func (e *WrapperReport) MarshalBSON() ([]byte, error) {
	fro := FinalyReportObject{
		CommonPropertiesObjectSTIX:       e.CommonPropertiesObjectSTIX,
		CommonPropertiesDomainObjectSTIX: wrappersobjectstix.NewCommonPropertiesDomainObjectSTIX(),
		Name:                             e.Name,
		Description:                      e.Description,
		ReportTypes:                      e.ReportTypes,
		ObjectRefs:                       e.ObjectRefs,
		ReportOutsideSpecification:       e.ReportOutsideSpecification,
	}

	fro.Revoked = e.Revoked
	fro.Defanged = e.Defanged
	fro.Сonfidence = e.Сonfidence
	fro.Lang = e.Lang
	fro.SpecVersion = e.SpecVersion
	fro.Labels = e.Labels
	fro.Extensions = e.Extensions
	fro.CreatedByRef = e.CreatedByRef
	fro.ExternalReferences = e.ExternalReferences
	fro.ObjectMarkingRefs = e.ObjectMarkingRefs
	fro.GranularMarkings = e.GranularMarkings

	if published, err := time.Parse(time.RFC3339, e.Published); err == nil {
		fro.Published = published
	}

	if create, err := time.Parse(time.RFC3339, e.Created); err == nil {
		fro.Created = create
	}

	if modified, err := time.Parse(time.RFC3339, e.Modified); err == nil {
		fro.Created = modified
	}

	return bson.Marshal(fro)
}

func (e *ReportOutsideSpecification) Get() *ReportOutsideSpecification {
	return e
}

func (e *ReportOutsideSpecification) GetObjectType() string {
	return e.ObjectType
}

// SetValueObjectType устанавливает значение для поля ObjectType
func (e *ReportOutsideSpecification) SetValueObjectType(v string) {
	e.ObjectType = v
}

// SetAnyObjectType устанавливает ЛЮБОЕ значение для поля ObjectType
func (e *ReportOutsideSpecification) SetAnyObjectType(i interface{}) {
	e.ObjectType = fmt.Sprint(i)
}

func (e *ReportOutsideSpecification) GetCaseId() string {
	return e.CaseId
}

// SetValueCaseId устанавливает значение для поля CaseId
func (e *ReportOutsideSpecification) SetValueCaseId(v string) {
	e.CaseId = v
}

// SetAnyCaseId устанавливает ЛЮБОЕ значение для поля CaseId
func (e *ReportOutsideSpecification) SetAnyCaseId(i interface{}) {
	e.CaseId = fmt.Sprint(i)
}

func (e *ReportOutsideSpecification) GetRootId() string {
	return e.RootId
}

// SetValueRootId устанавливает значение для поля RootId
func (e *ReportOutsideSpecification) SetValueRootId(v string) {
	e.RootId = v
}

// SetAnyRootId устанавливает ЛЮБОЕ значение для поля RootId
func (e *ReportOutsideSpecification) SetAnyRootId(i interface{}) {
	e.RootId = fmt.Sprint(i)
}

func (e *ReportOutsideSpecification) GetObjectId() string {
	return e.ObjectId
}

// SetValueObjectId устанавливает значение для поля ObjectId
func (e *ReportOutsideSpecification) SetValueObjectId(v string) {
	e.ObjectId = v
}

// SetAnyObjectId устанавливает ЛЮБОЕ значение для поля ObjectId
func (e *ReportOutsideSpecification) SetAnyObjectId(i interface{}) {
	e.ObjectId = fmt.Sprint(i)
}

func (e *ReportOutsideSpecification) GetStartDate() string {
	return e.StartDate
}

// SetValueStartDate устанавливает значение в формате RFC3339 для поля StartDate
func (e *ReportOutsideSpecification) SetValueStartDate(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.StartDate = v

	return nil
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (e *ReportOutsideSpecification) SetAnyStartDate(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.StartDate = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *ReportOutsideSpecification) GetEndDate() string {
	return e.EndDate
}

// SetValueEndDate устанавливает значение в формате RFC3339 для поля EndDate
func (e *ReportOutsideSpecification) SetValueEndDate(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.EndDate = v

	return nil
}

// SetAnyEndDate устанавливает ЛЮБОЕ значение для поля EndDate
func (e *ReportOutsideSpecification) SetAnyEndDate(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.EndDate = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *ReportOutsideSpecification) GetImpactStatus() string {
	return e.ImpactStatus
}

// SetValueImpactStatus устанавливает значение для поля ImpactStatus
func (e *ReportOutsideSpecification) SetValueImpactStatus(v string) {
	e.ImpactStatus = v
}

// SetAnyImpactStatus устанавливает ЛЮБОЕ значение для поля ImpactStatus
func (e *ReportOutsideSpecification) SetAnyImpactStatus(i interface{}) {
	e.ImpactStatus = fmt.Sprint(i)
}

func (e *ReportOutsideSpecification) GetResolutionStatus() string {
	return e.ResolutionStatus
}

// SetValueResolutionStatus устанавливает значение для поля ResolutionStatus
func (e *ReportOutsideSpecification) SetValueResolutionStatus(v string) {
	e.ResolutionStatus = v
}

// SetAnyResolutionStatus устанавливает ЛЮБОЕ значение для поля ResolutionStatus
func (e *ReportOutsideSpecification) SetAnyResolutionStatus(i interface{}) {
	e.ResolutionStatus = fmt.Sprint(i)
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e ReportOutsideSpecification) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'object_type': '%s'\n", ws, e.ObjectType))
	str.WriteString(fmt.Sprintf("%s'root_id': '%s'\n", ws, e.RootId))
	str.WriteString(fmt.Sprintf("%s'object_id': '%v'\n", ws, e.ObjectId))
	str.WriteString(fmt.Sprintf("%s'case_id': '%s'\n", ws, e.CaseId))
	str.WriteString(fmt.Sprintf("%s'start_date': '%s'\n", ws, e.StartDate))
	str.WriteString(fmt.Sprintf("%s'end_date': '%v'\n", ws, e.EndDate))
	str.WriteString(fmt.Sprintf("%s'impact_status': '%s'\n", ws, e.ImpactStatus))
	str.WriteString(fmt.Sprintf("%s'resolution_status': '%v'\n", ws, e.ResolutionStatus))

	return str.String()
}
