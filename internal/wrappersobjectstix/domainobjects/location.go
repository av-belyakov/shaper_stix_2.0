package domainobjects

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix"
)

// NewWrapperLocationDomainObjectsSTIX формирует новый объект 'location'
func NewWrapperLocationDomainObjectsSTIX() *WrapperLocation {
	return &WrapperLocation{
		methodstixobjects.NewLocationDomainObjectsSTIX(),
	}
}

func (e *WrapperLocation) Get() *WrapperLocation {
	return e
}

func (e *WrapperLocation) ToStringBeautiful(num int) string {
	return e.LocationDomainObjectsSTIX.ToStringBeautiful(num)
}

func (e *WrapperLocation) MarshalBSON() ([]byte, error) {
	flo := FinalyLocationObjects{
		CommonPropertiesObjectSTIX:       e.CommonPropertiesObjectSTIX,
		CommonPropertiesDomainObjectSTIX: wrappersobjectstix.NewCommonPropertiesDomainObjectSTIX(),
		Name:                             e.Name,
		Description:                      e.Description,
		Latitude:                         e.Latitude,
		Longitude:                        e.Longitude,
		Precision:                        e.Precision,
		Country:                          e.Country,
		AdministrativeArea:               e.AdministrativeArea,
		City:                             e.City,
		StreetAddress:                    e.StreetAddress,
		PostalCode:                       e.PostalCode,
		Region:                           e.Region,
	}

	if created, err := time.Parse(time.RFC3339, e.GetCreated()); err == nil {
		flo.Created = created
	}

	if modified, err := time.Parse(time.RFC3339, e.GetModified()); err == nil {
		flo.Modified = modified
	}

	return bson.Marshal(flo)
}
