package main

import "github.com/emirpasic/gods/trees/redblacktree"

// github.com/EndlessCheng/codeforces-go
type StockPrice struct {
	*redblacktree.Tree
	prices   map[int]int
	now, cur int
}

func Constructor() StockPrice {
	return StockPrice{redblacktree.NewWithIntComparator(), map[int]int{}, 0, 0}
}

func (s *StockPrice) Update(timestamp, price int) {
	if p := s.prices[timestamp]; p > 0 {
		s.remove(p) // 更新价格前先把旧的删掉
	}
	s.put(price) // 记录价格
	s.prices[timestamp] = price // 记录时间戳对应价格
	if timestamp >= s.now {
		s.now, s.cur = timestamp, price // 更新最新时间及价格
	}
}

func (s *StockPrice) Current() int { return s.cur }
func (s *StockPrice) Maximum() int { return s.Right().Key.(int) }
func (s *StockPrice) Minimum() int { return s.Left().Key.(int) }

func (s *StockPrice) put(v int) {
	c := 0
	if cnt, has := s.Get(v); has {
		c = cnt.(int)
	}
	s.Put(v, c+1)
}

func (s *StockPrice) remove(v int) {
	if cnt, _ := s.Get(v); cnt.(int) > 1 {
		s.Put(v, cnt.(int)-1)
	} else {
		s.Remove(v)
	}
}
