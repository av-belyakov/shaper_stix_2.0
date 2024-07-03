package mongodbapi

// GetChanInput возвращает канал для приема данных В модуле
func (mdb *MongoDBModule) GetChanInput() <-chan ChanInput {
	return mdb.ChanInputToModule
}

// GetChanOutput возвращает канал для приема данных ИЗ модуля
func (mdb *MongoDBModule) GetChanOutput() <-chan ChanOutput {
	return mdb.ChanOutputFromModule
}

// SendingDataToModule отправляет данные В модуль
func (mdb *MongoDBModule) SendingDataToModule(data ChanInput) {
	mdb.ChanInputToModule <- data
}

// SendingDataFromModule отправляет данные ИЗ модуля
func (mdb *MongoDBModule) SendingDataFromModule(data ChanOutput) {
	mdb.ChanOutputFromModule <- data
}
