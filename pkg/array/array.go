package array

func InArrayInt64s(int64s []int64, item int64) bool {
	for _, i := range int64s {
		if i == item {
			return true
		}
	}
	return false
}
