package mongodbapi

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBModule содержит описание каналов для взаимодействия с БД MongoDB
// ChanInputModule - канал для отправки данных В модуль
// ChanOutputModule - канал для принятия данных ИЗ модуля
type MongoDBModule struct {
	ChanInputModule  chan ChanInputMongoDB
	ChanOutputModule chan ChanOutputMongoDB
}

// ConnectionDescriptorMongoDB дескриптор соединения с БД MongoDB
// databaseName - имя базы данных
// connection - дескриптор соединения
// ctx - контекст переносит крайний срок, сигнал отмены и другие значения через границы API
// ctxCancel - метод закрытия контекста
type ConnectionDescriptorMongoDB struct {
	databaseName string
	connection   *mongo.Client
	ctx          context.Context
	ctxCancel    context.CancelFunc
}
