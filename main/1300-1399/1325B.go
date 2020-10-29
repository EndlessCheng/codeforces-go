package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1325B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		c := map[int]bool{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			c[v] = true
		}
		Fprintln(out, len(c))
	}
}

//func main() { CF1325B(os.Stdin, os.Stdout) }
