package datamodels

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

type GetterID interface {
	GetID() string
}

type GetterType interface {
	GetType() string
}

type GetterCommonPropertiesObjectSTIX interface {
	GetterID
	GetterType
}
