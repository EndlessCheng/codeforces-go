package main

import "math"

// https://space.bilibili.com/206214
func maximumSaleItems1(items [][]int, budget int) (ans int) {
	f := make([]int, budget+1)
	minPrice := math.MaxInt

	for _, p := range items {
		factor, price := p[0], p[1]
		minPrice = min(minPrice, price)

		cnt := 0 // 统计 factor 的倍数（包括 factor）
		for _, q := range items {
			if q[0]%factor == 0 {
				cnt++
			}
		}

		// 视作一个体积为 price，价值为 cnt 的物品
		for j := budget; j >= price; j-- {
			f[j] = max(f[j], f[j-price]+cnt)
		}
	}

	for i, cnt := range f {
		ans = max(ans, cnt+(budget-i)/minPrice)
	}
	return
}

func maximumSaleItems(items [][]int, budget int) (ans int) {
	maxFactor := 0
	minPrice := math.MaxInt
	for _, p := range items {
		maxFactor = max(maxFactor, p[0])
		minPrice = min(minPrice, p[1])
	}

	cntFactor := make([]int, maxFactor+1)
	for _, p := range items {
		cntFactor[p[0]]++
	}
	cntMulti := make([]int, maxFactor+1)
	f := make([]int, budget+1)
	sumPrice := 0

	for _, p := range items {
		factor, price := p[0], p[1]

		if cntMulti[factor] == 0 { // 之前没有计算过
			for j := factor; j <= maxFactor; j += factor {
				cntMulti[factor] += cntFactor[j]
			}
		}
		cnt := cntMulti[factor]

		// 视作一个体积为 price，价值为 cnt 的物品
		// 优化：已遍历的物品的体积和至多为 sumPrice，大于这个值的体积和无法凑出来
		sumPrice = min(sumPrice+price, budget)
		for j := sumPrice; j >= price; j-- {
			f[j] = max(f[j], f[j-price]+cnt)
		}
	}

	for i, cnt := range f {
		ans = max(ans, cnt+(budget-i)/minPrice)
	}
	return
}
