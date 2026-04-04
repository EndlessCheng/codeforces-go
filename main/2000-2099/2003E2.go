package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2003E2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		sum := make([]int, n+2)
		isL := make([]bool, n+2)
		isR := make([]bool, n+2)
		for range m {
			var l, r int
			Fscan(in, &l, &r)
			isL[l] = true
			isR[r] = true
			sum[l+1]++
			sum[r+1]--
		}
		for i := 1; i <= n; i++ {
			sum[i] += sum[i-1]
		}

		f := make([][][2]int, n+2)
		for i := range f {
			f[i] = make([][2]int, n+2)
			for j := range f[i] {
				f[i][j] = [2]int{-1e9, -1e9}
			}
		}
		f[0][0][0] = 0
		for i := 1; i <= n; i++ {
			for j := range i + 1 {
				if j > 0 && !isL[i] {
					f[i][j][1] = max(f[i-1][j-1][1], f[i-1][j-1][0]) + j - 1
				}
				if j != i && !isR[i] {
					f[i][j][0] = f[i-1][j][0] + i - 1
					if sum[i] == 0 {
						f[i][j][0] = max(f[i][j][0], f[i-1][j][1]+i-1)
					}
				}
			}
		}

		ans := -1
		for i := range n + 1 {
			ans = max(ans, f[n][i][0], f[n][i][1])
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2003E2(bufio.NewReader(os.Stdin), os.Stdout) }
