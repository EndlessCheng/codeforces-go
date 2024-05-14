[本题视频讲解](https://www.bilibili.com/video/BV1ae411e7fn/)

分类讨论：

- 如果车能直接攻击到皇后，答案是 $1$。
- 如果象能直接攻击到皇后，答案是 $1$。
- 如果车被象挡住，那么移走象，车就可以攻击到皇后，答案是 $2$。小知识：这在国际象棋中称作「闪击」。
- 如果象被车挡住，那么移走车，象就可以攻击到皇后，答案是 $2$。
- 如果车不能直接攻击到皇后，那么车可以水平移动或者垂直移动，其中一种方式必定不会被象挡住，可以攻击到皇后，答案是 $2$。

对于车，如果和皇后在同一水平线或者同一竖直线，且中间没有象，那么就可以直接攻击到皇后。

对于象，如果和皇后在同一斜线，且中间没有车，那么就可以直接攻击到皇后。判断的技巧我在[【基础算法精讲 16】](https://www.bilibili.com/video/BV1mY411D7f6/) 中有讲到，欢迎收看。

```py [sol-Python3]
class Solution:
    def minMovesToCaptureTheQueen(self, a: int, b: int, c: int, d: int, e: int, f: int) -> int:
        def ok(l: int, m: int, r: int) -> bool:
            return not min(l, r) < m < max(l, r)

        if a == e and (c != e or ok(b, d, f)) or \
           b == f and (d != f or ok(a, c, e)) or \
           c + d == e + f and (a + b != e + f or ok(c, a, e)) or \
           c - d == e - f and (a - b != e - f or ok(c, a, e)):
            return 1
        return 2
```

```java [sol-Java]
public class Solution {
    public int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
        if (a == e && (c != e || ok(b, d, f)) ||
            b == f && (d != f || ok(a, c, e)) ||
            c + d == e + f && (a + b != e + f || ok(c, a, e)) ||
            c - d == e - f && (a - b != e - f || ok(c, a, e))) {
            return 1;
        }
        return 2;
    }

    private boolean ok(int l, int m, int r) {
        return m < Math.min(l, r) || m > Math.max(l, r);
    }
}
```

```cpp [sol-C++]
class Solution {
    bool ok(int l, int m, int r) {
        return m < min(l, r) || m > max(l, r);
    }

public:
    int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
        if (a == e && (c != e || ok(b, d, f)) ||
            b == f && (d != f || ok(a, c, e)) ||
            c + d == e + f && (a + b != e + f || ok(c, a, e)) ||
            c - d == e - f && (a - b != e - f || ok(c, a, e))) {
            return 1;
        }
        return 2;
    }
};
```

```go [sol-Go]
func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	if a == e && (c != e || ok(b, d, f)) ||
		b == f && (d != f || ok(a, c, e)) ||
		c+d == e+f && (a+b != e+f || ok(c, a, e)) ||
		c-d == e-f && (a-b != e-f || ok(c, a, e)) {
		return 1
	}
	return 2
}

func ok(l, m, r int) bool {
	return m < min(l, r) || m > max(l, r)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果要输出移动方案呢？

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
