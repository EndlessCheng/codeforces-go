package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n, d int
	Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	cnt := [1e6 + 1]int{}
	ans := 1
	dup := 1
	j := 0
	for i, v := range a {
		cnt[v]++
		dup = dup * cnt[v] % mod
		for a[j] < v-d {
			j++
		}
		ans = ans * (i - j + 1) % mod
	}
	Fprint(out, ans*pow(dup, mod-2)%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
