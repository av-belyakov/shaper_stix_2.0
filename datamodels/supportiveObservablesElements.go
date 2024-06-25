package datamodels

import "strings"

// SupportiveObservables вспомогательный тип для для обработки observables
type SupportiveObservables struct {
	previousFieldReports      string
	listAcceptedFields        []string
	listAcceptedFieldsReports []string
	observableTmp             ObservableMessage
	observables               []ObservableMessage
}

// NewSupportiveObservables формирует вспомогательный объект для обработки
// thehive объектов типа observables
func NewSupportiveObservables() *SupportiveObservables {
	return &SupportiveObservables{
		listAcceptedFields:        []string(nil),
		listAcceptedFieldsReports: []string(nil),
		observableTmp:             *NewObservableMessage(),
		observables:               []ObservableMessage(nil)}
}

// GetObservables возвращает []datamodels.ObservableMessage, однако, метод
// выполняет еще очень важное действие, перемещает содержимое из o.observableTmp в
// список o.observables, так как observables автоматически пополняется только при
// совпадении значений в listAcceptedFields. Соответственно при завершении
// JSON объекта, последние добавленные значения остаются observableTmp
func (o *SupportiveObservables) GetObservables() []ObservableMessage {
	o.listAcceptedFields = []string(nil)
	o.listAcceptedFieldsReports = []string(nil)

	if o.observableTmp.DataType != "" {
		o.observables = append(o.observables, o.observableTmp)
	}

	return o.observables
}

// GetObservableTmp возвращает временный объект observable
func (o *SupportiveObservables) GetObservableTmp() *ObservableMessage {
	return &o.observableTmp
}

func (o *SupportiveObservables) HandlerValue(fieldBranch string, i interface{}, f func(interface{})) {
	//если поле повторяется то считается что это уже новый объект
	if fieldBranch != "observables.tags" && fieldBranch != "observables.attachment.hashes" && o.isExistFieldBranch(fieldBranch) {
		o.listAcceptedFields = []string(nil)
		o.listAcceptedFieldsReports = []string(nil)
		o.previousFieldReports = ""
		o.observables = append(o.observables, o.observableTmp)
		o.observableTmp = ObservableMessage{
			CommonObservableType: CommonObservableType{
				StartDate:            "1970-01-01T00:00:00+00:00",
				UnderliningCreatedAt: "1970-01-01T00:00:00+00:00",
				UnderliningUpdatedAt: "1970-01-01T00:00:00+00:00",
			},
			Reports: map[string]ReportTaxonomies{},
		}
	}

	o.listAcceptedFields = append(o.listAcceptedFields, fieldBranch)

	f(i)
}

func (o *SupportiveObservables) isExistFieldBranch(value string) bool {
	for _, v := range o.listAcceptedFields {
		if v == value {
			return true
		}
	}

	return false
}

func (o *SupportiveObservables) HandlerReportValue(fieldBranch string, i interface{}) {
	fields := strings.Split(fieldBranch, ".")
	if len(fields) != 5 {
		return
	}

	//пока обрабатываем только taxonomies
	if fields[3] != "taxonomies" {
		return
	}

	if _, ok := o.observableTmp.Reports[fields[2]]; !ok {
		o.observableTmp.Reports[fields[2]] = ReportTaxonomies{Taxonomies: make([]Taxonomy, 1)}
		o.previousFieldReports = fields[2]
		o.listAcceptedFieldsReports = []string{}
	}

	//для того чтобы понять нужно ли создавать новый элемент среза
	//используем хранилище listAcceptedFields для временного хранения
	//наименований полей, создаем новый элемент среза, если попадается
	// повторяющееся свойство структуры Taxonomy
	if o.previousFieldReports == fields[2] && o.isExistFieldBranchReports(fields[4]) {
		tmpSlice := o.observableTmp.Reports[fields[2]]
		tmpSlice.Taxonomies = append(tmpSlice.Taxonomies, Taxonomy{})
		o.observableTmp.Reports[fields[2]] = tmpSlice

		o.listAcceptedFieldsReports = []string{}
	}

	o.listAcceptedFieldsReports = append(o.listAcceptedFieldsReports, fields[4])
	lastNum := len(o.observableTmp.Reports[fields[2]].Taxonomies) - 1
	if lastNum < 0 {
		lastNum = 0
	}

	switch fields[4] {
	case "level":
		o.observableTmp.Reports[fields[2]].Taxonomies[lastNum].SetAnyLevel(i)

	case "namespace":
		o.observableTmp.Reports[fields[2]].Taxonomies[lastNum].SetAnyNamespace(i)

	case "predicate":
		o.observableTmp.Reports[fields[2]].Taxonomies[lastNum].SetAnyPredicate(i)

	case "value":
		o.observableTmp.Reports[fields[2]].Taxonomies[lastNum].SetAnyValue(i)
	}
}

func (o *SupportiveObservables) isExistFieldBranchReports(value string) bool {
	for _, v := range o.listAcceptedFieldsReports {
		if v == value {
			return true
		}
	}

	return false
}
