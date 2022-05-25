package formatter

import "strings"

func init() {
	RegisterFormatter("str2array", Str2Array)
}

func Str2Array(v interface{}) interface{} {

	if s, ok := v.(string); ok {
		if len(s) == 0 {
			return []string{}
		} else {
			return strings.Split(s, ",")
		}
	}

	return v

}
