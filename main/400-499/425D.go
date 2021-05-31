package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF425D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var xy, yx [1e5 + 1]map[int]bool
	for i := 0; i <= 1e5; i++ {
		xy[i] = map[int]bool{}
		yx[i] = map[int]bool{}
	}
	var n, ans int
	Fscan(in, &n)
	a := make([]struct{ x, y int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
		xy[a[i].x][a[i].y] = true
		yx[a[i].y][a[i].x] = true
	}
	for _, p := range a {
		x, y := p.x, p.y
		if len(xy[x]) < len(yx[y]) {
			for yy := range xy[x] {
				if yx[y][x+yy-y] && yx[yy][x+yy-y] {
					ans++
				}
			}
		} else {
			for xx := range yx[y] {
				if xy[x][y+xx-x] && xy[xx][y+xx-x] {
					ans++
				}
			}
		}
	}
	Fprint(out, (ans-n)/2)
}

//func main() { CF425D(os.Stdin, os.Stdout) }
