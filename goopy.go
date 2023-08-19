package goopy

func GenTS(target interface{}, indent int) (res string) {
	res = gen_ts(target, indent, indent)
	return
}

func GenJS(target interface{}) (res string) {
	return
}
