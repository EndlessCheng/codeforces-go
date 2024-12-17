由于矩形的高没有限制，所以我们只需考虑点的**横坐标**。

矩形越宽，覆盖的点越多，所以 $x_2$ 应该**恰好等于** $x_1+w$。

算法如下：

1. 把横坐标按照从小到大的顺序排序。
2. 为方便计算，假设第一个矩形左边还有一个矩形，初始化 $x_2 = -1$，因为所有横坐标都是非负数。
3. 遍历横坐标 $x=\textit{points}[i][0]$，如果 $x > x_2$，我们需要一个新的 $x_1=x$ 的矩形，答案加一，然后把 $x_2$ 更新为 $x+w$。
4. 遍历结束，返回答案。

[视频讲解](https://www.bilibili.com/video/BV1et42177VM/) 第二题。

```py [sol-Python3]
class Solution:
    def minRectanglesToCoverPoints(self, points: List[List[int]], w: int) -> int:
        points.sort(key=lambda p: p[0])
        ans = 0
        x2 = -1
        for x, _ in points:
            if x > x2:
                ans += 1
                x2 = x + w
        return ans
```

```java [sol-Java]
class Solution {
    public int minRectanglesToCoverPoints(int[][] points, int w) {
        Arrays.sort(points, (p, q) -> p[0] - q[0]);
        int ans = 0;
        int x2 = -1;
        for (int[] p : points) {
            if (p[0] > x2) {
                ans++;
                x2 = p[0] + w;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minRectanglesToCoverPoints(vector<vector<int>>& points, int w) {
        ranges::sort(points, {}, [](auto& p) { return p[0]; });
        int ans = 0;
        int x2 = -1;
        for (auto& p : points) {
            if (p[0] > x2) {
                ans++;
                x2 = p[0] + w;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return (*(int**)a)[0] - (*(int**)b)[0];
}

int minRectanglesToCoverPoints(int** points, int pointsSize, int* pointsColSize, int w) {
    qsort(points, pointsSize, sizeof(int*), cmp);
    int ans = 0;
    int x2 = -1;
    for (int i = 0; i < pointsSize; i++) {
        if (points[i][0] > x2) {
            ans++;
            x2 = points[i][0] + w;
        }
    }
    return ans;
}
```

```go [sol-Go]
func minRectanglesToCoverPoints(points [][]int, w int) (ans int) {
	slices.SortFunc(points, func(p, q []int) int { return p[0] - q[0] })
	x2 := -1
	for _, p := range points {
		if p[0] > x2 {
			ans++
			x2 = p[0] + w
		}
	}
	return
}
```

```js [sol-JavaScript]
var minRectanglesToCoverPoints = function(points, w) {
    points.sort((p, q) => p[0] - q[0]);
    let ans = 0;
    let x2 = -1;
    for (const [x, _] of points) {
        if (x > x2) {
            ans++;
            x2 = x + w;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_rectangles_to_cover_points(mut points: Vec<Vec<i32>>, w: i32) -> i32 {
        points.sort_unstable_by(|p, q| p[0].cmp(&q[0]));
        let mut ans = 0;
        let mut x2 = -1;
        for p in points {
            if p[0] > x2 {
                ans += 1;
                x2 = p[0] + w;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{points}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

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
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
