package copypasta

import (
	"math"
	"math/bits"
)

// 注：由于用的是指针写法，必要时禁止 GC，能加速不少
// func init() { debug.SetGCPercent(-1) }

// 异或字典树
// 一棵（所有叶节点深度都相同的）二叉树
// 模板题 LC421 https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/
// LC1707 https://leetcode.cn/problems/maximum-xor-with-an-element-from-array/
// LC1803 https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/
// LC2479 利用先序遍历的特点 https://leetcode.cn/problems/maximum-xor-of-two-non-overlapping-subtrees/
// https://codeforces.com/problemset/problem/706/D
// 数组前缀异或数组后缀的最大值（前后缀不重叠）https://codeforces.com/problemset/problem/282/E
// https://codeforces.com/contest/1446/problem/C
// todo https://codeforces.com/problemset/problem/1055/F
//  转换 https://codeforces.com/contest/1720/problem/D2
//  异或和 ≥k 的最短区间 https://acm.hdu.edu.cn/showproblem.php?pid=6955
type trie01Node struct {
	son [2]*trie01Node
	cnt int // 子树叶子数
	min int // 子树最小值
}

type trie01 struct{ root *trie01Node }

func newTrie01() *trie01 { return &trie01{&trie01Node{min: math.MaxInt32}} }

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
			o.son[b] = &trie01Node{min: math.MaxInt32}
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

// https://codeforces.com/problemset/problem/282/E
// LC1938 https://leetcode-cn.com/problems/maximum-genetic-difference-query/
func (t *trie01) del(v int) *trie01Node {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		o = o.son[v>>i&1]
		o.cnt--
	}
	return o
}

// v 与 trie 上所有数的最大异或值，trie 不能是空的
// 模板题 LC421 https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/
// 离线 LC1707 https://leetcode-cn.com/problems/maximum-xor-with-an-element-from-array/ 注：可以通过记录子树最小值来在线查询
// todo 模板题：树上最长异或路径 https://www.luogu.com.cn/problem/P4551
// todo 好题：区间异或第 k 大 https://www.luogu.com.cn/problem/P5283
// EXTRA: minXor: 若要求 a[i] 与数组 a 中元素的最小异或值，可以先把 a[i] 从 trie01 中删掉，然后搜索一遍即可，最后把 a[i] 重新插入
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

// 上面也可以用哈希表做
// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/discuss/91049/Java-O(n)-solution-using-bit-manipulation-and-HashMap/95535
func findMaximumXOR(a []int) (ans int) {
	mask := 0
	for i := trieBitLen - 1; i >= 0; i-- {
		mask |= 1 << i
		try := ans | 1<<i
		leftPart := map[int]bool{a[0] & mask: true}
		for _, v := range a[1:] {
			if leftPart[v&mask^try] { // 对每个 v' = v&mask，判断是否有 w' 满足 v' ^ w' = try
				ans = try
				break
			}
			leftPart[v&mask] = true
		}
	}
	return
}

// v 与 trie 上所有不超过 limit 的数的最大异或值
// 不存在时返回 -1
// https://codeforces.com/problemset/problem/979/D
// LC1707 https://leetcode-cn.com/problems/maximum-xor-with-an-element-from-array/
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

// 求与 v 异或值不超过 limit 的元素个数
// 核心原理是，当 limit+1 的某一位是 1 的时候，若该位异或值取 0，则后面的位是可以取任意数字的
// 如果在 limit 上而不是 limit+1 上讨论，就要单独处理走到叶子的情况了（恰好等于 limit）
// LC1803 https://leetcode-cn.com/problems/count-pairs-with-xor-in-a-range/
// 补集 https://codeforces.com/problemset/problem/665/E
// https://codeforces.com/problemset/problem/817/E
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
func countLimitXOR(a []int, limit int) int {
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	ans := 0
	for ; limit > 0; limit >>= 1 {
		tmp := cnt
		cnt = map[int]int{}
		for v, c := range tmp {
			cnt[v>>1] += c
			if limit&1 > 0 {
				ans += c * tmp[(limit-1)^v]
			}
		}
	}
	return ans >> 1
}

// v 与 trie 上所有数异或不超过 limit 的最大异或值
// 不存在时返回 -1
// 原理同 countLimitXOR
func (t *trie01) maxXorWithLimitXor(v, limit int) (ans int) {
	limit++ // 改成 <
	lastO, lastI, lastAns := (*trie01Node)(nil), -2, 0
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

// 持久化
// 注意为了拷贝一份 trie01Node，这里的接收器不是指针
// roots := make([]*trie01Node, n+1)
// roots[0] = trie01Node{}.put(0, trieBitLen-1)
// roots[i+1] = roots[i].put(v, trieBitLen-1)
// 模板题（最大异或和） https://www.luogu.com.cn/problem/P4735 https://www.acwing.com/problem/content/258/
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
// n*(k-logn) + 2^(logn+1) - 1, 这里 logn = int(log_2(n))
// 实际使用的时候，可以简单地用 n*(k+2-logn) 代替
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
