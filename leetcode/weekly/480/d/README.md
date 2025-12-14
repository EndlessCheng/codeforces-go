## 方法一：线段树

**提示**：如果问题的不修改版本可以用**分治**解决，那么问题的带修版本可以用**线段树**解决。

对于 $s$ 的某个子串 $t$ 的最小删除次数，可以拆分为：

- $t$ 的左半段的最小删除次数。
- $t$ 的右半段的最小删除次数。
- 如果 $t$ 左半段的最后一个字母，等于 $t$ 右半段的第一个字母，要多删一次。

例如 $t = \texttt{ABA} + \texttt{AAB}$，左半删除 $0$ 次，右半删除 $1$ 次，合并后中间的 $\texttt{AAA}$ 要多删除一个 $\texttt{A}$，所以 $t = \texttt{ABAAAB}$ 一共要删除 $0+1+1=2$ 次，得到 $\texttt{ABAB}$。

所以线段树的每个节点，要维护：

- 区间左端点字母。
- 区间右端点字母。
- 区间最小删除次数。

对于线段树的叶子节点，保存 $s$ 对应位置的字母，最小删除次数为 $0$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Data:
    __slots__ = 'lc', 'rc', 'del_cnt'

    def __init__(self, lc='', rc='', del_cnt=0):
        self.lc = lc  # 区间左端点字母
        self.rc = rc  # 区间右端点字母
        self.del_cnt = del_cnt  # 区间删除次数


# 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree:
    def __init__(self, s: str):
        self.n = n = len(s)
        self.t = [Data() for _ in range(2 << (n - 1).bit_length())]
        self._build(s, 1, 0, n - 1)

    def _merge(self, l: Data, r: Data) -> Data:
        # 端点相同时，合并后多删一次
        return Data(l.lc, r.rc, l.del_cnt + r.del_cnt + (l.rc == r.lc))

    def _maintain(self, node: int) -> None:
        self.t[node] = self._merge(self.t[node * 2], self.t[node * 2 + 1])

    def _build(self, a: str, node: int, l: int, r: int) -> None:
        if l == r:  # 叶子节点
            self.t[node].lc = self.t[node].rc = ord(a[l]) - ord('A')
            return
        m = (l + r) // 2
        self._build(a, node * 2, l, m)  # 初始化左子树
        self._build(a, node * 2 + 1, m + 1, r)  # 初始化右子树
        self._maintain(node)

    def _flip(self, node: int, l: int, r: int, i: int) -> None:
        if l == r:  # 叶子（到达目标）
            self.t[node].lc ^= 1
            self.t[node].rc ^= 1
            return
        m = (l + r) // 2
        if i <= m:  # i 在左子树
            self._flip(node * 2, l, m, i)
        else:  # i 在右子树
            self._flip(node * 2 + 1, m + 1, r, i)
        self._maintain(node)

    def flip(self, i: int) -> None:
        self._flip(1, 0, self.n - 1, i)

    def _query(self, node: int, l: int, r: int, ql: int, qr: int) -> Data:
        if ql <= l and r <= qr:  # 当前子树完全在 [ql, qr] 内
            return self.t[node]
        m = (l + r) // 2
        if qr <= m:  # [ql, qr] 在左子树
            return self._query(node * 2, l, m, ql, qr)
        if ql > m:  # [ql, qr] 在右子树
            return self._query(node * 2 + 1, m + 1, r, ql, qr)
        l_res = self._query(node * 2, l, m, ql, qr)
        r_res = self._query(node * 2 + 1, m + 1, r, ql, qr)
        return self._merge(l_res, r_res)

    def query(self, ql: int, qr: int) -> int:
        return self._query(1, 0, self.n - 1, ql, qr).del_cnt


class Solution:
    def minDeletions(self, s: str, queries: List[List[int]]) -> List[int]:
        t = SegmentTree(s)
        ans = []
        for q in queries:
            if q[0] == 1:
                t.flip(q[1])
            else:
                ans.append(t.query(q[1], q[2]))
        return ans
