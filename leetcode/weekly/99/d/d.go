package main

// github.com/EndlessCheng/codeforces-go
type FreqStack struct{}

var (
	cnt map[int]int
	stk map[int][]int
	mxC int
)

func Constructor() (_ FreqStack) {
	cnt = map[int]int{}
	stk = map[int][]int{}
	mxC = 0
	return
}

func (FreqStack) Push(v int) {
	cnt[v]++
	c := cnt[v]
	if c > mxC {
		mxC = c
	}
	stk[c] = append(stk[c], v)
}

func (FreqStack) Pop() int {
	s := stk[mxC]
	v := s[len(s)-1]
	if cnt[v] == 1 {
		delete(cnt, v)
	} else {
		cnt[v]--
	}
	if len(s) == 1 {
		delete(stk, mxC)
		mxC--
	} else {
		stk[mxC] = s[:len(s)-1]
	}
	return v
}
