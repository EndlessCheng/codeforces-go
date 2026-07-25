package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2238C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pa := make([]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &pa[i])
			pa[i]--
		}

		h := make([]int, n)
		h2 := make([]int, n)
		ans := n
		for i := n - 1; i >= 0; i-- {
			ans += h2[i]
			p := pa[i]
			v := h[i] + 1
			if v > h[p] {
				h2[p] = h[p]
				h[p] = v
			} else if v > h2[p] {
				h2[p] = v
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2238C(bufio.NewReader(os.Stdin), os.Stdout) }
