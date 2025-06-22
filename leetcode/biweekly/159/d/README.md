## 思路

> **离线**：按照自己定义的某种顺序回答询问，而不是按照输入顺序 $\textit{queries}[0],\textit{queries}[1],\cdots$ 回答询问。

把询问按照 $u$ 分组，这样我们可以一边 DFS 这棵树，一边回答询问。递归完子树 $x$ 后，就可以回答关于 $x$ 的询问了。

路径异或和的意思是，计算从根节点 $0$ 到每个节点的路径点权异或和。

定义 $s(x)$ 表示根节点 $0$ 到节点 $x$ 的路径异或和。

⚠**注意**：路径异或和一定是以 $0$ 为起点的路径，而不是以询问的 $u$ 为起点的路径。

对于子树 $x$ 中的节点 $z$，把 $s(z)$ 添加到一个**有序集合**中。

在有序集合中查询第 $k$ 小，就是答案。

## 启发式合并

怎么计算子树 $x$ 的有序集合 $A$？

对于 $x$ 的儿子子树 $y$，设 $y$ 对应的有序集合为 $B$，我们做个合并操作：把 $B$ 中的每个数加到 $A$ 中。

但这样做的话，如果树是一条链（或者其他高度很大的树），每次合并操作的时间复杂度是 $\mathcal{O}(n)$，总共会发生 $\mathcal{O}(n^2)$ 次合并，太慢了。

**启发式合并**：每次把小的集合合并到大的集合。如果集合 $B$ 更大，那么就交换集合 $A$ 和 $B$，然后再合并。相当于把 $A$ 的元素添加到 $B$ 中。

考察集合中的一个元素 $v$，如果 $v$ 要添加到另一个更大的集合，那么合并后，$v$ 所处的集合的大小，一定比 $v$ 之前所在集合的大小大一倍。

最坏情况下 $v$ 所处的集合大小为 $1\to 2\to 4\to 8\to \cdots n$，至多合并 $\mathcal{O}(\log n)$ 次。也就是说，每个元素至多参与 $\mathcal{O}(\log n)$ 次合并，所以总的合并次数为 $\mathcal{O}(n\log n)$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def kthSmallest(self, par: List[int], vals: List[int], queries: List[List[int]]) -> List[int]:
        n = len(par)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[par[i]].append(i)

        qs = [[] for _ in range(n)]
        for i, (x, k) in enumerate(queries):
            qs[x].append((k, i))

        ans = [-1] * len(queries)

        def dfs(x: int, xor: int) -> SortedSet:
            xor ^= vals[x]

            st = SortedSet()
            st.add(xor)
            for y in g[x]:
                set_y = dfs(y, xor)
                # 启发式合并：小集合并入大集合
                if len(set_y) > len(st):
                    st, set_y = set_y, st
                for v in set_y:
                    st.add(v)

            for k, qi in qs[x]:
                if k - 1 < len(st):
                    ans[qi] = st[k - 1]

            return st

        dfs(0, 0)
        return ans
```

```cpp [sol-C++]
#include <ext/pb_ds/assoc_container.hpp>

using namespace __gnu_pbds;
using ordered_set = tree<int, null_type, less<int>, rb_tree_tag, tree_order_statistics_node_update>;

class Solution {
public:
    vector<int> kthSmallest(vector<int>& par, vector<int>& vals, vector<vector<int>>& queries) {
        int n = par.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[par[i]].push_back(i);
        }

        int m = queries.size();
        vector<vector<pair<int, int>>> qs(n);
        for (int i = 0; i < m; i++) {
            int x = queries[i][0], k = queries[i][1];
            qs[x].emplace_back(k, i);
        }

        vector<int> ans(m, -1);
        // 必须返回指针，直接返回 ordered_set 会重新 build 一个新的有序集合，导致超时
        auto dfs = [&](this auto&& dfs, int x, int xor_) -> ordered_set* {
            xor_ ^= vals[x];

            ordered_set* st = new ordered_set();
            st->insert(xor_);
            for (int y : g[x]) {
                ordered_set* st_y = dfs(y, xor_);
                // 启发式合并：小集合并入大集合
                if (st_y->size() > st->size()) {
                    swap(st, st_y);
                }
                for (int v : *st_y) {
                    st->insert(v);
                }
                delete st_y;
            }

            for (auto& [k, idx] : qs[x]) {
                if (k - 1 < st->size()) {
                    ans[idx] = *st->find_by_order(k - 1);
                }
            }
            return st;
        };
        dfs(0, 0);

        return ans;
    }
};
```

```go [sol-Go]
// 泛型 Treap 模板（set 版本，不含重复元素）
type nodeS[K comparable] struct {
	son      [2]*nodeS[K]
	priority uint
	key      K
	subSize  int
}

func (o *nodeS[K]) size() int {
	if o != nil {
		return o.subSize
	}
	return 0
}

func (o *nodeS[K]) maintain() {
	o.subSize = 1 + o.son[0].size() + o.son[1].size()
}

