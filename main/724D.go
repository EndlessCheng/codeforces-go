package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF724D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var m int
	var s []byte
	Fscan(in, &m, &s)
	n := len(s)
	pos := [26][]int{}
	for i, b := range s {
		b -= 'a'
		pos[b] = append(pos[b], i)
	}
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n+1)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	mergeRange := func(l, r int) (merged bool) {
		if l < 0 {
			l = 0
		}
		for i := find(l); i < r; i = find(i + 1) {
			fa[i] = r
			merged = true
		}
		return
	}

	initFa(n)
	ans := make([]byte, 0, n)
outer:
	for i, ps := range pos {
		b := byte(i + 'a')
		left := len(ps)
		for j := 0; j < len(ps); j++ {
			check := find(0)
			found := false
			for ; j < len(ps); j++ {
				if ps[j]-m+1 > check {
					break
				}
				found = true
			}
			if !found {
				for _, p := range ps[j:] {
					mergeRange(p-m+1, p+1)
				}
				break
			}
			if j > 0 {
				j--
			}
			p := ps[j]
			if mergeRange(p-m+1, p+1) {
				ans = append(ans, b)
				left--
				if find(0) >= n-m+1 {
					break outer
				}
			}
		}
		ans = append(ans, bytes.Repeat([]byte{b}, left)...)
	}
	Fprint(out, string(ans))
}

//func main() {
//	CF724D(os.Stdin, os.Stdout)
//}
