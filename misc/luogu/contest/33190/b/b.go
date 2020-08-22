package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const mx = 1e6

var cnt [mx + 1]int

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

	n, k := r(), r()
	for ; n > 0; n-- {
		cnt[r()] += int(r())
	}

	ans := -1
	if k == 0 {
		for i := int(mx); i >= 0; i-- {
			if c := cnt[i]; c > 1 {
				if c*i > ans {
					ans = c * i
				}
			}
		}
	} else {
		for i := k; i <= mx; i++ {
			if c := min(cnt[i-k], cnt[i]); c > 0 {
				if s := c * int(2*i-k); s > ans {
					ans = s
				}
			}
		}
	}
	if ans < 0 {
		Fprint(out, "NO")
		return
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
