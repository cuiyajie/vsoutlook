package utils

import "golang.org/x/exp/slices"

func StringContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func SliceIndex(limit int, filter func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if filter(i) {
			return i
		}
	}
	return -1
}

func Filter[T any](input []T, filter func(t T) bool) (result []T) {
	for _, x := range input {
		if filter(x) {
			result = append(result, x)
		}
	}
	return
}

func Remove[E comparable](s1 []E, s2 E) []E {
	idx := slices.Index(s1, s2)
	if idx > -1 {
		return slices.Delete(s1, idx, idx+1)
	}
	return s1
}
