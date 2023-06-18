下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，不仅讲做法，还会教你如何一步步思考，记得关注哦~

---

## 前置知识：模运算

如果让你计算 $1234\cdot 6789$ 的**个位数**，你会如何计算？

由于只有个位数会影响到乘积的个位数，那么 $4\cdot 9=36$ 的个位数 $6$ 就是答案。

对于 $1234+6789$ 的个位数，同理，$4+9=13$ 的个位数 $3$ 就是答案。

你能把这个结论抽象成数学等式吗？

一般地，涉及到取模的题目，通常会用到如下等式（上面计算的是 $m=10$）：

$$
(a+b)\bmod m = ((a\bmod m) + (b\bmod m)) \bmod m
$$

$$
(a\cdot b) \bmod m=((a\bmod m)\cdot  (b\bmod m)) \bmod m
$$

证明：根据**带余除法**，任意整数 $a$ 都可以表示为 $a=km+r$，这里 $r$ 相当于 $a\bmod m$。那么设 $a=k_1m+r_1,\ b=k_2m+r_2$。

第一个等式：

$$
\begin{aligned}
&\ (a+b) \bmod m\\
=&\ ((k_1+k_2) m+r_1+r_2)\bmod m\\
=&\ (r_1+r_2)\bmod m\\
=&\ ((a\bmod m) + (b\bmod m)) \bmod m
\end{aligned}
$$

第二个等式：

$$
\begin{aligned}
&\ (a\cdot b) \bmod m\\
=&\ (k_1k_2m^2+(k_1r_2+k_2r_1)m+r_1r_2)\bmod m\\
=&\ (r_1r_2)\bmod m\\
=&\ ((a\bmod m)\cdot  (b\bmod m)) \bmod m
\end{aligned}
$$

**根据这两个恒等式，可以随意地对代码中的加法和乘法的结果取模**。

## 前置知识：位运算

详见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

## 前置知识：动态规划入门

详见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)

## 思路

