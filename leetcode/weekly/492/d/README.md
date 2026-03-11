写个分治：

- 定义 $\textit{dfs}(\ell,r)$ 表示子串 $[\ell,r)$ 的最小费用。
- 计算子串 $[\ell,r)$ 中有多少个 $\texttt{1}$，从而算出不拆分子串时的费用。
- 如果子串 $[\ell,r)$ 的长度 $r-\ell$ 是偶数，可以拆分，对应的费用为 $\textit{dfs}(\ell,m) + \textit{dfs}(m,r)$，其中 $m = \left\lfloor\dfrac{\ell+r}{2}\right\rfloor$。
- 两种情况取最小值，即为 $\textit{dfs}(\ell,r)$ 的返回值。

递归入口：$\textit{dfs}(0,n)$。

递归边界：由于 $1$ 是奇数，所以递归边界已经蕴含在 $r-\ell$ 是奇数的情况中了。

代码实现时，子串有多少个 $\texttt{1}$ 可以用**前缀和**快速计算。原理请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

[本题视频讲解](https://www.bilibili.com/video/BV1H6NMzdEbo/?t=47m48s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def minCost(self, s: str, encCost: int, flatCost: int) -> int:
        n = len(s)
        pre = list(accumulate(map(int, s), initial=0))

        # 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
        def dfs(l: int, r: int) -> int:
            # 不拆分
            x = pre[r] - pre[l]
            res = (r - l) * x * encCost if x else flatCost

            # 拆分
            if (r - l) % 2 == 0:
                m = (l + r) // 2
                res = min(res, dfs(l, m) + dfs(m, r))

            return res

        return dfs(0, n)
```

```java [sol-Java]
class Solution {
    public long minCost(String s, int encCost, int flatCost) {
        int n = s.length();
        int[] sum = new int[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (s.charAt(i) - '0');
        }
        return dfs(0, n, sum, encCost, flatCost);
    }

    // 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
    private long dfs(int l, int r, int[] sum, int encCost, int flatCost) {
        // 不拆分
        int x = sum[r] - sum[l];
        long res = x > 0 ? (long) (r - l) * x * encCost : flatCost;

        // 拆分
        if ((r - l) % 2 == 0) {
            int m = (l + r) / 2;
            res = Math.min(res, dfs(l, m, sum, encCost, flatCost) + dfs(m, r, sum, encCost, flatCost));
        }

        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(string s, int encCost, int flatCost) {
        int n = s.size();
        vector<int> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (s[i] - '0');
        }

        // 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
        auto dfs = [&](this auto&& dfs, int l, int r) -> long long {
            // 不拆分
            int x = sum[r] - sum[l];
            long long res = x ? 1LL * (r - l) * x * encCost : flatCost;

            // 拆分
            if ((r - l) % 2 == 0) {
                int m = (l + r) / 2;
                res = min(res, dfs(l, m) + dfs(m, r));
            }

            return res;
        };

        return dfs(0, n);
    }
};
```

```go [sol-Go]
func minCost(s string, encCost, flatCost int) int64 {
	n := len(s)
	sum := make([]int, n+1)
	for i, ch := range s {
		sum[i+1] = sum[i] + int(ch-'0')
	}

	// 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
	var dfs func(int, int) int
	dfs = func(l, r int) int {
		// 不拆分
		res := flatCost
		if x := sum[r] - sum[l]; x > 0 {
			res = (r - l) * x * encCost
		}

		// 拆分
		if (r-l)%2 == 0 {
			m := (l + r) / 2
			res = min(res, dfs(l, m)+dfs(m, r))
		}

		return res
	}
	return int64(dfs(0, n))
}
```

## 优化

当 $x = 0$ 时，由于拆分后，子串的 $x$ 仍然是 $0$，所以费用是 $2\cdot \textit{flatCost}$，大于不拆分的费用 $\textit{flatCost}$。如果继续拆分，费用只会越来越大，都会大于不拆分的费用 $\textit{flatCost}$。

所以当 $x=0$ 时，直接返回 $\textit{flatCost}$ 即可。

```py [sol-Python3]
class Solution:
    def minCost(self, s: str, encCost: int, flatCost: int) -> int:
        n = len(s)
        pre = list(accumulate(map(int, s), initial=0))

        # 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
        def dfs(l: int, r: int) -> int:
            x = pre[r] - pre[l]
            if x == 0:
                return flatCost           

            # 不拆分
            res = (r - l) * x * encCost

            # 拆分
            if (r - l) % 2 == 0:
                m = (l + r) // 2
                res = min(res, dfs(l, m) + dfs(m, r))

            return res

        return dfs(0, n)
```

```java [sol-Java]
class Solution {
    public long minCost(String s, int encCost, int flatCost) {
        int n = s.length();
        int[] sum = new int[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (s.charAt(i) - '0');
        }
        return dfs(0, n, sum, encCost, flatCost);
    }

    // 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
    private long dfs(int l, int r, int[] sum, int encCost, int flatCost) {
        int x = sum[r] - sum[l];
        if (x == 0) {
            return flatCost;
        }

        // 不拆分
        long res = (long) (r - l) * x * encCost;

        // 拆分
        if ((r - l) % 2 == 0) {
            int m = (l + r) / 2;
            res = Math.min(res, dfs(l, m, sum, encCost, flatCost) + dfs(m, r, sum, encCost, flatCost));
        }

        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(string s, int encCost, int flatCost) {
        int n = s.size();
        vector<int> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (s[i] - '0');
        }

        // 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
        auto dfs = [&](this auto&& dfs, int l, int r) -> long long {
            int x = sum[r] - sum[l];
            if (x == 0) {
                return flatCost;
            }

            // 不拆分
            long long res = 1LL * (r - l) * x * encCost;

            // 拆分
            if ((r - l) % 2 == 0) {
                int m = (l + r) / 2;
                res = min(res, dfs(l, m) + dfs(m, r));
            }

            return res;
        };

        return dfs(0, n);
    }
};
```

```go [sol-Go]
func minCost(s string, encCost, flatCost int) int64 {
	n := len(s)
	sum := make([]int, n+1)
	for i, ch := range s {
		sum[i+1] = sum[i] + int(ch-'0')
	}

	// 计算子串 [l, r) 的最小费用，注意区间是左闭右开，方便计算
	var dfs func(int, int) int
	dfs = func(l, r int) int {
		x := sum[r] - sum[l]
		if x == 0 {
			return flatCost
		}

		// 不拆分
		res := (r - l) * x * encCost

		// 拆分
		if (r-l)%2 == 0 {
			m := (l + r) / 2
			res = min(res, dfs(l, m)+dfs(m, r))
		}

		return res
	}
	return int64(dfs(0, n))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。搜索树是一棵高为 $\mathcal{O}(\log n)$ 的二叉树，从上到下，每一层至多有 $1,2,4,8,16,\ldots,n$ 个节点，根据等比数列求和公式，搜索树一共有 $\mathcal{O}(n)$ 个节点。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
