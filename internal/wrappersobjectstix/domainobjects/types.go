package domainobjects

import (
	"time"

	"github.com/av-belyakov/methodstixobjects/datamodels/commonproperties"
	"github.com/av-belyakov/methodstixobjects/datamodels/domainobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix"
)

// WrapperIdentity тип содержит дополнительный метод MarshalBSON
type WrapperIdentity struct {
	*domainobjectsstix.IdentityDomainObjectsSTIX
}

type FinalyIdentityObjects struct {
	commonproperties.CommonPropertiesObjectSTIX
	wrappersobjectstix.CommonPropertiesDomainObjectSTIX
	Name               string                          `bson:"name" required:"true"`
	Description        string                          `bson:"description"`
	ContactInformation string                          `bson:"contact_information"`
	Roles              []string                        `bson:"roles"`
	IdentityClass      stixhelpers.OpenVocabTypeSTIX   `bson:"identity_class"`
	Sectors            []stixhelpers.OpenVocabTypeSTIX `bson:"sectors"`
}

// WrapperLocation тип содержит дополнительный метод MarshalBSON
type WrapperLocation struct {
	*domainobjectsstix.LocationDomainObjectsSTIX
}

type FinalyLocationObjects struct {
	commonproperties.CommonPropertiesObjectSTIX
	wrappersobjectstix.CommonPropertiesDomainObjectSTIX
	Latitude           float32                       `bson:"latitude"`
	Longitude          float32                       `bson:"longitude"`
	Precision          float32                       `bson:"precision"`
	Name               string                        `bson:"name"`
	Description        string                        `bson:"description"`
	Country            string                        `bson:"country"`
	AdministrativeArea string                        `bson:"administrative_area"`
	City               string                        `bson:"city"`
	StreetAddress      string                        `bson:"street_address"`
	PostalCode         string                        `bson:"postal_code"`
	Region             stixhelpers.OpenVocabTypeSTIX `bson:"region"`
}

// WrapperIndicator тип содержит дополнительный метод MarshalBSON
type WrapperIndicator struct {
	*domainobjectsstix.IndicatorDomainObjectsSTIX
}

type FinalyIndicatorObject struct {
	commonproperties.CommonPropertiesObjectSTIX
	wrappersobjectstix.CommonPropertiesDomainObjectSTIX
	Name            string                                       `bson:"name" required:"true"`
	Pattern         string                                       `bson:"pattern" required:"true"`
	PatternVersion  string                                       `bson:"pattern_version"`
	Description     string                                       `bson:"description"`
	ValidFrom       time.Time                                    `bson:"valid_from" required:"true"`
	ValidUntil      time.Time                                    `bson:"valid_until"`
	PatternType     stixhelpers.OpenVocabTypeSTIX                `bson:"pattern_type" required:"true"`
	KillChainPhases []stixhelpers.KillChainPhasesTypeElementSTIX `bson:"kill_chain_phases"`
	IndicatorTypes  []stixhelpers.OpenVocabTypeSTIX              `bson:"indicator_types"`
}

// WrapperReport содержит STIX объект 'report' и дополнительные расширеные свойства
// кроме того этот тип содержит дополнительный метод MarshalBSON
type WrapperReport struct {
	*domainobjectsstix.ReportDomainObjectsSTIX
	ReportOutsideSpecification ReportOutsideSpecification
}

type FinalyReportObject struct {
	commonproperties.CommonPropertiesObjectSTIX
	wrappersobjectstix.CommonPropertiesDomainObjectSTIX
	Name                       string                           `bson:"name" required:"true"`
	Description                string                           `bson:"description"`
	Published                  time.Time                        `bson:"published" required:"true"`
	ReportTypes                []stixhelpers.OpenVocabTypeSTIX  `bson:"report_types"`
	ObjectRefs                 []stixhelpers.IdentifierTypeSTIX `bson:"object_refs" required:"true"`
	ReportOutsideSpecification `bson:"outside_specification"`
}

// ReportOutsideSpecification содержит дополнительные свойства
type ReportOutsideSpecification struct {
	ObjectType                  string `bson:"object_type"`
	RootId                      string `bson:"root_id"`
	ObjectId                    string `bson:"object_id"`
	CaseId                      string `bson:"case_id"`
	StartDate                   string `bson:"start_date"`
	EndDate                     string `bson:"end_date"`
	ImpactStatus                string `bson:"impact_status"`
	ResolutionStatus            string `bson:"resolution_status"`
	AdditionalName              string `bson:"additional_name"`
	DecisionsMadeComputerThreat string `bson:"decisions_made_computer_threat"`
	ComputerThreatType          string `bson:"computer_threat_type"`
}
