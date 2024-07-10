package wrappersobjectstix

import (
	"fmt"

	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

func (e *CommonOutsideSpecification) Get() *CommonOutsideSpecification {
	return e
}

// GetTlp возвращает номер TLP
func (e *CommonOutsideSpecification) GetTlp() int {
	return e.Tlp
}

// SetAnyTlp устанавливает номер TLP
func (e *CommonOutsideSpecification) SetAnyTlp(i interface{}) {
	e.Tlp = supportingfunctions.ConversionAnyToInt(i)
}

// SetValueTlp устанавливает номер TLP
func (e *CommonOutsideSpecification) SetValueTlp(v int) {
	e.Tlp = v
}

// GetElementId возвращает идентификатор ElementId
func (e *CommonOutsideSpecification) GetElementId() string {
	return e.ElementId
}

// SetAnyElementId устанавливает идентификатор ElementId
func (e *CommonOutsideSpecification) SetAnyElementId(i interface{}) {
	e.ElementId = fmt.Sprint(i)
}

// SetValueElementId устанавливает идентификатор ElementId
func (e *CommonOutsideSpecification) SetValueElementId(v string) {
	e.ElementId = v
}
