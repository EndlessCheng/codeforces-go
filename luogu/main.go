package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// Fast Fourier transform
// https://zhuanlan.zhihu.com/p/31584464
type fft struct {
	n               int
	omega, omegaInv []complex128
}

func newFFT(n int) *fft {
	omega := make([]complex128, n)
	omegaInv := make([]complex128, n)
	for i := range omega {
		sin, cos := math.Sincos(2 * math.Pi * float64(i) / float64(n))
		omega[i] = complex(cos, sin)
		omegaInv[i] = complex(cos, -sin)
	}
	return &fft{n, omega, omegaInv}
}

func (f *fft) transform(a, omega []complex128) {
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

func (f *fft) dft(a []complex128) {
	f.transform(a, f.omega)
}

func (f *fft) idft(a []complex128) {
	f.transform(a, f.omegaInv)
	for i := range a {
		a[i] /= complex(float64(f.n), 0)
	}
}

// 入参出参都是从低到高的系数
func convolution(a, b []int) []int {
	n, m := len(a), len(b)
	limit := n + m - 1
	f := newFFT(limit)
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
		conv[i] = int(math.Round(real(cmplxA[i])))
	}
	return conv
}

func solve(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, m := read()+1, read()+1
	a := make([]int, n)
	for i := range a {
		a[i] = read()
	}
	b := make([]int, m)
	for i := range b {
		b[i] = read()
	}
	c := convolution(a, b)
	Fprint(out, c[0])
	for _, v := range c[1:] {
		Fprint(out, " ", v)
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
