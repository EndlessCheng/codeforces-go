package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1763C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int64, n)
		mx := int64(0)
		for i := range a {
			Fscan(in, &a[i])
			mx = max(mx, a[i])
		}
		if n > 3 {
			Fprintln(out, n*mx)
		} else if n == 2 {
			Fprintln(out, max(a[0]+a[1], 2*abs(a[0]-a[1])))
		} else {
			Fprintln(out, max(a[0]+a[1]+a[2], max(max(3*a[0], 3*a[2]), max(3*abs(a[0]-a[1]), 3*abs(a[1]-a[2])))))
		}
	}
}

//func main() { CF1763C(os.Stdin, os.Stdout) }
