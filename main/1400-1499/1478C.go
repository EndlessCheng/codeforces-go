package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1478C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		n *= 2
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
		s := int64(0)
		for i := 0; i < n; i += 2 {
			v := (a[i] - s) / int64(n-i)
			if a[i] != a[i+1] || a[i] <= s || (a[i]-s)%int64(n-i) > 0 || i > 0 && v >= a[i-2] {
				Fprintln(out, "NO")
				continue o
			}
			a[i] = v
			s += v * 2
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1478C(os.Stdin, os.Stdout) }
