package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF939E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, op, l int
	var v, s int64
	a := []int64{}
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &op); op == 1 {
			s -= v
			Fscan(in, &v)
			s += v
			for ; l < len(a) && a[l]*int64(l+1) < s; l++ {
				s += a[l]
			}
			a = append(a, v)
		} else {
			Fprintf(out, "%.8f\n", float64(v)-float64(s)/float64(l+1))
		}
	}
}

//func main() { CF939E(os.Stdin, os.Stdout) }
