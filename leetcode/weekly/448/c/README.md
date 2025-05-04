**关键思路**：一段连续的合并操作执行完后，合并的时间是 $\textit{time}$ 的一个连续**子数组**，我们需要知道子数组的左端点，方便计算合并后的时间。

从左到右模拟旅行的过程。我们需要知道如下信息：

- 还剩下 $\textit{leftK}$ 次合并操作可以用。
- 当前在 $\textit{position}[i]$。
- 合并到 $\textit{time}[i]$ 的这段时间的左端点为 $\textit{pre}$。

定义 $\textit{dfs}(\textit{leftK},i,\textit{pre})$ 表示在上述情况下，完成剩余旅程需要的最小时间。

⚠**注意**：每段路程的耗时是两部分的乘积：合并到 $\textit{time}[i]$ 的时间，当前位置到下一个位置的距离。**这两个数据相对 $i$ 是一左一右的关系，并不是都在 $i$ 的右边**！

枚举合并后，下一个位置的下标 $\textit{nxt}=i+1,i+2,\ldots, \min(n-1, i+1+\textit{leftK})$，问题变成：

- 还剩下 $\textit{leftK} - (\textit{nxt}-i-1)$ 次合并操作可以用。
- 当前在 $\textit{position}[\textit{nxt}]$。
- 合并到 $\textit{time}[\textit{nxt}]$ 的这段时间的左端点为 $i+1$。注意我们删除的时间下标范围是 $[i+1, \textit{nxt}-1]$，这段时间合并到 $\textit{time}[\textit{nxt}]$ 中。所以合并后的下标范围为 $[i+1, \textit{nxt}]$。

子问题为 $\textit{dfs}(\textit{leftK} - (\textit{nxt}-i-1), \textit{nxt}, i+1)$。

设 $s$ 为 $\textit{time}$ 的 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)，对于当前的 $\textit{time}[i]$ 来说，$[\textit{pre},i]$ 合并之后，$\textit{time}[i]$ 变成 $s[i+1] - s[\textit{pre}]$。从 $i$ 到 $\textit{nxt}$，花费的时间为 $(\textit{position}[\textit{nxt}] - \textit{position}[i])\cdot (s[i+1] - s[\textit{pre}])$。

取最小值，得

$$
\textit{dfs}(\textit{leftK},i,\textit{pre}) = \min_{\textit{nxt}=i+1}^{\min(n-1, i+1+\textit{leftK})} \textit{dfs}(\textit{leftK} - (\textit{nxt}-i-1), \textit{nxt}, i+1) + (\textit{position}[\textit{nxt}] - \textit{position}[i])\cdot (s[i+1] - s[\textit{pre}])
$$

递归边界：$\textit{dfs}(0,n-1,\textit{pre})=0$，其余 $\textit{dfs}(\textit{leftK},n-1,\textit{pre})=\infty$。

递归入口：$\textit{dfs}(k,0,0)$，即答案。

## 写法一：记忆化搜索

原理见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

