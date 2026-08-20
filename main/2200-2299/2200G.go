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

	var T, n, v0 int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &v0)
		sum := 0
		f := []int{1}
		for range n {
			Fscan(in, &s)
			v, _ := strconv.Atoi(s[1:])
			switch s[0] {
			case '+':
				sum += v
				continue
			case '-':
				sum -= v
				continue
			case '/':
				v = pow(v, M-2)
			}
			v0 = v0 * v % M
			f = append(f, 0)
			for i := len(f) - 1; i > 0; i-- {
				f[i] = (f[i] + f[i-1]*v) % M
			}
		}

		ans := 0
		for i, v := range f {
			ans = (ans + v*fac[i]%M*fac[len(f)-1-i]) % M
		}
		Fprintln(out, (ans*(sum%M+M)%M*pow(fac[len(f)], M-2)+v0)%M)
	}
}

//func main() { cf2200G(bufio.NewReader(os.Stdin), os.Stdout) }
