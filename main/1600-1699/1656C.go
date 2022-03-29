package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1656C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c := map[int]int{}
		for ; n > 0; n-- {
			Fscan(in, &v)
			c[v]++
		}
		if c[1] > 0 {
			for v := range c {
				if c[v-1] > 0 {
					Fprintln(out, "NO")
					continue o
				}
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1656C(os.Stdin, os.Stdout) }
