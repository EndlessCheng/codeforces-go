package main

import (
	"index/suffixarray"
	"math"
	"math/rand"
	"slices"
)

// https://space.bilibili.com/206214
type node struct {
	son  [26]*node
	fail *node // 当 cur.son[i] 不能匹配 target 中的某个字符时，cur.fail.son[i] 即为下一个待匹配节点（等于 root 则表示没有匹配）
	last *node // 后缀链接（suffix link），用来快速跳到一定是某个 words[k] 的最后一个字母的节点（等于 root 则表示没有）
	len  int
	cost int
}

type acam struct {
	root *node
}

func (ac *acam) put(s string, cost int) {
	cur := ac.root
	for _, b := range s {
		b -= 'a'
		if cur.son[b] == nil {
			cur.son[b] = &node{cost: math.MaxInt}
		}
		cur = cur.son[b]
	}
	cur.len = len(s)
	cur.cost = min(cur.cost, cost)
}

func (ac *acam) buildFail() {
	ac.root.fail = ac.root
	ac.root.last = ac.root
	q := []*node{}
	for i, son := range ac.root.son[:] {
		if son == nil {
			ac.root.son[i] = ac.root
		} else {
			son.fail = ac.root // 第一层的失配指针，都指向根节点 ∅
			son.last = ac.root
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
			if son.fail.len > 0 {
				son.last = son.fail
			} else {
				// 沿着 last 往上走，可以直接跳到一定是某个 words[k] 的最后一个字母的节点（如果跳到 root 表示没有匹配）
				son.last = son.fail.last
			}
			q = append(q, son)
		}
	}
}

