package main

// github.com/EndlessCheng/codeforces-go
const trieBitLen = 30

func bin(v int) []byte {
	s := make([]byte, trieBitLen+1)
	for i := range s {
		s[i] = byte(v >> (trieBitLen - i) & 1)
	}
	return s
}

type node struct {
	son [2]*node
	sz  int
}

func (o *node) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node) maintain() {
	o.sz = o.son[0].size() + o.son[1].size()
}

func (o *node) put(s []byte) {
	c := s[0]
	if o.son[c] == nil {
		o.son[c] = &node{}
	}
	if len(s) == 1 {
		o.son[c].sz++
	} else {
		o.son[c].put(s[1:])
	}
	o.maintain()
}

func (o *node) count(mx, v []byte) (ans int) {
	if o == nil {
		return
	}
	cur := v[0]
	if mx[0] == 0 {
		if len(mx) == 1 {
			return o.son[cur].size()
		}
		return o.son[cur].count(mx[1:], v[1:])
	}
	ans = o.son[cur].size()
	if len(mx) == 1 {
		ans += o.son[cur^1].size()
	} else {
		ans += o.son[cur^1].count(mx[1:], v[1:])
	}
	return
}

func countPairs(a []int, low, high int) (ans int) {
	mxH, mxL := bin(high), bin(low-1)
	t := &node{}
	for _, v := range a {
		s := bin(v)
		ans += t.count(mxH, s) - t.count(mxL, s)
		t.put(s)
	}
	return
}
