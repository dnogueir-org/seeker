package common

func IsStringInSlice(s string, sl []string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}
