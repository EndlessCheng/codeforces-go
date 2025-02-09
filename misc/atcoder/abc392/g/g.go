package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
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
			for i := 0; i < m; i++ {
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

func run(in io.Reader, out io.Writer) {
	const mx int = 1e6 + 1
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	cnt := make([]int, mx)
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	limit := 1 << bits.Len(uint(mx*2-1))
	A := make([]complex128, limit)
	for i, c := range cnt {
		A[i] = complex(float64(c), 0)
	}
	t := newFFT(limit)
	t.dft(A)
	for i := range A {
		A[i] *= A[i]
	}
	t.idft(A)
	C := make([]int, mx*2-1)
	for i := range C {
		C[i] = int(math.Round(real(A[i])))
	}
	for _, v := range a {
		// ans += (C[v*2] - 1) / 2
		ans += C[v*2] / 2
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
