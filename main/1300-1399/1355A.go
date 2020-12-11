package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1355A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var v, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &v, &n)
	o:
		for ; n > 1; n-- {
			mi, mx := v%10, v%10
			for w := v; w > 0; w /= 10 {
				x := w % 10
				if x == 0 {
					break o
				}
				if x < mi {
					mi = x
				} else if x > mx {
					mx = x
				}
			}
			v += mi * mx
		}
		Fprintln(out, v)
	}
}

//func main() { CF1355A(os.Stdin, os.Stdout) }
