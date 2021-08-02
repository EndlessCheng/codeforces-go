package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF678E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	p := make([][]float64, n)
	for i := range p {
		p[i] = make([]float64, n)
		for j := range p[i] {
			Fscan(in, &p[i][j])
		}
	}

	f := make([]float64, 1<<n)
	f[1] = 1
	for i := range f {
		for s, lb := i, 0; s > 0; s ^= lb {
			lb = s & -s
			x := bits.TrailingZeros(uint(lb))
			for t, lb2 := s^lb, 0; t > 0; t ^= lb2 {
				lb2 = t & -t
				y := bits.TrailingZeros(uint(lb2))
				f[i] = math.Max(f[i], f[i^lb]*p[y][x]+f[i^lb2]*p[x][y])
			}
		}
	}
	Fprintf(out, "%.8f", f[1<<n-1])
}

//func main() { CF678E(os.Stdin, os.Stdout) }
