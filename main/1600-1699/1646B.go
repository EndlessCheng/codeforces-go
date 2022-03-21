package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1646B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		for l, r, s, t := 1, n-1, int64(a[0]), int64(0); l < r; l++ {
			s += int64(a[l])
			t += int64(a[r])
			if s < t {
				Fprintln(out, "YES")
				continue o
			}
			r--
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1646B(os.Stdin, os.Stdout) }
