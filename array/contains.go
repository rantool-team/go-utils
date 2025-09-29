package array

import "slices"

func Contains[T comparable](arr []T, element T) bool {
	return slices.Contains(arr, element)
}
