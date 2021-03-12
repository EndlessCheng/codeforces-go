package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
type fft58 struct {
	n               int
	omega, omegaInv []complex128
}

func newFFT58(n int) *fft58 {
	omega := make([]complex128, n)
	omegaInv := make([]complex128, n)
	for i := range omega {
		sin, cos := math.Sincos(2 * math.Pi * float64(i) / float64(n))
		omega[i] = complex(cos, sin)
		omegaInv[i] = complex(cos, -sin)
	}
	return &fft58{n, omega, omegaInv}
}

func (f *fft58) transform(a, omega []complex128) {
	for i, j := 0, 0; i < f.n; i++ {
		if i > j {
			a[i], a[j] = a[j], a[i]
		}
		for l := f.n >> 1; ; l >>= 1 {
			j ^= l
			if j >= l {
				break
			}
		}
	}
	for l := 2; l <= f.n; l <<= 1 {
		m := l >> 1
		for st := 0; st < f.n; st += l {
			p := a[st:]
			for i := 0; i < m; i++ {
				t := omega[f.n/l*i] * p[m+i]
				p[m+i] = p[i] - t
				p[i] += t
			}
		}
	}
}

func (f *fft58) dft(a []complex128) {
	f.transform(a, f.omega)
}

func (f *fft58) idft(a []complex128) {
	f.transform(a, f.omegaInv)
	for i := range a {
		a[i] /= complex(float64(f.n), 0)
	}
}

func convolution58(a, b []int) []int {
	n, m := len(a), len(b)
	limit := 1 << uint(bits.Len(uint(n+m-1)))
	f := newFFT58(limit)
	cmplxA := make([]complex128, limit)
	for i, v := range a {
		cmplxA[i] = complex(float64(v), 0)
	}
	cmplxB := make([]complex128, limit)
	for i, v := range b {
		cmplxB[i] = complex(float64(v), 0)
	}
	f.dft(cmplxA)
	f.dft(cmplxB)
	for i := range cmplxA {
		cmplxA[i] *= cmplxB[i]
	}
	f.idft(cmplxA)
	conv := make([]int, n+m-1)
	for i := range conv {
		conv[i] = int(int64(math.Round(real(cmplxA[i]))) % 1009)
	}
	return conv
}

func convolutionN58(coefs [][]int) []int {
	n := len(coefs)
	if n == 1 {
		return coefs[0]
	}
	return convolution58(convolutionN58(coefs[:n/2]), convolutionN58(coefs[n/2:]))
}

func CF958F3(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, x int
	Fscan(in, &n, &m, &k)
	coefs := make([][]int, m)
	for i := range coefs {
		coefs[i] = []int{1}
	}
	for ; n > 0; n-- {
		Fscan(in, &x)
		x--
		coefs[x] = append(coefs[x], 1)
	}
	Fprint(out, convolutionN58(coefs)[k])
}

//func main() { CF958F3(os.Stdin, os.Stdout) }
