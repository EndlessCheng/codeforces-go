package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF883E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var m, ans int
	var s, t string
	Fscan(in, &m, &s, &m)
	has := ['z' + 1]bool{}
	ps := []int{}
	for i, b := range s {
		if b == '*' {
			ps = append(ps, i)
		} else {
			has[b] = true
		}
	}
	a := []string{}
o:
	for ; m > 0; m-- {
		Fscan(in, &t)
		for j, b := range t {
			if s[j] == '*' {
				if has[b] {
					continue o
				}
			} else if t[j] != s[j] {
				continue o
			}
		}
		a = append(a, t)
	}

o2:
	for b := byte('a'); b <= 'z'; b++ {
		if has[b] {
			continue
		}
	o3:
		for _, s := range a {
			for _, p := range ps {
				if s[p] == b {
					continue o3
				}
			}
			continue o2
		}
		ans++
	}
	Fprint(out, ans)
}

//func main() { CF883E(os.Stdin, os.Stdout) }
