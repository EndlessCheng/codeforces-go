package main

import "github.com/emirpasic/gods/trees/redblacktree"

// github.com/EndlessCheng/codeforces-go
type MyCalendar struct{}

var t *redblacktree.Tree

func Constructor() (_ MyCalendar) {
	t = redblacktree.NewWithIntComparator()
	t.Put(-1, -1) // 哨兵
	return
}

func (MyCalendar) Book(start, end int) bool {
	o, _ := t.Floor(start)
	if o.Value.(int) > start {
		return false
	}
	if it := t.IteratorAt(o); it.Next() && it.Key().(int) < end {
		return false
	}
	t.Put(start, end)
	return true
}
