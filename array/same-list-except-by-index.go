package array

func SameListExcept[T any](list []T, indexToNotInclude int) []T {
	newList := make([]T, 0, len(list)-1)

	newList = append(newList, list[:indexToNotInclude]...)

	newList = append(newList, list[indexToNotInclude+1:]...)

	return newList
}
