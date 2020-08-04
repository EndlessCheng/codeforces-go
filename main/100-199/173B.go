package _00_199

import (
	"bufio"
	. "fmt"
	"io"
)

type dPair173 struct{ x, y, dir int }
type deque173 struct{ l, r []dPair173 }

func (q *deque173) empty() bool      { return len(q.l) == 0 && len(q.r) == 0 }
func (q *deque173) pushL(v dPair173) { q.l = append(q.l, v) }
func (q *deque173) pushR(v dPair173) { q.r = append(q.r, v) }
func (q *deque173) popL() (v dPair173) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func CF173B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	dir4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var n, m int
	Fscan(in, &n, &m)
	mat := make([][]byte, n)
	for i := range mat {
		Fscan(in, &mat[i])
	}
	const mx, inf = 1000, 1e9
	dist := [mx][mx][4]int{}
	for i, di := range dist {
		for j, dij := range di {
			for k := range dij {
				dist[i][j][k] = inf
			}
		}
	}
	dist[n-1][m-1][3] = 0
	q := &deque173{}
	q.pushL(dPair173{n - 1, m - 1, 3})
	for !q.empty() {
		p := q.popL()
		x, y, dir := p.x, p.y, p.dir
		d := dist[x][y][dir]
		if xx, yy := x+dir4[dir][0], y+dir4[dir][1]; xx >= 0 && xx < n && yy >= 0 && yy < m && d < dist[xx][yy][dir] {
			dist[xx][yy][dir] = d
			q.pushL(dPair173{xx, yy, dir})
		}
		if mat[x][y] == '#' {
			for i := 0; i < 4; i++ {
				if i != dir && d+1 < dist[x][y][i] {
					dist[x][y][i] = d + 1
					q.pushR(dPair173{x, y, i})
				}
			}
		}
	}
	if dist[0][0][3] == inf {
		dist[0][0][3] = -1
	}
	Fprint(out, dist[0][0][3])
}

//func main() {
//	CF173B(os.Stdin, os.Stdout)
//}
