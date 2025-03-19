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
	var n, v int
	const k = 1
	f := [k + 1][]int{}
	for i := range f {
		f[i] = []int{-2e9}
	}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		for i := k; i >= 0; i-- {
			j := sort.SearchInts(f[i], v)
			if j < len(f[i]) {
				f[i][j] = v
			} else {
				f[i] = append(f[i], v)
			}
			if i > 0 {
				// 注意这里没有二分，完全取决于 len(g)
				g := f[i-1]
				j = len(g)
				w := g[len(g)-1] + 1
				if j < len(f[i]) {
					f[i][j] = min(f[i][j], w)
				} else {
					f[i] = append(f[i], w)
				}
			}
		}
	}

	ans := 0
	for _, g := range f {
		ans = max(ans, len(g)-1)
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
