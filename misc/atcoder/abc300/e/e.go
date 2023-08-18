package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	const inv5 = 598946612
	var n int
	Fscan(in, &n)
	memo := map[int]int{}
	var f func(int) int
	f = func(x int) (res int) {
		if x == 1 {
			return 1
		}
		if v, ok := memo[x]; ok {
			return v
		}
		for i := 2; i <= 6; i++ {
			if x%i == 0 {
				res += f(x / i)
			}
		}
		res = res * inv5 % mod
		memo[x] = res
		return
	}
	Fprint(out, f(n))
}

func main() { run(os.Stdin, os.Stdout) }
