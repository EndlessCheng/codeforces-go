package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF676D(_r io.Reader, _w io.Writer) {
	doorTable := [125]string{}
	doorTable['+'] = "++++"
	doorTable['-'] = "-|-|"
	doorTable['|'] = "|-|-"
	doorTable['^'] = "^>v<"
	doorTable['>'] = ">v<^"
	doorTable['v'] = "v<^>"
	doorTable['<'] = "<^>v"
	doorTable['U'] = "URDL"
	doorTable['R'] = "RDLU"
	doorTable['D'] = "DLUR"
	doorTable['L'] = "LURD"
	dir4 := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // 上右下左
	dirTable := [125][]int{}
	dirTable['+'] = []int{0, 1, 2, 3}
	dirTable['-'] = []int{1, 3}
	dirTable['|'] = []int{0, 2}
	dirTable['^'] = []int{0}
	dirTable['>'] = []int{1}
	dirTable['v'] = []int{2}
	dirTable['<'] = []int{3}
	dirTable['U'] = []int{1, 2, 3}
	dirTable['R'] = []int{0, 2, 3}
	dirTable['D'] = []int{0, 1, 3}
	dirTable['L'] = []int{0, 1, 2}
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
	g := make([]string, n)
	for i := range g {
		Fscan(in, &g[i])
	}
	Fscan(in, &ax, &ay, &bx, &by)
	bx--
	by--

	type stat struct{ x, y, rot int }
	qs := [4][]stat{{{ax - 1, ay - 1, 0}}}
	vis := map[stat]bool{}
	isAllQueueEmpty := func() bool {
		for _, q := range qs {
			if len(q) > 0 {
				return false
			}
		}
		return true
	}
	for time := 0; !isAllQueueEmpty(); time++ {
		curQ := make([]stat, len(qs[time&3]))
		copy(curQ, qs[time&3])
		qs[time&3] = []stat{}
		for _, s := range curQ {
			if s.x == bx && s.y == by {
				Fprint(out, time)
				return
			}
			if vis[s] {
				continue
			}
			vis[s] = true
			for rotTimes := 0; rotTimes < 4; rotTimes++ {
				rot := (s.rot + rotTimes) & 3
				door := doorTable[g[s.x][s.y]][rot]
				for _, i := range dirTable[door] {
					dir := dir4[i]
					x, y := s.x+dir[0], s.y+dir[1]
					if x < 0 || x >= n || y < 0 || y >= m || g[x][y] == '*' {
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

func main() {
	CF676D(os.Stdin, os.Stdout)
}
