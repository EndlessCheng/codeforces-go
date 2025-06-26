## 方法一：有序集合 + 启发式合并

> **离线**：按照自己定义的某种顺序回答询问，而不是按照输入顺序 $\textit{queries}[0],\textit{queries}[1],\cdots$ 回答询问。

离线，把询问按照 $u$ 分组，这样我们可以一边 DFS 这棵树，一边回答询问。递归完子树 $x$ 后，就可以回答关于 $x$ 的询问了。

定义 $s(x)$ 表示根节点 $0$ 到节点 $x$ 的路径异或和。

⚠**注意**：路径异或和一定是以 $0$ 为起点的路径，而不是以询问的 $u$ 为起点的路径。

对于子树 $x$ 中的节点 $z$，把 $s(z)$ 添加到一个**有序集合**中。

有序集合中的第 $k$ 小，即为答案。

### 启发式合并

怎么计算子树 $x$ 的有序集合 $A$？

对于 $x$ 的儿子子树 $y$，设 $y$ 对应的有序集合为 $B$，我们做个合并操作：把 $B$ 中的每个数加到 $A$ 中。

但这样做的话，如果树是一条链（或者其他高度很大的树），每次合并操作的时间复杂度是 $\mathcal{O}(n)$，总共会发生 $\mathcal{O}(n^2)$ 次合并，太慢了。

**启发式合并**：每次把小的集合合并到大的集合。如果集合 $B$ 更大，那么就交换集合 $A$ 和 $B$，然后再合并。相当于把 $A$ 的元素添加到 $B$ 中。

这样做的合并次数是多少？

考察集合中的一个元素 $v$，如果 $v$ 要添加到另一个更大的集合，那么合并后，$v$ 所处的集合的大小，至少比 $v$ 之前所在集合的大小大一倍。

最坏情况下，$v$ 所处的集合大小的变化情况为 $1\to 2\to 4\to 8\to \cdots n$，合并 $\mathcal{O}(\log n)$ 次。也就是说，每个元素会参与 $\mathcal{O}(\log n)$ 次合并，所以总的合并次数为 $\mathcal{O}(n\log n)$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1qeNRzjEEk/?t=18m07s)，欢迎点赞关注~

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
using ordered_set = tree<int, null_type, less<>, rb_tree_tag, tree_order_statistics_node_update>;

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

        ordered_set* st = dfs(0, 0);
        delete st;

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

## 方法二：DFS 时间戳 + 莫队 + 树状数组第 k 小

### 前置知识

