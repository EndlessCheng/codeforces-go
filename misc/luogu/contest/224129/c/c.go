package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	buf := make([]byte, 4096)
	_i, _n := 0, 0
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	rsn := func(n int) []byte {
		b := rc()
		for ; 'a' > b || b > 'z'; b = rc() {
		}
		s := make([]byte, 0, n)
		s = append(s, b)
		for i := 1; i < n; i++ {
			s = append(s, rc())
		}
		return s
	}
	rd := func() (x int) {
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, m := rd(), rd()
	s := rsn(n)
	l := make([]int, n)
	r := make([]int, n)
	cntN := bytes.Count(s, []byte{'n'})
	cntV := 0

	ans := 0
	for i, b := range s {
		if b == 'v' {
			cntV++
		} else if b == 'n' {
			cntN--
		} else {
			ans += cntV * cntN
			l[i] = cntV
			r[i] = cntN
		}
	}

	for ; m > 0; m-- {
		i := rd()
		if s[i] != s[i-1] && (s[i] == 'a' || s[i-1] == 'a') {
			if s[i-1] == 'a' {
				if s[i] == 'n' {
					ans -= l[i-1]
					l[i] = l[i-1]
					r[i] = r[i-1] - 1
				} else {
					ans += r[i-1]
					r[i] = r[i-1]
					l[i] = l[i-1] + 1
				}
			} else if s[i] == 'a' {
				if s[i-1] == 'v' {
					ans -= r[i]
					r[i-1] = r[i]
					l[i-1] = l[i] - 1
				} else {
					ans += l[i]
					l[i-1] = l[i]
					r[i-1] = r[i] + 1
				}
			}
		}
		s[i-1], s[i] = s[i], s[i-1]
		Fprintln(out, ans)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
