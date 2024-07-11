package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3901(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, v, l, r int
	Fscan(in, &n, &q)
	pre := make([]int, n+1)
	left := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		left[i] = max(left[i-1], pre[v])
		pre[v] = i
	}
	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		if left[r] < l {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { p3901(bufio.NewReader(os.Stdin), os.Stdout) }
