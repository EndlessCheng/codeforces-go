package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1554C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		m++ // n^mex >= m+1
		mex := 0
		for i := 29; i >= 0; i-- {
			x, y := n>>i&1, m>>i&1
			if x > 0 && y == 0 {
				break
			}
			if x == 0 && y > 0 {
				mex |= 1 << i
			}
		}
		Fprintln(out, mex)
	}
}

//func main() { CF1554C(os.Stdin, os.Stdout) }
