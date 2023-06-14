package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1700C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := int64(0)
		a0 := a[0]
		for i := 1; i < n; i++ {
			d := a[i] - a[i-1]
			ans += abs(d)
			if d < 0 {
				a0 += d
			}
		}
		ans += abs(a0)
		Fprintln(out, ans)
	}
}

//func main() { CF1700C(os.Stdin, os.Stdout) }
