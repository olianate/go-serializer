package serializer

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/olianate/go-serializer/formatter"
)

func Serialize(data interface{}) interface{} {
	if data == nil {
		return data
	}
	v := reflect.ValueOf(data)
	if v.IsZero() {
		return data
	}

	t := reflect.TypeOf(data)
	// return value if implemented json.Marshaler
	if t.Implements(reflect.TypeOf((*json.Marshaler)(nil)).Elem()) {
		return data
	}

	switch t.Kind() {
	case reflect.Ptr:
		return Serialize(v.Elem().Interface())

	case reflect.Struct:
		valuesMap := make(map[string]interface{})

		numOfFilelds := t.NumField()
		for i := 0; i < numOfFilelds; i++ {
			fieldName := t.Field(i).Name
			if !startWithSupperCase(fieldName) {
				continue
			}

			tag := t.Field(i).Tag
			jsonTag := strings.Split(tag.Get("json"), ",")

			if len(jsonTag) > 0 {
				if jsonTag[0] == "-" {
					continue
				}
				fieldName = orString(jsonTag[0], fieldName)
			}

			f := tag.Get("format")
			var retval interface{}
			if formatter, ok := formatter.Formatters[f]; ok {
				retval = formatter(v.Field(i).Interface())
			} else {
				retval = Serialize(v.Field(i).Interface())
			}

			if retmap, ok := retval.(map[string]interface{}); t.Field(i).Anonymous && ok {
				for key, val := range retmap {
					valuesMap[key] = val
				}
			} else if !(reflect.ValueOf(retval).IsZero() && len(jsonTag) > 1 && jsonTag[1] == "omitempty") {
				valuesMap[fieldName] = retval
			}

		}

		return valuesMap

	case reflect.Slice:
		length := v.Len()
		res := make([]interface{}, length)
		for i := 0; i < length; i++ {
			res[i] = Serialize(v.Index(i).Interface())
		}

		return res

	default:
		return data

	}

}
