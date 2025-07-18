package main

import (
	"slices"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
type trieNode struct {
	son     map[string]*trieNode
	name    string // 文件夹名称
	deleted bool   // 删除标记
}

func deleteDuplicateFolder(paths [][]string) (ans [][]string) {
	root := &trieNode{}
	for _, path := range paths {
		// 把 path 插到字典树中，见 208. 实现 Trie
		cur := root
		for _, s := range path {
			if cur.son == nil {
				cur.son = map[string]*trieNode{}
			}
			if cur.son[s] == nil {
				cur.son[s] = &trieNode{}
			}
			cur = cur.son[s]
			cur.name = s
		}
	}

	exprToNode := map[string]*trieNode{} // 子树括号表达式 -> 子树根节点
	var genExpr func(*trieNode) string
	genExpr = func(node *trieNode) string {
		if node.son == nil { // 叶子
			return node.name // 表达式就是文件夹名
		}

		expr := make([]string, 0, len(node.son)) // 预分配空间
		for _, son := range node.son {
			// 每个子树的表达式外面套一层括号
			expr = append(expr, "("+genExpr(son)+")")
		}
		slices.Sort(expr)

		subTreeExpr := strings.Join(expr, "") // 按字典序拼接所有子树的表达式
		n := exprToNode[subTreeExpr]
		if n != nil { // 哈希表中有 subTreeExpr，说明有重复的文件夹
			n.deleted = true    // 哈希表中记录的节点标记为删除
			node.deleted = true // 当前节点标记为删除
		} else {
			exprToNode[subTreeExpr] = node
		}

		return node.name + subTreeExpr
	}
	for _, son := range root.son {
		genExpr(son)
	}

	// 在字典树上回溯，仅访问未被删除的节点，并将路径记录到答案中
	// 类似 257. 二叉树的所有路径
	path := []string{}
	var dfs func(*trieNode)
	dfs = func(node *trieNode) {
		if node.deleted {
			return
		}
		path = append(path, node.name)
		ans = append(ans, slices.Clone(path))
		for _, son := range node.son {
			dfs(son)
		}
		path = path[:len(path)-1] // 恢复现场
	}
	for _, son := range root.son {
		dfs(son)
	}
	return
}
