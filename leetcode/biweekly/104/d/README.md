视频讲解见[【双周赛 104】](https://www.bilibili.com/video/BV1fV4y1r7e6/)第四题。

## 思路

由于元素的顺序不影响答案，先排序。

设有 $a,b,c,d,e$ 五个数，顺序从小到大。

如果把 $d$ 当成最大值：

- 如果只选 $d$ 单独一个数，那么力量为 $d^3$。
- 选 $a$ 为最小值，由于中间的 $b$ 和 $c$ 可选可不选，一共有 $2^2$ 种方案，所以力量总和为 $d^2\cdot a\cdot 2^2$。
- 选 $b$ 为最小值，由于中间的 $c$ 可选可不选，一共有 $2^1$ 种方案，所以力量总和为 $d^2\cdot b\cdot 2^1$。
- 选 $c$ 为最小值，只有 $2^0=1$ 种方案，所以力量总和为 $d^2\cdot c\cdot 2^0$。 

因此，当 $d$ 为最大值时，$d$ 及其左侧元素对答案的贡献为

$$
d^3 + d^2\cdot (a\cdot 2^2 + b\cdot 2^1 + c\cdot 2^0)
$$

令 $s=a\cdot 2^2 + b\cdot 2^1 + c\cdot 2^0$，上式为

$$
d^3 + d^2\cdot s = d^2\cdot(d+s)
$$

继续，把 $e$ 当成最大值，观察 $s$ 如何变化，也就是 $a,b,c,d$ 作为最小值的贡献：

$$
\begin{aligned}
&\ a\cdot 2^3 + b\cdot 2^2 + c\cdot 2^1 + d\cdot 2^0\\
=&\ 2\cdot(a\cdot 2^2 + b\cdot 2^1 + c\cdot 2^0) + d\cdot 2^0\\
=&\ 2\cdot s + d\\
\end{aligned}
$$

这意味着，我们不需要枚举最小值，只需要枚举最大值，就可以把 $s$ 递推计算出来。

> 有关取模的知识点，见文末的讲解。

```py [sol1-Python3]
class Solution:
    def sumOfPower(self, nums: List[int]) -> int:
        MOD = 10 ** 9 + 7
        nums.sort()
        ans = s = 0
        for x in nums:  # x 作为最大值
            ans = (ans + x * x * (x + s)) % MOD
            s = (s * 2 + x) % MOD  # 递推计算下一个 s
        return ans
```

```java [sol1-Java]
class Solution {
    public int sumOfPower(int[] nums) {
        final long MOD = (long) 1e9 + 7;
        Arrays.sort(nums);
        long ans = 0, s = 0;
        for (long x : nums) { // x 作为最大值
            ans = (ans + x * x % MOD * (x + s)) % MOD; // 中间模一次防止溢出
            s = (s * 2 + x) % MOD; // 递推计算下一个 s
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
        for (long long x: nums) { // x 作为最大值
            ans = (ans + x * x % MOD * (x + s)) % MOD; // 中间模一次防止溢出
            s = (s * 2 + x) % MOD; // 递推计算下一个 s
        }
        return ans;
    }
};
```

```go [sol1-Go]
func sumOfPower(nums []int) (ans int) {
	const mod = 1_000_000_007
	sort.Ints(nums)
	s := 0
	for _, x := range nums { // x 作为最大值
		ans = (ans + x*x%mod*(x+s)) % mod // 中间模一次防止溢出
		s = (s*2 + x) % mod // 递推计算下一个 s
	}
	return
}
```

```js [sol1-JavaScript]
var sumOfPower = function (nums) {
    const mod = 1e9 + 7;
    nums.sort((a, b) => a - b);
    let ans = 0, s = 0;
    for (const x of nums) { // x 作为最大值
        const bx = BigInt(x);
        ans += Number(bx * bx * (bx + BigInt(s)) % BigInt(mod));
        s = (s * 2 + x) % mod // 递推计算下一个 s
    }
    return ans % mod;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈空间，仅用到若干额外变量。

## 思考题

把「子序列」改成「子数组」，要怎么做？

## 专题训练：贡献法

- [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/)
- [1856. 子数组最小乘积的最大值](https://leetcode.cn/problems/maximum-subarray-min-product/)
- [2104. 子数组范围和](https://leetcode.cn/problems/sum-of-subarray-ranges/)
- [2281. 巫师的总力量和](https://leetcode.cn/problems/sum-of-total-strength-of-wizards/)

[往期每日一题题解（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

## 附：模运算

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
