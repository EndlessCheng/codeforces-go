package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1342F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		u := 1 << n
		sum := make([]int, u)
		for i := range n {
			Fscan(in, &v)
			hb := 1 << i
			for j, s := range sum[:hb] {
				sum[hb|j] = s + v
			}
		}

		type pair struct{ sz, minS int }
		type pr struct{ i, j int }
		f := make([][]pair, u)
		from := make([][]pr, u)
		for i, s := range sum {
			f[i] = make([]pair, n)
			for x := uint32(i); x > 0; x &= x - 1 {
				f[i][bits.TrailingZeros32(x)] = pair{1, s}
			}
			from[i] = make([]pr, n)
		}

		maxSave, j0 := 0, 0
		for s, fs := range f {
			for x := uint32(s); x > 0; x &= x - 1 {
				j := bits.TrailingZeros32(x)
				p := fs[j]
				sz := p.sz
				t := u - 1 ^ s
				for sub := t; sub>>j > 0; sub = (sub - 1) & t {
					if sum[sub] <= p.minS {
						continue
					}
					// 为了让序列尽可能长，当前组的 core 越靠前越好（为后面留出位置空间）
					k := bits.TrailingZeros32(uint32(sub >> j << j))
					ss := s | sub
					q := &f[ss][k]
					if sz+1 > q.sz || sz+1 == q.sz && sum[sub] < q.minS {
						*q = pair{sz + 1, sum[sub]}
						from[ss][k] = pr{s, j}
					}
				}
				if s == u-1 && sz > maxSave {
					maxSave, j0 = sz, j
				}
			}
		}

		Fprintln(out, n-maxSave)
		idx := make([]int, n)
		for i := range idx {
			idx[i] = i + 1
		}
		for i, j := u-1, j0; i > 0; {
			p := from[i][j]
			for s := uint(i ^ p.i ^ 1<<j); s > 0; s &= s - 1 {
				k := bits.TrailingZeros(s)
				Fprintln(out, idx[k], idx[j])
				for k++; k < n; k++ {
					idx[k]--
				}
			}
			i, j = p.i, p.j
		}
	}
}

//func main() { cf1342F(bufio.NewReader(os.Stdin), os.Stdout) }
