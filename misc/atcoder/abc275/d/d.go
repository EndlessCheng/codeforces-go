package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	dp := map[int]int{0: 1}
	var f func(int) int
	f = func(x int) int {
		if v, ok := dp[x]; ok {
			return v
		}
		res := f(x/2) + f(x/3)
		dp[x] = res
		return res
	}
	Fprint(out, f(n))
}

func main() { run(os.Stdin, os.Stdout) }
