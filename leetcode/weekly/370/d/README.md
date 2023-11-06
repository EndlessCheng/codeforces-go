请看 [视频讲解](https://www.bilibili.com/video/BV1Fc411R7xA/) 第四题。

定义 $b[i] = \textit{nums}[i] - i$，问题变成从 $b$ 中选出一个非降子序列，求对应的 $\textit{nums}$ 的元素和的最大值。

如果 $i$ 是子序列最后一个数的下标，考虑倒数第二个数的下标 $j$，如果 $b[j]\le b[i]$，那么就找到了一个子问题：子序列最后一个数的下标是 $j$ 时，对应的 $\textit{nums}$ 的元素和的最大值。

据此，定义 $f[i]$ 表示子序列最后一个数的下标是 $i$ 时，对应的 $\textit{nums}$ 的元素和的最大值。那么答案就是 $\max(f)$。

状态转移方程为

$$
f[i] = \max_{j} \max(f[j],0) + \textit{nums}[i] 
$$

其中 $j$ 满足 $j < i$ 且 $b[j]\le b[i]$。如果 $f[j]<0$，则和 $0$ 取最大值，表示只选 $\textit{nums}[i]$ 一个数，前面的数都不选。

这可以用**权值树状数组**（或者权值线段树）来优化。树状数组用来维护前缀最大值：设下标为 $x=b[i]$，维护的值为 $\max(f[x], f[x-1], f[x-2], \cdots)$。具体请看视频讲解。

代码实现时需要先把 $\textit{nums}[i]-i$ **离散化**，再使用树状数组。

> 离散化：把最小元素映射为 $1$，次小元素映射为 $2$，依此类推。

```py [sol-Python3]
class Solution:
    def maxBalancedSubsequenceSum(self, nums: List[int]) -> int:
        b = sorted(set(x - i for i, x in enumerate(nums)))  # 离散化 nums[i]-i
        t = BIT(len(b) + 1)
        for i, x in enumerate(nums):
            j = bisect_left(b, x - i) + 1  # nums[i]-i 离散化后的值（从 1 开始）
            f = max(t.pre_max(j), 0) + x
            t.update(j, f)
        return t.pre_max(len(b))  # 所有 f 的最大值

# 树状数组模板（维护前缀最大值）
class BIT:
    def __init__(self, n: int):
        self.tree = [-inf] * n

    def update(self, i: int, val: int) -> None:
        while i < len(self.tree):
            self.tree[i] = max(self.tree[i], val)
            i += i & -i

    def pre_max(self, i: int) -> int:
        mx = -inf
        while i > 0:
            mx = max(mx, self.tree[i])
            i &= i - 1
        return mx
```

```java [sol-Java]
class Solution {
    public long maxBalancedSubsequenceSum(int[] nums) {
        int n = nums.length;
        int[] b = new int[n];
        for (int i = 0; i < n; i++) {
            b[i] = nums[i] - i;
        }
        Arrays.sort(b);
 
        BIT t = new BIT(b.length + 1);
        for (int i = 0; i < n; i++) {
            // j 为 nums[i]-i 离散化后的值（从 1 开始）
            int j = Arrays.binarySearch(b, nums[i] - i) + 1;
            long f = Math.max(t.preMax(j), 0) + nums[i];
            t.update(j, f);
        }
        return t.preMax(b.length);
    }
}

// 树状数组模板（维护前缀最大值）
class BIT {
    private long[] tree;

    public BIT(int n) {
        tree = new long[n];
        Arrays.fill(tree, Long.MIN_VALUE);
    }

    public void update(int i, long val) {
        while (i < tree.length) {
            tree[i] = Math.max(tree[i], val);
            i += i & -i;
        }
    }

    public long preMax(int i) {
        long res = Long.MIN_VALUE;
        while (i > 0) {
            res = Math.max(res, tree[i]);
            i &= i - 1;
        }
        return res;
    }
}
```

```cpp [sol-C++]
// 树状数组模板（维护前缀最大值）
class BIT {
    vector<long long> tree;
public:
    BIT(int n) : tree(n, LLONG_MIN) {}

    void update(int i, long long val) {
        while (i < tree.size()) {
            tree[i] = max(tree[i], val);
            i += i & -i;
        }
    }

    long long pre_max(int i) {
        long long res = LLONG_MIN;
        while (i > 0) {
            res = max(res, tree[i]);
            i &= i - 1;
        }
        return res;
    }
};

class Solution {
public:
    long long maxBalancedSubsequenceSum(vector<int> &nums) {
        int n = nums.size();
        // 离散化 nums[i]-i
        auto b = nums;
        for (int i = 0; i < n; i++) {
            b[i] -= i;
        }
        sort(b.begin(), b.end());
        b.erase(unique(b.begin(), b.end()), b.end()); // 去重

        BIT t = BIT(b.size() + 1);
        for (int i = 0; i < n; i++) {
            // j 为 nums[i]-i 离散化后的值（从 1 开始）
            int j = lower_bound(b.begin(), b.end(), nums[i] - i) - b.begin() + 1;
            long long f = max(t.pre_max(j), 0LL) + nums[i];
            t.update(j, f);
        }
        return t.pre_max(b.size());
    }
};
```

```go [sol-Go]
// 树状数组模板（维护前缀最大值）
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) int {
	mx := math.MinInt
	for ; i > 0; i &= i - 1 {
		mx = max(mx, f[i])
	}
	return mx
}

func maxBalancedSubsequenceSum(nums []int) int64 {
	// 离散化 nums[i]-i
	b := slices.Clone(nums)
	for i := range b {
		b[i] -= i
	}
	slices.Sort(b)
	b = slices.Compact(b) // 去重

	// 初始化树状数组
	t := make(fenwick, len(b)+1)
	for i := range t {
		t[i] = math.MinInt
	}

	for i, x := range nums {
		j := sort.SearchInts(b, x-i) + 1 // nums[i]-i 离散化后的值（从 1 开始）
		f := max(t.preMax(j), 0) + x
		t.update(j, f)
	}
	return int64(t.preMax(len(b)))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

#### 相似题目

- [2407. 最长递增子序列 II](https://leetcode.cn/problems/longest-increasing-subsequence-ii/)
