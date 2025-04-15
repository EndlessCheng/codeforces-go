## 题意解读

题目本质上是求：$\textit{nums}_1$ 和 $\textit{nums}_2$ 的长度恰好为 $3$ 的**公共子序列**的个数。

你可能想到了 [1143. 最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/)。但本题 $n$ 太大，写 $\mathcal{O}(n^2)$ 的 DP 太慢。

## 核心思路

本题是排列，所有元素互不相同。如果可以通过某种方法，把 $\textit{nums}_1$ 变成 $[0,1,2,\dots,n-1]$，我们就能把「公共子序列问题」变成「严格递增子序列问题」，后者有更好的性质，可以更快地求解。

此外，本题子序列长度为 $3$，对于 $3$ 个数的问题，通常可以枚举中间那个数。

## 前置知识：置换

**置换**是一个排列到另一个排列的双射。

以示例 2 为例，定义如下置换 $P(x)$：

$$
\begin{pmatrix}
x & 0 & 1 & 2 & 3 & 4       \\
P(x) & 1 & 2 & 4 & 3 & 0    \\
\end{pmatrix}
$$

把 $\textit{nums}_1=[4,0,1,3,2]$ 中的每个元素 $x$ 替换为 $P(x)$，可以得到一个单调递增的排列 $A=[0,1,2,3,4]$。

把 $\textit{nums}_2=[4,1,0,2,3]$ 中的每个元素 $x$ 替换为 $P(x)$，可以得到一个新的排列 $B=[0,2,1,4,3]$。

> 解释：比如 $1$ 替换成 $P(1)=2$，意思是所有等于 $1$ 的元素（$\textit{nums}_1[2]$ 和 $\textit{nums}_2[1]$）都要替换成 $2$。替换后 $A[2]=2$，$B[1]=2$。

在置换之前，$(4,0,3)$ 是两个排列的公共子序列。

在置换之后，$(P(4),P(0),P(3))=(0,1,3)$ 也是两个新的排列的公共子序列。

⚠**注意**：置换不是排序，是**映射**（可以理解成**重命名**），原来的公共子序列在映射后，子序列元素的**位置没变**，只是数值变了，仍然是公共子序列。所以置换不会改变公共子序列的个数。

## 思路

把 $\textit{nums}_1$ 置换成排列 $A=[0,1,2,\dots, n-1]$，设这一置换为 $P(x)$。把 $P(x)$ 也应用到 $\textit{nums}_2$ 上，得到排列 $B$。

置换后，我们要找的长为 $3$ 的公共子序列，一定是严格递增的。由于 $A$ 的所有子序列都是严格递增的，我们**只需关注** $B$。现在问题变成：

- $B$ 中有多少个长为 $3$ 的严格递增子序列？

对于长为 $3$ 的严格递增子序列 $(x,y,z)$，枚举中间元素 $y$。现在问题变成：

- 在 $B$ 中，元素 $y$ 的左侧有多少个比 $y$ 小的数 $x$？右侧有多少个比 $y$ 大的数 $z$？

枚举 $y=B[i]$，设 $i$ 左侧有 $\textit{less}_y$ 个元素比 $y$ 小，那么 $i$ 左侧有 $i-\textit{less}_y$ 个元素比 $y$ 大。在整个排列 $B$ 中，比 $y$ 大的数有 $n-1-y$ 个，减去 $i-\textit{less}_y$，得到 $i$ 右侧有 $n-1-y-(i-\textit{less}_y)$ 个数比 $y$ 大。所以（根据乘法原理）中间元素是 $y$ 的长为 $3$ 的严格递增子序列的个数为

$$
\textit{less}_y\cdot(n-1-y-(i-\textit{less}_y))
$$

枚举 $y=B[i]$，计算上式，加入答案。

如何计算 $\textit{less}_y$？这可以用**值域树状数组**（或者有序集合）。关于树状数组的原理，请看 [带你发明树状数组！附数学证明](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/)。

