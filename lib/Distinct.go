package lib

func Distinct(slice []int) (result []int) {
	k := make(map[int]bool)
	for _, v := range slice {
		if _, val := k[v]; !val {
			k[v] = true
			result = append(result, v)
		}
	}
	return
}
