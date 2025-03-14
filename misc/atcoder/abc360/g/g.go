package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	const k = 1
	f := [k + 1][]int{}
	f[0] = []int{a[0]}
	for j := 1; j <= k; j++ {
		f[j] = []int{0}
	}
	for _, v := range a[1:] {
		for j := k; j >= 0; j-- {
			p := sort.SearchInts(f[j], v)
			if p < len(f[j]) {
				f[j][p] = v
			} else {
				f[j] = append(f[j], v)
			}
			if j > 0 {
				g := f[j-1]
				p = len(g)
				w := g[len(g)-1] + 1
				if p < len(f[j]) {
					f[j][p] = min(f[j][p], w)
				} else {
					f[j] = append(f[j], w)
				}
			}
		}
	}

	ans := 0
	for _, g := range f {
		ans = max(ans, len(g))
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
