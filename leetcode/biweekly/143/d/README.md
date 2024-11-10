首先，由于数位乘积中的质因子只有 $2,3,5,7$，如果 $t$ 包含其他质因子，那么直接返回 $-1$。如果 $t$ 只包含质因子 $2,3,5,7$，那么答案一定存在。

下文把 $\textit{num}$ 简记为 $s$。其长度为 $n$。

例如 $s=123$，并假设答案的长度也为 $3$。

仿照 [数位 DP](https://www.bilibili.com/video/BV1rS4y1s721/?t=20m05s) 的思路，写一个爆搜（回溯）：

- 最高位如果填 $1$，那么第二位不能填 $1$（不然小于 $s$ 了），至少要填 $2$。
- 最高位如果填的数大于 $1$，那么第二位，以及第三位，填的数字不受到 $s$ 的约束，可以填 $[1,9]$ 中的任意数字。

这启发我们也像数位 DP 那样，在回溯的过程中，用一个参数 $\textit{isLimit}$ 表示「是否受到 $s$ 的约束」。

如何判断所填数位之积是 $t$ 的倍数呢？

比如 $t=10$：

- 如果最高位填 $2$，那么后面两位所填数字，只需满足乘积是 $t/2=5$ 的倍数。
- 如果最高位填 $5$，那么后面两位所填数字，只需满足乘积是 $t/5=2$ 的倍数。
- 如果最高位填 $6$，由于 $6$ 中有因子 $2$，那么后面两位所填数字，只需满足乘积是 $t/2=5$ 的倍数。

一般地，如果填的数字是 $d$，那么余下的数位，需要满足乘积是 $\dfrac{t}{\text{GCD}(t,d)}$ 的倍数。

综上所述，写一个带 $\textit{vis}$ 的爆搜（回溯），参数有：

- $i$：表示当前填到 $s$ 的第 $i$ 个数位了。
- $t$：表示 $[i,n-1]$ 所填数位之积必须是 $t$ 的倍数。
- $\textit{isLimit}$：表示是否受到 $s$ 的约束。如果为 $\texttt{false}$，那么当前数位可以填 $[1,9]$ 中的数；如果为 $\texttt{true}$，那么当前数位只能填 $[\max(s[i],1),9]$ 中的数。这里和 $1$ 取最大值是因为 $s[i]$ 可能为 $0$，但我们不能填 $0$。在受到 $s$ 约束的情况下，如果填的数字为 $s[i]$，那么后面仍然会受到 $s$ 的约束。

递归边界：当 $i=n$ 时，如果 $t=1$，说明当前填法是符合要求的，返回 $\texttt{true}$，否则返回 $\texttt{false}$。

递归入口：$\textit{dfs}(0,t,\texttt{true})$。从最高位开始填，一开始受到 $s$ 的约束。

如果下面的 DFS 返回 $\texttt{true}$，说明找到了答案，也直接返回 $\texttt{true}$。

### 细节

如果 $s$ 很短但 $t$ 很大，答案是会比 $s$ 还长的。

为了兼顾这种情况，我们可以往 $s$ 的前面添加

$$
\max(\textit{cnt}-n+1,1)
$$

个前导 $0$。其中 $\textit{cnt}$ 是 $t$ 的质因子个数。

注意至少要添加 $1$ 个前导零，因为可能有 $s=999$ 这种情况，即使 $t=2$，答案（$1112$）长度也比 $s$ 要长。

注意添加前导零会影响可以填入的数字，当 $\textit{isLimit}=\texttt{true}$ 且 $i < \textit{cnt}$ 时，我们可以填入 $0$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1cgmBYqEhu/?t=28m31s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def smallestNumber(self, s: str, t: int) -> str:
        tmp = t
        cnt = 0
        for p in 2, 3, 5, 7:
            while tmp % p == 0:
                tmp //= p
                cnt += 1
        if tmp > 1:  # t 包含其他质因子
            return "-1"

        # 补前导零（至少一个）
        cnt = max(cnt - len(s) + 1, 1)
        s = '0' * cnt + s

        n = len(s)
        ans = [0] * n

        @cache  # 仅仅作为 vis 标记使用
        def dfs(i: int, t: int, is_limit: bool) -> bool:
            if i == n:
                return t == 1

            x = int(s[i])
            # 如果没有约束，那么 1~9 随便填（注意这意味着前面填了大于 0 的数）
            low = x if is_limit and (x or i < cnt) else 1
            for d in range(low, 10):
                ans[i] = d  # 直接覆盖，无需恢复现场
                new_t = t // gcd(t, d) if d > 1 else t
                if dfs(i + 1, new_t, is_limit and d == x):
                    return True
            return False

        dfs(0, t, True)
        dfs.cache_clear()  # 防止爆内存
        return ''.join(map(str, ans)).lstrip('0')  # 去掉前导零
```

```py [sol-Python3 写法二]
class Solution:
    def smallestNumber(self, s: str, t: int) -> str:
        tmp = t
        cnt = 0
        for p in 2, 3, 5, 7:
            while tmp % p == 0:
                tmp //= p
                cnt += 1
        if tmp > 1:  # t 包含其他质因子
            return "-1"

        # 补前导零（至少一个）
        cnt = max(cnt - len(s) + 1, 1)
        s = '0' * cnt + s

        n = len(s)
        ans = [0] * n
        vis = [set() for _ in range(n)]

        def dfs(i: int, t: int, is_limit: bool) -> bool:
            if i == n:
                return t == 1

            if not is_limit:
                if t in vis[i]:
                    return False
                vis[i].add(t)

            x = int(s[i])
            # 如果没有约束，那么 1~9 随便填（注意这意味着前面填了大于 0 的数）
            low = x if is_limit and (x or i < cnt) else 1
            for d in range(low, 10):
                ans[i] = d  # 直接覆盖，无需恢复现场
                new_t = t // gcd(t, d) if d > 1 else t
                if dfs(i + 1, new_t, is_limit and d == x):
                    return True
            return False

        dfs(0, t, True)
        return ''.join(map(str, ans)).lstrip('0')  # 去掉前导零
```

```java [sol-Java]
class Solution {
    public String smallestNumber(String s, long t) {
        long tmp = t;
        int cnt = 0;
        for (int p : new int[]{2, 3, 5, 7}) {
            while (tmp % p == 0) {
                tmp /= p;
                cnt++;
            }
        }
        if (tmp > 1) { // t 包含其他质因子
            return "-1";
        }

        // 补前导零（至少一个）
        cnt = Math.max(cnt - s.length() + 1, 1);
        s = "0".repeat(cnt) + s;

        int n = s.length();
        char[] ans = new char[n];
        Set<Long>[] vis = new HashSet[n];
        Arrays.setAll(vis, i -> new HashSet<>());
        dfs(0, t, true, cnt, s.toCharArray(), ans, vis);
        for (int i = 0; ; i++) {
            if (ans[i] != '0') {
                return new String(ans, i, n - i); // 去掉前导零
            }
        }
    }

    private boolean dfs(int i, long t, boolean isLimit, int cnt, char[] s, char[] ans, Set<Long>[] vis) {
        if (i == s.length) {
            return t == 1;
        }
        if (!isLimit && !vis[i].add(t)) {
            return false;
        }

        int x = s[i] - '0';
        // 如果没有约束，那么 1~9 随便填（注意这意味着前面填了大于 0 的数）
        int low = isLimit && (x > 0 || i < cnt) ? x : 1;
        for (int d = low; d <= 9; d++) {
            ans[i] = (char) ('0' + d); // 直接覆盖，无需恢复现场
            long newT = d > 1 ? t / gcd(t, d) : t;
            if (dfs(i + 1, newT, isLimit && d == x, cnt, s, ans, vis)) {
                return true;
            }
        }
        return false;
    }

    private long gcd(long a, long b) {
        while (a != 0) {
            long tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
// 超时了，代码仅供参考
class Solution {
public:
    string smallestNumber(string s, long long t) {
        long long tmp = t;
        int cnt = 0;
        for (int p : {2, 3, 5, 7}) {
            while (tmp % p == 0) {
                tmp /= p;
                cnt++;
            }
        }
        if (tmp > 1) {
            return "-1";
        }

        cnt = max(cnt - (int) s.length() + 1, 1);
        s = string(cnt, '0') + s;

        int n = s.length();
        string ans(n, 0);
        vector<unordered_set<long long>> vis(n);
        auto dfs = [&](auto&& dfs, int i, long long t, bool is_limit) -> bool {
            if (i == n) {
                return t == 1;
            }
            if (!is_limit && !vis[i].insert(t).second) {
                return false;
            }

            int x = s[i] - '0';
            int low = is_limit && (x > 0 || i < cnt) ? x : 1;

            for (int d = low; d <= 9; d++) {
                ans[i] = '0' + d; // 直接覆盖，无需恢复现场
                long long new_t = d > 1 ? t / gcd(t, d) : t;
                if (dfs(dfs, i + 1, new_t, is_limit && d == x)) {
                    return true;
                }
            }
            return false;
        };
        dfs(dfs, 0, t, true);

        auto it = ranges::find_if(ans, [](char c) { return c != '0'; });
        return string(it, ans.end()); // 去掉前导零
    }
};
```

```go [sol-Go]
func smallestNumber(s string, t int64) string {
	tmp, cnt := int(t), 0
	for _, p := range []int{2, 3, 5, 7} {
		for tmp%p == 0 {
			tmp /= p
			cnt++
		}
	}
	if tmp > 1 { // t 包含其他质因子
		return "-1"
	}

	// 补前导零（至少一个）
	cnt = max(cnt-len(s)+1, 1)
	s = strings.Repeat("0", cnt) + s

	n := len(s)
	ans := make([]byte, len(s))
	type pair struct{ i, t int }
	vis := map[pair]bool{}

	var dfs func(int, int, bool) bool
	dfs = func(i, t int, isLimit bool) bool {
		if i == n {
			return t == 1
		}
		if !isLimit {
			p := pair{i, t}
			if vis[p] {
				return false
			}
			vis[p] = true
		}

		x := int(s[i] - '0')
		low := 1 // 如果没有约束，那么 1~9 随便填（注意这意味着前面填了大于 0 的数）
		if isLimit && (x > 0 || i < cnt) {
			low = x
		}
		for d := low; d <= 9; d++ {
			ans[i] = '0' + byte(d) // 直接覆盖，无需恢复现场
			newT := t
			if d > 1 {
				newT = t / gcd(t, d)
			}
			if dfs(i+1, newT, isLimit && d == x) {
				return true
			}
		}
		return false
	}
	dfs(0, int(t), true)

	i := bytes.LastIndexByte(ans, '0')
	return string(ans[i+1:])
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log t)$，其中 $n$ 是 $s$ 的长度。注意在 DFS 中，只有 $\mathcal{O}(\log t)$ 个不同的 $t$。
- 空间复杂度：$\mathcal{O}(n\log t)$。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
