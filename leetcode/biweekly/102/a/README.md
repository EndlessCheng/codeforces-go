## 方法一

遍历每一列，求出数字转成字符串后的最大长度。

```py [sol-Python3]
class Solution:
    def findColumnWidth(self, grid: List[List[int]]) -> List[int]:
        return [max(len(str(x)) for x in col)
                for col in zip(*grid)]
```

```java [sol-Java]
class Solution {
    public int[] findColumnWidth(int[][] grid) {
        int n = grid[0].length;
        int[] ans = new int[n];
        for (int j = 0; j < n; j++) {
            for (int[] row : grid) {
                ans[j] = Math.max(ans[j], Integer.toString(row[j]).length());
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findColumnWidth(vector<vector<int>>& grid) {
        int n = grid[0].size();
        vector<int> ans(n);
        for (int j = 0; j < n; j++) {
            for (auto& row : grid) {
                ans[j] = max(ans[j], (int) to_string(row[j]).length());
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findColumnWidth(grid [][]int) []int {
    ans := make([]int, len(grid[0]))
    for j := range grid[0] {
        for _, row := range grid {
            ans[j] = max(ans[j], len(strconv.Itoa(row[j])))
        }
    }
    return ans
}
```

```js [sol-JavaScript]
var findColumnWidth = function(grid) {
    const n = grid[0].length;
    const ans = Array(n).fill(0);
    for (let j = 0; j < n; j++) {
        for (const row of grid) {
            ans[j] = Math.max(ans[j], row[j].toString().length);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_column_width(grid: Vec<Vec<i32>>) -> Vec<i32> {
        (0..grid[0].len()).map(|j| {
            grid.iter().map(|row| row[j].to_string().len()).max().unwrap() as i32
        }).collect()
    }
}
```

也可以手动计算长度。

```py [sol-Python3]
class Solution:
    def findColumnWidth(self, grid: List[List[int]]) -> List[int]:
        ans = [0] * len(grid[0])
        for j, col in enumerate(zip(*grid)):
            for x in col:
                x_len = int(x <= 0)
                x = abs(x)
                while x:
                    x_len += 1
                    x //= 10
                ans[j] = max(ans[j], x_len)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] findColumnWidth(int[][] grid) {
        int n = grid[0].length;
        int[] ans = new int[n];
        for (int j = 0; j < n; j++) {
            for (int[] row : grid) {
                int len = row[j] <= 0 ? 1 : 0;
                for (int x = row[j]; x != 0; x /= 10) {
                    len++;
                }
                ans[j] = Math.max(ans[j], len);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findColumnWidth(vector<vector<int>>& grid) {
        int n = grid[0].size();
        vector<int> ans(n);
        for (int j = 0; j < n; j++) {
            for (auto& row : grid) {
                int len = row[j] <= 0;
                for (int x = row[j]; x; x /= 10) {
                    len++;
                }
                ans[j] = max(ans[j], len);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findColumnWidth(grid [][]int) []int {
    ans := make([]int, len(grid[0]))
    for j := range grid[0] {
        for _, row := range grid {
            xLen := 0
            if row[j] <= 0 {
                xLen = 1
            }
            for x := row[j]; x != 0; x /= 10 {
                xLen++
            }
            ans[j] = max(ans[j], xLen)
        }
    }
    return ans
}
```

