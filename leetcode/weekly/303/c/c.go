package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

// https://space.bilibili.com/206214/dynamic
type pair struct {
	r int
	f string
}

type FoodRatings struct {
	fr map[string]int
	fc map[string]string
	ct map[string]*redblacktree.Tree
}

func Constructor(foods, cuisines []string, ratings []int) FoodRatings {
	fr := map[string]int{}
	fc := map[string]string{}
	ct := map[string]*redblacktree.Tree{}
	for i, f := range foods {
		r, c := ratings[i], cuisines[i]
		fr[f] = r
		fc[f] = c
		if ct[c] == nil {
			ct[c] = redblacktree.NewWith(func(x, y interface{}) int {
				a, b := x.(pair), y.(pair)
				if a.r != b.r {
					return utils.IntComparator(b.r, a.r)
				}
				return utils.StringComparator(a.f, b.f)
			})
		}
		ct[c].Put(pair{r, f}, nil)
	}
	return FoodRatings{fr, fc, ct}
}

func (r FoodRatings) ChangeRating(f string, newRating int) {
	t := r.ct[r.fc[f]]
	t.Remove(pair{r.fr[f], f})
	t.Put(pair{newRating, f}, nil)
	r.fr[f] = newRating
}

func (r FoodRatings) HighestRated(cuisine string) string {
	return r.ct[cuisine].Left().Key.(pair).f
}
