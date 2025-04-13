package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func fwtXOR(a []int, rsh int) {
	n := len(a)
	for l, k := 2, 1; l <= n; l, k = l<<1, k<<1 {
		for i := 0; i < n; i += l {
			for j := 0; j < k; j++ {
				a[i+j], a[i+j+k] = (a[i+j]+a[i+j+k])>>rsh, (a[i+j]-a[i+j+k])>>rsh
			}
		}
	}
}

func fwtXOR3(a []int) []int {
	fwtXOR(a, 0)
	for i, x := range a {
		a[i] *= x * x
	}
	fwtXOR(a, 1)
	return a
}

func uniqueXorTriplets(nums []int) (ans int) {
	cnt := make([]int, 1<<bits.Len(uint(slices.Max(nums))))
	for _, x := range nums {
		cnt[x]++
	}
	for _, c := range fwtXOR3(cnt) {
		if c > 0 {
			ans++
		}
	}
	return
}

func uniqueXorTriplets1(nums []int) (ans int) {
	u := 1 << bits.Len(uint(slices.Max(nums)))
	has := make([]bool, u)
	for i, x := range nums {
		for _, y := range nums[i:] {
			has[x^y] = true
		}
	}

	has3 := make([]bool, u)
	for xy, b := range has {
		if !b {
			continue
		}
		for _, z := range nums {
			has3[xy^z] = true
		}
	}

	for _, b := range has3 {
		if b {
			ans++
		}
	}
	return
}
