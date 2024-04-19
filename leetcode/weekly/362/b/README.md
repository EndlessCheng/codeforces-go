由于可以往 $8$ 个方向走，那么最快可以在

$$
\max\{|sx-fx|, |sy-fy|\}
$$

秒后到达终点（先斜着走再直走）。

上式只要小于等于 $t$ 就能恰好到达终点。因为我们可以在到达终点附近时，在终点周围不断「绕圈」消耗时间，这样可以直到最后一秒才走到终点。

**特殊情况**：如果起点和终点重合，那么 $t=1$ 的情况是无法回到起点的；如果 $t\ne 1$，我们可以同样地在起点不断「绕圈」，最后回到起点。

请看 [视频讲解](https://www.bilibili.com/video/BV1U34y1N7Pe/) 第二题。

```py [sol-Python3]
class Solution:
    def isReachableAtTime(self, sx: int, sy: int, fx: int, fy: int, t: int) -> bool:
        if sx == fx and sy == fy:
            return t != 1
        return max(abs(sx - fx), abs(sy - fy)) <= t
```

```java [sol-Java]
public class Solution {
    public boolean isReachableAtTime(int sx, int sy, int fx, int fy, int t) {
        if (sx == fx && sy == fy) {
            return t != 1;
        }
        return Math.max(Math.abs(sx - fx), Math.abs(sy - fy)) <= t;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isReachableAtTime(int sx, int sy, int fx, int fy, int t) {
        if (sx == fx && sy == fy)
            return t != 1;
        return max(abs(sx - fx), abs(sy - fy)) <= t;
    }
};
```

```go [sol-Go]
func isReachableAtTime(sx, sy, fx, fy, t int) bool {
	if sx == fx && sy == fy {
		return t != 1
	}
	return max(abs(sx-fx), abs(sy-fy)) <= t
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```js [sol-JavaScript]
var isReachableAtTime = function(sx, sy, fx, fy, t) {
    if (sx === fx && sy === fy) {
        return t !== 1;
    }
    return Math.max(Math.abs(sx - fx), Math.abs(sy - fy)) <= t;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果改成只能往 $4$ 个方向走呢？

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
