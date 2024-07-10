package mongodbapi

import (
	"fmt"
	"runtime"

	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (w *Wrappers) AddNewSITXObject(
	data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	qp := QueryParameters{
		NameDB:         w.NameDB,
		ConnectDB:      w.ConnDB,
		CollectionName: "stix_object_collection",
	}

	//********************************************************************
	// В место простого Insert сделать следующее:
	// 1. Получить список STIX объектов по case id (для report), и по
	//common_outside_specification.element_id для всех остальных
	//
	// 2. Выполнить сравнение найденых объектов с имеющимися STIX объектами
	// 		ВАЖНО!!! STIX объекты при вставке в БД проходят пред обработку методом
	// MarshalBSON, где меняются такие типы как например, время. Объекты для сравнения
	// находятся в типе Wrapper<тип_STIX>, соответственно для корректного сравнения,
	// после получения данных из БД нужно сделать метод UnmarshalBSON, что бы привести
	// данные из типа FinalyIndicatorObject в Wrapper<тип_STIX>
	//
	// 2.1 Для объектов по которым не было найдено ни какой информации
	// создать STIX объекты типа !!!! relationship !!!!
	//
	// 2.2 Для объектов по которым была найдена какая то информация выполнить
	// сравнение с объектами из БД и обновить объекты из БД новыми данными,
	// ИСКЛЮЧАЯ изменение 'commonpropertiesobjectstix.id' что бы не потерять
	// связь с другими объектами в том числе через relationship
	//
	// 3. Выполнить update в БД для данных которые уже есть и insert для
	// данных которых еще нет в БД
	//
	//********************************************************************

	//создаем объект Relationship для установки обратной связи между
	//объектом Report и обрабатываемым объектом
	// relationship := methodstixobjects.NewRelationshipObjectSTIX()
	// relationship.SetValueID(fmt.Sprintf("relationship-%s", uuid.NewString()))
	// relationship.SetValueCreated(supportingfunctions.GetDateTimeFormatRFC3339(time.Now().UnixMilli()))
	//исходный объект, то есть обрабатываемый в настоящее время
	// relationship.SetValueSourceRef(stixhelpers.IdentifierTypeSTIX(newObject.GetID()))
	//целевой объект, то есть объект Report
	// relationship.SetValueTargetRef(stixhelpers.IdentifierTypeSTIX(reportWrap.GetID()))

	if _, err := qp.InsertData([]interface{}{data}, []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "commonpropertiesobjectstix.type", Value: 1},
				{Key: "commonpropertiesobjectstix.id", Value: 1},
				{Key: "outside_specification.case_id", Value: 1},
			},
			Options: &options.IndexOptions{},
		},
	}); err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	//счетчик
	counting <- datamodels.DataCounterSettings{
		DataType: "update count insert MongoDB",
		DataMsg:  "subject_case",
		Count:    1,
	}
}
