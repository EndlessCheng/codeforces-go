**技巧**：如果一个问题可以用**分治**解决，那么这个问题的带修改版本可以用**线段树**解决。

如何用分治计算一个字符串的最长连续相同子串的长度？

把字符串分成左右两部分 $A$ 和 $B$：

1. 计算 $A$ 的最长连续相同子串的长度。这可以递归解决。
2. 计算 $B$ 的最长连续相同子串的长度。这可以递归解决。
3. 计算左端点在 $A$ 中，右端点在 $B$ 中的最长连续相同子串的长度。这相当于求出 $A$ 的最长连续相同**后缀**的长度，以及 $B$ 的最长连续相同**前缀**的长度。如果 $A$ 的最后一个字符等于 $B$ 的第一个字符，就可以把 $A$ 的最长连续相同后缀与 $B$ 的最长连续相同前缀拼起来。

这样思考后，发现本题可以用分治解决，那么带修改版本就可以用线段树解决了。

对于线段树的区间 $[L,R]$，维护如下信息：

- $\textit{mx}$：该区间的最长连续相同子串的长度。这等于上面分治过程的三种情况的最大值。
- $\textit{pre}$：该区间最长连续相同前缀的长度。计算方法见下。
- $\textit{suf}$：该区间最长连续相同后缀的长度。计算方法见下。

我们合并区间 $A$ 和 $B$，设合并后的大区间为 $C$。

- 如果 $A.\textit{pre}$ 等于 $A$ 的长度，且 $A$ 的最后一个字符等于 $B$ 的第一个字符，那么 $C.\textit{pre} = A.\textit{pre} + B.\textit{pre}$。否则，$C.\textit{pre} = A.\textit{pre}$。
- 如果 $B.\textit{suf}$ 等于 $B$ 的长度，且 $A$ 的最后一个字符等于 $B$ 的第一个字符，那么 $C.\textit{suf} = A.\textit{suf} + B.\textit{suf}$。否则，$C.\textit{suf} = B.\textit{suf}$。

```py [sol-Python3]
# 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree:
    def __init__(self, s: list[str]):
        self._n = n = len(s)
        self._s = s
        self._tree = [None] * (2 << (n - 1).bit_length())
        self._build(s, 1, 0, n - 1)

    def _maintain(self, node: int, l: int, m: int, r: int) -> None:
        a_mx, a_pre, a_suf = self._tree[node * 2]
        b_mx, b_pre, b_suf = self._tree[node * 2 + 1]
        same = self._s[m] == self._s[m + 1]  # 左区间的最后一个字符 == 右区间的第一个字符
        self._tree[node] = (
            max(a_mx, b_mx, a_suf + b_pre) if same else max(a_mx, b_mx),
            a_pre + b_pre if same and a_pre == m - l + 1 else a_pre,
            a_suf + b_suf if same and b_suf == r - m else b_suf,
        )

    def _build(self, s: List[str], node: int, l: int, r: int) -> None:
        if l == r:  # 叶子
            self._tree[node] = (1, 1, 1)  # 初始化叶节点的值
            return
        m = (l + r) // 2
        self._build(s, node * 2, l, m)  # 初始化左子树
        self._build(s, node * 2 + 1, m + 1, r)  # 初始化右子树
        self._maintain(node, l, m, r)

    def _update(self, node: int, l: int, r: int, i: int, val: str) -> None:
        if l == r:  # 叶子（到达目标）
            self._s[i] = val
            return
        m = (l + r) // 2
        if i <= m:  # i 在左子树
            self._update(node * 2, l, m, i, val)
        else:  # i 在右子树
            self._update(node * 2 + 1, m + 1, r, i, val)
        self._maintain(node, l, m, r)

    def update(self, i: int, val: str) -> None:
        self._update(1, 0, self._n - 1, i, val)

    def query_all(self) -> int:
        return self._tree[1][0]


class Solution:
    def longestRepeating(self, s: str, queryCharacters: str, queryIndices: List[int]) -> List[int]:
        t = SegmentTree(list(s))
        ans = []
        for i, ch in zip(queryIndices, queryCharacters):
            t.update(i, ch)
            ans.append(t.query_all())
        return ans
```

```java [sol-Java]
// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree {
    private record Data(int mx, int pre, int suf) {
    }

    private final int n;
    private final char[] s;
    private final Data[] tree;

    public SegmentTree(char[] s) {
        n = s.length;
        this.s = s;
        tree = new Data[2 << (32 - Integer.numberOfLeadingZeros(n - 1))];
        build(1, 0, n - 1);
    }

    public void update(int i, char val) {
        update(1, 0, n - 1, i, val);
    }

    public int queryAll() {
        return tree[1].mx;
    }

    private void maintain(int node, int l, int m, int r) {
        Data left = tree[node * 2];
        Data right = tree[node * 2 + 1];
        int mx = Math.max(left.mx, right.mx);
        int pre = left.pre;
        int suf = right.suf;
        if (s[m] == s[m + 1]) { // 左区间的最后一个字符 == 右区间的第一个字符
            mx = Math.max(mx, left.suf + right.pre);
            if (left.pre == m - l + 1) {
                pre += right.pre;
            }
            if (right.suf == r - m) {
                suf += left.suf;
            }
        }
        tree[node] = new Data(mx, pre, suf);
    }

    private void build(int node, int l, int r) {
        if (l == r) { // 叶子
            tree[node] = new Data(1, 1, 1); // 初始化叶节点的值
            return;
        }
        int m = (l + r) >>> 1;
        build(node * 2, l, m); // 初始化左子树
        build(node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node, l, m, r);
    }

    private void update(int node, int l, int r, int i, char val) {
        if (l == r) { // 叶子（到达目标）
            s[i] = val;
            return;
        }
        int m = (l + r) >>> 1;
        if (i <= m) { // i 在左子树
            update(node * 2, l, m, i, val);
        } else { // i 在右子树
            update(node * 2 + 1, m + 1, r, i, val);
        }
        maintain(node, l, m, r);
    }
}

class Solution {
    public int[] longestRepeating(String s, String queryCharacters, int[] queryIndices) {
        SegmentTree t = new SegmentTree(s.toCharArray());
        int q = queryIndices.length;
        int[] ans = new int[q];
        for (int i = 0; i < q; i++) {
            t.update(queryIndices[i], queryCharacters.charAt(i));
            ans[i] = t.queryAll();
        }
        return ans;
    }
}
```

