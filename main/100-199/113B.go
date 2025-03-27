package main

import (
	. "fmt"
	"io"
	"runtime/debug"
)

// https://github.com/EndlessCheng
func init() { debug.SetGCPercent(-1) }

func cf113B(in io.Reader, out io.Writer) {
	calcPi := func(s []byte) []int {
		pi := make([]int, len(s))
		match := 0
		for i := 1; i < len(pi); i++ {
			v := s[i]
			for match > 0 && s[match] != v {
				match = pi[match-1]
			}
			if s[match] == v {
				match++
			}
			pi[i] = match
		}
		return pi
	}
	kmpSearch := func(text, pattern []byte) (end []int) {
		pi := calcPi(pattern)
		match := 0
		for i := range text {
			v := text[i]
			for match > 0 && pattern[match] != v {
				match = pi[match-1]
			}
			if pattern[match] == v {
				match++
			}
			if match == len(pi) {
				end = append(end, i)
				match = pi[match-1]
			}
		}
		return
	}

	var s, p, q []byte
	Fscan(in, &s, &p, &q)
	sufEnd := make([]bool, len(s))
	for _, i := range kmpSearch(s, q) {
		sufEnd[i] = true
	}

	ans := 0
	type node struct {
		son [26]*node
		end bool
	}
	rt := &node{}
	for _, i := range kmpSearch(s, p) {
		i -= len(p) - 1
		o := rt
		for j, b := range s[i:] {
			b -= 'a'
			if o.son[b] == nil {
				o.son[b] = &node{}
			}
			o = o.son[b]
			if !o.end && sufEnd[i+j] {
				o.end = true
				ans++
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf113B(bufio.NewReader(os.Stdin), os.Stdout) }
