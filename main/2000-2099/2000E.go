package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2000E(in io.Reader, out io.Writer) {
	var T, n, m, k, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k, &w)
		a := make([]int, w)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.SortFunc(a, func(a, b int) int { return b - a })
		b := make([]int, 0, n*m)
		for i := range n {
			for j := range m {
				b = append(b, (min(i, n-k)-max(i-k, -1))*(min(j, m-k)-max(j-k, -1)))
			}
		}
		slices.SortFunc(b, func(a, b int) int { return b - a })
		ans := 0
		for i, v := range a {
			ans += v * b[i]
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2000E(bufio.NewReader(os.Stdin), os.Stdout) }
