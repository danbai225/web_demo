package str

import "strings"

func RemoveSpecial(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	return str
}
func TimeStr2Date(timeStr string) string {
	if len(timeStr) >= 10 {
		return timeStr[:10]
	}
	if len(timeStr) == 10 {
		return timeStr
	}
	if timeStr == "" {
		return ""
	}
	return "1001-01-01"
}
