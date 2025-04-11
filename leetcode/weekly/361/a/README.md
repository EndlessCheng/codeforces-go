## 方法一：枚举

枚举 $[\textit{low},\textit{high}]$ 中的整数 $i$。

设 $i$ 的十进制字符串为 $s$。如果 $s$ 的长度为偶数，且 $s$ 左半字符 ASCII 值之和等于 $s$ 右半字符 ASCII 值之和，那么 $i$ 是一个对称整数，答案加一。

```py [sol-Python3]
class Solution:
    def countSymmetricIntegers(self, low: int, high: int) -> int:
        ans = 0
        for i in range(low, high + 1):
            s = str(i)
            n = len(s)
            if n % 2 == 0 and sum(map(ord, s[:n // 2])) == sum(map(ord, s[n // 2:])):
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countSymmetricIntegers(int low, int high) {
        int ans = 0;
        for (int i = low; i <= high; i++) {
            char[] s = Integer.toString(i).toCharArray();
            int n = s.length;
            if (n % 2 > 0) {
                continue;
            }
            int diff = 0;
            for (int j = 0; j < n / 2; j++) {
                diff += s[j];
            }
            for (int j = n / 2; j < n; j++) {
                diff -= s[j];
            }
            if (diff == 0) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countSymmetricIntegers(int low, int high) {
        int ans = 0;
        for (int i = low; i <= high; i++) {
            string s = to_string(i);
            int n = s.size();
            if (n % 2 == 0 && reduce(s.begin(), s.begin() + n / 2) == reduce(s.begin() + n / 2, s.end())) {
                ans++;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSymmetricIntegers(low, high int) (ans int) {
    for i := low; i <= high; i++ {
        s := strconv.Itoa(i)
        n := len(s)
        if n%2 > 0 {
            continue
        }
        diff := 0
        for _, c := range s[:n/2] {
            diff += int(c)
        }
        for _, c := range s[n/2:] {
            diff -= int(c)
        }
        if diff == 0 {
            ans++
        }
    }
    return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((\textit{high} - \textit{low})\log \textit{high})$。
- 空间复杂度：$\mathcal{O}(\log \textit{high})$。

## 方法二：上下界数位 DP

如果 $\textit{high} = 10^{18}$，方法一就超时了，怎么办？

**前置知识**：

[数位 DP v1.0 视频讲解](https://www.bilibili.com/video/BV1rS4y1s721/)

[数位 DP v2.0 视频讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/)（上下界数位 DP）

**状态定义**：$\textit{dfs}(i, \textit{start}, \textit{diff},\textit{limitLow},\textit{limitHigh})$ 表示构造第 $i$ 位及其之后数位的合法方案数，其余参数的含义为：

- $\textit{start}$ 表示从左往右第一个非零数字的下标。用于判断现在处于对称整数的左半还是右半。
- $\textit{diff}$ 表示左半数字与右半数字之差。如果最终 $\textit{diff}=0$，则说明我们成功构造出一个对称整数。
- $\textit{limitHigh}$ 表示当前是否受到了 $\textit{high}$ 的约束（我们要构造的数字不能超过 $\textit{high}$）。若为真，则第 $i$ 位填入的数字至多为 $\textit{high}[i]$，否则至多为 $9$，这个数记作 $\textit{hi}$。如果在受到约束的情况下填了 $\textit{hi}$，那么后续填入的数字仍会受到 $\textit{high}$ 的约束。例如 $\textit{high}=123$，那么 $i=0$ 填的是 $1$ 的话，$i=1$ 的这一位至多填 $2$。
- $\textit{limitLow}$ 表示当前是否受到了 $\textit{low}$ 的约束（我们要构造的数字不能低于 $\textit{low}$）。若为真，则第 $i$ 位填入的数字至少为 $\textit{low}[i]$，否则至少为 $0$，这个数记作 $\textit{lo}$。如果在受到约束的情况下填了 $\textit{lo}$，那么后续填入的数字仍会受到 $\textit{low}$ 的约束。

**状态转移**：

- 如果前面没有填数字，且剩余数位个数是奇数，那么当前数位不能填数字（因为对称整数的长度必须是偶数），往后递归。在这种情况下，如果 $\textit{lo}>0$，那么必须填数字，但这不合法，直接返回 $0$。
- 否则，枚举第 $i$ 位填数字 $d=\textit{lo},\textit{lo}+1,\ldots,\textit{hi}$。如果之前没有填过数字，且当前填的数字大于 $0$，那么记录 $\textit{start}=i$。如果 $i< \dfrac{\textit{start}+n}{2}$，说明我们在左半，把 $\textit{diff}$ 增加 $d$，否则把 $\textit{diff}$ 减少 $d$。

**递归终点**：$i=n$ 时，如果 $\textit{diff}=0$，说明我们成功构造出一个对称整数，返回 $1$，否则返回 $0$。

**递归入口**：$\textit{dfs}(0, -1, 0, \texttt{true}, \texttt{true})$，表示：

- 从最高位开始。
- 一开始没有填任何数字。
- 一开始要受到 $\textit{low}$ 和 $\textit{high}$ 的约束（否则就可以随意填了，这肯定不行）。

```py [sol-Python3]
class Solution:
    def countSymmetricIntegers(self, low: int, high: int) -> int:
        high = list(map(int, str(high)))  # 避免在 dfs 中频繁调用 int()
        n = len(high)
        low = list(map(int, str(low).zfill(n)))  # 补前导零，和 high 对齐

        @cache
        def dfs(i: int, start: int, diff: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1 if diff == 0 else 0

            lo = low[i] if limit_low else 0
            hi = high[i] if limit_high else 9

            # 如果前面没有填数字，且剩余数位个数是奇数，那么当前数位不能填数字
            if start < 0 and (n - i) % 2:
                # 如果必须填数字（lo > 0），不合法，返回 0
                return 0 if lo else dfs(i + 1, start, diff, True, False)

            res = 0
            is_left = start < 0 or i < (start + n) // 2
            for d in range(lo, hi + 1):
                res += dfs(i + 1,
                           i if start < 0 and d else start,  # 记录第一个填数字的位置
                           diff + (d if is_left else -d),  # 左半 + 右半 -
                           limit_low and d == lo,
                           limit_high and d == hi)
            return res

        return dfs(0, -1, 0, True, True)
```

```java [sol-Java]
class Solution {
    private char[] lowS, highS;
    private int n, m, diffLh;
    private int[][][] memo;

    public int countSymmetricIntegers(int low, int high) {
        lowS = String.valueOf(low).toCharArray();
        highS = String.valueOf(high).toCharArray();
        n = highS.length;
        m = n / 2;
        diffLh = n - lowS.length;

        // dfs 中的 start <= diffLh，-9m <= diff <= 9m
        memo = new int[n][diffLh + 1][m * 18 + 1];
        for (int[][] mat : memo) {
            for (int[] row : mat) {
                Arrays.fill(row, -1);
            }
        }

        // 初始化 diff = m * 9，避免出现负数导致 memo 下标越界
        return dfs(0, -1, m * 9, true, true);
    }

    private int dfs(int i, int start, int diff, boolean limitLow, boolean limitHigh) {
        if (i == n) {
            return diff == m * 9 ? 1 : 0;
        }

        // start 当 isNum 用
        if (start != -1 && !limitLow && !limitHigh && memo[i][start][diff] != -1) {
            return memo[i][start][diff];
        }

        int lo = limitLow && i >= diffLh ? lowS[i - diffLh] - '0' : 0;
        int hi = limitHigh ? highS[i] - '0' : 9;

        // 如果前面没有填数字，且剩余数位个数是奇数，那么当前数位不能填数字
        if (start < 0 && (n - i) % 2 > 0) {
            // 如果必须填数字（lo > 0），不合法，返回 0
            return lo > 0 ? 0 : dfs(i + 1, start, diff, true, false);
        }

        int res = 0;
        boolean isLeft = start < 0 || i < (start + n) / 2;
        for (int d = lo; d <= hi; d++) {
            res += dfs(i + 1,
                       start < 0 && d > 0 ? i : start, // 记录第一个填数字的位置
                       diff + (isLeft ? d : -d), // 左半 +，右半 -
                       limitLow && d == lo,
                       limitHigh && d == hi);
        }

        if (start != -1 && !limitLow && !limitHigh) {
            memo[i][start][diff] = res;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countSymmetricIntegers(int low, int high) {
        string low_s = to_string(low), high_s = to_string(high);
        int n = high_s.size(), m = n / 2;
        int diff_lh = n - low_s.size();

        // dfs 中的 start <= diff_lh，-9m <= diff <= 9m
        vector memo(n, vector(diff_lh + 1, vector<int>(m * 18 + 1, -1)));
        auto dfs = [&](this auto&& dfs, int i, int start, int diff, bool limit_low, bool limit_high) -> int {
            if (i == n) {
                return diff == m * 9;
            }

            // start 当 is_num 用
            if (start != -1 && !limit_low && !limit_high && memo[i][start][diff] != -1) {
                return memo[i][start][diff];
            }

            int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
            int hi = limit_high ? high_s[i] - '0' : 9;

            // 如果前面没有填数字，且剩余数位个数是奇数，那么当前数位不能填数字
            if (start < 0 && (n - i) % 2) {
                // 如果必须填数字（lo > 0），不合法，返回 0
                return lo > 0 ? 0 : dfs(i + 1, start, diff, true, false);
            }

            int res = 0;
            bool is_left = start < 0 || i < (start + n) / 2;
            for (int d = lo; d <= hi; d++) {
                res += dfs(i + 1,
                           start < 0 && d > 0 ? i : start, // 记录第一个填数字的位置
                           diff + (is_left ? d : -d), // 左半 +，右半 -
                           limit_low && d == lo,
                           limit_high && d == hi);
            }

            if (start != -1 && !limit_low && !limit_high) {
                memo[i][start][diff] = res;
            }
            return res;
        };

        // 初始化 diff = m * 9，避免出现负数导致 memo 下标越界
        return dfs(0, -1, m * 9, true, true);
    }
};
```

```go [sol-Go]
func countSymmetricIntegers(low, high int) int {
    lowS := strconv.Itoa(low)
    highS := strconv.Itoa(high)
    n := len(highS)
    m := n / 2
    diffLH := n - len(lowS)

    memo := make([][][]int, n)
    for i := range memo {
        memo[i] = make([][]int, diffLH+1) // start <= diffLH
        for j := range memo[i] {
            memo[i][j] = make([]int, m*18+1) // -9m <= diff <= 9m
            for k := range memo[i][j] {
                memo[i][j][k] = -1
            }
        }
    }
    var dfs func(int, int, int, bool, bool) int
    dfs = func(i, start, diff int, limitLow, limitHigh bool) (res int) {
        if i == n {
            if diff != 0 {
                return 0
            }
            return 1
        }

        // start 当 isNum 用
        if start != -1 && !limitLow && !limitHigh {
            p := &memo[i][start][diff+m*9]
            if *p != -1 {
                return *p
            }
            defer func() { *p = res }()
        }

        lo := 0
        if limitLow && i >= diffLH {
            lo = int(lowS[i-diffLH] - '0')
        }
        hi := 9
        if limitHigh {
            hi = int(highS[i] - '0')
        }

        // 如果前面没有填数字，且剩余数位个数是奇数，那么当前数位不能填数字
        if start < 0 && (n-i)%2 > 0 {
            if lo > 0 {
                return 0 // 必须填数字但 lo > 0，不合法
            }
            return dfs(i+1, start, diff, true, false)
        }

        isLeft := start < 0 || i < (start+n)/2
        for d := lo; d <= hi; d++ {
            newStart := start
            if start < 0 && d > 0 {
                newStart = i // 记录第一个填数字的位置
            }
            newDiff := diff
            if isLeft {
                newDiff += d
            } else {
                newDiff -= d
            }
            res += dfs(i+1, newStart, newDiff, limitLow && d == lo, limitHigh && d == hi)
        }
        return
    }
    return dfs(0, -1, 0, true, true)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(D^2n^3)$，其中 $D=10$，$n=\mathcal{O}(\log \textit{high})$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题 $i$ 有 $\mathcal{O}(n)$ 个，$\textit{start}$ 有 $\mathcal{O}(n)$ 个，$\textit{diff}$ 有 $\mathcal{O}(Dn)$ 个，所以状态个数为 $\mathcal{O}(Dn^3)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以总的时间复杂度为 $\mathcal{O}(D^2n^3)$。
- 空间复杂度：$\mathcal{O}(Dn^3)$。保存多少状态，就需要多少空间。

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
