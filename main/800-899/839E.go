package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf839E(in io.Reader, out io.Writer) {
	var n, k, v int
	Fscan(in, &n, &k)
	g := make([]int, n)
	for i := range n {
		for j := range n {
			Fscan(in, &v)
			g[i] |= v << j
		}
	}

	memo := map[int]int{0: 0}
	var dfs func(int) int
	dfs = func(s int) int {
		if v, ok := memo[s]; ok {
			return v
		}
		tz := bits.TrailingZeros(uint(s))
		memo[s] = max(dfs(s&(s-1)), dfs(s&g[tz])+1)
		return memo[s]
	}
	ans := dfs(1<<n - 1)
	Fprintf(out, "%.6f", float64(k*k*(ans-1))/float64(ans*2))
}

//func main() { cf839E(bufio.NewReader(os.Stdin), os.Stdout) }
