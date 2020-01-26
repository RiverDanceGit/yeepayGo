package util

func StrInArray(search string, strArr []string) bool {
	return inArray(len(strArr), func(i int) bool { return strArr[i] == search })
}

func inArray(n int, f func(int) bool) bool {
	for i := 0; i < n; i++ {
		if f(i) {
			return true
		}
	}
	return false
}
