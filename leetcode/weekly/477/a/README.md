不断地把 $n$ 除以 $10$（下取整）直到 $0$，例如 $123\to 12\to 1\to 0$。在这个过程中的 $n\bmod 10$，即为每个数位。

对于不为零的数位 $d$，通过 $d\cdot 10^k$ 组合起来，其中 $k$ 是当前 $x$ 的十进制长度。$k$ 从 $0$ 开始，每添加一个数位就加一。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def sumAndMultiply(self, n: int) -> int:
        x, s, pow10 = 0, 0, 1
        while n:
            n, d = divmod(n, 10)
            if d > 0:
                x += d * pow10
                s += d
                pow10 *= 10
        return x * s
```

```java [sol-Java]
class Solution {
    public long sumAndMultiply(int n) {
        int x = 0;
        int sum = 0;
        for (int pow10 = 1; n > 0; n /= 10) {
            int d = n % 10;
            if (d > 0) {
                x += d * pow10;
                sum += d;
                pow10 *= 10;
            }
        }
        return (long) x * sum;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long sumAndMultiply(int n) {
        int x = 0, sum = 0, pow10 = 1;
        for (; n > 0; n /= 10) {
            int d = n % 10;
            if (d > 0) {
                x += d * pow10;
                sum += d;
                pow10 *= 10;
            }
        }
        return 1LL * x * sum;
    }
};
```

```go [sol-Go]
func sumAndMultiply(n int) int64 {
	x, sum, pow10 := 0, 0, 1
	for ; n > 0; n /= 10 {
		d := n % 10
		if d > 0 {
			x += d * pow10
			sum += d
			pow10 *= 10
		}
	}
	return int64(x) * int64(sum)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。$n$ 的十进制长度为 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

[3726. 移除十进制表示中的所有零](https://leetcode.cn/problems/remove-zeros-in-decimal-representation/)

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
