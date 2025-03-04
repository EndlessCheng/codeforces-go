package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func cf835D(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	n := len(s)
	isPal := make([][]bool, n)
	f := make([][]int, n)
	for i := range f {
		isPal[i] = make([]bool, n)
		for j := range isPal[i] {
			isPal[i][j] = true
		}
		f[i] = make([]int, n)
	}

	ans := make([]int, n+1)
	ans[1] = n
	for l := n - 2; l >= 0; l-- {
		f[l][l] = 1
		for r := l + 1; r < n; r++ {
			isPal[l][r] = s[l] == s[r] && isPal[l+1][r-1]
			if isPal[l][r] {
				f[l][r] = f[l][(l+r-1)/2] + 1
				ans[f[l][r]]++
			}
		}
	}
	for i := n; i > 1; i-- {
		ans[i-1] += ans[i]
	}
	for _, v := range ans[1:] {
		Fprint(out, v, " ")
	}
}

//func main() { cf835D(os.Stdin, os.Stdout) }
