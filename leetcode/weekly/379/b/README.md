分类讨论：

- 如果车能直接攻击到皇后，答案是 $1$。
- 如果象能直接攻击到皇后，答案是 $1$。
- 其他情况，答案一定是 $2$：
    - 如果车能攻击到皇后，但被象挡住，那么移走象，车就可以攻击到皇后，答案是 $2$。小知识：这在国际象棋中叫做「闪击」。
    - 如果象能攻击到皇后，但被车挡住，那么移走车，象就可以攻击到皇后，答案是 $2$。
    - 如果车不能攻击到皇后，那么车可以水平移动或者垂直移动，其中一种方式必定不会被象挡住，可以攻击到皇后，答案是 $2$。

判断能否直接攻击到：

- 对于车，如果和皇后在同一行或者同一列，且中间没有象，那么就可以直接攻击到皇后。
- 对于象，如果和皇后在同一斜线，且中间没有车，那么就可以直接攻击到皇后。判断的技巧我在[【基础算法精讲 16】](https://www.bilibili.com/video/BV1mY411D7f6/) 中有讲到，欢迎收看。

```py [sol-Python3]
class Solution:
    def minMovesToCaptureTheQueen(self, a: int, b: int, c: int, d: int, e: int, f: int) -> int:
        # m 在 l 和 r 之间（写不写等号都可以）
        def in_between(l: int, m: int, r: int) -> bool:
            return min(l, r) < m < max(l, r)

        # 车直接攻击到皇后 or 象直接攻击到皇后
        if a == e and (c != e or not in_between(b, d, f)) or \
           b == f and (d != f or not in_between(a, c, e)) or \
           c + d == e + f and (a + b != e + f or not in_between(c, a, e)) or \
           c - d == e - f and (a - b != e - f or not in_between(c, a, e)):
            return 1
        return 2
```

```java [sol-Java]
class Solution {
    public int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
        if (a == e && (c != e || !inBetween(b, d, f)) || // 车直接攻击到皇后（同一行）
            b == f && (d != f || !inBetween(a, c, e)) || // 车直接攻击到皇后（同一列）
            c + d == e + f && (a + b != e + f || !inBetween(c, a, e)) || // 象直接攻击到皇后
            c - d == e - f && (a - b != e - f || !inBetween(c, a, e))) {
            return 1;
        }
        return 2;
    }

    // m 在 l 和 r 之间（写不写等号都可以）
    private boolean inBetween(int l, int m, int r) {
        return Math.min(l, r) < m && m < Math.max(l, r);
    }
}
```

```cpp [sol-C++]
class Solution {
    // m 在 l 和 r 之间（写不写等号都可以）
    bool in_between(int l, int m, int r) {
        return min(l, r) < m && m < max(l, r);
    }

public:
    int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
        if (a == e && (c != e || !in_between(b, d, f)) || // 车直接攻击到皇后（同一行）
            b == f && (d != f || !in_between(a, c, e)) || // 车直接攻击到皇后（同一列）
            c + d == e + f && (a + b != e + f || !in_between(c, a, e)) || // 象直接攻击到皇后
            c - d == e - f && (a - b != e - f || !in_between(c, a, e))) {
            return 1;
        }
        return 2;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))
#define MAX(a, b) ((b) > (a) ? (b) : (a))

// m 在 l 和 r 之间（写不写等号都可以）
int inBetween(int l, int m, int r) {
    return MIN(l, r) < m && m < MAX(l, r);
}

int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
    if (a == e && (c != e || !inBetween(b, d, f)) || // 车直接攻击到皇后（同一行）
        b == f && (d != f || !inBetween(a, c, e)) || // 车直接攻击到皇后（同一列）
        c + d == e + f && (a + b != e + f || !inBetween(c, a, e)) || // 象直接攻击到皇后
        c - d == e - f && (a - b != e - f || !inBetween(c, a, e))) {
        return 1;
    }
    return 2;
}
```

```go [sol-Go]
// m 在 l 和 r 之间（写不写等号都可以）
func inBetween(l, m, r int) bool {
    return min(l, r) < m && m < max(l, r)
}

func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
    if a == e && (c != e || !inBetween(b, d, f)) || // 车直接攻击到皇后（同一行）
        b == f && (d != f || !inBetween(a, c, e)) || // 车直接攻击到皇后（同一列）
        c+d == e+f && (a+b != e+f || !inBetween(c, a, e)) || // 象直接攻击到皇后
        c-d == e-f && (a-b != e-f || !inBetween(c, a, e)) {
        return 1
    }
    return 2
}
```

```js [sol-JavaScript]
// m 在 l 和 r 之间（写不写等号都可以）
function inBetween(l, m, r) {
    return Math.min(l, r) < m && m < Math.max(l, r);
}

var minMovesToCaptureTheQueen = function(a, b, c, d, e, f) {
    if (a === e && (c !== e || !inBetween(b, d, f)) || // 车直接攻击到皇后（同一行）
        b === f && (d !== f || !inBetween(a, c, e)) || // 车直接攻击到皇后（同一列）
        c + d === e + f && (a + b !== e + f || !inBetween(c, a, e)) || // 象直接攻击到皇后
        c - d === e - f && (a - b !== e - f || !inBetween(c, a, e))) {
        return 1;
    }
    return 2;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_moves_to_capture_the_queen(a: i32, b: i32, c: i32, d: i32, e: i32, f: i32) -> i32 {
        // m 在 l 和 r 之间（写不写等号都可以）
        let in_between = |l: i32, m: i32, r: i32| l.min(r) < m && m < l.max(r); 

        if a == e && (c != e || !in_between(b, d, f)) || // 车直接攻击到皇后（同一行）
            b == f && (d != f || !in_between(a, c, e)) || // 车直接攻击到皇后（同一列）
            c + d == e + f && (a + b != e + f || !in_between(c, a, e)) || // 象直接攻击到皇后
            c - d == e - f && (a - b != e - f || !in_between(c, a, e)) {
            return 1;
        }
        2
    }
}
```

另一种写法，如果 $(m-l)\cdot(m-r) > 0$，那么整数 $m$ 不在整数 $l$ 和 $r$ 之间，

```py [sol-Python3]
class Solution:
    def minMovesToCaptureTheQueen(self, a: int, b: int, c: int, d: int, e: int, f: int) -> int:
        # 车直接攻击到皇后 or 象直接攻击到皇后
        if a == e and (c != e or (d - b) * (d - f) > 0) or \
           b == f and (d != f or (c - a) * (c - e) > 0) or \
           c + d == e + f and (a + b != e + f or (a - c) * (a - e) > 0) or \
           c - d == e - f and (a - b != e - f or (a - c) * (a - e) > 0):
            return 1
        return 2
```

```java [sol-Java]
class Solution {
    public int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
        if (a == e && (c != e || (d - b) * (d - f) > 0) || // 车直接攻击到皇后（同一行）
            b == f && (d != f || (c - a) * (c - e) > 0) || // 车直接攻击到皇后（同一列）
            c + d == e + f && (a + b != e + f || (a - c) * (a - e) > 0) || // 象直接攻击到皇后
            c - d == e - f && (a - b != e - f || (a - c) * (a - e) > 0)) {
            return 1;
        }
        return 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
        if (a == e && (c != e || (d - b) * (d - f) > 0) || // 车直接攻击到皇后（同一行）
            b == f && (d != f || (c - a) * (c - e) > 0) || // 车直接攻击到皇后（同一列）
            c + d == e + f && (a + b != e + f || (a - c) * (a - e) > 0) || // 象直接攻击到皇后
            c - d == e - f && (a - b != e - f || (a - c) * (a - e) > 0)) {
            return 1;
        }
        return 2;
    }
};
```

```c [sol-C]
int minMovesToCaptureTheQueen(int a, int b, int c, int d, int e, int f) {
    if (a == e && (c != e || (d - b) * (d - f) > 0) || // 车直接攻击到皇后（同一行）
        b == f && (d != f || (c - a) * (c - e) > 0) || // 车直接攻击到皇后（同一列）
        c + d == e + f && (a + b != e + f || (a - c) * (a - e) > 0) || // 象直接攻击到皇后
        c - d == e - f && (a - b != e - f || (a - c) * (a - e) > 0)) {
        return 1;
    }
    return 2;
}
```

```go [sol-Go]
func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
    if a == e && (c != e || (d-b)*(d-f) > 0) || // 车直接攻击到皇后（同一行）
        b == f && (d != f || (c-a)*(c-e) > 0) || // 车直接攻击到皇后（同一列）
        c+d == e+f && (a+b != e+f || (a-c)*(a-e) > 0) || // 象直接攻击到皇后
        c-d == e-f && (a-b != e-f || (a-c)*(a-e) > 0) {
        return 1
    }
    return 2
}
```

```js [sol-JavaScript]
var minMovesToCaptureTheQueen = function(a, b, c, d, e, f) {
    if (a === e && (c !== e || (d - b) * (d - f) > 0) || // 车直接攻击到皇后（同一行）
        b === f && (d !== f || (c - a) * (c - e) > 0) || // 车直接攻击到皇后（同一列）
        c + d === e + f && (a + b !== e + f || (a - c) * (a - e) > 0) || // 象直接攻击到皇后
        c - d === e - f && (a - b !== e - f || (a - c) * (a - e) > 0)) {
        return 1;
    }
    return 2;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_moves_to_capture_the_queen(a: i32, b: i32, c: i32, d: i32, e: i32, f: i32) -> i32 {
        if a == e && (c != e || (d - b) * (d - f) > 0) || // 车直接攻击到皇后（同一行）
            b == f && (d != f || (c - a) * (c - e) > 0) || // 车直接攻击到皇后（同一列）
            c + d == e + f && (a + b != e + f || (a - c) * (a - e) > 0) || // 象直接攻击到皇后
            c - d == e - f && (a - b != e - f || (a - c) * (a - e) > 0) {
            return 1;
        }
        2
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果要输出具体移动方案呢？

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
