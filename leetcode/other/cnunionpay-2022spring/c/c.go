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
	})

	for _, p := range product {
		if p > price {
			cnt := p - price
			ans += (p + price + 1) * cnt / 2
			limit -= cnt
		}
	}
	if limit > len(product) {
		limit = len(product)
	}
	return (ans + limit*price) % (1e9 + 7)
}
