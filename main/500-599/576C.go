package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF576C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }

	var n int
	Fscan(in, &n)
	g := [1001][1001][]int{}
	ps := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &ps[i].x, &ps[i].y)
		x, y := ps[i].x/1000, ps[i].y/1000
		if x&1 > 0 {
			y = 1000 - y // 这样写，下面就只需要在奇数行时交换一下最左最右的两个点
		}
		g[x][y] = append(g[x][y], i)
	}
	for i := 0; i < 1001; i++ {
		for j := 0; j < 1001; j++ {
			ids := g[i][j]
			if len(ids) == 0 {
				continue
			}
			miI, mxI := ids[0], ids[0]
			for _, id := range ids[1:] {
				if ps[id].x < ps[miI].x {
					miI = id
				} else if ps[id].x > ps[mxI].x {
					mxI = id
				}
			}
			if i&1 > 0 {
				miI, mxI = mxI, miI
			}
			Fprint(out, miI, " ")
			for _, id := range ids {
				if id != miI && id != mxI {
					Fprint(out, id, " ")
				}
			}
			if mxI != miI {
				Fprint(out, mxI, " ")
			}
		}
	}
}

//func main() { CF576C(os.Stdin, os.Stdout) }
