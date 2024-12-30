package main

import (
	"index/suffixarray"
	"unsafe"
)

// https://space.bilibili.com/206214
func answerString(s string, k int) string {
	if k == 1 {
		return s
	}
	sa := (*struct{_[]byte;sa[]int32})(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
	n := len(s)
	i := int(sa[n-1])
	return s[i:min(i+n-k+1, n)]
}

func answerString2(s string, k int) string {
	if k == 1 {
		return s
	}
	n := len(s)
	i, j := 0, 1
	for j < n {
		k := 0
		for j+k < n && s[i+k] == s[j+k] {
			k++
		}
		if j+k < n && s[i+k] < s[j+k] {
			i, j = j, max(j+1, i+k+1)
		} else {
			j += k + 1
		}
	}
	return s[i:min(i+n-k+1, n)]
}

func answerString3(s string, k int) (ans string) {
	if k == 1 {
		return s
	}
	n := len(s)
	for i := range n {
		ans = max(ans, s[i:min(i+n-k+1, n)])
	}
	return
}
