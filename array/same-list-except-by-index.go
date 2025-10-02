package array

func SameListExcept[T any](list []T, indexToNotInclude int) []T {
	var newList []T
	for index, value := range list {
		if index == indexToNotInclude {
			continue
		}
		newList = append(newList, value)
	}
	return newList
}
