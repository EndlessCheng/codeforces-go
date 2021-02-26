package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1393C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c, mx := make([]int, n+1), 0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if c[v]++; c[v] > mx {
				mx = c[v]
			}
		}
		cntMx := 0
		for _, v := range c {
			if v == mx {
				cntMx++
			}
		}
		Fprintln(out, cntMx-1+(n-mx*cntMx)/(mx-1))
	}
}

//func main() { CF1393C(os.Stdin, os.Stdout) }
