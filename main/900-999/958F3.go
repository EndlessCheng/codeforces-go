package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
)

type fft958 struct {
	n               int
	omega, omegaInv []complex128
}

func newFFT958(n int) *fft958 {
	omega := make([]complex128, n)
	omegaInv := make([]complex128, n)
	for i := range omega {
		sin, cos := math.Sincos(2 * math.Pi * float64(i) / float64(n))
		omega[i] = complex(cos, sin)
		omegaInv[i] = complex(cos, -sin)
	}
	return &fft958{n, omega, omegaInv}
}

func (f *fft958) transform(a, omega []complex128) {
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

func (f *fft958) dft(a []complex128) {
	f.transform(a, f.omega)
}

func (f *fft958) idft(a []complex128) {
	f.transform(a, f.omegaInv)
	for i := range a {
		a[i] /= complex(float64(f.n), 0)
	}
}

func convolution958(a, b []int) []int {
	n, m := len(a), len(b)
	limit := 1 << uint(bits.Len(uint(n+m-1)))
	f := newFFT958(limit)
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

func convolutionN958(coefs [][]int) []int {
	var f func(l, r int) []int
	f = func(l, r int) []int {
		if l == r {
			return coefs[l-1]
		}
		mid := (l + r) >> 1
		return convolution958(f(l, mid), f(mid+1, r))
	}
	return f(1, len(coefs))
}

// github.com/EndlessCheng/codeforces-go
func Sol958F3(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

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
	Fprint(out, convolutionN958(coefs)[k])
}

//func main() {
//	Sol958F3(os.Stdin, os.Stdout)
//}
