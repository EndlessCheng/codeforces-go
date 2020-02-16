package main

type ProductOfNumbers struct {
	p []int
}

func Constructor() (p ProductOfNumbers) {
	p.p = []int{1}
	return
}

func (p *ProductOfNumbers) Add(v int) {
	if v != 0 {
		p.p = append(p.p, p.p[len(p.p)-1]*v)
	} else {
		p.p = []int{1}
	}
}

func (p *ProductOfNumbers) GetProduct(k int) (ans int) {
	if k < len(p.p) {
		return p.p[len(p.p)-1] / p.p[len(p.p)-1-k]
	}
	return 0
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
