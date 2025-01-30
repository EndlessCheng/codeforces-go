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
	const mod = 1_000_000_007
	var s string
	var k int
	Fscan(in, &s, &k)
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, bool) int
	dfs = func(i, mask int, isLimit bool) (res int) {
		c := bits.OnesCount(uint(mask))
		if c > k {
			return 0
		}
		if i == len(s) {
			if c < k {
				return 0
			}
			return 1
		}
		if !isLimit && mask > 0 {
			p := &memo[i][c]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		if mask == 0 {
			res += dfs(i+1, 0, false)
		}
		up := 15
		if isLimit {
			if s[i] <= '9' {
				up = int(s[i] - '0')
			} else {
				up = 10 + int(s[i]-'A')
			}
		}
		d := 0
		if mask == 0 {
			d = 1
		}
		for ; d <= up; d++ {
			res += dfs(i+1, mask|1<<d, isLimit && d == up)
		}
		return res % mod
	}
	Fprint(out, dfs(0, 0, true))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
