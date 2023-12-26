## 提示 1

枚举操作次数，从操作 $0$ 次到操作 $n-1$ 次。

## 提示 2

如果不操作，第 $i$ 个巧克力必须花费 $\textit{nums}[i]$ 收集，总花费为所有 $\textit{nums}[i]$ 之和。例如示例 2 不操作是最优的。

如果只操作一次，第 $i$ 个巧克力可以在操作前购买，也可以在操作后购买，取最小值，即 $\min(\textit{nums}[i], \textit{nums}[(i+1)\bmod n])$。

如果操作两次，购买第 $i$ 个巧克力的花费为 $\min(\textit{nums}[i], \textit{nums}[(i+1)\bmod n], \textit{nums}[(i+2) \bmod n])$。例如示例 1，我们可以操作两次，这样每块巧克力都只需要 $1$ 的花费，总成本为 $2x+1+1+1=13$。

依此类推。

## 提示 3

如果暴力枚举操作次数，再枚举每个巧克力，再计算购买这个巧克力的最小花费，总的时间复杂度是 $\mathcal{O}(n^3)$。

一个初步的优化是，用 $\mathcal{O}(n^2)$ 的时间预处理所有子数组的最小值，保存到一个二维数组中。这样做需要 $\mathcal{O}(n^2)$ 的时间和空间。

但其实不需要预处理，还有更简单的做法：

1. 用一个长为 $n$ 的数组 $s$ 统计不同操作次数下的总成本。
2. 写一个二重循环，枚举子数组的左端点 $i$ 和右端点 $j$。
3. 在枚举右端点的同时，维护从 $\textit{nums}[i]$ 到 $\textit{nums}[j]$ 的最小值 $\textit{mn}$。
4. 把 $\textit{mn}$ 加到 $s[j-i]$ 中，这是因为长为 $j-i+1$ 的子数组恰好对应着操作 $j-i$ 次时要计算的子数组。
5. 最后输出 $\min(s)$。

```py [sol-Python3]
class Solution:
    def minCost(self, nums: List[int], x: int) -> int:
        n = len(nums)
        s = list(range(0, n * x, x))  # s[k] 统计操作 k 次的总成本
        for i, mn in enumerate(nums):  # 子数组左端点
            for j in range(i, n + i):  # 子数组右端点（把数组视作环形的）
                mn = min(mn, nums[j % n])  # 维护从 nums[i] 到 nums[j] 的最小值
                s[j - i] += mn  # 累加操作 j-i 次的花费
        return min(s)
```

```java [sol-Java]
class Solution {
    public long minCost(int[] nums, int x) {
        int n = nums.length;
        long[] s = new long[n]; // s[k] 统计操作 k 次的总成本
        for (int i = 0; i < n; i++) {
            s[i] = (long) i * x;
        }
        for (int i = 0; i < n; i++) { // 子数组左端点
            int mn = nums[i];
            for (int j = i; j < n + i; j++) { // 子数组右端点（把数组视作环形的）
                mn = Math.min(mn, nums[j % n]); // 维护从 nums[i] 到 nums[j] 的最小值
                s[j - i] += mn; // 累加操作 j-i 次的花费
            }
        }
        long ans = Long.MAX_VALUE;
        for (long v : s) {
            ans = Math.min(ans, v);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(vector<int> &nums, int x) {
        int n = nums.size();
        vector<long long> s(n); // s[k] 统计操作 k 次的总成本
        for (int i = 0; i < n; i++) {
            s[i] = (long long) i * x;
        }
        for (int i = 0; i < n; i++) { // 子数组左端点
            int mn = nums[i];
            for (int j = i; j < n + i; j++) { // 子数组右端点（把数组视作环形的）
                mn = min(mn, nums[j % n]); // 维护从 nums[i] 到 nums[j] 的最小值
                s[j - i] += mn; // 累加操作 j-i 次的花费
            }
        }
        return *min_element(s.begin(), s.end());
    }
};
```

```go [sol-Go]
func minCost(nums []int, x int) int64 {
	n := len(nums)
	s := make([]int64, n) // s[k] 统计操作 k 次的总成本
	for i := range s {
		s[i] = int64(i) * int64(x)
	}
	for i, mn := range nums { // 子数组左端点
		for j := i; j < n+i; j++ { // 子数组右端点（把数组视作环形的）
			mn = min(mn, nums[j%n]) // 维护从 nums[i] 到 nums[j] 的最小值
			s[j-i] += int64(mn) // 累加操作 j-i 次的花费
		}
	}
	return slices.Min(s)
}
```

```js [sol-JavaScript]
var minCost = function(nums, x) {
    const n = nums.length;
    const s = Array(n).fill(0); // s[k] 统计操作 k 次的总成本
    for (let i = 0; i < n; i++) {
        s[i] = i * x;
    }
    for (let i = 0; i < n; i++) { // 子数组左端点
        let mn = nums[i];
        for (let j = i; j < n + i; j++) { // 子数组右端点（把数组视作环形的）
            mn = Math.min(mn, nums[j % n]); // 维护从 nums[i] 到 nums[j] 的最小值
            s[j - i] += mn; // 累加操作 j-i 次的花费
        }
    }
    return Math.min(...s);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_cost(nums: Vec<i32>, x: i32) -> i64 {
        let n = nums.len();
        // s[k] 统计操作 k 次的总成本
        let mut s: Vec<i64> = (0..n).map(|i| i as i64 * x as i64).collect();
        for i in 0..n { // 子数组左端点
            let mut mn = nums[i];
            for j in i..(n + i) { // 子数组右端点（把数组视作环形的）
                mn = mn.min(nums[j % n]); // 维护从 nums[i] 到 nums[j] 的最小值
                s[j - i] += mn as i64; // 累加操作 j-i 次的花费
            }
        }
        *s.iter().min().unwrap()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
