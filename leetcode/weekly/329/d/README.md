下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

### 如何思考

划分出**第一个**子数组，问题变成一个规模更小的子问题。

由于「划分出长为 $x$ 和 $y$ 的子数组」和「划分出长为 $y$ 和 $x$ 的子数组」之后，剩余的子问题是相同的，因此这题适合用动态规划解决。

### 具体算法

定义 $f[i+1]$ 表示划分 $\textit{nums}$ 的前 $i$ 个数的最小代价，从 $i$ 开始倒序枚举最后一个子数组的开始位置 $j$，同时用一个数组 $\textit{cnt}$ 维护每个元素的出现次数，用一个变量 $\textit{unique}$ 维护只出现一次的元素个数。

如果一个元素首次遇到，那么 $\textit{unique}$ 加一，如果第二次遇到，那么 $\textit{unique}$ 减一。

重要性为子数组的长度减去只出现一次的元素个数加 $k$，即

$$
i-j+1 - \textit{unique}_j + k
$$

这里 $\textit{unique}_j$ 表示枚举到 $j$ 时的 $\textit{unique}$ 值。

加上前面子数组的最小代价，所有结果取最小值，得

$$
\begin{aligned}
f[i+1] &= \min\limits_{j=0}^{i} f[j] + i-j+1 - \textit{unique}_j + k \\ 
       &= i+1+k+ \min\limits_{j=0}^{i} f[j] -j - \textit{unique}_j
\end{aligned}
$$

初始值 $f[0] = 0$，答案为 $f[n]$。

### 优化

注意到 $f[j]$ 每次都要减去 $j$，而 $f[i+1]$ 最后还要加上 $i+1$，如果定义 $f'[i] = f[i]-i$，则有

$$
f'[i+1] = k+\min\limits_{j=0}^{i} f'[j] - \textit{unique}_j
$$

由于 $f'[n] = f[n]-n$，所以最后答案为 $f'[n]+n$。

```py [sol1-Python3]
class Solution:
    def minCost(self, nums: List[int], k: int) -> int:
        n = len(nums)
        f = [0] * (n + 1)
        for i in range(n):
            cnt, unique, mn = [0] * n, 0, inf
            for j in range(i, -1, -1):
                x = nums[j]
                cnt[x] += 1
                if cnt[x] == 1: unique += 1
                elif cnt[x] == 2: unique -= 1
                mn = min(mn, f[j] - unique)
                # if f[j]-unique < mn: mn = f[j]-unique  # 手写 min 会快很多
            f[i + 1] = k + mn
        return f[n] + n
```

```java [sol1-Java]
class Solution {
    public int minCost(int[] nums, int k) {
        int n = nums.length;
        int[] f = new int[n + 1], cnt = new int[n];
        for (int i = 0; i < n; ++i) {
            Arrays.fill(cnt, 0);
            int unique = 0, mn = Integer.MAX_VALUE;
            for (int j = i; j >= 0; --j) {
                int x = nums[j];
                if (++cnt[x] == 1) ++unique; // 首次出现
                else if (cnt[x] == 2) --unique; // 不再唯一
                mn = Math.min(mn, f[j] - unique);
            }
            f[i + 1] = k + mn;
        }
        return f[n] + n;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minCost(vector<int> &nums, int k) {
        int n = nums.size(), f[n + 1], cnt[n];
        f[0] = 0;
        for (int i = 0; i < n; ++i) {
            memset(cnt, 0, sizeof(cnt));
            int unique = 0, mn = INT_MAX;
            for (int j = i; j >= 0; --j) {
                int x = nums[j];
                if (++cnt[x] == 1) ++unique; // 首次出现
                else if (cnt[x] == 2) --unique; // 不再唯一
                mn = min(mn, f[j] - unique);
            }
            f[i + 1] = k + mn;
        }
        return f[n] + n;
    }
};
```

```go [sol1-Go]
func minCost(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		cnt, unique, mn := make([]int, n), 0, math.MaxInt
		for j := i; j >= 0; j-- {
			x := nums[j]
			cnt[x]++
			if cnt[x] == 1 { // 首次出现
				unique++
			} else if cnt[x] == 2 { // 不再唯一
				unique--
			}
			mn = min(mn, f[j]-unique)
		}
		f[i+1] = k + mn
	}
	return f[n] + n
}

func min(a, b int) int { if a > b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

### 思考题

你能想出复杂度更低的算法吗？
