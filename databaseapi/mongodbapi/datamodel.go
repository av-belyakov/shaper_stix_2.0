package mongodbapi

// CommonChanMongoDB общие поля пользовательской структуры
// Section - секция обработки данных
// Command - команда
// AppTaskID - внутренний идентификатор задачи
// ObjectId - идентификатор объекта
// ObjectType - тип объекта
type CommonChanMongoDB struct {
	Section    string
	Command    string
	AppTaskId  string
	ObjectId   string
	ObjectType string
}

// ChanInputMongoDB для данных передаваемых В модуль
type ChanInputMongoDB struct {
	CommonChanMongoDB
	Data interface{}
}

// ChanInputMongoDB для данных передаваемых ИЗ модуля
type ChanOutputMongoDB struct {
	CommonChanMongoDB
	Data interface{}
}
