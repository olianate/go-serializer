package serializer

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/olianate/go-serializer/formatter"
)

type Serializable interface {
	Serializer() interface{}
}

func Serialize(item interface{}) interface{} {

	t := reflect.TypeOf(item)
	v := reflect.ValueOf(item)

	if t.Implements(reflect.TypeOf((*json.Marshaler)(nil)).Elem()) {
		return item
	}

	if v.IsZero() {
		return item
	}

	switch t.Kind() {
	case reflect.Ptr:
		return Serialize(v.Elem().Interface())
	case reflect.Struct:
		valuesMap := make(map[string]interface{})
		numOfFields := t.NumField()
		for i := 0; i < numOfFields; i++ {
			// format in tag

			tag := t.Field(i).Tag

			jsontag := strings.Split(tag.Get("json"), ",")

			fieldname := t.Field(i).Name
			if !startWithSupperCase(fieldname) {
				continue
			}

			if len(jsontag) > 0 {
				if jsontag[0] == "-" {
					continue
				}
				fieldname, _ = firstNotEmpty(jsontag[0], fieldname)
			}

			format := tag.Get("format")
			var retval interface{}
			if formatter, ok := formatter.Formatters[format]; ok {
				retval = formatter(v.Field(i).Interface())
			} else {
				retval = Serialize(v.Field(i).Interface())
			}

			if retmap, ok := retval.(map[string]interface{}); t.Field(i).Anonymous && ok {
				for key, val := range retmap {
					valuesMap[key] = val
				}
			} else if !(reflect.ValueOf(retval).IsZero() && len(jsontag) > 1 && jsontag[1] == "omitempty") {
				valuesMap[fieldname] = retval
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
		return item

	}

}
