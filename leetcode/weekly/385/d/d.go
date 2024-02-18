package main

// https://space.bilibili.com/206214
func countPrefixSuffixPairs(words []string) (ans int64) {
	type pair struct{ x, y byte }
	type node struct {
		son map[pair]*node
		cnt int
	}
	root := &node{son: map[pair]*node{}}
	for _, s := range words {
		cur := root
		for i := range s {
			p := pair{s[i], s[len(s)-1-i]}
			if cur.son[p] == nil {
				cur.son[p] = &node{son: map[pair]*node{}}
			}
			cur = cur.son[p]
			ans += int64(cur.cnt)
		}
		cur.cnt++
	}
	return
}
