package test

import "strings"

func escapeChar(char string) string {
	res := strings.ReplaceAll(char, "\n", "")
	res = strings.ReplaceAll(res, "\t", "")
	res = strings.ReplaceAll(res, ";", "")
	res = strings.ReplaceAll(res, " ", "")

	return res
}
