package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1095C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, k)
	for i := range a {
		a[i] = 1
	}
	n -= k
	for i := range a {
		for ; a[i] <= n; a[i] <<= 1 {
			n -= a[i]
		}
		if n == 0 {
			Fprintln(out, "YES")
			for _, v := range a {
				Fprint(out, v, " ")
			}
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { CF1095C(os.Stdin, os.Stdout) }
