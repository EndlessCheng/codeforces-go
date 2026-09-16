package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf976D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	d := make([]int, n+2)
	for i := 1; i <= n; i++ {
		Fscan(in, &d[i])
	}

	var v, w []int
	for i := 1; i <= n; i++ {
		for s := d[i]; s > d[i-1]; s-- {
			for t := d[n+1-i] + 1; t > s; t-- {
				v = append(v, s)
				w = append(w, t)
			}
		}
	}

	Fprintln(out, len(v))
	for i := range v {
		Fprintln(out, v[i], w[i])
	}
}

//func main() { cf976D(bufio.NewReader(os.Stdin), os.Stdout) }
