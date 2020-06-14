package main

type TreeAncestor struct {
}

const mx = 17
var pa [][mx]int

func Constructor(n int, parent []int) (t TreeAncestor) {
	pa = make([][mx]int, n)
	for i, p := range parent {
		pa[i][0] = p
	}
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	return
}

func (*TreeAncestor) GetKthAncestor(v int, k int) (ans int) {
	for i := 0; i < mx && v != -1; i++ {
		if k>>i&1 > 0 {
			v = pa[v][i]
		}
	}
	return v
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
