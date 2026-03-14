package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1685C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		n *= 2
		sum := make([]int, n+1)
		l, r := -1, -1
		mxI := 0
		for i, b := range s {
			sum[i+1] = sum[i] + 1 - int(b%2*2)
			if sum[i+1] < 0 {
				if l < 0 {
					l = i + 1
				}
				r = i + 1
			}
			if sum[i+1] > sum[mxI] {
				mxI = i + 1
			}
		}

		if l == -1 {
			Fprintln(out, 0)
			continue
		}

		mx := 0
		x, y := 0, n
		for i, v := range sum {
			if i < l && v > sum[x] {
				x = i
			}
			if i > r && v > sum[y] {
				y = i
			}
			if l <= i && i <= r {
				mx = max(mx, sum[i])
			}
		}

		if sum[x]+sum[y] >= mx {
			Fprintln(out, 1)
			Fprintln(out, x+1, y)
		} else {
			Fprintln(out, 2)
			Fprintln(out, 1, mxI)
			Fprintln(out, mxI+1, n)
		}
	}
}

//func main() { cf1685C(bufio.NewReader(os.Stdin), os.Stdout) }
