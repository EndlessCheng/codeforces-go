**前置知识**：

- [数位 DP v1.0 模板讲解](https://www.bilibili.com/video/BV1rS4y1s721/?t=19m36s)
- [数位 DP v2.0 模板讲解](https://www.bilibili.com/video/BV1Fg4y1Q7wv/?t=31m28s)（上下界数位 DP）

注意到，数位之和是很小的，最多 $9n\le 81$（$n$ 是 $r$ 的十进制长度），所以可以用数位之和作为 DP 的状态之一。

注意到，$9$ 个 $0$ 到 $9$ 的数相乘，只有 $K(9)=3026$ 种不同的乘积（证明见文末），所以可以直接用数位乘积作为 DP 的状态之一。

**状态定义**。$\textit{dfs}(i,m,s,\textit{limitLow},\textit{limitHigh})$ 表示在如下约束时的美丽整数个数：

- 前 $i$ 个数位已经填完了。
- 之前填的数位乘积为 $m$。
- 之前填的数位之和为 $s$。
- 另外两个参数见模板讲解。

枚举当前数位填 $d$，那么 $m$ 变成 $m\cdot d$，$s$ 变成 $s+d$，继续递归。

**递归边界**。如果 $i=n$：

- 如果 $s = 0$ 或者 $m\bmod s \ne 0$，返回 $0$。
- 否则找到了一个美丽整数，返回 $1$。

**递归入口**：$\textit{dfs}(0,1,0,\texttt{true},\texttt{true})$。一开始没有填数字，数位乘积为 $1$（乘法单位元），数位之和为 $0$（加法单位元）。

代码实现时，如果 $\textit{limitLow}=\texttt{true}$，且 $i$ 比 $r$ 和 $l$ 的十进制长度之差还小，那么当前数位可以不填。这样就无需 $\textit{isNum}$ 参数了。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1JYQ8YWEvD/?t=30m42s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def beautifulNumbers(self, l: int, r: int) -> int:
        low = list(map(int, str(l)))
        high = list(map(int, str(r)))
        n = len(high)
        diff_lh = n - len(low)  # 这样写无需给 low 补前导零，也无需 is_num 参数

        @cache
        def dfs(i: int, m: int, s: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1 if s and m % s == 0 else 0

            lo = low[i - diff_lh] if limit_low and i >= diff_lh else 0
            hi = high[i] if limit_high else 9

            res = 0
            if limit_low and i < diff_lh:
                res += dfs(i + 1, 1, 0, True, False)  # 什么也不填
                d = 1  # 下面循环从 1 开始
            else:
                d = lo
            # 枚举填数字 d
            for d in range(d, hi + 1):
                res += dfs(i + 1, m * d, s + d, limit_low and d == lo, limit_high and d == hi)
            return res

        return dfs(0, 1, 0, True, True)
```

```java [sol-Java]
class Solution {
    public int beautifulNumbers(int l, int r) {
        char[] low = String.valueOf(l).toCharArray(); // 无需补前导零
        char[] high = String.valueOf(r).toCharArray();
        Map<Long, Integer> memo = new HashMap<>();
        return dfs(0, 1, 0, true, true, low, high, memo);
    }

    private int dfs(int i, int m, int s, boolean limitLow, boolean limitHigh,
                    char[] low, char[] high, Map<Long, Integer> memo) {
        if (i == high.length) {
            return s > 0 && m % s == 0 ? 1 : 0;
        }
        long mask = (long) m << 32 | i << 16 | s; // 三个 int 压缩成一个 long
        if (!limitLow && !limitHigh && memo.containsKey(mask)) {
            return memo.get(mask);
        }

        int diffLh = high.length - low.length;
        int lo = limitLow && i >= diffLh ? low[i - diffLh] - '0' : 0;
        int hi = limitHigh ? high[i] - '0' : 9;

        int res = 0;
        int d = lo;
        if (limitLow && i < diffLh) { // 利用 limitLow 和 i 可以判断出 isNum 是否为 true，所以 isNum 可以省略
            res = dfs(i + 1, 1, 0, true, false, low, high, memo); // 什么也不填
            d = 1; // 下面循环从 1 开始
        }
        for (; d <= hi; d++) {
            res += dfs(i + 1, m * d, s + d, limitLow && d == lo, limitHigh && d == hi, low, high, memo);
        }

        if (!limitLow && !limitHigh) {
            memo.put(mask, res);
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int beautifulNumbers(int l, int r) {
        string low = to_string(l), high = to_string(r);
        int n = high.size();
        int diff_lh = n - low.size(); // 这样写无需给 low 补前导零，也无需 is_num 参数
        unordered_map<long long, int> memo;

        auto dfs = [&](this auto&& dfs, int i, int m, int s, bool limit_low, bool limit_high) -> int {
            if (i == n) {
                return s && m % s == 0;
            }
            long long mask = (long long) m << 32 | i << 16 | s; // 三个 int 压缩成一个 long long
            if (!limit_low && !limit_high && memo.contains(mask)) {
                return memo[mask];
            }

            int lo = limit_low && i >= diff_lh ? low[i - diff_lh] - '0' : 0;
            int hi = limit_high ? high[i] - '0' : 9;

            int res = 0;
            int d = lo;
            if (limit_low && i < diff_lh) {
                res = dfs(i + 1, 1, 0, true, false); // 什么也不填
                d = 1; // 下面循环从 1 开始
            }
            for (; d <= hi; d++) {
                res += dfs(i + 1, m * d, s + d, limit_low && d == lo, limit_high && d == hi);
            }

            if (!limit_low && !limit_high) {
                memo[mask] = res;
            }
            return res;
        };

        return dfs(0, 1, 0, true, true);
    }
};
```

```go [sol-Go]
func beautifulNumbers(l, r int) int {
	low := strconv.Itoa(l)
	high := strconv.Itoa(r)
	n := len(high)
	diffLH := n - len(low) // 这样写无需给 low 补前导零，也无需 isNum 参数

	type tuple struct{ i, m, s int }
	memo := map[tuple]int{}
	var dfs func(int, int, int, bool, bool) int
	dfs = func(i, m, s int, limitLow, limitHigh bool) (res int) {
		if i == n {
			if s == 0 || m%s > 0 {
				return 0
			}
			return 1
		}
		if !limitLow && !limitHigh {
			t := tuple{i, m, s}
			if v, ok := memo[t]; ok {
				return v
			}
			defer func() { memo[t] = res }()
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(low[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(high[i] - '0')
		}

		d := lo
		if limitLow && i < diffLH {
			res = dfs(i+1, 1, 0, true, false) // 什么也不填
			d = 1 // 下面循环从 1 开始
		}
		// 枚举填数字 d
		for ; d <= hi; d++ {
			res += dfs(i+1, m*d, s+d, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}
	return dfs(0, 1, 0, true, true)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^6D^2)$，其中 $n$ 是 $r$ 的十进制长度，$D=10$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数为 $\mathcal{O}(n^6D)$，即 $i$ 的个数 $\mathcal{O}(n)$、$m$ 的个数 $\mathcal{O}(n^4)$（见文末的证明）、$s$ 的个数 $\mathcal{O}(nD)$ 三者相乘。单个状态的计算时间为 $\mathcal{O}(D)$，所以总的时间复杂度为 $\mathcal{O}(n^6D^2)$。
- 空间复杂度：$\mathcal{O}(n^6D)$。保存多少状态，就需要多少空间。

## 写法二

转换成 $[1,r]$ 的美丽整数个数，减去 $[1,l-1]$ 的美丽整数个数。

使用 [数位 DP v1.0 模板](https://www.bilibili.com/video/BV1rS4y1s721/?t=19m36s) 解决，这样可以把 $\textit{dfs}$ 写在外面，多个测试数据之间可以**复用**记忆化搜索的结果。

为了保证可以复用，需要把个位数的下标统一为 $0$，十位数的下标统一为 $1$，依此类推。

```py [sol-Python3]
memo = {}
high = []

def dfs(i: int, m: int, s: int, is_limit: bool, is_num: bool) -> int:
    if i < 0:
        return 1 if s and m % s == 0 else 0

    t = (m, i, s)
    if not is_limit and is_num and t in memo:
        return memo[t]

    # 什么也不填
    res = 0 if is_num else dfs(i - 1, m, s, False, False)
    hi = high[i] if is_limit else 9
    # 枚举填哪个数字
    for d in range(0 if is_num else 1, hi + 1):
        res += dfs(i - 1, m * d, s + d, is_limit and d == hi, True)

    if not is_limit and is_num:
        memo[t] = res
    return res

class Solution:
    def beautifulNumbers(self, l: int, r: int) -> int:
        def calc(r: int) -> int:
            global high
            high = list(map(int, str(r)))[::-1]
            return dfs(len(high) - 1, 1, 0, True, False)
        return calc(r) - calc(l - 1)
```

```java [sol-Java]
class Solution {
    private static final Map<Long, Integer> memo = new HashMap<>();

    public int beautifulNumbers(int l, int r) {
        return calc(r) - calc(l - 1);
    }

    private int calc(int r) {
        char[] high = new StringBuilder(String.valueOf(r)).reverse().toString().toCharArray();
        return dfs(high.length - 1, 1, 0, true, false, high);
    }

    private int dfs(int i, int m, int s, boolean isLimit, boolean isNum, char[] high) {
        if (i < 0) {
            return s > 0 && m % s == 0 ? 1 : 0;
        }

        long mask = (long) m << 32 | i << 16 | s;
        if (!isLimit && isNum && memo.containsKey(mask)) {
            return memo.get(mask);
        }

        // 什么也不填
        int res = isNum ? 0 : dfs(i - 1, m, s, false, false, high);
        int hi = isLimit ? high[i] - '0' : 9;
        // 枚举填哪个数字
        for (int d = isNum ? 0 : 1; d <= hi; d++) {
            res += dfs(i - 1, m * d, s + d, isLimit && d == hi, true, high);
        }

        if (!isLimit && isNum) {
            memo.put(mask, res);
        }
        return res;
    }
}
```

```cpp [sol-C++]
unordered_map<long long, int> memo;
string high;

int dfs(int i, int m, int s, bool is_limit, bool is_num) {
    if (i < 0) {
        return s && m % s == 0;
    }

    long long mask = (long long) m << 32 | i << 16 | s;
    if (!is_limit && is_num && memo.contains(mask)) {
        return memo[mask];
    }

    // 什么也不填
    int res = is_num ? 0 : dfs(i - 1, m, s, false, false);
    int hi = is_limit ? high[i] - '0' : 9;
    // 枚举填哪个数字
    for (int d = 1 - is_num; d <= hi; d++) {
        res += dfs(i - 1, m * d, s + d, is_limit && d == hi, true);
    }

    if (!is_limit && is_num) {
        memo[mask] = res;
    }
    return res;
}

class Solution {
    int calc(int r) {
        high = to_string(r);
        ranges::reverse(high);
        return dfs(high.size() - 1, 1, 0, true, false);
    }

public:
    int beautifulNumbers(int l, int r) {
        return calc(r) - calc(l - 1);
    }
};
```

```go [sol-Go]
type tuple struct{ i, m, s int }

var memo = map[tuple]int{}
var high []byte

func dfs(i, m, s int, isLimit, isNum bool) (res int) {
	if i < 0 {
		if s == 0 || m%s > 0 {
			return 0
		}
		return 1
	}
	if !isLimit && isNum {
		t := tuple{i, m, s}
		if v, ok := memo[t]; ok {
			return v
		}
		defer func() { memo[t] = res }()
	}

	hi := 9
	if isLimit {
		hi = int(high[i] - '0')
	}

	d := 0
	if !isNum {
		res = dfs(i-1, m, s, false, false) // 什么也不填
		d = 1
	}
	// 枚举填数字 d
	for ; d <= hi; d++ {
		res += dfs(i-1, m*d, s+d, isLimit && d == hi, true)
	}
	return
}

func calc(r int) int {
	high = []byte(strconv.Itoa(r))
	slices.Reverse(high)
	return dfs(len(high)-1, 1, 0, true, false)
}

func beautifulNumbers(l, r int) int {
	return calc(r) - calc(l-1)
}
```

#### 复杂度分析

同写法一。

## 有多少个不同乘积

定义 $K(n)$ 表示 $n$ 个在 $[0,9]$ 中的整数相乘，有多少个不同的乘积。

如果有 $0$，那么乘积为 $0$。

下面讨论 $n$ 个在 $[1,9]$ 中的整数相乘，有多少个不同的乘积。

考虑乘积的质因子分解，即

$$
2^a3^b5^c7^d
$$

假设我们选了 $k$ 个在 $\{1,2,3,4,6,8,9\}$ 中的数，那么有 $n-k$ 个 $5$ 或 $7$，即 $c+d=n-k$。知道了 $c$ 的个数就知道了 $d$ 的个数，$c$ 的范围为 $[0,n-k]$，所以 $(c,d)$ 有 $n-k+1$ 个。

下面讨论在 $\{1,2,3,4,6,8,9\}$ 中（允许重复地）选 $k$ 个数，有多少个不同的 $2^a3^b$。

先不考虑数字 $6$。设我们选了 $m$ 个 $\{1,2,4,8\}$ 中的数：

- 如果 $m=k$，那么 $a$ 的范围为 $[0,3k]$，$b=0$，这有 $3k+1$ 个 $(a,b)$。
- 如果 $m=k-1$，那么 $a$ 的范围为 $[0,3(k-1)]$，$b$ 的范围为 $[1,2]$（选 $3$ 或者 $9$），这新增了 $(3(k-1)+1)\cdot 2$ 个 $(a,b)$。
- 如果 $m=k-2$，那么 $a$ 的范围为 $[0,3(k-2)]$，$b$ 的范围为 $[2,4]$，其中 $b=2$ 的情况之前算过，所以新增了 $(3(k-2)+1)\cdot 2$ 个 $(a,b)$。
- 如果 $m=k-3$，那么 $a$ 的范围为 $[0,3(k-3)]$，$b$ 的范围为 $[3,6]$，其中 $b=3,4$ 的情况之前算过，所以新增了 $(3(k-3)+1)\cdot 2$ 个 $(a,b)$。
- 依此类推，每次会新增 $(3m+1)\cdot 2$ 个 $(a,b)$，一直到 $m=0$。

累加得

$$
(3k+1) + \sum_{m=0}^{k-1} (3m+1)\cdot 2 = 3k^2+2k+1
$$

然后考虑数字 $6$，设我们选了 $m$ 个 $\{1,2,4,8\}$ 中的数：

- 如果 $m=k-1$，那么 $a$ 的范围为 $[0,3(k-1)]$，$b=0$（因为得留一个数出来选 $6$）。只有当 $a=3(k-1)$ 的时候，和 $6$ 相乘才会产生 $1$ 个新的 $(a,b)$。
- 如果 $m=k-2$，那么 $a$ 的范围为 $[0,3(k-2)]$，如果选两个 $6$，这不会产生新的 $(a,b)$。只有选一个 $6$，且 $a=3(k-2),\ b=2$ 的时候，和 $6$ 相乘才会产生 $1$ 个新的 $(a,b)$。
- 依此类推，对于 $[0,k-1]$ 中的每个 $m$，都会产生 $1$ 个新的 $(a,b)$，一共产生了 $k$ 个新的 $(a,b)$。

所以一共有

$$
3k^2+3k+1
$$

个不同的 $2^a3^b$。

枚举 $k$，得到 $2^a3^b5^c7^d$ 的个数为

$$
\sum_{k=0}^{n} (n-k+1)(3k^2+3k+1) = \left[\dfrac{(n+1)(n+2)}{2}\right]^2
$$

> 巧合的是，上式刚好是 $[1,n+1]$ 的立方和。例如 $n=8$ 的时候，结果是今年的年份 $2025$。

所以

$$
K(n) = \left[\dfrac{(n+1)(n+2)}{2}\right]^2 + 1
$$

## 附：如何用代码计算

```py
st = {1}  # 空集的乘积（乘法单位元）
for _ in range(9):  # 9 个数相乘
    st = set(x * d for x in st for d in range(10))  # 每个数从 0 到 9
print(len(st))  # 3026
```

更多相似题目，见下面动态规划题单中的「**十、数位 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
