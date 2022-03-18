package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1618E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		b := make([]int64, n)
		tot := int64(0)
		for i := range b {
			Fscan(in, &b[i])
			tot += b[i]
		}
		d := int64(n) * int64(n+1) / 2
		if tot%d > 0 {
			Fprintln(out, "NO")
			continue
		}
		tot /= d
		a := make([]interface{}, n)
		for i, v := range b {
			s := tot + b[(i+n-1)%n] - v
			if s <= 0 || s%int64(n) > 0 {
				Fprintln(out, "NO")
				continue o
			}
			a[i] = s / int64(n)
		}
		Fprintln(out, "YES")
		Fprintln(out, a...)
	}
}

//func main() { CF1618E(os.Stdin, os.Stdout) }
