## 写法一

用一个变量 $\textit{ok}$ 表示当前数字能否加入答案。初始值为 $\texttt{true}$，每遍历一个数就取反，这样我们可以选一个数，跳过一个数，选一个数，跳过一个数，……

对于下标为奇数的行，倒序遍历（或者将其反转）。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1HKcue9ETm/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def zigzagTraversal(self, grid: List[List[int]]) -> List[int]:
        ans = []
        ok = True
        for i, row in enumerate(grid):
            if i % 2:
                row.reverse()
            for x in row:
                if ok:
                    ans.append(x)
                ok = not ok
        return ans
```

```py [sol-Python3 一行]
class Solution:
    def zigzagTraversal(self, grid: List[List[int]]) -> List[int]:
        return list(chain(*(row[::-1] if i % 2 else row for i, row in enumerate(grid))))[::2]
```

```java [sol-Java]
class Solution {
    public List<Integer> zigzagTraversal(int[][] grid) {
        List<Integer> ans = new ArrayList<>();
        boolean ok = true;
        for (int i = 0; i < grid.length; i++) {
            if (i % 2 == 0) {
                for (int x : grid[i]) {
                    if (ok) {
                        ans.add(x);
                    }
                    ok = !ok;
                }
            } else {
                for (int j = grid[i].length - 1; j >= 0; j--) {
                    if (ok) {
                        ans.add(grid[i][j]);
                    }
                    ok = !ok;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> zigzagTraversal(vector<vector<int>>& grid) {
        vector<int> ans;
        bool ok = true;
        for (int i = 0; i < grid.size(); i++) {
            auto& row = grid[i];
            if (i % 2) {
                ranges::reverse(row);
            }
            for (int x : row) {
                if (ok) {
                    ans.push_back(x);
                }
                ok = !ok;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func zigzagTraversal(grid [][]int) (ans []int) {
	ok := true
	for i, row := range grid {
		if i%2 > 0 {
			slices.Reverse(row)
		}
		for _, x := range row {
			if ok {
				ans = append(ans, x)
			}
			ok = !ok
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

## 写法二

前 $2k$ 行一共有 $2kn$ 个数，这必然是偶数。所以 $0,2,4,\cdots$ 行必然是从第一个数开始选的。

对于 $1,3,5,\cdots$ 行：

- 如果 $n$ 是偶数，从下标 $n-1$ 开始选。
- 如果 $n$ 是奇数，从下标 $n-2$ 开始选。

综合一下，从下标 $n-1-n\bmod 2$ 开始选。

```py [sol-Python3]
class Solution:
    def zigzagTraversal(self, grid: List[List[int]]) -> List[int]:
        end = -1 - len(grid[0]) % 2
        ans = []
        for i, row in enumerate(grid):
            ans.extend(row[end::-2] if i % 2 else row[::2])
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def zigzagTraversal(self, grid: List[List[int]]) -> List[int]:
        end = -1 - len(grid[0]) % 2
        return list(chain(*(row[end::-2] if i % 2 else row[::2] for i, row in enumerate(grid))))
```

```java [sol-Java]
class Solution {
    public List<Integer> zigzagTraversal(int[][] grid) {
        int n = grid[0].length;
        int end = n - 1 - n % 2;
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < grid.length; i++) {
            if (i % 2 > 0) {
                for (int j = end; j >= 0; j -= 2) {
                    ans.add(grid[i][j]);
                }
            } else {
                for (int j = 0; j < n; j += 2) {
                    ans.add(grid[i][j]);
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> zigzagTraversal(vector<vector<int>>& grid) {
        int n = grid[0].size();
        int end = n - 1 - n % 2;
        vector<int> ans;
        for (int i = 0; i < grid.size(); i++) {
            if (i % 2) {
                for (int j = end; j >= 0; j -= 2) {
                    ans.push_back(grid[i][j]);
                }
            } else {
                for (int j = 0; j < n; j += 2) {
                    ans.push_back(grid[i][j]);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func zigzagTraversal(grid [][]int) (ans []int) {
	n := len(grid[0])
	end := n - 1 - n%2
	for i, row := range grid {
		if i%2 > 0 {
			for j := end; j >= 0; j -= 2 {
				ans = append(ans, row[j])
			}
		} else {
			for j := 0; j < n; j += 2 {
				ans = append(ans, row[j])
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。Python 忽略切片的空间。

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
