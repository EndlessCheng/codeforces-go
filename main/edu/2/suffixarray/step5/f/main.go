package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"math/bits"
	"os"
	"reflect"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var s []byte
	var m int
	Fscan(in, &s, &m)
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
	st := make([][19]int, n)
	for i, v := range height {
		st[i][0] = v
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = min(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	_q := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return min(st[l][k], st[r-1<<k][k]) }
	lcp := func(i, j int) int {
		if i == j {
			return n - i
		}
		ri, rj := rank[i], rank[j]
		if ri > rj {
			ri, rj = rj, ri
		}
		return _q(ri+1, rj+1)
	}

	ans := 1
	for l := 1; l < n; l++ {
		for i := 0; i+l < n; i += l {
			repLen := lcp(i, i+l)
			repCnt := repLen/l + 1
			if p := i - (l - repLen%l); p >= 0 && lcp(p, p+l) >= l {
				repCnt++
			}
			if repCnt > ans {
				ans = repCnt
			}
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
