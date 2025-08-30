package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func minimumTeachings1(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	learned := make([][]bool, m)
	for i, list := range languages {
		learned[i] = make([]bool, n+1)
		for _, x := range list {
			learned[i][x] = true
		}
	}

	todoList := [][2]int{}
next:
	for _, f := range friendships {
		u, v := f[0]-1, f[1]-1 // 减一，下标从 0 开始
		for _, x := range languages[u] {
			if learned[v][x] { // 两人可以相互沟通，无需学习语言
				continue next
			}
		}
		todoList = append(todoList, [2]int{u, v})
	}

	ans := m
	for k := 1; k <= n; k++ { // 枚举需要教的语言 k
		set := map[int]struct{}{}
		for _, f := range todoList {
			u, v := f[0], f[1]
			if !learned[u][k] { // u 需要学习语言 k
				set[u] = struct{}{}
			}
			if !learned[v][k] { // v 需要学习语言 k
				set[v] = struct{}{}
			}
		}
		ans = min(ans, len(set)) // len(set) 是需要学习语言 k 的人数
	}
	return ans
}

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	learned := make([][]bool, m)
	for i, list := range languages {
		learned[i] = make([]bool, n+1)
		for _, x := range list {
			learned[i][x] = true
		}
	}

	vis := make([]bool, m)
	cnt := make([]int, n+1)
	total := 0
	add := func(u int) {
		if vis[u] {
			return
		}
		total++
		vis[u] = true // 避免重复统计
		for _, x := range languages[u] {
			cnt[x]++
		}
	}

next:
	for _, f := range friendships {
		u, v := f[0]-1, f[1]-1 // 减一，下标从 0 开始
		for _, x := range languages[u] {
			if learned[v][x] { // 两人可以相互沟通，无需学习语言
				continue next
			}
		}
		add(u)
		add(v)
	}

	return total - slices.Max(cnt)
}
