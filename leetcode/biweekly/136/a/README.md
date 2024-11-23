## 方法一：两次遍历

遍历 $\textit{pick}$，用一个 $n\times 11$ 大小的矩阵，统计每个玩家得到的每种颜色的球的个数。

然后遍历每个玩家，如果该玩家至少有一种颜色的球大于玩家编号，则把答案加一。

```py [sol-Python3]
class Solution:
    def winningPlayerCount(self, n: int, pick: List[List[int]]) -> int:
        cnts = [[0] * 11 for _ in range(n)]
        for x, y in pick:
            cnts[x][y] += 1

        ans = 0
        for i, cnt in enumerate(cnts):
            if any(c > i for c in cnt):
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int winningPlayerCount(int n, int[][] pick) {
        int[][] cnts = new int[n][11];
        for (int[] p : pick) {
            cnts[p[0]][p[1]]++;
        }

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int c : cnts[i]) {
                if (c > i) {
                    ans++;
                    break;
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
    int winningPlayerCount(int n, vector<vector<int>>& pick) {
        vector<array<int, 11>> cnts(n);
        for (auto& p : pick) {
            cnts[p[0]][p[1]]++;
        }

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int c : cnts[i]) {
                if (c > i) {
                    ans++;
                    break;
                }
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int winningPlayerCount(int n, int** pick, int pickSize, int* pickColSize) {
    int (*cnts)[11] = calloc(n, sizeof(int[11]));
    for (int i = 0; i < pickSize; i++) {
        int* p = pick[i];
        cnts[p[0]][p[1]]++;
    }

    int ans = 0;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < 11; j++) {
            if (cnts[i][j] > i) {
                ans++;
                break;
            }
        }
    }

    free(cnts);
    return ans;
}
```

```go [sol-Go]
func winningPlayerCount(n int, pick [][]int) (ans int) {
    cnts := make([][11]int, n)
    for _, p := range pick {
        cnts[p[0]][p[1]]++
    }

    for i, cnt := range cnts {
        for _, c := range cnt {
            if c > i {
                ans++
                break
            }
        }
    }
    return
}
```

```js [sol-JavaScript]
var winningPlayerCount = function(n, pick) {
    const cnts = Array.from({ length: n }, () => Array(11).fill(0));
    for (const [x, y] of pick) {
        cnts[x][y]++;
    }

    let ans = 0;
    for (let i = 0; i < n; i++) {
        if (cnts[i].some(c => c > i)) {
            ans++;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn winning_player_count(n: i32, pick: Vec<Vec<i32>>) -> i32 {
        let mut cnts = vec![[0; 11]; n as usize];
        for p in pick {
            cnts[p[0] as usize][p[1] as usize] += 1;
        }

        let mut ans = 0;
        for (i, cnt) in cnts.iter().enumerate() {
            if cnt.iter().any(|&c| c > i) {
                ans += 1;
            }
        }
        ans
    }
}
```

```rust [sol-Rust 写法二]
impl Solution {
    pub fn winning_player_count(n: i32, pick: Vec<Vec<i32>>) -> i32 {
        let mut cnts = vec![[0; 11]; n as usize];
        for p in pick {
            cnts[p[0] as usize][p[1] as usize] += 1;
        }

        cnts.iter()
            .enumerate()
            .filter(|(i, cnt)| cnt.iter().any(|c| c > i))
            .count() as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU+m)$，其中 $m$ 是 $\textit{pick}$ 的长度，$U$ 是 $y_i$ 的最大值。
- 空间复杂度：$\mathcal{O}(nU)$。

## 方法二：一次遍历

额外创建一个布尔数组 $\textit{won}$。

在统计 $\textit{cnts}[x][y]$ 的过程中，如果发现玩家 $x$ 获得了至少 $x + 1$ 个相同颜色的球，那么答案加一，同时标记 $\textit{won}[x]=\texttt{true}$，避免重复计入答案。

```py [sol-Python3]
class Solution:
    def winningPlayerCount(self, n: int, pick: List[List[int]]) -> int:
        ans = 0
        cnts = [[0] * 11 for _ in range(n)]
        won = [False] * n
        for x, y in pick:
            cnts[x][y] += 1
            if not won[x] and cnts[x][y] > x:
                won[x] = True
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int winningPlayerCount(int n, int[][] pick) {
        int ans = 0;
        int[][] cnts = new int[n][11];
        boolean[] won = new boolean[n];
        for (int[] p : pick) {
            int x = p[0];
            int y = p[1];
            if (!won[x] && ++cnts[x][y] > x) {
                won[x] = true;
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int winningPlayerCount(int n, vector<vector<int>>& pick) {
        int ans = 0;
        vector<array<int, 11>> cnts(n);
        vector<int> won(n);
        for (auto& p : pick) {
            int x = p[0], y = p[1];
            if (!won[x] && ++cnts[x][y] > x) {
                won[x] = true;
                ans++;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int winningPlayerCount(int n, int** pick, int pickSize, int* pickColSize) {
    int ans = 0;
    int (*cnts)[11] = calloc(n, sizeof(int[11]));
    bool* won = calloc(n, sizeof(bool));
    for (int i = 0; i < pickSize; i++) {
        int x = pick[i][0], y = pick[i][1];
        if (!won[x] && ++cnts[x][y] > x) {
            won[x] = true;
            ans++;
        }
    }
    free(cnts);
    free(won);
    return ans;
}
```

```go [sol-Go]
func winningPlayerCount(n int, pick [][]int) (ans int) {
    cnts := make([][11]int, n)
    won := make([]bool, n)
    for _, p := range pick {
        x, y := p[0], p[1]
        cnts[x][y]++
        if !won[x] && cnts[x][y] > x {
            won[x] = true
            ans++
        }
    }
    return
}
```

```js [sol-JavaScript]
var winningPlayerCount = function(n, pick) {
    let ans = 0;
    const cnts = Array.from({ length: n }, () => Array(11).fill(0));
    const won = Array(n).fill(false);
    for (const [x, y] of pick) {
        if (!won[x] && ++cnts[x][y] > x) {
            won[x] = true;
            ans++;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn winning_player_count(n: i32, pick: Vec<Vec<i32>>) -> i32 {
        let mut ans = 0;
        let mut cnts = vec![[0; 11]; n as usize];
        let mut won = vec![false; n as usize];
        for p in pick {
            let x = p[0] as usize;
            let y = p[1] as usize;
            cnts[x][y] += 1;
            if !won[x] && cnts[x][y] > x {
                won[x] = true;
                ans += 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU+m)$，其中 $m$ 是 $\textit{pick}$ 的长度，$U$ 是 $y_i$ 的最大值。
- 空间复杂度：$\mathcal{O}(nU)$。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
