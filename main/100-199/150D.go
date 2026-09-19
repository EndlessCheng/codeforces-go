package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf150D(in io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(in, &n)
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	Fscan(in, &s)
	f := make([][][]int, n+2)
	for i := range f {
		f[i] = make([][]int, n+2)
		for j := range f[i] {
			f[i][j] = make([]int, n+2)
			for k := range f[i][j] {
				f[i][j][k] = -1e18
			}
		}
	}
	for i := 1; i <= n; i++ {
		f[i+1][i][0] = 0
		f[i][i][1] = 0
		if a[1] != -1 {
			f[i][i][0] = a[1]
		}
	}

	for i := n; i > 0; i-- {
		for j := i + 1; j <= n; j++ {
			for k := range j - i + 2 {
				if k >= 2 && s[i-1] == s[j-1] {
					f[i][j][k] = max(f[i][j][k], f[i+1][j-1][k-2])
				}
				for l := i; l < j; l++ {
					f[i][j][k] = max(f[i][j][k], f[i][l][k]+f[l+1][j][0], f[i][l][0]+f[l+1][j][k])
				}
				if k > 0 && a[k] != -1 {
					f[i][j][0] = max(f[i][j][0], f[i][j][k]+a[k])
				}
			}
		}
	}

	mx := make([]int, n+1)
	for i := 1; i <= n; i++ {
		mx[i] = mx[i-1]
		for j := 1; j <= i; j++ {
			mx[i] = max(mx[i], mx[j-1]+f[j][i][0])
		}
	}
	Fprint(out, mx[n])
}

//func main() { cf150D(bufio.NewReader(os.Stdin), os.Stdout) }
