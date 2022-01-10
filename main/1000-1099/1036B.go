package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1036B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, k int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		if k < n || k < m {
			Fprintln(out, -1)
		} else if n&1 != m&1 {
			Fprintln(out, k-1)
		} else {
			Fprintln(out, k-(k-n)&1*2)
		}
	}
}

//func main() { CF1036B(os.Stdin, os.Stdout) }
