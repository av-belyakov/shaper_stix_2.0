package mongodbapi

import (
	"fmt"

	"shaper_stix/datamodels"
)

// AddAttackPatternDO объект "Attack Pattern", по терминалогии STIX, описывающий способы
// компрометации цели
func (w *Wrappers) AddAttackPatternDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Attack-Pattern
	//****************************************************************
	fmt.Println("Additing Attack-Pattern Domain Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "AttackPattern",
	}
}

// AddCampaignDO объект "Campaign", по терминалогии STIX, это набор действий определяющих
// злонамеренную деятельность или атаки
func (w *Wrappers) AddCampaignDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Campaign
	//****************************************************************
	fmt.Println("Additing Campaign Domain Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Campaign",
	}
}

// AddCourseOfActionDO объект "Course of Action", по терминалогии STIX, описывающий
// совокупность действий
func (w *Wrappers) AddCourseOfActionDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа CourseOfAction
	//****************************************************************
	fmt.Println("Additing CourseOfAction Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "CourseOfAction",
	}
}

// AddGroupingDO объект "Grouping", по терминалогии STIX, объединяет различные объекты
// STIX в рамках какого то общего контекста
func (w *Wrappers) AddGroupingDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Grouping
	//****************************************************************
	fmt.Println("Additing Grouping Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Grouping",
	}
}

// AddIdentityDO объект "Identity", по терминалогии STIX, содержит основную идентификационную
// информацию физичиских лиц, организаций и т.д.
func (w *Wrappers) AddIdentityDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Identity
	//****************************************************************
	fmt.Println("Additing Identity Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Identity",
	}
}

// AddIndicatorDO объект "Indicator", по терминалогии STIX, содержит шаблон который может быть
// использован для обнаружения подозрительной или вредоносной киберактивности
func (w *Wrappers) AddIndicatorDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Indicator
	//****************************************************************
	fmt.Println("Additing Indicator Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Indicator",
	}
}

// AddInfrastructureDO объект "Infrastructure", по терминалогии STIX, содержит описание
// любых систем, программных служб, а так же любые связанные с ними физические или
// виртуальные ресурсы, предназначенные для поддержки какой-либо цели
func (w *Wrappers) AddInfrastructureDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Infrastructure
	//****************************************************************
	fmt.Println("Additing Infrastructure Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Infrastructure",
	}
}

// AddIntrusionSetDO объект "Intrusion Set", по терминалогии STIX, содержит сгруппированный
// набор враждебного поведения и ресурсов с общими свойствами, который, как считается,
// управляется одной организацией
func (w *Wrappers) AddIntrusionSetDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа IntrusionSet
	//****************************************************************
	fmt.Println("Additing IntrusionSet Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "IntrusionSet",
	}
}

// AddLocationDO объект "Location", по терминалогии STIX, содержит описание географического
// местоположения
func (w *Wrappers) AddLocationDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Location
	//****************************************************************
	fmt.Println("Additing Location Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Location",
	}
}

// AddMalwareDO объект "Malware", по терминалогии STIX, содержит подробную информацию о
// функционировании вредоносной программы
func (w *Wrappers) AddMalwareDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Malware
	//****************************************************************
	fmt.Println("Additing Malware Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Malware",
	}
}

// AddMalwareAnalysisDO объект "Malware Analysis", по терминалогии STIX, содержит анализ
// вредоносных программ захватывающих метаданные и результаты конкретного статического
// или динамического анализа, выполненного на экземпляре вредоносного ПО или семействе
// вредоносных программ
func (w *Wrappers) AddMalwareAnalysisDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа MalwareAnalysis
	//****************************************************************
	fmt.Println("Additing MalwareAnalysis Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "MalwareAnalysis",
	}
}

// AddNoteDO объект "Note", по терминалогии STIX, содержит текстовую информации дополняющую
// текущий контекст анализа либо содержащей результаты дополнительного анализа которые не
// может быть описан в терминах объектов STIX
func (w *Wrappers) AddNoteDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Note
	//****************************************************************
	fmt.Println("Additing Note Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Note",
	}
}

// AddObservedDataDO объект "Observed Data", по терминалогии STIX, содержит информацию о
// сущностях связанных с кибер безопасностью, таких как файлы, системы или сети.
// Наблюдаемые данные это не результат анализа или заключение искусственного интеллекта,
// это просто сырая информация без какого-либо контекста.
func (w *Wrappers) AddObservedDataDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа ObservedData
	//****************************************************************
	fmt.Println("Additing ObservedData Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "ObservedData",
	}
}

// AddOpinionDO объект "Opinion", по терминалогии STIX, содержит оценку информации в
// приведенной в каком либо другом объекте STIX, которую произвел другой участник
// анализа.
func (w *Wrappers) AddOpinionDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Opinion
	//****************************************************************
	fmt.Println("Additing Opinion Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Opinion",
	}
}

// AddReportDO объект "Report", по терминалогии STIX, содержит совокупность данных об
// угрозах, сосредоточенных на одной или нескольких темах, таких как описание
// исполнителя, вредоносного ПО или метода атаки, включая контекст и связанные с ним
// детали. Применяется для группировки информации связанной с кибер угрозой. Может
// быть использован для дальнейшей публикации данной информации как истории
// расследования.
func (w *Wrappers) AddReportDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Report
	//****************************************************************
	fmt.Println("Additing Report Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Report",
	}
}