```cpp [sol-C++]
struct Data {
    int mx, pre, suf;
};

// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
class SegmentTree {
    int n;
    string s;
    vector<Data> tree;

    void maintain(int node, int l, int m, int r) {
        Data& left = tree[node * 2];
        Data& right = tree[node * 2 + 1];
        int mx = max(left.mx, right.mx);
        int pre = left.pre;
        int suf = right.suf;
        if (s[m] == s[m + 1]) { // 左区间的最后一个字符 == 右区间的第一个字符
            mx = max(mx, left.suf + right.pre);
            if (left.pre == m - l + 1) {
                pre += right.pre;
            }
            if (right.suf == r - m) {
                suf += left.suf;
            }
        }
        tree[node] = {mx, pre, suf};
    }

    void build(int node, int l, int r) {
        if (l == r) { // 叶子
            tree[node] = {1, 1, 1}; // 初始化叶节点的值
            return;
        }
        int m = (l + r) >> 1;
        build(node * 2, l, m); // 初始化左子树
        build(node * 2 + 1, m + 1, r); // 初始化右子树
        maintain(node, l, m, r);
    }

    void update(int node, int l, int r, int i, char val) {
        if (l == r) { // 叶子（到达目标）
            s[i] = val;
            return;
        }
        int m = (l + r) >> 1;
        if (i <= m) { // i 在左子树
            update(node * 2, l, m, i, val);
        } else { // i 在右子树
            update(node * 2 + 1, m + 1, r, i, val);
        }
        maintain(node, l, m, r);
    }

public:
    SegmentTree(const string& s) : n(s.size()), s(s), tree(2 << bit_width(s.size() - 1)) {
        build(1, 0, n - 1);
    }

    void update(int i, char val) {
        update(1, 0, n - 1, i, val);
    }

    int query_all() const {
        return tree[1].mx;
    }
};

class Solution {
public:
    vector<int> longestRepeating(string s, string queryCharacters, vector<int>& queryIndices) {
        SegmentTree t(s);
        int q = queryIndices.size();
        vector<int> ans(q);
        for (int i = 0; i < q; i++) {
            t.update(queryIndices[i], queryCharacters[i]);
            ans[i] = t.query_all();
        }
        return ans;
    }
};
```

```go [sol-Go]
type data struct{ mx, pre, suf int }

// 详细注释见 https://leetcode.cn/circle/discuss/mOr1u6/
type seg struct {
	s []byte
	d []data
}

func newSegmentTree(s []byte) *seg {
	n := len(s)
	t := &seg{s, make([]data, 2<<bits.Len(uint(n-1)))}
	t.build(1, 0, n-1)
	return t
}

func (t *seg) maintain(node, l, m, r int) {
	a, b := t.d[node*2], t.d[node*2+1]
	mx := max(a.mx, b.mx)
	pre := a.pre
	suf := b.suf
	if t.s[m] == t.s[m+1] { // 左区间的最后一个字符 == 右区间的第一个字符
		mx = max(mx, a.suf+b.pre)
		if pre == m-l+1 {
			pre += b.pre
		}
		if suf == r-m {
			suf += a.suf
		}
	}
	t.d[node] = data{mx, pre, suf}
}

func (t *seg) build(node, l, r int) {
	if l == r { // 叶子
		t.d[node] = data{1, 1, 1} // 初始化叶节点的值
		return
	}
	m := (l + r) >> 1
	t.build(node*2, l, m)     // 初始化左子树
	t.build(node*2+1, m+1, r) // 初始化右子树
	t.maintain(node, l, m, r)
}

func (t *seg) update(node, l, r, i int, ch byte) {
	if l == r { // 叶子（到达目标）
		t.s[i] = ch
		return
	}
	m := (l + r) >> 1
	if i <= m { // i 在左子树
		t.update(node*2, l, m, i, ch)
	} else { // i 在右子树
		t.update(node*2+1, m+1, r, i, ch)
	}
	t.maintain(node, l, m, r)
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	n := len(s)
	t := newSegmentTree([]byte(s))
	ans := make([]int, len(queryIndices))
	for k, i := range queryIndices {
		t.update(1, 0, n-1, i, queryCharacters[k])
		ans[k] = t.d[1].mx
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 是 $s$ 的长度，$q$ 是 $\textit{queryCharacters}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 相似题目

- [3525. 求出数组的 X 值 II](https://leetcode.cn/problems/find-x-value-of-array-ii/)
- [4017. 数组中的峰值 II](https://leetcode.cn/problems/peaks-in-array-ii/)

## 专题训练

见下面数据结构题单的「**§8.3 线段树**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
