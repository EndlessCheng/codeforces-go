package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	const m int = 1e9 + 7
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % m
			}
			x = x * x % m
		}
		return res
	}
	var n, k, s int
	Fscan(_r, &n, &k)
	ans := make([]int, k+1)
	for i := k; i > 0; i-- {
		v := pow(k/i, n)
		for j := 2 * i; j <= k; j += i {
			v -= ans[j]
		}
		ans[i] = v % m
		s += i * ans[i]
	}
	Fprint(_w, (s%m+m)%m)
}

func main() { run(os.Stdin, os.Stdout) }
