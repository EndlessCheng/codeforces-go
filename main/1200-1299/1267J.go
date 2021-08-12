package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1267J(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c := make([]int, n)
		for ; n > 0; n-- {
			Fscan(in, &v)
			c[v-1]++
		}
		sort.Ints(c)
		c = c[sort.SearchInts(c, 1):]
	o:
		for sz := c[0] + 1; ; sz-- {
			for _, c := range c {
				if (sz-c%sz)%sz > (c-1)/sz+1 {
					continue o
				}
			}
			ans := len(c)
			for _, c := range c {
				ans += (c - 1) / sz
			}
			Fprintln(out, ans)
			break
		}
	}
}

//func main() { CF1267J(os.Stdin, os.Stdout) }
