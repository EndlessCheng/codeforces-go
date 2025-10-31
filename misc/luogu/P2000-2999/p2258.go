package P2000_2999

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func p2258(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, m, r, c int
	Fscan(in, &n, &m, &r, &c)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	ans := int(1e9)
	for sub := 1; sub < 1<<n; sub++ {
		if bits.OnesCount(uint(sub)) != r {
			continue
		}
		b := [][]int{}
		for s := uint(sub); s > 0; s &= s - 1 {
			b = append(b, a[bits.TrailingZeros(s)])
		}
		col := make([]int, m)
		for j := 0; j < m; j++ {
			for i := 1; i < len(b); i++ {
				col[j] += abs(b[i][j] - b[i-1][j])
			}
		}
		row := make([][]int, m)
		for j := range row {
			row[j] = make([]int, j)
			for k := 0; k < j; k++ {
				for _, r := range b {
					row[j][k] += abs(r[k] - r[j])
				}
			}
		}
		f := slices.Clone(col)
		for i := 2; i <= c; i++ {
			for j := m - 1 - (c - i); j >= i-1; j-- {
				f[j] = 1e9
				for k := i - 2; k < j; k++ {
					f[j] = min(f[j], f[k]+row[j][k])
				}
				f[j] += col[j]
			}
		}
		ans = min(ans, slices.Min(f[c-1:]))
	}
	Fprint(out, ans)
}

//func main() { p2258(bufio.NewReader(os.Stdin), os.Stdout) }
