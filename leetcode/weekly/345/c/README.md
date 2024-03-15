## 方法一：网格图 DFS

从第一列的任一单元格 $(i,0)$ 开始递归。枚举往右上/右/右下三个方向走，如果走一步后，没有出界，且格子值大于 $\textit{grid}[i][j]$，则可以走，继续递归。

在递归过程中，记录能访问到的最大列号，作为答案。

代码实现时，为避免重复递归之前访问过的格子，可以用一个 $\textit{vis}$ 数组标记访问过的格子。但实际上，可以把 $\textit{grid}[i][j]$ 置为 $0$ 从而无需创建 $\textit{vis}$ 数组。这是因为网格值均为正数，并且我们只能访问到比当前格子值更大的格子，所以置为 $0$ 会导致永远无法访问到该格子，这正是我们所希望的。

```py [sol-Python3]
class Solution:
    def maxMoves(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        ans = 0
        def dfs(i: int, j: int) -> None:
            nonlocal ans
            ans = max(ans, j)
            if ans == n - 1:  # ans 已达到最大值
                return
            for k in i - 1, i, i + 1:  # 向右上/右/右下走一步
                if 0 <= k < m and grid[k][j + 1] > grid[i][j]:
                    dfs(k, j + 1)
            grid[i][j] = 0
        for i in range(m):
            dfs(i, 0)  # 从第一列的任一单元格出发
        return ans
```

```java [sol-Java]
class Solution {
    private int ans;

    public int maxMoves(int[][] grid) {
        for (int i = 0; i < grid.length; i++) {
            dfs(i, 0, grid); // 从第一列的任一单元格出发
        }
        return ans;
    }

    private void dfs(int i, int j, int[][] grid) {
        ans = Math.max(ans, j);
        if (ans == grid[0].length - 1) { // ans 已达到最大值
            return;
        }
        // 向右上/右/右下走一步
        for (int k = Math.max(i - 1, 0); k < Math.min(i + 2, grid.length); k++) {
            if (grid[k][j + 1] > grid[i][j]) {
                dfs(k, j + 1, grid);
            }
        }
        grid[i][j] = 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxMoves(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size();
        int ans = 0;
        function<void(int, int)> dfs = [&](int i, int j) {
            ans = max(ans, j);
            if (ans == n - 1) { // ans 已达到最大值
                return;
            }
            // 向右上/右/右下走一步
            for (int k = max(i - 1, 0); k < min(i + 2, m); k++) {
                if (grid[k][j + 1] > grid[i][j]) {
                    dfs(k, j + 1);
                }
            }
            grid[i][j] = 0;
        };
        for (int i = 0; i < m; i++) {
            dfs(i, 0); // 从第一列的任一单元格出发
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		ans = max(ans, j)
		if ans == n-1 { // ans 已达到最大值
			return
		}
		// 向右上/右/右下走一步
		for k := max(i-1, 0); k < min(i+2, m); k++ {
			if grid[k][j+1] > grid[i][j] {
				dfs(k, j+1)
			}
		}
		grid[i][j] = 0
	}
	for i := range grid {
		dfs(i, 0) // 从第一列的任一单元格出发
	}
	return
}
```

