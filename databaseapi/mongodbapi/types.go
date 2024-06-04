package mongodbapi

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBModule содержит описание каналов для взаимодействия с БД MongoDB
// ChanInputToModule - канал для отправки данных В модуль
// ChanOutputFromModule - канал для приема данных ИЗ модуля
type MongoDBModule struct {
	chanInputToModule    chan ChanInput
	chanOutputFromModule chan ChanOutput
}

// ConnectionDescriptorMongoDB дескриптор соединения с БД MongoDB
// databaseName - имя базы данных
// connection - дескриптор соединения
// ctx - контекст переносит крайний срок, сигнал отмены и другие значения через границы API
// ctxCancel - метод закрытия контекста
// ctxRoute - контекст для роута
type ConnectionDescriptorMongoDB struct {
	databaseName string
	connection   *mongo.Client
	ctx          context.Context
	ctxCancel    context.CancelFunc
	ctxRoute     context.Context
}
