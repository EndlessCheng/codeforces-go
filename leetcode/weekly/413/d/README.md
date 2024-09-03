## 数组的异或值

下文用 $\oplus$ 表示异或。

考虑数组的异或值（最后剩下的元素）是由哪些元素异或得到的。

例如数组为 $[a,b,c]$，那么操作一次后变成 $[a\oplus b,\ b\oplus c]$，再操作一次，得到 $a\oplus b\oplus b\oplus c$。其中 $b$ 异或了 $2$ 次。

为方便描述，下面把 $a\oplus b$ 简记为 $ab$，把 $a\oplus b\oplus b\oplus c$ 简记为 $ab^2c$。

又例如数组为 $[a,b,c,d]$，那么操作一次后变成 $[ab,bc,cd]$，再操作一次，变成 $[ab^2c,bc^2d]$，再操作一次，得到 $ab^3c^3d$。

可以发现，$ab^3c^3d$ 相当于数组 $[a,b,c]$ 的异或值，再异或 $[b,c,d]$ 的异或值。

> 当然，你可以用组合数计算出幂次，那就是另一个思路了，具体见 [本题视频讲解](https://www.bilibili.com/video/BV142Hae7E5y/)。

## 第一个区间 DP

定义 $f[i][j]$ 表示下标从 $i$ 到 $j$ 的子数组的「数组的异或值」，根据上面的讨论，有

$$
f[i][j] = f[i][j-1]\oplus f[i+1][j]
$$

初始值：$f[i][i]=\textit{nums}[i]$。

## 第二个区间 DP

为了回答询问，我们需要计算下标从 $i$ 到 $j$ 的子数组中的所有子数组的 $f$ 值的最大值，将其记作 $\textit{mx}[i][j]$。

分类讨论：

- 取 $f[i][j]$ 作为最大值。
- 最大值对应的子数组，右端点不是 $j$，那么问题变成下标从 $i$ 到 $j-1$ 的子数组中的所有子数组的 $f$ 值的最大值，即 $\textit{mx}[i][j-1]$。
- 最大值对应的子数组，左端点不是 $i$，那么问题变成下标从 $i+1$ 到 $j$ 的子数组中的所有子数组的 $f$ 值的最大值，即 $\textit{mx}[i+1][j]$。

三者取最大值，得

$$
\textit{mx}[i][j] = \max(f[i][j], \textit{mx}[i][j-1], \textit{mx}[i+1][j])
$$

初始值：$\textit{mx}[i][i]=\textit{nums}[i]$。

回答询问时直接查询 $\textit{mx}$ 数组。

下面代码为什么 $i$ 要倒序循环，请看 [区间 DP【基础算法精讲 22】](https://www.bilibili.com/video/BV1Gs4y1E7EU/)。

[本题视频讲解](https://www.bilibili.com/video/BV142Hae7E5y/)（第四题），欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumSubarrayXor(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        n = len(nums)
        f = [[0] * n for _ in range(n)]
        mx = [[0] * n for _ in range(n)]
        for i in range(n - 1, -1, -1):
            mx[i][i] = f[i][i] = nums[i]
            for j in range(i + 1, n):
                f[i][j] = f[i][j - 1] ^ f[i + 1][j]
                mx[i][j] = max(f[i][j], mx[i + 1][j], mx[i][j - 1])
        return [mx[l][r] for l, r in queries]
```

```java [sol-Java]
class Solution {
    public int[] maximumSubarrayXor(int[] nums, int[][] queries) {
        int n = nums.length;
        int[][] f = new int[n][n];
        int[][] mx = new int[n][n];
        for (int i = n - 1; i >= 0; i--) {
            f[i][i] = nums[i];
            mx[i][i] = nums[i];
            for (int j = i + 1; j < n; j++) {
                f[i][j] = f[i][j - 1] ^ f[i + 1][j];
                mx[i][j] = Math.max(f[i][j], Math.max(mx[i + 1][j], mx[i][j - 1]));
            }
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            ans[i] = mx[queries[i][0]][queries[i][1]];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maximumSubarrayXor(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        vector<vector<int>> f(n, vector<int>(n));
        vector<vector<int>> mx(n, vector<int>(n));
        for (int i = n - 1; i >= 0; i--) {
            f[i][i] = nums[i];
            mx[i][i] = nums[i];
            for (int j = i + 1; j < n; j++) {
                f[i][j] = f[i][j - 1] ^ f[i + 1][j];
                mx[i][j] = max({f[i][j], mx[i + 1][j], mx[i][j - 1]});
            }
        }

        vector<int> ans;
        for (auto& q : queries) {
            ans.push_back(mx[q[0]][q[1]]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumSubarrayXor(nums []int, queries [][]int) []int {
	n := len(nums)
	f := make([][]int, n)
	mx := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		mx[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][i] = nums[i]
		mx[i][i] = nums[i]
		for j := i + 1; j < n; j++ {
			f[i][j] = f[i][j-1] ^ f[i+1][j]
			mx[i][j] = max(f[i][j], mx[i+1][j], mx[i][j-1])
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = mx[q[0]][q[1]]
	}
	return ans
}
```

## 优化

去掉 $f$ 的第一个维度。

进一步地，直接把 $\textit{nums}$ 当作 $f$ 数组。

```py [sol-Python3]
class Solution:
    def maximumSubarrayXor(self, f: List[int], queries: List[List[int]]) -> List[int]:
        n = len(f)
        mx = [[0] * n for _ in range(n)]
        for i in range(n - 1, -1, -1):
            mx[i][i] = f[i]
            for j in range(i + 1, n):
                f[j] ^= f[j - 1]
                mx[i][j] = max(f[j], mx[i + 1][j], mx[i][j - 1])
        return [mx[l][r] for l, r in queries]
```

```py [sol-Python3 写法二]
class Solution:
    def maximumSubarrayXor(self, f: List[int], queries: List[List[int]]) -> List[int]:
        n = len(f)
        mx = [[0] * n for _ in range(n)]
        for i in range(n - 1, -1, -1):
            mx[i][i] = f[i]
            for j in range(i + 1, n):
                f[j] ^= f[j - 1]
                res = f[j]
                if mx[i + 1][j] > res:
                    res = mx[i + 1][j]
                if mx[i][j - 1] > res:
                    res = mx[i][j - 1]
                mx[i][j] = res
        return [mx[l][r] for l, r in queries]
```

```java [sol-Java]
class Solution {
    public int[] maximumSubarrayXor(int[] f, int[][] queries) {
        int n = f.length;
        int[][] mx = new int[n][n];
        for (int i = n - 1; i >= 0; i--) {
            mx[i][i] = f[i];
            for (int j = i + 1; j < n; j++) {
                f[j] ^= f[j - 1];
                mx[i][j] = Math.max(f[j], Math.max(mx[i + 1][j], mx[i][j - 1]));
            }
        }

        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            ans[i] = mx[queries[i][0]][queries[i][1]];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maximumSubarrayXor(vector<int>& f, vector<vector<int>>& queries) {
        int n = f.size();
        vector<vector<int>> mx(n, vector<int>(n));
        for (int i = n - 1; i >= 0; i--) {
            mx[i][i] = f[i];
            for (int j = i + 1; j < n; j++) {
                f[j] ^= f[j - 1];
                mx[i][j] = max({f[j], mx[i + 1][j], mx[i][j - 1]});
            }
        }

        vector<int> ans;
        for (auto& q : queries) {
            ans.push_back(mx[q[0]][q[1]]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumSubarrayXor(f []int, queries [][]int) []int {
	n := len(f)
	mx := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		mx[i] = make([]int, n)
		mx[i][i] = f[i]
		for j := i + 1; j < n; j++ {
			f[j] ^= f[j-1]
			mx[i][j] = max(f[j], mx[i+1][j], mx[i][j-1])
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = mx[q[0]][q[1]]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度
- 空间复杂度：$\mathcal{O}(n^2)$。返回值不计入。

## 思考题

如果把 $\textit{nums}$ 的长度增大到 $10^6$，且只要求计算 $\textit{nums}$ 的「异或值」，你能想出一个更快的做法吗？

见 [视频讲解](https://www.bilibili.com/video/BV142Hae7E5y/) 的最后。证明需要用到 Lucas 定理。

更多相似题目，见下面动态规划题单中的「**区间 DP**」。

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
