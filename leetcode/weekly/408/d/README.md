如果从矩形【上边界/左边界】到矩形【右边界/下边界】的路被圆堵死，则无法从矩形左下角移动到矩形右上角。

怎么判断呢？

首先考虑圆心都在矩形内部的情况。如果圆和圆相交或相切，则相当于在两个圆之间架起了一座桥。如果圆和矩形边界相交或相切，则相当于在矩形边界和圆之间架起了一座桥。如果可以从矩形【上边界/左边界】通过桥到达矩形【右边界/下边界】，则说明路被堵死，无法从矩形左下角移动到矩形右上角。

也可以把桥理解成切割线，如果能把从矩形左下角到矩形右上角的路径**切断**，则无法从矩形左下角移动到矩形右上角。

用图论的术语来说，就是把圆抽象成节点，在相交或相切的节点之间连边，得到一张无向图。如果从与【上边界/左边界】相交的节点出发，**DFS** 这张图，到达与【右边界/下边界】相交的节点，则说明无法从矩形左下角移动到矩形右上角。

需要注意，本题没有保证圆心一定在矩形内部，如何处理这种情况呢？

![lc3235-c.png](https://pic.leetcode.cn/1722649636-dihkoU-lc3235-c.png)

注：把两圆的两个交点连起来，该线段与 $O_1O_2$ 相交得到的交点作为点 $A$ 也可以，但这种情况点 $A$ 横纵坐标的分母会是一个 $10^{18}$ 数量级的数，在与 $\textit{X}$ 或 $\textit{Y}$ 相乘时会产生 $10^{27}$ 数量级的数，超出了 64 位整数的范围，需要用大整数实现，更麻烦。

如何判断圆是否与矩形边界相交相切？

![lc3235-2-c.png](https://pic.leetcode.cn/1722579370-cPlOGI-lc3235-2-c.png)

⚠**注意**：$y>Y$ 的情况属于情况一，实际上无需判断。

### 具体做法

从与矩形【上边界/左边界】相交/相切的圆开始 DFS。

如果当前 DFS 到了圆 $i$：

- 先判断其是否与矩形【右边界/下边界】相交或相切，如果是，则 DFS 返回 $\texttt{true}$。
- 否则，判断其是否与其他圆 $j$ 相交或相切，如果是，则判断点 $A$ 是否严格在矩形内，如果在，则递归 $j$，如果收到了 $\texttt{true}$，则 DFS 返回 $\texttt{true}$。

最后，如果最外层调用 DFS 的地方收到了 $\texttt{true}$，则表示无法从矩形左下角移动到矩形右上角，返回 $\texttt{false}$。

代码实现时，可以在递归之前，特判圆包含矩形左下角或者矩形右上角的情况，此时可以直接返回 $\texttt{false}$。

```py [sol-Python3]
class Solution:
    def canReachCorner(self, X: int, Y: int, circles: List[List[int]]) -> bool:
        vis = [False] * len(circles)
        def dfs(i: int) -> bool:
            x1, y1, r1 = circles[i]
            # 圆 i 是否与矩形右边界/下边界相交相切
            if y1 <= Y and abs(x1 - X) <= r1 or x1 <= X and y1 <= r1:
                return True
            vis[i] = True
            for j, (x2, y2, r2) in enumerate(circles):
                # 在两圆相交相切的前提下，点 A 是否严格在矩形内
                if not vis[j] and \
                   (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2) <= (r1 + r2) * (r1 + r2) and \
                   x1 * r2 + x2 * r1 < (r1 + r2) * X and \
                   y1 * r2 + y2 * r1 < (r1 + r2) * Y and \
                   dfs(j):
                    return True
            return False

        for i, (x, y, r) in enumerate(circles):
            # 圆 i 包含矩形左下角 or
            # 圆 i 包含矩形右上角 or
            # 圆 i 与矩形上边界/左边界相交相切
            if x * x + y * y <= r * r or \
               (x - X) * (x - X) + (y - Y) * (y - Y) <= r * r or \
               not vis[i] and (x <= X and abs(y - Y) <= r or y <= Y and x <= r) and dfs(i):
                return False
        return True
```

```java [sol-Java]
class Solution {
    public boolean canReachCorner(int X, int Y, int[][] circles) {
        boolean[] vis = new boolean[circles.length];
        for (int i = 0; i < circles.length; i++) {
            long x = circles[i][0], y = circles[i][1], r = circles[i][2];
            if (inCircle(x, y, r, 0, 0) || // 圆 i 包含矩形左下角
                inCircle(x, y, r, X, Y) || // 圆 i 包含矩形右上角
                // 圆 i 是否与矩形上边界/左边界相交相切
                !vis[i] && (x <= X && Math.abs(y - Y) <= r || y <= Y && x <= r) && dfs(i, X, Y, circles, vis)) {
                return false;
            }
        }
        return true;
    }

    // 判断点 (x,y) 是否在圆 (ox,oy,r) 内
    private boolean inCircle(long ox, long oy, long r, long x, long y) {
        return (ox - x) * (ox - x) + (oy - y) * (oy - y) <= r * r;
    }

    private boolean dfs(int i, int X, int Y, int[][] circles, boolean[] vis) {
        long x1 = circles[i][0], y1 = circles[i][1], r1 = circles[i][2];
        // 圆 i 是否与矩形右边界/下边界相交相切
        if (y1 <= Y && Math.abs(x1 - X) <= r1 || x1 <= X && y1 <= r1) {
            return true;
        }
        vis[i] = true;
        for (int j = 0; j < circles.length; j++) {
            long x2 = circles[j][0], y2 = circles[j][1], r2 = circles[j][2];
            // 在两圆相交相切的前提下，点 A 是否严格在矩形内
            if (!vis[j] &&
                (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2) <= (r1 + r2) * (r1 + r2) &&
                x1 * r2 + x2 * r1 < (r1 + r2) * X &&
                y1 * r2 + y2 * r1 < (r1 + r2) * Y &&
                dfs(j, X, Y, circles, vis)) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 判断点 (x,y) 是否在圆 (ox,oy,r) 内
    bool in_circle(long long ox, long long oy, long long r, long long x, long long y) {
        return (ox - x) * (ox - x) + (oy - y) * (oy - y) <= r * r;
    }

public:
    bool canReachCorner(int X, int Y, vector<vector<int>>& circles) {
        int n = circles.size();
        vector<int> vis(n);
        auto dfs = [&](auto&& dfs, int i) -> bool {
            long long x1 = circles[i][0], y1 = circles[i][1], r1 = circles[i][2];
            // 圆 i 是否与矩形右边界/下边界相交相切
            if (y1 <= Y && abs(x1 - X) <= r1 || x1 <= X && y1 <= r1) {
                return true;
            }
            vis[i] = true;
            for (int j = 0; j < n; j++) {
                long long x2 = circles[j][0], y2 = circles[j][1], r2 = circles[j][2];
                // 在两圆相交相切的前提下，点 A 是否严格在矩形内
                if (!vis[j] && (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2) <= (r1 + r2) * (r1 + r2) &&
                    x1 * r2 + x2 * r1 < (r1 + r2) * X &&
                    y1 * r2 + y2 * r1 < (r1 + r2) * Y &&
                    dfs(dfs, j)) {
                    return true;
                }
            }
            return false;
        };
        for (int i = 0; i < n; i++) {
            long long x = circles[i][0], y = circles[i][1], r = circles[i][2];
            if (in_circle(x, y, r, 0, 0) || // 圆 i 包含矩形左下角
                in_circle(x, y, r, X, Y) || // 圆 i 包含矩形右上角
                // 圆 i 是否与矩形上边界/左边界相交相切
                !vis[i] && (x <= X && abs(y - Y) <= r || y <= Y && x <= r) && dfs(dfs, i)) {
                return false;
            }
        }
        return true;
    }
};
```

```c [sol-C]
// 判断点 (x,y) 是否在圆 (ox,oy,r) 内
bool inCircle(long long ox, long long oy, long long r, long long x, long long y) {
    return (ox - x) * (ox - x) + (oy - y) * (oy - y) <= r * r;
}

bool dfs(int i, int X, int Y, int** circles, int circlesSize, bool* vis) {
    long long x1 = circles[i][0], y1 = circles[i][1], r1 = circles[i][2];
    // 圆 i 是否与矩形右边界/下边界相交相切
    if (y1 <= Y && abs(x1 - X) <= r1 || x1 <= X && y1 <= r1) {
        return true;
    }
    vis[i] = true;
    for (int j = 0; j < circlesSize; j++) {
        long long x2 = circles[j][0], y2 = circles[j][1], r2 = circles[j][2];
        // 在两圆相交相切的前提下，点 A 是否严格在矩形内
        if (!vis[j] && (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2) <= (r1 + r2) * (r1 + r2) &&
            x1 * r2 + x2 * r1 < (r1 + r2) * X &&
            y1 * r2 + y2 * r1 < (r1 + r2) * Y &&
            dfs(j, X, Y, circles, circlesSize, vis)) {
            return true;
        }
    }
    return false;
}

bool canReachCorner(int X, int Y, int** circles, int circlesSize, int* circlesColSize) {
    bool* vis = calloc(circlesSize, sizeof(bool));
    for (int i = 0; i < circlesSize; i++) {
        long long x = circles[i][0], y = circles[i][1], r = circles[i][2];
        if (inCircle(x, y, r, 0, 0) || // 圆 i 包含矩形左下角
            inCircle(x, y, r, X, Y) || // 圆 i 包含矩形右上角
            // 圆 i 是否与矩形上边界/左边界相交相切
            !vis[i] && (x <= X && abs(y - Y) <= r || y <= Y && x <= r) && dfs(i, X, Y, circles, circlesSize, vis)) {
            free(vis);
            return false;
        }
    }
    free(vis);
    return true;
}
```

```go [sol-Go]
// 判断点 (x,y) 是否在圆 (ox,oy,r) 内
func inCircle(ox, oy, r, x, y int) bool {
    return (ox-x)*(ox-x)+(oy-y)*(oy-y) <= r*r
}

func canReachCorner(X, Y int, circles [][]int) bool {
    vis := make([]bool, len(circles))
    var dfs func(int) bool
    dfs = func(i int) bool {
        x1, y1, r1 := circles[i][0], circles[i][1], circles[i][2]
        // 圆 i 是否与矩形右边界/下边界相交相切
        if y1 <= Y && abs(x1-X) <= r1 || x1 <= X && y1 <= r1 {
            return true
        }
        vis[i] = true
        for j, c := range circles {
            x2, y2, r2 := c[0], c[1], c[2]
            // 在两圆相交相切的前提下，点 A 是否严格在矩形内
            if !vis[j] && (x1-x2)*(x1-x2)+(y1-y2)*(y1-y2) <= (r1+r2)*(r1+r2) &&
                x1*r2+x2*r1 < (r1+r2)*X &&
                y1*r2+y2*r1 < (r1+r2)*Y &&
                dfs(j) {
                return true
            }
        }
        return false
    }
    for i, c := range circles {
        x, y, r := c[0], c[1], c[2]
        if inCircle(x, y, r, 0, 0) || // 圆 i 包含矩形左下角
            inCircle(x, y, r, X, Y) || // 圆 i 包含矩形右上角
            // 圆 i 是否与矩形上边界/左边界相交相切
            !vis[i] && (x <= X && abs(y-Y) <= r || y <= Y && x <= r) && dfs(i) {
            return false
        }
    }
    return true
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```js [sol-JavaScript]
var canReachCorner = function(X, Y, circles) {
    // 判断点 (x, y) 是否在圆 (ox, oy, r) 内
    function inCircle(ox, oy, r, x, y) {
        return BigInt(ox - x) * BigInt(ox - x) +
               BigInt(oy - y) * BigInt(oy - y) <= BigInt(r) * BigInt(r);
    }

    const BX = BigInt(X), BY = BigInt(Y);
    const vis = new Array(circles.length).fill(false);
    function dfs(i) {
        let [x1, y1, r1] = circles[i];
        // 圆 i 是否与矩形右边界/下边界相交相切
        if (y1 <= Y && Math.abs(x1 - X) <= r1 || x1 <= X && y1 <= r1) {
            return true;
        }
        x1 = BigInt(x1);
        y1 = BigInt(y1);
        r1 = BigInt(r1);
        vis[i] = true;
        for (let j = 0; j < circles.length; j++) {
            if (!vis[j]) {
                let [x2, y2, r2] = circles[j];
                x2 = BigInt(x2);
                y2 = BigInt(y2);
                r2 = BigInt(r2);
                // 在两圆相交相切的前提下，点 A 是否严格在矩形内
                if ((x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2) <= (r1 + r2) * (r1 + r2) &&
                    x1 * r2 + x2 * r1 < (r1 + r2) * BX &&
                    y1 * r2 + y2 * r1 < (r1 + r2) * BY &&
                    dfs(j)) {
                    return true;
                }
            }
        }
        return false;
    }

    for (let i = 0; i < circles.length; i++) {
        const [x, y, r] = circles[i];
        if (inCircle(x, y, r, 0, 0) || // 圆 i 包含矩形左下角
            inCircle(x, y, r, X, Y) || // 圆 i 包含矩形右上角
            // 圆 i 是否与矩形上边界/左边界相交相切
            !vis[i] && (x <= X && Math.abs(y - Y) <= r || y <= Y && x <= r) && dfs(i)) {
            return false;
        }
    }
    return true;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn can_reach_corner(x_corner: i32, y_corner: i32, circles: Vec<Vec<i32>>) -> bool {
        let X = x_corner as i64;
        let Y = y_corner as i64;
        let mut vis = vec![false; circles.len()];
        for i in 0..circles.len() {
            let x = circles[i][0] as i64;
            let y = circles[i][1] as i64;
            let r = circles[i][2] as i64;
            if Self::in_circle(x, y, r, 0, 0) || // 圆 i 包含矩形左下角
               Self::in_circle(x, y, r, X, Y) || // 圆 i 包含矩形右上角
               // 圆 i 是否与矩形上边界/左边界相交相切
               !vis[i] && (x <= X && (y - Y).abs() <= r || y <= Y && x <= r) && Self::dfs(i, X, Y, &circles, &mut vis) {
                return false;
            }
        }
        true
    }

    // 判断点 (x,y) 是否在圆 (ox,oy,r) 内
    fn in_circle(ox: i64, oy: i64, r: i64, x: i64, y: i64) -> bool {
        (ox - x) * (ox - x) + (oy - y) * (oy - y) <= r * r
    }

    fn dfs(i: usize, x: i64, y: i64, circles: &Vec<Vec<i32>>, vis: &mut Vec<bool>) -> bool {
        let x1 = circles[i][0] as i64;
        let y1 = circles[i][1] as i64;
        let r1 = circles[i][2] as i64;
        // 圆 i 是否与矩形右边界/下边界相交相切
        if y1 <= y && (x1 - x).abs() <= r1 || x1 <= x && y1 <= r1 {
            return true;
        }
        vis[i] = true;
        for (j, c2) in circles.iter().enumerate() {
            let x2 = c2[0] as i64;
            let y2 = c2[1] as i64;
            let r2 = c2[2] as i64;
            // 在两圆相交相切的前提下，点 A 是否严格在矩形内
            if !vis[j] && (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2) <= (r1 + r2) * (r1 + r2) &&
               x1 * r2 + x2 * r1 < (r1 + r2) * x &&
               y1 * r2 + y2 * r1 < (r1 + r2) * y &&
               Self::dfs(j, x, y, circles, vis) {
                return true;
            }
        }
        false
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{circles}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

注：本题也可以用并查集实现，但效率不如 DFS。

更多相似题目，见下面图论题单中的「**DFS**」和数据结构题单中的「**并查集**」。

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
