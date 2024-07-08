package domainobjects

import (
	"time"

	"github.com/av-belyakov/methodstixobjects/datamodels/commonproperties"
	"github.com/av-belyakov/methodstixobjects/datamodels/domainobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix"
)

// WrapperReport содержит STIX объект 'report' и дополнительные расширеные свойства
type WrapperReport struct {
	*domainobjectsstix.ReportDomainObjectsSTIX
	ReportOutsideSpecification ReportOutsideSpecification
}

type FinalyIndicatorObject struct {
	commonproperties.CommonPropertiesObjectSTIX
	wrappersobjectstix.CommonPropertiesDomainObjectSTIX
	Name            string                                       `json:"name" bson:"name" required:"true"`
	Pattern         string                                       `json:"pattern" bson:"pattern" required:"true"`
	PatternVersion  string                                       `json:"pattern_version" bson:"pattern_version"`
	Description     string                                       `json:"description" bson:"description"`
	ValidFrom       time.Time                                    `json:"valid_from" bson:"valid_from" required:"true"`
	ValidUntil      time.Time                                    `json:"valid_until" bson:"valid_until"`
	PatternType     stixhelpers.OpenVocabTypeSTIX                `json:"pattern_type" bson:"pattern_type" required:"true"`
	KillChainPhases []stixhelpers.KillChainPhasesTypeElementSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
	IndicatorTypes  []stixhelpers.OpenVocabTypeSTIX              `json:"indicator_types" bson:"indicator_types"`
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
