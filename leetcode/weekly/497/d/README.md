## 一

如果子序列的 GCD 等于 $p$，那么 $p$ 是子序列每个元素的因子，或者说，子序列每个元素都是 $p$ 的倍数。

对于 $\textit{nums}$ 中的不是 $p$ 的倍数的数，一定不能在子序列中。

参与 GCD 的元素越多，GCD 越小。把所有是 $p$ 的倍数的数都选上，如果 GCD 不等于 $p$，那么选更少的数，GCD 也无法等于 $p$。

所以我们必须知道所有是 $p$ 的倍数的数的 GCD。本题有单点修改操作，这可以用**线段树**维护。

然而，本题不能把 $\textit{nums}$ 的所有数都选上，这是本题的难点。

换句话说，如果 $\textit{nums}$ 的每个数都是 $p$ 的倍数，且 $\textit{nums}$ 所有数的 GCD 等于 $p$，我们必须删除一个数（不选一个数）。

## 二

万一删除一个数后，剩余元素的 GCD 大于 $p$ 呢？

这样的数组长什么样？删除任意一个数，都会导致剩余元素的 GCD 大于 $p$。

设 $a[i] = \dfrac{nums[i]}{p}$，那么 $a$ 的所有数的 GCD 等于 $1$。

如果删除 $a[i]$ 后，剩余 $n-1$ 个数的 GCD 大于 $1$，那么剩余 $n-1$ 个数都包含某个质因子 $q_i$，且 $a[i]$ 不含质因子 $q_i$。

要让每个 $a[i]$ 都具有这样的性质，至少要有 $n$ 个**不同**的质数，每个 $a[i]$ 至少是 $n-1$ 个不同质数的乘积。

由于 $2\times 3\times 5 \times 7 \times 11 \times 13 = 30030$，再乘一个质数就超过值域上界 $5\times 10^4$，所以 $a[i]$ 至多是 $6$ 个不同质数的乘积。所以如果「删除一个数后，所有数的 GCD 大于 $p$」，那么必须满足 $n-1\le 6$，即 $n\le 7$。

如果 $n>7$，那么一定可以删除一个数，使得剩余元素的 GCD 等于 $p$。

所以我们只需处理 $n\le 7$ 的情况，暴力枚举删除的数即可。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
# 线段树有两个下标，一个是线段树节点的下标，另一个是线段树维护的区间的下标
# 节点的下标：从 1 开始，如果你想改成从 0 开始，需要把左右儿子下标分别改成 node*2+1 和 node*2+2
# 区间的下标：从 0 开始
class SegmentTree:
    def __init__(self, arr, target_gcd: int, default=0):
        # 线段树维护一个长为 n 的数组（下标从 0 到 n-1）
        # arr 可以是 list 或者 int
        # 如果 arr 是 int，视作数组大小，默认值为 default
        if isinstance(arr, int):
            arr = [default] * arr
        n = len(arr)
        self._target_gcd = target_gcd
        self._n = n
        self._tree = [0] * (2 << (n - 1).bit_length())
        self._build(arr, 1, 0, n - 1)

    # 合并左右儿子的 val 到当前节点的 val
    def _maintain(self, node: int) -> None:
        self._tree[node] = gcd(self._tree[node * 2], self._tree[node * 2 + 1])

    # 用 a 初始化线段树
    # 时间复杂度 O(n)
    def _build(self, a: List[int], node: int, l: int, r: int) -> None:
        if l == r:  # 叶子
            self._tree[node] = a[l] if a[l] % self._target_gcd == 0 else 0  # 初始化叶节点的值
            return
        m = (l + r) // 2
        self._build(a, node * 2, l, m)  # 初始化左子树
        self._build(a, node * 2 + 1, m + 1, r)  # 初始化右子树
        self._maintain(node)

    def _update(self, node: int, l: int, r: int, i: int, val: int) -> None:
        if l == r:  # 叶子（到达目标）
            self._tree[node] = val if val % self._target_gcd == 0 else 0
            return
        m = (l + r) // 2
        if i <= m:  # i 在左子树
            self._update(node * 2, l, m, i, val)
        else:  # i 在右子树
            self._update(node * 2 + 1, m + 1, r, i, val)
        self._maintain(node)

    def _query(self, node: int, l: int, r: int, ql: int, qr: int) -> int:
        if ql > qr:
            return 0
        if ql <= l and r <= qr:  # 当前子树完全在 [ql, qr] 内
            return self._tree[node]
        m = (l + r) // 2
        if qr <= m:  # [ql, qr] 在左子树
            return self._query(node * 2, l, m, ql, qr)
        if ql > m:  # [ql, qr] 在右子树
            return self._query(node * 2 + 1, m + 1, r, ql, qr)
        l_res = self._query(node * 2, l, m, ql, qr)
        r_res = self._query(node * 2 + 1, m + 1, r, ql, qr)
        return gcd(l_res, r_res)

    # 更新 a[i]
    # 时间复杂度 O(log n)
    def update(self, i: int, val: int) -> None:
        self._update(1, 0, self._n - 1, i, val)

    # 返回用 gcd 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
    # 时间复杂度 O(log n)
    def query(self, ql: int, qr: int) -> int:
        return self._query(1, 0, self._n - 1, ql, qr)

    def query_all(self) -> int:
        return self._tree[1]

    def check(self, n: int) -> bool:
        return any(gcd(self.query(0, i - 1), self.query(i + 1, n - 1)) == self._target_gcd for i in range(n))


