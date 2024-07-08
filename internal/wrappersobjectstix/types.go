package wrappersobjectstix

import (
	"time"

	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

// CommonPropertiesDomainObjectSTIX свойства общие, для всех объектов STIX
type CommonPropertiesDomainObjectSTIX struct {
	Revoked            bool                                           `bson:"revoked"`
	Defanged           bool                                           `bson:"defanged"`
	Сonfidence         int                                            `bson:"confidence"`
	Lang               string                                         `bson:"lang"`
	SpecVersion        string                                         `bson:"spec_version" required:"true"`
	Created            time.Time                                      `bson:"created" required:"true"`
	Modified           time.Time                                      `bson:"modified" required:"true"`
	Labels             []string                                       `bson:"labels"`
	Extensions         map[string]string                              `bson:"extensions"`
	CreatedByRef       stixhelpers.IdentifierTypeSTIX                 `bson:"created_by_ref"`
	ExternalReferences []stixhelpers.ExternalReferenceTypeElementSTIX `bson:"external_references"`
	ObjectMarkingRefs  []stixhelpers.IdentifierTypeSTIX               `bson:"object_marking_refs"`
	GranularMarkings   []stixhelpers.GranularMarkingsTypeSTIX         `bson:"granular_markings"`
}
