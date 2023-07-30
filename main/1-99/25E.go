package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF25E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	perm3 := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	a := make([]string, 3)
	for i := range a {
		Fscan(in, &a[i])
	}

	getMatch := func(s string) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	kmpSearch := func(text, pattern string) bool {
		match := getMatch(pattern)
		lenP := len(pattern)
		c := 0
		for _, v := range text {
			for c > 0 && pattern[c] != byte(v) {
				c = match[c-1]
			}
			if pattern[c] == byte(v) {
				c++
			}
			if c == lenP {
				return true
			}
		}
		return false
	}
	merge := func(s, t string) string {
		if kmpSearch(s, t) {
			return s
		}
		if kmpSearch(t, s) {
			return t
		}
		match := getMatch(t + "#" + s)
		return s + t[match[len(match)-1]:]
	}

	ans := int(1e9)
	for _, p := range perm3 {
		s := merge(merge(a[p[0]], a[p[1]]), a[p[2]])
		if len(s) < ans {
			ans = len(s)
		}
	}
	Fprint(out, ans)
}

//func main() { CF25E(os.Stdin, os.Stdout) }
