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
