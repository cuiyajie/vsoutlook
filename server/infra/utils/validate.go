package utils

import "golang.org/x/exp/slices"

func ValidateEnumFallback(val int, pool []int, fallback int) int {
	if !slices.Contains(pool, val) {
		return fallback
	}
	return val
}
