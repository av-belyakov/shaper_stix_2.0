package datamodels

import (
	"fmt"
	"strings"

	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

//****************** CommonObservableType ******************

func (o *CommonObservableType) Get() *CommonObservableType {
	return o
}

func (o *CommonObservableType) GetIoc() bool {
	return o.Ioc
}

// SetValueIoc устанавливает BOOL значение для поля Ioc
func (o *CommonObservableType) SetValueIoc(v bool) {
	o.Ioc = v
}

// SetAnyIoc устанавливает ЛЮБОЕ значение для поля Ioc
func (o *CommonObservableType) SetAnyIoc(i interface{}) {
	if v, ok := i.(bool); ok {
		o.Ioc = v
	}
}

func (o *CommonObservableType) GetSighted() bool {
	return o.Sighted
}

// SetValueSighted устанавливает BOOL значение для поля Sighted
func (o *CommonObservableType) SetValueSighted(v bool) {
	o.Sighted = v
}

// SetAnySighted устанавливает ЛЮБОЕ значение для поля Sighted
func (o *CommonObservableType) SetAnySighted(i interface{}) {
	if v, ok := i.(bool); ok {
		o.Sighted = v
	}
}

func (o *CommonObservableType) GetIgnoreSimilarity() bool {
	return o.IgnoreSimilarity
}

// SetValueIgnoreSimilarity устанавливает BOOL значение для поля IgnoreSimilarity
func (o *CommonObservableType) SetValueIgnoreSimilarity(v bool) {
	o.IgnoreSimilarity = v
}

// SetAnyIgnoreSimilarity устанавливает ЛЮБОЕ значение для поля IgnoreSimilarity
func (o *CommonObservableType) SetAnyIgnoreSimilarity(i interface{}) {
	if v, ok := i.(bool); ok {
		o.IgnoreSimilarity = v
	}
}

func (o *CommonObservableType) GetTlp() uint64 {
	return o.Tlp
}

// SetValueTlp устанавливает INT значение для поля Tlp
func (o *CommonObservableType) SetValueTlp(v uint64) {
	o.Tlp = v
}

// SetAnyTlp устанавливает ЛЮБОЕ значение для поля Tlp
func (o *CommonObservableType) SetAnyTlp(i interface{}) {
	if v, ok := i.(float32); ok {
		o.Tlp = uint64(v)

		return
	}

	if v, ok := i.(float64); ok {
		o.Tlp = uint64(v)

		return
	}

	if v, ok := i.(uint64); ok {
		o.Tlp = v
	}
}

func (o *CommonObservableType) GetUnderliningCreatedAt() string {
	return o.UnderliningCreatedAt
}

// SetValueUnderliningCreatedAt устанавливает значение в формате RFC3339 для поля CreatedAt
func (o *CommonObservableType) SetValueUnderliningCreatedAt(v string) {
	o.UnderliningCreatedAt = v
}

// SetAnyUnderliningCreatedAt устанавливает ЛЮБОЕ значение для поля CreatedAt
func (o *CommonObservableType) SetAnyUnderliningCreatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UnderliningCreatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonObservableType) GetUnderliningUpdatedAt() string {
	return o.UnderliningUpdatedAt
}

// SetValueUnderliningUpdatedAt устанавливает значениев формате RFC3339 для поля UpdatedAt
func (o *CommonObservableType) SetValueUnderliningUpdatedAt(v string) {
	o.UnderliningUpdatedAt = v
}

// SetAnyUnderliningUpdatedAt устанавливает ЛЮБОЕ значение для поля UpdatedAt
func (o *CommonObservableType) SetAnyUnderliningUpdatedAt(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.UnderliningUpdatedAt = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonObservableType) GetStartDate() string {
	return o.StartDate
}

// SetValueStartDate устанавливает значениев формате RFC3339 для поля StartDate
func (o *CommonObservableType) SetValueStartDate(v string) {
	o.StartDate = v
}

// SetAnyStartDate устанавливает ЛЮБОЕ значение для поля StartDate
func (o *CommonObservableType) SetAnyStartDate(i interface{}) {
	tmp := supportingfunctions.ConversionAnyToInt(i)
	o.StartDate = supportingfunctions.GetDateTimeFormatRFC3339(int64(tmp))
}

func (o *CommonObservableType) GetUnderliningCreatedBy() string {
	return o.UnderliningCreatedBy
}

// SetValueUnderliningCreatedBy устанавливает STRING значение для поля CreatedBy
func (o *CommonObservableType) SetValueUnderliningCreatedBy(v string) {
	o.UnderliningCreatedBy = v
}

