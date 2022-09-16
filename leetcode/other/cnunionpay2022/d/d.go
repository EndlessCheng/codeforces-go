package main

import (
	"container/heap"
)

// https://space.bilibili.com/206214
type Good struct {
	data   hp
	expire hp2
	left   int
}

type VendingMachine struct {
	good     map[string]*Good
	discount map[string]int
}

func Constructor() VendingMachine {
	return VendingMachine{map[string]*Good{}, map[string]int{}}
}

func (v VendingMachine) AddItem(time int, number int, item string, price int, duration int) {
	if v.good[item] == nil {
		v.good[item] = &Good{}
	}
	it := v.good[item]
	t := &tuple{price, time + duration, number}
	heap.Push(&it.data, t)
	heap.Push(&it.expire, t)
	it.left += number
}

func (v VendingMachine) Sell(time int, customer string, item string, number int) int64 {
	it := v.good[item]
	if it == nil {
		return -1
	}

	// 清除过期商品
	for len(it.expire) > 0 && it.expire[0].expire < time {
		t := heap.Pop(&it.expire).(*tuple)
		it.left -= t.left
		t.left = 0
	}

	if it.left < number {
		return -1
	}
	it.left -= number

	// 计算花费
	cost := 0
	for len(it.data) > 0 {
		t := it.data[0]
		if t.left >= number {
			cost += number * t.price
			t.left -= number
			break
		}
		cost += t.left * t.price
		number -= t.left
		t.left = 0
		heap.Pop(&it.data)
	}

	// 计算折扣
	if v.discount[customer] == 0 {
		v.discount[customer] = 100
	}
	ans := (cost*v.discount[customer] + 99) / 100
	if v.discount[customer] > 70 {
		v.discount[customer]--
	}
	return int64(ans)
}

type tuple struct{ price, expire, left int }
type hp []*tuple
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.price < b.price || a.price == b.price && a.expire < b.expire }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(*tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
type hp2 []*tuple
func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].expire < h[j].expire }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(*tuple)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
