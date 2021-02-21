package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1179B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	x1, y1, x2, y2, d := 1, 1, n, m, true
	for left := n * m; left > 0; left-- {
		Fprintln(out, x1, y1)
		if d {
			if y1 < m {
				y1++
			} else {
				x1++
			}
		} else {
			if y1 > 1 {
				y1--
			} else {
				x1++
			}
		}
		left--
		if left == 0 {
			break
		}
		Fprintln(out, x2, y2)
		if d {
			if y2 > 1 {
				y2--
			} else {
				x2--
				d = !d
			}
		} else {
			if y2 < m {
				y2++
			} else {
				x2--
				d = !d
			}
		}
	}
}

//func main() { CF1179B(os.Stdin, os.Stdout) }
