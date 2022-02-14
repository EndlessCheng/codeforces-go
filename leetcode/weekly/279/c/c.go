package main

import "bytes"

// 懒标记法

// github.com/EndlessCheng/codeforces-go
type Bitset struct{}

var (
	s    []byte
	flip bool
	cnt1 int
)

func Constructor(size int) (_ Bitset) {
	s = bytes.Repeat([]byte{'0'}, size)
	flip, cnt1 = false, 0
	return
}

func (Bitset) Fix(i int) {
	if s[i] == '1' == flip {
		s[i] ^= 1
		cnt1++
	}
}

func (Bitset) Unfix(i int) {
	if s[i] == '0' == flip {
		s[i] ^= 1
		cnt1--
	}
}

func (Bitset) Flip() {
	flip = !flip
	cnt1 = len(s) - cnt1
}

func (Bitset) All() bool  { return cnt1 == len(s) }
func (Bitset) One() bool  { return cnt1 > 0 }
func (Bitset) Count() int { return cnt1 }

func (Bitset) ToString() string {
	if flip {
		t := make([]byte, len(s))
		for i, ch := range s {
			t[i] = ch ^ 1
		}
		return string(t)
	}
	return string(s)
}
