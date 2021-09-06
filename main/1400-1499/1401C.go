package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1401C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		min := int(1e9)
		for i := range a {
			Fscan(in, &a[i])
			if a[i] < min {
				min = a[i]
			}
		}
		b := append([]int(nil), a...)
		sort.Ints(b)
		for i, v := range b {
			if v != a[i] && v%min > 0 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1401C(os.Stdin, os.Stdout) }
