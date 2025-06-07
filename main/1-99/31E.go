package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf31E(in io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(in, &n, &s)
	f := make([][]int, n*2+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i, b := range s {
		for j, fv := range f[i] {
			res := 0
			if j < n {
				res = f[i][j+1] + int(b-'0')*int(math.Pow10(j))
			}
			if j > n-1-i {
				res = max(res, fv+int(b-'0')*int(math.Pow10(n*2-1-i-j)))
			}
			f[i+1][j] = res
		}
	}

	ans := make([]byte, n*2)
	for i, j := n*2-1, 0; i >= 0; i-- {
		if j < n && f[i+1][j] == f[i][j+1]+int(s[i]-'0')*int(math.Pow10(j)) {
			ans[i] = 'M'
			j++
		} else {
			ans[i] = 'H'
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { cf31E(os.Stdin, os.Stdout) }
