package memorytemporarystorage

import (
	"sync"
	"sync/atomic"
	"time"
)

var once sync.Once
var cst CommonStorageTemporary

func NewTemporaryStorage() *CommonStorageTemporary {
	once.Do(func() {
		cst = CommonStorageTemporary{}
	})

	return &cst
}

func (cst *CommonStorageTemporary) GetAlertCounter() StorageCounter {
	return cst.alertCounter
}

func (cst *CommonStorageTemporary) GetCaseCounter() StorageCounter {
	return cst.caseCounter
}

func (cst *CommonStorageTemporary) IncrementAcceptedEvents() {
	atomic.AddUint64(&cst.dataCounter.acceptedEvents, 1)
}

func (cst *CommonStorageTemporary) GetAcceptedEvents() uint64 {
	return cst.dataCounter.acceptedEvents
}

func (cst *CommonStorageTemporary) IncrementProcessedEvents() {
	atomic.AddUint64(&cst.dataCounter.processedEvents, 1)
}

func (cst *CommonStorageTemporary) GetProcessedEvents() uint64 {
	return cst.dataCounter.processedEvents
}

func (cst *CommonStorageTemporary) IncrementAlertInsertMongoDB() {
	atomic.AddUint64(&cst.alertCounter.insertMongoDB, 1)
}

func (cst *CommonStorageTemporary) GetAlertInsertMongoDB() uint64 {
	return cst.alertCounter.insertMongoDB
}

func (cst *CommonStorageTemporary) IncrementCaseInsertMongoDB() {
	atomic.AddUint64(&cst.caseCounter.insertMongoDB, 1)
}

func (cst *CommonStorageTemporary) GetCaseInsertMongoDB() uint64 {
	return cst.caseCounter.insertMongoDB
}

// SetStartTimeDataCounter добавляет время начала сетчика
func (cst *CommonStorageTemporary) SetStartTimeDataCounter(t time.Time) {
	cst.dataCounter.startTime = t
}

// GetStartTimeDataCounter возвращает время начала сетчика
func (cst *CommonStorageTemporary) GetStartTimeDataCounter() time.Time {
	return cst.dataCounter.startTime
}
