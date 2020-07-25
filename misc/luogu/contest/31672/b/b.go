package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	buf := make([]byte, 4096)
	_i := len(buf)
	rc := func() byte {
		if _i == len(buf) {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int32) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int32(b&15)
		}
		return
	}
	const mod = 998244353
	const mx int = 2e5

	cnt := [mx + 1]int{}
	sum := [mx + 2]int{}
	for n := r(); n > 0; n-- {
		cnt[r()]++
	}
	for i := 1; i <= mx; i++ {
		sum[i+1] = sum[i] + cnt[i]
	}
	ans := 0
	for i := 1; i <= mx; i++ {
		c := cnt[i]
		if c <= 1 {
			continue
		}
		if c >= 3 {
			ans += c * (c - 1) * (c - 2) / 6
		}
		r := i << 1
		if r > mx+1 {
			r = mx + 1
		}
		ans += c * (c - 1) / 2 * (sum[r] - c)
	}
	Fprint(out, ans%mod)
}

func main() { run(os.Stdin, os.Stdout) }
