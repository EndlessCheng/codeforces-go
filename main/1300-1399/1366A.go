package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 另一种写法是 min(a,b,(a+b)/3)

// github.com/EndlessCheng/codeforces-go
func CF1366A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var t, a, b int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &a, &b)
		if a > b {
			a, b = b, a
		}
		c := 2*a - b
		if c < 0 {
			c = 0
		}
		ans := 0
		for x := c / 3; x <= c/3+1; x++ {
			y := min(a-2*x, (b-x)/2)
			if y >= 0 && x+y > ans {
				ans = x + y
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1366A(os.Stdin, os.Stdout) }
