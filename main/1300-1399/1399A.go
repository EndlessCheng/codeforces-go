package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1399A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		for i := 1; i < n; i++ {
			if a[i]-a[i-1] > 1 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1399A(os.Stdin, os.Stdout) }
