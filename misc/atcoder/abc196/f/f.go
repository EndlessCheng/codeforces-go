package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
const P = 998244353

func pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % P
		}
		x = x * x % P
	}
	return
}

var omega, omegaInv [31]int

func init() {
	const g, invG = 3, 332748118
	for i := 1; i < len(omega); i++ {
		omega[i] = pow(g, (P-1)/(1<<i))
		omegaInv[i] = pow(invG, (P-1)/(1<<i))
	}
}

type ntt struct {
	n    int
	invN int
}

func (t ntt) transform(a, omega []int) {
	for i, j := 0, 0; i < t.n; i++ {
		if i > j {
			a[i], a[j] = a[j], a[i]
		}
		for l := t.n >> 1; ; l >>= 1 {
			j ^= l
			if j >= l {
				break
			}
		}
	}
	for l, li := 2, 1; l <= t.n; l <<= 1 {
		m := l >> 1
		wn := omega[li]
		li++
		for st := 0; st < t.n; st += l {
			b := a[st:]
			for i, w := 0, 1; i < m; i++ {
				d := b[m+i] * w % P
				b[m+i] = (b[i] - d + P) % P
				b[i] = (b[i] + d) % P
				w = w * wn % P
			}
		}
	}
}

func (t ntt) dft(a []int) {
	t.transform(a, omega[:])
}

func (t ntt) idft(a []int) {
	t.transform(a, omegaInv[:])
	for i, v := range a {
		a[i] = v * t.invN % P
	}
}

type poly []int

func (a poly) resize(n int) poly {
	b := make(poly, n)
	copy(b, a)
	return b
}

func (a poly) conv(b poly) poly {
	n, m := len(a), len(b)
	limit := 1 << bits.Len(uint(n+m-1))
	A := a.resize(limit)
	B := b.resize(limit)
	t := ntt{limit, pow(limit, P-2)}
	t.dft(A)
	t.dft(B)
	for i, v := range A {
		A[i] = v * B[i] % P
	}
	t.idft(A)
	return A[:n+m-1]
}

func run(in io.Reader, out io.Writer) {
	var s, t []byte
	Fscan(in, &s, &t)
	n, m := len(s), len(t)

	a := make(poly, n)
	sum := make([]int, n+1)
	for i, c := range s {
		a[i] = int(c - '0')
		sum[i+1] = sum[i] + a[i]
	}

	b := make(poly, m)
	sumT := 0
	for i, c := range t {
		b[i] = int(c - '0')
		sumT += b[i]
	}

	ans := m
	slices.Reverse(b)
	for i, v := range a.conv(b)[m-1 : n] {
		ans = min(ans, sum[i+m]-sum[i]+sumT-v*2)
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
