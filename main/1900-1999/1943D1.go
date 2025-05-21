package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1943D1(in io.Reader, out io.Writer) {
	var T, n, mx, mod int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &mx, &mod)
		f := [2][][]int{}
		for i := range f {
			f[i] = make([][]int, mx+1)
			for j := range f[i] {
				f[i][j] = make([]int, mx+1)
			}
		}
		for pre, row := range f[0] {
			for pre2 := pre; pre2 <= mx; pre2++ {
				row[pre2] = 1
			}
		}
		// 在 f 上原地后缀和的写法见 https://codeforces.com/problemset/submission/1943/320574526
		suf := make([]int, mx+2)
		for i := 1; i <= n; i++ {
			for pre, row := range f[i%2] {
				for j := mx; j >= 0; j-- {
					suf[j] = suf[j+1] + f[(i-1)%2][j][pre]
				}
				for pre2 := range row {
					row[pre2] = suf[max(pre-pre2, 0)] % mod
				}
			}
		}
		Fprintln(out, f[n%2][0][0])
	}
}

//func main() { cf1943D1(bufio.NewReader(os.Stdin), os.Stdout) }
