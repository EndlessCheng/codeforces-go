package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var s string
	var k int
	Fscan(in, &s, &k)
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, k)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var f func(int, int, bool) int
	f = func(i, sum int, isLimit bool) (res int) {
		if i == len(s) {
			if sum > 0 {
				return
			}
			return 1
		}
		if !isLimit {
			p := &memo[i][sum]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for d := 0; d <= up; d++ {
			res += f(i+1, (sum+d)%k, isLimit && d == up)
		}
		res %= mod
		return
	}
	Fprint(out, (f(0, 0, true)+mod-1)%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
