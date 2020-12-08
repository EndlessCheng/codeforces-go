package main

// github.com/EndlessCheng/codeforces-go
func finalPrices(prices []int) (ans []int) {
o:
	for i, v := range prices {
		for _, w := range prices[i+1:] {
			if w <= v {
				ans = append(ans, v-w)
				continue o
			}
		}
		ans = append(ans, v)
	}
	return
}