// SetAnyUnderliningCreatedBy устанавливает ЛЮБОЕ значение для поля CreatedBy
func (o *CommonObservableType) SetAnyUnderliningCreatedBy(i interface{}) {
	o.UnderliningCreatedBy = fmt.Sprint(i)
}

func (o *CommonObservableType) GetUnderliningUpdatedBy() string {
	return o.UnderliningUpdatedBy
}

// SetValueUnderliningUpdatedBy устанавливает STRING значение для поля UpdatedBy
func (o *CommonObservableType) SetValueUnderliningUpdatedBy(v string) {
	o.UnderliningUpdatedBy = v
}

// SetAnyUnderliningUpdatedBy устанавливает ЛЮБОЕ значение для поля UpdatedBy
func (o *CommonObservableType) SetAnyUnderliningUpdatedBy(i interface{}) {
	o.UnderliningUpdatedBy = fmt.Sprint(i)
}

func (o *CommonObservableType) GetUnderliningId() string {
	return o.UnderliningId
}

// SetValueUnderliningId устанавливает STRING значение для поля UnderliningId
func (o *CommonObservableType) SetValueUnderliningId(v string) {
	o.UnderliningId = v
}

// SetAnyUnderliningId устанавливает ЛЮБОЕ значение для поля UnderliningId
func (o *CommonObservableType) SetAnyUnderliningId(i interface{}) {
	o.UnderliningId = fmt.Sprint(i)
}

func (o *CommonObservableType) GetUnderliningType() string {
	return o.UnderliningType
}

// SetValueUnderliningType устанавливает STRING значение для поля UnderliningType
func (o *CommonObservableType) SetValueUnderliningType(v string) {
	o.UnderliningType = v
}

// SetAnyUnderliningType устанавливает ЛЮБОЕ значение для поля UnderliningType
func (o *CommonObservableType) SetAnyUnderliningType(i interface{}) {
	o.UnderliningType = fmt.Sprint(i)
}

func (o *CommonObservableType) GetData() string {
	return o.Data
}

// SetValueData устанавливает STRING значение для поля Data
func (o *CommonObservableType) SetValueData(v string) {
	o.Data = v
}

// SetAnyData устанавливает ЛЮБОЕ значение для поля Data
func (o *CommonObservableType) SetAnyData(i interface{}) {
	o.Data = fmt.Sprint(i)
}

func (o *CommonObservableType) GetDataType() string {
	return o.DataType
}

// SetValueDataType устанавливает STRING значение для поля DataType
func (o *CommonObservableType) SetValueDataType(v string) {
	o.DataType = v
}

// SetAnyDataType устанавливает ЛЮБОЕ значение для поля DataType
func (o *CommonObservableType) SetAnyDataType(i interface{}) {
	o.DataType = fmt.Sprint(i)
}

func (o *CommonObservableType) GetMessage() string {
	return o.Message
}

// SetValueMessage устанавливает STRING значение для поля Message
func (o *CommonObservableType) SetValueMessage(v string) {
	o.Message = v
}

// SetAnyMessage устанавливает ЛЮБОЕ значение для поля Message
func (o *CommonObservableType) SetAnyMessage(i interface{}) {
	o.Message = fmt.Sprint(i)
}

func (om CommonObservableType) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	ws := supportingfunctions.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'_createdAt': '%s'\n", ws, om.UnderliningCreatedAt))
	str.WriteString(fmt.Sprintf("%s'_createdBy': '%s'\n", ws, om.UnderliningCreatedBy))
	str.WriteString(fmt.Sprintf("%s'_id': '%s'\n", ws, om.UnderliningId))
	str.WriteString(fmt.Sprintf("%s'_type': '%s'\n", ws, om.UnderliningType))
	str.WriteString(fmt.Sprintf("%s'_updatedAt': '%s'\n", ws, om.UnderliningUpdatedAt))
	str.WriteString(fmt.Sprintf("%s'_updatedBy': '%s'\n", ws, om.UnderliningUpdatedBy))
	str.WriteString(fmt.Sprintf("%s'data': '%s'\n", ws, om.Data))
	str.WriteString(fmt.Sprintf("%s'dataType': '%s'\n", ws, om.DataType))
	str.WriteString(fmt.Sprintf("%s'ignoreSimilarity': '%v'\n", ws, om.IgnoreSimilarity))
	str.WriteString(fmt.Sprintf("%s'ioc': '%v'\n", ws, om.Ioc))
	str.WriteString(fmt.Sprintf("%s'message': '%s'\n", ws, om.Message))
	str.WriteString(fmt.Sprintf("%s'sighted': '%v'\n", ws, om.Sighted))
	str.WriteString(fmt.Sprintf("%s'startDate': '%s'\n", ws, om.StartDate))
	str.WriteString(fmt.Sprintf("%s'tlp': '%d'\n", ws, om.Tlp))

	return str.String()
}

