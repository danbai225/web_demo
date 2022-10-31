package array

import (
	"strconv"
	"strings"
)

func InArrayInt64s(int64s []int64, item int64) bool {
	for _, i := range int64s {
		if i == item {
			return true
		}
	}
	return false
}

func String2IntArr(string string, delimiter string) []int64 {
	split := strings.Split(string, delimiter)
	int64s := make([]int64, 0)
	for _, s := range split {
		i, _ := strconv.ParseInt(s, 10, 64)
		int64s = append(int64s, i)
	}
	return int64s
}
