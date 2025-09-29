package array

func SameListExcept[T any](list []T, indexToNotInclude int) []T {
	newList := []T{}
	for index, value := range list {
		if index == indexToNotInclude {
			break
		}

		newList = append(newList, value)
	}

	return newList
}
