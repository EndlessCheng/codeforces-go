下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

### 前置知识：模运算

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

### 提示 1

元素的顺序不影响答案，先排序。

### 提示 2

枚举 $x=\textit{nums}[i]$ 作为最大值。

$x$ 一个数的力量为 $x^3$。

其余不超过 $\textit{nums}[i]$ 的数字作为最小值时，贡献是多少？

例如有 $a,b,c,d,e$ 五个数，且顺序从小到大。

枚举到 $d$ 时：

- $a$ 作为最小值，中间的 $b$ 和 $c$ 可选可不选，一共有 $2^2$ 种方案，所以 $a$ 的贡献是 $a\cdot 2^2$。
- $b$ 作为最小值，中间的 $c$ 可选可不选，一共有 $2^1$ 种方案，所以 $b$ 的贡献是 $b\cdot 2^1$。
- $c$ 作为最小值，一共有 $2^0$ 种方案，所以 $c$ 的贡献是 $c\cdot 2^0$。 

这些元素作为最小值的贡献为 

$$
a\cdot 2^2 + b\cdot 2^1 + c\cdot 2^0
$$

记作 $\textit{s}$。

那么 $d$ 及其左侧元素对答案的贡献为

$$
d^3 + d^2\cdot s = d^2\cdot(d+s)
$$

继续，枚举到 $e$ 时：

$a,b,c,d$ 作为最小值的贡献为

$$
\begin{aligned}
&\ a\cdot 2^3 + b\cdot 2^2 + c\cdot 2^1 + d\cdot 2^0\\
=&\ 2\cdot(a\cdot 2^2 + b\cdot 2^1 + c\cdot 2^0) + d\cdot 2^0\\
=&\ 2\cdot s + d\\
\end{aligned}
$$

上式为新的 $s$，即

$$
s_{\textit{new}} = 2\cdot s + \textit{nums}[i]
$$

利用这一递推式，就可以 $\mathcal{O}(1)$ 地计算出排序后的每个 $\textit{nums}[i]$ 左侧元素作为最小值的贡献。

```py [sol1-Python3]
class Solution:
    def sumOfPower(self, nums: List[int]) -> int:
        MOD = 10 ** 9 + 7
        nums.sort()
        ans = s = 0
        for x in nums:
            ans = (ans + x * x * (x + s)) % MOD
            s = (s * 2 + x) % MOD
        return ans
```

```java [sol1-Java]
class Solution {
    public int sumOfPower(int[] nums) {
        final long MOD = (long) 1e9 + 7;
        Arrays.sort(nums);
        long ans = 0, s = 0;
        for (long x : nums) {
            ans = (ans + x * x % MOD * (x + s)) % MOD; // 中间模一次防止溢出
            s = (s * 2 + x) % MOD;
        }
        return (int) ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int sumOfPower(vector<int> &nums) {
        const int MOD = 1e9 + 7;
        sort(nums.begin(), nums.end());
        int ans = 0, s = 0;
        for (long long x: nums) {
            ans = (ans + x * x % MOD * (x + s)) % MOD; // 中间模一次防止溢出
            s = (s * 2 + x) % MOD;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func sumOfPower(nums []int) (ans int) {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	s := 0
	for _, x := range nums {
		ans = (ans + x*x%mod*(x+s)) % mod // 中间模一次防止溢出
		s = (s*2 + x) % mod
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序时的栈空间，仅用到若干额外变量。

### 思考题

把「子序列」改成「子数组」，要怎么做？

### 专题训练：贡献法

- [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/)
- [1856. 子数组最小乘积的最大值](https://leetcode.cn/problems/maximum-subarray-min-product/)
- [2104. 子数组范围和](https://leetcode.cn/problems/sum-of-subarray-ranges/)
- [2281. 巫师的总力量和](https://leetcode.cn/problems/sum-of-total-strength-of-wizards/)


