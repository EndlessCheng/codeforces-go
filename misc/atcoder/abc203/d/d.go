package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	s := make([][]int, n+1)
	for i := range s {
		s[i] = make([]int, n+1)
	}
	Fprint(out, sort.Search(1e9, func(up int) bool {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				s[i][j] = s[i][j-1] + s[i-1][j] - s[i-1][j-1]
				if a[i-1][j-1] <= up {
					s[i][j]++
				}
				if i >= k && j >= k && s[i][j]-s[i][j-k]-s[i-k][j]+s[i-k][j-k] >= (k*k+1)/2 {
					return true
				}
			}
		}
		return false
	}))
}

func main() { run(os.Stdin, os.Stdout) }
