package main

import (
	"index/suffixarray"
	"reflect"
	"strings"
	"unsafe"
)

func stringMatching(a []string) (ans []string) {
	for i, s := range a {
		for j, s2 := range a {
			if j != i && strings.Contains(s2, s) {
				ans = append(ans, s)
				break
			}
		}
	}
	return
}

// O(∑len(a[i]))
func stringMatchingSA(a []string) (ans []string) {
	s := "#" + strings.Join(a, "#")
	n := len(s)
	lens := make([]int, n)
	cnt := 0
	for i := 1; i < n; i++ {
		if s[i-1] == '#' {
			lens[i] = len(a[cnt])
			cnt++
		}
	}
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, ri := range rank {
		if h > 0 {
			h--
		}
		if ri > 0 {
			for j := int(sa[ri-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[ri] = h
	}
	for i, p := range sa {
		if l := lens[p]; l > 0 {
			if height[i] >= l || i+1 < n && height[i+1] >= l {
				ans = append(ans, s[p:int(p)+l])
			}
		}
	}
	return
}