```js [sol-JavaScript]
var findColumnWidth = function(grid) {
    const n = grid[0].length;
    const ans = Array(n).fill(0);
    for (let j = 0; j < n; j++) {
        for (const row of grid) {
            let len = row[j] <= 0 ? 1 : 0;
            for (let x = Math.abs(row[j]); x; x = Math.floor(x / 10)) {
                len++;
            }
            ans[j] = Math.max(ans[j], len);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_column_width(grid: Vec<Vec<i32>>) -> Vec<i32> {
        let n = grid[0].len();
        let mut ans = vec![0; n];
        for j in 0..n {
            for row in &grid {
                let mut len = if row[j] <= 0 { 1 } else { 0 };
                let mut x = row[j];
                while x != 0 {
                    len += 1;
                    x /= 10;
                }
                ans[j] = ans[j].max(len);
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log U)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数，$U$ 为 $\textit{grid}[i][j]$ 的绝对值的最大值。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。Python 忽略 `zip(*grid)` 的空间。

## 方法二：优化

方法一需要对每个数字都计算长度。但实际上，由于数字的绝对值越大，数字的长度就越长，所以只需要对每一列的最小值或最大值求长度。

设列最小值和列最大值分别为 $\textit{mn}$ 和 $\textit{mx}$。

由于负数中的负号也算一个长度，我们可以取

$$
\max(\textit{mx}, -10\cdot \textit{mn})
$$

的长度作为答案。

或者，为避免乘法溢出，取

$$
\max\left(\left\lfloor\dfrac{\textit{mx}}{10}\right\rfloor, -\textit{mn}\right)
$$

的长度**加一**作为答案。此时要把 $0$ 的长度视作 $0$。

注意上式在一整列全为负数或者全为正数时也是正确的。

```py [sol-Python3]
class Solution:
    def findColumnWidth(self, grid: List[List[int]]) -> List[int]:
        return [len(str(max(max(col), -10 * min(col))))
                for col in zip(*grid)]
```

```py [sol-Python3 写法二]
class Solution:
    def findColumnWidth(self, grid: List[List[int]]) -> List[int]:
        ans = []
        for col in zip(*grid):
            x_len = 1
            x = max(max(col) // 10, -min(col))
            while x:
                x_len += 1
                x //= 10
            ans.append(x_len)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] findColumnWidth(int[][] grid) {
        int n = grid[0].length;
        int[] ans = new int[n];
        for (int j = 0; j < n; j++) {
            int mn = 0;
            int mx = 0;
            for (int[] row : grid) {
                mn = Math.min(mn, row[j]);
                mx = Math.max(mx, row[j]);
            }
            int len = 1;
            for (int x = Math.max(mx / 10, -mn); x > 0; x /= 10) {
                len++;
            }
            ans[j] = len;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findColumnWidth(vector<vector<int>>& grid) {
        int n = grid[0].size();
        vector<int> ans(n);
        for (int j = 0; j < n; j++) {
            int mn = 0, mx = 0;
            for (auto& row : grid) {
                mn = min(mn, row[j]);
                mx = max(mx, row[j]);
            }
            int len = 1;
            for (int x = max(mx / 10, -mn); x; x /= 10) {
                len++;
            }
            ans[j] = len;
        }
        return ans;
    }
};
```

```go [sol-Go]
func findColumnWidth(grid [][]int) []int {
    ans := make([]int, len(grid[0]))
    for j := range grid[0] {
        mn, mx := 0, 0
        for _, row := range grid {
            mn = min(mn, row[j])
            mx = max(mx, row[j])
        }
        xLen := 1
        for x := max(mx/10, -mn); x > 0; x /= 10 {
            xLen++
        }
        ans[j] = xLen
    }
    return ans
}
```

```js [sol-JavaScript]
var findColumnWidth = function(grid) {
    const n = grid[0].length;
    const ans = Array(n);
    for (let j = 0; j < n; j++) {
        let mn = 0, mx = 0;
        for (const row of grid) {
            mn = Math.min(mn, row[j]);
            mx = Math.max(mx, row[j]);
        }
        let len = 1;
        for (let x = Math.max(Math.floor(mx / 10), -mn); x; x = Math.floor(x / 10)) {
            len++;
        }
        ans[j] = len;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_column_width(grid: Vec<Vec<i32>>) -> Vec<i32> {
        let n = grid[0].len();
        let mut ans = vec![0; n];
        for j in 0..n {
            let mut mn = 0;
            let mut mx = 0;
            for row in &grid {
                mn = mn.min(row[j]);
                mx = mx.max(row[j]);
            }
            let mut len = 1;
            let mut x = (mx / 10).max(-mn);
            while x > 0 {
                len += 1;
                x /= 10;
            }
            ans[j] = len;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(m+\log U))$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数，$U$ 为 $\textit{grid}[i][j]$ 的绝对值的最大值。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。Python 忽略 `zip(*grid)` 的空间。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
