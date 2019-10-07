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
func Sol427D(reader io.Reader, writer io.Writer) {
	calcLCP := func(s []byte, sa []int) (lcp []int) {
		n := len(s)
		rank := make([]int, n+1)
		for i := range rank {
			rank[sa[i]] = i
		}
		lcp = make([]int, n, n+1)
		h := 0
		for i := range lcp {
			j := sa[rank[i]-1]
			if h > 0 {
				h--
			}
			for ; j+h < n && i+h < n; h++ {
				if s[j+h] != s[i+h] {
					break
				}
			}
			lcp[rank[i]-1] = h
		}
		return
	}
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var s, t []byte
	Fscan(in, &s, &t)
	lenS1 := len(s)
	s = append(s, '$')
	s = append(s, t...)
	n := len(s)

	sa := *(*[]int)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").UnsafeAddr()))
	sa = append([]int{len(s)}, sa...)
	lcp := calcLCP(s, sa)
	lcp = append(lcp, 0)

	const inf int = 1e9
	ans := inf
	for i := 2; i < n; i++ {
		if (sa[i] < lenS1) != (sa[i+1] < lenS1) && lcp[i-1] < lcp[i] && lcp[i+1] < lcp[i] {
			ans = min(ans, max(lcp[i-1], lcp[i+1])+1)
		}
	}
	if ans == inf {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() {
//	Sol427D(os.Stdin, os.Stdout)
//}
