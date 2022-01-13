package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1117C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	dir4 := []struct{ x, y int64 }{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}

	var sx, sy, tx, ty, n int64
	var s string
	Fscan(in, &sx, &sy, &tx, &ty, &n, &s)
	l, r := int64(1), int64(1e14+1)
	for l < r {
		mid := (l + r) / 2
		x, y := sx, sy
		for i, b := range s {
			d := mid / n
			if i < int(mid%n) {
				d++
			}
			x += dir4[b].x * d
			y += dir4[b].y * d // 把风的位移和人的位移分开算
		}
		if abs(x-tx)+abs(y-ty) > mid { // 这样简单判断人的位移即可
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l > 1e14 {
		l = -1
	}
	Fprint(out, l)
}

//func main() { CF1117C(os.Stdin, os.Stdout) }
