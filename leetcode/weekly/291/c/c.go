package main

import (
	"index/suffixarray"
	"reflect"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func countDistinct(nums []int, k, p int) (ans int) {
	n := len(nums)
	s := make([]byte, n)
	right := make([]int, n)
	r, cnt := 0, 0
	for l, num := range nums {
		s[l] = byte(num)
		for ; r < n && (cnt < k || nums[r]%p != 0); r++ {
			if nums[r]%p == 0 {
				cnt++
			}
		}
		ans += r - l
		right[l] = r
		if num%p == 0 {
			cnt--
		}
	}

	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		ans -= min(h, right[sa[rk]]-int(sa[rk]))
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func countDistinct2(nums []int, k, p int) int {
	set := map[[200]int]struct{}{}
	for i := range nums {
		arr, idx, cnt := [200]int{}, 0, 0
		for _, v := range nums[i:] {
			if v%p == 0 {
				if cnt++; cnt > k {
					break
				}
			}
			arr[idx] = v
			idx++
			set[arr] = struct{}{}
		}
	}
	return len(set)
}
