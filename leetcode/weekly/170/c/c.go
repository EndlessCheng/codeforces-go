package main

import "sort"

func bfsWithDepth(g [][]int, st int, do func(v, dep int)) {
	visited := make([]bool, len(g))
	visited[st] = true
	type pair struct{ v, dep int }
	queue := []pair{{st, 0}}
	for len(queue) > 0 {
		var p pair
		p, queue = queue[0], queue[1:]
		do(p.v, p.dep)
		for _, w := range g[p.v] {
			if !visited[w] {
				visited[w] = true
				queue = append(queue, pair{w, p.dep + 1})
			}
		}
	}
}

func watchedVideosByFriends(watchedVideos [][]string, g [][]int, st int, level int) []string {
	cntMap := map[string]int{}
	bfsWithDepth(g, st, func(v, dep int) {
		if dep == level {
			for _, video := range watchedVideos[v] {
				cntMap[video]++
			}
		}
	})

	cnt := make([][]string, 200)
	for k, v := range cntMap {
		cnt[v] = append(cnt[v], k)
	}
	ans := []string{}
	for _, c := range cnt {
		if len(c) > 0 {
			sort.Strings(c)
			ans = append(ans, c...)
		}
	}
	return ans
}
