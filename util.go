package serializer

func startWithSupperCase(values string) bool {
	rs := []rune(values)
	if len(rs) > 0 && rs[0] >= 'A' && rs[0] <= 'Z' {
		return true

	}
	return false
}

func orString(elments ...string) string {
	for _, e := range elments {
		if e != "" {
			return e
		}
	}

	return ""
}