/********** 			Domain Objects STIX			**********/

// GetAttackPatternDomainObjectsSTIX возвращает объект STIX типа 'attack-pattern'
func (estix *ElementSTIXObject) GetAttackPatternDomainObjectsSTIX() *mstixo.AttackPatternDomainObjectsSTIX {
	if result, ok := estix.Data.(AttackPatternDomainObjectsSTIX); ok {
		return &result.AttackPatternDomainObjectsSTIX
	}

	return nil
}

// GetCampaignDomainObjectsSTIX возвращает объект STIX типа 'campaign'
func (estix *ElementSTIXObject) GetCampaignDomainObjectsSTIX() *mstixo.CampaignDomainObjectsSTIX {
	if result, ok := estix.Data.(CampaignDomainObjectsSTIX); ok {
		return &result.CampaignDomainObjectsSTIX
	}

	return nil
}

// GetCourseOfActionDomainObjectsSTIX возвращает объект STIX типа 'course-of-action'
func (estix *ElementSTIXObject) GetCourseOfActionDomainObjectsSTIX() *mstixo.CourseOfActionDomainObjectsSTIX {
	if result, ok := estix.Data.(CourseOfActionDomainObjectsSTIX); ok {
		return &result.CourseOfActionDomainObjectsSTIX
	}

	return nil
}

// GetGroupingDomainObjectsSTIX возвращает объект STIX типа 'grouping'
func (estix *ElementSTIXObject) GetGroupingDomainObjectsSTIX() *mstixo.GroupingDomainObjectsSTIX {
	if result, ok := estix.Data.(GroupingDomainObjectsSTIX); ok {
		return &result.GroupingDomainObjectsSTIX
	}

	return nil
}

// GetIdentityDomainObjectsSTIX возвращает объект STIX типа 'identity'
func (estix *ElementSTIXObject) GetIdentityDomainObjectsSTIX() *mstixo.IdentityDomainObjectsSTIX {
	if result, ok := estix.Data.(IdentityDomainObjectsSTIX); ok {
		return &result.IdentityDomainObjectsSTIX
	}

	return nil
}

// GetIndicatorDomainObjectsSTIX возвращает объект STIX типа 'indicator'
func (estix *ElementSTIXObject) GetIndicatorDomainObjectsSTIX() *mstixo.IndicatorDomainObjectsSTIX {
	if result, ok := estix.Data.(IndicatorDomainObjectsSTIX); ok {
		return &result.IndicatorDomainObjectsSTIX
	}

	return nil
}

// GetInfrastructureDomainObjectsSTIX возвращает объект STIX типа 'infrastructure'
func (estix *ElementSTIXObject) GetInfrastructureDomainObjectsSTIX() *mstixo.InfrastructureDomainObjectsSTIX {
	if result, ok := estix.Data.(InfrastructureDomainObjectsSTIX); ok {
		return &result.InfrastructureDomainObjectsSTIX
	}

	return nil
}

// GetIntrusionSetDomainObjectsSTIX возвращает объект STIX типа 'intrusion-set'
func (estix *ElementSTIXObject) GetIntrusionSetDomainObjectsSTIX() *mstixo.IntrusionSetDomainObjectsSTIX {
	if result, ok := estix.Data.(IntrusionSetDomainObjectsSTIX); ok {
		return &result.IntrusionSetDomainObjectsSTIX
	}

	return nil
}

// GetLocationDomainObjectsSTIX возвращает объект STIX типа 'location'
func (estix *ElementSTIXObject) GetLocationDomainObjectsSTIX() *mstixo.LocationDomainObjectsSTIX {
	if result, ok := estix.Data.(LocationDomainObjectsSTIX); ok {
		return &result.LocationDomainObjectsSTIX
	}

	return nil
}

// GetMalwareDomainObjectsSTIX возвращает объект STIX типа 'malware'
func (estix *ElementSTIXObject) GetMalwareDomainObjectsSTIX() *mstixo.MalwareDomainObjectsSTIX {
	if result, ok := estix.Data.(MalwareDomainObjectsSTIX); ok {
		return &result.MalwareDomainObjectsSTIX
	}

	return nil
}

