根据题意，点 $i$ 的度等于 $\textit{matrix}[i]$ 中的 $1$ 的个数。

所以统计矩阵每一行的 $1$ 的个数即可。

```py [sol-Python3]
class Solution:
    def findDegrees(self, matrix: List[List[int]]) -> List[int]:
        return list(map(sum, matrix))
```

```java [sol-Java]
class Solution {
    public int[] findDegrees(int[][] matrix) {
        int[] ans = new int[matrix.length];
        for (int i = 0; i < matrix.length; i++) {
            for (int x : matrix[i]) {
                ans[i] += x;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findDegrees(vector<vector<int>>& matrix) {
        vector<int> ans;
        for (auto& row : matrix) {
            ans.push_back(reduce(row.begin(), row.end(), 0));
        }
        return ans;
    }
};
```

```go [sol-Go]
func findDegrees(matrix [][]int) []int {
	ans := make([]int, len(matrix))
	for i, row := range matrix {
		for _, x := range row {
			ans[i] += x
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别是 $\textit{matrix}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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
