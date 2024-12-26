**题意**：每次可以爬 $\textit{zero}$ 或 $\textit{one}$ 个台阶，返回爬 $\textit{low}$ 到 $\textit{high}$ 个台阶的方案数。

定义 $f[i]$ 表示构造长为 $i$ 的字符串的方案数，其中构造空串的方案数为 $1$，即 $f[0]=1$。

有两类得到长为 $i$ 的字符串的方法：

- 如果 $i\ge \textit{zero}$，那么可以在长为 $i-\textit{zero}$ 的字符串末尾添加 $\textit{zero}$ 个 $0$，方案数为 $f[i-\textit{zero}]$。
- 如果 $i\ge \textit{one}$，那么可以在长为 $i-\textit{one}$ 的字符串末尾添加 $\textit{one}$ 个 $1$，方案数为 $f[i-\textit{one}]$。
- 两类方案互斥（第 $i$ 个字符不能既是 $0$ 又是 $1$），所以用**加法原理**，得

$$
f[i] = f[i-\textit{zero}] + f[i-\textit{one}]
$$

对比一下 [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/)，相当于本题的 $\textit{zero}=1,\ \textit{one}=2$，即 $f[i]=f[i-1]+f[i-2]$。

初始值：$f[0]=1$，表示得到空串的方案数是 $1$，即什么也不做，也算一种方案。

答案：$\sum\limits_{i=\textit{low}}^{\textit{high}} f[i]$。

代码中用到了取模，不了解或者写错的同学请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

## 写法一：记忆化搜索

### 答疑

**问**：为什么要把每个 $\textit{dfs}(i)$ 都算一遍？能不能只算一次 $\textit{dfs}(\textit{high})$，然后累加 $\textit{memo}$ 数组中的计算结果？

**答**：这是不对的。比如 $\textit{zero} = \textit{one} = 2$，这种写法不会递归到 $i=\textit{high}-1$ 这个状态，$\textit{memo}[\textit{high}-1]$ 仍然是其初始值 $-1$。

**问**：那我不累加 $\textit{memo}[i]=-1$ 的结果可以吗？

**答**：仍然不行。比如 $\textit{zero} = \textit{one} = 2$ 且 $\textit{high}=5$，我们可以得到长为 $2$ 和 $4$ 的字符串，但这些答案并不能从 $\textit{dfs}(5)$ 中算出来。

```py [sol-Python3]
class Solution:
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        MOD = 1_000_000_007
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int) -> int:
            if i < 0:
                return 0
            if i == 0:
                return 1
            return (dfs(i - zero) + dfs(i - one)) % MOD
        return sum(dfs(i) for i in range(low, high + 1)) % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int countGoodStrings(int low, int high, int zero, int one) {
        int[] memo = new int[high + 1];
        Arrays.fill(memo, -1); // -1 表示没有计算过
        int ans = 0;
        for (int i = low; i <= high; i++) {
            ans = (ans + dfs(i, zero, one, memo)) % MOD;
        }
        return ans;
    }

    private int dfs(int i, int zero, int one, int[] memo) {
        if (i < 0) {
            return 0;
        }
        if (i == 0) {
            return 1;
        }
        if (memo[i] != -1) { // 之前计算过
            return memo[i];
        }
        return memo[i] = (dfs(i - zero, zero, one, memo) + dfs(i - one, zero, one, memo)) % MOD;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countGoodStrings(int low, int high, int zero, int one) {
        const int MOD = 1'000'000'007;
        vector<int> memo(high + 1, -1); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i) -> int {
            if (i < 0) {
                return 0;
            }
            if (i == 0) {
                return 1;
            }
            int& res = memo[i]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }
            return res = (dfs(dfs, i - zero) + dfs(dfs, i - one)) % MOD;
        };
        int ans = 0;
        for (int i = low; i <= high; i++) {
            ans = (ans + dfs(dfs, i)) % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countGoodStrings(low int, high int, zero int, one int) (ans int) {
    const mod = 1_000_000_007
    memo := make([]int, high+1)
    for i := range memo {
        memo[i] = -1 // -1 表示没有计算过
    }
    var dfs func(int) int
    dfs = func(i int) int {
        if i < 0 {
            return 0
        }
        if i == 0 {
            return 1
        }
        p := &memo[i]
        if *p == -1 { // 没有计算过
            *p = (dfs(i-zero) + dfs(i-one)) % mod
        }
        return *p
    }
    for i := low; i <= high; i++ {
        ans += dfs(i)
    }
    return ans % mod
}
```

