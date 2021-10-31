package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1594B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7

	var T, n, k int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		ans := int64(0)
		for p := int64(1); k > 0; k >>= 1 {
			ans = (ans + k&1*p) % mod
			p = p * n % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1594B(os.Stdin, os.Stdout) }
