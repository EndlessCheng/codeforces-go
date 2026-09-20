package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf67C(in io.Reader, out io.Writer) {
	var a, b, c, d int
	var s, t string
	Fscan(in, &a, &b, &c, &d, &s, &t)

	n, m := len(s), len(t)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i := 1; i <= m; i++ {
		f[0][i] = f[0][i-1] + a
	}
	pt := [26]int{}
	for i := 1; i <= n; i++ {
		f[i][0] = f[i-1][0] + b
		ps := [26]int{}
		for j := 1; j <= m; j++ {
			x := c
			if s[i-1] == t[j-1] {
				x = 0
			}
			f[i][j] = min(f[i][j-1]+a, f[i-1][j]+b, f[i-1][j-1]+x)
			k, l := pt[t[j-1]-'a'], ps[s[i-1]-'a']
			if k > 0 && l > 0 {
				f[i][j] = min(f[i][j], f[k-1][l-1]+(i-k-1)*b+(j-l-1)*a+d)
			}
			ps[t[j-1]-'a'] = j
		}
		pt[s[i-1]-'a'] = i
	}
	Fprint(out, f[n][m])
}

//func main() { cf67C(bufio.NewReader(os.Stdin), os.Stdout) }
