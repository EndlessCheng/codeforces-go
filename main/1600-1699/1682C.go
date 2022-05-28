package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1682C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		c := map[int]int{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			c[v]++
		}
		s := 0
		for _, c := range c {
			if c > 2 {
				c = 2
			}
			s += c
		}
		Fprintln(out, (s+1)/2)
	}
}

//func main() { CF1682C(os.Stdin, os.Stdout) }
