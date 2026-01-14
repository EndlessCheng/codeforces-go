枚举从 $n$ 个矩形中选出两个矩形。

如果两个矩形有交集，那么交集也是矩形。

求出这个交集矩形的左下角和右上角，就可以算出交集矩形的长和宽。

- 左下角横坐标：两个矩形左下角横坐标的最大值。
- 左下角纵坐标：两个矩形左下角纵坐标的最大值。
- 右上角横坐标：两个矩形右上角横坐标的最小值。
- 右上角纵坐标：两个矩形右上角纵坐标的最小值。

长和宽的最小值，就是能放入交集区域的正方形的最大边长 $\textit{side}$。

最大的 $\textit{side}$ 的平方（正方形面积）即为答案。

## 优化前

```py [sol-Python3]
class Solution:
    def largestSquareArea(self, bottomLeft: List[List[int]], topRight: List[List[int]]) -> int:
        max_side = 0
        for i, ((bx, by), (tx, ty)) in enumerate(zip(bottomLeft, topRight)):
            for j in range(i):
                bx2, by2 = bottomLeft[j]
                tx2, ty2 = topRight[j]
                width = min(tx, tx2) - max(bx, bx2)  # 右上横坐标 - 左下横坐标
                height = min(ty, ty2) - max(by, by2)  # 右上纵坐标 - 左下纵坐标
                side = min(width, height)
                max_side = max(max_side, side)
        return max_side ** 2
```

```java [sol-Java]
class Solution {
    public long largestSquareArea(int[][] bottomLeft, int[][] topRight) {
        int maxSide = 0;
        for (int i = 0; i < bottomLeft.length; i++) {
            int[] b1 = bottomLeft[i];
            int[] t1 = topRight[i];
            for (int j = 0; j < i; j++) {
                int[] b2 = bottomLeft[j];
                int[] t2 = topRight[j];
                int width = Math.min(t1[0], t2[0]) - Math.max(b1[0], b2[0]);
                int height = Math.min(t1[1], t2[1]) - Math.max(b1[1], b2[1]);
                int side = Math.min(width, height);
                maxSide = Math.max(maxSide, side);
            }
        }
        return (long) maxSide * maxSide;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long largestSquareArea(vector<vector<int>>& bottomLeft, vector<vector<int>>& topRight) {
        int max_side = 0;
        for (int i = 0; i < bottomLeft.size(); i++) {
            auto& b1 = bottomLeft[i];
            auto& t1 = topRight[i];
            for (int j = 0; j < i; j++) {
                auto& b2 = bottomLeft[j];
                auto& t2 = topRight[j];
                int width = min(t1[0], t2[0]) - max(b1[0], b2[0]);
                int height = min(t1[1], t2[1]) - max(b1[1], b2[1]);
                int side = min(width, height);
                max_side = max(max_side, side);
            }
        }
        return 1LL * max_side * max_side;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))
#define MAX(a, b) ((b) > (a) ? (b) : (a))

long long largestSquareArea(int** bottomLeft, int bottomLeftSize, int* bottomLeftColSize, int** topRight, int topRightSize, int* topRightColSize) {
    int max_side = 0;
    for (int i = 0; i < bottomLeftSize; i++) {
        int* b1 = bottomLeft[i];
        int* t1 = topRight[i];
        for (int j = 0; j < i; j++) {
            int* b2 = bottomLeft[j];
            int* t2 = topRight[j];
            int width = MIN(t1[0], t2[0]) - MAX(b1[0], b2[0]);
            int height = MIN(t1[1], t2[1]) - MAX(b1[1], b2[1]);
            int side = MIN(width, height);
            max_side = MAX(max_side, side);
        }
    }
    return 1LL * max_side * max_side;
}
```

```go [sol-Go]
func largestSquareArea(bottomLeft, topRight [][]int) int64 {
	maxSide := 0
	for i, b1 := range bottomLeft {
		t1 := topRight[i]
		for j, b2 := range bottomLeft[:i] {
			t2 := topRight[j]
			width := min(t1[0], t2[0]) - max(b1[0], b2[0])
			height := min(t1[1], t2[1]) - max(b1[1], b2[1])
			side := min(width, height)
			maxSide = max(maxSide, side)
		}
	}
	return int64(maxSide) * int64(maxSide)
}
```