```

```java [sol-Java]
// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree {
    private static class Data {
        byte lc; // 区间左端点字母
        byte rc; // 区间右端点字母
        int del = 0; // 区间删除次数
    }

    private Data merge(Data l, Data r) {
        Data res = new Data();
        res.lc = l.lc;
        res.rc = r.rc;
        res.del = l.del + r.del + (l.rc == r.lc ? 1 : 0); // 端点相同时，合并后多删一次
        return res;
    }

    private final int n;
    private final Data[] t;

    public SegmentTree(String s) {
        n = s.length();
        t = new Data[2 << (32 - Integer.numberOfLeadingZeros(n - 1))];
        build(s.toCharArray(), 1, 0, n - 1);
    }

    public void flip(int i) {
        flip(1, 0, n - 1, i);
    }

    public int query(int ql, int qr) {
        return query(1, 0, n - 1, ql, qr).del;
    }

    private void maintain(int node) {
        t[node] = merge(t[node * 2], t[node * 2 + 1]);
    }

    private void build(char[] a, int node, int l, int r) {
        t[node] = new Data();
        if (l == r) { // 叶子
            t[node].lc = t[node].rc = (byte) (a[l] - 'A');
            return;
        }
        int m = (l + r) / 2;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    private void flip(int node, int l, int r, int i) {
        if (l == r) { // 叶子（到达目标）
            t[node].lc ^= 1;
            t[node].rc ^= 1;
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) { // i 在左子树
            flip(node * 2, l, m, i);
        } else { // i 在右子树
            flip(node * 2 + 1, m + 1, r, i);
        }
        maintain(node);
    }

    private Data query(int node, int l, int r, int ql, int qr) {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return t[node];
        }
        int m = (l + r) / 2;
        if (qr <= m) { // [ql, qr] 在左子树
            return query(node * 2, l, m, ql, qr);
        }
        if (ql > m) { // [ql, qr] 在右子树
            return query(node * 2 + 1, m + 1, r, ql, qr);
        }
        Data lRes = query(node * 2, l, m, ql, qr);
        Data rRes = query(node * 2 + 1, m + 1, r, ql, qr);
        return merge(lRes, rRes);
    }
}

