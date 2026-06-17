package main

import (
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func numOfStrings1(patterns []string, word string) (ans int) {
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			ans++
		}
	}
	return
}

//

type node struct {
	son  [26]*node
	fail *node // 当 node.son[i] 失配时，node.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
	last *node // 后缀链接（suffix link），用来快速跳到一定是某个模式串末尾的节点（等于 root 则表示匹配结束）
	cnt  int   // node 是 cnt 个模式串的末尾
}

type acam struct {
	root *node
}

// 把模式串 pattern 插入 AC 自动机（代码和字典树一样）
func (ac *acam) put(pattern string) {
	cur := ac.root
	for _, ch := range pattern {
		ch -= 'a'
		if cur.son[ch] == nil {
			cur.son[ch] = &node{}
		}
		cur = cur.son[ch]
	}
	cur.cnt++
}

// BFS，构建 AC 自动机的 fail 和 last，方便快速查询
func (ac *acam) buildFail() {
	ac.root.fail = ac.root
	ac.root.last = ac.root

	q := []*node{}
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
			continue
		}
		son.fail = ac.root // 第一层的 fail 都指向根节点
		son.last = ac.root
		q = append(q, son)
	}

	// BFS
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i, son := range cur.son[:] {
			if son == nil {
				// 把虚拟子节点 cur.son[i] 设置为 cur.fail.son[i]
				// 方便失配时直接跳到下一个可能匹配的位置（但不一定是某个模式串的末尾）
				cur.son[i] = cur.fail.son[i]
				continue
			}
			son.fail = cur.fail.son[i] // 计算失配位置
			if son.fail.cnt > 0 {
				son.last = son.fail
			} else {
				// 沿着 last 往上走，可以直接跳到一定是某个模式串末尾的节点（如果跳到 root 表示匹配结束）
				son.last = son.fail.last
			}
			q = append(q, son)
		}
	}
}

func numOfStrings(patterns []string, word string) (ans int) {
	ac := &acam{root: &node{}}
	for _, pattern := range patterns {
		ac.put(pattern)
	}
	ac.buildFail()

	cur := ac.root
	for _, ch := range word {
		cur = cur.son[ch-'a'] // 如果没有匹配，相当于移动到 fail 的 son[ch-'a']
		// 可能匹配更短的模式串，要继续在 last 链上找
		for match := cur; match.cnt >= 0; match = match.last {
			ans += match.cnt
			match.cnt = -1 // 避免重复统计
		}
	}
	return
}
