package main

// github.com/EndlessCheng/codeforces-go
type activity struct {
	priceLimit, discount, left, userLimit int
	userCnt                               map[int]int
}

type DiscountSystem []*activity

func Constructor() DiscountSystem {
	return make([]*activity, 1001)
}

func (acts DiscountSystem) AddActivity(actId, priceLimit, discount, number, userLimit int) {
	acts[actId] = &activity{priceLimit, discount, number, userLimit, map[int]int{}}
}

func (acts DiscountSystem) RemoveActivity(actId int) {
	acts[actId] = nil
}

func (acts DiscountSystem) Consume(userId, cost int) int {
	maxDiscount := -1
	var best *activity
	for _, a := range acts {
		if a != nil && a.left > 0 && a.discount > maxDiscount && a.priceLimit <= cost && a.userCnt[userId] < a.userLimit {
			maxDiscount, best = a.discount, a
		}
	}
	if best != nil {
		best.left--
		best.userCnt[userId]++
		cost -= best.discount
	}
	return cost
}
