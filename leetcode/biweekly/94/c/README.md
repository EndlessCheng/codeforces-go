看到「最小化最大值」就要先尝试二分答案，这是因为：

- 最大元素越大，约束条件越宽松，越能够满足要求。
- 最大元素越小，约束条件越严苛，越不能满足要求。

下文把 $\textit{divisor}_1$ 和 $\textit{divisor}_2$ 简写成 $d_1$ 和 $d_2$，记 $\textit{LCM}$ 为 $d_1$ 和 $d_2$ 的最小公倍数。

由于：

- 能被 $d_2$ 整除但不能被 $d_1$ 整除的数，能在 $\textit{arr}_1$ 中且不能在 $\textit{arr}_2$ 中；
- 能被 $d_1$ 整除但不能被 $d_2$ 整除的数，能在 $\textit{arr}_2$ 中且不能在 $\textit{arr}_1$ 中；
- 既不能被 $d_1$ 整除也不能被 $d_2$ 整除的数，可以在 $\textit{arr}_1$ 和 $\textit{arr}_2$ 中。

因此二分答案 $x$，则：

- 有 $\left\lfloor\dfrac{x}{d_2}\right\rfloor - \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor$ 个数是 $\textit{arr}_1$ 独享的；
- 有 $\left\lfloor\dfrac{x}{d_1}\right\rfloor - \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor$ 个数是 $\textit{arr}_2$ 独享的；
- 有 $x - \left\lfloor\dfrac{x}{d_1}\right\rfloor - \left\lfloor\dfrac{x}{d_2}\right\rfloor + \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor$ 个数（根据容斥原理）是 $\textit{arr}_1$ 和 $\textit{arr}_2$ 共享的。

去掉独享的，剩余的数字只能在共享中选择，因此二分判定条件为

$$
x - \left\lfloor\dfrac{x}{d_1}\right\rfloor - \left\lfloor\dfrac{x}{d_2}\right\rfloor + \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor \ge \max\left(\textit{uniqueCnt}_1 - \left\lfloor\dfrac{x}{d_2}\right\rfloor + \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor, 0\right) + \max\left(\textit{uniqueCnt}_2 - \left\lfloor\dfrac{x}{d_1}\right\rfloor + \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor, 0\right)
$$

代码实现时，由于 $d_i$ 越小，不能在数组中的数越多，数组的最大元素越大，所以最坏情况下 $d_1=d_2=2$，只能取奇数。此时二分上界为 $(\textit{uniqueCnt}_1 + \textit{uniqueCnt}_2)\cdot 2-1$。

关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

[视频讲解](https://www.bilibili.com/video/BV1Dd4y1h72z/) 第三题。

```py [sol1-Python3]
class Solution:
    def minimizeSet(self, d1: int, d2: int, uniqueCnt1: int, uniqueCnt2: int) -> int:
        lcm = math.lcm(d1, d2)
        def check(x: int) -> bool:
            left1 = max(uniqueCnt1 - x // d2 + x // lcm, 0)
            left2 = max(uniqueCnt2 - x // d1 + x // lcm, 0)
            common = x - x // d1 - x // d2 + x // lcm
            return common >= left1 + left2
        return bisect_left(range((uniqueCnt1 + uniqueCnt2) * 2 - 1), True, key=check)
```

```go [sol1-Go]
func minimizeSet(d1, d2, uniqueCnt1, uniqueCnt2 int) int {
	lcm := d1 / gcd(d1, d2) * d2
	return sort.Search((uniqueCnt1+uniqueCnt2)*2-1, func(x int) bool {
		left1 := max(uniqueCnt1-x/d2+x/lcm, 0)
		left2 := max(uniqueCnt2-x/d1+x/lcm, 0)
		common := x - x/d1 - x/d2 + x/lcm
		return common >= left1+left2
	})
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$O(\log(\textit{divisor}_1+\textit{divisor}_2) + \log(\textit{uniqueCnt}_1+\textit{uniqueCnt}_2))$。
- 空间复杂度：$O(1)$。

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
