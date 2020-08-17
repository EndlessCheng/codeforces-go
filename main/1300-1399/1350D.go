package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1350D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, k int
o:
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k)
		ok := false
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			if a[i] == k {
				ok = true
			}
		}
		if !ok {
			Fprintln(out, "no")
			continue
		}
		if n == 1 {
			Fprintln(out, "yes")
			continue
		}
		for i := 0; i < n-1; i++ {
			if a[i] >= k && (a[i+1] >= k || i < n-2 && a[i+2] >= k) {
				Fprintln(out, "yes")
				continue o
			}
		}
		Fprintln(out, "no")
	}
}

//func main() { CF1350D(os.Stdin, os.Stdout) }
