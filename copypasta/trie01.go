package copypasta

import (
	"math"
	"math/bits"
	"runtime/debug"
)

// 异或字典树
// 一棵（所有叶节点深度都相同的）二叉树
// 模板题 LC421 https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/
// LC1707 https://leetcode.cn/problems/maximum-xor-with-an-element-from-array/
// LC1803 https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/
// LC2479 利用先序遍历的特点 https://leetcode.cn/problems/maximum-xor-of-two-non-overlapping-subtrees/
// https://codeforces.com/problemset/problem/1847/C 1400
// https://codeforces.com/problemset/problem/706/D 1800
// https://codeforces.com/problemset/problem/923/C 1800 字典序最小
// https://codeforces.com/problemset/problem/2093/G 1900
// https://codeforces.com/problemset/problem/817/E 2000
// https://codeforces.com/problemset/problem/842/D 2000
// https://codeforces.com/problemset/problem/665/E 2100
// - LC3632 https://leetcode.cn/problems/subarrays-with-xor-at-least-k/ （会员题）
// https://codeforces.com/problemset/problem/1446/C 2100
// https://codeforces.com/problemset/problem/282/E 2200 数组前缀异或数组后缀的最大值（前后缀不重叠，但这要求可以无视）
// https://codeforces.com/problemset/problem/979/D 2200
// https://codeforces.com/problemset/problem/888/G 2300
// https://codeforces.com/problemset/problem/1720/D2 2400 转换 
// https://codeforces.com/problemset/problem/1777/F 2400 启发式合并
// https://codeforces.com/problemset/problem/1983/F 2500
// https://codeforces.com/problemset/problem/241/B 2700 求个数以及和
// https://codeforces.com/problemset/problem/1849/F 2700
// https://codeforces.com/problemset/problem/1055/F 2900
// https://codeforces.com/problemset/problem/1616/H 3000
// https://www.luogu.com.cn/problem/P10218
// https://acm.hdu.edu.cn/showproblem.php?pid=6955 异或和 ≥k 的最短区间

// 指针写法，关闭 GC 可以得到明显加速
// 如果仍然超时，可以改成纯数组写法
func init() { debug.SetGCPercent(-1) }

type trie01Node struct {
	son [2]*trie01Node
	cnt int // 子树叶子数（元素个数）
	min int // 子树最小值
}

type trie01 struct{ root *trie01Node }

func newTrie01() *trie01 { return &trie01{&trie01Node{min: math.MaxInt}} }

const trieBitLen = 31 // 30 for 1e9, 63 for int64, or bits.Len(MAX_VAL)

func (trie01) bin(v int) []byte {
	s := make([]byte, trieBitLen)
	for i := range s {
		s[i] = byte(v >> (trieBitLen - 1 - i) & 1)
	}
	return s
}

func (t *trie01) put(v int) *trie01Node {
	o := t.root
	if v < o.min {
		o.min = v
	}
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &trie01Node{min: math.MaxInt}
		}
		o = o.son[b]
		o.cnt++
		if v < o.min {
			o.min = v
		}
	}
	//o.val = v
	return o
}

// https://codeforces.com/problemset/problem/282/E 2200
// LC1938 https://leetcode.cn/problems/maximum-genetic-difference-query/
func (t *trie01) del(v int) *trie01Node {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		o = o.son[v>>i&1]
		o.cnt--
	}
	return o
}