[本题视频讲解](https://www.bilibili.com/video/BV1avVwz5EbY/?t=12m34s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minTravelTime(self, l: int, n: int, k: int, position: List[int], time: List[int]) -> int:
        s = list(accumulate(time, initial=0))
        @cache
        def dfs(left_k: int, i: int, pre: int) -> int:
            if i == n - 1:
                return inf if left_k else 0
            t = s[i + 1] - s[pre]
            return min(dfs(left_k - (nxt - i - 1), nxt, i + 1) + (position[nxt] - position[i]) * t
                       for nxt in range(i + 1, min(n, i + 2 + left_k)))
        return dfs(k, 0, 0)
```

```java [sol-Java]
class Solution {
    public int minTravelTime(int l, int n, int k, int[] position, int[] time) {
        int[] s = new int[n];
        for (int i = 0; i < n - 1; i++) {
            s[i + 1] = s[i] + time[i];
        }

        int[][][] memo = new int[k + 1][n - 1][n - 1];
        return dfs(k, 0, 0, position, s, memo);
    }

    private int dfs(int leftK, int i, int pre, int[] position, int[] s, int[][][] memo) {
        int n = position.length;
        if (i == n - 1) {
            return leftK > 0 ? Integer.MAX_VALUE / 2 : 0;
        }
        if (memo[leftK][i][pre] > 0) {
            return memo[leftK][i][pre];
        }
        int res = Integer.MAX_VALUE;
        int t = s[i + 1] - s[pre];
        for (int nxt = i + 1; nxt < Math.min(n, i + 2 + leftK); nxt++) {
            int r = dfs(leftK - (nxt - i - 1), nxt, i + 1, position, s, memo) + (position[nxt] - position[i]) * t;
            res = Math.min(res, r);
        }
        return memo[leftK][i][pre] = res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minTravelTime(int l, int n, int k, vector<int>& position, vector<int>& time) {
        vector<int> s(n);
        partial_sum(time.begin(), time.end() - 1, s.begin() + 1);

        vector memo(k + 1, vector(n - 1, vector<int>(n - 1)));
        auto dfs = [&](this auto&& dfs, int left_k, int i, int pre) -> int {
            if (i == n - 1) {
                return left_k ? INT_MAX / 2 : 0;
            }
            int& res = memo[left_k][i][pre]; // 注意这里是引用
            if (res) {
                return res;
            }
            res = INT_MAX;
            int t = s[i + 1] - s[pre];
            for (int nxt = i + 1; nxt < min(n, i + 2 + left_k); nxt++) {
                res = min(res, dfs(left_k - (nxt - i - 1), nxt, i + 1) + (position[nxt] - position[i]) * t);
            }
            return res;
        };
        return dfs(k, 0, 0);
    }
};
```

```go [sol-Go]
func minTravelTime(l, n, k int, position, time []int) int {
	s := make([]int, n)
	for i, t := range time[:n-1] {
		s[i+1] = s[i] + t
	}

	memo := make([][][]int, k+1)
	for i := range memo {
		memo[i] = make([][]int, n-1)
		for j := range memo[i] {
			memo[i][j] = make([]int, n-1)
		}
	}
	var dfs func(int, int, int) int
	dfs = func(leftK, i, pre int) int {
		if i == n-1 {
			if leftK > 0 {
				return math.MaxInt / 2
			}
			return 0
		}
		p := &memo[leftK][i][pre]
		if *p > 0 {
			return *p
		}
		res := math.MaxInt
		t := s[i+1] - s[pre]
		for nxt := i + 1; nxt < min(n, i+2+leftK); nxt++ {
			res = min(res, dfs(leftK-(nxt-i-1), nxt, i+1)+(position[nxt]-position[i])*t)
		}
		*p = res
		return res
	}
	return dfs(k, 0, 0)
}
```

## 写法二：递推

```py [sol-Python3]
class Solution:
    def minTravelTime(self, l: int, n: int, k: int, position: List[int], time: List[int]) -> int:
        s = list(accumulate(time, initial=0))
        f = [[[inf] * n for _ in range(n)] for _ in range(k + 1)]
        f[0][-1] = [0] * n
        for left_k in range(k + 1):
            for i in range(n - 2, -1, -1):
                for pre in range(i + 1):
                    t = s[i + 1] - s[pre]
                    f[left_k][i][pre] = min(f[left_k - (nxt - i - 1)][nxt][i + 1] + (position[nxt] - position[i]) * t
                                            for nxt in range(i + 1, min(n, i + 2 + left_k)))
        return f[k][0][0]
```

```java [sol-Java]
class Solution {
    public int minTravelTime(int l, int n, int k, int[] position, int[] time) {
        int[] s = new int[n];
        for (int i = 0; i < n - 1; i++) {
            s[i + 1] = s[i] + time[i];
        }

        int[][][] f = new int[k + 1][n][n];
        for (int leftK = 1; leftK <= k; leftK++) {
            Arrays.fill(f[leftK][n - 1], Integer.MAX_VALUE / 2);
        }
        for (int leftK = 0; leftK <= k; leftK++) {
            for (int i = n - 2; i >= 0; i--) {
                for (int pre = 0; pre <= i; pre++) {
                    int res = Integer.MAX_VALUE;
                    int t = s[i + 1] - s[pre];
                    for (int nxt = i + 1; nxt < Math.min(n, i + 2 + leftK); nxt++) {
                        int r = f[leftK - (nxt - i - 1)][nxt][i + 1] + (position[nxt] - position[i]) * t;
                        res = Math.min(res, r);
                    }
                    f[leftK][i][pre] = res;
                }
            }
        }
        return f[k][0][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minTravelTime(int l, int n, int k, vector<int>& position, vector<int>& time) {
        vector<int> s(n);
        partial_sum(time.begin(), time.end() - 1, s.begin() + 1);

        vector f(k + 1, vector(n, vector<int>(n, INT_MAX / 2)));
        for (int pre = 0; pre < n; pre++) {
            f[0][n - 1][pre] = 0;
        }
        for (int left_k = 0; left_k <= k; left_k++) {
            for (int i = n - 2; i >= 0; i--) {
                for (int pre = 0; pre <= i; pre++) {
                    int res = INT_MAX;
                    int t = s[i + 1] - s[pre];
                    for (int nxt = i + 1; nxt < min(n, i + 2 + left_k); nxt++) {
                        res = min(res, f[left_k - (nxt - i - 1)][nxt][i + 1] + (position[nxt] - position[i]) * t);
                    }
                    f[left_k][i][pre] = res;
                }
            }
        }
        return f[k][0][0];
    }
};
```

```go [sol-Go]
func minTravelTime(l, n, k int, position, time []int) int {
	s := make([]int, n)
	for i, t := range time[:n-1] {
		s[i+1] = s[i] + t
	}

	f := make([][][]int, k+1)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
		}
	}
	for leftK := 1; leftK <= k; leftK++ {
		for pre := range n {
			f[leftK][n-1][pre] = math.MaxInt / 2
		}
	}

	for leftK := range f {
		for i := n - 2; i >= 0; i-- {
			for pre := range i + 1 {
				t := s[i+1] - s[pre]
				res := math.MaxInt
				for nxt := i + 1; nxt < min(n, i+2+leftK); nxt++ {
					res = min(res, f[leftK-(nxt-i-1)][nxt][i+1]+(position[nxt]-position[i])*t)
				}
				f[leftK][i][pre] = res
			}
		}
	}
	return f[k][0][0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(kn^3)$。
- 空间复杂度：$\mathcal{O}(kn^2)$。

更多相似题目，见下面动态规划题单的「**§5.3 约束划分个数**」。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
