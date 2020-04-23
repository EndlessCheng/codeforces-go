package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF676D(_r io.Reader, _w io.Writer) {
	doorTable := [...]string{
		'+': "++++",
		'-': "-|-|",
		'|': "|-|-",
		'^': "^>v<",
		'>': ">v<^",
		'v': "v<^>",
		'<': "<^>v",
		'U': "URDL",
		'R': "RDLU",
		'D': "DLUR",
		'L': "LURD",
	}
	dir4 := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // 上右下左
	dirTable := [...][]int{
		'+': {0, 1, 2, 3},
		'-': {1, 3},
		'|': {0, 2},
		'^': {0},
		'>': {1},
		'v': {2},
		'<': {3},
		'U': {1, 2, 3},
		'R': {0, 2, 3},
		'D': {0, 1, 3},
		'L': {0, 1, 2},
	}
	canBack := func(from int, backs []int) bool {
		for _, back := range backs {
			if (back+2)&3 == from {
				return true
			}
		}
		return false
	}

	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, ax, ay, bx, by int
	Fscan(in, &n, &m)
	g := make([][]byte, n)
	for i := range g {
		Fscan(in, &g[i])
	}
	Fscan(in, &ax, &ay, &bx, &by)
	bx--
	by--

	type stat struct{ x, y, rot int }
	qs := [4][]stat{{{ax - 1, ay - 1, 0}}}
	allEmpty := func() bool {
		for _, q := range qs {
			if len(q) > 0 {
				return false
			}
		}
		return true
	}
	vis := [1000][1000][4]bool{}
	for time := 0; !allEmpty(); time++ {
		q := qs[time&3]
		qs[time&3] = []stat{}
		for _, s := range q {
			if s.x == bx && s.y == by {
				Fprint(out, time)
				return
			}
			if vis[s.x][s.y][s.rot] {
				continue
			}
			vis[s.x][s.y][s.rot] = true
			for rotTimes := 0; rotTimes < 4; rotTimes++ {
				rot := (s.rot + rotTimes) & 3
				door := doorTable[g[s.x][s.y]][rot]
				for _, i := range dirTable[door] {
					d := dir4[i]
					x, y := s.x+d[0], s.y+d[1]
					if x < 0 || x >= n || y < 0 || y >= m || g[x][y] == '*' || vis[x][y][rot] {
						continue
					}
					door1 := doorTable[g[x][y]][rot]
					if !canBack(i, dirTable[door1]) {
						continue
					}
					tarTime := time + rotTimes + 1
					qs[tarTime&3] = append(qs[tarTime&3], stat{x, y, rot})
				}
			}
		}
	}
	Fprint(out, -1)
}

func main() { CF676D(os.Stdin, os.Stdout) }
