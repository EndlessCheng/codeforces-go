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

	// 下面将原题中的编号 1 称为编号 0
	f := make([]float64, 1<<n)
	f[1] = 1 // 只有一个人（编号 0），此时编号 0 获胜的概率为 1
	for i := 3; i < 1<<n; i += 2 { // 只计算集合中有编号 0 的（即奇数），因为 f[集合不含编号 0] 一定是 0
		for s, lb := i, 0; s > 0; s ^= lb {
			lb = s & -s
			x := bits.TrailingZeros(uint(lb))
			for t, lb2 := s^lb, 0; t > 0; t ^= lb2 {
				lb2 = t & -t
				y := bits.TrailingZeros(uint(lb2))
				// 若 x 和 y 均不为 0，那么相当于先让 x 和 y 比，然后胜者后面去和编号 0 比
				// 枚举胜者是谁，然后相加
				f[i] = math.Max(f[i], f[i^lb]*p[y][x]+f[i^lb2]*p[x][y])
			}
		}
	}
	Fprintf(out, "%.8f", f[1<<n-1])
}

//func main() { CF678E(os.Stdin, os.Stdout) }
