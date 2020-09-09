package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1407A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c := [2]int{}
		for ; n > 0; n-- {
			Fscan(in, &v)
			c[v]++
		}
		if c[0] >= c[1] {
			Fprintln(out, c[0])
			Fprintln(out, strings.Repeat("0 ", c[0]))
		} else {
			c[1] &^= 1
			Fprintln(out, c[1])
			Fprintln(out, strings.Repeat("1 ", c[1]))
		}
	}
}

//func main() { CF1407A(os.Stdin, os.Stdout) }
