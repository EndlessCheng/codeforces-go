package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"reflect"
	"sort"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func CF271D(_r io.Reader, out io.Writer) {
	var s, t []byte
	var k, ans int
	Fscan(bufio.NewReader(_r), &s, &t, &k)
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	sum := make([]int, n+1)
	for i, b := range s {
		sum[i+1] = sum[i] + int(t[b-'a']&15^1)
	}
	for i, p := range sa {
		ans += sort.SearchInts(sum[int(p)+height[i]+1:], sum[p]+k+1)
	}
	Fprint(out, ans)
}

// 暴力求法
func CF271D_trie(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type node struct{ son [26]*node }

	var s, t []byte
	var k, ans int
	Fscan(in, &s, &t, &k)
	root := &node{}
	for i := range s {
		o, c := root, 0
		for _, b := range s[i:] {
			b -= 'a'
			if t[b] == '0' {
				if c++; c > k {
					break
				}
			}
			if o.son[b] == nil {
				o.son[b] = &node{}
				ans++
			}
			o = o.son[b]
		}
	}
	Fprint(out, ans)
}

//func main() { CF271D(os.Stdin, os.Stdout) }