```js [sol-JavaScript]
var maxMoves = function(grid) {
    const m = grid.length, n = grid[0].length;
    let ans = 0;
    function dfs(i, j) {
        ans = Math.max(ans, j);
        if (ans === n - 1) { // ans 已达到最大值
            return;
        }
        // 向右上/右/右下走一步
        for (let k = Math.max(i - 1, 0); k < Math.min(i + 2, m); k++) {
            if (grid[k][j + 1] > grid[i][j]) {
                dfs(k, j + 1);
            }
        }
        grid[i][j] = 0;
    }
    for (let i = 0; i < m; i++) {
        dfs(i, 0); // 从第一列的任一单元格出发
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_moves(mut grid: Vec<Vec<i32>>) -> i32 {
        fn dfs(i: usize, j: usize, ans: &mut usize, grid: &mut Vec<Vec<i32>>) {
            if j > *ans {
                *ans = j
            }
            if *ans == grid[0].len() - 1 { // ans 已达到最大值
                return;
            }
            // 向右上/右/右下走一步
            for k in i.saturating_sub(1)..grid.len().min(i + 2) {
                if grid[k][j + 1] > grid[i][j] {
                    dfs(k, j + 1, ans, grid);
                }
            }
            grid[i][j] = 0;
        }
        let mut ans = 0;
        for i in 0..grid.len() {
            dfs(i, 0, &mut ans, &mut grid); // 从第一列的任一单元格出发
        }
        ans as i32
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。每个格子至多递归一次。
- 空间复杂度：$\mathcal{O}(n)$。递归需要 $\mathcal{O}(n)$ 的栈空间。

## 方法二：网格图 BFS

双数组写法，见 [视频讲解](https://www.bilibili.com/video/BV1hG4y1277i/)。

一开始把所有 $(i,0)$ 都加入一个列表。每一轮循环，遍历列表中的坐标，把右边这一列的能到达的格子坐标加入一个新列表。注意只有之前没入队的格子才能入队，可以用一个 $\textit{vis}$ 数组，在入队前标记这个格子。

返回可以到达的最大列号，即为最大移动次数。

```py [sol-Python3]
class Solution:
    def maxMoves(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        vis = [-1] * m
        q = range(m)
        for j in range(n - 1):
            tmp = q
            q = []
            for i in tmp:
                for k in i - 1, i, i + 1:
                    if 0 <= k < m and vis[k] != j and grid[k][j + 1] > grid[i][j]:
                        vis[k] = j  # 第 k 行目前最右访问到了 j
                        q.append(k)
            if not q:  # 无法再往右走了
                return j
        return n - 1
```

```java [sol-Java]
class Solution {
    public int maxMoves(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        int[] vis = new int[m];
        Arrays.fill(vis, -1);
        List<Integer> q = new ArrayList<>(m);
        for (int i = 0; i < m; i++) {
            q.add(i);
        }
        for (int j = 0; j < n - 1; j++) {
            List<Integer> nxt = new ArrayList<>();
            for (int i : q) {
                for (int k = Math.max(i - 1, 0); k < Math.min(i + 2, m); k++) {
                    if (vis[k] != j && grid[k][j + 1] > grid[i][j]) {
                        vis[k] = j; // 第 k 行目前最右访问到了 j
                        nxt.add(k);
                    }
                }
            }
            if (nxt.isEmpty()) { // 无法再往右走了
                return j;
            }
            q = nxt;
        }
        return n - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxMoves(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size();
        vector<int> vis(m, -1), q(m);
        iota(q.begin(), q.end(), 0);
        for (int j = 0; j < n - 1; j++) {
            vector<int> nxt;
            for (int i : q) {
                for (int k = max(i - 1, 0); k < min(i + 2, m); k++) {
                    if (vis[k] != j && grid[k][j + 1] > grid[i][j]) {
                        vis[k] = j; // 第 k 行目前最右访问到了 j
                        nxt.push_back(k);
                    }
                }
            }
            if (nxt.empty()) { // 无法再往右走了
                return j;
            }
            q = move(nxt);
        }
        return n - 1;
    }
};
```

```go [sol-Go]
func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	vis := make([]int, m)
	q := make([]int, m)
	for i := range q {
		q[i] = i
	}
	for j := 0; j < n-1; j++ {
		tmp := q
		q = nil
		for _, i := range tmp {
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				if vis[k] != j+1 && grid[k][j+1] > grid[i][j] {
					vis[k] = j + 1 // 第 k 行目前最右访问到了 j
					q = append(q, k)
				}
			}
		}
		if q == nil { // 无法再往右走了
			return j
		}
	}
	return n - 1
}
```

```js [sol-JavaScript]
var maxMoves = function(grid) {
    const m = grid.length, n = grid[0].length;
    const vis = Array(m).fill(-1);
    let q = [...Array(m).keys()];
    for (let j = 0; j < n - 1; j++) {
        let nxt = [];
        for (const i of q) {
            for (let k = Math.max(i - 1, 0); k < Math.min(i + 2, m); k++) {
                if (vis[k] !== j && grid[k][j + 1] > grid[i][j]) {
                    vis[k] = j; // 第 k 行目前最右访问到了 j
                    nxt.push(k);
                }
            }
        }
        if (nxt.length === 0) { // 无法再往右走了
            return j;
        }
        q = nxt;
    }
    return n - 1;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_moves(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        let mut vis = vec![n; m];
        let mut q = (0..m).collect::<Vec<_>>();
        for j in 0..n - 1 {
            let mut nxt = Vec::new();
            for &i in &q {
                for k in i.saturating_sub(1)..m.min(i + 2) {
                    if vis[k] != j && grid[k][j + 1] > grid[i][j] {
                        vis[k] = j; // 第 k 行目前最右访问到了 j
                        nxt.push(k);
                    }
                }
            }
            if nxt.is_empty() { // 无法再往右走了
                return j as i32;
            }
            q = nxt;
        }
        (n - 1) as i32
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。每个格子至多入队一次。
- 空间复杂度：$\mathcal{O}(m)$。

## 空间优化

由于 $\textit{grid}[i][j]>0$，我们可以把要入队的格子值变为其相反数，从而判断哪些格子在队列中。此外，一个数一旦变成负数就不会比当前格子值大了，这可以保证一个格子值只会被标记（改成相反数）一次。

这样就不需要队列和 $\textit{vis}$ 数组了。

```py [sol-Python3]
class Solution:
    def maxMoves(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        for row in grid:
            row[0] *= -1  # 入队标记
        for j in range(n - 1):
            for i, row in enumerate(grid):
                if row[j] > 0:  # 不在队列中
                    continue
                for k in i - 1, i, i + 1:
                    if 0 <= k < m and grid[k][j + 1] > -row[j]:
                        grid[k][j + 1] *= -1  # 入队标记
            if all(row[j + 1] > 0 for row in grid):  # 无法再往右走了
                return j
        return n - 1
```

```java [sol-Java]
class Solution {
    public int maxMoves(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        for (int[] row : grid) {
            row[0] *= -1; // 入队标记
        }
        for (int j = 0; j < n - 1; j++) {
            boolean ok = false;
            for (int i = 0; i < m; i++) {
                if (grid[i][j] > 0) { // 不在队列中
                    continue;
                }
                for (int k = Math.max(i - 1, 0); k < Math.min(i + 2, m); k++) {
                    if (grid[k][j + 1] > -grid[i][j]) {
                        grid[k][j + 1] *= -1; // 入队标记
                        ok = true;
                    }
                }
            }
            if (!ok) { // 无法再往右走了
                return j;
            }
        }
        return n - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxMoves(vector<vector<int>> &grid) {
        int m = grid.size(), n = grid[0].size();
        for (auto &row: grid) {
            row[0] *= -1; // 入队标记
        }
        for (int j = 0; j < n - 1; j++) {
            bool ok = false;
            for (int i = 0; i < m; i++) {
                if (grid[i][j] > 0) continue; // 不在队列中
                for (int k = max(i - 1, 0); k < min(i + 2, m); k++) {
                    if (grid[k][j + 1] > -grid[i][j]) {
                        grid[k][j + 1] *= -1; // 入队标记
                        ok = true;
                    }
                }
            }
            if (!ok) { // 无法再往右走了
                return j;
            }
        }
        return n - 1;
    }
};
```

```go [sol-Go]
func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for _, row := range grid {
		row[0] *= -1 // 入队标记
	}
	for j := 0; j < n-1; j++ {
		ok := false
		for i := 0; i < m; i++ {
			if grid[i][j] > 0 { // 不在队列中
				continue
			}
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				if grid[k][j+1] > -grid[i][j] {
					grid[k][j+1] *= -1 // 入队标记
					ok = true
				}
			}
		}
		if !ok { // 无法再往右走了
			return j
		}
	}
	return n - 1
}
```

```js [sol-JavaScript]
var maxMoves = function(grid) {
    const m = grid.length, n = grid[0].length;
    for (const row of grid) {
        row[0] *= -1; // 入队标记
    }
    for (let j = 0; j < n - 1; j++) {
        let ok = false;
        for (let i = 0; i < m; i++) {
            if (grid[i][j] > 0) continue; // 不在队列中
            for (let k = Math.max(i - 1, 0); k < Math.min(i + 2, m); k++) {
                if (grid[k][j + 1] > -grid[i][j]) {
                    grid[k][j + 1] *= -1; // 入队标记
                    ok = true;
                }
            }
        }
        if (!ok) { // 无法再往右走了
            return j;
        }
    }
    return n - 1;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_moves(mut grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        for row in grid.iter_mut() {
            row[0] *= -1; // 入队标记
        }
        for j in 0..n - 1 {
            let mut ok = false;
            for i in 0..m {
                if grid[i][j] > 0 { // 不在队列中
                    continue;
                }
                for k in i.saturating_sub(1)..m.min(i + 2) {
                    if grid[k][j + 1] > -grid[i][j] {
                        grid[k][j + 1] *= -1; // 入队标记
                        ok = true;
                    }
                }
            }
            if !ok { // 无法再往右走了
                return j as i32;
            }
        }
        (n - 1) as i32
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

## 题单

- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- 更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
