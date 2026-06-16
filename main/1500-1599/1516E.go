package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1516E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
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

	var n, k int
	Fscan(in, &n, &k)
	b := make([]int, 2*k+1)
	b[0] = 1
	for i := 1; i <= 2*k; i++ {
		b[i] = b[i-1] * (n - i + 1) % mod * pow(i, mod-2) % mod
	}

	s := make([][]int, 2*k-1)
	for i := range s {
		s[i] = make([]int, k)
	}
	s[0][0] = 1
	for i := 1; i < 2*k-1; i++ {
		for j := 0; j <= i/2; j++ {
			s[i][j] = (i + 1) * s[i-1][j] % mod
			if j > 0 {
				s[i][j] = (s[i][j] + (i+1)*s[i-2][j-1]) % mod
			}
		}
	}

	d := make([]int, k+1)
	d[0] = 1
	if k >= 1 {
		d[1] = n * (n - 1) % mod * pow(2, mod-2) % mod
	}
	for i := 2; i <= k; i++ {
		d[i] = d[i-2]
		for t := 1; t <= i; t++ {
			d[i] = (d[i] + s[i+t-2][t-1]*b[i+t]) % mod
		}
	}

	for i := 1; i <= k; i++ {
		Fprint(out, d[i], " ")
	}
}

//func main() { cf1516E(bufio.NewReader(os.Stdin), os.Stdout) }