值域树状数组的意思是，把元素值视作下标。添加一个值为 $3$ 的数，就是调用树状数组的 $\texttt{update}(3,1)$。查询小于 $3$ 的元素个数，即小于等于 $2$ 的元素个数，就是调用树状数组的 $\texttt{pre}(2)$。完整的树状数组模板，见 [数据结构题单](https://leetcode.cn/circle/discuss/mOr1u6/)。

由于本题元素值是从 $0$ 开始的，但树状数组的下标是从 $1$ 开始的，所以把元素值转成下标，要加一。

```py [sol-Python3]
class FenwickTree:
    def __init__(self, n: int):
        self.tree = [0] * (n + 1)  # 使用下标 1 到 n

    # a[i] 增加 val
    # 1 <= i <= n
    def update(self, i: int, val: int) -> None:
        while i < len(self.tree):
            self.tree[i] += val
            i += i & -i

    # 计算前缀和 a[1] + ... + a[i]
    # 1 <= i <= n
    def pre(self, i: int) -> int:
        res = 0
        while i > 0:
            res += self.tree[i]
            i &= i - 1
        return res

class Solution:
    def goodTriplets(self, nums1: List[int], nums2: List[int]) -> int:
        n = len(nums1)
        p = [0] * n
        for i, x in enumerate(nums1):
            p[x] = i

        ans = 0
        t = FenwickTree(n)
        for i, y in enumerate(nums2):
            y = p[y]
            less = t.pre(y)
            ans += less * (n - 1 - y - (i - less))
            t.update(y + 1, 1)
        return ans
```

```py [sol-Python3 SortedList]
class Solution:
    def goodTriplets(self, nums1: List[int], nums2: List[int]) -> int:
        n = len(nums1)
        p = [0] * n
        for i, x in enumerate(nums1):
            p[x] = i

        ans = 0
        sl = SortedList()
        for i, y in enumerate(nums2):
            y = p[y]
            less = sl.bisect_left(y)  # sl 的 [0,less-1] 中的数都是小于 y 的，这有 less 个
            ans += less * (n - 1 - y - (i - less))
            sl.add(y)
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
    public void update(int i, long val) {
        for (; i < tree.length; i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    public int pre(int i) {
        int res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }
}

class Solution {
    public long goodTriplets(int[] nums1, int[] nums2) {
        int n = nums1.length;
        int[] p = new int[n];
        for (int i = 0; i < n; i++) {
            p[nums1[i]] = i;
        }

        long ans = 0;
        FenwickTree t = new FenwickTree(n);
        for (int i = 0; i < n - 1; i++) {
            int y = p[nums2[i]];
            int less = t.pre(y);
            ans += (long) less * (n - 1 - y - (i - less));
            t.update(y + 1, 1);
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
    void update(int i, T val) {
        for (; i < tree.size(); i += i & -i) {
            tree[i] += val;
        }
    }

    // 求前缀和 a[1] + ... + a[i]
    // 1 <= i <= n
    T pre(int i) const {
        T res = 0;
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }
};

class Solution {
public:
    long long goodTriplets(vector<int>& nums1, vector<int>& nums2) {
        int n = nums1.size();
        vector<int> p(n);
        for (int i = 0; i < n; i++) {
            p[nums1[i]] = i;
        }

        long long ans = 0;
        FenwickTree<int> t(n);
        for (int i = 0; i < n - 1; i++) {
            int y = p[nums2[i]];
            int less = t.pre(y);
            ans += 1LL * less * (n - 1 - y - (i - less));
            t.update(y + 1, 1);
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
func (f fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func goodTriplets(nums1, nums2 []int) (ans int64) {
	n := len(nums1)
	p := make([]int, n)
	for i, x := range nums1 {
		p[x] = i
	}

	t := newFenwickTree(n)
	for i, y := range nums2[:n-1] {
		y = p[y]
		less := t.pre(y)
		ans += int64(less) * int64(n-1-y-(i-less))
		t.update(y+1, 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$。其中 $n$ 是 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 进阶问题

把 $3$ 改成 $4$ 怎么做？改成 $k$ 怎么做？

欢迎在评论区分享你的思路/代码。

## 相似题目

- [1713. 得到子序列的最少操作次数](https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence/)

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
