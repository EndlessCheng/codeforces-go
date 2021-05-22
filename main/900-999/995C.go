package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
)

// github.com/EndlessCheng/codeforces-go
func CF995C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type point struct {
		x, y int64
		i    int
	}

	var n int
	Fscan(in, &n)
	a := make([]point, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
		a[i].i = i
	}
	ans := make([]interface{}, n)
	for {
		rand.Shuffle(n, func(i, j int) { a[i], a[j] = a[j], a[i] })
		var x, y int64
		for _, p := range a {
			if x*p.x+y*p.y < 0 {
				x += p.x
				y += p.y
				ans[p.i] = 1
			} else {
				x -= p.x
				y -= p.y
				ans[p.i] = -1
			}
		}
		if x*x+y*y <= 2.25e12 {
			Fprintln(out, ans...)
			return
		}
	}
}

//func main() { CF995C(os.Stdin, os.Stdout) }
