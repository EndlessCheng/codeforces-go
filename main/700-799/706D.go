package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF706D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type node struct {
		ch [2]*node
		c  int
	}

	var q, x int
	var op string
	root := &node{}
	for i, o := 29, root; i >= 0; i-- {
		o.ch[0] = &node{}
		o = o.ch[0]
		o.c++
	}
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &op, &x); op[0] == '+' {
			for j, o := 29, root; j >= 0; j-- {
				b := x >> j & 1
				if o.ch[b] == nil {
					o.ch[b] = &node{}
				}
				o = o.ch[b]
				o.c++
			}
		} else if op[0] == '-' {
			for i, o := 29, root; i >= 0; i-- {
				o = o.ch[x>>i&1]
				o.c--
			}
		} else {
			ans := 0
			for j, o := 29, root; j >= 0; j-- {
				b := x >> j & 1
				if o.ch[b^1] != nil && o.ch[b^1].c > 0 {
					ans |= 1 << j
					b ^= 1
				}
				o = o.ch[b]
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { CF706D(os.Stdin, os.Stdout) }
