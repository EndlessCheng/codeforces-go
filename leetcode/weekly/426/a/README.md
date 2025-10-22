**题意**：返回 $\ge n$ 且二进制全为 $1$ 的最小整数。

**思路**：计算 $n$ 的二进制长度 $m$，那么答案的二进制长度至少是 $m$。由于长为 $m$ 的全为 $1$ 的二进制数 $\ge n$，满足要求，所以答案的二进制长度就是 $m$，所以答案为

$$
2^m - 1
$$

上式的意思是，把 $1$ 左移 $m$，得到 $1\underbrace{00\cdots 0}_{m\ 个\ 0}$，然后减一，得到 $\underbrace{11\cdots 1}_{m\ 个\ 1}$。

[本题视频讲解](https://www.bilibili.com/video/BV1tAzoY1EUN/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def smallestNumber(self, n: int) -> int:
        return (1 << n.bit_length()) - 1
```

```java [sol-Java 写法一]
class Solution {
    public int smallestNumber(int n) {
        int bitLength = 32 - Integer.numberOfLeadingZeros(n);
        return (1 << bitLength) - 1;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public int smallestNumber(int n) {
        return (Integer.highestOneBit(n) << 1) - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int smallestNumber(int n) {
        return (1 << bit_width((uint32_t) n)) - 1;
    }
};
```

```c [sol-C]
int smallestNumber(int n) {
    int bit_length = 32 - __builtin_clz(n);
    return (1 << bit_length) - 1;
}
```

```go [sol-Go]
func smallestNumber(n int) int {
	return 1<<bits.Len(uint(n)) - 1
}
```

```js [sol-JavaScript]
var smallestNumber = function(n) {
    const bitLength = 32 - Math.clz32(n);
    return (1 << bitLength) - 1;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn smallest_number(n: i32) -> i32 {
        let bit_length = 32 - n.leading_zeros();
        (1 << bit_length) - 1
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

把 $\ge n$ 改成 $>n$ 怎么做？

答案见评论。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
