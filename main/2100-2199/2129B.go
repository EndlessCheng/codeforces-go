package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2129B(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([]int, n+1)
		for i := range n {
			Fscan(in, &v)
			pos[v] = i
		}

		a := make([]int, n)
		for v = 1; v <= n; v++ {
			i := pos[v]
			inv1, inv2 := 0, 0
			for _, w := range a[:i] {
				if w == 0 || w > v {
					inv1++
				}
				if w > n*2-v {
					inv2++
				}
			}
			for _, w := range a[i+1:] {
				if w > 0 && w < v {
					inv1++
				}
				if w < n*2-v {
					inv2++
				}
			}
			if inv1 < inv2 {
				a[i] = v
			} else {
				a[i] = n*2 - v
			}
		}

		ans := 0
		for i, v := range a {
			for _, w := range a[:i] {
				if w > v {
					ans++
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2129B(bufio.NewReader(os.Stdin), os.Stdout) }
