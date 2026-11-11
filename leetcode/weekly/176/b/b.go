package main

type ProductOfNumbers []int

func Constructor() ProductOfNumbers {
	return []int{1}
}

func (p *ProductOfNumbers) Add(num int) {
	if num != 0 {
		*p = append(*p, (*p)[len(*p)-1]*num)
	} else {
		*p = []int{1}
	}
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	n := len(*p)
	if n <= k {
		return 0
	}
	return (*p)[n-1] / (*p)[n-1-k]
}
