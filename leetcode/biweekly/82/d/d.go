package main

import "sort"

// https://space.bilibili.com/206214/dynamic
func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	left := make([]int, n) // left[i] 为左侧大于 nums[i] 的最近元素位置（不存在时为 -1）
	st := []int{}
	for i, v := range nums {
		for len(st) > 0 && nums[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	right := make([]int, n) // right[i] 为右侧大于 nums[i] 的最近元素位置（不存在时为 n）
	st = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] >= nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}

	for i, num := range nums {
		k := right[i] - left[i] - 1
		if num > threshold/k {
			return k
		}
	}
	return -1
}

func validSubarraySize2(nums []int, threshold int) int {
	n := len(nums)
	type pair struct{ v, i int }
	a := make([]pair, n)
	for i, v := range nums {
		a[i] = pair{v, i}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })

	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	for _, p := range a {
		i := p.i
		j := find(i + 1)
		fa[i] = j // 合并 i 和 i+1
		sz[j] += sz[i]
		if p.v > threshold/(sz[j]-1) {
			return sz[j] - 1
		}
	}
	return -1
}
