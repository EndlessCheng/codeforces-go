package main

import (
	. "fmt"
	"io"
	"slices"
	"strings"
)

// https://github.com/EndlessCheng
func cf137D(in io.Reader, out io.Writer) {
	var s []byte
	var k int
	Fscan(in, &s, &k)
	n := len(s)
	mc := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		mc[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			mc[i][j] = mc[i+1][j-1]
			if s[i] != s[j] {
				mc[i][j]++
			}
		}
	}

	fr := make([][]int, k)
	for i := range fr {
		fr[i] = make([]int, n)
	}
	f := mc[0]
	minF, minK := f[n-1], 0
	for i := 1; i < k; i++ {
		for r := n - 1; r >= i; r-- {
			f[r] = 1e9
			for l := i; l <= r; l++ {
				v := f[l-1] + mc[l][r]
				if v < f[r] {
					f[r] = v
					fr[i][r] = l
				}
			}
		}
		if f[n-1] < minF {
			minF, minK = f[n-1], i
		}
	}
	Fprintln(out, minF)

	path := []string{}
	for i, r := minK, n-1; i >= 0; i-- {
		l := fr[i][r]
		t := s[l : r+1]
		for j, m := 0, len(t); j < m/2; j++ {
			c := t[m-1-j]
			if t[j] != c {
				t[j] = c
			}
		}
		path = append(path, string(t))
		r = l - 1
	}
	slices.Reverse(path)
	Fprint(out, strings.Join(path, "+"))
}

//func main() { cf137D(os.Stdin, os.Stdout) }
