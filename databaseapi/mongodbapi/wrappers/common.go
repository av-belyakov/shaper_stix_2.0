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
