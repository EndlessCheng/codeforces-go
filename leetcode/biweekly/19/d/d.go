package main

// github.com/EndlessCheng/codeforces-go
func minJumps(arr []int) (ans int) {
	pos := map[int][]int{} // 元素值 -> [下标]
	for i, x := range arr {
		pos[x] = append(pos[x], i)
	}

	n := len(arr)
	vis := make([]bool, n)
	vis[0] = true
	q := []int{0} // 起点
	for ; ; ans++ {
		tmp := q
		q = nil
		for _, i := range tmp {
			if i == n-1 { // 到达终点
				return
			}

			// 往右
			if !vis[i+1] {
				vis[i+1] = true
				q = append(q, i+1)
			}

			// 往左
			if i > 0 && !vis[i-1] {
				vis[i-1] = true
				q = append(q, i-1)
			}

			// 访问所有元素值为 arr[i] 的点（下标）
			x := arr[i]
			if pos[x] == nil { // 之前访问过
				continue
			}
			for _, j := range pos[x] {
				if !vis[j] {
					vis[j] = true
					q = append(q, j)
				}
			}
			delete(pos, x) // 避免重复访问
		}
	}
}
