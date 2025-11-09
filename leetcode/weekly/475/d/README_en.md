## Where to Break the Cycle?

Let $M = \textit{nums}[i]$ be the maximum value in $\textit{nums}$. If there are multiple occurrences of the maximum, pick any one of them.

Suppose we already have an optimal partition. Consider the subarray that contains $M$, and let the minimum value in this subarray be at index $j$.

- If $i \le j$, we can assign the elements **to the left of** $i$ to the previous subarray without making the result worse. In this case, $M$ is at the **leftmost** position of its subarray.
- If $i \ge j$, we can assign the elements **to the right of** $i$ to the next subarray without making the result worse. In this case, $M$ is at the **rightmost** position of its subarray.

Therefore, there always exists an optimal partition in which $M$ is either at the leftmost or rightmost end of a subarray.

In other words, we can break the circular array either between $(i-1, i)$ or between $(i, i+1)$, thus converting it into a non-circular problem. The final answer is the maximum result obtained from these two breaking options.

## Non-Circular Array

For each partitioned subarray:

- If the minimum value is **to the left** of the maximum value, it can be viewed as a **buy low, sell high** transaction.
- If the maximum value is **to the left** of the minimum value, it can be viewed as a **short high, cover low** transaction.

The problem restricts the total number of subarrays to at most $k$, which corresponds to at most $k$ transactions. 

This is **exactly the same** as [3573. Best Time to Buy and Sell Stock V](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-v/).

```py [sol-Python3]
fmax = lambda a, b: b if b > a else a

class Solution:
    # 3573. Best Time to Buy and Sell Stock V
    def maximumProfit(self, prices: List[int], k: int) -> int:
        f = [[-inf] * 3 for _ in range(k + 2)]
        for j in range(1, k + 2):
            f[j][0] = 0
        for p in prices:
            for j in range(k + 1, 0, -1):
                f[j][0] = fmax(f[j][0], fmax(f[j][1] + p, f[j][2] - p))
                f[j][1] = fmax(f[j][1], f[j - 1][0] - p)
                f[j][2] = fmax(f[j][2], f[j - 1][0] + p)
        return f[-1][0]

    def maximumScore(self, nums: List[int], k: int) -> int:
        max_i = nums.index(max(nums))
        ans1 = self.maximumProfit(nums[max_i:] + nums[:max_i], k)  # nums[max_i] is the first element.
        ans2 = self.maximumProfit(nums[max_i + 1:] + nums[:max_i + 1], k)  # nums[max_i] is the last element.
        return fmax(ans1, ans2)
```

```java [sol-Java]
class Solution {
    public long maximumScore(int[] nums, int k) {
        int n = nums.length;
        int maxI = 0;
        for (int i = 1; i < n; i++) {
            if (nums[i] > nums[maxI]) {
                maxI = i;
            }
        }

        long ans1 = maximumProfit(nums, maxI, maxI + n, k); // nums[maxI] is the first element.
        long ans2 = maximumProfit(nums, maxI + 1, maxI + 1 + n, k); // nums[maxI] is the last element.
        return Math.max(ans1, ans2);
    }

    // 3573. Best Time to Buy and Sell Stock V
    private long maximumProfit(int[] prices, int l, int r, int k) {
        int n = prices.length;
        long[][] f = new long[k + 2][3];
        for (int j = 1; j <= k + 1; j++) {
            f[j][1] = Long.MIN_VALUE / 2;
        }
        f[0][0] = Long.MIN_VALUE / 2;
        for (int i = l; i < r; i++) {
            int p = prices[i % n];
            for (int j = k + 1; j > 0; j--) {
                f[j][0] = Math.max(f[j][0], Math.max(f[j][1] + p, f[j][2] - p));
                f[j][1] = Math.max(f[j][1], f[j - 1][0] - p);
                f[j][2] = Math.max(f[j][2], f[j - 1][0] + p);
            }
        }
        return f[k + 1][0];
    }
}
```

```cpp [sol-C++]
class Solution {
    // 3573. Best Time to Buy and Sell Stock V
    long long maximumProfit(vector<int>& prices, int l, int r, int k) {
        int n = prices.size();
        vector<array<long long, 3>> f(k + 2, {LLONG_MIN / 2, LLONG_MIN / 2, LLONG_MIN / 2});
        for (int j = 1; j <= k + 1; j++) {
            f[j][0] = 0;
        }
        for (int i = l; i < r; i++) {
            int p = prices[i % n];
            for (int j = k + 1; j > 0; j--) {
                f[j][0] = max(f[j][0], max(f[j][1] + p, f[j][2] - p));
                f[j][1] = max(f[j][1], f[j - 1][0] - p);
                f[j][2] = max(f[j][2], f[j - 1][0] + p);
            }
        }
        return f[k + 1][0];
    }

public:
    long long maximumScore(vector<int>& nums, int k) {
        int n = nums.size();
        int max_i = ranges::max_element(nums) - nums.begin();
        long long ans1 = maximumProfit(nums, max_i, max_i + n, k); // nums[max_i] is the first element.
        long long ans2 = maximumProfit(nums, max_i + 1, max_i + 1 + n, k); // nums[max_i] is the last element.
        return max(ans1, ans2);
    }
};
```

```go [sol-Go]
// 3573. Best Time to Buy and Sell Stock V
func maximumProfit(prices []int, l, r, k int) int64 {
	n := len(prices)
	f := make([][3]int, k+2)
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2
		f[j][2] = math.MinInt / 2
	}
	f[0][0] = math.MinInt / 2
	for i := l; i < r; i++ {
		p := prices[i%n]
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p, f[j][2]-p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
			f[j][2] = max(f[j][2], f[j-1][0]+p)
		}
	}
	return int64(f[k+1][0])
}

func maximumScore(nums []int, k int) int64 {
	n := len(nums)
	maxI := 0
	for i, x := range nums {
		if x > nums[maxI] {
			maxI = i
		}
	}

	ans1 := maximumProfit(nums, maxI, maxI+n, k) // nums[maxI] is the first element.
	ans2 := maximumProfit(nums, maxI+1, maxI+1+n, k) // nums[maxI] is the last element.
	return max(ans1, ans2)
}
```

#### Complexity Analysis

- **Time Complexity:** $\mathcal{O}(nk)$, where $n$ is the length of $\textit{nums}$.
- **Space Complexity:** $\mathcal{O}(k)$.
