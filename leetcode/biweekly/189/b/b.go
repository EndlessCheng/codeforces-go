package main

import (
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
func minOperations1(s string) int {
	n := len(s)
	ans := math.MaxInt
	for rot := range n {
		op := rot
		for i := range n / 2 {
			d := abs(int(s[(rot+i)%n]) - int(s[(rot+n-1-i)%n]))
			op += min(d, 26-d)
		}
		ans = min(ans, op)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//

type fft struct {
	n        int
	omega    []complex128
	omegaInv []complex128
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

func (t *fft) transform(a, omega []complex128) {
	n := t.n
	for i, j := 0, 0; i < n; i++ {
		if i > j { // 保证同一对元素只交换一次
			a[i], a[j] = a[j], a[i]
		}
		for l := n / 2; ; l /= 2 {
			j ^= l
			if j >= l {
				break
			}
		}
	}
	for l := 2; l <= n; l *= 2 {
		m := l / 2
		for st := 0; st < n; st += l {
			b := a[st:]
			for i := range m {
				v := omega[n/l*i] * b[m+i]
				b[m+i] = b[i] - v
				b[i] += v
			}
		}
	}
}

func (t *fft) dft(a []complex128) {
	t.transform(a, t.omega)
}

func (t *fft) idft(a []complex128) {
	t.transform(a, t.omegaInv)
	cn := complex(float64(t.n), 0)
	for i := range a {
		a[i] /= cn
	}
}

// 计算 a 的自卷积
func selfPolyConvFFT(a []int) []int {
	n := len(a)
	limit := 1 << bits.Len(uint(n*2-1))
	A := make([]complex128, limit)
	for i, v := range a {
		A[i] = complex(float64(v), 0)
	}

	t := newFFT(limit)
	t.dft(A)
	for i, x := range A {
		A[i] *= x
	}
	t.idft(A)

	conv := make([]int, n*2-1)
	for i := range conv {
		conv[i] = int(math.Round(real(A[i])))
	}
	return conv
}

// 计算 a 的循环自卷积
func selfCyclicConvFFT(a []int) []int {
	n := len(a)
	conv := selfPolyConvFFT(a)
	for k := range n - 1 {
		conv[k] += conv[n+k]
	}
	return conv[:n]
}

func minOperations(s string) int {
	n := len(s)
	convSum := make([]int, n)
	a := make([]int, n)
	total := 0
	for k := range 13 {
		for i := range n {
			x := int(s[i] - 'a')
			if k <= x && x < k+13 {
				a[i] = 1
				total++
			} else {
				a[i] = 0
			}
		}
		c := selfCyclicConvFFT(a)
		for i, v := range c {
			convSum[i] += v
		}
	}

	ans := math.MaxInt
	for rot := range n {
		c := (rot*2 - 1 + n) % n
		ans = min(ans, rot-convSum[c])
	}
	return ans + total
}
