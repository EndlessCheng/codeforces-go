package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF448C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var f func([]int, int) int
	f = func(a []int, low int) int {
		if len(a) == 0 {
			return 0
		}
		mi := int(1e9)
		for _, v := range a {
			mi = min(mi, v)
		}
		c, p := mi-low, 0
		for i, v := range a {
			if v == mi {
				c += f(a[p:i], mi)
				p = i + 1
			}
		}
		c += f(a[p:], mi)
		return min(c, len(a))
	}
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fprint(out, f(a, 0))
}

//func main() { CF448C(os.Stdin, os.Stdout) }
