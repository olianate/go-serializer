package formatter

type FormatterHandler func(v interface{}) interface{}

var Formatters = map[string]FormatterHandler{}

func RegisterFormatter(key string, fn FormatterHandler) {
	Formatters[key] = fn
}
