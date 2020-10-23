package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1409C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x, &y)
		for d := 1; ; d++ {
			if (y-x)%d == 0 && (y-x)/d < n {
				for v := y; v > 0 && n > 0; v -= d {
					Fprint(out, v, " ")
					n--
				}
				for v := y + d; n > 0; v += d {
					Fprint(out, v, " ")
					n--
				}
				break
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1409C(os.Stdin, os.Stdout) }
