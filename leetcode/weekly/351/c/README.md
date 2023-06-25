根据题意，需要在**每**两个 $1$ 之间画一条分割线，有 $x$ 个 $0$ 就可以画 $x+1$ 条分割线。

根据乘法原理，答案为所有分割线的方案数的乘积。这里讲解见[【周赛 351】](https://www.bilibili.com/video/BV1du41187ZN/)第三题。

特别地，如果数组中没有 $1$，那么答案为 $0$。如果数组只有一个 $1$，那么答案为 $1$。

如果你对取模有疑问，可以看文末的讲解。

```py [sol-Python3]
class Solution:
    def numberOfGoodSubarraySplits(self, nums: List[int]) -> int:
        MOD = 10 ** 9 + 7
        ans, pre = 1, -1
        for i, x in enumerate(nums):
            if x == 0: continue
            if pre >= 0:
                ans = ans * (i - pre) % MOD
            pre = i
        return 0 if pre < 0 else ans
```

```java [sol-Java]
class Solution {
    public int numberOfGoodSubarraySplits(int[] nums) {
        final long MOD = (long) 1e9 + 7;
        long ans = 1;
        int pre = -1, n = nums.length;
        for (int i = 0; i < n; i++) {
            if (nums[i] == 0) continue;
            if (pre >= 0) ans = ans * (i - pre) % MOD;
            pre = i;
        }
        return pre < 0 ? 0 : (int) ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfGoodSubarraySplits(vector<int> &nums) {
        const int MOD = 1e9 + 7;
        int ans = 1, pre = -1, n = nums.size();
        for (int i = 0; i < n; i++) {
            if (nums[i] == 0) continue;
            if (pre >= 0) ans = (long) ans * (i - pre) % MOD;
            pre = i;
        }
        return pre < 0 ? 0 : ans;
    }
};
```

```go [sol-Go]
func numberOfGoodSubarraySplits(nums []int) int {
	const mod int = 1e9 + 7
	ans, pre := 1, -1
	for i, x := range nums {
		if x > 0 {
			if pre >= 0 {
				ans = ans * (i - pre) % mod
			}
			pre = i
		}
	}
	if pre < 0 { // 整个数组都是 0，没有好子数组
		return 0
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

#### 相似题目

- [2147. 分隔长廊的方案数](https://leetcode.cn/problems/number-of-ways-to-divide-a-long-corridor/)

#### 算法小课堂：模运算

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

