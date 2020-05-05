package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF432D(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var s []byte
	Fscan(_r, &s)
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		z[i] = max(0, min(z[i-l], r-i+1))
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}
	z[0] = n
	sorted := make([]int, len(z))
	copy(sorted, z)
	sort.Ints(sorted)
	ans := [][2]int{}
	for l := 1; l <= n; l++ {
		if z[n-l] == l {
			i := sort.SearchInts(sorted, l)
			ans = append(ans, [2]int{l, n - i})
		}
	}
	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

//func main() { CF432D(os.Stdin, os.Stdout) }
