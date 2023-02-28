package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF27C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	mx := make([]int, n)
	mn := make([]int, n)
	Fscan(in, &a[0])
	for i := 1; i < n; i++ {
		Fscan(in, &a[i])
		if a[i] > a[mx[i-1]] {
			mx[i] = i
		} else {
			mx[i] = mx[i-1]
		}
		if a[i] < a[mn[i-1]] {
			mn[i] = i
		} else {
			mn[i] = mn[i-1]
		}
	}
	up, low := n-1, n-1
	for i := n - 2; i > 0; i-- {
		if a[mn[i-1]] < a[i] && a[i] > a[low] {
			Fprintln(out, 3)
			Fprintln(out, mn[i-1]+1, i+1, low+1)
			return
		}
		if a[mx[i-1]] > a[i] && a[i] < a[up] {
			Fprintln(out, 3)
			Fprintln(out, mx[i-1]+1, i+1, up+1)
			return
		}
	}
	Fprint(out, 0)
}

//func main() { CF27C(os.Stdin, os.Stdout) }
