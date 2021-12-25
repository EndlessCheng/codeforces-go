package main

// github.com/EndlessCheng/codeforces-go
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) (ans []string) {
	n := len(recipes)
	idx := make(map[string]int, n+len(supplies))
	for i, r := range recipes {
		idx[r] = i + 1
	}
	for _, s := range supplies {
		idx[s] = -1
	}

	g := make([][]int, n)
	deg := make([]int, n)
next:
	for i, in := range ingredients {
		for _, s := range in {
			if idx[s] == 0 { // 没有原材料
				deg[i] = -1
				continue next
			}
		}
		for _, s := range in {
			if j := idx[s]; j > 0 {
				g[j-1] = append(g[j-1], i) // 建图
				deg[i]++
			}
		}
	}

	// 跑拓扑排序
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
			ans = append(ans, recipes[i])
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
				ans = append(ans, recipes[w])
			}
		}
	}
	return
}
