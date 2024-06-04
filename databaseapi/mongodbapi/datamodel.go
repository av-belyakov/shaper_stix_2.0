package mongodbapi

// CommonChan общие поля пользовательской структуры
// Section - секция обработки данных
// Command - команда
// AppTaskID - внутренний идентификатор задачи
// ObjectId - идентификатор объекта
// ObjectType - тип объекта
type CommonChan struct {
	Section    string
	Command    string
	AppTaskId  string
	ObjectId   string
	ObjectType string
}

// ChanInput для данных передаваемых В модуль
type ChanInput struct {
	CommonChan
	Data interface{}
}

// ChanInputMongoDB для данных передаваемых ИЗ модуля
type ChanOutput struct {
	CommonChan
	Data interface{}
}
