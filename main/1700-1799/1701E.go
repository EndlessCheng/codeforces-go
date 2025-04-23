package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1701E(in io.Reader, out io.Writer) {
	var T, n, m int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &s, &t)
		f := make([][3]int, n+1)
		nf := make([][3]int, n+1)
		f[0][0] = 1
		for i, x := range s {
			if i < m {
				f[i+1] = [3]int{1e9, 1e9, 1e9}
			}
			nf[0][0] = f[0][0] + 2
			nf[0][1] = nf[0][0]
			nf[0][2] = f[0][2] + 1
			for j, y := range t[:min(i+1, m)] {
				nf[j+1][0] = f[j+1][0] + 2
				nf[j+1][1] = nf[j+1][0]
				nf[j+1][2] = min(f[j+1][2]+1, nf[j+1][0])
				if x == y {
					nf[j+1][0] = min(nf[j+1][0], f[j][0] + 1)
					nf[j+1][1] = min(nf[j+1][1], f[j][1])
					nf[j+1][2] = min(nf[j+1][2], f[j][2]+1, f[j][1])
				}
			}
			f, nf = nf, f
		}
		ans := f[m][2]
		if ans > n {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1701E(bufio.NewReader(os.Stdin), os.Stdout) }
