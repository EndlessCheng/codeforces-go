## 分析

举个例子，如果最左两个位置和最右两个位置分别填 $3$ 和 $7$，我们会得到一个 $37000\cdots00073$ 的数字（中间的 $0$ 表示还没有填入数字），假设这个数模 $k$ 的结果为 $1$。如果最左最右两个位置分别填 $9$ 和 $5$，我们会得到一个 $95000\cdots00059$ 的数字（中间的 $0$ 表示还没有填入数字），假设这个数模 $k$ 的结果也为 $1$。

无论后面的数字怎么填，这两种情况其实是一样的，因为我们只关心回文数**最终**模 $k$ 的值是否为 $0$。

把「当前从右到左填到第 $i$ 位」和「已填入的数字模 $k$ 的值为 $j$」作为状态 $(i,j)$。如果我们从状态 $(i,j)$ 开始 DFS，无论后面怎么填，回文数**最终**模 $k$ 的值不为 $0$，那么当我们第二次 DFS 到状态 $(i,j)$ 时，根据上面的例子，我们可以直接得出，无论后面怎么填，回文数**最终**模 $k$ 的值仍然不会为 $0$。

这启发我们用图上的 DFS 思考。在 DFS 中的过程中，不访问之前访问过的点。

## 建图

枚举在第 $i$ 位填入数字 $d$（从大到小枚举 $d$），那么同时也在第 $n-1-i$ 位填入了 $d$。

填入数字后，回文数模 $k$ 的值变成了

$$
j_2 = (j + d\cdot (10^i + 10^{n-1-i}))\bmod k
$$

我们可以在 $(i,j)$ 和 $(i+1, j_2)$ 之间连边，得到一个有向图。

一开始什么数也没填，所以 $j=0$；最终模 $k$ 要等于 $0$，所以 $j$ 也等于 $0$。所以答案是一条从起点 $(0,0)$ 到终点 $(m,0)$ 的**字典序最大路径**，其中 $m=\left\lceil\dfrac{n}{2}\right\rceil$，因为我们只需填一半的数字，另一半可以镜像得到。

用 DFS 搜索，每次从 $d=9$ 开始倒着枚举，即可得到字典序最大路径。

注意特判 $n$ 为奇数且 $i=m-1$ 的情况，此时填入数字后，回文数模 $k$ 的值变成了

$$
(j + d\cdot 10^i)\bmod k
$$

代码实现时，可以预处理 $10^i \bmod k$。

为什么可以在中途取模，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

本题 [视频讲解](https://www.bilibili.com/video/BV1hH4y1c7T5/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def largestPalindrome(self, n: int, k: int) -> str:
        pow10 = [1] * n
        for i in range(1, n):
            pow10[i] = pow10[i - 1] * 10 % k

        ans = [''] * n
        m = (n + 1) // 2
        vis = [[False] * k for _ in range(m + 1)]
        def dfs(i: int, j: int) -> bool:
            if i == m:
                return j == 0
            vis[i][j] = True
            for d in range(9, -1, -1):  # 贪心：从大到小枚举
                if n % 2 and i == m - 1:  # 正中间
                    j2 = (j + d * pow10[i]) % k
                else:
                    j2 = (j + d * (pow10[i] + pow10[-1 - i])) % k
                if not vis[i + 1][j2] and dfs(i + 1, j2):
                    ans[i] = ans[-1 - i] = str(d)
                    return True
            return False
        dfs(0, 0)
        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String largestPalindrome(int n, int k) {
        int[] pow10 = new int[n];
        pow10[0] = 1;
        for (int i = 1; i < n; i++) {
            pow10[i] = pow10[i - 1] * 10 % k;
        }

        char[] ans = new char[n];
        int m = (n + 1) / 2;
        boolean[][] vis = new boolean[m + 1][k];
        dfs(0, 0, n, k, m, pow10, ans, vis);
        return new String(ans);
    }

    private boolean dfs(int i, int j, int n, int k, int m, int[] pow10, char[] ans, boolean[][] vis) {
        if (i == m) {
            return j == 0;
        }
        vis[i][j] = true;
        for (int d = 9; d >= 0; d--) { // 贪心：从大到小枚举
            int j2;
            if (n % 2 > 0 && i == m - 1) { // 正中间
                j2 = (j + d * pow10[i]) % k;
            } else {
                j2 = (j + d * (pow10[i] + pow10[n - 1 - i])) % k;
            }
            if (!vis[i + 1][j2] && dfs(i + 1, j2, n, k, m, pow10, ans, vis)) {
                ans[i] = ans[n - 1 - i] = (char) ('0' + d);
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string largestPalindrome(int n, int k) {
        vector<int> pow10(n);
        pow10[0] = 1;
        for (int i = 1; i < n; i++) {
            pow10[i] = pow10[i - 1] * 10 % k;
        }

        string ans(n, 0);
        int m = (n + 1) / 2;
        vector<vector<int>> vis(m + 1, vector<int>(k));
        auto dfs = [&](auto&& dfs, int i, int j) -> bool {
            if (i == m) {
                return j == 0;
            }
            vis[i][j] = true;
            for (int d = 9; d >= 0; d--) { // 贪心：从大到小枚举
                int j2;
                if (n % 2 && i == m - 1) { // 正中间
                    j2 = (j + d * pow10[i]) % k;
                } else {
                    j2 = (j + d * (pow10[i] + pow10[n - 1 - i])) % k;
                }
                if (!vis[i + 1][j2] && dfs(dfs, i + 1, j2)) {
                    ans[i] = ans[n - 1 - i] = '0' + d;
                    return true;
                }
            }
            return false;
        };
        dfs(dfs, 0, 0);
        return ans;
    }
};
```

```go [sol-Go]
func largestPalindrome(n, k int) string {
	pow10 := make([]int, n)
	pow10[0] = 1
	for i := 1; i < n; i++ {
		pow10[i] = pow10[i-1] * 10 % k
	}

	ans := make([]byte, n)
	m := (n + 1) / 2
	vis := make([][]bool, m+1)
	for i := range vis {
		vis[i] = make([]bool, k)
	}
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i == m {
			return j == 0
		}
		vis[i][j] = true
		for d := 9; d >= 0; d-- { // 贪心：从大到小枚举
			var j2 int
			if n%2 > 0 && i == m-1 { // 正中间
				j2 = (j + d*pow10[i]) % k
			} else {
				j2 = (j + d*(pow10[i]+pow10[n-1-i])) % k
			}
			if !vis[i+1][j2] && dfs(i+1, j2) {
				ans[i] = '0' + byte(d)
				ans[n-1-i] = ans[i]
				return true
			}
		}
		return false
	}
	dfs(0, 0)
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nkD)$，其中 $D=10$。
- 空间复杂度：$\mathcal{O}(nk)$。

## 相似题目

见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) 中的「**专题：输出具体方案（打印方案）**」。

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
