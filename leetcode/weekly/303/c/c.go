package main

import (
	"container/heap"
)

// https://space.bilibili.com/206214
type pair struct {
	rating int
	s      string
}

type FoodRatings struct {
	foodMap    map[string]pair
	cuisineMap map[string]*hp
}

func Constructor(foods, cuisines []string, ratings []int) FoodRatings {
	foodMap := map[string]pair{}
	cuisineMap := map[string]*hp{}
	for i, food := range foods {
		rating, cuisine := ratings[i], cuisines[i]
		foodMap[food] = pair{rating, cuisine}
		if cuisineMap[cuisine] == nil {
			cuisineMap[cuisine] = &hp{}
		}
		heap.Push(cuisineMap[cuisine], pair{rating, food})
	}
	return FoodRatings{foodMap, cuisineMap}
}

func (r FoodRatings) ChangeRating(food string, newRating int) {
	p := r.foodMap[food]
	// 直接添加新数据，后面 highestRated 再删除旧的
	heap.Push(r.cuisineMap[p.s], pair{newRating, food})
	p.rating = newRating
	r.foodMap[food] = p
}

func (r FoodRatings) HighestRated(cuisine string) string {
	h := r.cuisineMap[cuisine]
	// 堆顶的食物评分不等于其实际值
	for h.Len() > 0 && (*h)[0].rating != r.foodMap[(*h)[0].s].rating {
		heap.Pop(h)
	}
	return (*h)[0].s
}

type hp []pair
func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.rating > b.rating || a.rating == b.rating && a.s < b.s
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
