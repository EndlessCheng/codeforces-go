标记所有 $\textit{edges}[i][1]$，这些队伍都不是冠军。

然后遍历每个节点 $i$，如果恰好有一个 $i$ 没有被标记，说明没有可以击败 $i$ 的队伍，$i$ 队是冠军。否则返回 $-1$。

附：[视频讲解](https://www.bilibili.com/video/BV1Fc411R7xA/) 第二题。

```py [sol-Python3]
class Solution:
    def findChampion(self, n: int, edges: List[List[int]]) -> int:
        is_weak = [False] * n
        for _, y in edges:
            is_weak[y] = True  # 不是冠军

        ans = -1
        for i, weak in enumerate(is_weak):
            if weak:
                continue
            if ans != -1:
                return -1  # 冠军只能有一个
            ans = i
        return ans
```

```java [sol-Java]
class Solution {
    public int findChampion(int n, int[][] edges) {
        boolean[] isWeak = new boolean[n];
        for (int[] e : edges) {
            isWeak[e[1]] = true; // 不是冠军
        }

        int ans = -1;
        for (int i = 0; i < n; i++) {
            if (isWeak[i]) {
                continue;
            }
            if (ans != -1) {
                return -1; // 冠军只能有一个
            }
            ans = i;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findChampion(int n, vector<vector<int>> &edges) {
        vector<int> is_weak(n);
        for (auto &e : edges) {
            is_weak[e[1]] = true; // 不是冠军
        }

        int ans = -1;
        for (int i = 0; i < n; i++) {
            if (is_weak[i]) {
                continue;
            }
            if (ans != -1) {
                return -1; // 冠军只能有一个
            }
            ans = i;
        }
        return ans;
    }
};
```

```c [sol-C]
int findChampion(int n, int** edges, int edgesSize, int* edgesColSize) {
    bool* is_weak = calloc(n, sizeof(bool));
    for (int i = 0; i < edgesSize; i++) {
        is_weak[edges[i][1]] = true; // 不是冠军
    }

    int ans = -1;
    for (int i = 0; i < n; i++) {
        if (is_weak[i]) {
            continue;
        }
        if (ans != -1) {
            free(is_weak);
            return -1; // 冠军只能有一个
        }
        ans = i;
    }
    free(is_weak);
    return ans;
}
```

```go [sol-Go]
func findChampion(n int, edges [][]int) int {
    weak := make([]bool, n)
    for _, e := range edges {
        weak[e[1]] = true // 不是冠军
    }

    ans := -1
    for i, w := range weak {
        if w {
            continue
        }
        if ans != -1 {
            return -1 // 冠军只能有一个
        }
        ans = i
    }
    return ans
}
```

```js [sol-JavaScript]
var findChampion = function(n, edges) {
    const isWeak = Array(n).fill(false);
    for (const [, y] of edges) {
        isWeak[y] = true; // 不是冠军
    }

    let ans = -1;
    for (let i = 0; i < n; i++) {
        if (isWeak[i]) {
            continue;
        }
        if (ans !== -1) {
            return -1; // 冠军只能有一个
        }
        ans = i;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_champion(n: i32, edges: Vec<Vec<i32>>) -> i32 {
        let mut is_weak = vec![false; n as usize];
        for e in &edges {
            is_weak[e[1] as usize] = true; // 不是冠军
        }

        let mut ans = -1;
        for (i, &weak) in is_weak.iter().enumerate() {
            if weak {
                continue;
            }
            if ans != -1 {
                return -1; // 冠军只能有一个
            }
            ans = i as i32;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
