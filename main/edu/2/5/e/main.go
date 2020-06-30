package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"os"
	"reflect"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	s := make([]byte, n)
	for i := range s {
		Fscan(in, &s[i])
	}

	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	var h int
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

	type pair struct{ v, i int }
	posL := make([]int, n)
	stack := []pair{{-1, -1}}
	for i, v := range height {
		for {
			if top := stack[len(stack)-1]; top.v < v {
				posL[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}
	posR := make([]int, n)
	stack = []pair{{-1, n}}
	for i := n - 1; i >= 0; i-- {
		v := height[i]
		for {
			if top := stack[len(stack)-1]; top.v < v {
				posR[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}

	ans, l, p := int64(n), n, 0
	for i, h := range height {
		if s := int64(h) * int64(posR[i]-posL[i]); s > ans {
			ans, l, p = s, h, int(sa[i])
		}
	}
	Fprintln(out, ans)
	Fprintln(out, l)
	for _, b := range s[p : p+l] {
		Fprint(out, b, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
