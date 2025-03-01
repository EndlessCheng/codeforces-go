package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
	}
	mask := make([]int, n)
	for ; m > 0; m-- {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v][w] = wt
		g[w][v] = wt
		mask[v] |= 1 << w
		mask[w] |= 1 << v
	}

	ans := k
	mst := 1
	var dfs func(int, int)
	dfs = func(i, s int) {
		if i == n {
			ans = min(ans, s%k)
			return
		}
		for v := 0; v < n; v++ {
			if mst>>v&1 == 0 {
				continue
			}
			for j := uint8(mask[v] &^ mst); j > 0; j &= j - 1 {
				w := bits.TrailingZeros8(j)
				mst |= 1 << w
				dfs(i+1, s+g[v][w])
				mst ^= 1 << w
			}
		}
	}
	dfs(1, 0)
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
