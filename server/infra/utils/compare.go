package utils

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](args ...T) T {
	if len(args) == 0 {
		return *new(T)
	}
	max := args[0]
	for _, arg := range args[1:] {
		if arg > max {
			max = arg
		}
	}
	return max
}

func Min[T constraints.Ordered](args ...T) T {
	if len(args) == 0 {
		return *new(T)
	}
	min := args[0]
	for _, arg := range args[1:] {
		if arg < min {
			min = arg
		}
	}
	return min
}

func MaxIn[T constraints.Ordered](args []T) T {
	if len(args) == 0 {
		return *new(T)
	}
	max := args[0]
	for _, arg := range args[1:] {
		if arg > max {
			max = arg
		}
	}
	return max
}

func MinIn[T constraints.Ordered](args []T) T {
	if len(args) == 0 {
		return *new(T)
	}
	min := args[0]
	for _, arg := range args[1:] {
		if arg < min {
			min = arg
		}
	}
	return min
}
