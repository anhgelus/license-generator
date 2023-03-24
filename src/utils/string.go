package utils

func GenerateSumeUp(name string, m map[string]string, s string) {
	println(name + ":")
	for id, content := range m {
		println(s + " " + id + ": " + content)
	}
}

func StringArrayToString(a []string) string {
	str := ""
	for i, s := range a {
		if i == len(a)-1 {
			str = str + " and " + s
		} else if i == 0 {
			str = s
		} else {
			str = str + ", " + s
		}
	}
	return str
}
