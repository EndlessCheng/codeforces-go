package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF615C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	f := func(text, pattern []byte) (mx, end int) {
		lenP := len(pattern)
		match := make([]int, lenP)
		for i, c := 1, 0; i < lenP; i++ {
			v := pattern[i]
			for c > 0 && pattern[c] != v {
				c = match[c-1]
			}
			if pattern[c] == v {
				c++
			}
			match[i] = c
		}
		c := 0
		for i, v := range text {
			for c > 0 && pattern[c] != v {
				c = match[c-1]
			}
			if pattern[c] == v {
				if c++; c > mx {
					mx, end = c, i
				}
			}
			if c == lenP {
				return
			}
		}
		return
	}

	var s, t []byte
	Fscan(in, &s, &t)
	n := len(s)
	rev := append([]byte(nil), s...)
	for i := 0; i < n/2; i++ {
		rev[i], rev[n-1-i] = rev[n-1-i], rev[i]
	}
	ans := [][2]int{}
	for len(t) > 0 {
		c1, end1 := f(s, t)
		if c1 == 0 {
			Fprint(out, -1)
			return
		}
		c2, end2 := f(rev, t)
		if c1 >= c2 {
			ans = append(ans, [2]int{end1 - c1 + 2, end1 + 1})
			t = t[c1:]
		} else {
			ans = append(ans, [2]int{n - end2 + c2 - 1, n - end2})
			t = t[c2:]
		}
	}
	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

//func main() { CF615C(os.Stdin, os.Stdout) }
