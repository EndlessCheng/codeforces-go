package main

type RLEIterator struct {
	a []int
}

func Constructor(a []int) (r RLEIterator) {
	r.a = a
	return
}

func (r *RLEIterator) Next(n int) (ans int) {
	ans = -1
	for len(r.a) > 0 {
		if r.a[0] < n {
			n -= r.a[0]
			r.a = r.a[2:]
		} else {
			ans = r.a[1]
			r.a[0] -= n
			break
		}
	}
	return
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
