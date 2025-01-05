package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1778C(in io.Reader, out io.Writer) {
	var T, n, k int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s, &t)
		u := 0
		for _, b := range s {
			u |= 1 << (b - 'a')
		}
		ans := 0
		for sub, ok := u, true; ok; ok = sub != u {
			if bits.OnesCount(uint(sub)) <= k {
				res, cnt := 0, 0
				for i, b := range s {
					if sub>>(b-'a')&1 > 0 || b == t[i] {
						cnt++
						res += cnt
					} else {
						cnt = 0
					}
				}
				ans = max(ans, res)
			}
			sub = (sub - 1) & u
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1778C(bufio.NewReader(os.Stdin), os.Stdout) }
