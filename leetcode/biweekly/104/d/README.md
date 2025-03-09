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

记得取模。关于取模的知识点，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

视频讲解见[【双周赛 104】](https://www.bilibili.com/video/BV1fV4y1r7e6/)第四题。

```py [sol-Python3]
class Solution:
    def sumOfPower(self, nums: List[int]) -> int:
        MOD = 1_000_000_007
        nums.sort()
        ans = s = 0
        for x in nums:  # x 作为最大值
            ans = (ans + x * x * (x + s)) % MOD
            s = (s * 2 + x) % MOD  # 递推计算下一个 s
        return ans
```

```java [sol-Java]
class Solution {
    public int sumOfPower(int[] nums) {
        final int MOD = 1_000_000_007;
        Arrays.sort(nums);
        long ans = 0;
        long s = 0;
        for (long x : nums) { // x 作为最大值
            ans = (ans + x * x % MOD * (x + s)) % MOD; // 中间模一次防止溢出
            s = (s * 2 + x) % MOD; // 递推计算下一个 s
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumOfPower(vector<int>& nums) {
        const int MOD = 1'000'000'007;
        ranges::sort(nums);
        int ans = 0, s = 0;
        for (long long x : nums) { // x 作为最大值
            ans = (ans + x * x % MOD * (x + s)) % MOD; // 中间模一次防止溢出
            s = (s * 2 + x) % MOD; // 递推计算下一个 s
        }
        return ans;
    }
};
```

```go [sol-Go]
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

```js [sol-JavaScript]
var sumOfPower = function(nums) {
    const MOD = 1_000_000_007;
    nums.sort((a, b) => a - b);
    let ans = 0, s = 0;
    for (const x of nums) { // x 作为最大值
        const bx = BigInt(x);
        ans += Number(bx * bx * (bx + BigInt(s)) % BigInt(MOD));
        s = (s * 2 + x) % MOD // 递推计算下一个 s
    }
    return ans % MOD;
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

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