// GetMalwareAnalysisDomainObjectsSTIX возвращает объект STIX типа 'malware-analysis'
func (estix *ElementSTIXObject) GetMalwareAnalysisDomainObjectsSTIX() *mstixo.MalwareAnalysisDomainObjectsSTIX {
	if result, ok := estix.Data.(MalwareAnalysisDomainObjectsSTIX); ok {
		return &result.MalwareAnalysisDomainObjectsSTIX
	}

	return nil
}

// GetNoteDomainObjectsSTIX возвращает объект STIX типа 'note'
func (estix *ElementSTIXObject) GetNoteDomainObjectsSTIX() *mstixo.NoteDomainObjectsSTIX {
	if result, ok := estix.Data.(NoteDomainObjectsSTIX); ok {
		return &result.NoteDomainObjectsSTIX
	}

	return nil
}

// GetObservedDataDomainObjectsSTIX возвращает объект STIX типа 'observed-data'
func (estix *ElementSTIXObject) GetObservedDataDomainObjectsSTIX() *mstixo.ObservedDataDomainObjectsSTIX {
	if result, ok := estix.Data.(ObservedDataDomainObjectsSTIX); ok {
		return &result.ObservedDataDomainObjectsSTIX
	}

	return nil
}

// GetOpinionDomainObjectsSTIX возвращает объект STIX типа 'opinion'
func (estix *ElementSTIXObject) GetOpinionDomainObjectsSTIX() *mstixo.OpinionDomainObjectsSTIX {
	if result, ok := estix.Data.(OpinionDomainObjectsSTIX); ok {
		return &result.OpinionDomainObjectsSTIX
	}

	return nil
}

// GetReportDomainObjectsSTIX возвращает объект STIX типа 'report'
func (estix *ElementSTIXObject) GetReportDomainObjectsSTIX() *mstixo.ReportDomainObjectsSTIX {
	if result, ok := estix.Data.(ReportDomainObjectsSTIX); ok {
		return &result.ReportDomainObjectsSTIX
	}

	return nil
}

// GetThreatActorDomainObjectsSTIX возвращает объект STIX типа 'threat-actor'
func (estix *ElementSTIXObject) GetThreatActorDomainObjectsSTIX() *mstixo.ThreatActorDomainObjectsSTIX {
	if result, ok := estix.Data.(ThreatActorDomainObjectsSTIX); ok {
		return &result.ThreatActorDomainObjectsSTIX
	}

	return nil
}

// GetToolDomainObjectsSTIX возвращает объект STIX типа 'tool'
func (estix *ElementSTIXObject) GetToolDomainObjectsSTIX() *mstixo.ToolDomainObjectsSTIX {
	if result, ok := estix.Data.(ToolDomainObjectsSTIX); ok {
		return &result.ToolDomainObjectsSTIX
	}

	return nil
}

// GetVulnerabilityDomainObjectsSTIX возвращает объект STIX типа 'vulnerability'
func (estix *ElementSTIXObject) GetVulnerabilityDomainObjectsSTIX() *mstixo.VulnerabilityDomainObjectsSTIX {
	if result, ok := estix.Data.(VulnerabilityDomainObjectsSTIX); ok {
		return &result.VulnerabilityDomainObjectsSTIX
	}

	return nil
}

/********** 			Relationship Objects STIX			**********/

