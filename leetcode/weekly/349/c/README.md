下午两点[【biIibiIi@灵茶山艾府】](https://b23.tv/JMcHRRp)直播讲题，记得关注哦~

---

### 提示 1

枚举操作次数。

### 提示 2

如果不操作，第 $i$ 个巧克力必须花费 $\textit{nums}[i]$ 收集，总成本为所有 $\textit{nums}[i]$ 之和。

如果操作一次，第 $i$ 个巧克力可以花费 $\min(\textit{nums}[i], \textit{nums}[(i+1)\bmod n])$ 收集。**注意在求和的情况下，把题意理解成循环左移还是循环右移，算出的结果都是一样的。**（样例 1 解释中的类型变更是反过来的，但计算结果是正确的。）

如果操作两次，第 $i$ 个巧克力可以花费 $\min(\textit{nums}[i], \textit{nums}[(i+1)\bmod n],  \textit{nums}[(i+2) \bmod n])$ 收集。

依此类推。

### 提示 3

如果暴力枚举，总的时间复杂度是 $\mathcal{O}(n^3)$。

优化办法有三种：

1. 用 $\mathcal{O}(n^2)$ 的时间预处理所有子数组的最小值，存到一个二维数组中。这样做需要 $\mathcal{O}(n^2)$ 的空间。
2. 用 ST 表优化上述过程。但还有更简单的做法。
3. 用一个长为 $n$ 的数组 $\textit{sum}$ 统计操作 $i$ 次的总花费，这样就可以一边枚举子数组，一边求最小值，一边累加成本了。该方法只需要 $\mathcal{O}(n)$ 的空间。

```py [sol-Python3]
class Solution:
    def minCost(self, nums: List[int], x: int) -> int:
        n = len(nums)
        s = list(range(0, n * x, x))  # s[i] 对应操作 i 次的总成本
        for i, mn in enumerate(nums):  # 子数组左端点
            for j in range(i, n + i):  # 子数组右端点（把数组视作环形的）
                mn = min(mn, nums[j % n])  # 注：手动 if 比大小会快很多
                s[j - i] += mn  # 累加操作 j-i 次的成本
        return min(s)
```

```java [sol-Java]
class Solution {
    public long minCost(int[] nums, int x) {
        int n = nums.length;
        var sum = new long[n];
        for (int i = 0; i < n; i++)
            sum[i] = (long) i * x; // 操作 i 次
        for (int i = 0; i < n; i++) { // 子数组左端点
            int mn = nums[i];
            for (int j = i; j < n + i; j++) { // 子数组右端点（把数组视作环形的）
                mn = Math.min(mn, nums[j % n]); // 从 nums[i] 到 nums[j%n] 的最小值
                sum[j - i] += mn; // 累加操作 j-i 次的成本
            }
        }
        long ans = Long.MAX_VALUE;
        for (var s : sum) ans = Math.min(ans, s);
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(vector<int> &nums, int x) {
        int n = nums.size();
        long long sum[n];
        for (int i = 0; i < n; i++)
            sum[i] = (long long) i * x; // 操作 i 次
        for (int i = 0; i < n; i++) { // 子数组左端点
            int mn = nums[i];
            for (int j = i; j < n + i; j++) { // 子数组右端点（把数组视作环形的）
                mn = min(mn, nums[j % n]); // 从 nums[i] 到 nums[j%n] 的最小值
                sum[j - i] += mn; // 累加操作 j-i 次的成本
            }
        }
        return *min_element(sum, sum + n);
    }
};
```

```go [sol-Go]
func minCost(nums []int, x int) int64 {
	n := len(nums)
	sum := make([]int, n)
	for i := range sum {
		sum[i] = i * x // 操作 i 次
	}
	for i, mn := range nums { // 子数组左端点
		for j := i; j < n+i; j++ { // 子数组右端点（把数组视作环形的）
			mn = min(mn, nums[j%n]) // 从 nums[i] 到 nums[j%n] 的最小值
			sum[j-i] += mn // 累加操作 j-i 次的成本
		}
	}
	ans := math.MaxInt
	for _, s := range sum {
		ans = min(ans, s)
	}
	return int64(ans)
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
