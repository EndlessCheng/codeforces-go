## 方法一：间断点 + 二分查找

由于 $\textit{nums}$ 是有序的，如果 $\textit{nums}[i+1]-\textit{nums}[i] > \textit{maxDiff}$，那么编号 $\le i$ 的节点无法跳到编号 $\ge i+1$ 的节点。我把这样的 $i$ 叫做「间断点」。

遍历 $\textit{nums}$，把所有间断点记录在 $\textit{idx}$ 中。

对于询问 $[u,v]$，我们在 $\textit{idx}$ 中**二分查找** $\ge u$ 的第一个间断点，以及 $\ge v$ 的第一个间断点。如果这两个间断点相同，则可以从 $u$ 到达 $v$，否则不能。

关于二分查找的原理，见 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

[本题视频讲解](https://www.bilibili.com/video/BV1BgjAzcE7k/?t=3m12s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def pathExistenceQueries(self, n: int, nums: List[int], maxDiff: int, queries: List[List[int]]) -> List[bool]:
        idx = [i for i, (x, y) in enumerate(pairwise(nums)) if y - x > maxDiff]
        return [bisect_left(idx, u) == bisect_left(idx, v) for u, v in queries]
```

```java [sol-Java]
class Solution {
    public boolean[] pathExistenceQueries(int n, int[] nums, int maxDiff, int[][] queries) {
        List<Integer> idx = new ArrayList<>();
        for (int i = 0; i < n - 1; i++) {
            if (nums[i + 1] - nums[i] > maxDiff) {
                idx.add(i); // 间断点
            }
        }

        boolean[] ans = new boolean[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            ans[i] = binarySearch(idx, q[0]) == binarySearch(idx, q[1]);
        }
        return ans;
    }

    private int binarySearch(List<Integer> idx, int target) {
        // idx 中没有重复元素，可以用库函数二分
        int i = Collections.binarySearch(idx, target);
        return i < 0 ? ~i : i;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> pathExistenceQueries(int n, vector<int>& nums, int maxDiff, vector<vector<int>>& queries) {
        vector<int> idx;
        for (int i = 0; i < n - 1; i++) {
            if (nums[i + 1] - nums[i] > maxDiff) {
                idx.push_back(i); // 间断点
            }
        }

        vector<bool> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            auto& q = queries[i];
            ans[i] = ranges::lower_bound(idx, q[0]) == ranges::lower_bound(idx, q[1]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	idx := []int{}
	for i := range n - 1 {
		if nums[i+1]-nums[i] > maxDiff {
			idx = append(idx, i) // 间断点
		}
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = sort.SearchInts(idx, q[0]) == sort.SearchInts(idx, q[1])
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 方法二：记录连通块的编号

遍历 $\textit{nums}$，计算每个节点所在连通块的编号，从 $0$ 开始。

如果 $\textit{nums}[i]-\textit{nums}[i-1] > \textit{maxDiff}$，那么 $i$ 在一个新的连通块中，编号加一。否则编号不变。

设 $i$ 所在连通块的编号为 $\textit{id}[i]$，我们有

$$
\textit{id}[i] =
\begin{cases}
0, & i = 0 \\ 
\textit{id}[i-1], & \textit{nums}[i]-\textit{nums}[i-1] \le \textit{maxDiff}     \\
\textit{id}[i-1] + 1, & \textit{nums}[i]-\textit{nums}[i-1] > \textit{maxDiff}     \\
\end{cases}
$$

```py [sol-Python3]
class Solution:
    def pathExistenceQueries(self, n: int, nums: List[int], maxDiff: int, queries: List[List[int]]) -> List[bool]:
        id = [0] * n  # 每个节点所在连通块的编号
        for i in range(1, n):
            id[i] = id[i - 1]
            if nums[i] - nums[i - 1] > maxDiff:
                id[i] += 1  # 找到了一个新的连通块

        return [id[u] == id[v] for u, v in queries]
```

```py [sol-Python3 写法二]
class Solution:
    def pathExistenceQueries(self, n: int, nums: List[int], maxDiff: int, queries: List[List[int]]) -> List[bool]:
        id = list(accumulate((y - x > maxDiff for x, y in pairwise(nums)), initial=0))
        return [id[u] == id[v] for u, v in queries]
```

```java [sol-Java]
class Solution {
    public boolean[] pathExistenceQueries(int n, int[] nums, int maxDiff, int[][] queries) {
        int[] id = new int[n]; // 每个节点所在连通块的编号
        for (int i = 1; i < n; i++) {
            id[i] = id[i - 1];
            if (nums[i] - nums[i - 1] > maxDiff) {
                id[i]++; // 找到了一个新的连通块
            }
        }

        boolean[] ans = new boolean[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int[] q = queries[i];
            ans[i] = id[q[0]] == id[q[1]];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> pathExistenceQueries(int n, vector<int>& nums, int maxDiff, vector<vector<int>>& queries) {
        vector<int> id(n); // 每个节点所在连通块的编号
        for (int i = 1; i < n; i++) {
            id[i] = id[i - 1] + (nums[i] - nums[i - 1] > maxDiff); // 如果 >，那么找到了一个新的连通块
        }

        vector<bool> ans(queries.size());
        for (int i = 0; i < queries.size(); i++) {
            auto& q = queries[i];
            ans[i] = id[q[0]] == id[q[1]];
        }
        return ans;
    }
};
```

```go [sol-Go]
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	id := make([]int, n) // 每个节点所在连通块的编号
	for i := 1; i < n; i++ {
		id[i] = id[i-1]
		if nums[i]-nums[i-1] > maxDiff {
			id[i]++ // 找到了一个新的连通块
		}
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = id[q[0]] == id[q[1]]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

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
