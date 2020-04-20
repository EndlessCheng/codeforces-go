package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	sign := func(b bool) int {
		if b {
			return 1
		}
		return -1
	}
	xor := func(b1, b2 bool) bool { return b1 && !b2 || !b1 && b2 }

	solve := func(_case int) string {
		var xx, yy int
		Fscan(in, &xx, &yy)

		// 考虑从 x,y 走到原点。这样考虑时，方向和从原点走到 x,y 的方向是相反的
		ds := [][2]int{}
		for x, y := abs(xx), abs(yy); x > 0 || y > 0; {
			// 同奇偶时无法走回原点
			if x%2 == y%2 {
				return "IMPOSSIBLE"
			}

			if x+y == 1 { // 1,0 或 0,1
				ds = append(ds, [2]int{-x, -y})
				break
			}

			dx, dy := 0, 0
			if x%2 == 1 {
				// 选择能走成异奇偶的走法：若 x+1 能走成异奇偶，则 dx=1，否则 dx=-1
				dx = sign((x+1)/2%2 != y/2%2)
			} else {
				// 同上
				dy = sign((y+1)/2%2 != x/2%2)
			}
			x = (x + dx) / 2
			y = (y + dy) / 2
			ds = append(ds, [2]int{dx, dy})
		}

		ans := make([]byte, len(ds))
		for i, d := range ds {
			if d[0] != 0 {
				if xor(d[0] < 0, xx < 0) {
					ans[i] = 'E'
				} else {
					ans[i] = 'W'
				}
			} else {
				if xor(d[1] < 0, yy < 0) {
					ans[i] = 'N'
				} else {
					ans[i] = 'S'
				}
			}
		}
		return string(ans)
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: %s\n", _case, solve(_case))
	}
}

func main() { run(os.Stdin, os.Stdout) }
