package wrappersobjectstix

import (
	"time"

	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

func NewCommonPropertiesDomainObjectSTIX() CommonPropertiesDomainObjectSTIX {
	var (
		create   time.Time
		modified time.Time
	)

	if newTime, err := time.Parse(time.RFC3339, "1970-01-01T00:00:00+00:00"); err == nil {
		create = newTime
		modified = newTime
	}

	return CommonPropertiesDomainObjectSTIX{
		Created:            create,
		Modified:           modified,
		Labels:             []string(nil),
		Extensions:         make(map[string]string),
		ExternalReferences: []stixhelpers.ExternalReferenceTypeElementSTIX(nil),
		ObjectMarkingRefs:  []stixhelpers.IdentifierTypeSTIX(nil),
		GranularMarkings:   []stixhelpers.GranularMarkingsTypeSTIX(nil),
	}
}
