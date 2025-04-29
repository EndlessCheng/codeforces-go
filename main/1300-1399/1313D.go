package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
func cf1313D(in io.Reader, out io.Writer) {
	var n, up, l, r int
	Fscan(in, &n, &up, &up)
	type pair struct{ v, i int }
	a := make([]pair, 0, n*2)
	for i := 1; i <= n; i++ {
		Fscan(in, &l, &r)
		a = append(a, pair{l<<1 | 1, i}, pair{(r + 1) << 1, i}) // 先出后进
	}
	slices.SortFunc(a, func(a, b pair) int { return a.v - b.v })

	f := make([]int, 1<<up)
	for i := 1; i < 1<<up; i++ {
		f[i] = -1e9
	}
	idx := make([]int, up)
	for ai, p := range a {
		d := 0
		if ai > 0 {
			d = p.v>>1 - a[ai-1].v>>1
		}
		k := 0
		if p.v&1 > 0 { // 入
			for i, j := range idx {
				if j == 0 {
					k = i
					idx[i] = p.i
					break
				}
			}
			for j := len(f) - 1; j >= 0; j-- {
				odd := bits.OnesCount8(uint8(j))&1 > 0
				if j>>k&1 > 0 { // 和 k 有关
					d2 := 0
					if !odd { // k 进来之前是奇数
						d2 = d
					}
					f[j] = f[j^1<<k] + d2
				} else if odd { // 和 k 无关
					f[j] += d
				}
			}
		} else { // 出
			for i, j := range idx {
				if j == p.i {
					k = i
					idx[i] = 0
					break
				}
			}
			for j := range f {
				if j>>k&1 > 0 { // 不合法
					f[j] = -1e9
				} else if bits.OnesCount8(uint8(j))&1 > 0 {
					f[j] = max(f[j]+d, f[j|1<<k])
				} else {
					f[j] = max(f[j], f[j|1<<k]+d)
				}
			}
		}
	}
	Fprint(out, f[0])
}

//func main() { cf1313D(bufio.NewReader(os.Stdin), os.Stdout) }
