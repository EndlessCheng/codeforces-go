package main

import "slices"

// https://space.bilibili.com/206214
func wordSquares(words []string) (ans [][]string) {
	slices.Sort(words) // 保证答案有序

	path := [4]int{}
	onPath := make([]bool, len(words))

	var dfs func(int)
	dfs = func(i int) {
		if i == 4 {
			top := words[path[0]]
			left := words[path[1]]
			right := words[path[2]]
			bottom := words[path[3]]
			if top[0] == left[0] && top[3] == right[0] && bottom[0] == left[3] && bottom[3] == right[3] {
				ans = append(ans, []string{top, left, right, bottom})
			}
			return
		}

		for j, on := range onPath {
			if !on {
				path[i] = j      // 从没有选的下标中选一个
				onPath[j] = true // 已选上
				dfs(i + 1)
				onPath[j] = false // 恢复现场
			}
		}
	}

	dfs(0)
	return
}
