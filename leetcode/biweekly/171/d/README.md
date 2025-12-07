## 前置知识

1. [【套路】教你解决定长滑窗！适用于所有定长滑窗题目！](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)
2. [带你发明树状数组！附数学证明](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/)

## 思路

滑窗的同时，实时维护窗口内的逆序对个数 $\textit{inv}$。

- 元素 $x$ 进入窗口时，$\textit{inv}$ 增加了窗口内的大于 $x$ 的元素个数。
- 元素 $x$ 离开窗口时，$\textit{inv}$ 减少了窗口内的小于 $x$ 的元素个数。

这可以用**值域树状数组**动态维护。维护什么？维护元素的**出现次数的前缀和**。

由于元素范围很大，我们又只需要知道元素的相对大小（元素的绝对大小不重要）。所以可以先离散化，把元素映射到 $n$ 以内，例如 $300,300,100,800$ 离散化后就是 $2,2,1,3$，保留了元素的相对大小。注：这里从 $1$ 开始，是因为我的树状数组模板是从 $1$ 开始的。

**小优化**：如果循环中发现 $\textit{ans}=0$，那么已经达到最小值，直接跳出循环。

> 另见「Python3 有序集合」写法。

[本题视频讲解](https://www.bilibili.com/video/BV1wr2fBpENB/?t=29m48s)，欢迎点赞关注~

```py [sol-Python3]
# 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
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


class Solution:
    def minInversionCount(self, nums: List[int], k: int) -> int:
        # 离散化
        sorted_nums = sorted(set(nums))
        mp = {x: i for i, x in enumerate(sorted_nums, 1)}  # 树状数组下标从 1 开始
        for i, x in enumerate(nums):
            nums[i] = mp[x]

        t = FenwickTree(len(sorted_nums))
        inv = 0
        ans = inf

        for i, x in enumerate(nums):
            # 1. 入
            t.update(x, 1)
            inv += min(i + 1, k) - t.pre(x)  # 窗口大小 - (<=x 的元素个数) = (>x 的元素个数)

            left = i + 1 - k
            if left < 0:  # 尚未形成第一个窗口
                continue

            # 2. 更新答案
            ans = min(ans, inv)
            if ans == 0:  # 已经最小了，无需再计算
                break

            # 3. 出
            out = nums[left]
            inv -= t.pre(out - 1)  # < out 的元素个数
            t.update(out, -1)

        return ans
```

```py [sol-Python3 有序集合]
class Solution:
    def minInversionCount(self, nums: List[int], k: int) -> int:
        sl = SortedList()
        inv = 0
        ans = inf

        for i, x in enumerate(nums):
            # 1. 入
            sl.add(x)
            inv += len(sl) - sl.bisect_right(x)  # 窗口大小 - (<=x 的元素个数) = (>x 的元素个数)

            left = i + 1 - k
            if left < 0:  # 尚未形成第一个窗口
                continue

            # 2. 更新答案
            ans = min(ans, inv)
            if ans == 0:  # 已经最小了，无需再计算
                break

            # 3. 出
            out = nums[left]
            inv -= sl.bisect_left(out)  # < out 的元素个数
            sl.discard(out)

        return ans
```

```java [sol-Java]
// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
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
}

class Solution {
    public long minInversionCount(int[] nums, int k) {
        // 离散化
        int n = nums.length;
        int[] sorted = nums.clone();
        Arrays.sort(sorted);
        for (int i = 0; i < n; i++) {
            nums[i] = Arrays.binarySearch(sorted, nums[i]) + 1; // 树状数组下标从 1 开始
        }

        FenwickTree t = new FenwickTree(sorted.length);
        long inv = 0;
        long ans = Long.MAX_VALUE;

        for (int i = 0; i < n; i++) {
            // 1. 入
            int in = nums[i];
            t.update(in, 1);
            inv += Math.min(i + 1, k) - t.pre(in); // 窗口大小 - (<=x 的元素个数) = (>x 的元素个数)

            int left = i + 1 - k;
            if (left < 0) { // 尚未形成第一个窗口
                continue;
            }

            // 2. 更新答案
            ans = Math.min(ans, inv);
            if (ans == 0) { // 已经最小了，无需再计算
                break;
            }

            // 3. 出
            int out = nums[left];
            inv -= t.pre(out - 1); // < out 的元素个数
            t.update(out, -1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
// 完整模板见 https://leetcode.cn/circle/discuss/mOr1u6/
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
};

class Solution {
public:
    long long minInversionCount(vector<int>& nums, int k) {
        // 离散化
        vector<int> sorted = nums;
        ranges::sort(sorted);
        sorted.erase(ranges::unique(sorted).begin(), sorted.end());
        for (int& x : nums) {
            x = ranges::lower_bound(sorted, x) - sorted.begin() + 1; // 树状数组下标从 1 开始
        }

        FenwickTree<int> t(sorted.size());
        long long ans = LLONG_MAX, inv = 0;

        for (int i = 0; i < nums.size(); i++) {
            // 1. 入
            int in = nums[i];
            t.update(in, 1);
            inv += min(i + 1, k) - t.pre(in); // 窗口大小 - (<=x 的元素个数) = (>x 的元素个数)

            int left = i + 1 - k;
            if (left < 0) { // 尚未形成第一个窗口
                continue;
            }

            // 2. 更新答案
            ans = min(ans, inv);
            if (ans == 0) { // 已经最小了，无需再计算
                break;
            }

            // 3. 出
            int out = nums[left];
            inv -= t.pre(out - 1); // < out 的元素个数
            t.update(out, -1);
        }
        return ans;
    }
};
```

```go [sol-Go]
type fenwick []int

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func minInversionCount(nums []int, k int) int64 {
	// 离散化
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	for i, x := range nums {
		nums[i] = sort.SearchInts(sorted, x) + 1 // 树状数组下标从 1 开始
	}

	t := make(fenwick, len(sorted)+1)
	inv := 0
	ans := math.MaxInt

	for i, in := range nums {
		// 1. 入
		t.update(in, 1)
		inv += min(i+1, k) - t.pre(in) // 窗口大小 - (<=x 的元素个数) = (>x 的元素个数)

		left := i + 1 - k
		if left < 0 { // 尚未形成第一个窗口
			continue
		}

		// 2. 更新答案
		ans = min(ans, inv)
		if ans == 0 { // 已经最小了，无需再计算
			break
		}

		// 3. 出
		out := nums[left]
		inv -= t.pre(out - 1) // < out 的元素个数
		t.update(out, -1)
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§8.2 逆序对**」。

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
