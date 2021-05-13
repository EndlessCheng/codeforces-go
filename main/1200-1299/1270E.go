package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1270E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n, &x, &y)
	a := make([]struct{ x, y int }, n)
	for i := 1; i < n; i++ {
		Fscan(in, &a[i].x, &a[i].y)
		a[i].x -= x
		a[i].y -= y
	}
	for {
		for _, p := range a {
			if p.x&1 > 0 || p.y&1 > 0 {
				var x, y []interface{}
				for i, p := range a {
					if p.x&1 == 0 && p.y&1 == 0 {
						x = append(x, i+1)
					} else if p.x&1 > 0 && p.y&1 > 0 {
						y = append(y, i+1)
					}
				}
				if len(x)+len(y) < n {
					x = append(x, y...)
				}
				Fprintln(out, len(x))
				Fprintln(out, x...)
				return
			}
		}
		for i := range a {
			a[i].x >>= 1
			a[i].y >>= 1
		}
	}
}

//func main() { CF1270E(os.Stdin, os.Stdout) }
