package main

// github.com/EndlessCheng/codeforces-go
type Cashier struct {
	productToPrice map[int]int
	customers      int // 顾客个数
	n              int
	discount       float64
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	productToPrice := make(map[int]int, len(products))
	for i, p := range products {
		productToPrice[p] = prices[i]
	}
	return Cashier{
		productToPrice: productToPrice,
		n:              n,
		discount:       1 - float64(discount)/100,
	}
}

func (c *Cashier) GetBill(product []int, amount []int) (ans float64) {
	for i, p := range product {
		ans += float64(c.productToPrice[p] * amount[i])
	}
	c.customers++
	if c.customers%c.n == 0 { // 第 n,2n,3n,... 个顾客会得到 discount 的折扣
		ans *= c.discount
	}
	return
}
