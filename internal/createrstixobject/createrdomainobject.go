package createrstixobject

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix/domainobjects"
)

// CreateLocationDomainObjectsSTIX формирует объект 'location'
func CreateLocationDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjects.WrapperLocation {
	location := domainobjects.NewWrapperLocationDomainObjectsSTIX()
	location.SetValueCreated(observable.GetUnderliningCreatedAt())
	location.SetValueModified(observable.GetUnderliningUpdatedAt())
	location.SetValueID(fmt.Sprintf("location--%s", uuid.NewString()))
	location.SetValueSpecVersion("2.1")
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
func CreateIndicatorSnortIdDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjects.WrapperIndicator {
	indicator := domainobjects.NewWrapperIndicatorDomainObjectsSTIX()
	indicator.SetValueCreated(observable.GetUnderliningCreatedAt())
	indicator.SetValueModified(observable.GetUnderliningUpdatedAt())
	indicator.SetValueID(fmt.Sprintf("indicator--%s", uuid.NewString()))
	indicator.SetValueSpecVersion("2.1")
	indicator.SetValueName("snort_sid")
	indicator.SetValueDescription("list of signatures of the Snort computer attack detection tool")
	indicator.SetValuePattern(observable.Data)
	indicator.SetValuePatternType("list of numbers")

	if observable.Message != "" {
		indicator.SetValueDescription(observable.Message)
	}

	return indicator
}

// CreateIndicatorYaraDomainObjectsSTIX формирует объект 'indicator' с описанием правил в формате YARA
func CreateIndicatorYaraDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjects.WrapperIndicator {
	indicator := domainobjects.NewWrapperIndicatorDomainObjectsSTIX()
	indicator.SetValueCreated(observable.GetUnderliningCreatedAt())
	indicator.SetValueModified(observable.GetUnderliningUpdatedAt())
	indicator.SetValueID(fmt.Sprintf("indicator--%s", uuid.NewString()))
	indicator.SetValueSpecVersion("2.1")
	indicator.SetValueName("yara")
	indicator.SetValueDescription("yara rule")
	indicator.SetValuePattern(observable.Data)
	indicator.SetValuePatternType("string")

	if observable.Message != "" {
		indicator.SetValueDescription(observable.Message)
	}

	return indicator
}

// CreateIndicatorHashDomainObjectsSTIX формирует объект 'indicator' с хеш-суммой
func CreateIndicatorHashDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjects.WrapperIndicator {
	indicator := domainobjects.NewWrapperIndicatorDomainObjectsSTIX()
	indicator.SetValueCreated(observable.GetUnderliningCreatedAt())
	indicator.SetValueModified(observable.GetUnderliningUpdatedAt())
	indicator.SetValueID(fmt.Sprintf("indicator--%s", uuid.NewString()))
	indicator.SetValueSpecVersion("2.1")
	indicator.SetValueName("hash")
	indicator.SetValueDescription("hash sum")
	indicator.SetValuePattern(observable.Data)
	indicator.SetValuePatternType("string")

	if len(observable.Tags) > 0 {
		indicator.SetValueDescription(observable.Tags[0])
	}

	return indicator
}

// CreateIndicatorUserAgentDomainObjectsSTIX формирует объект 'indicator' с описанием User-agent
func CreateIndicatorUserAgentDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjects.WrapperIndicator {
	indicator := domainobjects.NewWrapperIndicatorDomainObjectsSTIX()
	indicator.SetValueCreated(observable.GetUnderliningCreatedAt())
	indicator.SetValueModified(observable.GetUnderliningUpdatedAt())
	indicator.SetValueID(fmt.Sprintf("indicator--%s", uuid.NewString()))
	indicator.SetValueSpecVersion("2.1")
	indicator.SetValueName("user-agent")
	indicator.SetValueDescription("user-agent")
	indicator.SetValuePattern(observable.Data)
	indicator.SetValuePatternType("string")

	if len(observable.Tags) > 0 {
		indicator.SetValueDescription(observable.Tags[0])
	}

	return indicator
}

// CreateIdentityDomainObjectsSTIX формирует объект 'identity'
func CreateIdentityDomainObjectsSTIX(observable datamodels.ObservableMessage) *domainobjects.WrapperIdentity {
	identity := domainobjects.NewWrapperIdentityDomainObjectsSTIX()
	identity.SetValueCreated(observable.GetUnderliningCreatedAt())
	identity.SetValueModified(observable.GetUnderliningUpdatedAt())
	identity.SetValueID(fmt.Sprintf("identity--%s", uuid.NewString()))
	identity.SetValueSpecVersion("2.1")
	identity.SetValueName("phone-number")
	identity.SetValueDescription("personal phone number")
	identity.SetValueContactInformation(observable.Data)

	if observable.Message != "" {
		identity.SetValueDescription(observable.Message)
	}

	return identity
}