1. [DFS 时间戳——处理树上问题的有力工具](https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/solutions/1625899/dfs-shi-jian-chuo-chu-li-shu-shang-wen-t-x1kk/)
2. [莫队算法](https://oi-wiki.org/misc/mo-algo/)
3. [分析：莫队算法的块大小取多少合适？](https://zhuanlan.zhihu.com/p/1920472309522740969)
4. [树状数组](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/)
5. [权值树状数组第 k 小](https://oi-wiki.org/ds/fenwick/#%E5%8D%95%E7%82%B9%E4%BF%AE%E6%94%B9%E6%9F%A5%E8%AF%A2%E5%85%A8%E5%B1%80%E7%AC%AC-k-%E5%B0%8F)

### 思路

通过 DFS 时间戳，把树上问题转化成子数组中的区间（无重复）第 $k$ 小。

用莫队算法离线回答询问。

现在我们需要一个数据结构，支持：

- 添加元素。
- 删除元素。
- 查询第 $k$ 小（不含重复元素）。

这可以用权值树状数组解决。

为了保证权值树状数组中没有重复元素，我们需要维护在莫队区间中的元素出现次数：

- 如果元素 $x$ 的出现次数从 $0$ 变成 $1$，那么把 $x$ 加到权值树状数组中。
- 如果元素 $x$ 的出现次数从 $1$ 变成 $0$，那么把 $x$ 从权值树状数组中移除。

```py [sol-Python3]
class FenwickTree:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)  # 使用下标 1 到 n
        self.hb = 1 << (n.bit_length() - 1)

    # a[i] 增加 val
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def update(self, i: int, val: int) -> None:
        while i < len(self.tree):
            self.tree[i] += val
            i += i & -i

    # 求第 k 小（k 从 1 开始）
    # 如果不存在第 k 小，返回 n+1
    # 时间复杂度 O(log n)
    def kth(self, k: int) -> int:
        res = 0
        b = self.hb
        while b > 0:
            next_ = res | b
            if next_ < len(self.tree) and k > self.tree[next_]:
                k -= self.tree[next_]
                res = next_
            b >>= 1
        return res + 1

class Solution:
    def kthSmallest(self, par: List[int], vals: List[int], queries: List[List[int]]) -> List[int]:
        n = len(par)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[par[i]].append(i)

        a = [0] * n
        in_ = [0] * n
        out = [0] * n  # 左闭右开
        dfn = 0

        def dfs(x: int, xor: int) -> None:
            nonlocal dfn
            in_[x] = dfn
            xor ^= vals[x]
            a[dfn] = xor
            dfn += 1
            for y in g[x]:
                dfs(y, xor)
            out[x] = dfn

        dfs(0, 0)

        # 排序去重
        b = sorted(set(a))
        # 离散化
        a = [bisect_left(b, v) + 1 for v in a]  # 从 1 开始

        nq = len(queries)
        block_size = ceil(n / sqrt(nq * 2))  # 块大小

        qs = []
        for i, (x, k) in enumerate(queries):
            l, r = in_[x], out[x]  # 左闭右开
            qs.append((l // block_size, l, r, k, i))

        # 奇偶化排序
        qs.sort(key=lambda q: (q[0], q[2] if q[0] % 2 == 0 else -q[2]))

        m = len(b)
        t = FenwickTree(m + 1)
        cnt = [0] * (m + 1)

        def move(i: int, delta: int) -> None:
            v = a[i]
            if delta > 0:
                if cnt[v] == 0:
                    t.update(v, 1)
                cnt[v] += 1
            else:
                cnt[v] -= 1
                if cnt[v] == 0:
                    t.update(v, -1)

        ans = [-1] * nq
        l = r = 0
        for _, ql, qr, k, i in qs:
            while l < ql:
                move(l, -1)
                l += 1
            while l > ql:
                l -= 1
                move(l, 1)
            while r < qr:
                move(r, 1)
                r += 1
            while r > qr:
                r -= 1
                move(r, -1)

            res = t.kth(k) - 1  # 离散化时 +1 了，所以要 -1
            if res < m:
                ans[i] = b[res]
        return ans
```

```java [sol-Java]
class FenwickTree {
    private final int[] tree;
    private final int hb;

    // 使用下标 1 到 n
    public FenwickTree(int n) {
        tree = new int[n + 1];
        hb = Integer.highestOneBit(n);
    }

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public void update(int i, int val) {
        for (; i < tree.length; i += i & -i) {
            tree[i] += val;
        }
    }

    // 求第 k 小（k 从 1 开始）
    // 如果不存在第 k 小，返回 n+1
    // 时间复杂度 O(log n)
    public int kth(int k) {
        int res = 0;
        for (int b = hb; b > 0; b >>= 1) {
            int next = res | b;
            if (next < tree.length && k > tree[next]) {
                k -= tree[next];
                res = next;
            }
        }
        return res + 1;
    }
}

class Solution {
    public int[] kthSmallest(int[] par, int[] vals, int[][] queries) {
        int n = par.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            g[par[i]].add(i);
        }

        int[] a = new int[n];
        int[] in = new int[n];
        int[] out = new int[n]; // 左闭右开
        dfs(0, 0, g, vals, a, in, out);

        // 排序
        int[] b = Arrays.copyOf(a, n);
        Arrays.sort(b);

        // 离散化
        for (int i = 0; i < n; i++) {
            a[i] = Arrays.binarySearch(b, a[i]) + 1; // 从 1 开始
        }

        int nq = queries.length;
        int blockSize = (int) Math.ceil(n / Math.sqrt(nq * 2)); // 块大小

        record Query(int bid, int l, int r, int k, int qid) {}
        Query[] qs = new Query[nq];
        for (int i = 0; i < nq; i++) {
            int[] q = queries[i];
            int l = in[q[0]], r = out[q[0]]; // 左闭右开
            qs[i] = new Query(l / blockSize, l, r, q[1], i);
        }

        Arrays.sort(qs, (x, y) -> {
            if (x.bid != y.bid) {
                return x.bid - y.bid;
            }
            // 奇偶化排序
            return x.bid % 2 == 0 ? x.r - y.r : y.r - x.r;
        });

        int[] cnt = new int[n + 1];
        FenwickTree t = new FenwickTree(n + 1);

        int[] ans = new int[nq];
        int l = 0, r = 0;
        for (Query q : qs) {
            int ql = q.l, qr = q.r, k = q.k, i = q.qid;
            while (l < ql) move(a[l++], -1, cnt, t);
            while (l > ql) move(a[--l],  1, cnt, t);
            while (r < qr) move(a[r++],  1, cnt, t);
            while (r > qr) move(a[--r], -1, cnt, t);

            int res = t.kth(k) - 1; // 离散化时 +1 了，所以要 -1
            ans[i] = res < n ? b[res] : -1;
        }
        return ans;
    }

    private int dfn = 0;

    private void dfs(int x, int xorVal, List<Integer>[] g, int[] vals, int[] a, int[] in, int[] out) {
        in[x] = dfn;
        xorVal ^= vals[x];
        a[dfn++] = xorVal;
        for (int y : g[x]) {
            dfs(y, xorVal, g, vals, a, in, out);
        }
        out[x] = dfn;
    }

    private void move(int v, int delta, int[] cnt, FenwickTree t) {
        if (delta > 0) {
            if (cnt[v] == 0) {
                t.update(v, 1);
            }
            cnt[v]++;
        } else {
            cnt[v]--;
            if (cnt[v] == 0) {
                t.update(v, -1);
            }
        }
    }
}
```

```cpp [sol-C++]
template<typename T>
class FenwickTree {
    vector<T> tree;
    int hb;

public:
    // 使用下标 1 到 n
    FenwickTree(size_t n) : tree(n + 1), hb(1 << (bit_width(n) - 1)) {}

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        for (; i < tree.size(); i += i & -i) {
            tree[i] += val;
        }
    }

    // 求第 k 小（k 从 1 开始）
    // 如果不存在第 k 小，返回 n+1
    // 时间复杂度 O(log n)
    int kth(int k) {
        int res = 0;
        for (int b = hb; b > 0; b >>= 1) {
            int next = res | b;
            if (next < tree.size() && k > tree[next]) {
                k -= tree[next];
                res = next;
            }
        }
        return res + 1;
    }
};

class Solution {
public:
    vector<int> kthSmallest(vector<int>& par, vector<int>& vals, vector<vector<int>>& queries) {
        int n = par.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[par[i]].push_back(i);
        }

        vector<int> a(n);
        vector<pair<int, int>> ranges(n); // 左闭右开 [l,r)
        int dfn = 0;
        auto dfs = [&](this auto&& dfs, int x, int xor_val) -> void {
            ranges[x].first = dfn;
            xor_val ^= vals[x];
            a[dfn++] = xor_val;
            for (int y : g[x]) {
                dfs(y, xor_val);
            }
            ranges[x].second = dfn;
        };
        dfs(0, 0);

        // 排序去重
        vector<int> b = a;
        ranges::sort(b);
        b.erase(ranges::unique(b).begin(), b.end());

        // 离散化
        for (int& v : a) {
            v = ranges::lower_bound(b, v) - b.begin() + 1; // 从 1 开始
        }

        int nq = queries.size();
        int block_size = ceil(n / sqrt(nq * 2)); // 块大小

        struct Query { int bid, l, r, k, qid; };
        vector<Query> qs(nq);
        for (int i = 0; i < nq; i++) {
            auto& q = queries[i];
            auto [l, r] = ranges[q[0]]; // 左闭右开
            qs[i] = {l / block_size, l, r, q[1], i};
        }

        ranges::sort(qs, [](auto& a, auto& b) {
            if (a.bid != b.bid) {
                return a.bid < b.bid;
            }
            // 奇偶化排序
            return a.bid % 2 == 0 ? a.r < b.r : a.r > b.r;
        });

        int m = b.size();
        FenwickTree<int> t(m + 1);
        vector<int> cnt(m + 1);
        auto move = [&](int i, int delta) {
            int v = a[i];
            if (delta > 0) {
                if (cnt[v] == 0) {
                    t.update(v, 1);
                }
                cnt[v]++;
            } else {
                cnt[v]--;
                if (cnt[v] == 0) {
                    t.update(v, -1);
                }
            }
        };

        vector<int> ans(nq, -1);
        int l = 0, r = 0;
        for (auto& [_, ql, qr, k, i] : qs) {
            while (l < ql) move(l++, -1);
            while (l > ql) move(--l, 1);
            while (r < qr) move(r++, 1);
            while (r > qr) move(--r, -1);

            int res = t.kth(k) - 1; // 离散化时 +1 了，所以要 -1
            if (res < m) {
                ans[i] = b[res];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) kth(hb, k int) (res int) {
	for b := hb; b > 0; b >>= 1 {
		next := res | b
		if next < len(f) && k > f[next] {
			k -= f[next]
			res = next
		}
	}
	return res + 1
}

func kthSmallest(par []int, vals []int, queries [][]int) []int {
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	a := make([]int, n)
	ranges := make([]struct{ l, r int }, n) // 左闭右开 [l,r)
	dfn := 0
	var dfs func(int, int)
	dfs = func(x, xor int) {
		ranges[x].l = dfn
		xor ^= vals[x]
		a[dfn] = xor
		dfn++
		for _, y := range g[x] {
			dfs(y, xor)
		}
		ranges[x].r = dfn
	}
	dfs(0, 0)

	// 排序去重
	b := slices.Clone(a)
	slices.Sort(b)
	b = slices.Compact(b)
	// 离散化
	for i, v := range a {
		a[i] = sort.SearchInts(b, v) + 1 // 从 1 开始
	}

	nq := len(queries)
	blockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(nq*2))))
	type query struct{ bid, l, r, k, qid int } // 左闭右开 [l,r)
	qs := make([]query, nq)
	for i, q := range queries {
		p := ranges[q[0]]
		qs[i] = query{p.l / blockSize, p.l, p.r, q[1], i}
	}
	slices.SortFunc(qs, func(a, b query) int {
		if a.bid != b.bid {
			return a.bid - b.bid
		}
		// 奇偶化排序
		if a.bid%2 == 0 {
			return a.r - b.r
		}
		return b.r - a.r
	})

	m := len(b)
	hb := 1 << (bits.Len(uint(m)) - 1)
	t := make(fenwick, m+1)
	cnt := make([]int, m+1)
	move := func(i, delta int) {
		v := a[i]
		if delta > 0 {
			if cnt[v] == 0 {
				t.update(v, 1)
			}
			cnt[v]++
		} else {
			cnt[v]--
			if cnt[v] == 0 {
				t.update(v, -1)
			}
		}
	}

	ans := make([]int, len(qs))
	l, r := 0, 0
	for _, q := range qs {
		for ; l < q.l; l++ {
			move(l, -1)
		}
		for l > q.l {
			l--
			move(l, 1)
		}
		for ; r < q.r; r++ {
			move(r, 1)
		}
		for r > q.r {
			r--
			move(r, -1)
		}

		res := t.kth(hb, q.k) - 1 // 离散化时 +1 了，所以要 -1
		if res < m {
			ans[q.qid] = b[res]
		} else {
			ans[q.qid] = -1
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt q\log n + q\log q)$，其中 $n$ 是 $\textit{par}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。做了 $\mathcal{O}(n\sqrt q)$ 次查询第 $k$ 小，每次需要 $\mathcal{O}(\log n)$ 时间。排序需要 $\mathcal{O}(n\log n + q\log q)$ 时间。
- 空间复杂度：$\mathcal{O}(n + q)$。

**注**：复杂度看上去高，但常数小，实际比方法一块（除了 Python，在方法三中更快）。

## 方法三：DFS 时间戳 + 莫队 + 分块求第 k 小

在方法二中：

- 每次 $\texttt{move}$ 需要 $\mathcal{O}(\log n)$ 时间维护树状数组。
- 每次询问需要 $\mathcal{O}(\log n)$ 时间查询树状数组。

由于 $\texttt{move}$ 很多但询问很少，这两个操作耗费的时间没有做到完全平衡，可以继续优化。

改用 [分块](https://oi-wiki.org/ds/decompose/) 维护第 $k$ 小。

设 $\textit{cnt}$ 的长度为 $m$，把 $\textit{cnt}$ 分割成若干段，每一段的大小为 $B=\left\lfloor\sqrt m\right\rfloor$。额外用一个 $\textit{blockUniqueCnt}$ 数组维护每一段的不同元素个数。

- 每次 $\texttt{move}$ 只需要 $\mathcal{O}(1)$ 时间。
- 每次询问需要 $\mathcal{O}(\sqrt n)$ 时间：
   - 遍历 $\textit{blockUniqueCnt}$。
   - 如果 $k> \textit{blockUniqueCnt}[i]$，把 $k$ 减少 $\textit{blockUniqueCnt}[i]$，继续遍历。
   - 否则，答案在 $\textit{blockUniqueCnt}[i]$ 对应的 $\textit{cnt}$ 数组的子数组 $[i\cdot B, (i+1)\cdot B)$ 中，遍历这个子数组即可。

```py [sol-Python3]
class Solution:
    def kthSmallest(self, par: List[int], vals: List[int], queries: List[List[int]]) -> List[int]:
        n = len(par)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[par[i]].append(i)

        a = [0] * n
        in_ = [0] * n
        out = [0] * n  # 左闭右开
        dfn = 0

        def dfs(x: int, xor: int) -> None:
            nonlocal dfn
            in_[x] = dfn
            xor ^= vals[x]
            a[dfn] = xor
            dfn += 1
            for y in g[x]:
                dfs(y, xor)
            out[x] = dfn

        dfs(0, 0)

        # 排序去重
        b = sorted(set(a))
        # 离散化
        a = [bisect_left(b, v) for v in a]  # 从 0 开始

        nq = len(queries)
        q_block_size = ceil(n / sqrt(nq * 2))  # 块大小

        qs = []
        for i, (x, k) in enumerate(queries):
            l, r = in_[x], out[x]  # 左闭右开
            qs.append((l // q_block_size, l, r, k, i))

        # 奇偶化排序
        qs.sort(key=lambda q: (q[0], q[2] if q[0] % 2 == 0 else -q[2]))

        m = len(b)
        c_block_size = isqrt(m)
        block_unique_cnt = [0] * ((m - 1) // c_block_size + 1)
        cnt = [0] * (m + 1)

        def move(i: int, delta: int) -> None:
            v = a[i]
            if delta > 0:
                if cnt[v] == 0:
                    block_unique_cnt[v // c_block_size] += 1
                cnt[v] += 1
            else:
                cnt[v] -= 1
                if cnt[v] == 0:
                    block_unique_cnt[v // c_block_size] -= 1

        ans = [-1] * nq
        l = r = 0
        for _, ql, qr, k, qid in qs:
            while l < ql:
                move(l, -1)
                l += 1
            while l > ql:
                l -= 1
                move(l, 1)
            while r < qr:
                move(r, 1)
                r += 1
            while r > qr:
                r -= 1
                move(r, -1)

            for i, c in enumerate(block_unique_cnt):
                if k <= c:
                    for j in count(i * c_block_size):
                        if cnt[j] == 0:
                            continue
                        k -= 1
                        if k == 0:
                            ans[qid] = b[j]
                            break
                    break
                k -= c
        return ans
```

```java [sol-Java]
class Solution {
    public int[] kthSmallest(int[] par, int[] vals, int[][] queries) {
        int n = par.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            g[par[i]].add(i);
        }

        int[] a = new int[n];
        int[] in = new int[n];
        int[] out = new int[n]; // 左闭右开
        dfs(0, 0, g, vals, a, in, out);

        // 排序
        int[] b = Arrays.copyOf(a, n);
        Arrays.sort(b);

        // 离散化
        for (int i = 0; i < n; i++) {
            a[i] = Arrays.binarySearch(b, a[i]); // 从 0 开始
        }

        int nq = queries.length;
        int qBlockSize = (int) Math.ceil(n / Math.sqrt(nq * 2)); // 块大小

        record Query(int bid, int l, int r, int k, int qid) {}
        Query[] qs = new Query[nq];
        for (int i = 0; i < nq; i++) {
            int[] q = queries[i];
            int l = in[q[0]], r = out[q[0]]; // 左闭右开
            qs[i] = new Query(l / qBlockSize, l, r, q[1], i);
        }

        Arrays.sort(qs, (x, y) -> {
            if (x.bid != y.bid) {
                return x.bid - y.bid;
            }
            // 奇偶化排序
            return x.bid % 2 == 0 ? x.r - y.r : y.r - x.r;
        });

        int[] cnt = new int[n + 1];
        int cBlockSize = (int) Math.sqrt(n);
        int[] blockUniqueCnt = new int[(n - 1) / cBlockSize + 1];

        int[] ans = new int[nq];
        int l = 0, r = 0;
        for (Query q : qs) {
            int ql = q.l, qr = q.r;
            while (l < ql) move(a[l++], -1, cnt, blockUniqueCnt, cBlockSize);
            while (l > ql) move(a[--l],  1, cnt, blockUniqueCnt, cBlockSize);
            while (r < qr) move(a[r++],  1, cnt, blockUniqueCnt, cBlockSize);
            while (r > qr) move(a[--r], -1, cnt, blockUniqueCnt, cBlockSize);

            int k = q.k;
            for (int i = 0; i < blockUniqueCnt.length; i++) {
                if (k <= blockUniqueCnt[i]) {
                    for (int j = i * cBlockSize; ; j++) {
                        if (cnt[j] > 0 && --k == 0) {
                            ans[q.qid] = b[j];
                            break;
                        }
                    }
                    break;
                }
                k -= blockUniqueCnt[i];
            }
            if (k > 0) {
                ans[q.qid] = -1;
            }
        }
        return ans;
    }

    private int dfn = 0;

    private void dfs(int x, int xorVal, List<Integer>[] g, int[] vals, int[] a, int[] in, int[] out) {
        in[x] = dfn;
        xorVal ^= vals[x];
        a[dfn++] = xorVal;
        for (int y : g[x]) {
            dfs(y, xorVal, g, vals, a, in, out);
        }
        out[x] = dfn;
    }

    private void move(int v, int delta, int[] cnt, int[] blockUniqueCnt, int cBlockSize) {
        if (delta > 0) {
            if (cnt[v] == 0) {
                blockUniqueCnt[v / cBlockSize]++;
            }
            cnt[v]++;
        } else {
            cnt[v]--;
            if (cnt[v] == 0) {
                blockUniqueCnt[v / cBlockSize]--;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> kthSmallest(vector<int>& par, vector<int>& vals, vector<vector<int>>& queries) {
        int n = par.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            g[par[i]].push_back(i);
        }

        vector<int> a(n);
        vector<pair<int, int>> ranges(n); // 左闭右开 [l,r)
        int dfn = 0;
        auto dfs = [&](this auto&& dfs, int x, int xor_val) -> void {
            ranges[x].first = dfn;
            xor_val ^= vals[x];
            a[dfn++] = xor_val;
            for (int y : g[x]) {
                dfs(y, xor_val);
            }
            ranges[x].second = dfn;
        };
        dfs(0, 0);

        // 排序去重
        vector<int> b = a;
        ranges::sort(b);
        b.erase(ranges::unique(b).begin(), b.end());

        // 离散化
        for (int& v : a) {
            v = ranges::lower_bound(b, v) - b.begin(); // 从 0 开始
        }

        int nq = queries.size();
        int q_block_size = ceil(n / sqrt(nq * 2)); // 块大小

        struct Query { int bid, l, r, k, qid; };
        vector<Query> qs(nq);
        for (int i = 0; i < nq; i++) {
            auto& q = queries[i];
            auto [l, r] = ranges[q[0]]; // 左闭右开
            qs[i] = {l / q_block_size, l, r, q[1], i};
        }

        ranges::sort(qs, [](auto& a, auto& b) {
            if (a.bid != b.bid) {
                return a.bid < b.bid;
            }
            // 奇偶化排序
            return a.bid % 2 == 0 ? a.r < b.r : a.r > b.r;
        });

        int m = b.size();
        int c_block_size = sqrt(m);
        vector<int> block_unique_cnt((m - 1) / c_block_size + 1);
        vector<int> cnt(m + 1);

        auto move = [&](int i, int delta) {
            int v = a[i];
            if (delta > 0) {
                if (cnt[v] == 0) {
                    block_unique_cnt[v / c_block_size]++;
                }
                cnt[v]++;
            } else {
                cnt[v]--;
                if (cnt[v] == 0) {
                    block_unique_cnt[v / c_block_size]--;
                }
            }
        };

        vector<int> ans(nq, -1);
        int l = 0, r = 0;
        for (auto& [_, ql, qr, k, qid] : qs) {
            while (l < ql) move(l++, -1);
            while (l > ql) move(--l, 1);
            while (r < qr) move(r++, 1);
            while (r > qr) move(--r, -1);

            for (int i = 0; i < block_unique_cnt.size(); i++) {
                if (k <= block_unique_cnt[i]) {
                    for (int j = i * c_block_size; ; j++) {
                        if (cnt[j] && --k == 0) {
                            ans[qid] = b[j];
                            break;
                        }
                    }
                    break;
                }
                k -= block_unique_cnt[i];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func kthSmallest(par []int, vals []int, queries [][]int) []int {
	n := len(par)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par[i]
		g[p] = append(g[p], i)
	}

	a := make([]int, n)
	ranges := make([]struct{ l, r int }, n) // 左闭右开 [l,r)
	dfn := 0
	var dfs func(int, int)
	dfs = func(x, xor int) {
		ranges[x].l = dfn
		xor ^= vals[x]
		a[dfn] = xor
		dfn++
		for _, y := range g[x] {
			dfs(y, xor)
		}
		ranges[x].r = dfn
	}
	dfs(0, 0)

	// 排序去重
	b := slices.Clone(a)
	slices.Sort(b)
	b = slices.Compact(b)
	// 离散化
	for i, v := range a {
		a[i] = sort.SearchInts(b, v) // 从 0 开始
	}

	nq := len(queries)
	qBlockSize := int(math.Ceil(float64(n) / math.Sqrt(float64(nq*2))))
	type query struct{ bid, l, r, k, qid int } // 左闭右开 [l,r)
	qs := make([]query, nq)
	for i, q := range queries {
		p := ranges[q[0]]
		qs[i] = query{p.l / qBlockSize, p.l, p.r, q[1], i}
	}
	slices.SortFunc(qs, func(a, b query) int {
		if a.bid != b.bid {
			return a.bid - b.bid
		}
		// 奇偶化排序
		if a.bid%2 == 0 {
			return a.r - b.r
		}
		return b.r - a.r
	})

	m := len(b)
	cBlockSize := int(math.Sqrt(float64(m)))
	blockUniqueCnt := make([]int, (m-1)/cBlockSize+1)
	cnt := make([]int, m+1)
	move := func(i, delta int) {
		v := a[i]
		if delta > 0 {
			if cnt[v] == 0 {
				blockUniqueCnt[v/cBlockSize]++
			}
			cnt[v]++
		} else {
			cnt[v]--
			if cnt[v] == 0 {
				blockUniqueCnt[v/cBlockSize]--
			}
		}
	}

	ans := make([]int, len(qs))
	l, r := 0, 0
	for _, q := range qs {
		for ; l < q.l; l++ {
			move(l, -1)
		}
		for l > q.l {
			l--
			move(l, 1)
		}
		for ; r < q.r; r++ {
			move(r, 1)
		}
		for r > q.r {
			r--
			move(r, -1)
		}

		k := q.k
		for i, c := range blockUniqueCnt {
			if k <= c {
				for j := i * cBlockSize; ; j++ {
					if cnt[j] == 0 {
						continue
					}
					k--
					if k == 0 {
						ans[q.qid] = b[j]
						break
					}
				}
				break
			}
			k -= c
		}
		if k > 0 {
			ans[q.qid] = -1
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt q + q\sqrt n + n\log n + q\log q)$，其中 $n$ 是 $\textit{par}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。排序需要 $\mathcal{O}(n\log n + q\log q)$ 时间。
- 空间复杂度：$\mathcal{O}(n + q)$。

## 思考题

把题目改成：对于每个询问 $[u,k]$，计算在子树 $u$ 中，从 $u$ 往下走的第 $k$ 小路径异或和。（本题是从 $0$ 往下走，这里改成从 $u$ 往下走。）

解答见 [视频讲解](https://www.bilibili.com/video/BV1qeNRzjEEk/?t=18m07s)。

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
