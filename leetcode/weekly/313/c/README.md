[视频讲解](https://www.bilibili.com/video/BV1kd4y1q7fC) 第三题。

设 $n_1$ 为 $\textit{num}_1$ 的二进制表示的长度，$c_1$ 为 $\textit{num}_1$ 的置位数，$c_2$ 为 $\textit{num}_2$ 的置位数。

基本思路：

$x$ 的置位数和 $\textit{num}_2$ 相同，意味着 $x$ 的二进制表示中有 $c_2$ 个 $1$，我们需要合理地分配这 $c_2$ 个 $1$。

为了让异或和尽量小，这些 $1$ 应当从高位到低位匹配 $\textit{num}_1$ 中的 $1$；如果匹配完了还有多余的 $1$，那么就从低位到高位把 $0$ 改成 $1$。

分类讨论：

- 如果 $c_2\ge n_1$，$x$ 只能是 $2^{c_2}-1$，任何其他方案都会使异或和变大；
- 如果 $c_2=c_1$，那么 $x=\textit{num}_1$；
- 如果 $c_2<c_1$，那么将 $\textit{num}_1$ 的最低的 $c_1-c_2$ 个 $1$ 变成 $0$，其结果就是 $x$；
- 如果 $c_2>c_1$，那么将 $\textit{num}_1$ 的最低的 $c_2-c_1$ 个 $0$ 变成 $1$，其结果就是 $x$；

下面代码用到了一些位运算技巧，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol1-Python3]
class Solution:
    def minimizeXor(self, num1: int, num2: int) -> int:
        c1 = num1.bit_count()
        c2 = num2.bit_count()
        while c2 < c1:
            num1 &= num1 - 1  # 最低的 1 变成 0
            c2 += 1
        while c2 > c1:
            num1 |= num1 + 1  # 最低的 0 变成 1
            c2 -= 1
        return num1
```

```java [sol1-Java]
class Solution {
    public int minimizeXor(int num1, int num2) {
        var c1 = Integer.bitCount(num1);
        var c2 = Integer.bitCount(num2);
        for (; c2 < c1; ++c2) num1 &= num1 - 1; // 最低的 1 变成 0
        for (; c2 > c1; --c2) num1 |= num1 + 1; // 最低的 0 变成 1
        return num1;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minimizeXor(int num1, int num2) {
        int c1 = __builtin_popcount(num1);
        int c2 = __builtin_popcount(num2);
        for (; c2 < c1; ++c2) num1 &= num1 - 1; // 最低的 1 变成 0
        for (; c2 > c1; --c2) num1 |= num1 + 1; // 最低的 0 变成 1
        return num1;
    }
};
```

```go [sol1-Go]
func minimizeXor(num1, num2 int) int {
	c1 := bits.OnesCount(uint(num1))
	c2 := bits.OnesCount(uint(num2))
	for ; c2 < c1; c2++ {
		num1 &= num1 - 1 // 最低的 1 变成 0
	}
	for ; c2 > c1; c2-- {
		num1 |= num1 + 1 // 最低的 0 变成 1
	}
	return num1
}
```

#### 复杂度分析

- 时间复杂度：$O(|\log\textit{num}_1 - \log\textit{num}_2|)$。
- 空间复杂度：$O(1)$，仅用到若干变量。

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
