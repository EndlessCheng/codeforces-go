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

	cnt := map[int32]int{}
	for n := r(); n > 0; n-- {
		v := r()
		cnt[v] += int(v)
	}

	ans := 0
	for v, c := range cnt {
		s := c
		if v&1 == 0 {
			for v = v / 2 * 3; v&1 == 0; v = v / 2 * 3 {
				c = cnt[v]
				if c == 0 {
					break
				}
				s += c
			}
			s += cnt[v]
		}
		if s > ans {
			ans = s
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
