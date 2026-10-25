package main

import "slices"

// https://space.bilibili.com/206214
func countStaircase(nums []int) (ans int) {
	n := len(nums)
	maxNum := slices.Max(nums)

	pos := make([][]int, maxNum+1)
	idxInPos := make([]int, n) // nums[i] 在 pos[nums[i]] 中的下标
	for i, x := range nums {
		idxInPos[i] = len(pos[x])
		pos[x] = append(pos[x], i)
	}

	firsts := []int{}
	firstPos := make([]int, maxNum+1)
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		firstPos[x] = i

		// 把 firsts 中的 x 移到最前面
		j := slices.Index(firsts, x)
		if j != 0 { // 优化：j = 0 时无需移动
			if j >= 0 {
				firsts = append(firsts[:j], firsts[j+1:]...)
			}
			firsts = slices.Insert(firsts, 0, x)
		}

		m := len(firsts)
		if m*(m+1)/2 > n { // firsts 不需要这么多元素
			firsts = firsts[:len(firsts)-1]
		}

		mx, mn := 0, n
		r := i // 子数组右开端点
		for cnt, x := range firsts {
			r += cnt + 1
			// [i,r) 必须包含恰好 cnt+1 个 x
			ps := pos[x]
			k := idxInPos[firstPos[x]] + cnt
			if k >= len(ps) { // [i,n) 中的 x 的个数小于 cnt+1
				break
			}

			// 更新 mx 和 mn 后，必须满足 mx < r <= mn
			if k < len(ps)-1 {
				mn = min(mn, ps[k+1])
			}
			if r > mn { // 继续遍历，r 更大，mn 更小，无法找到合法子数组
				break
			}

			mx = max(mx, ps[k])
			if mx < r {
				ans++
			}
		}
	}
	return ans
}
