package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1136C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v int
	Fscan(in, &n, &m)
	c := make([]map[int]int, n+m-1)
	for i := range c {
		c[i] = map[int]int{}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fscan(in, &v)
			c[i+j][v]++
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fscan(in, &v)
			c[i+j][v]--
		}
	}
	for _, c := range c {
		for _, c := range c {
			if c != 0 {
				Fprint(out, "NO")
				return
			}
		}
	}
	Fprint(out, "YES")
}

//func main() { CF1136C(os.Stdin, os.Stdout) }