// GetRelationshipObjectSTIX возвращает объект STIX типа 'relationship'
func (estix *ElementSTIXObject) GetRelationshipObjectSTIX() *RelationshipObjectSTIX {
	if result, ok := estix.Data.(RelationshipObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetSightingObjectSTIX возвращает объект STIX типа 'sighting'
func (estix *ElementSTIXObject) GetSightingObjectSTIX() *SightingObjectSTIX {
	if result, ok := estix.Data.(SightingObjectSTIX); ok {
		return &result
	}

	return nil
}

/********** 			Cyber-observable Objects STIX			**********/

// GetArtifactCyberObservableObjectSTIX возвращает объект STIX типа 'artifact'
func (estix *ElementSTIXObject) GetArtifactCyberObservableObjectSTIX() *ArtifactCyberObservableObjectSTIX {
	if result, ok := estix.Data.(ArtifactCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetAutonomousSystemCyberObservableObjectSTIX возвращает объект STIX типа 'autonomous-system'
func (estix *ElementSTIXObject) GetAutonomousSystemCyberObservableObjectSTIX() *AutonomousSystemCyberObservableObjectSTIX {
	if result, ok := estix.Data.(AutonomousSystemCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetDirectoryCyberObservableObjectSTIX возвращает объект STIX типа 'directory'
func (estix *ElementSTIXObject) GetDirectoryCyberObservableObjectSTIX() *DirectoryCyberObservableObjectSTIX {
	if result, ok := estix.Data.(DirectoryCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetDomainNameCyberObservableObjectSTIX возвращает объект STIX типа 'domain-name'
func (estix *ElementSTIXObject) GetDomainNameCyberObservableObjectSTIX() *DomainNameCyberObservableObjectSTIX {
	if result, ok := estix.Data.(DomainNameCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetEmailAddressCyberObservableObjectSTIX возвращает объект STIX типа 'email-addr'
func (estix *ElementSTIXObject) GetEmailAddressCyberObservableObjectSTIX() *EmailAddressCyberObservableObjectSTIX {
	if result, ok := estix.Data.(EmailAddressCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetEmailMessageCyberObservableObjectSTIX возвращает объект STIX типа 'email-message'
func (estix *ElementSTIXObject) GetEmailMessageCyberObservableObjectSTIX() *EmailMessageCyberObservableObjectSTIX {
	if result, ok := estix.Data.(EmailMessageCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetFileCyberObservableObjectSTIX возвращает объект STIX типа 'file'
func (estix *ElementSTIXObject) GetFileCyberObservableObjectSTIX() *FileCyberObservableObjectSTIX {
	if result, ok := estix.Data.(FileCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetIPv4AddressCyberObservableObjectSTIX возвращает объект STIX типа 'ipv4-addr'
func (estix *ElementSTIXObject) GetIPv4AddressCyberObservableObjectSTIX() *IPv4AddressCyberObservableObjectSTIX {
	if result, ok := estix.Data.(IPv4AddressCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetIPv6AddressCyberObservableObjectSTIX возвращает объект STIX типа 'ipv6-addr'
func (estix *ElementSTIXObject) GetIPv6AddressCyberObservableObjectSTIX() *IPv6AddressCyberObservableObjectSTIX {
	if result, ok := estix.Data.(IPv6AddressCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetMACAddressCyberObservableObjectSTIX возвращает объект STIX типа 'mac-addr'
func (estix *ElementSTIXObject) GetMACAddressCyberObservableObjectSTIX() *MACAddressCyberObservableObjectSTIX {
	if result, ok := estix.Data.(MACAddressCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetMutexCyberObservableObjectSTIX возвращает объект STIX типа 'mutex'
func (estix *ElementSTIXObject) GetMutexCyberObservableObjectSTIX() *MutexCyberObservableObjectSTIX {
	if result, ok := estix.Data.(MutexCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetNetworkTrafficCyberObservableObjectSTIX возвращает объект STIX типа 'network-traffic'
func (estix *ElementSTIXObject) GetNetworkTrafficCyberObservableObjectSTIX() *NetworkTrafficCyberObservableObjectSTIX {
	if result, ok := estix.Data.(NetworkTrafficCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetProcessCyberObservableObjectSTIX возвращает объект STIX типа 'process'
func (estix *ElementSTIXObject) GetProcessCyberObservableObjectSTIX() *ProcessCyberObservableObjectSTIX {
	if result, ok := estix.Data.(ProcessCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetSoftwareCyberObservableObjectSTIX возвращает объект STIX типа 'software'
func (estix *ElementSTIXObject) GetSoftwareCyberObservableObjectSTIX() *SoftwareCyberObservableObjectSTIX {
	if result, ok := estix.Data.(SoftwareCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetURLCyberObservableObjectSTIX возвращает объект STIX типа 'url'
func (estix *ElementSTIXObject) GetURLCyberObservableObjectSTIX() *URLCyberObservableObjectSTIX {
	if result, ok := estix.Data.(URLCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetUserAccountCyberObservableObjectSTIX возвращает объект STIX типа 'user-account'
func (estix *ElementSTIXObject) GetUserAccountCyberObservableObjectSTIX() *UserAccountCyberObservableObjectSTIX {
	if result, ok := estix.Data.(UserAccountCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetWindowsRegistryKeyCyberObservableObjectSTIX возвращает объект STIX типа 'windows-registry-key'
func (estix *ElementSTIXObject) GetWindowsRegistryKeyCyberObservableObjectSTIX() *WindowsRegistryKeyCyberObservableObjectSTIX {
	if result, ok := estix.Data.(WindowsRegistryKeyCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

// GetX509CertificateCyberObservableObjectSTIX возвращает объект STIX типа 'x509-certificate'
func (estix *ElementSTIXObject) GetX509CertificateCyberObservableObjectSTIX() *X509CertificateCyberObservableObjectSTIX {
	if result, ok := estix.Data.(X509CertificateCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}