仿照 [排列型回溯](https://www.bilibili.com/video/BV1mY411D7f6/) 的定义方式，我们需要知道当前还有哪些数（下标）可以选，以及上一个选的数（下标）是多少。

定义 $\textit{dfs}(i,j)$ 表示当前可以选的下标集合为 $i$，上一个选的数的下标是 $j$ 时，可以构造出多少个特别排列。

枚举当前要选的数的下标 $k$，如果 $\textit{nums}[k]$ 与 $\textit{nums}[j]$ 满足题目整除的要求，则

$$
\textit{dfs}(i,j) = \sum_{k\in i} \textit{dfs}(i\setminus \{k\},k)
$$

递归边界：$\textit{dfs}(0,j) = 1$，表示找到了一个特别排列。

递归入口：$\textit{dfs}(U\setminus \{j\},j)$，其中全集 $U=\{0,1,2,\cdots,n-1\}$。枚举特别排列的第一个数的下标 $j$，累加所有 $\textit{dfs}(U\setminus \{j\},j)$，即为答案。

```py [sol-Python3]
class Solution:
    def specialPerm(self, nums: List[int]) -> int:
        MOD = 10 ** 9 + 7
        @cache
        def dfs(i: int, j: int) -> int:
            if i == 0: return 1  # 找到一个特别排列
            res = 0
            for k, x in enumerate(nums):
                if i >> k & 1 and (nums[j] % x == 0 or x % nums[j] == 0):
                    res += dfs(i ^ (1 << k), k)
            return res
        n = len(nums)
        return sum(dfs(((1 << n) - 1) ^ (1 << j), j) for j in range(n)) % MOD
```

```go [sol-Go]
func specialPerm(nums []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(nums)
	m := 1 << n
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1 // 找到一个特别排列
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		for k, x := range nums {
			if i>>k&1 > 0 && (nums[j]%x == 0 || x%nums[j] == 0) {
				res = (res + dfs(i^(1<<k), k)) % mod
			}
		}
		*p = res
		return
	}
	for j := range nums {
		ans = (ans + dfs((m-1)^(1<<j), j)) % mod
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^22^n)$，其中 $n$ 为 $\textit{nums}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n2^n)$，单个状态的计算时间为 $\mathcal{O}(n)$，因此时间复杂度为 $\mathcal{O}(n^22^n)$。
- 空间复杂度：$\mathcal{O}(n2^n)$。

## 1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

做法：

- $\textit{dfs}$ 改成 $f$ 数组；
- 递归改成循环（每个参数都对应一层循环）；
- 递归边界改成 $f$ 数组的初始值。

具体来说，$f[i][j]$ 的含义和状态转移方程 $\textit{dfs}(i,j)$ 是一样的，即

$$
f[i][j] =\sum_{k\in i} f[i\setminus \{k\}][k]
$$

初始值 $f[0][j]=1$。（翻译自 $\textit{dfs}(0,j)=1$。）

答案为 $f[U\setminus \{j\}][j]$ 之和。（翻译自 $\textit{dfs}(U\setminus \{j\},j)$。）

```py [sol-Python3]
class Solution:
    def specialPerm(self, nums: List[int]) -> int:
        MOD = 10 ** 9 + 7
        n = len(nums)
        m = 1 << n
        f = [[0] * n for _ in range(m)]
        f[0] = [1] * n
        for i in range(1, m):
            for k, x in enumerate(nums):
                if i >> k & 1 == 0: continue
                for j, y in enumerate(nums):
                    if x % y == 0 or y % x == 0:
                        f[i][j] += f[i ^ (1 << k)][k]
        return sum(f[(m - 1) ^ (1 << j)][j] for j in range(n)) % MOD
```

```go [sol-Go]
func specialPerm(nums []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(nums)
	m := 1 << n
	f := make([][]int, m)
	f[0] = make([]int, n)
	for j := range f[0] {
		f[0][j] = 1
	}
	for i := 1; i < m; i++ {
		f[i] = make([]int, n)
		for j, x := range nums {
			for k, y := range nums {
				if i>>k&1 > 0 && (x%y == 0 || y%x == 0) {
					f[i][j] = (f[i][j] + f[i^(1<<k)][k]) % mod
				}
			}
		}
	}
	for j := range nums {
		ans = (ans + f[(m-1)^(1<<j)][j]) % mod
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^22^n)$，其中 $n$ 为 $\textit{nums}$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n2^n)$，单个状态的计算时间为 $\mathcal{O}(n)$，因此时间复杂度为 $\mathcal{O}(n^22^n)$。
- 空间复杂度：$\mathcal{O}(n2^n)$。

## 状压 DP 题单

- [996. 正方形数组的数目](https://leetcode.cn/problems/number-of-squareful-arrays/)，和本题没啥区别
- [2172. 数组的最大与和](https://leetcode.cn/problems/maximum-and-sum-of-array/)，[题解](https://leetcode.cn/problems/maximum-and-sum-of-array/solution/zhuang-tai-ya-suo-dp-by-endlesscheng-5eqn/)
- [1125. 最小的必要团队](https://leetcode.cn/problems/smallest-sufficient-team/)，[题解](https://leetcode.cn/problems/smallest-sufficient-team/solution/zhuang-ya-0-1-bei-bao-cha-biao-fa-vs-shu-qode/)
- [2305. 公平分发饼干](https://leetcode.cn/problems/fair-distribution-of-cookies/)，[题解](https://leetcode.cn/problems/fair-distribution-of-cookies/solution/by-endlesscheng-80ao/)
- [1494. 并行课程 II](https://leetcode.cn/problems/parallel-courses-ii/)，[题解](https://leetcode.cn/problems/parallel-courses-ii/solution/zi-ji-zhuang-ya-dpcong-ji-yi-hua-sou-suo-oxwd/)
- [LCP 53. 守护太空城](https://leetcode.cn/problems/EJvmW4/)，[题解](https://leetcode.cn/problems/EJvmW4/solution/by-endlesscheng-pk2q/)
- [1879. 两个数组最小的异或值之和](https://leetcode.cn/problems/minimum-xor-sum-of-two-arrays/)
- [1986. 完成任务的最少工作时间段](https://leetcode.cn/problems/minimum-number-of-work-sessions-to-finish-the-tasks/)
