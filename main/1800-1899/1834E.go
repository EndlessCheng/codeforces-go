package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1834E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int64) int64 { return a / gcd(a, b) * b }

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		vis := map[int64]bool{}
		set := []int64{}
		for _, v := range a {
			for j, w := range set {
				set[j] = lcm(w, int64(v))
			}
			set = append(set, int64(v))
			j := 0
			for _, w := range set[1:] {
				if set[j] != w {
					j++
					set[j] = w
				}
			}
			set = set[:j+1]
			for len(set) > 0 && set[0] > int64(n*24) {
				set = set[1:]
			}
			for _, w := range set {
				vis[w] = true
			}
		}
		mex := int64(1)
		for vis[mex] {
			mex++
		}
		Fprintln(out, mex)
	}
}

//func main() { CF1834E(os.Stdin, os.Stdout) }
