package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2050E(in io.Reader, out io.Writer) {
	var T int
	var a, b, c []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)
		m := len(b)
		f := make([]int, m+1)
		for j, y := range b {
			f[j+1] = f[j] + b2i50(y != c[j])
		}
		for i, x := range a {
			f[0] += b2i50(x != c[i])
			for j, y := range b {
				f[j+1] = min(f[j+1]+b2i50(x != c[i+j+1]), f[j]+b2i50(y != c[i+j+1]))
			}
		}
		Fprintln(out, f[m])
	}
}

//func main() { cf2050E(bufio.NewReader(os.Stdin), os.Stdout) }
func b2i50(b bool) int { if b { return 1 }; return 0 }

//

func cf2050E2(in io.Reader, out io.Writer) {
	var T int
	var a, b, c []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c)
		n, m := len(a), len(b)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, m+1)
		}
		for j, y := range b {
			f[0][j+1] = f[0][j] + b2i50(y != c[j])
		}
		for i, x := range a {
			f[i+1][0] = f[i][0] + b2i50(x != c[i])
			for j, y := range b {
				f[i+1][j+1] = min(f[i][j+1]+b2i50(x != c[i+j+1]), f[i+1][j]+b2i50(y != c[i+j+1]))
			}
		}
		Fprintln(out, f[n][m])
	}
}
