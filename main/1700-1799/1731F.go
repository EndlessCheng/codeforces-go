package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1731F(in io.Reader, out io.Writer) {
	const mod = 998244353
	const inv6 = 166374059
	var n, k int
	Fscan(in, &n, &k)
	if n == 1 {
		Fprint(out, 0)
		return
	}

	a := make([]int, n+1)
	a[2] = 1
	y := 1
	for i := 3; i <= n; i++ {
		y = y * k % mod
		a[i] = (a[i-1] + y*(i-1)) % mod
	}

	x := (k + 1) * k % mod
	x = x * (k - 1) % mod
	x = x * inv6 % mod
	Fprint(out, a[n]*x%mod)
}

//func main() { cf1731F(bufio.NewReader(os.Stdin), os.Stdout) }
