package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2075D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const u = 58
	var f, f2 [u][u]int
	for i := range f {
		for j := range f[i] {
			f[i][j] = 8e18
			f2[i][j] = 8e18
		}
	}
	f[0][0] = 0
	f2[0][0] = 0
	for i := 1; i < u; i++ {
		for j := u - 1; j >= 0; j-- {
			for k := u - 1; k >= 0; k-- {
				// 至少装满
				f[j][k] = min(f[j][k], f[j][max(k-i, 0)]+1<<i, f[max(j-i, 0)][k]+1<<i)
				// 恰好装满
				if k >= i {
					f2[j][k] = min(f2[j][k], f2[j][k-i]+1<<i)
				}
				if j >= i {
					f2[j][k] = min(f2[j][k], f2[j-i][k]+1<<i)
				}
			}
		}
	}

	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)
		if x > y {
			x, y = y, x
		}
		n, m := bits.Len(uint(x)), bits.Len(uint(y))
		ans := f[n][m]
		for a := bits.Len(uint(x ^ y>>(m-n))); a < n; a++ {
			ans = min(ans, f2[a][a+m-n])
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2075D(bufio.NewReader(os.Stdin), os.Stdout) }
