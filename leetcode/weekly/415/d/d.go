package main

import "math/rand"

// https://space.bilibili.com/206214
func minValidStrings(words []string, target string) (ans int) {
	n := len(target)

	// 多项式字符串哈希（方便计算子串哈希值）
	// 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack（注意 Go1.20 之后的版本，每次随机的数都不一样）
	powBase := make([]int, n+1)  // powBase[i] = base^i
	preHash := make([]int, n+1)  // 前缀哈希值 preHash[i] = hash(target[:i])
	powBase[0] = 1
	for i, b := range target {
		powBase[i+1] = powBase[i] * base % mod
		preHash[i+1] = (preHash[i]*base + int(b)) % mod // 秦九韶算法计算多项式哈希
	}
	// 计算子串 target[l:r] 的哈希值，注意这是左闭右开区间 [l,r)
	// 计算方法类似前缀和
	subHash := func(l, r int) int {
		return ((preHash[r]-preHash[l]*powBase[r-l])%mod + mod) % mod
	}

	maxLen := 0
	for _, w := range words {
		maxLen = max(maxLen, len(w))
	}
	sets := make([]map[int]bool, maxLen)
	for i := range sets {
		sets[i] = map[int]bool{}
	}
	for _, w := range words {
		h := 0
		for j, b := range w {
			h = (h*base + int(b)) % mod
			sets[j][h] = true // 注意 j 从 0 开始
		}
	}

	curR := 0 // 已建造的桥的右端点
	nxtR := 0 // 下一座桥的右端点的最大值
	for i := range target {
		for nxtR < n && nxtR-i < maxLen && sets[nxtR-i][subHash(i, nxtR+1)] {
			nxtR++
		}
		if i == curR { // 到达已建造的桥的右端点
			if i == nxtR { // 无论怎么造桥，都无法从 i 到 i+1
				return -1
			}
			curR = nxtR // 建造下一座桥
			ans++
		}
	}
	return
}

func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0 // z-box 左右边界（闭区间）
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	return z
}

// 桥的概念，见我在 45 或 1326 题下的题解
func jump(maxJumps []int) (ans int) {
	curR := 0 // 已建造的桥的右端点
	nxtR := 0 // 下一座桥的右端点的最大值
	for i, maxJump := range maxJumps {
		nxtR = max(nxtR, i+maxJump)
		if i == curR { // 到达已建造的桥的右端点
			if i == nxtR { // 无论怎么造桥，都无法从 i 到 i+1
				return -1
			}
			curR = nxtR // 造一座桥
			ans++
		}
	}
	return
}

func minValidStrings2(words []string, target string) int {
	maxJumps := make([]int, len(target))
	for _, word := range words {
		z := calcZ(word + "#" + target)
		for i, z := range z[len(word)+1:] {
			maxJumps[i] = max(maxJumps[i], z)
		}
	}
	return jump(maxJumps)
}

// node 表示从根到 node 的字符串，也是某个 words[i] 的前缀
type node struct {
	son  [26]*node
	fail *node // 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
	len  int   // node 代表的字符串的长度，也是 node 在 trie 中的深度
}

type acam struct {
	root *node
}

func (ac *acam) put(s string) {
	cur := ac.root
	for _, b := range s {
		b -= 'a'
		if cur.son[b] == nil {
			cur.son[b] = &node{len: cur.len + 1}
		}
		cur = cur.son[b]
	}
}

func (ac *acam) buildFail() {
	ac.root.fail = ac.root
	q := []*node{}
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
		} else {
			son.fail = ac.root // 第一层的失配指针，都指向根节点 ∅
			q = append(q, son)
		}
	}
	// BFS
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i, son := range cur.son[:] {
			if son == nil {
				// 虚拟子节点 cur.son[i]，和 cur.fail.son[i] 是同一个
				// 方便失配时直接跳到下一个可能匹配的位置（但不一定是某个 words[k] 的最后一个字母）
				cur.son[i] = cur.fail.son[i]
				continue
			}
			son.fail = cur.fail.son[i] // 计算失配位置
			q = append(q, son)
		}
	}
}

func minValidStringsACAM(words []string, target string) int {
	ac := &acam{root: &node{}}
	for _, w := range words {
		ac.put(w)
	}
	ac.buildFail()

	n := len(target)
	f := make([]int, n+1)
	cur := ac.root
	for i, b := range target {
		cur = cur.son[b-'a']
		if cur == ac.root { // 没有任何字符串的前缀与 target[:i+1] 的后缀匹配
			return -1
		}
		f[i+1] = f[i+1-cur.len] + 1
	}
	return f[n]
}
