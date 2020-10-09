package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1422C(_r io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	var d []byte
	Fscan(bufio.NewReader(_r), &d)
	var ans, k, p10, n int64 = 0, 0, 1, int64(len(d))
	for i := n - 1; i >= 0; i-- {
		ans += int64(d[i]&15) * (i*(i+1)/2%mod*p10%mod + k)
		k = (k + (n-i)*p10) % mod
		p10 = p10 * 10 % mod
	}
	Fprint(out, ans%mod)
}

//func main() { CF1422C(os.Stdin, os.Stdout) }
