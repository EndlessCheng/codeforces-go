package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1336C(in io.Reader, out io.Writer) {
	const mod = 998244353
	var s, t string
	Fscan(bufio.NewReader(in), &s, &t)
	n, m := len(s), len(t)
	f := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		f[i] = make([]int, n)
		if i >= m || s[0] == t[i] {
			f[i][i] = 2
		}
		for j := i + 1; j < n; j++ {
			if i >= m || s[j-i] == t[i] {
				f[i][j] = f[i+1][j]
			}
			if j >= m || s[j-i] == t[j] {
				f[i][j] = (f[i][j] + f[i][j-1]) % mod
			}
		}
	}
	ans := 0
	for _, v := range f[0][m-1:] {
		ans = (ans + v) % mod
	}
	Fprint(out, ans)
}

//func main() { CF1336C(os.Stdin, os.Stdout) }
