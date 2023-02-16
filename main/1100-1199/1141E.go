package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1141E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var h, v int64
	Fscan(in, &h, &n)
	s := make([]int64, n+1)
	p := []int{0}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		s[i] = s[i-1] - v
		if s[i] > s[p[len(p)-1]] {
			p = append(p, i)
		}
	}
	if h <= s[p[len(p)-1]] {
		i := sort.Search(len(p), func(i int) bool { return s[p[i]] >= h })
		Fprint(out, p[i])
	} else if s[n] < 1 {
		Fprint(out, -1)
	} else {
		loop := (h-s[p[len(p)-1]]-1)/s[n] + 1
		h -= loop * s[n]
		i := sort.Search(len(p), func(i int) bool { return s[p[i]] >= h })
		Fprint(out, loop*int64(n)+int64(p[i]))
	}
}

//func main() { CF1141E(os.Stdin, os.Stdout) }
