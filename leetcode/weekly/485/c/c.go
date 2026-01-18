package main

import "container/heap"

// https://space.bilibili.com/206214
type pair struct{ userId, itemId int }
type AuctionSystem struct {
	itemH  map[int]*hp  // itemId -> [(bidAmount, userId)]
	amount map[pair]int // (userId, itemId) -> bidAmount
}

func Constructor() AuctionSystem {
	return AuctionSystem{map[int]*hp{}, map[pair]int{}}
}

func (a AuctionSystem) AddBid(userId, itemId, bidAmount int) {
	a.amount[pair{userId, itemId}] = bidAmount

	if a.itemH[itemId] == nil {
		a.itemH[itemId] = &hp{}
	}
	heap.Push(a.itemH[itemId], hpPair{bidAmount, userId})
}

func (a AuctionSystem) UpdateBid(userId, itemId, newAmount int) {
	a.AddBid(userId, itemId, newAmount)
	// 堆中重复的元素在 GetHighestBidder 中删除（懒更新）
}

func (a AuctionSystem) RemoveBid(userId, itemId int) {
	delete(a.amount, pair{userId, itemId})
	// 堆中元素在 GetHighestBidder 中删除（懒删除）
}

func (a AuctionSystem) GetHighestBidder(itemId int) (ans int) {
	h := a.itemH[itemId]
	if h == nil {
		return -1
	}
	for h.Len() > 0 {
		if (*h)[0].bidAmount == a.amount[pair{(*h)[0].userId, itemId}] {
			return (*h)[0].userId
		}
		// 货不对板，堆顶出价与实际不符
		heap.Pop(h)
	}
	return -1
}

type hpPair struct{ bidAmount, userId int }
type hp []hpPair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.bidAmount > b.bidAmount || a.bidAmount == b.bidAmount && a.userId > b.userId
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(hpPair)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
