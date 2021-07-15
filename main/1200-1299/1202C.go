package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1202C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type xy struct{ x, y int }
	dir4 := []xy{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dir4c := []xy{
		'A': {-1, 0},
		'D': {1, 0},
		'S': {0, -1},
		'W': {0, 1},
	}
	type pair struct{ miX, mxX, miY, mxY int }
	update := func(p pair, x, y int) pair {
		return pair{
			min(p.miX, x),
			max(p.mxX, x),
			min(p.miY, y),
			max(p.mxY, y),
		}
	}

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		suf := make([]pair, n+1)
		sufX, sufY := 0, 0
		for i := n - 1; i >= 0; i-- {
			d := dir4c[s[i]]
			sufX -= d.x
			sufY -= d.y
			suf[i] = update(suf[i+1], sufX, sufY)
		}
		q := suf[0]
		ans := int64(q.mxX-q.miX+1) * int64(q.mxY-q.miY+1)
		x, y := 0, 0
		p := pair{}
		for i, b := range s {
			q := suf[i]
			for _, d := range dir4 {
				x, y := x+d.x, y+d.y
				shiftX, shiftY := x-sufX, y-sufY
				pp := update(p, x, y)
				miX := min(pp.miX, q.miX+shiftX)
				mxX := max(pp.mxX, q.mxX+shiftX)
				miY := min(pp.miY, q.miY+shiftY)
				mxY := max(pp.mxY, q.mxY+shiftY)
				if a := int64(mxX-miX+1) * int64(mxY-miY+1); a < ans {
					ans = a
				}
			}
			d := dir4c[b]
			sufX += d.x
			sufY += d.y
			x += d.x
			y += d.y
			p = update(p, x, y)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1202C(os.Stdin, os.Stdout) }