```js [sol-JavaScript]
var countGoodStrings = function(low, high, zero, one) {
    const MOD = 1_000_000_007;
    const memo = Array(high + 1).fill(-1); // -1 表示没有计算过
    function dfs(i) {
        if (i < 0) {
            return 0;
        }
        if (i === 0) {
            return 1;
        }
        if (memo[i] !== -1) { // 之前计算过
            return memo[i];
        }
        return memo[i] = (dfs(i - zero) + dfs(i - one)) % MOD;
    }
    let ans = 0;
    for (let i = low; i <= high; i++) {
        ans = (ans + dfs(i)) % MOD;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_good_strings(low: i32, high: i32, zero: i32, one: i32) -> i32 {
        const MOD: i32 = 1_000_000_007;
        
        fn dfs(i: i32, zero: i32, one: i32, memo: &mut Vec<i32>) -> i32 {
            if i < 0 {
                return 0;
            }
            if i == 0 {
                return 1;
            }
            if memo[i as usize] != -1 { // 之前计算过
                return memo[i as usize];
            }
            memo[i as usize] = (dfs(i - zero, zero, one, memo) + dfs(i - one, zero, one, memo)) % MOD;
            memo[i as usize]
        }

        let mut ans = 0;
        let mut memo = vec![-1; (high + 1) as usize]; // -1 表示没有计算过
        for i in low..=high {
            ans = (ans + dfs(i, zero, one, &mut memo)) % MOD;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{high})$。
- 空间复杂度：$\mathcal{O}(\textit{high})$。

## 写法二：递推

### 答疑

**问**：代码中计算 $f[i]$ 的那两段代码是什么意思？

**答**：相当于把递推式 $f[i] = f[i-\textit{zero}] + f[i-\textit{one}]$ 拆分成了两步：

1. 第一步，把 $f[i-\textit{zero}]$ 赋值给 $f[i]$。
2. 第二步，把 $f[i-\textit{one}]$ 加到 $f[i]$ 中。
   
由于注意第一步必须满足 $i-\textit{zero}\ge 0$，否则无法转移。第二步也同理，必须满足 $i-\textit{one}\ge 0$。

```py [sol-Python3]
class Solution:
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        MOD = 1_000_000_007
        f = [1] + [0] * high  # f[i] 表示构造长为 i 的字符串的方案数
        for i in range(1, high + 1):
            if i >= zero: f[i] = f[i - zero]
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
            if (i >= zero) f[i] = f[i - zero];
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
            if (i >= zero) f[i] = f[i - zero];
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
        if i >= zero { f[i] = f[i-zero] }
        if i >= one  { f[i] = (f[i] + f[i-one]) % mod }
        if i >= low  { ans = (ans + f[i]) % mod }
    }
    return
}
```

```js [sol-JavaScript]
var countGoodStrings = function(low, high, zero, one) {
    const MOD = 1_000_000_007;
    const f = Array(high + 1).fill(0); // f[i] 表示构造长为 i 的字符串的方案数
    f[0] = 1; // 构造空串的方案数为 1
    let ans = 0;
    for (let i = 1; i <= high; i++) {
        if (i >= zero) f[i] = f[i - zero];
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
                f[i] = f[i - zero as usize];
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

## 优化

如果 $\textit{zero}$ 和 $\textit{one}$ 都是偶数，比如都是 $2$，由于偶数+偶数=偶数，无论如何，我们都不可能得到长为奇数的字符串。

如果此时 $\textit{low}=3,\ \textit{high}=7$，那么我们只需要计算字符串长度为 $4$ 和 $6$ 的情况。或者说，问题的规模可以缩小为 $\textit{low}'=2,\ \textit{high}'=3,\ \textit{zero}'=1,\ \textit{one}'=1$。

一般地，设 $\textit{zero}$ 和 $\textit{one}$ 的最大公约数（GCD）为 $g$，那么问题的规模可以缩小为

$$
\begin{cases} 
\textit{low}' = \left\lceil\dfrac{\textit{low}}{g}\right\rceil \\[2ex]
\textit{high}' = \left\lfloor\dfrac{\textit{high}}{g}\right\rfloor \\[2ex]
\textit{zero}' = \dfrac{\textit{zero}}{g} \\[2ex]
\textit{one}' = \dfrac{\textit{one}}{g} \\
\end{cases}
$$

关于上取整的计算，当 $a$ 和 $b$ 均为正整数时，我们有

$$
\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a-1}{b}\right\rfloor + 1
$$

讨论 $a$ 被 $b$ 整除，和不被 $b$ 整除两种情况，可以证明上式的正确性。

```py [sol-Python3]
class Solution:
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        g = gcd(zero, one)
        low = (low - 1) // g + 1
        high //= g
        zero //= g
        one //= g

        MOD = 1_000_000_007
        f = [1] + [0] * high  # f[i] 表示构造长为 i 的字符串的方案数
        for i in range(1, high + 1):
            if i >= zero: f[i] = f[i - zero]
            if i >= one:  f[i] = (f[i] + f[i - one]) % MOD
        return sum(f[low:]) % MOD
```

```java [sol-Java]
class Solution {
    public int countGoodStrings(int low, int high, int zero, int one) {
        int g = gcd(zero, one);
        low = (low - 1) / g + 1;
        high /= g;
        zero /= g;
        one /= g;

        final int MOD = 1_000_000_007;
        int ans = 0;
        int[] f = new int[high + 1]; // f[i] 表示构造长为 i 的字符串的方案数
        f[0] = 1; // 构造空串的方案数为 1
        for (int i = 1; i <= high; i++) {
            if (i >= zero) f[i] = f[i - zero];
            if (i >= one)  f[i] = (f[i] + f[i - one]) % MOD;
            if (i >= low)  ans = (ans + f[i]) % MOD;
        }
        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countGoodStrings(int low, int high, int zero, int one) {
        int g = gcd(zero, one);
        low = (low - 1) / g + 1;
        high /= g;
        zero /= g;
        one /= g;
    
        const int MOD = 1'000'000'007;
        int ans = 0;
        vector<int> f(high + 1); // f[i] 表示构造长为 i 的字符串的方案数
        f[0] = 1; // 构造空串的方案数为 1
        for (int i = 1; i <= high; i++) {
            if (i >= zero) f[i] = f[i - zero];
            if (i >= one)  f[i] = (f[i] + f[i - one]) % MOD;
            if (i >= low)  ans = (ans + f[i]) % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countGoodStrings(low, high, zero, one int) (ans int) {
    g := gcd(zero, one)
    low = (low-1)/g + 1
    high /= g
    zero /= g
    one /= g

    const mod = 1_000_000_007
    f := make([]int, high+1) // f[i] 表示构造长为 i 的字符串的方案数
    f[0] = 1 // 构造空串的方案数为 1
    for i := 1; i <= high; i++ {
        if i >= zero { f[i] = f[i-zero] }
        if i >= one  { f[i] = (f[i] + f[i-one]) % mod }
        if i >= low  { ans = (ans + f[i]) % mod }
    }
    return
}

func gcd(a, b int) int {
    for a != 0 {
        a, b = b%a, a
    }
    return b
}
```

```js [sol-JavaScript]
var countGoodStrings = function(low, high, zero, one) {
    const gcd = (a, b) => b === 0 ? a : gcd(b, a % b);
    const g = gcd(zero, one);
    low = Math.ceil(low / g);
    high = Math.floor(high / g);
    zero /= g;
    one /= g;

    const MOD = 1_000_000_007;
    const f = Array(high + 1).fill(0); // f[i] 表示构造长为 i 的字符串的方案数
    f[0] = 1; // 构造空串的方案数为 1
    let ans = 0;
    for (let i = 1; i <= high; i++) {
        if (i >= zero) f[i] = f[i - zero];
        if (i >= one)  f[i] = (f[i] + f[i - one]) % MOD;
        if (i >= low)  ans = (ans + f[i]) % MOD;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_good_strings(mut low: i32, mut high: i32, mut zero: i32, mut one: i32) -> i32 {
        let g = gcd(zero, one);
        low = (low - 1) / g + 1;
        high /= g;
        zero /= g;
        one /= g;

        const MOD: i32 = 1_000_000_007;
        let mut ans = 0;
        let mut f = vec![0; (high + 1) as usize]; // f[i] 表示构造长为 i 的字符串的方案数
        f[0] = 1; // 构造空串的方案数为 1
        for i in 1..=high as usize {
            if i >= zero as usize {
                f[i] = f[i - zero as usize];
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

fn gcd(a: i32, b: i32) -> i32 {
    if b == 0 { a } else { gcd(b, a % b) }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}\left(\dfrac{\textit{high}}{g}\right)$。其中 $g$ 为 $\textit{zero}$ 和 $\textit{one}$ 的 GCD。
- 空间复杂度：$\mathcal{O}\left(\dfrac{\textit{high}}{g}\right)$。

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
