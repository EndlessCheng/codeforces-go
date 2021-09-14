package main

import "runtime/debug"

// github.com/EndlessCheng/codeforces-go
func init() { debug.SetGCPercent(-1) }

type StreamChecker struct{}

type node struct {
	son  [26]*node
	fail *node
	end  bool
}

var root, cur *node

func Constructor(words []string) (s StreamChecker) {
	root = &node{}
	for _, w := range words {
		o := root
		for i := range w {
			b := w[i] - 'a'
			if o.son[b] == nil {
				o.son[b] = &node{}
			}
			o = o.son[b]
		}
		o.end = true
	}
	q := []*node{}
	for _, son := range root.son {
		if son != nil {
			son.fail = root
			q = append(q, son)
		}
	}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		if o.fail == nil {
			o.fail = root
		}
		for i, son := range o.son {
			if son != nil {
				son.fail = o.fail.son[i]
				q = append(q, son)
			} else {
				o.son[i] = o.fail.son[i]
			}
		}
	}
	cur = root
	return
}

func (StreamChecker) Query(b byte) bool {
	cur = cur.son[b-'a']
	if cur == nil {
		cur = root
		return false
	}
	for f := cur; f != nil; f = f.fail {
		if f.end {
			return true
		}
	}
	return false
}
