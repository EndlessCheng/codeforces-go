package main

// https://space.bilibili.com/206214
func sumPrefixScores(words []string) []int {
	type node struct {
		son   [26]*node
		ids   []int
		score int
	}
	root := &node{}
	for i, word := range words {
		cur := root
		for _, c := range word {
			c -= 'a'
			if cur.son[c] == nil {
				cur.son[c] = &node{}
			}
			cur = cur.son[c]
			cur.score++ // 更新所有前缀的分数
		}
		cur.ids = append(cur.ids, i)
	}

	ans := make([]int, len(words))
	var dfs func(*node, int)
	dfs = func(node *node, sum int) {
		sum += node.score // 累加分数，即可得到答案
		for _, i := range node.ids {
			ans[i] = sum
		}
		for _, child := range node.son {
			if child != nil {
				dfs(child, sum)
			}
		}
	}
	dfs(root, 0)
	return ans
}
