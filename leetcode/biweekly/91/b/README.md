**题意**：每次可以爬 $\textit{zero}$ 或 $\textit{one}$ 个台阶，返回爬 $\textit{low}$ 到 $\textit{high}$ 个台阶的方案数。

定义 $f[i]$ 表示构造长为 $i$ 的字符串的方案数，其中构造空串的方案数为 $1$，即 $f[0]=1$。

有两类得到长为 $i$ 的字符串的方法：

- 如果 $i\ge \textit{zero}$，那么可以在长为 $i-\textit{zero}$ 的字符串末尾添加 $\textit{zero}$ 个 $0$，方案数为 $f[i-\textit{zero}]$。
- 如果 $i\ge \textit{one}$，那么可以在长为 $i-\textit{one}$ 的字符串末尾添加 $\textit{one}$ 个 $1$，方案数为 $f[i-\textit{one}]$。
- 两类方案互相独立，相加得

$$
f[i] = f[i-\textit{zero}] + f[i-\textit{one}]
$$

对比一下 [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/)，相当于本题的 $\textit{zero}=1,\ \textit{one}=2$，即 $f[i]=f[i-1]+f[i-2]$。

答案为 $\sum\limits_{i=\textit{low}}^{\textit{high}} f[i]$。

代码中用到了取模，不了解或者写错的同学请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
class Solution:
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        MOD = 1_000_000_007
        f = [1] + [0] * high  # f[i] 表示构造长为 i 的字符串的方案数
        for i in range(1, high + 1):
            if i >= zero: f[i] = (f[i] + f[i - zero]) % MOD
            if i >= one:  f[i] = (f[i] + f[i - one]) % MOD
        return sum(f[low:]) % MOD
```

```java [sol-Java]
class Solution {
    public int countGoodStrings(int low, int high, int zero, int one) {
        final int MOD = 1_000_000_007;
        int ans = 0;
        int[] f = new int[high + 1]; // f[i] 表示构造长为 i 的字符串的方案数
        f[0] = 1; // 构造空串的方案数为 1
        for (int i = 1; i <= high; i++) {
            if (i >= zero) f[i] = (f[i] + f[i - zero]) % MOD;
            if (i >= one)  f[i] = (f[i] + f[i - one]) % MOD;
            if (i >= low)  ans = (ans + f[i]) % MOD;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countGoodStrings(int low, int high, int zero, int one) {
        const int MOD = 1'000'000'007;
        int ans = 0;
        vector<int> f(high + 1); // f[i] 表示构造长为 i 的字符串的方案数
        f[0] = 1; // 构造空串的方案数为 1
        for (int i = 1; i <= high; i++) {
            if (i >= zero) f[i] = (f[i] + f[i - zero]) % MOD;
            if (i >= one)  f[i] = (f[i] + f[i - one]) % MOD;
            if (i >= low)  ans = (ans + f[i]) % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countGoodStrings(low, high, zero, one int) (ans int) {
	const mod = 1_000_000_007
	f := make([]int, high+1) // f[i] 表示构造长为 i 的字符串的方案数
	f[0] = 1 // 构造空串的方案数为 1
	for i := 1; i <= high; i++ {
		if i >= zero { f[i] = (f[i] + f[i-zero]) % mod }
		if i >= one  { f[i] = (f[i] + f[i-one]) % mod }
		if i >= low  { ans = (ans + f[i]) % mod }
	}
	return
}
```

```js [sol-JavaScript]
var countGoodStrings = function (low, high, zero, one) {
    const MOD = 1_000_000_007;
    const f = Array(high + 1).fill(0); // f[i] 表示构造长为 i 的字符串的方案数
    f[0] = 1; // 构造空串的方案数为 1
    let ans = 0;
    for (let i = 1; i <= high; i++) {
        if (i >= zero) f[i] = (f[i] + f[i - zero]) % MOD;
        if (i >= one)  f[i] = (f[i] + f[i - one]) % MOD;
        if (i >= low)  ans = (ans + f[i]) % MOD;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_good_strings(low: i32, high: i32, zero: i32, one: i32) -> i32 {
        const MOD: i32 = 1_000_000_007;
        let mut ans = 0;
        let mut f = vec![0; (high + 1) as usize]; // f[i] 表示构造长为 i 的字符串的方案数
        f[0] = 1; // 构造空串的方案数为 1
        for i in 1..=high as usize {
            if i >= zero as usize {
                f[i] = (f[i] + f[i - zero as usize]) % MOD;
            }
            if i >= one as usize {
                f[i] = (f[i] + f[i - one as usize]) % MOD;
            }
            if i >= low as usize {
                ans = (ans + f[i]) % MOD;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{high})$。
- 空间复杂度：$\mathcal{O}(\textit{high})$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
