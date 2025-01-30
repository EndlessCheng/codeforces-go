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
	var dfs func(int, int, bool, bool) int
	dfs = func(i, mask int, isLimit, isNum bool) (res int) {
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
		if !isLimit && isNum {
			p := &memo[i][c]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		if !isNum {
			res += dfs(i+1, mask, false, false)
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
		if !isNum {
			d = 1
		}
		for ; d <= up; d++ {
			res += dfs(i+1, mask|1<<d, isLimit && d == up, true)
		}
		return res % mod
	}
	Fprint(out, dfs(0, 0, true, false))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
