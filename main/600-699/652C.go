package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF652C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, m, x, y int
	Fscan(in, &n, &m)
	pos := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &x)
		pos[x] = i
	}
	l := make([]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &x, &y)
		x, y = pos[x], pos[y]
		if x > y {
			x, y = y, x
		}
		l[y] = max(l[y], x)
	}
	ans := int64(0)
	for i, maxL := 1, 0; i <= n; i++ {
		maxL = max(maxL, l[i])
		ans += int64(i - maxL)
	}
	Fprint(out, ans)
}

//func main() { CF652C(os.Stdin, os.Stdout) }
