package mongodbapi

// GetChanInput возвращает канал для приема данных В модуле
func (mdb *MongoDBModule) GetChanInput() <-chan ChanInput {
	return mdb.chanInputToModule
}

// GetChanOutput возвращает канал для приема данных ИЗ модуля
func (mdb *MongoDBModule) GetChanOutput() <-chan ChanOutput {
	return mdb.chanOutputFromModule
}

// SendingDataToModule отправляет данные В модуль
func (mdb *MongoDBModule) SendingDataToModule(data ChanInput) {
	mdb.chanInputToModule <- data
}

// SendingDataFromModule отправляет данные ИЗ модуля
func (mdb *MongoDBModule) SendingDataFromModule(data ChanOutput) {
	mdb.chanOutputFromModule <- data
}
