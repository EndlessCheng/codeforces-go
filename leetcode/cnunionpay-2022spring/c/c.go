package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxInvestment(product []int, limit int) (ans int) {
	price := sort.Search(1e7, func(price int) bool {
		cnt := 0
		for _, p := range product {
			if p > price {
				cnt += p - price
			}
		}
		return cnt <= limit
	}) + 1 // 这里算的是最后一次交易价格 +1 后的值

	for _, p := range product {
		if p >= price {
			cnt := p - price + 1
			ans = (ans + (price+p)*cnt/2) % (1e9 + 7)
			limit -= cnt
		}
	}
	if limit > len(product) {
		limit = len(product)
	}
	return (ans + limit*(price-1)) % (1e9 + 7)
}
