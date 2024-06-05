package wrappersobjectstix

import (
	"fmt"
	"time"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/domainobjectsstix"
	"go.mongodb.org/mongo-driver/bson"
)

// WrapperReport содержит STIX объект 'report' и дополнительные расширеные свойства
type WrapperReport struct {
	*domainobjectsstix.ReportDomainObjectsSTIX
	ReportOutsideSpecification ReportOutsideSpecification
}

// ReportOutsideSpecification содержит дополнительные свойства
type ReportOutsideSpecification struct {
	ObjectType       string
	CaseId           string
	StartDate        string
	EndDate          string
	ImpactStatus     string
	ResolutionStatus string
}

// NewWrapperReportDomainObjectsSTIX формирует новый объект 'report' с расширеными
// свойствами выходящими за пределы спецификации STIX 2.1.
func NewWrapperReportDomainObjectsSTIX() *WrapperReport {
	return &WrapperReport{
		methodstixobjects.NewReportDomainObjectsSTIX(),
		ReportOutsideSpecification{},
	}
}

func (wr *WrapperReport) MarshalBSON(i interface{}) ([]byte, error) {
	/*

		Тут надо сделать обработчик формирующий BSON

	*/

	b, err := bson.Marshal(i)

	return b, err
}

/*
func (wr *WrapperReport) MarshalJSON(i interface{}) ([]byte, error) {
	//
	//
	//   Тут надо сделать обработчик формирующий JSON
	//
	//

	b, err := json.Marshal(i)

	return b, err
}
*/

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
