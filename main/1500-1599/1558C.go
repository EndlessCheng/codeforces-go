package main

import (
	"bufio"
	. "fmt"
	"io"
	. "slices"
)

// https://space.bilibili.com/206214
func CF1558C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			a[i]--
			if a[i]%2 != i%2 {
				a[0] = -1
			}
		}
		if a[0] < 0 {
			Fprintln(out, -1)
			continue
		}
		ans := []any{}
		for i := n - 1; i > 0; i -= 2 {
			p1 := Index(a, i)
			Reverse(a[:p1+1])
			p2 := Index(a, i-1)
			Reverse(a[:p2])
			Reverse(a[:p2+2])
			Reverse(a[:3])
			Reverse(a[:i+1])
			ans = append(ans, p1+1, p2, p2+2, 3, i+1)
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { CF1558C(os.Stdin, os.Stdout) }
