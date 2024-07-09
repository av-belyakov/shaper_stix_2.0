package domainobjects

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix"
)

// NewWrapperIdentityDomainObjectsSTIX формирует новый объект 'identity'
func NewWrapperIdentityDomainObjectsSTIX() *WrapperIdentity {
	return &WrapperIdentity{
		methodstixobjects.NewIdentityDomainObjectsSTIX(),
	}
}

func (e *WrapperIdentity) Get() *WrapperIdentity {
	return e
}

func (e *WrapperIdentity) ToStringBeautiful(num int) string {
	return e.IdentityDomainObjectsSTIX.ToStringBeautiful(num)
}

func (e *WrapperIdentity) MarshalBSON() ([]byte, error) {
	fio := FinalyIdentityObjects{
		CommonPropertiesObjectSTIX:       e.CommonPropertiesObjectSTIX,
		CommonPropertiesDomainObjectSTIX: wrappersobjectstix.NewCommonPropertiesDomainObjectSTIX(),
		Name:                             e.Name,
		Description:                      e.Description,
		ContactInformation:               e.ContactInformation,
		Roles:                            e.Roles,
		IdentityClass:                    e.IdentityClass,
		Sectors:                          e.Sectors,
	}

	if created, err := time.Parse(time.RFC3339, e.GetCreated()); err == nil {
		fio.Created = created
	}

	if modified, err := time.Parse(time.RFC3339, e.GetModified()); err == nil {
		fio.Modified = modified
	}

	return bson.Marshal(fio)
}
