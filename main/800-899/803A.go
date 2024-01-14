package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf803A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)
	if n*n < k {
		Fprint(out, -1)
		return
	}

	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}
	for i, r := range a {
		if k > 0 {
			r[i] = 1
			k--
			for j := i + 1; j < n && k > 1; j++ {
				r[j] = 1
				a[j][i] = 1
				k -= 2
			}
		}
		for _, v := range r {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf803A(os.Stdin, os.Stdout) }
