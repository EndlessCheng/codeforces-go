package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2084G1(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		p := make([]int, n+1)
		for i := range p {
			p[i] = -1
		}

		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			p[a[i]] = i & 1
		}

		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, n+1)
			for j := range f[i] {
				f[i][j] = -1e18
			}
		}
		f[0][0] = 0
		for i := 1; i <= n; i++ {
			if p[i] != 1 {
				for j := range i {
					f[i][j] = max(f[i][j], f[i-1][j]+i*((i-j)+(n+1)/2-j))
				}
			}
			if p[i] != 0 {
				for j := 1; j <= i; j++ {
					f[i][j] = max(f[i][j], f[i-1][j-1]+i*(j+n/2-(i-j)))
				}
			}
		}
		Fprintln(out, f[n][(n+1)/2])
	}
}

//func main() { cf2084G1(bufio.NewReader(os.Stdin), os.Stdout) }
