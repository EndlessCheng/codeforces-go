package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1618D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		s := 0
		for _, v := range a[:n-k*2] {
			s += v
		}
		for i := n - k; i < n; i++ {
			s += a[i-k] / a[i]
		}
		Fprintln(out, s)
	}
}

//func main() { CF1618D(os.Stdin, os.Stdout) }
