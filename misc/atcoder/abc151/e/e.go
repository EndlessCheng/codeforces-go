package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
const mod = 1_000_000_007

func run(in io.Reader, out io.Writer) {
	var n, k, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	c := 1
	for i := k - 1; i < n; i++ {
		ans = (ans + (a[i]-a[n-1-i])*c) % mod
		c = c * (i + 1) % mod * pow(i+2-k, mod-2) % mod
	}
	Fprint(out, (ans+mod)%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