func minimumCostACAM(target string, words []string, costs []int) int {
	ac := &acam{root: &node{}}
	for i, w := range words {
		ac.put(w, costs[i])
	}
	ac.buildFail()

	n := len(target)
	f := make([]int, n+1)
	cur := ac.root
	for i, b := range target {
		cur = cur.son[b-'a'] // 如果没有匹配相当于移动到 fail 的 son[b-'a']
		i++
		f[i] = math.MaxInt / 2
		if cur.len > 0 { // 匹配到了一个尽可能长的 words[k]
			f[i] = min(f[i], f[i-cur.len]+cur.cost)
		}
		// 还可能匹配其余更短的 words[k]，要在 last 链上找
		for match := cur.last; match != ac.root; match = match.last {
			f[i] = min(f[i], f[i-match.len]+match.cost)
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}

//

func minimumCost(target string, words []string, costs []int) int {
	n := len(target)
	const mod1 = 1_070_777_777
	const mod2 = 1_000_000_007
	base1 := 9e8 - rand.Intn(1e8)
	base2 := 9e8 - rand.Intn(1e8)

	type hPair struct{ h1, h2 int }
	powBase := make([]hPair, n+1)
	preHash := make([]hPair, n+1)
	powBase[0] = hPair{1, 1}
	for i, b := range target {
		powBase[i+1] = hPair{powBase[i].h1 * base1 % mod1, powBase[i].h2 * base2 % mod2}
		preHash[i+1] = hPair{(preHash[i].h1*base1 + int(b)) % mod1, (preHash[i].h2*base2 + int(b)) % mod2}
	}

	// 计算子串 target[l:r] 的哈希值
	// 空串的哈希值为 0
	subHash := func(l, r int) hPair {
		h1 := ((preHash[r].h1-preHash[l].h1*powBase[r-l].h1)%mod1 + mod1) % mod1
		h2 := ((preHash[r].h2-preHash[l].h2*powBase[r-l].h2)%mod2 + mod2) % mod2
		return hPair{h1, h2}
	}

	calcHash := func(t string) (p hPair) {
		for _, b := range t {
			p.h1 = (p.h1*base1 + int(b)) % mod1
			p.h2 = (p.h2*base2 + int(b)) % mod2
		}
		return
	}

	minCost := make([]map[hPair]int, n+1) // [words[i] 的长度][words[i] 的哈希值] -> 最小成本
	lens := map[int]struct{}{}            // 所有 words[i] 的长度集合
	for i, w := range words {
		m := len(w)
		lens[m] = struct{}{}
		h := calcHash(w)
		if minCost[m] == nil {
			minCost[m] = map[hPair]int{}
		}
		if minCost[m][h] == 0 {
			minCost[m][h] = costs[i]
		} else {
			minCost[m][h] = min(minCost[m][h], costs[i])
		}
	}

	// 有 O(√L) 个不同的长度
	sortedLens := make([]int, 0, len(lens))
	for l := range lens {
		sortedLens = append(sortedLens, l)
	}
	slices.Sort(sortedLens)

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, sz := range sortedLens {
			if sz > i {
				break
			}
			if cost, ok := minCost[sz][subHash(i-sz, i)]; ok {
				f[i] = min(f[i], f[i-sz]+cost)
			}
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}

//

func minimumCostHash2(target string, words []string, costs []int) int {
	n := len(target)

	// 多项式字符串哈希（方便计算子串哈希值）
	// 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	const mod = 1_070_777_777
	base := 9e8 - rand.Intn(1e8) // 随机 base，防止 hack（注意 Go1.20 之后的版本，每次随机的数都不一样）
	powBase := make([]int, n+1)  // powBase[i] = base^i
	preHash := make([]int, n+1)  // 前缀哈希值 preHash[i] = hash(s[:i])
	powBase[0] = 1
	for i, b := range target {
		powBase[i+1] = powBase[i] * base % mod
		preHash[i+1] = (preHash[i]*base + int(b)) % mod // 秦九韶算法计算多项式哈希
	}
	// 计算子串 target[l:r] 的哈希值，注意这是左闭右开区间 [l,r) 其中 0<=l<=r<=len(s)
	// 空串的哈希值为 0
	// 计算方法类似前缀和
	subHash := func(l, r int) int {
		return ((preHash[r]-preHash[l]*powBase[r-l])%mod + mod) % mod
	}

	minCost := make([]map[int]int, n+1) // words[i] 的哈希值 -> 最小成本
	lens := map[int]struct{}{}          // 所有 words[i] 的长度集合
	for i, w := range words {
		m := len(w)
		lens[m] = struct{}{}
		// 计算 w 的哈希值
		h := 0
		for _, b := range w {
			h = (h*base + int(b)) % mod
		}
		if minCost[m] == nil {
			minCost[m] = map[int]int{}
		}
		if minCost[m][h] == 0 {
			minCost[m][h] = costs[i]
		} else {
			minCost[m][h] = min(minCost[m][h], costs[i])
		}
	}

	// 有 O(√L) 个不同的长度
	sortedLens := make([]int, 0, len(lens))
	for l := range lens {
		sortedLens = append(sortedLens, l)
	}
	slices.Sort(sortedLens)

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, sz := range sortedLens {
			if sz > i {
				break
			}
			if cost := minCost[sz][subHash(i-sz, i)]; cost > 0 {
				f[i] = min(f[i], f[i-sz]+cost)
			}
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}

func minimumCostSA(target string, words []string, costs []int) int {
	minCost := map[string]uint16{}
	for i, w := range words {
		c := uint16(costs[i])
		if minCost[w] == 0 {
			minCost[w] = c
		} else {
			minCost[w] = min(minCost[w], c)
		}
	}

	n := len(target)
	type pair struct{ l, cost uint16 }
	from := make([][]pair, n+1)
	sa := suffixarray.New([]byte(target))
	for w, c := range minCost {
		for _, l := range sa.Lookup([]byte(w), -1) {
			r := l + len(w)
			from[r] = append(from[r], pair{uint16(l), c})
		}
	}

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2
		for _, p := range from[i] {
			f[i] = min(f[i], f[p.l]+int(p.cost))
		}
	}
	if f[n] == math.MaxInt/2 {
		return -1
	}
	return f[n]
}
