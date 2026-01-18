package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1380G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
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

	var n int
	Fscan(in, &n)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
	}
	slices.Sort(s)
	for i := 1; i < n; i++ {
		s[i+1] += s[i]
	}

	invN := pow(n, mod-2)
	for k := 1; k <= n; k++ {
		ans := 0
		for i := n - k; i > 0; i -= k {
			ans += s[i]
		}
		Fprint(out, ans%mod*invN%mod, " ")
	}
}

//func main() { cf1380G(bufio.NewReader(os.Stdin), os.Stdout) }
