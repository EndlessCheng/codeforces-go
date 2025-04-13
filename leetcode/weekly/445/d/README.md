**前置知识**：

[数位 DP v1.0 视频讲解](https://www.bilibili.com/video/BV1rS4y1s721/)

[数位 DP v2.0 视频讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/)（上下界数位 DP）

把 $l$ 和 $r$ 转成 $b$ 进制，然后套数位 DP v2.0 模板。

**状态定义**：$\textit{dfs}(i, \textit{pre}, \textit{limitLow},\textit{limitHigh})$ 表示构造第 $i$ 位及其之后数位的合法方案数，其余参数的含义为：

- $\textit{pre}$ 表示前一个位置填的数字，初始值为 $0$。
- $\textit{limitHigh}$ 表示当前是否受到了 $\textit{high}$ 的约束（我们要构造的数字不能超过 $\textit{high}$）。若为真，则第 $i$ 位填入的数字至多为 $\textit{high}[i]$，否则至多为 $b-1$，这个数记作 $\textit{hi}$。如果在受到约束的情况下填了 $\textit{hi}$，那么后续填入的数字仍会受到 $\textit{high}$ 的约束。例如 $\textit{high}=123$，那么 $i=0$ 填的是 $1$ 的话，$i=1$ 的这一位至多填 $2$。
- $\textit{limitLow}$ 表示当前是否受到了 $\textit{low}$ 的约束（我们要构造的数字不能低于 $\textit{low}$）。若为真，则第 $i$ 位填入的数字至少为 $\textit{low}[i]$，否则至少为 $0$，这个数记作 $\textit{lo}$。如果在受到约束的情况下填了 $\textit{lo}$，那么后续填入的数字仍会受到 $\textit{low}$ 的约束。

**状态转移**：枚举第 $i$ 位填数字 $d=\max(\textit{lo},\textit{pre}),\textit{lo}+1,\ldots,\textit{hi}$。继续递归，把 $\textit{i}$ 加一，把 $\textit{pre}$ 置为 $d$。

**递归终点**：$i=n$ 时，找到了一个合法数字，返回 $1$。

**递归入口**：$\textit{dfs}(0, 0, \texttt{true}, \texttt{true})$，表示：

- 从最高位开始。
- 假设第一个数字的前面是 $0$，这样第一个数字不会受到 $\textit{pre}$ 的约束。
- 一开始要受到 $\textit{low}$ 和 $\textit{high}$ 的约束（否则就可以随意填了，这肯定不行）。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1e3dBYLEDz/?t=29m32s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countNumbers(self, l: str, r: str, b: int) -> int:
        # 把 s 转成 b 进制
        def trans(s: str) -> List[int]:
            x = int(s)
            digits = []
            while x:
                x, r = divmod(x, b)
                digits.append(r)
            digits.reverse()
            return digits

        high = trans(r)
        n = len(high)
        low = trans(l)
        low = [0] * (n - len(low)) + low

        @cache
        def dfs(i: int, pre: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1

            lo = low[i] if limit_low else 0
            hi = high[i] if limit_high else b - 1

            res = 0
            for d in range(max(lo, pre), hi + 1):
                res += dfs(i + 1, d, limit_low and d == lo, limit_high and d == hi)
            return res % 1_000_000_007

        return dfs(0, 0, True, True)
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    private static final int MOD = 1_000_000_007;

    public int countNumbers(String l, String r, int b) {
        char[] low = trans(l, b);
        char[] high = trans(r, b);
        int[][] memo = new int[high.length][b];
        for (int[] row : memo) {
            Arrays.fill(row, -1);
        }
        return dfs(0, 0, true, true, b, low, high, memo);
    }

    // 把十进制字符串 s 转成 b 进制字符数组
    private char[] trans(String s, int b) {
        return new BigInteger(s).toString(b).toCharArray();
    }

    private int dfs(int i, int pre, boolean limitLow, boolean limitHigh, int b, char[] low, char[] high, int[][] memo) {
        if (i == high.length) {
            return 1;
        }
        if (!limitLow && !limitHigh && memo[i][pre] >= 0) {
            return memo[i][pre];
        }

        int diffLH = high.length - low.length;
        int lo = limitLow && i >= diffLH ? low[i - diffLH] - '0' : 0;
        int hi = limitHigh ? high[i] - '0' : b - 1;

        long res = 0;
        for (int d = Math.max(lo, pre); d <= hi; d++) {
            res += dfs(i + 1, d, limitLow && d == lo, limitHigh && d == hi, b, low, high, memo);
        }
        res %= MOD;

        if (!limitLow && !limitHigh) {
            memo[i][pre] = (int) res;
        }
        return (int) res;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 把十进制字符串 s 转成 b 进制
    // 用小学学过的【竖式除法】计算，读者可以先用竖式除法算算 1234÷10，再对照下面的代码
    vector<int> trans(string& s, int b) {
        for (char& c : s) {
            c -= '0';
        }
        vector<int> digits;
        while (!s.empty()) {
            string nxt_s; // 用竖式除法计算 s / b 得到的商（十进制）
            int rem = 0; // s % b
            for (char c : s) {
                rem = rem * 10 + c;
                int q = rem / b; // 商
                if (q || !nxt_s.empty()) {
                    nxt_s.push_back(q);
                }
                rem = rem % b;
            }
            digits.push_back(rem);
            s = move(nxt_s);
        }
        ranges::reverse(digits);
        return digits;
    }

public:
    int countNumbers(string l, string r, int b) {
        vector<int> low = trans(l, b);
        vector<int> high = trans(r, b);
        int n = high.size();
        int diff_lh = n - low.size();

        vector memo(n, vector<int>(b, -1));
        auto dfs = [&](this auto&& dfs, int i, int pre, bool limit_low, bool limit_high) -> int {
            if (i == n) {
                return 1;
            }
            if (!limit_low && !limit_high && memo[i][pre] >= 0) {
                return memo[i][pre];
            }

            int lo = limit_low && i >= diff_lh ? low[i - diff_lh] : 0;
            int hi = limit_high ? high[i] : b - 1;

            long long res = 0;
            for (int d = max(lo, pre); d <= hi; d++) {
                res += dfs(i + 1, d, limit_low && d == lo, limit_high && d == hi);
            }
            res %= 1'000'000'007;

            if (!limit_low && !limit_high) {
                memo[i][pre] = res;
            }
            return res;
        };
        return dfs(0, 0, true, true);
    }
};
```

```go [sol-Go]
func trans(s string, b int) string {
	x := big.Int{}
	fmt.Fscan(strings.NewReader(s), &x)
	return x.Text(b) // 转成 b 进制
}

func countNumbers(l, r string, b int) int {
	const mod = 1_000_000_007
	lowS := trans(l, b)
	highS := trans(r, b)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, b)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int
	dfs = func(i, pre int, limitLow, limitHigh bool) (res int) {
		if i == n {
			return 1
		}
		if !limitLow && !limitHigh {
			p := &memo[i][pre]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := b - 1
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		for d := max(lo, pre); d <= hi; d++ {
			res += dfs(i+1, d, limitLow && d == lo, limitHigh && d == hi)
		}
		return res % mod
	}
	return dfs(0, 0, true, true)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(bn^2)$，其中 $n$ 是 $r$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^2)$，单个状态的计算时间为 $\mathcal{O}(b)$，所以总的时间复杂度为 $\mathcal{O}(bn^2)$。
- 空间复杂度：$\mathcal{O}(n^2)$。保存多少状态，就需要多少空间。

更多相似题目，见下面动态规划题单中的「**十、数位 DP**」。

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