```js [sol-JavaScript]
var largestSquareArea = function(bottomLeft, topRight) {
    let maxSide = 0;
    for (let i = 0; i < bottomLeft.length; i++) {
        const [bx, by] = bottomLeft[i];
        const [tx, ty] = topRight[i];
        for (let j = 0; j < i; j++) {
            const [bx2, by2] = bottomLeft[j];
            const [tx2, ty2] = topRight[j];
            const width = Math.min(tx, tx2) - Math.max(bx, bx2); // 右上横坐标 - 左下横坐标
            const height = Math.min(ty, ty2) - Math.max(by, by2); // 右上纵坐标 - 左下纵坐标
            const side = Math.min(width, height);
            maxSide = Math.max(maxSide, side);
        }
    }
    return maxSide * maxSide;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn largest_square_area(bottom_left: Vec<Vec<i32>>, top_right: Vec<Vec<i32>>) -> i64 {
        let mut max_side = 0;
        for i in 0..bottom_left.len() {
            let b1 = &bottom_left[i];
            let t1 = &top_right[i];
            for j in 0..i {
                let b2 = &bottom_left[j];
                let t2 = &top_right[j];
                let width = t1[0].min(t2[0]) - b1[0].max(b2[0]); // 右上横坐标 - 左下横坐标
                let height = t1[1].min(t2[1]) - b1[1].max(b2[1]); // 右上纵坐标 - 左下纵坐标
                let side = width.min(height);
                max_side = max_side.max(side);
            }
        }
        (max_side as i64) * (max_side as i64)
    }
}
```

## 优化

外层循环枚举的矩形，如果其长或宽 $\le \textit{maxSide}$，那么 $\textit{maxSide}$ 不会变大，直接 `continue`。

```py [sol-Python3]
class Solution:
    def largestSquareArea(self, bottomLeft: List[List[int]], topRight: List[List[int]]) -> int:
        max_side = 0
        for i, ((bx, by), (tx, ty)) in enumerate(zip(bottomLeft, topRight)):
            if tx - bx <= max_side or ty - by <= max_side:
                continue  # 最优性剪枝：max_side 不可能变大
            for j in range(i):
                bx2, by2 = bottomLeft[j]
                tx2, ty2 = topRight[j]
                width = min(tx, tx2) - max(bx, bx2)  # 右上横坐标 - 左下横坐标
                height = min(ty, ty2) - max(by, by2)  # 右上纵坐标 - 左下纵坐标
                side = min(width, height)
                max_side = max(max_side, side)
        return max_side ** 2
```

```java [sol-Java]
class Solution {
    public long largestSquareArea(int[][] bottomLeft, int[][] topRight) {
        int maxSide = 0;
        for (int i = 0; i < bottomLeft.length; i++) {
            int[] b1 = bottomLeft[i];
            int[] t1 = topRight[i];
            if (t1[0] - b1[0] <= maxSide || t1[1] - b1[1] <= maxSide) {
                continue; // 最优性剪枝：maxSide 不可能变大
            }
            for (int j = 0; j < i; j++) {
                int[] b2 = bottomLeft[j];
                int[] t2 = topRight[j];
                int width = Math.min(t1[0], t2[0]) - Math.max(b1[0], b2[0]); // 右上横坐标 - 左下横坐标
                int height = Math.min(t1[1], t2[1]) - Math.max(b1[1], b2[1]); // 右上纵坐标 - 左下纵坐标
                int side = Math.min(width, height);
                maxSide = Math.max(maxSide, side);
            }
        }
        return (long) maxSide * maxSide;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long largestSquareArea(vector<vector<int>>& bottomLeft, vector<vector<int>>& topRight) {
        int max_side = 0;
        for (int i = 0; i < bottomLeft.size(); i++) {
            auto& b1 = bottomLeft[i];
            auto& t1 = topRight[i];
            if (t1[0] - b1[0] <= max_side || t1[1] - b1[1] <= max_side) {
                continue; // 最优性剪枝：max_side 不可能变大
            }
            for (int j = 0; j < i; j++) {
                auto& b2 = bottomLeft[j];
                auto& t2 = topRight[j];
                int width = min(t1[0], t2[0]) - max(b1[0], b2[0]); // 右上横坐标 - 左下横坐标
                int height = min(t1[1], t2[1]) - max(b1[1], b2[1]); // 右上纵坐标 - 左下纵坐标
                int side = min(width, height);
                max_side = max(max_side, side);
            }
        }
        return 1LL * max_side * max_side;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))
#define MAX(a, b) ((b) > (a) ? (b) : (a))

long long largestSquareArea(int** bottomLeft, int bottomLeftSize, int* bottomLeftColSize, int** topRight, int topRightSize, int* topRightColSize) {
    int max_side = 0;
    for (int i = 0; i < bottomLeftSize; i++) {
        int* b1 = bottomLeft[i];
        int* t1 = topRight[i];
        if (t1[0] - b1[0] <= max_side || t1[1] - b1[1] <= max_side) {
            continue; // 最优性剪枝：max_side 不可能变大
        }
        for (int j = 0; j < i; j++) {
            int* b2 = bottomLeft[j];
            int* t2 = topRight[j];
            int width = MIN(t1[0], t2[0]) - MAX(b1[0], b2[0]); // 右上横坐标 - 左下横坐标
            int height = MIN(t1[1], t2[1]) - MAX(b1[1], b2[1]); // 右上纵坐标 - 左下纵坐标
            int side = MIN(width, height);
            max_side = MAX(max_side, side);
        }
    }
    return 1LL * max_side * max_side;
}
```

