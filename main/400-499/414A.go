package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF414A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	if k < n/2 {
		Fprint(out, -1)
		return
	}
	if n == 1 {
		if k > 0 {
			Fprint(out, -1)
		} else {
			Fprint(out, 1)
		}
		return
	}
	k -= n/2 - 1
	Fprint(out, k, k*2)
	for i := 1; i < n-1; i++ {
		if i != k && i != k*2 {
			Fprint(out, " ", i)
		} else {
			n++
		}
	}
}

//func main() { CF414A(os.Stdin, os.Stdout) }
