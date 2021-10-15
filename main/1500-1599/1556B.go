package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1556B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c1, s0, s1 := 0, 0, 0
		for i, p := 0, 0; i < n; i++ {
			if Fscan(in, &v); v&1 > 0 {
				c1++
				s0 += abs(i - p)
				s1 += abs(i - (p ^ 1))
				p += 2
			}
		}
		if abs(n-c1*2) > 1 {
			Fprintln(out, -1)
		} else if n&1 == 0 {
			Fprintln(out, min(s0, s1))
		} else if c1*2 > n {
			Fprintln(out, s0)
		} else {
			Fprintln(out, s1)
		}
	}
}

//func main() { CF1556B(os.Stdin, os.Stdout) }
