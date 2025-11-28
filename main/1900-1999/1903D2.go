package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1903D2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, v int
	Fscan(in, &n, &q)
	const w = 20
	const u = 1 << w
	low := n << w
	f := [u][w][2]int{}
	for range n {
		Fscan(in, &v)
		low -= v
		for j := uint32(u - 1 ^ v); j > 0; j &= j - 1 {
			i := bits.TrailingZeros32(j)
			f[v][i][0]++ // 恰好满足 v 但 i 位是 0 的数的个数
			f[v][i][1] += v & (1<<i - 1)
		}
	}

	// 算完 SOS DP 后，「恰好满足」就变成「至少满足 s 中的比特位是 1（其余比特位可能还有 1）」
	for i := range w {
		for s := 0; s < u; s++ {
			s |= 1 << i
			for j := range w {
				f[s^1<<i][j][0] += f[s][j][0]
				f[s^1<<i][j][1] += f[s][j][1]
			}
		}
	}

	for range q {
		Fscan(in, &v)
		if v >= low {
			Fprintln(out, u+(v-low)/n)
			continue
		}
		ans, cnt := 0, 0
		for i := w - 1; i >= 0; i-- {
			p := f[ans][i] // 至少满足 ans 但 i 位是 0 的数的个数
			c := (cnt+p[0])<<i - p[1] // 减掉已经有的
			if c <= v {
				v -= c
				cnt += p[0]
				ans |= 1 << i
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1903D2(bufio.NewReader(os.Stdin), os.Stdout) }
