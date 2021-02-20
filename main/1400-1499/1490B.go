package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1490B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		c := [3]int{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			c[v%3]++
		}
		Fprintln(out, max(max(c[1]-c[0], c[2]-c[1]), c[0]-c[2]))
	}
}

//func main() { CF1490B(os.Stdin, os.Stdout) }
