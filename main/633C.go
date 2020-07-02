package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF633C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type node struct {
		son [26]*node
		s   []byte
	}

	var n, m int
	var s, t []byte
	root := &node{}
	for Fscan(in, &n, &s, &m); m > 0; m-- {
		Fscan(in, &t)
		o := root
		for i := len(t) - 1; i >= 0; i-- {
			c := t[i]
			if c < 'a' {
				c += 32
			}
			c -= 'a'
			if o.son[c] == nil {
				o.son[c] = &node{}
			}
			o = o.son[c]
		}
		o.s = t
	}

	ans := [][]byte{}
	dp := make([]int8, n)
	var f func(int) int8
	f = func(p int) (res int8) {
		if p == n {
			return 1
		}
		dv := &dp[p]
		if *dv != 0 {
			return *dv
		}
		defer func() { *dv = res }()
		o := root
		for i := p; i < len(s); i++ {
			o = o.son[s[i]-'a']
			if o == nil {
				return -1
			}
			if o.s != nil && f(i+1) == 1 {
				ans = append(ans, o.s)
				return 1
			}
		}
		return -1
	}
	f(0)
	for i := len(ans) - 1; i >= 0; i-- {
		Fprint(out, string(ans[i]), " ")
	}
}

//func main() { CF633C(os.Stdin, os.Stdout) }
