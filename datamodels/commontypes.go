package datamodels

import "time"

// MessageLogging содержит информацию используемую при логировании
// MsgData - сообщение
// MsgType - тип сообщения
type MessageLogging struct {
	MsgData, MsgType string
}

// DataCounterSettings содержит информацию для подсчета
type DataCounterSettings struct {
	DataType string
	DataMsg  string
	Count    int
}

// ChanOutputDecodeJSON содержит данные получаемые при декодировании JSON формата
// обрабатываемого обработчиком HandlerMessageFromHive
// ExclusionRuleWorked - информирует что сработало правило исключения значения из списка
// передаваемых данных
// UUID - уникальный идентификатор в формате UUID
// FieldName - наименование поля
// ValueType - тип передаваемого значения (string, int и т.д.)
// Value - любые передаваемые данные
// FieldBranch - 'путь' до значения в как в JSON формате, например 'event.details.customFields.class'
type ChanOutputDecodeJSON struct {
	ExclusionRuleWorked bool
	UUID                string
	FieldName           string
	ValueType           string
	Value               interface{}
	FieldBranch         string
}

// CommonObservableType общие поля наблюдаемого сообщения
// Ioc - индикатор компрометации
// Sighted - видящий
// IgnoreSimilarity - игнорировать сходство
// Tlp - tlp
// UnderliningCreatedAt - время создания
// UnderliningUpdatedAt - время обновления
// StartDate - дата начала
// UnderliningCreatedBy - кем создан
// UnderliningUpdatedBy - кем обновлен
// UnderliningId - уникальный идентификатор
// UnderliningType - тип
// Data - данные
// DataType - тип данных
// Message - сообщение
type CommonObservableType struct {
	Ioc                  bool   `json:"ioc,omitempty" bson:"ioc"`
	Sighted              bool   `json:"sighted,omitempty" bson:"sighted"`
	IgnoreSimilarity     bool   `json:"ignoreSimilarity,omitempty" bson:"ignoreSimilarity"`
	Tlp                  uint64 `json:"tlp,omitempty" bson:"tlp"`
	UnderliningCreatedAt string `json:"_createdAt,omitempty" bson:"_createdAt"` //формат RFC3339
	UnderliningUpdatedAt string `json:"_updatedAt,omitempty" bson:"_updatedAt"` //формат RFC3339
	StartDate            string `json:"startDate,omitempty" bson:"startDate"`   //формат RFC3339
	UnderliningCreatedBy string `json:"_createdBy,omitempty" bson:"_createdBy"`
	UnderliningUpdatedBy string `json:"_updatedBy,omitempty" bson:"_updatedBy"`
	UnderliningId        string `json:"_id,omitempty" bson:"_id"`
	UnderliningType      string `json:"_type,omitempty" bson:"_type"`
	Data                 string `json:"data,omitempty" bson:"data"`
	DataType             string `json:"dataType,omitempty" bson:"dataType"`
	Message              string `json:"message,omitempty" bson:"message"`
}

// DifferentObjectType содержит перечисление полей и их значения, которые были изменены в произвольном типе
// SourceReceivingChanges - источник от которого были получены изменения
// ModifiedTime - время выполнения модификации
// UserNameModifiedObject - пользователь выполнивший модификацию
// CollectionName - наименование коллекции в которой выполнялись модификации
// DocumentID - идентификатор документа в котором выполнялись модификации
// FieldList - перечень полей подвергшихся изменениям
type DifferentObjectType struct {
	SourceReceivingChanges string                    `json:"source_receiving_changes" bson:"source_receiving_changes"`
	ModifiedTime           time.Time                 `json:"modified_time" bson:"modified_time"`
	UserNameModifiedObject string                    `json:"user_name_modified_object" bson:"user_name_modified_object"`
	CollectionName         string                    `json:"collection_name" bson:"collection_name"`
	DocumentID             string                    `json:"document_id" bson:"document_id"`
	FieldList              []OldFieldValueObjectType `json:"field_list" bson:"field_list"`
}

// OldFieldValueObjectType содержит старое значение полей, до их модификации
// FeildType - тип поля
// Path - полный путь к объекту подвергшемуся модификации
// Value - предыдущее значение поля, которое подверглось модификации
type OldFieldValueObjectType struct {
	FeildType string      `json:"feild_type" bson:"feild_type"`
	Path      string      `json:"path" bson:"path"`
	Value     interface{} `json:"value" bson:"value"`
}

// ElementSTIXObject может содержать любой из STIX объектов с указанием его типа
// DataType - тип STIX объекта
// Data - непосредственно сам STIX объект
type ElementSTIXObject struct {
	DataType string
	Data     HandlerSTIXObject
}

// HandlerSTIXObject набор интерфейсов реализующих различные обработчики
type HandlerSTIXObject interface {
	GetterObjectSTIX
	GetterID
	GetterType
	//	ComparatorSTIXObject ДЛЯ СРАВНЕНИЯ ОБЪЕКТОВ
	ToStringBeautifulReader
}

// GetterObjectSTIX интерфейс реализующий метод применяемый для получения любого объекта
type GetterObjectSTIX interface {
	GetObject() interface{}
}

// GetterID интерфейс реализующий реализующий метод применяемый для получения id объекта
type GetterID interface {
	GetID() string
}

// GetterType интерфейс реализующий реализующий метод применяемый для получения типа объекта
type GetterType interface {
	GetType() string
}

// ToStringBeautifulReader интерфейс реализующий метод для вывода объекта в строковом виде
type ToStringBeautifulReader interface {
	ToStringBeautiful(int) string
}

// ComparatorSTIXObject интерфейс реализующий обработчик для сравнения STIX объектов одного типа
type ComparatorSTIXObject interface {
	ComparisonTypeCommonFields(interface{}, string) (bool, DifferentObjectType, error)
}
