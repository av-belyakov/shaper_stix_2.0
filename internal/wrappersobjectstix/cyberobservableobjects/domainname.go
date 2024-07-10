package cyberobservableobjects

import (
	"go.mongodb.org/mongo-driver/bson"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix"
)

// NewWrapperDomainNameCyberObservableObjectSTIX формирует новый объект 'location'
func NewWrapperDomainNameCyberObservableObjectSTIX() *WrapperDomainName {
	return &WrapperDomainName{
		methodstixobjects.NewDomainNameCyberObservableObjectSTIX(),
		wrappersobjectstix.CommonOutsideSpecification{},
	}
}

func (e *WrapperDomainName) Get() *WrapperDomainName {
	return e
}

func (e *WrapperDomainName) GetObject() interface{} {
	return e
}

func (e *WrapperDomainName) ToStringBeautiful(num int) string {
	return e.DomainNameCyberObservableObjectSTIX.ToStringBeautiful(num)
}

func (e *WrapperDomainName) MarshalBSON() ([]byte, error) {
	fdno := FinalyDomainNameObjects{
		CommonPropertiesObjectSTIX:                        e.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: e.OptionalCommonPropertiesCyberObservableObjectSTIX,
		Value:          e.Value,
		ResolvesToRefs: e.ResolvesToRefs,
	}

	return bson.Marshal(fdno)
}
