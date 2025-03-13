## 题意

从 $[1,\infty]$ 中选 $n$ 个不同的整数，要求任意两数之和都不等于 $k$。

计算这 $n$ 个数的最小总和。

> 注意 $k$ 是可以选的，因为题目约束的是两数之和，不是一个数。

## 思路

#### 第一部分

对于 $[1,k-1]$ 内的数字：

- $1$ 和 $k-1$ 只能选其中一个，不能都选，不然两数之和为 $k$。既然只能选一个，那么选 $1$ 比选 $k-1$ 更好。注意不能都不选，这会导致后面要选更大的数，不是最优的。
- $2$ 和 $k-2$ 只能选其中一个，选 $2$，理由同上。
- $3$ 和 $k-3$ 只能选其中一个，选 $3$，理由同上。
- ……
- 一直到 $\left\lfloor\dfrac{k}{2}\right\rfloor$，无论 $k$ 是奇数还是偶数，它都可以选。

设 $m=\min\left(\left\lfloor\dfrac{k}{2}\right\rfloor, n\right)$，那么答案的第一部分是从 $1$ 到 $m$，根据等差数列求和公式，元素和为

$$
\dfrac{m(m+1)}{2}
$$

#### 第二部分

此时还剩下 $n-m$ 个数，只能在 $[k,\infty]$ 中选。这些数没有其他约束，选最小的 $n-m$ 个数即可。所以答案的第二部分是从 $k$ 到 $k+n-m-1$，根据等差数列求和公式，元素和为

$$
\dfrac{(2k+n-m-1)(n-m)}{2}
$$

综上所述，答案为

$$
\dfrac{m(m+1) + (2k+n-m-1)(n-m)}{2}
$$

[视频讲解](https://www.bilibili.com/video/BV1Rx4y1f75Y/) 第二题。

```py [sol-Python3]
class Solution:
    def minimumSum(self, n: int, k: int) -> int:
        m = min(k // 2, n)
        return (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) // 2
```

```java [sol-Java]
class Solution {
    public int minimumSum(int n, int k) {
        int m = Math.min(k / 2, n);
        return (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) / 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSum(int n, int k) {
        int m = min(k / 2, n);
        return (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) / 2;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int minimumSum(int n, int k) {
    int m = MIN(k / 2, n);
    return (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) / 2;
}
```

```go [sol-Go]
func minimumSum(n, k int) int {
    m := min(k/2, n)
    return (m*(m+1) + (k*2+n-m-1)*(n-m)) / 2
}
```

```js [sol-JavaScript]
var minimumSum = function(n, k) {
    const m = Math.min(Math.floor(k / 2), n);
    return (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) / 2;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimum_sum(n: i32, k: i32) -> i32 {
        let m = n.min(k / 2);
        (m * (m + 1) + (k * 2 + n - m - 1) * (n - m)) / 2
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
