package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1537B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &x, &x)
		Fprintln(out, 1, 1, n, m)
	}
}

//func main() { CF1537B(os.Stdin, os.Stdout) }
