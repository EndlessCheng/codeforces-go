package main

import "sort"

func canBeEqual(target []int, arr []int) (ans bool) {
	equal := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		sort.Ints(a)
		sort.Ints(b)
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}
	return equal(target, arr)
}
