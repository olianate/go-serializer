package serializer

func startWithSupperCase(values string) bool {
	rs := []rune(values)
	if len(rs) > 0 && rs[0] >= 'A' && rs[0] <= 'Z' {
		return true

	}
	return false
}

func firstNotEmpty(values ...string) (string, int) {

	for i, s := range values {
		if s != "" {
			return s, i
		}
	}

	return "", -1
}
