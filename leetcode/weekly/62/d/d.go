package main

import "runtime/debug"

func init() { debug.SetGCPercent(-1) }

const sep = string('z' + 1)

type WordFilter struct{}

type node struct {
	son [27]*node
	i   int
}

var root *node

func Constructor(words []string) (t WordFilter) {
	root = &node{}
	for i, s := range words {
		n := len(s)
		s += sep + s
		for j := 0; j <= n; j++ {
			o := root
			for _, c := range s[j:] {
				c -= 'a'
				if o.son[c] == nil {
					o.son[c] = &node{}
				}
				o = o.son[c]
				o.i = i
			}
		}
	}
	return
}

func (WordFilter) F(p, s string) int {
	o := root
	for _, c := range s + sep + p {
		o = o.son[c-'a']
		if o == nil {
			return -1
		}
	}
	return o.i
}
