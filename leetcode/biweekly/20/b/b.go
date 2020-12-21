package main

// github.com/EndlessCheng/codeforces-go
type Cashier struct{}

var (
	ps            [201]int
	cnt, sep, dis int
)

func Constructor(n, discount int, products, prices []int) (_ Cashier) {
	cnt, sep, dis = 0, n, discount
	for i, product := range products {
		ps[product] = prices[i]
	}
	return
}

func (Cashier) GetBill(product, amount []int) (ans float64) {
	s := 0
	for i, p := range product {
		s += amount[i] * ps[p]
	}
	ans = float64(s)
	cnt++
	if cnt%sep == 0 {
		ans -= float64(dis) * ans / 100
	}
	return
}
