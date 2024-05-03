这题和 [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/) 其实是一样的，把那道题的 $1$ 和 $2$ 替换成 $\textit{zero}$ 和 $\textit{one}$ 你就认识了。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1gd4y1b7qj/?t=1m49s) 第二题。

```py [sol-Python3]
class Solution:
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        MOD = 1_000_000_007
        f = [1] + [0] * high  # f[i] 表示构造长为 i 的字符串的方案数，其中构造空串的方案数为 1
        for i in range(1, high + 1):
            if i >= one:  f[i] = (f[i] + f[i - one]) % MOD
            if i >= zero: f[i] = (f[i] + f[i - zero]) % MOD
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
            if (i >= one)  f[i] = (f[i] + f[i - one]) % MOD;
            if (i >= zero) f[i] = (f[i] + f[i - zero]) % MOD;
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
            if (i >= one)  f[i] = (f[i] + f[i - one]) % MOD;
            if (i >= zero) f[i] = (f[i] + f[i - zero]) % MOD;
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
		if i >= one  { f[i] = (f[i] + f[i-one]) % mod }
		if i >= zero { f[i] = (f[i] + f[i-zero]) % mod }
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
        if (i >= one)  f[i] = (f[i] + f[i - one]) % MOD;
        if (i >= zero) f[i] = (f[i] + f[i - zero]) % MOD;
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
            if i >= one as usize {
                f[i] = (f[i] + f[i - one as usize]) % MOD;
            }
            if i >= zero as usize {
                f[i] = (f[i] + f[i - zero as usize]) % MOD;
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

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
