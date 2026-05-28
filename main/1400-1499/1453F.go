package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1453F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			a[i] += i
		}

		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, n+1)
		}
		for i := 2; i <= n; i++ {
			for j := 1; j <= n; j++ {
				f[i][j] = 1e9
			}
			cnt := 0
			for j := i - 1; j >= 1; j-- {
				if a[j] >= i {
					f[i][a[j]] = min(f[i][a[j]], f[j][i-1]+cnt)
					cnt++
				}
			}
			for j := i + 1; j <= n; j++ {
				f[i][j] = min(f[i][j], f[i][j-1])
			}
		}
		Fprintln(out, f[n][n])
	}
}

//func main() { cf1453F(bufio.NewReader(os.Stdin), os.Stdout) }
