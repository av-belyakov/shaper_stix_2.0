package cyberobservableobjects

import (
	"go.mongodb.org/mongo-driver/bson"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix"
)

// NewWrapperIPv4AddressCyberObservableObjectSTIX формирует новый объект 'ipv4-address'
func NewWrapperIPv4AddressCyberObservableObjectSTIX() *WrapperIPv4Address {
	return &WrapperIPv4Address{
		methodstixobjects.NewIPv4AddressCyberObservableObjectSTIX(),
		wrappersobjectstix.CommonOutsideSpecification{},
	}
}

func (e *WrapperIPv4Address) Get() *WrapperIPv4Address {
	return e
}

func (e *WrapperIPv4Address) GetObject() interface{} {
	return e
}

func (e *WrapperIPv4Address) ToStringBeautiful(num int) string {
	return e.IPv4AddressCyberObservableObjectSTIX.ToStringBeautiful(num)
}

func (e *WrapperIPv4Address) MarshalBSON() ([]byte, error) {
	fipo := FinalyIPv4AddressObjects{
		CommonPropertiesObjectSTIX:                        e.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: e.OptionalCommonPropertiesCyberObservableObjectSTIX,
		Value:          e.Value,
		ResolvesToRefs: e.ResolvesToRefs,
		BelongsToRefs:  e.BelongsToRefs,
	}

	return bson.Marshal(fipo)
}
