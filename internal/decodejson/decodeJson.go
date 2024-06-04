package decodejson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"

	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

type DecodeJsonMessageSettings struct {
	Logging  chan<- datamodels.MessageLogging
	Counting chan<- datamodels.DataCounterSettings
}

func NewDecodeJsonMessageSettings(
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) *DecodeJsonMessageSettings {
	return &DecodeJsonMessageSettings{
		Logging:  logging,
		Counting: counting,
	}
}

func (s *DecodeJsonMessageSettings) HandlerJsonMessage(b []byte, id, subject string) (chan datamodels.ChanOutputDecodeJSON, chan bool) {
	chanOutputJsonData := make(chan datamodels.ChanOutputDecodeJSON)
	chanDone := make(chan bool)

	//ПРЕДНАЗНАЧЕНО для записи принимаемых объектов в лог-файл
	str, err := supportingfunctions.NewReadReflectJSONSprint(b)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		s.Logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l+2),
			MsgType: "error",
		}
	}

	s.Logging <- datamodels.MessageLogging{
		MsgData: fmt.Sprintf("\t---------------\n%s\n", str),
		MsgType: "objects",
	}

	go func() {
		var (
			f         string
			l         int
			err       error
			isAllowed bool
		)

		//для карт
		_, f, l, _ = runtime.Caller(0)
		listMap := map[string]interface{}{}
		if err = json.Unmarshal(b, &listMap); err == nil {
			if len(listMap) == 0 {
				s.Logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'error decoding the json message, it may be empty' %s:%d", f, l+2),
					MsgType: "error",
				}

				return
			}

			_ = reflectMap(chanOutputJsonData, listMap, 0, "", id)
		} else {
			// для срезов
			_, f, l, _ = runtime.Caller(0)
			listSlice := []interface{}{}
			if err = json.Unmarshal(b, &listSlice); err != nil {
				s.Logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l+2),
					MsgType: "error",
				}

				return
			}

			if len(listSlice) == 0 {
				s.Logging <- datamodels.MessageLogging{
					MsgData: fmt.Sprintf("'error decoding the json message, it may be empty' %s:%d", f, l+2),
					MsgType: "error",
				}

				return
			}

			_ = reflectSlice(chanOutputJsonData, listSlice, 0, "", id)
		}

		close(chanOutputJsonData)

		// сетчик обработаных сообщений
		s.Counting <- datamodels.DataCounterSettings{
			DataType: "update processed events",
			DataMsg:  subject,
			Count:    1,
		}

		//останавливаем обработчик формирующий верифицированный объект
		chanDone <- isAllowed

		close(chanDone)
	}()

	return chanOutputJsonData, chanDone
}

func reflectAnySimpleType(
	chanOutMispFormat chan<- datamodels.ChanOutputDecodeJSON,
	name interface{},
	anyType interface{},
	fieldBranch string,
	id string) interface{} {

	var nameStr string
	r := reflect.TypeOf(anyType)

	if n, ok := name.(int); ok {
		nameStr = fmt.Sprint(n)
	} else if n, ok := name.(string); ok {
		nameStr = n
	}

	if r == nil {
		return anyType
	}

	switch r.Kind() {
	case reflect.String:
		result := reflect.ValueOf(anyType).String()
		chanOutMispFormat <- datamodels.ChanOutputDecodeJSON{
			UUID:        id,
			FieldName:   nameStr,
			ValueType:   "string",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		result := reflect.ValueOf(anyType).Int()
		chanOutMispFormat <- datamodels.ChanOutputDecodeJSON{
			UUID:        id,
			FieldName:   nameStr,
			ValueType:   "int",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result := reflect.ValueOf(anyType).Uint()
		chanOutMispFormat <- datamodels.ChanOutputDecodeJSON{
			UUID:        id,
			FieldName:   nameStr,
			ValueType:   "uint",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Float32, reflect.Float64:
		result := reflect.ValueOf(anyType).Float()
		chanOutMispFormat <- datamodels.ChanOutputDecodeJSON{
			UUID:        id,
			FieldName:   nameStr,
			ValueType:   "float",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Bool:
		result := reflect.ValueOf(anyType).Bool()
		chanOutMispFormat <- datamodels.ChanOutputDecodeJSON{
			UUID:        id,
			FieldName:   nameStr,
			ValueType:   "bool",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result
	}

	return anyType
}

func reflectMap(
	chanOutMispFormat chan<- datamodels.ChanOutputDecodeJSON,
	l map[string]interface{},
	num int,
	fieldBranch string,
	id string) map[string]interface{} {

	var (
		newMap  map[string]interface{}
		newList []interface{}
	)
	nl := map[string]interface{}{}

	for k, v := range l {
		var fbTmp string
		r := reflect.TypeOf(v)

		if r == nil {
			return nl
		}

		fbTmp = fieldBranch
		if fbTmp == "" {
			fbTmp += k
		} else {
			fbTmp += "." + k
		}

		switch r.Kind() {
		case reflect.Map:
			if v, ok := v.(map[string]interface{}); ok {
				newMap = reflectMap(chanOutMispFormat, v, num+1, fbTmp, id)
				nl[k] = newMap
			}

		case reflect.Slice:
			if v, ok := v.([]interface{}); ok {
				newList = reflectSlice(chanOutMispFormat, v, num+1, fbTmp, id)
				nl[k] = newList
			}

		default:
			nl[k] = reflectAnySimpleType(chanOutMispFormat, k, v, fbTmp, id)
		}
	}

	return nl
}

func reflectSlice(
	chanOutMispFormat chan<- datamodels.ChanOutputDecodeJSON,
	l []interface{},
	num int,
	fieldBranch string,
	id string) []interface{} {

	var (
		newMap  map[string]interface{}
		newList []interface{}
	)
	nl := make([]interface{}, 0, len(l))

	for k, v := range l {
		r := reflect.TypeOf(v)

		if r == nil {
			return nl
		}

		switch r.Kind() {
		case reflect.Map:
			if v, ok := v.(map[string]interface{}); ok {
				newMap = reflectMap(chanOutMispFormat, v, num+1, fieldBranch, id)

				nl = append(nl, newMap)
			}

		case reflect.Slice:
			if v, ok := v.([]interface{}); ok {
				newList = reflectSlice(chanOutMispFormat, v, num+1, fieldBranch, id)

				nl = append(nl, newList...)
			}

		default:
			nl = append(nl, reflectAnySimpleType(chanOutMispFormat, k, v, fieldBranch, id))
		}
	}

	return nl
}