class Solution:
    def countGoodSubseq(self, nums: List[int], p: int, queries: List[List[int]]) -> int:
        n = len(nums)
        cnt_p = sum(x % p == 0 for x in nums)

        t = SegmentTree(nums, p)
        ans = 0

        for i, x in queries:
            if nums[i] % p == 0:
                cnt_p -= 1
            if x % p == 0:
                cnt_p += 1
            nums[i] = x
            t.update(i, x)

            if t.query_all() == p and (cnt_p < n or n > 7 or t.check(n)):
                ans += 1

        return ans
```

```java [sol-Java]
// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
// 线段树有两个下标，一个是线段树节点的下标，另一个是线段树维护的区间的下标
// 节点的下标：从 1 开始，如果你想改成从 0 开始，需要把左右儿子下标分别改成 node*2+1 和 node*2+2
// 区间的下标：从 0 开始
class SegmentTree {
    private final int targetGcd;
    private final int n;
    private final int[] tree;

    // 线段树维护数组 a
    public SegmentTree(int[] a, int targetGcd) {
        this.targetGcd = targetGcd;
        n = a.length;
        tree = new int[2 << (32 - Integer.numberOfLeadingZeros(n - 1))];
        build(a, 1, 0, n - 1);
    }

    // 更新 a[i]
    // 时间复杂度 O(log n)
    public void update(int i, int val) {
        update(1, 0, n - 1, i, val);
    }

    // 返回用 gcd 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
    // 时间复杂度 O(log n)
    public int query(int ql, int qr) {
        return query(1, 0, n - 1, ql, qr);
    }

    public int queryAll() {
        return tree[1];
    }

    public boolean check(int n) {
        for (int i = 0; i < n; i++) {
            if (gcd(query(0, i - 1), query(i + 1, n - 1)) == targetGcd) {
                return true;
            }
        }
        return false;
    }

    // 合并左右儿子的 val 到当前节点的 val
    private void maintain(int node) {
        tree[node] = gcd(tree[node * 2], tree[node * 2 + 1]);
    }

