## 本题视频讲解

请看 [数位 DP 通用模板](https://www.bilibili.com/video/BV1rS4y1s721/?t=20m20s)。

除了讲解模板，还讲了如何使用该模板秒杀相关困难题目。

## 前置知识：位运算与集合论

集合可以用二进制表示，二进制从低到高第 $i$ 位为 $1$ 表示 $i$ 在集合中，为 $0$ 表示 $i$ 不在集合中。例如集合 $\{0,2,3\}$ 对应的二进制数为 $1101_{(2)}$。

设集合对应的二进制数为 $x$。本题需要用到两个位运算操作：

1. 判断元素 $d$ 是否在集合中：`x >> d & 1` 可以取出 $x$ 的第 $d$ 个比特位，如果是 $1$ 就说明 $d$ 在集合中。
2. 把元素 $d$ 添加到集合中：将 `x` 更新为 `x | (1 << d)`。

更多位运算的知识点，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

## 思路

将 $n$ 转换成字符串 $s$，定义 $\textit{dfs}(i,\textit{mask}, \textit{isLimit},\textit{isNum})$ 表示构造第 $i$ 位及其之后数位的合法方案数，其余参数的含义为：

- $\textit{mask}$ 表示前面选过的数字集合，换句话说，第 $i$ 位要选的数字不能在 $\textit{mask}$ 中。
- $\textit{isLimit}$ 表示当前是否受到了 $n$ 的约束（注意要构造的数字不能超过 $n$）。若为真，则第 $i$ 位填入的数字至多为 $s[i]$，否则可以是 $9$。如果在受到约束的情况下填了 $s[i]$，那么后续填入的数字仍会受到 $n$ 的约束。例如 $n=123$，如果 $i=0$ 填的是 $1$ 的话，$i=1$ 的这一位至多填 $2$。如果 $i=0$ 填的是 $1$，$i=1$ 填的是 $2$，那么 $i=2$ 的这一位至多填 $3$。
- $\textit{isNum}$ 表示 $i$ 前面的数位是否填了数字。若为假，则当前位可以跳过（不填数字），或者要填入的数字至少为 $1$；若为真，则要填入的数字可以从 $0$ 开始。例如 $n=123$，在 $i=0$ 时跳过的话，相当于后面要构造的是一个 $99$ 以内的数字了，如果 $i=1$ 不跳过，那么相当于构造一个 $10$ 到 $99$ 的两位数，如果 $i=1$ 跳过，相当于构造的是一个 $9$ 以内的数字。
- 为什么要定义 $\textit{isNum}$？因为 $010$ 和 $10$ 都是 $10$，如果认为第一个 $0$ 和第三个 $0$ 都是我们填入的数字，这就不符合题目要求了，但 $10$ 显然是符合题目要求的。

## 实现细节

**递归入口**：$\textit{dfs}(0,0,\texttt{true},\texttt{false})$，表示：

- 从 $s[0]$ 开始枚举；
- 一开始集合中没有数字（空集）；
- 一开始要受到 $n$ 的约束（否则就可以随意填了，这肯定不行）；
- 一开始没有填数字。

**递归中**：

- 如果 $\textit{isNum}$ 为假，说明前面没有填数字，那么当前也可以不填数字。一旦从这里递归下去，$\textit{isLimit}$ 就可以置为 `false` 了，这是因为 $s[0]$ 必然是大于 $0$ 的，后面就不受到 $n$ 的约束了。或者说，最高位不填数字，后面无论怎么填都比 $n$ 小。
- 如果 $\textit{isNum}$ 为真，那么当前必须填一个数字。枚举填入的数字，根据 $\textit{isNum}$ 和 $\textit{isLimit}$ 来决定填入数字的范围。

**递归终点**：当 $i$ 等于 $s$ 长度时，如果 $\textit{isNum}$ 为真，则表示得到了一个合法数字（因为不合法的数字不会递归到终点），返回 $1$，否则返回 $0$。

## 答疑

**问**：$\textit{isNum}$ 这个参数可以去掉吗？

**答**：本题由于 $\textit{mask}$ 中记录了数字，可以通过判断 $\textit{mask}$ 是否为 $0$（空集）来判断前面是否填了数字，所以对于本题来说，$\textit{isNum}$ 可以省略。

下面的代码保留了 $\textit{isNum}$，主要是为了方便大家掌握这个模板。因为有些题目不需要 $\textit{mask}$，但需要 $\textit{isNum}$。

**问**：记忆化四个状态有点麻烦，能不能只记忆化 $(i,\textit{mask})$ 这两个状态？

**答**：是可以的。比如 $n=234$，第一位填 $2$，第二位填 $3$，后面无论怎么递归，都不会再次递归到第一位填 $2$，第二位填 $3$ 的情况，所以不需要记录。又比如，第一位不填，第二位也不填，后面无论怎么递归也不会再次递归到这种情况，所以也不需要记录。

根据这个例子，我们可以只记录不受到 $\textit{isLimit}$ 或 $\textit{isNum}$ 约束时的状态 $(i,\textit{mask})$。比如 $n=234$，第一位（最高位）填的 $1$，那么继续递归，后面就可以随便填，所以状态 $(1,2)$ 就表示前面填了一个 $1$（对应的 $\textit{mask}=2$），从第二位往后随便填的方案数。

相当于我们记忆化的是 $(i,\textit{mask},\texttt{false},\texttt{true})$。

**问**：能不能只记忆化 $i$？

**答**：这是不行的。想一想，我们为什么要用记忆化？如果递归到同一个状态时，计算出的结果是一样的，那么第二次递归到同一个状态，就可以直接返回第一次计算的结果了。通过保存第一次计算的结果，来优化时间复杂度。

由于前面选的数字会影响后面选的数字，两次递归到相同的 $i$，如果前面选的数字不一样，计算出的结果就可能是不一样的。如果只记忆化 $i$，就可能会算出错误的结果。

```py [sol-Python3]
class Solution:
    def countSpecialNumbers(self, n: int) -> int:
        s = str(n)
        @cache  # 缓存装饰器，避免重复计算 dfs 的结果（记忆化）
        def dfs(i: int, mask: int, is_limit: bool, is_num: bool) -> int:
            if i == len(s):
                return 1 if is_num else 0  # is_num 为 True 表示得到了一个合法数字
            res = 0
            if not is_num:  # 可以跳过当前数位
                res = dfs(i + 1, mask, False, False)
            # 如果前面没有填数字，则必须从 1 开始（因为不能有前导零）
            low = 0 if is_num else 1
            # 如果前面填的数字都和 n 的一样，那么这一位至多填 s[i]（否则就超过 n 啦）
            up = int(s[i]) if is_limit else 9
            for d in range(low, up + 1):  # 枚举要填入的数字 d
                if mask >> d & 1 == 0:  # d 不在 mask 中，说明之前没有填过 d
                    res += dfs(i + 1, mask | (1 << d), is_limit and d == up, True)
            return res
        return dfs(0, 0, True, False)
```

```java [sol-Java]
class Solution {
    public int countSpecialNumbers(int n) {
        char[] s = Integer.toString(n).toCharArray();
        int[][] memo = new int[s.length][1 << 10];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(0, 0, true, false, s, memo);
    }

    private int dfs(int i, int mask, boolean isLimit, boolean isNum, char[] s, int[][] memo) {
        if (i == s.length) {
            return isNum ? 1 : 0; // isNum 为 true 表示得到了一个合法数字
        }
        if (!isLimit && isNum && memo[i][mask] != -1) {
            return memo[i][mask]; // 之前计算过
        }
        int res = 0;
        if (!isNum) { // 可以跳过当前数位
            res = dfs(i + 1, mask, false, false, s, memo);
        }
        // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
        int up = isLimit ? s[i] - '0' : 9;
        // 枚举要填入的数字 d
        // 如果前面没有填数字，则必须从 1 开始（因为不能有前导零）
        for (int d = isNum ? 0 : 1; d <= up; d++) {
            if ((mask >> d & 1) == 0) { // d 不在 mask 中，说明之前没有填过 d
                res += dfs(i + 1, mask | (1 << d), isLimit && d == up, true, s, memo);
            }
        }
        if (!isLimit && isNum) {
            memo[i][mask] = res; // 记忆化
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countSpecialNumbers(int n) {
        string s = to_string(n);
        int m = s.length();
        vector<vector<int>> memo(m, vector<int>(1 << 10, -1)); // -1 表示没有计算过
        auto dfs = [&](auto&& dfs, int i, int mask, bool is_limit, bool is_num) -> int {
            if (i == m) {
                return is_num; // is_num 为 true 表示得到了一个合法数字
            }
            if (!is_limit && is_num && memo[i][mask] != -1) {
                return memo[i][mask]; // 之前计算过
            }
            int res = 0;
            if (!is_num) { // 可以跳过当前数位
                res = dfs(dfs, i + 1, mask, false, false);
            }
            // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
            int up = is_limit ? s[i] - '0' : 9;
            // 枚举要填入的数字 d
            // 如果前面没有填数字，则必须从 1 开始（因为不能有前导零）
            for (int d = is_num ? 0 : 1; d <= up; d++) {
                if ((mask >> d & 1) == 0) { // d 不在 mask 中，说明之前没有填过 d
                    res += dfs(dfs, i + 1, mask | (1 << d), is_limit && d == up, true);
                }
            }
            if (!is_limit && is_num) {
                memo[i][mask] = res; // 记忆化
            }
            return res;
        };
        return dfs(dfs, 0, 0, true, false);
    }
};
```

```c [sol-C]
#define MAX_M 11
int memo[MAX_M][1 << 10];

int dfs(int i, int mask, bool is_limit, bool is_num, const char* s) {
    if (s[i] == '\0') {
        return is_num; // is_num 为 true 表示得到了一个合法数字
    }
    if (!is_limit && is_num && memo[i][mask] != -1) {
        return memo[i][mask]; // 之前计算过
    }
    int res = 0;
    if (!is_num) { // 可以跳过当前数位
        res = dfs(i + 1, mask, false, false, s);
    }
    // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
    int up = is_limit ? s[i] - '0' : 9;
    // 枚举要填入的数字 d
    // 如果前面没有填数字，则必须从 1 开始（因为不能有前导零）
    for (int d = is_num ? 0 : 1; d <= up; d++) {
        if ((mask >> d & 1) == 0) { // d 不在 mask 中，说明之前没有填过 d
            res += dfs(i + 1, mask | (1 << d), is_limit && d == up, 1, s);
        }
    }
    if (!is_limit && is_num) {
        memo[i][mask] = res; // 记忆化
    }
    return res;
}

int countSpecialNumbers(int n) {
    char s[MAX_M];
    sprintf(s, "%d", n); // 将数字 n 转换为字符串
    int m = strlen(s);
    memset(memo, -1, m * (1 << 10) * sizeof(int)); // -1 表示没有计算过
    return dfs(0, 0, true, false, s);
}
```

```go [sol-Go]
func countSpecialNumbers(n int) int {
    s := strconv.Itoa(n)
    m := len(s)
    memo := make([][1 << 10]int, m)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1 // -1 表示没有计算过
        }
    }
    var dfs func(int, int, bool, bool) int
    dfs = func(i, mask int, isLimit, isNum bool) (res int) {
        if i == m {
            if isNum {
                return 1 // 得到了一个合法数字
            }
            return
        }
        if !isLimit && isNum {
            p := &memo[i][mask]
            if *p >= 0 { // 之前计算过
                return *p
            }
            defer func() { *p = res }() // 记忆化
        }
        if !isNum { // 可以跳过当前数位
            res += dfs(i+1, mask, false, false)
        }
        d := 0
        if !isNum {
            d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
        }
        up := 9
        if isLimit {
            up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
        }
        for ; d <= up; d++ { // 枚举要填入的数字 d
            if mask>>d&1 == 0 { // d 不在 mask 中，说明之前没有填过 d
                res += dfs(i+1, mask|1<<d, isLimit && d == up, true)
            }
        }
        return
    }
    return dfs(0, 0, true, false)
}
```

```js [sol-JavaScript]
var countSpecialNumbers = function(n) {
    const s = n.toString();
    const m = s.length;
    const memo = Array.from({ length: m }, () => Array(1 << 10).fill(-1)); // -1 表示没有计算过
    function dfs(i, mask, isLimit, isNum) {
        if (i === m) {
            return isNum ? 1 : 0; // is_num 为 true 表示得到了一个合法数字
        }
        if (!isLimit && isNum && memo[i][mask] !== -1) {
            return memo[i][mask]; // 之前计算过
        }
        let res = 0;
        if (!isNum) { // 可以跳过当前数位
            res += dfs(i + 1, mask, false, false);
        }
        // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
        const up = isLimit ? +s[i] : 9;
        // 枚举要填入的数字 d
        // 如果前面没有填数字，则必须从 1 开始（因为不能有前导零）
        for (let d = isNum ? 0 : 1; d <= up; d++) {
            if ((mask >> d & 1) === 0) { // d 不在 mask 中，说明之前没有填过 d
                res += dfs(i + 1, mask | (1 << d), isLimit && d === up, true);
            }
        }
        if (!isLimit && isNum) {
            memo[i][mask] = res; // 记忆化
        }
        return res;
    }
    return dfs(0, 0, true, false);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_special_numbers(n: i32) -> i32 {
        fn dfs(i: usize, mask: usize, is_limit: bool, is_num: bool, s: &[u8], memo: &mut Vec<Vec<i32>>) -> i32 {
            if i == s.len() {
                return if is_num { 1 } else { 0 }; // is_num 为 true 表示得到了一个合法数字
            }
            if !is_limit && is_num && memo[i][mask] != -1 {
                return memo[i][mask]; // 之前计算过
            }
            let mut res = 0;
            if !is_num { // 可以跳过当前数位
                res = dfs(i + 1, mask, false, false, s, memo);
            }
            // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
            let up = if is_limit { s[i] - b'0' } else { 9 };
            // 枚举要填入的数字 d
            // 如果前面没有填数字，则必须从 1 开始（因为不能有前导零）
            let low = if is_num { 0 } else { 1 };
            for d in low..=up {
                if (mask >> d & 1) == 0 { // d 不在 mask 中，说明之前没有填过 d
                    res += dfs(i + 1, mask | (1 << d), is_limit && d == up, true, s, memo);
                }
            }
            if !is_limit && is_num {
                memo[i][mask] = res; // 记忆化
            }
            return res;
        }

        let s = n.to_string();
        let s = s.as_bytes();
        let mut memo = vec![vec![-1; 1 << 10]; s.len()]; // -1 表示没有计算过
        return dfs(0, 0, true, false, &s, &mut memo);
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mD2^D)$，其中 $m$ 为 $s$ 的长度，即 $\mathcal{O}(\log n)$，$D=10$。由于每个状态只会计算一次，因此动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数为 $\mathcal{O}(m2^D)$，单个状态的计算时间为 $\mathcal{O}(D)$，因此时间复杂度为 $\mathcal{O}(mD2^D)$。
- 空间复杂度：$\mathcal{O}(m2^D)$。

更多相似题目，见下面动态规划题单中的「**数位 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
