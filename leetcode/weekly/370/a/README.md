## 分析

分析题意：

- 如果 $\textit{grid}[i][j]=1$，说明 $i$ 队比 $j$ 队**强**。
- 如果 $\textit{grid}[i][j]=0$，说明 $i$ 队比 $j$ 队**弱**。
- 没有平手。

所以 a 队要是冠军，a 队就要比其它 $n-1$ 个队都要强。

## 方法一：横看成岭

如果 $\textit{grid}[i]$ 有 $n-1$ 个 $1$，即元素和为 $n-1$，说明 $i$ 队比其它 $n-1$ 个队都要强，$i$ 队是冠军。

也可以判断，对于这一行的所有不等于 $i$ 的 $j$，都有 $\textit{grid}[i][j]=1$。这样可以在遇到 $0$ 的时候，提前退出循环。

```py [sol-Python3]
class Solution:
    def findChampion(self, grid: List[List[int]]) -> int:
        for i, row in enumerate(grid):
            if sum(row) == len(grid) - 1:
                return i
```

```java [sol-Java]
class Solution {
    public int findChampion(int[][] grid) {
        next:
        for (int i = 0; ; i++) {
            for (int j = 0; j < grid.length; j++) {
                if (j != i && grid[i][j] == 0) {
                    continue next;
                }
            }
            return i;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findChampion(vector<vector<int>> &grid) {
        int n = grid.size();
        for (int i = 0; ; i++) {
            if (accumulate(grid[i].begin(), grid[i].end(), 0) == n - 1) {
                return i;
            }
        }
    }
};
```

```cpp [sol-C++ 写法二]
class Solution {
public:
    int findChampion(vector<vector<int>> &grid) {
        int n = grid.size();
        for (int i = 0; ; i++) {
            bool ok = true;
            for (int j = 0; j < n && ok; j++) {
                ok = j == i || grid[i][j];
            }
            if (ok) {
                return i;
            }
        }
    }
};
```

```c [sol-C]
int findChampion(int** grid, int gridSize, int* gridColSize) {
    int n = gridSize;
    for (int i = 0; ; i++) {
        bool ok = true;
        for (int j = 0; j < n && ok; j++) {
            ok = j == i || grid[i][j] == 1;
        }
        if (ok) {
            return i;
        }
    }
}
```

```go [sol-Go]
func findChampion(grid [][]int) int {
next:
    for i, row := range grid {
        for j, x := range row {
            if j != i && x == 0 {
                continue next
            }
        }
        return i
    }
    panic(-1)
}
```

```js [sol-JS]
var findChampion = function(grid) {
    for (let i = 0; ; i++) {
        if (_.sum(grid[i]) === grid.length - 1) {
            return i;
        }
    }
};
```

```js [sol-JS 写法二]
var findChampion = function(grid) {
    const n = grid.length;
    for (let i = 0; ; i++) {
        let ok = true;
        for (let j = 0; j < n && ok; j++) {
            ok = j === i || grid[i][j] === 1;
        }
        if (ok) {
            return i;
        }
    }
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_champion(grid: Vec<Vec<i32>>) -> i32 {
        let n = grid.len() as i32;
        for (i, row) in grid.iter().enumerate() {
            if row.iter().sum::<i32>() == n - 1 {
                return i as _;
            }
        }
        unreachable!()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{grid}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：侧成峰

如果第 $j$ 列的元素值都是 $0$，说明没有队伍可以击败 $j$ 队，$j$ 队是冠军。

```py [sol-Python3]
class Solution:
    def findChampion(self, grid: List[List[int]]) -> int:
        for j, col in enumerate(zip(*grid)):
            if 1 not in col:
                return j
```

```java [sol-Java]
class Solution {
    public int findChampion(int[][] grid) {
        next:
        for (int j = 0; ; j++) {
            for (int[] row : grid) {
                if (row[j] != 0) {
                    continue next;
                }
            }
            return j;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findChampion(vector<vector<int>> &grid) {
        int n = grid.size();
        for (int j = 0;; j++) {
            bool ok = true;
            for (int i = 0; i < n && ok; i++) {
                ok = grid[i][j] == 0;
            }
            if (ok) {
                return j;
            }
        }
    }
};
```

```c [sol-C]
int findChampion(int** grid, int gridSize, int* gridColSize) {
    int n = gridSize;
    for (int j = 0; ; j++) {
        bool ok = true;
        for (int i = 0; i < n && ok; i++) {
            ok = grid[i][j] == 0;
        }
        if (ok) {
            return j;
        }
    }
}
```

```go [sol-Go]
func findChampion(grid [][]int) int {
next:
    for j := range grid[0] {
        for _, row := range grid {
            if row[j] != 0 {
                continue next
            }
        }
        return j
    }
    panic(-1)
}
```

```js [sol-JS]
var findChampion = function(grid) {
    const n = grid.length;
    for (let j = 0; ; j++) {
        let ok = true;
        for (let i = 0; i < n && ok; i++) {
            ok = grid[i][j] === 0;
        }
        if (ok) {
            return j;
        }
    }
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_champion(grid: Vec<Vec<i32>>) -> i32 {
        for j in 0.. {
            let mut ok = true;
            for row in &grid {
                if row[j] != 0 {
                    ok = false;
                    break;
                }
            }
            if ok {
                return j as _;
            }
        }
        unreachable!()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{grid}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法三：打擂台

假设冠军是 $\textit{ans}=0$，我们从 $i=1$ 开始遍历，寻找可以击败 $\textit{ans}$ 的队伍，也就是 $\textit{grid}[i][\textit{ans}]=1$。

如果没有出现 $\textit{grid}[i][\textit{ans}]=1$，那么答案就是 $\textit{ans}$，否则冠军可能是 $i$，更新 $\textit{ans}=i$。然后从 $i+1$ 继续向后遍历，因为 $[1,i-1]$ 中没有比 $0$ 强的队，更别说比 $i$ 强了。重复上述过程，最后返回 $\textit{ans}$。

```py [sol-Python3]
class Solution:
    def findChampion(self, grid: List[List[int]]) -> int:
        ans = 0
        for i, row in enumerate(grid):
            if row[ans]:
                ans = i
        return ans
```

```java [sol-Java]
class Solution {
    public int findChampion(int[][] grid) {
        int ans = 0;
        for (int i = 1; i < grid.length; i++) {
            if (grid[i][ans] == 1) {
                ans = i;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findChampion(vector<vector<int>> &grid) {
        int ans = 0;
        for (int i = 1; i < grid.size(); i++) {
            if (grid[i][ans]) {
                ans = i;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int findChampion(int** grid, int gridSize, int* gridColSize) {
    int ans = 0;
    for (int i = 1; i < gridSize; i++) {
        if (grid[i][ans]) {
            ans = i;
        }
    }
    return ans;
}
```

```go [sol-Go]
func findChampion(grid [][]int) (ans int) {
    for i, row := range grid {
        if row[ans] == 1 {
            ans = i
        }
    }
    return
}
```

```js [sol-JS]
var findChampion = function(grid) {
    let ans = 0;
    for (let i = 1; i < grid.length; i++) {
        if (grid[i][ans]) {
            ans = i;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_champion(grid: Vec<Vec<i32>>) -> i32 {
        let mut ans = 0;
        for (i, row) in grid.iter().enumerate() {
            if row[ans] == 1 {
                ans = i;
            }
        }
        ans as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{grid}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
