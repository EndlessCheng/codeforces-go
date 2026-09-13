> **注**：本题有负数，不能用滑动窗口。

为方便计算，用所有子数组个数减去不遥远的子数组个数，即为遥远的子数组个数。

设 $\textit{nums}$ 的**前缀和**数组为 $s$。关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

不遥远的子数组 $[i,j)$ 需要满足 $i < j$ 且

$$
\textit{goal}-k+1 \le s[j] - s[i] \le \textit{goal}+k-1
$$

枚举 $j$，那么 $i$ 需要满足

$$
s[j] - \textit{goal}-k+1 \le s[i] \le s[j] - \textit{goal}+k-1
$$

用**值域树状数组**保存遍历过的 $s[i]$ 的出现次数，即可 $\mathcal{O}(\log n)$ 计算范围内的 $s[i]$ 的出现次数。

为方便使用树状数组，需要把 $s$ **离散化**。例如 $s = [0, 30, 10, 40, 20]$ 离散化后是 $[0,3,1,4,2]$。如何离散化？一种方法是，把 $s$ 复制一份，排序，然后在有序数组中二分 $s[i]$ 的位置，即为离散化后的值。

本题需要二分找到最后一个 $\le s[j] - \textit{goal}+k-1$ 的数的位置，这可以转化成二分第一个 $\ge s[j] - \textit{goal}+k$ 的数的位置 $p$，那么 $p-1$ 就是最后一个 $\le s[j] - \textit{goal}+k-1$ 的数的位置。

```py [sol-Python3]
# 模板来源 https://leetcode.cn/discuss/post/3583665/
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
        if l > r:
            return 0
        return self.pre(r) - self.pre(l - 1)


class Solution:
    def distantSubarrays(self, nums: list[int], goal: int, k: int) -> int:
        n = len(nums)
        s = list(accumulate(nums, initial=0))
        sorted_s = sorted(set(s))

        ans = n * (n + 1) // 2
        t = FenwickTree(len(sorted_s))
        for v in s:
            # 离散化后的值加一，方便使用树状数组
            l = bisect_right(sorted_s, v - goal - k) + 1
            r = bisect_left(sorted_s, v - goal + k)
            ans -= t.query(l, r)
            t.update(bisect_left(sorted_s, v) + 1, 1)
        return ans
```

```java [sol-Java]
// 模板来源 https://leetcode.cn/discuss/post/3583665/
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
        if (l > r) {
            return 0;
        }
        return pre(r) - pre(l - 1);
    }
}

class Solution {
    public long distantSubarrays(int[] nums, int goal, int k) {
        int n = nums.length;
        long[] sum = new long[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + nums[i];
        }

        long[] sorted = sum.clone();
        Arrays.sort(sorted);

        long ans = (long) n * (n + 1) / 2;
        FenwickTree t = new FenwickTree(n + 1);
        for (long s : sum) {
            // 离散化后的值加一，方便使用树状数组
            int l = lowerBound(sorted, s - goal - k + 1) + 1;
            int r = lowerBound(sorted, s - goal + k);
            ans -= t.query(l, r);
            t.update(lowerBound(sorted, s) + 1, 1);
        }
        return ans;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(long[] a, long target) {
        int left = -1;
        int right = a.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            int mid = (left + right) >>> 1; // 比 /2 快
            if (a[mid] >= target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
// 模板来源 https://leetcode.cn/discuss/post/3583665/
// 根据题目用 FenwickTree<int> t(n) 或者 FenwickTree<long long> t(n) 初始化
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
        T res{};
        for (; i > 0; i &= i - 1) {
            res += tree[i];
        }
        return res;
    }

    // 求区间和 a[l] + ... + a[r]
    // 1 <= l <= r <= n
    // 时间复杂度 O(log n)
    T query(int l, int r) const {
        if (l > r) {
            return 0;
        }
        return pre(r) - pre(l - 1);
    }
};

class Solution {
public:
    long long distantSubarrays(vector<int>& nums, int goal, int k) {
        int n = nums.size();
        vector<long long> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + nums[i];
        }

        auto sorted = sum;
        ranges::sort(sorted);
        sorted.erase(ranges::unique(sorted).begin(), sorted.end());

        long long ans = 1LL * n * (n + 1) / 2;
        FenwickTree<int> t(sorted.size());
        for (auto s : sum) {
            // 离散化后的值加一，方便使用树状数组
            int l = ranges::upper_bound(sorted, s - goal - k) - sorted.begin() + 1;
            int r = ranges::lower_bound(sorted, s - goal + k) - sorted.begin();
            ans -= t.query(l, r);
            t.update(ranges::lower_bound(sorted, s) - sorted.begin() + 1, 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
// 模板来源 https://leetcode.cn/discuss/post/3583665/
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
	if l > r {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func distantSubarrays(nums []int, goal, k int) int64 {
	n := len(nums)
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}

	sorted := slices.Clone(sum)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	ans := n * (n + 1) / 2
	t := newFenwickTree(len(sorted))
	for _, s := range sum {
		// 离散化后的值加一，方便使用树状数组
		l := sort.SearchInts(sorted, s-goal-k+1) + 1
		r := sort.SearchInts(sorted, s-goal+k)
		ans -= t.query(l, r)
		t.update(sort.SearchInts(sorted, s)+1, 1)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 有序集合做法

```py
class Solution:
    def distantSubarrays(self, nums: list[int], goal: int, k: int) -> int:
        n = len(nums)
        s = list(accumulate(nums, initial=0))

        ans = n * (n + 1) // 2
        sl = SortedList()
        for v in s:
            # 统计满足 v-goal-k < x < v-goal+k 的 s[i] 的个数
            l = sl.bisect_right(v - goal - k)
            r = sl.bisect_left(v - goal + k)
            ans -= max(r - l, 0)
            sl.add(v)
        return ans
```

## 其他做法

类似统计**逆序对**的个数，也可以在归并排序 $s$ 的同时，用三指针统计满足 $\textit{goal}-k+1 \le s[j] - s[i] \le \textit{goal}+k-1$ 的 $(i,j)$ 个数。

## 专题训练

见下面数据结构题单的「**§1.2 前缀和与哈希表**」「**§8.1 树状数组**」和「**§8.2 逆序对**」。

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
