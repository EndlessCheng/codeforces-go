package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF460C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, w int
	Fscan(in, &n, &m, &w)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fprint(out, sort.Search(1e9+m, func(low int) bool {
		low++
		cnt := 0
		diff := make([]int, n)
		sumD := 0
		for i, v := range a {
			sumD += diff[i]
			if sumD+v < low {
				d := low - sumD - v
				cnt += d
				if cnt > m {
					return true
				}
				sumD += d
				if i+w < n {
					diff[i+w] -= d
				}
			}
		}
		return false
	}))
}

//func main() { CF460C(os.Stdin, os.Stdout) }