func (o *nodeS[K]) rotate(d int) *nodeS[K] {
	x := o.son[d^1]
	o.son[d^1] = x.son[d]
	x.son[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treapS[K comparable] struct {
	rd         uint
	root       *nodeS[K]
	comparator func(a, b K) int
}

func (t *treapS[K]) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treapS[K]) size() int   { return t.root.size() }
func (t *treapS[K]) empty() bool { return t.size() == 0 }

func (t *treapS[K]) _put(o *nodeS[K], key K) *nodeS[K] {
	if o == nil {
		o = &nodeS[K]{priority: t.fastRand(), key: key}
	} else {
		c := t.comparator(key, o.key)
		if c != 0 {
			d := 0
			if c > 0 {
				d = 1
			}
			o.son[d] = t._put(o.son[d], key)
			if o.son[d].priority > o.priority {
				o = o.rotate(d ^ 1)
			}
		}
	}
	o.maintain()
	return o
}

func (t *treapS[K]) put(key K) { t.root = t._put(t.root, key) }

func (t *treapS[K]) _delete(o *nodeS[K], key K) *nodeS[K] {
	if o == nil {
		return nil
	}
	if c := t.comparator(key, o.key); c != 0 {
		d := 0
		if c > 0 {
			d = 1
		}
		o.son[d] = t._delete(o.son[d], key)
	} else {
		if o.son[1] == nil {
			return o.son[0]
		}
		if o.son[0] == nil {
			return o.son[1]
		}
		d := 0
		if o.son[0].priority > o.son[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.son[d] = t._delete(o.son[d], key)
	}
	o.maintain()
	return o
}

func (t *treapS[K]) delete(key K) { t.root = t._delete(t.root, key) }

func (t *treapS[K]) min() *nodeS[K] { return t.kth(0) }
func (t *treapS[K]) max() *nodeS[K] { return t.kth(t.size() - 1) }

func (t *treapS[K]) lowerBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else {
			kth += o.son[0].size()
			break
		}
	}
	return
}

func (t *treapS[K]) upperBoundIndex(key K) (kth int) {
	for o := t.root; o != nil; {
		c := t.comparator(key, o.key)
		if c < 0 {
			o = o.son[0]
		} else if c > 0 {
			kth += o.son[0].size() + 1
			o = o.son[1]
		} else {
			kth += o.son[0].size() + 1
			break
		}
	}
	return
}

func (t *treapS[K]) kth(k int) (o *nodeS[K]) {
	if k < 0 || k >= t.root.size() {
		return
	}
	for o = t.root; o != nil; {
		leftSize := o.son[0].size()
		if k < leftSize {
			o = o.son[0]
		} else {
			k -= leftSize + 1
			if k < 0 {
				break
			}
			o = o.son[1]
		}
	}
	return
}

func (t *treapS[K]) prev(key K) *nodeS[K] { return t.kth(t.lowerBoundIndex(key) - 1) }
func (t *treapS[K]) next(key K) *nodeS[K] { return t.kth(t.upperBoundIndex(key)) }

func (t *treapS[K]) find(key K) *nodeS[K] {
	o := t.kth(t.lowerBoundIndex(key))
	if o == nil || o.key != key {
		return nil
	}
	return o
}

func newSet[K cmp.Ordered]() *treapS[K] {
	return &treapS[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: cmp.Compare[K],
	}
}

func newSetWith[K comparable](comp func(a, b K) int) *treapS[K] {
	return &treapS[K]{
		rd:         uint(time.Now().UnixNano()),
		comparator: comp,
	}
}

func kthSmallest(par []int, vals []int, queries [][]int) []int {
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	type pair struct{ k, i int }
	qs := make([][]pair, n)
	for i, q := range queries {
		x, k := q[0], q[1]
		qs[x] = append(qs[x], pair{k, i})
	}

	ans := make([]int, len(queries))
	var dfs func(int, int) *treapS[int]
	dfs = func(x, xor int) *treapS[int] {
		xor ^= vals[x]

		set := newSet[int]()
		set.put(xor)
		for _, y := range g[x] {
			setY := dfs(y, xor)
			// 启发式合并：小集合并入大集合
			if setY.size() > set.size() {
				set, setY = setY, set
			}
			// 中序遍历 setY
			var f func(*nodeS[int])
			f = func(node *nodeS[int]) {
				if node == nil {
					return
				}
				f(node.son[0])
				set.put(node.key)
				f(node.son[1])
			}
			f(setY.root)
		}

		for _, p := range qs[x] {
			node := set.kth(p.k - 1)
			if node == nil {
				ans[p.i] = -1
			} else {
				ans[p.i] = node.key
			}
		}

		return set
	}
	dfs(0, 0)

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log^2 n + q\log n)$，其中 $n$ 是 $\textit{par}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。有 $\mathcal{O}(n\log n)$ 次合并，每次合并需要 $\mathcal{O}(\log n)$ 的时间插入有序集合。每次询问查询第 $k$ 小需要 $\mathcal{O}(\log n)$ 的时间。
- 空间复杂度：$\mathcal{O}(n + q)$。

## 思考题

把题目改成：对于每个询问 $[u,k]$，计算在子树 $u$ 中，从 $u$ 往下走的第 $k$ 小路径异或和。（本题是从 $0$ 往下走，这里改成从 $u$ 往下走。）

欢迎在评论区分享你的思路/代码。

## 相关题目

1. 数据结构题单的「**专题：离线算法**」。
2. 树题单的「**§3.10 树上启发式合并**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. 【本题相关】[链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
