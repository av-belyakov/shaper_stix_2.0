package cyberobservableobjects

import (
	"go.mongodb.org/mongo-driver/bson"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
)

// NewWrapperURLCyberObservableObjectSTIX формирует новый объект 'url'
func NewWrapperURLCyberObservableObjectSTIX() *WrapperURL {
	return &WrapperURL{
		methodstixobjects.NewURLCyberObservableObjectSTIX(),
	}
}

func (e *WrapperURL) Get() *WrapperURL {
	return e
}

func (e *WrapperURL) ToStringBeautiful(num int) string {
	return e.URLCyberObservableObjectSTIX.ToStringBeautiful(num)
}

func (e *WrapperURL) MarshalBSON() ([]byte, error) {
	furlo := FinalyURLObjects{
		CommonPropertiesObjectSTIX:                        e.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: e.OptionalCommonPropertiesCyberObservableObjectSTIX,
		Value: e.Value,
	}

	return bson.Marshal(furlo)
}
