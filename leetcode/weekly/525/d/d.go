package main

import "slices"

// https://space.bilibili.com/206214
func countStaircase(nums []int) (ans int) {
	n := len(nums)
	//pos := map[int][]int{}
	pos := make([][]int, slices.Max(nums)+1) // slice 比 map 快
	firsts := []int{}

	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		pos[x] = append(pos[x], i) // pos[x][len(pos[x])-k] 是右侧第 k 个 x 的下标

		// 把 firsts 中的 x 移到最前面
		j := slices.Index(firsts, x)
		if j >= 0 {
			firsts = append(firsts[:j], firsts[j+1:]...)
		}
		firsts = slices.Insert(firsts, 0, x)

		m := len(firsts)
		if m*(m+1)/2 > n { // firsts 不需要这么多元素
			firsts = firsts[:len(firsts)-1]
		}

		mx, mn := 0, n
		r := i // 子数组右开端点
		for k, x := range firsts {
			k++
			r += k
			// [i,r) 必须包含恰好 k 个 x
			ps := pos[x]
			sz := len(ps)
			if k > sz { // [i,n) 中的 x 的个数小于 k
				break
			}

			// 更新 mx 和 mn 后，必须满足 mx < r <= mn
			mx = max(mx, ps[sz-k])
			if k < len(ps) {
				mn = min(mn, ps[sz-k-1])
			}
			if mx >= mn || r > mn { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
				break
			}
			if mx < r {
				ans++
			}
		}
	}

	return
}
