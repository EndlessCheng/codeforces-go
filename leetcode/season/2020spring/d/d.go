package main

func minJump(a []int) (ans int) {
	n := len(a)
	vis := make([]bool, n)
	l := 0 // 类似网络流算法中的「当前弧优化」，往前跳的时候，已经访问过的点无需再访问
	for q := []int{0}; len(q) > 0; {
		ans++
		tmp := q
		q = []int{}
		for _, cur := range tmp {
			vis[cur] = true
			r := cur + a[cur]
			if r >= n {
				return
			}
			if !vis[r] {
				q = append(q, r)
			}
			for ; l < cur; l++ {
				if !vis[l] {
					q = append(q, l)
				}
			}
		}
	}
	return
}
