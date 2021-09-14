package main

import (
	"index/suffixarray"
	"reflect"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func largestMerge(s, t string) string {
	ans := make([]byte, 0, len(s)+len(t))
	for {
		if s == "" {
			ans = append(ans, t...)
			break
		}
		if t == "" {
			ans = append(ans, s...)
			break
		}
		if s > t {
			ans = append(ans, s[0])
			s = s[1:]
		} else {
			ans = append(ans, t[0])
			t = t[1:]
		}
	}
	return string(ans)
}

// 后缀数组 O(|s|+|t|) 做法
func largestMergeSA(s, t string) string {
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s + "#" + t))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n+1+len(t))
	for i := range rank {
		rank[sa[i]] = i
	}
	ans := []byte{}
	i, j := 0, n+1
	for {
		if s == "" {
			ans = append(ans, t...)
			break
		}
		if t == "" {
			ans = append(ans, s...)
			break
		}
		if rank[i] > rank[j] {
			ans = append(ans, s[0])
			s = s[1:]
			i++
		} else {
			ans = append(ans, t[0])
			t = t[1:]
			j++
		}
	}
	return string(ans)
}
