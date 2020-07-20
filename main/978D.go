package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF978D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n int
	Fscan(in, &n)
	if n < 3 {
		Fprint(out, 0)
		return
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := n + 1
	for i := -1; i <= 1; i++ {
	o:
		for j := -1; j <= 1; j++ {
			c, d := 0, a[1]+j-a[0]-i
			for k, v := range a {
				f := abs(a[0] + i + k*d - v)
				if f > 1 {
					continue o
				}
				if f == 1 {
					c++
				}
			}
			if c < ans {
				ans = c
			}
		}
	}
	if ans == n+1 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF978D(os.Stdin, os.Stdout) }
