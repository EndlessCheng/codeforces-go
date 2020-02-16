package main

type ProductOfNumbers struct {
	a []int
}

func Constructor() (p ProductOfNumbers) {
	return
}

func (p *ProductOfNumbers) Add(v int) {
	p.a = append(p.a, v)
}

func (p *ProductOfNumbers) GetProduct(k int) (ans int) {
	ans = 1
	for i := 0; i < k; i++ {
		ans *= p.a[len(p.a)-1-i]
	}
	return
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