// 返回 v 与 trie 上所有数的最小异或值，trie 不能是空的
// 注：若要求 a[i] 与数组 a 中元素的最小异或值，可以先把 a[i] 从 trie01 中删掉，算好后再把 a[i] 重新插入
// https://codeforces.com/problemset/problem/888/G 2300
func (t *trie01) minXor(v int) (ans int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 完全图，边权为 a[v]^a[w]，求 MST
// Boruvka 算法，分治连边
// O(nlognlogU)
// https://codeforces.com/problemset/problem/888/G 2300
func xorMST(a []int) (ans int) {
	var f func([]int, int)
	f = func(a []int, p int) {
		if a == nil || p < 0 {
			return
		}
		b := [2][]int{}
		for _, v := range a {
			k := v >> p & 1
			b[k] = append(b[k], v)
		}
		if b[0] != nil && b[1] != nil {
			// b[0] 与 b[1] 之间一条边权最小的边
			t := newTrie01()
			for _, v := range b[0] {
				t.put(v)
			}
			minXor := math.MaxInt
			for _, v := range b[1] {
				minXor = min(minXor, t.minXor(v))
			}
			ans += minXor
		}
		f(b[0], p-1)
		f(b[1], p-1)
	}
	f(a, trieBitLen-1)
	return
}

// 返回 v 与 trie 上所有数的最大异或值，trie 不能是空的
// 模板题 LC421 https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/
// 加约束 LC2935 https://leetcode.cn/problems/maximum-strong-pair-xor-ii/
// 离线 LC1707 https://leetcode.cn/problems/maximum-xor-with-an-element-from-array/ 注：可以通过记录子树最小值来在线查询
// https://codeforces.com/problemset/problem/1847/C 1400
// https://codeforces.com/problemset/problem/2093/G 1900
// todo 模板题：树上最长异或路径 https://www.luogu.com.cn/problem/P4551
func (t *trie01) maxXor(v int) (ans int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 返回 v 与 trie 上所有数的第 k 大异或值
// k 从 1 开始
// 如果 k 超过 trie 中元素个数，返回 0
// https://www.luogu.com.cn/problem/P5283 [十二省联考 2019] 异或粽子（堆）
// https://atcoder.jp/contests/abc252/tasks/abc252_h
func (t *trie01) maxXorKth(v, k int) (ans int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil {
			if k <= o.son[b^1].cnt {
				ans |= 1 << i
				b ^= 1
			} else {
				k -= o.son[b^1].cnt
			}
		}
		o = o.son[b]
	}
	return
}

// 动态维护任意两数异或的最小值 https://atcoder.jp/contests/abc308/tasks/abc308_g https://ac.nowcoder.com/acm/contest/53485/F
// 需要用平衡树

// maxXor 也可以用哈希表做
// https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/solution/tu-jie-jian-ji-gao-xiao-yi-tu-miao-dong-1427d/
// O(1) space 做法 https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/solution/lei-si-kuai-su-pai-xu-de-fen-zhi-suan-fa-9eqx/
func findMaximumXOR(a []int) (ans int) {
	highBit := trieBitLen - 1
	//highBit := bits.Len(uint(slices.Max(a))) - 1
	seen := map[int]bool{}
	for i := highBit; i >= 0; i-- {
		clear(seen)
		mask := ans | 1<<i
		for _, v := range a {
			v &= mask
			if seen[v^mask] { // 对每个 v' = v&mask，判断是否有 w' 满足 v' ^ w' = mask
				ans = mask
				break
			}
			seen[v] = true
		}
	}
	return
}

// v 与 trie 上所有不超过 limit 的数的最大异或值
// 不存在时返回 -1
// https://codeforces.com/problemset/problem/979/D 2200
// LC1707 https://leetcode.cn/problems/maximum-xor-with-an-element-from-array/
func (t *trie01) maxXorWithLimitVal(v, limit int) (ans int) {
	o := t.root
	if o.min > limit {
		return -1
	}
	// 由于上面的判断保证了必然存在一个值，后面是不需要判断 o 是否为空的
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil && o.son[b^1].min <= limit {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 求与 v 异或值 <= limit 的元素个数（遍历 v=a[i]，累加，可以得到 limit 的 rank）
// 核心原理是，当 limit+1 的某一位是 1 的时候，若该位异或值取 0，则后面的位是可以取任意数字的
// 为什么要这样写呢？如果在 limit 而不是 limit+1 上讨论，就要单独处理走到叶子的情况了（恰好等于 limit）
// 如果求的是 >= limit，那就用 trie 中元素中个数，减去 < limit 的个数
// LC1803 https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/
// https://codeforces.com/problemset/problem/817/E 2000
// https://codeforces.com/problemset/problem/665/E 2100
// - LC3632 https://leetcode.cn/problems/subarrays-with-xor-at-least-k/ （会员题）
// https://codeforces.com/problemset/problem/1983/F 2500
// https://codeforces.com/problemset/problem/241/B 2700 求个数以及和
// 另见 maxXorKth
func (t *trie01) countLimitXOR(v, limit int) (cnt int) {
	limit++ // 改成 limit+1（求与 v 异或值小于 limit 的元素个数）
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if limit>>i&1 > 0 {
			if o.son[b] != nil {
				cnt += o.son[b].cnt
			}
			b ^= 1
		}
		if o.son[b] == nil {
			return
		}
		o = o.son[b]
	}
	return
}

// 上面也可以用哈希表做
// 图解 https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/solutions/2045560/bu-hui-zi-dian-shu-zhi-yong-ha-xi-biao-y-p2pu/
func countLimitXOR(a []int, limit int) int {
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}

	ans := 0
	for limit++; limit > 0; limit /= 2 {
		tmp := cnt
		cnt = map[int]int{}
		for v, c := range tmp {
			if limit&1 > 0 {
				ans += c * tmp[(limit-1)^v]
			}
			cnt[v>>1] += c
		}
	}
	return ans / 2
}

// 写法二
// 跳过 limit+1 中的 0
func countLimitXOR2(a []int, limit int) int {
	limit++
	p := bits.TrailingZeros(uint(limit))
	cnt := map[int]int{}
	for _, x := range a {
		cnt[x>>p]++
	}
	limit >>= p

	ans := 0
	for limit > 0 {
		limit--
		p := bits.TrailingZeros(uint(limit))
		nxt := map[int]int{}
		for x, c := range cnt {
			ans += c * cnt[x^limit]
			nxt[x>>p] += c
		}
		cnt = nxt
		limit >>= p
	}
	return ans / 2
}

// v 与 trie 上所有数异或不超过 limit 的最大异或值
// 不存在时返回 -1
// 原理同 countLimitXOR
func (t *trie01) maxXorWithLimitXor(v, limit int) (ans int) {
	limit++ // 改成 <
	var lastO *trie01Node
	lastI, lastAns := -2, 0
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if limit>>i&1 > 0 {
			if o.son[b] != nil {
				lastO, lastI, lastAns = o.son[b], i-1, ans
			}
			if o.son[b^1] != nil {
				ans |= 1 << i
			}
			b ^= 1
		}
		if o.son[b] == nil {
			break
		}
		o = o.son[b]
	}

	if lastI < -1 {
		return -1
	}

	ans = lastAns
	o = lastO
	for i := lastI; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil {
			ans |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 计算 trie 中所有元素与 xor 异或后的 mex
// 如果只是计算 trie 中所有元素的 mex，传入 xor = 0
// ！需要保证 trie 中没有重复元素
// https://codeforces.com/problemset/problem/842/D 2000
func (t *trie01) mex(xor int) (mex int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0 && o != nil; i-- {
		b := xor >> i & 1
		// cnt 表示子树叶子数，即元素个数
		if o.son[b] != nil && o.son[b].cnt == 1<<i {
			mex |= 1 << i
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 持久化
// EXTRA: 可持久化异或字典树
// 注意为了拷贝一份 trie01Node，这里的接收器不是指针
// roots := make([]*trie01Node, n+1)
// roots[0] = trie01Node{}.put(0, trieBitLen-1)
// roots[i+1] = roots[i].put(v, trieBitLen-1)
// 模板题（最大异或和） https://www.luogu.com.cn/problem/P4735
// todo https://www.luogu.com.cn/problem/P4592
// todo https://codeforces.com/problemset/problem/1777/F
func (o trie01Node) put(v, k int) *trie01Node {
	if k < 0 {
		return &o
	}
	b := v >> k & 1
	if o.son[b] == nil {
		o.son[b] = &trie01Node{}
	}
	o.son[b] = o.son[b].put(v, k-1)
	//o.maintain()
	return &o
}

// n 个 [0, 2^k) 范围内的数构成的 0-1 trie 至多可以有多少个节点？
// n*(k-logn) + 2^(logn+1) - 1, 这里 logn = int(log_2(n)) = bit.Len(n)-1
// 实际使用的时候，可以简单地用 n*(k+3-bit.Len(n)) 代替
// O(max(n*(logU-logn),n))
// 构造方法：先用不超过 n 的最大的 2 的幂次个数来构建一个完全二叉树，然后把剩余的数放入二叉树的下一层
// 传入 n 和数据范围上限 maxV
// 返回 n 个数，每个数的范围在 [0, maxV] 中
// 当 maxV = 2^30-1 时，各个 n 下的 0-1 trie 节点数
//   n   节点数
// 1e5 1531071
// 2e5 2862143
// 3e5 4124287
// 4e5 5324287
// 5e5 6524287
// 6e5 7648575
// 7e5 8748575
// 8e5 9848575
// 9e5 10948575
// 1e6 12048575
// 当 maxV = 1e9 时，各个 n 下的 0-1 trie 节点数
//   n   节点数
// 1e5 1522076
// 2e5 2844147
// 3e5 4088288
// 4e5 5288288
// 5e5 6511723
// 6e5 7576570
// 7e5 8676570
// 8e5 9776570
// 9e5 10876570
// 1e6 12023441
func generateMaxNodes01TrieData(n, maxV int) []int {
	shift := bits.Len(uint(maxV)) - bits.Len(uint(n)) + 1
	a := make([]int, 0, n)
	// 构建一棵上半部分为完全二叉树，下半部分为一串 0...0 的 01-trie
	for i := 0; i<<shift <= maxV; i++ {
		v := i << shift
		a = append(a, v)
	}
	// 填充上半部分的下一层，由于下半部分的开头是 0，这里要用一个奇数 shift
	for i := 0; len(a) < n; i++ {
		v := (i<<1 | 1) << (shift - 1)
		a = append(a, v)
	}
	return a
}
