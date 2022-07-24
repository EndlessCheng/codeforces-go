package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

// https://space.bilibili.com/206214/dynamic
type pair struct {
	r int
	s string
}

type FoodRatings struct {
	fs map[string]pair
	cs map[string]*redblacktree.Tree
}

func Constructor(foods, cuisines []string, ratings []int) FoodRatings {
	fs := map[string]pair{}
	cs := map[string]*redblacktree.Tree{}
	for i, f := range foods {
		r, c := ratings[i], cuisines[i]
		fs[f] = pair{r, c}
		if cs[c] == nil {
			cs[c] = redblacktree.NewWith(func(x, y interface{}) int {
				a, b := x.(pair), y.(pair)
				if a.r != b.r {
					return utils.IntComparator(b.r, a.r)
				}
				return utils.StringComparator(a.s, b.s)
			})
		}
		cs[c].Put(pair{r, f}, nil)
	}
	return FoodRatings{fs, cs}
}

func (r FoodRatings) ChangeRating(food string, newRating int) {
	p := r.fs[food]
	t := r.cs[p.s]
	t.Remove(pair{p.r, food}) // 移除旧数据
	t.Put(pair{newRating, food}, nil) // 添加新数据
	p.r = newRating
	r.fs[food] = p
}

func (r FoodRatings) HighestRated(cuisine string) string {
	return r.cs[cuisine].Left().Key.(pair).s
}
