package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1422A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, a, b, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)
		Fprintln(out, max(max(a, b), c))
	}
}

//func main() { CF1422A(os.Stdin, os.Stdout) }
