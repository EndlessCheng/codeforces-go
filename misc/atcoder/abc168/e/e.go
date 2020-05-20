package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	const mod int = 1e9 + 7
	pow := func(n int) int {
		res, x := 1, 2
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	type pair struct{ x, y int }
	var n, x, y int
	Fscan(in, &n)
	cnt := map[pair]int{}
	cnt2 := map[pair]int{}
	cntO := 0
	for i := 0; i < n; i++ {
		Fscan(in, &x, &y)
		if x == 0 && y == 0 {
			cntO++
			continue
		}
		if y < 0 {
			x = -x
			y = -y
		} else if y == 0 {
			x = abs(x)
		}
		g := gcd(abs(x), y)
		x /= g
		y /= g
		if x > 0 {
			cnt[pair{x, y}]++
		} else {
			cnt2[pair{y, -x}]++
		}
	}

	ans := 1
	for p, c := range cnt {
		ans = ans * (pow(c) + pow(cnt2[p]) - 1) % mod
		delete(cnt2, p)
	}
	for _, c := range cnt2 {
		ans = ans * (pow(c)) % mod
	}
	Fprint(_w, (ans+cntO-1+mod)%mod)
}

func main() { run(os.Stdin, os.Stdout) }
