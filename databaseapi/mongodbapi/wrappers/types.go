package mongodbapi

import "go.mongodb.org/mongo-driver/mongo"

// Wrappers обертка для обработчиков
type Wrappers struct {
	AdditionalRequestParameters interface{}
	NameDB                      string
	ConnDB                      *mongo.Client
}
