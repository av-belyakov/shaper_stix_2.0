package createrstixobject

import (
	"fmt"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/domainobjectsstix"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/google/uuid"
)

// CreateLocationDomainObjectsSTIX формирует объект 'location'
func CreateLocationDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjectsstix.LocationDomainObjectsSTIX {
	location := methodstixobjects.NewLocationDomainObjectsSTIX()
	location.SetValueID(fmt.Sprintf("location-%s", uuid.NewString()))
	location.SetValueCountry(observable.Data)
	if len(observable.Tags) > 0 {
		location.SetValueName(observable.Tags[0])
	}
	if observable.Message != "" {
		location.SetValueDescription(observable.Message)
	}

	return location
}

// CreateIndicatorSnortIdDomainObjectsSTIX формирует объект 'indicator' с описанием правил СОА Snort
func CreateIndicatorSnortIdDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjectsstix.IndicatorDomainObjectsSTIX {
	indicatior := methodstixobjects.NewIndicatorDomainObjectsSTIX()
	indicatior.SetValueID(fmt.Sprintf("indicator-%s", uuid.NewString()))
	indicatior.SetValueName("snort_sid")
	indicatior.SetValueDescription("list of signatures of the Snort computer attack detection tool")
	indicatior.SetValuePattern(observable.Data)
	indicatior.SetValuePatternType("list of numbers")

	if observable.Message != "" {
		indicatior.SetValueDescription(observable.Message)
	}

	return indicatior
}

// CreateIndicatorYaraDomainObjectsSTIX формирует объект 'indicator' с описанием правил в формате YARA
func CreateIndicatorYaraDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjectsstix.IndicatorDomainObjectsSTIX {
	indicatior := methodstixobjects.NewIndicatorDomainObjectsSTIX()
	indicatior.SetValueID(fmt.Sprintf("indicator-%s", uuid.NewString()))
	indicatior.SetValueName("yara")
	indicatior.SetValueDescription("yara rule")
	indicatior.SetValuePattern(observable.Data)
	indicatior.SetValuePatternType("string")

	if observable.Message != "" {
		indicatior.SetValueDescription(observable.Message)
	}

	return indicatior
}

// CreateIndicatorHashDomainObjectsSTIX формирует объект 'indicator' с хеш-суммой
func CreateIndicatorHashDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjectsstix.IndicatorDomainObjectsSTIX {
	indicatior := methodstixobjects.NewIndicatorDomainObjectsSTIX()
	indicatior.SetValueID(fmt.Sprintf("indicator-%s", uuid.NewString()))
	indicatior.SetValueName("hash")
	indicatior.SetValueDescription("hash sum")
	indicatior.SetValuePattern(observable.Data)
	indicatior.SetValuePatternType("string")

	if len(observable.Tags) > 0 {
		indicatior.SetValueDescription(observable.Tags[0])
	}

	return indicatior
}

// CreateIndicatorUserAgentDomainObjectsSTIX формирует объект 'indicator' с описанием User-agent
func CreateIndicatorUserAgentDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjectsstix.IndicatorDomainObjectsSTIX {
	indicatior := methodstixobjects.NewIndicatorDomainObjectsSTIX()
	indicatior.SetValueID(fmt.Sprintf("indicator-%s", uuid.NewString()))
	indicatior.SetValueName("user-agent")
	indicatior.SetValueDescription("user-agent")
	indicatior.SetValuePattern(observable.Data)
	indicatior.SetValuePatternType("string")

	if len(observable.Tags) > 0 {
		indicatior.SetValueDescription(observable.Tags[0])
	}

	return indicatior
}

// CreateIdentityDomainObjectsSTIX формирует объект 'identity'
func CreateIdentityDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjectsstix.IdentityDomainObjectsSTIX {
	identity := methodstixobjects.NewIdentityDomainObjectsSTIX()
	identity.SetValueID(fmt.Sprintf("identity-%s", uuid.NewString()))
	identity.SetValueName("phone-number")
	identity.SetValueDescription("personal phone number")
	identity.SetValueContactInformation(observable.Data)

	if observable.Message != "" {
		identity.SetValueDescription(observable.Message)
	}

	return identity
}
