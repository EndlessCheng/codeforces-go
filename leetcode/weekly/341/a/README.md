遍历每一行，计算行的元素和 $s$。如果 $s$ 比最大和 $\textit{maxSum}$ 大，更新 $\textit{maxSum}=s$，顺带记录行号 $\textit{rowIdx}$。

由于我们是从上到下遍历行，所以找到的最大和的行号，一定是相同最大和的行号中最小的。

```py [sol-Python3]
class Solution:
    def rowAndMaximumOnes(self, mat: List[List[int]]) -> List[int]:
        row_idx = max_sum = -1
        for i, row in enumerate(mat):
            s = sum(row)
            if s > max_sum:
                row_idx, max_sum = i, s
        return [row_idx, max_sum]
```

```java [sol-Java]
class Solution {
    public int[] rowAndMaximumOnes(int[][] mat) {
        int rowIdx = -1;
        int maxSum = -1;
        for (int i = 0; i < mat.length; i++) {
            int s = Arrays.stream(mat[i]).sum();
            if (s > maxSum) {
                rowIdx = i;
                maxSum = s;
            }
        }
        return new int[]{rowIdx, maxSum};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> rowAndMaximumOnes(vector<vector<int>>& mat) {
        int row_idx = -1, max_sum = -1;
        for (int i = 0; i < mat.size(); i++) {
            int s = reduce(mat[i].begin(), mat[i].end());
            if (s > max_sum) {
                row_idx = i;
                max_sum = s;
            }
        }
        return {row_idx, max_sum};
    }
};
```

```c [sol-C]
int* rowAndMaximumOnes(int** mat, int matSize, int* matColSize, int* returnSize) {
    int row_idx = -1, max_sum = -1;
    for (int i = 0; i < matSize; i++) {
        int s = 0;
        for (int j = 0; j < matColSize[i]; j++) {
            s += mat[i][j];
        }
        if (s > max_sum) {
            row_idx = i;
            max_sum = s;
        }
    }

    *returnSize = 2;
    int* ans = malloc(2 * sizeof(int));
    ans[0] = row_idx;
    ans[1] = max_sum;
    return ans;
}
```

```go [sol-Go]
func rowAndMaximumOnes(mat [][]int) []int {
    rowIdx, maxSum := -1, -1
    for i, row := range mat {
        s := 0
        for _, x := range row {
            s += x
        }
        if s > maxSum {
            rowIdx, maxSum = i, s
        }
    }
    return []int{rowIdx, maxSum}
}
```

```js [sol-JavaScript]
var rowAndMaximumOnes = function(mat) {
    let rowIdx = -1, maxSum = -1;
    for (let i = 0; i < mat.length; i++) {
        const s = _.sum(mat[i]);
        if (s > maxSum) {
            rowIdx = i;
            maxSum = s;
        }
    }
    return [rowIdx, maxSum];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn row_and_maximum_ones(mat: Vec<Vec<i32>>) -> Vec<i32> {
        let mut row_idx = 0;
        let mut max_sum = -1;
        for (i, row) in mat.into_iter().enumerate() {
            let s = row.into_iter().sum();
            if s > max_sum {
                row_idx = i;
                max_sum = s;
            }
        }
        vec![row_idx as _, max_sum]
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{mat}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
