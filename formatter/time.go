package formatter

import "time"

func init() {
	RegisterFormatter("datetime", TimeFormatter("2006-01-02 15:04:05"))
	RegisterFormatter("time", TimeFormatter("2006-01-02"))
	RegisterFormatter("date", TimeFormatter("2006-01-02"))
}

func TimeFormatter(layout string) FormatterHandler {
	return func(v interface{}) interface{} {
		if t, ok := v.(time.Time); ok {
			return t.Format(layout)
		}

		if t, ok := v.(*time.Time); ok && t != nil {
			return t.Format(layout)
		}

		return v
	}
}
