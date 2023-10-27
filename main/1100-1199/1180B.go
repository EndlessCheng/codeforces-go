package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1180B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, t int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] >= 0 {
			a[i] = -a[i] - 1
		}
		if a[i] < a[t] {
			t = i
		}
	}
	if n%2 > 0 {
		a[t] = -a[t] - 1
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF1180B(os.Stdin, os.Stdout) }
