package main

/* 离线+字典树
离线，将查询按照 $\textit{node}_i$ 分组。
根据 $\textit{parents}$ 建树，然后从根开始 DFS，每访问一个节点 $v$，就将其插入 trie 树，然后回答当前节点 $v$ 对应的所有询问，最后在递归结束时将 $v$ 从 trie 中删去。
*/

// github.com/EndlessCheng/codeforces-go
type node struct {
	son [2]*node
	cnt int
}
type trie struct{ root *node }

func (t *trie) put(v int) *node {
	o := t.root
	for i := 17; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &node{}
		}
		o = o.son[b]
		o.cnt++
	}
	return o
}

func (t *trie) del(v int) *node {
	o := t.root
	for i := 17; i >= 0; i-- {
		o = o.son[v>>i&1]
		o.cnt-- // 删除操作只需要减少 cnt 就行，cnt 为 0 就视作删掉了该节点
	}
	return o
}

func (t *trie) maxXor(v int) (ans int) {
	o := t.root
	for i := 17; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil && o.son[b^1].cnt > 0 {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

func maxGeneticDifference(parents []int, queries [][]int) []int {
	n := len(parents)
	// 建树
	g := make([][]int, n)
	var root int
	for v, pa := range parents {
		if pa == -1 {
			root = v
		} else {
			g[pa] = append(g[pa], v)
		}
	}

	// 离线，将查询分组
	type query struct{ val, i int }
	qs := make([][]query, n)
	for i, q := range queries {
		qs[q[0]] = append(qs[q[0]], query{q[1], i})
	}

	ans := make([]int, len(queries))
	t := &trie{&node{}}
	// 遍历整棵树，每访问一个节点就将其插入 trie 树，访问结束时将其从 trie 中删去
	var dfs func(int)
	dfs = func(v int) {
		t.put(v)
		for _, q := range qs[v] {
			ans[q.i] = t.maxXor(q.val)
		}
		for _, w := range g[v] {
			dfs(w)
		}
		t.del(v)
	}
	dfs(root)
	return ans
}
