package cyberobservableobjects

import (
	"go.mongodb.org/mongo-driver/bson"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
)

// NewWrapperEmailAddressCyberObservableObjectSTIX формирует новый объект 'email-address'
func NewWrapperEmailAddressCyberObservableObjectSTIX() *WrapperEmailAddress {
	return &WrapperEmailAddress{
		methodstixobjects.NewEmailAddressCyberObservableObjectSTIX(),
	}
}

func (e *WrapperEmailAddress) Get() *WrapperEmailAddress {
	return e
}

func (e *WrapperEmailAddress) ToStringBeautiful(num int) string {
	return e.EmailAddressCyberObservableObjectSTIX.ToStringBeautiful(num)
}

func (e *WrapperEmailAddress) MarshalBSON() ([]byte, error) {
	feao := FinalyEmailAddressObjects{
		CommonPropertiesObjectSTIX:                        e.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: e.OptionalCommonPropertiesCyberObservableObjectSTIX,
		Value:        e.Value,
		DisplayName:  e.DisplayName,
		BelongsToRef: e.BelongsToRef,
	}

	return bson.Marshal(feao)
}
