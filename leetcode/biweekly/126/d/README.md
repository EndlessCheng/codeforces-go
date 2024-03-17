请看 [视频讲解](https://www.bilibili.com/video/BV11x421r7q5/) 第四题。

## 方法一：二维 0-1 背包

#### 观察

如果和为 $k$ 的子序列 $S$ 的长度是 $c$，那么有多少个子序列 $T$，会包含 $S$ 呢？即 $S$ 是 $T$ 的子序列。

例如示例 1，子序列 $S=[3]$，出现在子序列 $[1,2,3],[1,3],[2,3],[3]$ 中，这 $4$ 个子序列都可以是 $T$。

除了 $3$ 以外的每个数，都可以选择在/不在包含 $[3]$ 的子序列 $T$ 中。

所以有 $2^{n-c}$ 个子序列 $T$。

这意味着 $S$ 对答案的贡献是 $2^{n-c}$。

#### 思路

枚举和为 $k$ 的子序列的长度 $c$，问题变成：

- 有 $n$ 个物品，每个物品的体积是 $\textit{nums}[i]$，问恰好装满容量为 $k$ 的背包，且选择 $c$ 个物品的方案数。

这可以用二维 0-1 背包解决。不了解 0-1 背包的同学请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

定义 $f[i][j][c]$ 表示考虑前 $i$ 个物品，所选物品体积和是 $j$，选了 $c$ 个物品的方案数。

考虑第 $i$ 个物品选或不选：

- 不选：$f[i+1][j][c] = f[i][j][c]$。
- 选：$f[i+1][j][c] = f[i][j-\textit{nums}[i]][c-1]$。

两者相加，得

$$
f[i+1][j][c] = f[i][j][c] + f[i][j-\textit{nums}[i]][c-1]
$$

初始值：$f[0][0][0] = 1$。

答案：$\sum\limits_{c=1}^{n} f[n][k][c] \cdot 2^{n-c}$。

代码实现时，第一个维度可以优化掉。

```py [sol-Python3]
class Solution:
    def sumOfPower(self, nums: List[int], k: int) -> int:
        MOD = 1_000_000_007
        n = len(nums)
        f = [[0] * (n + 1) for _ in range(k + 1)]
        f[0][0] = 1
        for i, x in enumerate(nums):
            for j in range(k, x - 1, -1):
                for c in range(i + 1, 0, -1):
                    f[j][c] = (f[j][c] + f[j - x][c - 1]) % MOD

        ans = 0
        pow2 = 1
        for i in range(n, 0, -1):
            ans = (ans + f[k][i] * pow2) % MOD
            pow2 = pow2 * 2 % MOD
        return ans
```

```java [sol-Java]
class Solution {
    public int sumOfPower(int[] nums, int k) {
        final int MOD = 1_000_000_007;
        int n = nums.length;
        int[][] f = new int[k + 1][n + 1];
        f[0][0] = 1;
        for (int i = 0; i < n; i++) {
            for (int j = k; j >= nums[i]; j--) {
                for (int c = i + 1; c > 0; c--) {
                    f[j][c] = (f[j][c] + f[j - nums[i]][c - 1]) % MOD;
                }
            }
        }

        long ans = 0;
        int pow2 = 1;
        for (int i = n; i > 0; i--) {
            ans = (ans + (long) f[k][i] * pow2) % MOD;
            pow2 = pow2 * 2 % MOD;
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumOfPower(vector<int> &nums, int k) {
        const int MOD = 1'000'000'007;
        int n = nums.size();
        vector<vector<int>> f(k + 1, vector<int>(n + 1));
        f[0][0] = 1;
        for (int i = 0; i < n; i++) {
            for (int j = k; j >= nums[i]; j--) {
                for (int c = i + 1; c > 0; c--) {
                    f[j][c] = (f[j][c] + f[j - nums[i]][c - 1]) % MOD;
                }
            }
        }

        int ans = 0;
        int pow2 = 1;
        for (int i = n; i > 0; i--) {
            ans = (ans + (long long) f[k][i] * pow2) % MOD;
            pow2 = pow2 * 2 % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumOfPower(nums []int, k int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 1
	for i, x := range nums {
		for j := k; j >= x; j-- {
			for c := i + 1; c > 0; c-- {
				f[j][c] = (f[j][c] + f[j-x][c-1]) % mod
			}
		}
	}
	pow2 := 1
	for i := n; i > 0; i-- {
		ans = (ans + f[k][i]*pow2) % mod
		pow2 = pow2 * 2 % mod
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2k)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。

## 方法二：一维 0-1 背包

定义 $f[i][j]$ 表示考虑前 $i$ 个数，所选元素和是 $j$ 时的能量和。

对于 $\textit{nums}[i]$，它只有三种可能：

- 不在子序列 $S$，也不在子序列 $T$ 中：$f[i+1][j] = f[i][j]$。
- 不在子序列 $S$，在子序列 $T$ 中：$f[i+1][j] = f[i][j]$。
- 在子序列 $S$ 中：$f[i+1][j] = f[i][j-\textit{nums}[i]]$。

三者相加，得

$$
f[i+1][j] = f[i][j]\cdot 2 + f[i][j-\textit{nums}[i]]
$$

初始值：$f[0][0] = 1$。

答案：$f[n][k]$。

代码实现时，第一个维度可以优化掉。

```py [sol-Python3]
class Solution:
    def sumOfPower(self, nums: List[int], k: int) -> int:
        f = [1] + [0] * k
        for x in nums:
            for j in range(k, -1, -1):
                f[j] = (f[j] * 2 + (f[j - x] if j >= x else 0)) % 1_000_000_007
        return f[k]
```

```java [sol-Java]
class Solution {
    public int sumOfPower(int[] nums, int k) {
        long[] f = new long[k + 1];
        f[0] = 1;
        for (int x : nums) {
            for (int j = k; j >= 0; j--) {
                f[j] = (f[j] * 2 + (j >= x ? f[j - x] : 0)) % 1_000_000_007;
            }
        }
        return (int) f[k];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumOfPower(vector<int> &nums, int k) {
        vector<long long> f(k + 1);
        f[0] = 1;
        for (int x : nums) {
            for (int j = k; j >= 0; j--) {
                f[j] = (f[j] * 2 + (j >= x ? f[j - x] : 0)) % 1'000'000'007;
            }
        }
        return f[k];
    }
};
```

```go [sol-Go]
func sumOfPower(nums []int, k int) int {
	const mod = 1_000_000_007
	f := make([]int, k+1)
	f[0] = 1
	for _, x := range nums {
		for j := k; j >= 0; j-- {
			if j >= x {
				f[j] = (f[j]*2 + f[j-x]) % mod
			} else {
				f[j] = f[j] * 2 % mod
			}
		}
	}
	return f[k]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 题单：0-1 背包

- [2915. 和为目标值的最长子序列的长度](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/) 1659
- [416. 分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum/)
- [494. 目标和](https://leetcode.cn/problems/target-sum/)
- [2787. 将一个数字表示成幂的和的方案数](https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/) 1818
- [474. 一和零](https://leetcode.cn/problems/ones-and-zeroes/)（二维）
- [1049. 最后一块石头的重量 II](https://leetcode.cn/problems/last-stone-weight-ii/) 2092
- [879. 盈利计划](https://leetcode.cn/problems/profitable-schemes/) 2204
- [956. 最高的广告牌](https://leetcode.cn/problems/tallest-billboard/) 2381
- [2518. 好分区的数目](https://leetcode.cn/problems/number-of-great-partitions/) 2415
- [2742. 给墙壁刷油漆](https://leetcode.cn/problems/painting-the-walls/) 2425
- [2291. 最大股票收益](https://leetcode.cn/problems/maximum-profit-from-trading-stocks/)（会员题）
- [2431. 最大限度地提高购买水果的口味](https://leetcode.cn/problems/maximize-total-tastiness-of-purchased-fruits/)（会员题）

更多题单，请点我个人主页 - 讨论发布。
