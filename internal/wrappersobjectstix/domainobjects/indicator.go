package domainobjects

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix"
)

// NewWrapperIndicatorDomainObjectsSTIX формирует новый объект 'indicator'
func NewWrapperIndicatorDomainObjectsSTIX() *WrapperIndicator {
	return &WrapperIndicator{
		methodstixobjects.NewIndicatorDomainObjectsSTIX(),
	}
}

func (e *WrapperIndicator) Get() *WrapperIndicator {
	return e
}

func (e *WrapperIndicator) ToStringBeautiful(num int) string {
	return e.IndicatorDomainObjectsSTIX.ToStringBeautiful(num)
}

func (e *WrapperIndicator) MarshalBSON() ([]byte, error) {
	fio := FinalyIndicatorObject{
		CommonPropertiesObjectSTIX:       e.CommonPropertiesObjectSTIX,
		CommonPropertiesDomainObjectSTIX: wrappersobjectstix.NewCommonPropertiesDomainObjectSTIX(),
		Name:                             e.Name,
		Description:                      e.Description,
		Pattern:                          e.Pattern,
		PatternVersion:                   e.PatternVersion,
		PatternType:                      e.PatternType,
		KillChainPhases:                  e.KillChainPhases,
		IndicatorTypes:                   e.IndicatorTypes,
	}

	if created, err := time.Parse(time.RFC3339, e.GetCreated()); err == nil {
		fio.Created = created
	}

	if modified, err := time.Parse(time.RFC3339, e.GetModified()); err == nil {
		fio.Modified = modified
	}

	if validFrom, err := time.Parse(time.RFC3339, e.GetValidFrom()); err == nil {
		fio.ValidFrom = validFrom
	}

	if validFrom, err := time.Parse(time.RFC3339, e.GetValidFrom()); err == nil {
		fio.ValidFrom = validFrom
	}

	return bson.Marshal(fio)
}
