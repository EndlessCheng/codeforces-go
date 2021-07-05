package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1408D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m, x, y, maxDY int
	Fscan(in, &n, &m)
	a := make([]struct{ x, y int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	dy := [1e6 + 1]int{}
	for ; m > 0; m-- {
		Fscan(in, &x, &y)
		for _, p := range a {
			if x >= p.x && y >= p.y {
				dy[x-p.x] = max(dy[x-p.x], y-p.y+1)
			}
		}
	}
	ans := int(1e9)
	for x = 1e6; x >= 0; x-- {
		maxDY = max(maxDY, dy[x])
		if x+maxDY < ans {
			ans = x + maxDY
		}
	}
	Fprint(out, ans)
}

//func main() { CF1408D(os.Stdin, os.Stdout) }
