package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf17D(in io.Reader, out io.Writer) {
	calcPhi := func(n int) int {
		phi := n
		for i := 2; i*i <= n; i++ {
			if n%i > 0 {
				continue
			}
			for n /= i; n%i == 0; n /= i {
			}
			phi = phi / i * (i - 1)
		}
		if n > 1 {
			phi = phi / n * (n - 1)
		}
		return phi
	}
	pow := func(x, n, p int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % p
			}
			x = x * x % p
		}
		return res
	}
	var s, t string
	var b, n, p int
	Fscan(in, &s, &t, &p)
	phiP := calcPhi(p)

	for _, d := range s {
		b = (b*10 + int(d-'0')) % p
	}

	gr := false
	for _, d := range t {
		n = n*10 + int(d-'0')
		if n > phiP {
			n %= phiP
			gr = true
		}
	}
	n--
	if gr {
		n = (n+phiP)%phiP + phiP
	}

	ans := (b - 1 + p) * pow(b, n, p) % p
	if ans == 0 {
		ans = p
	}
	Fprint(out, ans)
}

//func main() { cf17D(bufio.NewReader(os.Stdin), os.Stdout) }
