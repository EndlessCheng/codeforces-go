### 本题视频讲解

见[【周赛 340】](https://www.bilibili.com/video/BV1iN411w7my/)。

### 前置知识：动态规划

见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

### 前置知识：单调栈

见我之前写的 [这篇题解](https://leetcode.cn/problems/sum-of-subarray-minimums/solution/gong-xian-fa-dan-diao-zhan-san-chong-shi-gxa5/) 的提示 3。

### 前置知识：二分查找

见 [二分查找【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 思路

暴力做法是从 $(0,0)$ 出发，向右/向下尝试移动到每个满足要求的格子。假设移动到了 $(i,j)$，那么问题就变成从 $(i,j)$ 出发，到右下角的最少移动格子数。这是一个和原问题相似的子问题，启发我们用递归来思考。

看了上面的视频后，你知道递归中有重叠子问题，从而可以用记忆化搜索来优化，而记忆化搜索又可以翻译成如下的递推。

定义 $f[i][j]$ 表示从 $(i,j)$ 出发，到右下角最少移动格子数。

设 $g=\textit{grid}[i][j]$，有

$$
f[i][j] = \min\left\{\min_{k=j+1}^{j+g} f[i][k], \min_{k=i+1}^{i+g} f[k][j]\right\} + 1
$$

$i$ 和 $j$ 均倒序遍历。答案为 $f[0][0]$。

但这样做时间复杂度是 $O(mn(m+n))$ 的，无法接受。

由于有「区间查询」、「单点更新」这两个操作，我们可以用线段树来优化。

但还有更「轻量级」的做法。

对于 $\min\limits_{k=j+1}^{j+g} f[i][k]$ 来说，在倒序遍历 $j$ 时，$k$ 的左边界 $j+1$ 是在**单调减小**的，我们可以用一个 $f$ 值底小顶大的单调栈来维护 $f[i][k]$ 及其下标 $k$。由于是倒序遍历，单调栈中的下标是底大顶小的，从那么在单调栈上二分查找最大的不超过 $j+g$ 的下标 $k$，对应的 $f[i][k]$ 就是 $[j+1, j+g]$ 范围内的最小值。

对于 $\min\limits_{k=i+1}^{i+g} f[k][j]$ 也同理，每一列都需要一个单调栈。

```py [sol1-Python3]
class Solution:
    def minimumVisitedCells(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        col_st = [[] for _ in range(n)]  # 每列的单调栈
        for i in range(m - 1, -1, -1):
            st = []  # 当前行的单调栈
            for j in range(n - 1, -1, -1):
                st2 = col_st[j]
                mn = inf
                g = grid[i][j]
                if i == m - 1 and j == n - 1:  # 特殊情况：已经是终点
                    mn = 0
                elif g:
                    # 在单调栈上二分
                    k = bisect_left(st, -(j + g), key=lambda p: p[1])
                    if k < len(st): mn = min(mn, st[k][0])
                    k = bisect_left(st2, -(i + g), key=lambda p: p[1])
                    if k < len(st2): mn = min(mn, st2[k][0])
                if mn == inf: continue

                mn += 1  # 加上 (i,j) 这个格子
                # 插入单调栈
                while st and mn <= st[-1][0]:
                    st.pop()
                st.append((mn, -j))  # 保证下标单调递增，方便调用 bisect_left
                while st2 and mn <= st2[-1][0]:
                    st2.pop()
                st2.append((mn, -i))  # 保证下标单调递增，方便调用 bisect_left
        return mn if mn < inf else -1  # 最后一个算出的 mn 就是 f[0][0]
```

```java [sol1-Java]
class Solution {
    public int minimumVisitedCells(int[][] grid) {
        int m = grid.length, n = grid[0].length, mn = 0;
        List<int[]>[] colSt = new ArrayList[n]; // 每列的单调栈
        Arrays.setAll(colSt, e -> new ArrayList<int[]>());
        for (int i = m - 1; i >= 0; --i) {
            var st = new ArrayList<int[]>(); // 当前行的单调栈
            for (int j = n - 1; j >= 0; --j) {
                var st2 = colSt[j];
                mn = Integer.MAX_VALUE;
                int g = grid[i][j];
                if (i == m - 1 && j == n - 1) // 特殊情况：已经是终点
                    mn = 0;
                else if (g > 0) {
                    // 在单调栈上二分
                    int k = search(st, j + g);
                    if (k < st.size()) mn = Math.min(mn, st.get(k)[0]);
                    k = search(st2, i + g);
                    if (k < st2.size()) mn = Math.min(mn, st2.get(k)[0]);
                }
                if (mn == Integer.MAX_VALUE) continue;

                ++mn; // 加上 (i,j) 这个格子
                // 插入单调栈
                while (!st.isEmpty() && mn <= st.get(st.size() - 1)[0])
                    st.remove(st.size() - 1);
                st.add(new int[]{mn, j});
                while (!st2.isEmpty() && mn <= st2.get(st2.size() - 1)[0])
                    st2.remove(st2.size() - 1);
                st2.add(new int[]{mn, i});
            }
        }
        return mn < Integer.MAX_VALUE ? mn : -1;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int search(List<int[]> st, int target) {
        int left = -1, right = st.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            int mid = (left + right) >>> 1;
            if (st.get(mid)[1] > target) left = mid; // 范围缩小到 (mid, right)
            else right = mid; // 范围缩小到 (left, mid)
        }
        return right;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minimumVisitedCells(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size(), mn;
        vector<vector<pair<int, int>>> col_st(n); // 每列的单调栈
        for (int i = m - 1; i >= 0; --i) {
            vector<pair<int, int>> st; // 当前行的单调栈
            for (int j = n - 1; j >= 0; --j) {
                auto &st2 = col_st[j];
                mn = INT_MAX;
                if (i == m - 1 && j == n - 1) // 特殊情况：已经是终点
                    mn = 0;
                else if (int g = grid[i][j]; g) {
                    // 在单调栈上二分
                    auto it = lower_bound(st.begin(), st.end(), j + g, [](const auto &a, const int b) {
                        return a.second > b;
                    });
                    if (it < st.end()) mn = min(mn, it->first);
                    it = lower_bound(st2.begin(), st2.end(), i + g, [](const auto &a, const int b) {
                        return a.second > b;
                    });
                    if (it < st2.end()) mn = min(mn, it->first);
                }
                if (mn == INT_MAX) continue;

                ++mn; // 加上 (i,j) 这个格子
                // 插入单调栈
                while (!st.empty() && mn <= st.back().first)
                    st.pop_back();
                st.emplace_back(mn, j);
                while (!st2.empty() && mn <= st2.back().first)
                    st2.pop_back();
                st2.emplace_back(mn, i);
            }
        }
        return mn < INT_MAX ? mn : -1;
    }
};
```

```go [sol1-Go]
func minimumVisitedCells(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	colSt := make([][]pair, n) // 每列的单调栈
	for i := m - 1; i >= 0; i-- {
		st := []pair{} // 当前行的单调栈
		for j := n - 1; j >= 0; j-- {
			st2 := colSt[j]
			mn = math.MaxInt
			if i == m-1 && j == n-1 { // 特殊情况：已经是终点
				mn = 0
			} else if g := grid[i][j]; g > 0 {
				// 在单调栈上二分
				k := j + g
				k = sort.Search(len(st), func(i int) bool { return st[i].i <= k })
				if k < len(st) {
					mn = min(mn, st[k].x)
				}
				k = i + g
				k = sort.Search(len(st2), func(i int) bool { return st2[i].i <= k })
				if k < len(st2) {
					mn = min(mn, st2[k].x)
				}
			}

			if mn < math.MaxInt {
				mn++ // 加上 (i,j) 这个格子
				// 插入单调栈
				for len(st) > 0 && mn <= st[len(st)-1].x {
					st = st[:len(st)-1]
				}
				st = append(st, pair{mn, j})
				for len(st2) > 0 && mn <= st2[len(st2)-1].x {
					st2 = st2[:len(st2)-1]
				}
				colSt[j] = append(st2, pair{mn, i})
			}
		}
	}
	// 最后一个算出的 mn 就是 f[0][0]
	if mn == math.MaxInt {
		return -1
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(mn\log(mn))$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(mn)$。