    // 用 a 初始化线段树
    // 时间复杂度 O(n)
    private void build(int[] a, int node, int l, int r) {
        if (l == r) { // 叶子
            tree[node] = a[l] % targetGcd == 0 ? a[l] : 0; // 初始化叶节点的值
            return;
        }
        int m = (l + r) / 2;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    private void update(int node, int l, int r, int i, int val) {
        if (l == r) { // 叶子（到达目标）
            tree[node] = val % targetGcd == 0 ? val : 0;
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) { // i 在左子树
            update(node * 2, l, m, i, val);
        } else { // i 在右子树
            update(node * 2 + 1, m + 1, r, i, val);
        }
        maintain(node);
    }

    private int query(int node, int l, int r, int ql, int qr) {
        if (ql > qr) {
            return 0;
        }
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node];
        }
        int m = (l + r) / 2;
        if (qr <= m) { // [ql, qr] 在左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 在右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        int lRes = query(node * 2, l, m, ql, qr);
        int rRes = query(node * 2 + 1, m + 1, r, ql, qr);
        return gcd(lRes, rRes);
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}

class Solution {
    public int countGoodSubseq(int[] nums, int p, int[][] queries) {
        int n = nums.length;
        int cntP = 0;
        for (int x : nums) {
            if (x % p == 0) {
                cntP++;
            }
        }

        SegmentTree t = new SegmentTree(nums, p);
        int ans = 0;

        for (int[] q : queries) {
            int i = q[0];
            int x = q[1];

            if (nums[i] % p == 0) {
                cntP--;
            }
            if (x % p == 0) {
                cntP++;
            }
            nums[i] = x;
            t.update(i, x);

            if (t.queryAll() == p && (cntP < n || n > 7 || t.check(n))) {
                ans++;
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
// 模板来源 https://leetcode.cn/circle/discuss/mOr1u6/
// 线段树有两个下标，一个是线段树节点的下标，另一个是线段树维护的区间的下标
// 节点的下标：从 1 开始，如果你想改成从 0 开始，需要把左右儿子下标分别改成 node*2+1 和 node*2+2
// 区间的下标：从 0 开始
template<typename T>
class SegmentTree {
    // 注：也可以去掉 template<typename T>，改在这里定义 T
    // 例如 using T = pair<int, int>;

    int target_gcd;
    int n;
    vector<T> tree;

    // 合并左右儿子的 val 到当前节点的 val
    void maintain(int node) {
        tree[node] = gcd(tree[node * 2], tree[node * 2 + 1]);
    }

    // 用 a 初始化线段树
    // 时间复杂度 O(n)
    void build(const vector<T>& a, int node, int l, int r) {
        if (l == r) { // 叶子
            tree[node] = a[l] % target_gcd == 0 ? a[l] : 0; // 初始化叶节点的值
            return;
        }
        int m = (l + r) / 2;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    void update(int node, int l, int r, int i, T val) {
        if (l == r) { // 叶子（到达目标）
            tree[node] = val % target_gcd == 0 ? val : 0;
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) { // i 在左子树
            update(node * 2, l, m, i, val);
        } else { // i 在右子树
            update(node * 2 + 1, m + 1, r, i, val);
        }
        maintain(node);
    }

    T query(int node, int l, int r, int ql, int qr) const {
        if (ql > qr) {
            return 0;
        }
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return tree[node];
        }
        int m = (l + r) / 2;
        if (qr <= m) { // [ql, qr] 在左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 在右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        T l_res = query(node * 2, l, m, ql, qr);
        T r_res = query(node * 2 + 1, m + 1, r, ql, qr);
        return gcd(l_res, r_res);
    }

public:
    // 线段树维护数组 a
    SegmentTree(const vector<T>& a, int target_gcd) : target_gcd(target_gcd), n(a.size()), tree(2 << bit_width(a.size() - 1)) {
        build(a, 1, 0, n - 1);
    }

    // 更新 a[i]
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        update(1, 0, n - 1, i, val);
    }

    // 返回用 gcd 合并所有 a[i] 的计算结果，其中 i 在闭区间 [ql, qr] 中
    // 时间复杂度 O(log n)
    T query(int ql, int qr) const {
        return query(1, 0, n - 1, ql, qr);
    }

    T query_all() const {
        return tree[1];
    }

    bool check(int n) const {
        for (int i = 0; i < n; i++) {
            if (gcd(query(0, i - 1), query(i + 1, n - 1)) == target_gcd) {
                return true;
            }
        }
        return false;
    }
};

class Solution {
public:
    int countGoodSubseq(vector<int>& nums, int p, vector<vector<int>>& queries) {
        int n = nums.size();
        int cnt_p = 0;
        for (int x : nums) {
            cnt_p += x % p == 0;
        }

        SegmentTree<int> t(nums, p);
        int ans = 0;

        for (auto& q : queries) {
            int i = q[0], x = q[1];

            cnt_p -= nums[i] % p == 0;
            cnt_p += x % p == 0;
            nums[i] = x;
            t.update(i, x);

            if (t.query_all() == p && (cnt_p < n || n > 7 || t.check(n))) {
                ans++;
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
var targetGcd int

type seg []struct{ l, r, gcd int }

func (t seg) maintain(o int) {
	t[o].gcd = gcd(t[o<<1].gcd, t[o<<1|1].gcd)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		if a[l]%targetGcd == 0 {
			t[o].gcd = a[l]
		}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i, val int) {
	cur := &t[o]
	if cur.l == cur.r {
		if val%targetGcd == 0 {
			cur.gcd = val
		} else {
			cur.gcd = 0
		}
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) int {
	if l > r {
		return 0
	}
	if l <= t[o].l && t[o].r <= r {
		return t[o].gcd
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return gcd(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func (t seg) check(n int) bool {
	for i := range n {
		if gcd(t.query(1, 0, i-1), t.query(1, i+1, n-1)) == targetGcd {
			return true
		}
	}
	return false
}

func countGoodSubseq(nums []int, p int, queries [][]int) (ans int) {
	targetGcd = p
	cntP := 0
	for _, x := range nums {
		if x%p == 0 {
			cntP++
		}
	}

	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)

	for _, q := range queries {
		i, x := q[0], q[1]

		if nums[i]%p == 0 {
			cntP--
		}
		if x%p == 0 {
			cntP++
		}
		nums[i] = x
		t.update(1, q[0], x)

		if t[1].gcd == p && (cntP < n || n > 7 || t.check(n)) {
			ans++
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U + q(\log n + \log U))$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度，$U=\max(\textit{nums})$。在线段树的 $\texttt{update}$ 过程中，从叶子到根的路径上，GCD 要么不变（此时退出辗转相除过程），要么至少减半（此时继续辗转相除过程），所以辗转相除的**总**循环次数是 $\mathcal{O}(\log n + \log U)$ 的。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§8.3 线段树**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)