```go [sol-Go]
func largestSquareArea(bottomLeft, topRight [][]int) int64 {
	maxSide := 0
	for i, b1 := range bottomLeft {
		t1 := topRight[i]
		if t1[0]-b1[0] <= maxSide || t1[1]-b1[1] <= maxSide {
			continue // 最优性剪枝：maxSide 不可能变大
		}
		for j, b2 := range bottomLeft[:i] {
			t2 := topRight[j]
			width := min(t1[0], t2[0]) - max(b1[0], b2[0])  // 右上横坐标 - 左下横坐标
			height := min(t1[1], t2[1]) - max(b1[1], b2[1]) // 右上纵坐标 - 左下纵坐标
			side := min(width, height)
			maxSide = max(maxSide, side)
		}
	}
	return int64(maxSide) * int64(maxSide)
}
```

```js [sol-JavaScript]
var largestSquareArea = function(bottomLeft, topRight) {
    let maxSide = 0;
    for (let i = 0; i < bottomLeft.length; i++) {
        const [bx, by] = bottomLeft[i];
        const [tx, ty] = topRight[i];
        if (tx - bx <= maxSide || ty - by <= maxSide) {
            continue; // 最优性剪枝：maxSide 不可能变大
        }
        for (let j = 0; j < i; j++) {
            const [bx2, by2] = bottomLeft[j];
            const [tx2, ty2] = topRight[j];
            const width = Math.min(tx, tx2) - Math.max(bx, bx2); // 右上横坐标 - 左下横坐标
            const height = Math.min(ty, ty2) - Math.max(by, by2); // 右上纵坐标 - 左下纵坐标
            const side = Math.min(width, height);
            maxSide = Math.max(maxSide, side);
        }
    }
    return maxSide * maxSide;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn largest_square_area(bottom_left: Vec<Vec<i32>>, top_right: Vec<Vec<i32>>) -> i64 {
        let mut max_side = 0;
        for i in 0..bottom_left.len() {
            let b1 = &bottom_left[i];
            let t1 = &top_right[i];
            if (t1[0] - b1[0] <= max_side || t1[1] - b1[1] <= max_side) {
                continue; // 最优性剪枝：max_side 不可能变大
            }
            for j in 0..i {
                let b2 = &bottom_left[j];
                let t2 = &top_right[j];
                let width = t1[0].min(t2[0]) - b1[0].max(b2[0]); // 右上横坐标 - 左下横坐标
                let height = t1[1].min(t2[1]) - b1[1].max(b2[1]); // 右上纵坐标 - 左下纵坐标
                let side = width.min(height);
                max_side = max_side.max(side);
            }
        }
        (max_side as i64) * (max_side as i64)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{bottomLeft}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
