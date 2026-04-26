**前置知识**：[树状数组](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/)。

**前置题目**：[376. 摆动序列](https://leetcode.cn/problems/wiggle-subsequence/)。

类似 376 题，定义：

- $\textit{f}_{\textit{inc}}[i]$ 表示以 $\textit{nums}[i]$ 结尾且最后两项**递增**的交替子序列的最大和。
- $\textit{f}_{\textit{dec}}[i]$ 表示以 $\textit{nums}[i]$ 结尾且最后两项**递减**的交替子序列的最大和。

对于 $\textit{f}_{\textit{inc}}[i]$，在 $[0,i-k]$ 中枚举子序列的倒数第二个数 $\textit{nums}[j]$，如果 $\textit{nums}[j] < \textit{nums}[i]$，那么问题变成以 $\textit{nums}[j]$ 结尾且最后两项递减的交替子序列的最大和，即 $\textit{f}_{\textit{dec}}[j]$。取最大值，有

$$
\textit{f}_{\textit{inc}}[i] = \max_{\substack{0\le j\le i-k \\ \textit{nums}[j] < \textit{nums}[i]}} \textit{f}_{\textit{dec}}[j] + \textit{nums}[i]
$$

同理有

$$
\textit{f}_{\textit{dec}}[i] = \max_{\substack{0\le j\le i-k \\ \textit{nums}[j] > \textit{nums}[i]}} \textit{f}_{\textit{inc}}[j] + \textit{nums}[i]
$$

这里要计算的是 $\textit{f}_{\textit{dec}}$ 的**值域前缀最大值**，以及 $\textit{f}_{\textit{inc}}$ 的**值域后缀最大值**，这可以用两个**值域树状数组**分别维护。

对于下标的限制，我们可以在遍历到 $\textit{nums}[i]$ 时，才把 $\textit{f}_{\textit{inc}}[i-k]$ 和 $\textit{f}_{\textit{dec}}[i-k]$ 添加到值域树状数组中，从而保证转移来源的下标 $\le i-k$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Fenwick:
    def __init__(self, n: int):
        self.f = [0] * n

    def update(self, i: int, val: int) -> None:
        f = self.f
        while i < len(f):
            f[i] = max(f[i], val)
            i += i & -i

    def pre_max(self, i: int) -> int:
        f = self.f
        res = 0
        while i > 0:
            res = max(res, f[i])
            i &= i - 1
        return res


class Solution:
    def maxAlternatingSum(self, nums: list[int], k: int) -> int:
        # 离散化 nums
        sorted_nums = sorted(set(nums))

        n = len(nums)
        f_inc = [0] * n  # f_inc[i] 表示以 nums[i] 结尾且最后两项递增的交替子序列的最大和
        f_dec = [0] * n  # f_dec[i] 表示以 nums[i] 结尾且最后两项递减的交替子序列的最大和

        # 值域树状数组
        m = len(sorted_nums)
        inc = Fenwick(m + 1)  # 维护 f_inc[i] 的最大值
        dec = Fenwick(m + 1)  # 维护 f_dec[i] 的最大值

        for i, x in enumerate(nums):
            if i >= k:
                # 在这个时候才把 f_inc[i-k] 和 f_dec[i-k] 添加到值域树状数组中，从而保证转移来源的下标 <= i-k
                j = nums[i - k]
                inc.update(m - j, f_inc[i - k])  # m-j 可以把后缀变成前缀
                dec.update(j + 1, f_dec[i - k])

            j = bisect_left(sorted_nums, x)
            nums[i] = j  # 注意这里修改了 nums[i]，这样上面的 nums[i-k] 无需二分

            f_inc[i] = dec.pre_max(j) + x          # 计算满足 nums[i'] < x 的 f_dec[i'] 的最大值
            f_dec[i] = inc.pre_max(m - 1 - j) + x  # 计算满足 nums[i'] > x 的 f_inc[i'] 的最大值

        return max(max(f_inc), max(f_dec))  # 枚举子序列以 nums[i] 结尾
```

```java [sol-Java]
class Fenwick {
    private final long[] f;

    Fenwick(int n) {
        f = new long[n];
    }

    public void update(int i, long val) {
        for (; i < f.length; i += i & -i) {
            f[i] = Math.max(f[i], val);
        }
    }

    public long preMax(int i) {
        long res = 0;
        for (; i > 0; i &= i - 1) {
            res = Math.max(res, f[i]);
        }
        return res;
    }
}

class Solution {
    public long maxAlternatingSum(int[] nums, int k) {
        int[] sorted = nums.clone();
        Arrays.sort(sorted);

        int n = nums.length;
        long[] fInc = new long[n]; // fInc[i] 表示以 nums[i] 结尾且最后两项递增的交替子序列的最大和
        long[] fDec = new long[n]; // fDec[i] 表示以 nums[i] 结尾且最后两项递减的交替子序列的最大和

        // 值域树状数组
        Fenwick inc = new Fenwick(n + 1); // 维护 fInc[i] 的最大值
        Fenwick dec = new Fenwick(n + 1); // 维护 fDec[i] 的最大值

        long ans = 0;

        for (int i = 0; i < n; i++) {
            int x = nums[i];

            if (i >= k) {
                // 在这个时候才把 fInc[i-k] 和 fDec[i-k] 添加到值域树状数组中，从而保证转移来源的下标 <= i-k
                int j = nums[i - k];
                inc.update(n - j, fInc[i - k]); // n-j 可以把后缀变成前缀
                dec.update(j + 1, fDec[i - k]);
            }

            int j = Arrays.binarySearch(sorted, x);
            nums[i] = j; // 注意这里修改了 nums[i]，这样上面的 nums[i-k] 无需二分

            fInc[i] = dec.preMax(j) + x;           // 计算满足 nums[i'] < x 的 fDec[i'] 的最大值
            fDec[i] = inc.preMax(n - 1 - j) + x; // 计算满足 nums[i'] > x 的 fInc[i'] 的最大值
            ans = Math.max(ans, Math.max(fInc[i], fDec[i])); // 枚举子序列以 nums[i] 结尾
        }

        return ans;
    }
}
```

```cpp [sol-C++]
template<typename T>
class Fenwick {
    vector<T> f;

public:
    Fenwick(int n) : f(n) {}

    void update(int i, T val) {
        for (; i < f.size(); i += i & -i) {
            f[i] = max(f[i], val);
        }
    }

    T pre_max(int i) const {
        T res = 0;
        for (; i > 0; i &= i - 1) {
            res = max(res, f[i]);
        }
        return res;
    }
};

class Solution {
public:
    long long maxAlternatingSum(vector<int>& nums, int k) {
        // 离散化 nums
        auto sorted = nums;
        ranges::sort(sorted);
        sorted.erase(ranges::unique(sorted).begin(), sorted.end());

        int n = nums.size();
        vector<long long> f_inc(n); // f_inc[i] 表示以 nums[i] 结尾且最后两项递增的交替子序列的最大和
        vector<long long> f_dec(n); // f_dec[i] 表示以 nums[i] 结尾且最后两项递减的交替子序列的最大和

        // 值域树状数组
        int m = sorted.size();
        Fenwick<long long> inc(m + 1); // 维护 f_inc[i] 的最大值
        Fenwick<long long> dec(m + 1); // 维护 f_dec[i] 的最大值

        long long ans = 0;

        for (int i = 0; i < n; i++) {
            if (i >= k) {
                // 在这个时候才把 f_inc[i-k] 和 f_dec[i-k] 添加到值域树状数组中，从而保证转移来源的下标 <= i-k
                int j = nums[i - k];
                inc.update(m - j, f_inc[i - k]); // m-j 可以把后缀变成前缀
                dec.update(j + 1, f_dec[i - k]);
            }

            int x = nums[i];
            int j = ranges::lower_bound(sorted, x) - sorted.begin();
            nums[i] = j; // 注意这里修改了 nums[i]，这样上面的 nums[i-k] 无需二分

            f_inc[i] = dec.pre_max(j) + x;         // 计算满足 nums[i'] < x 的 f_dec[i'] 的最大值
            f_dec[i] = inc.pre_max(m - 1 - j) + x; // 计算满足 nums[i'] > x 的 f_inc[i'] 的最大值
            ans = max(ans, max(f_inc[i], f_dec[i])); // 枚举子序列以 nums[i] 结尾
        }

        return ans;
    }
};
```

```go [sol-Go]
type fenwick []int64

func (f fenwick) update(i int, val int64) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

// [1, i] 中的最大值
func (f fenwick) preMax(i int) (res int64) {
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return
}

func maxAlternatingSum(nums []int, k int) (ans int64) {
	// 离散化 nums
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)

	n := len(nums)
	fInc := make([]int64, n) // fInc[i] 表示以 nums[i] 结尾且最后两项递增的交替子序列的最大和
	fDec := make([]int64, n) // fDec[i] 表示以 nums[i] 结尾且最后两项递减的交替子序列的最大和

	// 值域树状数组
	m := len(sorted)
	inc := make(fenwick, m+1) // 维护 fInc[i] 的最大值
	dec := make(fenwick, m+1) // 维护 fDec[i] 的最大值

	for i, x := range nums {
		if i >= k {
			// 在这个时候才把 fInc[i-k] 和 fDec[i-k] 添加到值域树状数组中，从而保证转移来源的下标 <= i-k
			j := nums[i-k]
			inc.update(m-j, fInc[i-k]) // m-j 可以把后缀变成前缀
			dec.update(j+1, fDec[i-k])
		}

		j := sort.SearchInts(sorted, x)
		nums[i] = j // 注意这里修改了 nums[i]，这样上面的 nums[i-k] 无需二分

		fInc[i] = dec.preMax(j) + int64(x)     // 计算满足 nums[i'] < x 的 fDec[i'] 的最大值
		fDec[i] = inc.preMax(m-1-j) + int64(x) // 计算满足 nums[i'] > x 的 fInc[i'] 的最大值
		ans = max(ans, fInc[i], fDec[i])       // 枚举子序列以 nums[i] 结尾
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**§11.4 树状数组/线段树优化 DP**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
