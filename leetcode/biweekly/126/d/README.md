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

定义 $f[i+1][j]$ 表示在 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 中选子序列，元素和恰好等于 $j$ 的子序列个数。

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

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
