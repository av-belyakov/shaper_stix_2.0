package datamodels

// *************************************************************
// *************   объект 'observables' из TheHive   ***********
// *************************************************************
// ObservablesMessageTheHive список наблюдаемых сообщений
// Observables - наблюдаемые сообщения
type ObservablesMessageTheHive struct {
	Observables []ObservableMessage `json:"observables" bson:"observables"`
}

// ObservableMessage наблюдаемое сообщение
// Tags - список тегов
// Attachment - приложенные данные
// Reports - список отчетов
type ObservableMessage struct {
	CommonObservableType
	Tags       []string                    `json:"tags" bson:"tags"`
	Attachment AttachmentData              `json:"attachment,omitempty" bson:"attachment"`
	Reports    map[string]ReportTaxonomies `json:"reports" bson:"reports"`
}

// ******************** AttachmentData ************************
// Size - размер
// Id - идентификатор
// Name - наименование
// ContentType - тип контента
// Hashes - список хешей
type AttachmentData struct {
	Size        uint64   `json:"size,omitempty" bson:"size"`
	Id          string   `json:"id,omitempty" bson:"id"`
	Name        string   `json:"name,omitempty" bson:"name"`
	ContentType string   `json:"contentType,omitempty" bson:"contentType"`
	Hashes      []string `json:"hashes,omitempty" bson:"hashes"`
}

// *************** ReportTaxonomies *****************
type ReportTaxonomies struct {
	Taxonomies []Taxonomy `json:"taxonomies,omitempty" bson:"taxonomies"`
}

// ******************* Taxonomy *******************
type Taxonomy struct {
	Level     string `json:"level,omitempty" bson:"level"`
	Namespace string `json:"namespace,omitempty" bson:"namespace"`
	Predicate string `json:"predicate,omitempty" bson:"predicate"`
	Value     string `json:"value,omitempty" bson:"value"`
}