// AddThreatActorDO объект "Threat Actor", по терминалогии STIX, содержит информацию о
// физических лицах или их группах и организациях которые могут действовать со злым
// умыслом.
func (w *Wrappers) AddThreatActorDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа ThreatActor
	//****************************************************************
	fmt.Println("Additing ThreatActor Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "ThreatActor",
	}
}

// AddToolDO объект "Tool", по терминалогии STIX, содержит информацию о легитимном
// ПО которое может быть использованно для реализации компьютерных угроз
func (w *Wrappers) AddToolDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Tool
	//****************************************************************
	fmt.Println("Additing Tool Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Tool",
	}
}

// AddVulnerabilityDO объект "Vulnerability", по терминологии STIX, содержит описание
// уязвимостей полученных в результате неверной формализации требований, ошибочном
// проектировании или некорректной реализации программного кода или логики в ПО, а
// также в компонентах оборудования
func (w *Wrappers) AddVulnerabilityDO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Vulnerability
	//****************************************************************
	fmt.Println("Additing Vulnerability Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "Vulnerability",
	}
}

/*
// AddNewCase добавляет новый кейс в БД
func (w *wrappers) AddNewCase(
	//data *datamodels.VerifiedTheHiveCase,
	data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	qp := QueryParameters{
		NameDB:         w.NameDB,
		ConnectDB:      w.ConnDB,
		CollectionName: "case_collection",
	}

	obj, ok := data.(*datamodels.VerifiedTheHiveCase)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type *datamodels.VerifiedTheHiveCase' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	//******************************************************************************
	//ищем документы подходящие под фильтр и удаляем их что бы избежать дублирования
	cur, err := qp.Find(bson.D{
		{Key: "source", Value: obj.GetSource()},
		{Key: "event.rootId", Value: obj.GetEvent().GetRootId()},
	})
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	listForDelete := []string{}
	for cur.Next(context.Background()) {
		var modelType struct {
			ID     string `bson:"@id"`
			Source string `bson:"source"`
		}

		if err := cur.Decode(&modelType); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}

			continue
		}

		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("MongoDB , document with ID:'%s', author: '%s' data will be deleted", modelType.ID, modelType.Source),
			MsgType: "warning",
		}

		listForDelete = append(listForDelete, modelType.ID)
	}

	if len(listForDelete) > 0 {
		if _, err := qp.DeleteManyData(
			bson.D{{
				Key:   "@id",
				Value: bson.D{{Key: "$in", Value: listForDelete}}}},
			options.Delete(),
		); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-1),
				MsgType: "error",
			}
		}
	}
	//******************************************************************************

	if _, err := qp.InsertData([]interface{}{data}, []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "@id", Value: 1},
			},
			Options: &options.IndexOptions{},
		}, {
			Keys: bson.D{
				{Key: "source", Value: 1},
				{Key: "event.rootId", Value: 1},
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

// AddNewCase выполняет замену уже существующего объекта типа Alert
// либо добавляет новый, если в БД нет объекта с заданными параметрами
func (w *wrappers) AddNewAlert(
	data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	qp := QueryParameters{
		NameDB:         w.NameDB,
		ConnectDB:      w.ConnDB,
		CollectionName: "alert_collection",
	}

	obj, ok := data.(*datamodels.VerifiedTheHiveAlert)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type *datamodels.VerifiedTheHiveCase' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	//поиск схожего документа
	currentData := datamodels.VerifiedTheHiveAlert{}
	err := qp.FindOne(bson.D{
		{Key: "source", Value: obj.GetSource()},
		{Key: "event.rootId", Value: obj.GetEvent().GetRootId()},
	}).Decode(&currentData)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		//если похожего документа нет в БД
		currentData = *obj
	} else {
		//если похожий документ есть, удаляем старый документ и выполняем
		//замену старых значений в полученном из БД документе новыми значениями
		if _, err := qp.DeleteManyData(
			bson.D{{
				Key:   "@id",
				Value: bson.D{{Key: "$in", Value: []string{obj.GetID()}}}}},
			options.Delete(),
		); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-1),
				MsgType: "error",
			}

			return
		}

		if _, err := currentData.GetEvent().ReplacingOldValues(*obj.GetEvent()); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("error replacing old values event '%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		}
		if _, err := currentData.GetAlert().ReplacingOldValues(*obj.GetAlert()); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("error replacing old values alert '%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		}
	}

	if _, err := qp.InsertData([]interface{}{currentData},
		[]mongo.IndexModel{
			{
				Keys: bson.D{
					{Key: "@id", Value: 1},
				},
				Options: &options.IndexOptions{},
			}, {
				Keys: bson.D{
					{Key: "source", Value: 1},
					{Key: "event.rootId", Value: 1},
				},
				Options: &options.IndexOptions{},
			}}); err != nil {
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
		DataMsg:  "subject_alert",
		Count:    1,
	}
}
*/
