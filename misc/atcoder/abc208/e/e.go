package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var s string
	var k int
	Fscan(in, &s, &k)
	type pair struct{ i, mul int }
	memo := map[pair]int{}
	var dfs func(int, int, bool, bool) int
	dfs = func(i, mul int, isLimit, isNum bool) (res int) {
		if i == len(s) {
			if isNum && mul <= k {
				return 1
			}
			return 0
		}
		if !isLimit && isNum {
			p := pair{i, mul}
			if r, ok := memo[p]; ok {
				return r
			}
			defer func() { memo[p] = res }()
		}
		if !isNum {
			res += dfs(i+1, mul, false, false)
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		d := 0
		if !isNum {
			d = 1
		}
		for ; d <= up; d++ {
			res += dfs(i+1, mul*d, isLimit && d == up, true)
		}
		return
	}
	Fprint(out, dfs(0, 1, true, false))
}

func main() { run(os.Stdin, os.Stdout) }
