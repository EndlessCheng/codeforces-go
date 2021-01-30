package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1475D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, tp int
	var low int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &low)
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := [3][]int64{}
		for _, v := range a {
			Fscan(in, &tp)
			b[tp] = append(b[tp], v)
		}
		for _, a := range b {
			sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
		}
		n2 := len(b[2])
		s2 := make([]int64, n2+1)
		for i, v := range b[2] {
			s2[i+1] = s2[i] + v
		}

		ans := int(1e9)
		b[1] = append(b[1], 0)
		for i, v := range b[1] {
			if j := sort.Search(n2+1, func(j int) bool { return s2[j] >= low }); j <= n2 && i+2*j < ans {
				ans = i + 2*j
			}
			low -= v
		}

		if ans == 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1475D(os.Stdin, os.Stdout) }
