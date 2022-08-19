package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1066E(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, m int
	var s, t string
	Fscan(bufio.NewReader(in), &n, &m, &s, &t)
	ans := int64(0)
	c1 := strings.Count(t, "1")
	for i, p2 := 0, 1; i < n && c1 > 0; i++ {
		ans = (ans + int64(int(s[n-1-i]&1)*c1)*int64(p2)) % mod
		c1 -= int(t[m-1-i] & 1)
		p2 = p2 * 2 % mod
	}
	Fprint(out, ans)
}

//func main() { CF1066E(os.Stdin, os.Stdout) }
