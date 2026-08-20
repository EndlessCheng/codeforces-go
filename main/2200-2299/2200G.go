package main

import (
	. "fmt"
	"io"
	"strconv"
)

// https://github.com/EndlessCheng
func cf2200G(in io.Reader, out io.Writer) {
	const M = 1_000_000_007
	const mx = 3002
	fac := [mx]int{1}
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % M
	}
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % M
			}
			x = x * x % M
		}
		return res
	}

	var T, n, x int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		sum := 0
		f := []int{1}
		for range n {
			Fscan(in, &s)
			y, _ := strconv.Atoi(s[1:])
			switch s[0] {
			case '+':
				sum += y
				continue
			case '-':
				sum -= y
				continue
			case '/':
				y = pow(y, M-2)
			}
			x = x * y % M
			f = append(f, 0)
			for i := len(f) - 1; i > 0; i-- {
				f[i] = (f[i] + f[i-1]*y) % M
			}
		}

		m := len(f) - 1
		eMul := 0
		for i, fi := range f {
			eMul = (eMul + fi*fac[i]%M*fac[m-i]) % M
		}
		eMul = eMul * pow(fac[m+1], M-2) % M
		Fprintln(out, ((sum%M+M)*eMul+x)%M)
	}
}

//func main() { cf2200G(bufio.NewReader(os.Stdin), os.Stdout) }
