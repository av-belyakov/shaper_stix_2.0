package supportingfunctions

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// GetListElementSTIXObject возвращает, из БД, список STIX объектов
func GetListElementSTIXObject(cur *mongo.Cursor) []*datamodels.ElementSTIXObject {
	elements := []*datamodels.ElementSTIXObject{}

	//fmt.Println("func 'GetListElementSTIXObject', START...")

	for cur.Next(context.Background()) {
		var modelType definingTypeSTIXObject
		if err := cur.Decode(&modelType); err != nil {

			//fmt.Println("func 'GetListElementSTIXObject', cur.Decode(&modelType), CONTINUE")

			continue
		}

		//fmt.Println("func 'GetListElementSTIXObject', modelType.Type:", modelType)

		switch modelType.Type {
		/* *** Domain Objects STIX *** */
		case "attack-pattern":
			tmpObj := mstixo.AttackPatternDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.AttackPatternDomainObjectsSTIX{AttackPatternDomainObjectsSTIX: tmpObj},
			})
		case "campaign":
			tmpObj := mstixo.CampaignDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.CampaignDomainObjectsSTIX{CampaignDomainObjectsSTIX: tmpObj},
			})

		case "course-of-action":
			tmpObj := mstixo.CourseOfActionDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.CourseOfActionDomainObjectsSTIX{CourseOfActionDomainObjectsSTIX: tmpObj},
			})

		case "grouping":
			tmpObj := mstixo.GroupingDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.GroupingDomainObjectsSTIX{GroupingDomainObjectsSTIX: tmpObj},
			})

		case "identity":
			tmpObj := mstixo.IdentityDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.IdentityDomainObjectsSTIX{IdentityDomainObjectsSTIX: tmpObj},
			})

		case "indicator":
			tmpObj := mstixo.IndicatorDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.IndicatorDomainObjectsSTIX{IndicatorDomainObjectsSTIX: tmpObj},
			})

		case "infrastructure":
			tmpObj := mstixo.InfrastructureDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.InfrastructureDomainObjectsSTIX{InfrastructureDomainObjectsSTIX: tmpObj},
			})

		case "intrusion-set":
			tmpObj := mstixo.IntrusionSetDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.IntrusionSetDomainObjectsSTIX{IntrusionSetDomainObjectsSTIX: tmpObj},
			})

		case "location":
			tmpObj := mstixo.LocationDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.LocationDomainObjectsSTIX{LocationDomainObjectsSTIX: tmpObj},
			})

		case "malware":
			tmpObj := mstixo.MalwareDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.MalwareDomainObjectsSTIX{MalwareDomainObjectsSTIX: tmpObj},
			})

		case "malware-analysis":
			tmpObj := mstixo.MalwareAnalysisDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.MalwareAnalysisDomainObjectsSTIX{MalwareAnalysisDomainObjectsSTIX: tmpObj},
			})

		case "note":
			tmpObj := mstixo.NoteDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.NoteDomainObjectsSTIX{NoteDomainObjectsSTIX: tmpObj},
			})

		case "observed-data":
			tmpObj := mstixo.ObservedDataDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ObservedDataDomainObjectsSTIX{ObservedDataDomainObjectsSTIX: tmpObj},
			})

		case "opinion":
			tmpObj := mstixo.OpinionDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.OpinionDomainObjectsSTIX{OpinionDomainObjectsSTIX: tmpObj},
			})

		case "report":
			tmpObj := mstixo.ReportDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ReportDomainObjectsSTIX{ReportDomainObjectsSTIX: tmpObj},
			})

		case "threat-actor":
			tmpObj := mstixo.ThreatActorDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ThreatActorDomainObjectsSTIX{ThreatActorDomainObjectsSTIX: tmpObj},
			})

		case "tool":
			tmpObj := mstixo.ToolDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ToolDomainObjectsSTIX{ToolDomainObjectsSTIX: tmpObj},
			})

		case "vulnerability":
			tmpObj := mstixo.VulnerabilityDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.VulnerabilityDomainObjectsSTIX{VulnerabilityDomainObjectsSTIX: tmpObj},
			})

		/* *** Relationship Objects *** */
		case "relationship":
			tmpObj := mstixo.RelationshipObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.RelationshipObjectSTIX{RelationshipObjectSTIX: tmpObj},
			})

		case "sighting":
			tmpObj := mstixo.SightingObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.SightingObjectSTIX{SightingObjectSTIX: tmpObj},
			})

		/* *** Cyber-observable Objects STIX *** */
		case "artifact":
			tmpObj := mstixo.ArtifactCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ArtifactCyberObservableObjectSTIX{ArtifactCyberObservableObjectSTIX: tmpObj},
			})

		case "autonomous-system":
			tmpObj := mstixo.AutonomousSystemCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.AutonomousSystemCyberObservableObjectSTIX{AutonomousSystemCyberObservableObjectSTIX: tmpObj},
			})

		case "directory":
			tmpObj := mstixo.DirectoryCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.DirectoryCyberObservableObjectSTIX{DirectoryCyberObservableObjectSTIX: tmpObj},
			})

		case "domain-name":
			tmpObj := mstixo.DomainNameCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.DomainNameCyberObservableObjectSTIX{DomainNameCyberObservableObjectSTIX: tmpObj},
			})

		case "email-addr":
			tmpObj := mstixo.EmailAddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.EmailAddressCyberObservableObjectSTIX{EmailAddressCyberObservableObjectSTIX: tmpObj},
			})

		case "email-message":
			tmpObj := mstixo.EmailMessageCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.EmailMessageCyberObservableObjectSTIX{EmailMessageCyberObservableObjectSTIX: tmpObj},
			})

		case "file":
			tmpObj := mstixo.FileCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.FileCyberObservableObjectSTIX{FileCyberObservableObjectSTIX: tmpObj},
			})

		case "ipv4-addr":
			tmpObj := datamodels.IPv4AddressCyberObservableSimilarObjectSTIX{}
			//tmpObj := datamodels.IPv4AddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			/*resolvesToRefs := make([]mstixo.IdentifierTypeSTIX, 0, len(tmpObj.ResolvesToRefs))
			for _, v := range tmpObj.ResolvesToRefs {
				resolvesToRefs = append(resolvesToRefs, v)
			}*/

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data: datamodels.IPv4AddressCyberObservableObjectSTIX{
					IPv4AddressCyberObservableObjectSTIX: mstixo.IPv4AddressCyberObservableObjectSTIX{
						CommonPropertiesObjectSTIX: mstixo.CommonPropertiesObjectSTIX{
							Type: tmpObj.CommonPropertiesObjectSTIX.Type,
							ID:   tmpObj.CommonPropertiesObjectSTIX.ID,
						},
						OptionalCommonPropertiesCyberObservableObjectSTIX: mstixo.OptionalCommonPropertiesCyberObservableObjectSTIX{
							SpecVersion:       tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.SpecVersion,
							ObjectMarkingRefs: tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.ObjectMarkingRefs,
							GranularMarkings:  tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.GranularMarkings,
							Defanged:          tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.Defanged,
						},
						Value:          tmpObj.Value,
						ResolvesToRefs: tmpObj.ResolvesToRefs,
						BelongsToRefs:  tmpObj.BelongsToRefs,
					},
				},
			})

			/*
				elements = append(elements, &datamodels.ElementSTIXObject{
					DataType: modelType.Type,
					Data: datamodels.IPv4AddressCyberObservableObjectSTIX{
						mstixo.IPv4AddressCyberObservableObjectSTIX{
							CommonPropertiesObjectSTIX: mstixo.CommonPropertiesObjectSTIX{
								Type: tmpObj.CommonPropertiesObjectSTIX.Type,
								ID:   tmpObj.CommonPropertiesObjectSTIX.ID,
							},
							OptionalCommonPropertiesCyberObservableObjectSTIX: mstixo.OptionalCommonPropertiesCyberObservableObjectSTIX{
								SpecVersion:       tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.SpecVersion,
								ObjectMarkingRefs: tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.ObjectMarkingRefs,
								GranularMarkings:  tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.GranularMarkings,
								Defanged:          tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.Defanged,
							},
							Value:          tmpObj.Value,
							ResolvesToRefs: tmpObj.ResolvesToRefs,
							BelongsToRefs:  tmpObj.BelongsToRefs,
						},
					},
				})
			*/

		case "ipv6-addr":
			tmpObj := mstixo.IPv6AddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.IPv6AddressCyberObservableObjectSTIX{IPv6AddressCyberObservableObjectSTIX: tmpObj},
			})

		case "mac-addr":
			tmpObj := mstixo.MACAddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.MACAddressCyberObservableObjectSTIX{MACAddressCyberObservableObjectSTIX: tmpObj},
			})

		case "mutex":
			tmpObj := mstixo.MutexCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.MutexCyberObservableObjectSTIX{MutexCyberObservableObjectSTIX: tmpObj},
			})

		case "network-traffic":
			tmpObj := mstixo.NetworkTrafficCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.NetworkTrafficCyberObservableObjectSTIX{NetworkTrafficCyberObservableObjectSTIX: tmpObj},
			})

		case "process":
			tmpObj := mstixo.ProcessCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ProcessCyberObservableObjectSTIX{ProcessCyberObservableObjectSTIX: tmpObj},
			})

		case "software":
			tmpObj := mstixo.SoftwareCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.SoftwareCyberObservableObjectSTIX{SoftwareCyberObservableObjectSTIX: tmpObj},
			})

		case "url":
			tmpObj := mstixo.URLCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.URLCyberObservableObjectSTIX{URLCyberObservableObjectSTIX: tmpObj},
			})

		case "user-account":
			tmpObj := mstixo.UserAccountCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.UserAccountCyberObservableObjectSTIX{UserAccountCyberObservableObjectSTIX: tmpObj},
			})

		case "windows-registry-key":
			tmpObj := mstixo.WindowsRegistryKeyCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.WindowsRegistryKeyCyberObservableObjectSTIX{WindowsRegistryKeyCyberObservableObjectSTIX: tmpObj},
			})

		case "x509-certificate":
			tmpObj := mstixo.X509CertificateCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.X509CertificateCyberObservableObjectSTIX{X509CertificateCyberObservableObjectSTIX: tmpObj},
			})
		}
	}

	//fmt.Println("func 'GetListElementSTIXObject' STOP elements ", elements)

	return elements
}