class Solution {
    public int[] minDeletions(String s, int[][] queries) {
        int size = 0;
        for (int[] q : queries) {
            size += q[0] - 1;
        }

        SegmentTree t = new SegmentTree(s);
        int[] ans = new int[size];
        int idx = 0;
        for (int[] q : queries) {
            if (q[0] == 1) {
                t.flip(q[1]);
            } else {
                ans[idx++] = t.query(q[1], q[2]);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree {
    struct T {
        char lc; // 区间左端点字母
        char rc; // 区间右端点字母
        int del = 0; // 区间删除次数
    };

    size_t n;
    vector<T> t;

    T merge(const T& l, const T& r) const {
        return T{l.lc, r.rc, l.del + r.del + (l.rc == r.lc)}; // 端点相同时，合并后多删一次
    }

    void maintain(int node) {
        t[node] = merge(t[node * 2], t[node * 2 + 1]);
    }

    void build(const string& a, int node, int l, int r) {
        if (l == r) { // 叶子
            t[node].lc = t[node].rc = a[l] - 'A';
            return;
        }
        int m = (l + r) / 2;
        build(a, node * 2, l, m); // 初始化左子树
        build(a, node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node);
    }

    void flip(int node, int l, int r, int i) {
        if (l == r) { // 叶子（到达目标）
            t[node].lc ^= 1;
            t[node].rc ^= 1;
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) { // i 在左子树
            flip(node * 2, l, m, i);
        } else { // i 在右子树
            flip(node * 2 + 1, m + 1, r, i);
        }
        maintain(node);
    }

    T query(int node, int l, int r, int ql, int qr) const {
        if (ql <= l && r <= qr) { // 当前子树完全在 [ql, qr] 内
            return t[node];
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
        return merge(l_res, r_res);
    }

public:
    SegmentTree(const string& s) : n(s.size()), t(2 << bit_width(n - 1)) {
        build(s, 1, 0, n - 1);
    }

    void flip(int i) {
        flip(1, 0, n - 1, i);
    }

    int query(int ql, int qr) const {
        return query(1, 0, n - 1, ql, qr).del;
    }
};

class Solution {
public:
    vector<int> minDeletions(string s, vector<vector<int>>& queries) {
        SegmentTree t(s);
        vector<int> ans;
        for (auto& q : queries) {
            if (q[0] == 1) {
                t.flip(q[1]);
            } else {
                ans.push_back(t.query(q[1], q[2]));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/ 
type data struct {
	lc  byte // 区间左端点字母
	rc  byte // 区间右端点字母
	del int  // 区间删除次数
}

type seg []data

func merge(l, r data) data {
	ans := l.del + r.del
	if l.rc == r.lc { // 端点相同时，合并后多删一次
		ans++
	}
	return data{l.lc, r.rc, ans}
}

func newSegmentTree(a string) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func (t seg) maintain(node int) {
	t[node] = merge(t[node*2], t[node*2+1])
}

func (t seg) build(a string, node, l, r int) {
	if l == r { // 叶子
		t[node].lc = a[l] - 'A'
		t[node].rc = t[node].lc
		return
	}
	m := (l + r) / 2
	t.build(a, node*2, l, m)     // 初始化左子树
	t.build(a, node*2+1, m+1, r) // 初始化右子树
	t.maintain(node)
}

func (t seg) flip(node, l, r, i int) {
	if l == r { // 叶子（到达目标）
		t[node].lc ^= 1
		t[node].rc ^= 1
		return
	}
	m := (l + r) / 2
	if i <= m { // i 在左子树
		t.flip(node*2, l, m, i)
	} else { // i 在右子树
		t.flip(node*2+1, m+1, r, i)
	}
	t.maintain(node)
}

func (t seg) query(node, l, r, ql, qr int) data {
	if ql <= l && r <= qr { // 当前子树完全在 [ql, qr] 内
		return t[node]
	}
	m := (l + r) / 2
	if qr <= m { // [ql, qr] 在左子树
		return t.query(node*2, l, m, ql, qr)
	}
	if ql > m { // [ql, qr] 在右子树
		return t.query(node*2+1, m+1, r, ql, qr)
	}
	return merge(t.query(node*2, l, m, ql, qr), t.query(node*2+1, m+1, r, ql, qr))
}

func minDeletions(s string, queries [][]int) (ans []int) {
	n := len(s)
	t := newSegmentTree(s)
	for _, q := range queries {
		if q[0] == 1 {
			t.flip(1, 0, n-1, q[1])
		} else {
			ans = append(ans, t.query(1, 0, n-1, q[1], q[2]).del)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 是 $s$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 方法二：树状数组

如果发现 $s[i-1]\ne s[i]$，人为规定删除右边的 $s[i]$。

这启发我们定义长为 $n-1$ 的数组 $b$（下标从 $1$ 到 $n-1$）：

$$
b[i] =
\begin{cases} 
0, & s[i-1]\ne s[i]     \\
1, & s[i-1]= s[i]     \\
\end{cases}
$$

子串 $[l,r]$ 的删除次数，等于 $b$ 中 $[l+1,r]$ 的元素和。

由于有修改操作，需要用树状数组维护。

修改 $s[i]$ 时，先撤销被影响的 $b[i]$ 和 $b[i+1]$，然后修改，然后添加新的 $b[i]$ 和 $b[i+1]$。

```py [sol-Python3]
class FenwickTree:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)  # 使用下标 1 到 n

    # a[i] 增加 val
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def update(self, i: int, val: int) -> None:
        t = self.tree
        while i < len(t):
            t[i] += val
            i += i & -i

    # 计算前缀和 a[1] + ... + a[i]
    # 1 <= i <= n
    # 时间复杂度 O(log n)
    def pre(self, i: int) -> int:
        t = self.tree
        res = 0
        while i > 0:
            res += t[i]
            i &= i - 1
        return res

    # 计算区间和 a[l] + ... + a[r]
    # 1 <= l <= r <= n
    # 时间复杂度 O(log n)
    def query(self, l: int, r: int) -> int:
        if r < l:
            return 0
        return self.pre(r) - self.pre(l - 1)


class Solution:
    def minDeletions(self, s: str, queries: List[List[int]]) -> List[int]:
        n = len(s)
        t = FenwickTree(n - 1)
        for i in range(1, n):
            if s[i - 1] == s[i]:  # 删除 i
                t.update(i, 1)

        s = list(s)
        ans = []
        for q in queries:
            if q[0] == 2:
                ans.append(t.query(q[1] + 1, q[2]))
                continue

            i = q[1]

            # 撤销旧的
            if i > 0 and s[i - 1] == s[i]:
                t.update(i, -1)
            if i < n - 1 and s[i] == s[i + 1]:
                t.update(i + 1, -1)

            s[i] = 'A' if s[i] == 'B' else 'B'  # A 变成 B，B 变成 A

            # 添加新的
            if i > 0 and s[i - 1] == s[i]:
                t.update(i, 1)
            if i < n - 1 and s[i] == s[i + 1]:
                t.update(i + 1, 1)

        return ans
```

```java [sol-Java]
class FenwickTree {
    private final int[] tree;

    public FenwickTree(int n) {
        tree = new int[n + 1]; // 使用下标 1 到 n
    }

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public void update(int i, int val) {
        for (; i < tree.length; i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    public int pre(int i) {
        int res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }

    // 求区间和 a[l] + ... + a[r]
    // 1 <= l <= r <= n
    // 时间复杂度 O(log n)
    public int query(int l, int r) {
        if (r < l) {
            return 0;
        }
        return pre(r) - pre(l - 1);
    }
}

class Solution {
    public int[] minDeletions(String S, int[][] queries) {
        char[] s = S.toCharArray();
        int n = s.length;
        FenwickTree t = new FenwickTree(n - 1);
        for (int i = 1; i < n; i++) {
            if (s[i - 1] == s[i]) { // 删除 i
                t.update(i, 1);
            }
        }

        int size = 0;
        for (int[] q : queries) {
            size += q[0] - 1;
        }

        int[] ans = new int[size];
        int idx = 0;
        for (int[] q : queries) {
            if (q[0] == 2) {
                ans[idx++] = t.query(q[1] + 1, q[2]);
                continue;
            }

            int i = q[1];

            // 撤销旧的
            if (i > 0 && s[i - 1] == s[i]) {
                t.update(i, -1);
            }
            if (i < n - 1 && s[i] == s[i + 1]) {
                t.update(i + 1, -1);
            }

            s[i] ^= 'A' ^ 'B'; // A 变成 B，B 变成 A

            // 添加新的
            if (i > 0 && s[i - 1] == s[i]) {
                t.update(i, 1);
            }
            if (i < n - 1 && s[i] == s[i + 1]) {
                t.update(i + 1, 1);
            }
        }

        return ans;
    }
}
```

```cpp [sol-C++]
template<typename T>
class FenwickTree {
    vector<T> tree;

public:
    // 使用下标 1 到 n
    FenwickTree(int n) : tree(n + 1) {}

    // a[i] 增加 val
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    void update(int i, T val) {
        for (; i < tree.size(); i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    // 时间复杂度 O(log n)
    T pre(int i) const {
        T res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }

    // 求区间和 a[l] + ... + a[r]
    // 1 <= l <= r <= n
    // 时间复杂度 O(log n)
    T query(int l, int r) const {
        if (r < l) {
            return 0;
        }
        return pre(r) - pre(l - 1);
    }
};

class Solution {
public:
    vector<int> minDeletions(string s, vector<vector<int>>& queries) {
        int n = s.size();
        FenwickTree<int> t(n - 1);
        for (int i = 1; i < n; i++) {
            if (s[i - 1] == s[i]) { // 删除 i
                t.update(i, 1);
            }
        }

        vector<int> ans;
        for (auto& q : queries) {
            if (q[0] == 2) {
                ans.push_back(t.query(q[1] + 1, q[2]));
                continue;
            }

            int i = q[1];

            // 撤销旧的
            if (i > 0 && s[i - 1] == s[i]) {
                t.update(i, -1);
            }
            if (i < n - 1 && s[i] == s[i + 1]) {
                t.update(i + 1, -1);
            }

            s[i] ^= 'A' ^ 'B'; // A 变成 B，B 变成 A

            // 添加新的
            if (i > 0 && s[i - 1] == s[i]) {
                t.update(i, 1);
            }
            if (i < n - 1 && s[i] == s[i + 1]) {
                t.update(i + 1, 1);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
type fenwick []int

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

// 求区间和 a[l] + ... + a[r]
// 1 <= l <= r <= n
// 时间复杂度 O(log n)
func (f fenwick) query(l, r int) int {
	if r < l {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func minDeletions(s string, queries [][]int) (ans []int) {
	n := len(s)
	t := newFenwickTree(n - 1)
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] { // 删除 i
			t.update(i, 1)
		}
	}

	bs := []byte(s)
	for _, q := range queries {
		if q[0] == 2 {
			ans = append(ans, t.query(q[1]+1, q[2]))
			continue
		}

		i := q[1]

		// 撤销旧的
		if i > 0 && bs[i-1] == bs[i] {
			t.update(i, -1)
		}
		if i < n-1 && bs[i] == bs[i+1] {
			t.update(i+1, -1)
		}

		bs[i] ^= 'A' ^ 'B' // A 变成 B，B 变成 A

		// 添加新的
		if i > 0 && bs[i-1] == bs[i] {
			t.update(i, 1)
		}
		if i < n-1 && bs[i] == bs[i+1] {
			t.update(i+1, 1)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n + q)\log n)$，其中 $n$ 是 $s$ 的长度，$q$ 是 $\textit{queries}$ 的长度。也可以 $\mathcal{O}(n)$ 建树，做到 $\mathcal{O}(n+q\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 专题训练

1. 数据结构题单的「**§8.3 线段树（无区间更新）**」中标有「**分治**」的题目。
2. 数据结构题单的「**§8.1 树状数组**」。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
