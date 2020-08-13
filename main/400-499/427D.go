package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"reflect"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func CF427D(_r io.Reader, _w io.Writer) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var s, t []byte
	Fscan(bufio.NewReader(_r), &s, &t)
	lenS1 := int32(len(s))
	s = append(append(s, '{'), t...)
	n := len(s)

	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
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

	ans := int(1e9)
	for i := 1; i+1 < n; i++ {
		if (sa[i-1] < lenS1) != (sa[i] < lenS1) && height[i-1] < height[i] && height[i] > height[i+1] {
			if l := 1 + max(height[i-1], height[i+1]); l < ans {
				ans = l
			}
		}
	}
	if ans == 1e9 {
		ans = -1
	}
	Fprint(_w, ans)
}

//func main() { CF427D(os.Stdin, os.Stdout) }