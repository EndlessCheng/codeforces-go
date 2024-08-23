请先阅读：[从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

从集合的视角看，$x$ 是每个 $\textit{nums}[i]$ 的**子集**。换句话说，$\textit{nums}[i]$ 一定是 $x$ 的**超集**。

例如 $x = 100100$，那么 $\textit{nums}[i]$ 一定在如下序列中：

$$
1\underline{00}1\underline{00},1\underline{00}1\underline{01},1\underline{00}1\underline{10},1\underline{00}1\underline{11},1\underline{01}1\underline{00},1\underline{01}1\underline{01},\cdots
$$

只看下划线上的数，是一个自然数序列

$$
0000,0001,0010,0011,0100,0101,\cdots
$$

为了让 $\textit{nums}[n-1]$ 尽量小，我们应当选择 $x$ 的超集中最小的 $n$ 个数。

所以把 $x$ 的二进制中的 $0$ 视作「空位」，把 $n-1$ 二进制的每个比特逐个填入空位，即为最小的 $\textit{nums}[n-1]$。

如果空位不足，往 $x$ 的前面添加前导零即可。

请看 [视频讲解](https://www.bilibili.com/video/BV1Pw4m1C79N/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minEnd(self, n: int, x: int) -> int:
        n -= 1  # 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
        i = j = 0
        while n >> j:
            # x 的第 i 个比特值是 0，即「空位」
            if (x >> i & 1) == 0:
                # 空位填入 n 的第 j 个比特值
                x |= (n >> j & 1) << i
                j += 1
            i += 1
        return x
```

```java [sol-Java]
class Solution {
    public long minEnd(int n, int x) {
        n--; // 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
        long ans = x;
        int i = 0, j = 0;
        while ((n >> j) > 0) {
            // x 的第 i 个比特值是 0，即「空位」
            if ((ans >> i & 1) == 0) {
                // 空位填入 n 的第 j 个比特值
                ans |= (long) (n >> j & 1) << i;
                j++;
            }
            i++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minEnd(int n, int x) {
        n--; // 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
        long long ans = x;
        int i = 0, j = 0;
        while (n >> j) {
            // x 的第 i 个比特值是 0，即「空位」
            if ((ans >> i & 1) == 0) {
                // 空位填入 n 的第 j 个比特值
                ans |= (long long) (n >> j & 1) << i;
                j++;
            }
            i++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minEnd(n, x int) int64 {
	n-- // 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
	i, j := 0, 0
	for n>>j > 0 {
		// x 的第 i 个比特值是 0，即「空位」
		if x>>i&1 == 0 {
			// 空位填入 n 的第 j 个比特值
			x |= n >> j & 1 << i
			j++
		}
		i++
	}
	return int64(x)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log x + \log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 优化

把 $x$ 取反，用 lowbit 枚举其中的 $1$，就是要填的空位。

```py [sol-Python3]
class Solution:
    def minEnd(self, n: int, x: int) -> int:
        n -= 1
        j = 0
        t = ~x
        while n >> j:
            lb = t & -t
            x |= (n >> j & 1) * lb
            j += 1
            t ^= lb
        return x
```

```java [sol-Java]
class Solution {
    public long minEnd(int n, int x) {
        n--;
        long ans = x;
        int j = 0;
        for (long t = ~x, lb; (n >> j) > 0; t ^= lb) {
            lb = t & -t;
            ans |= (long) (n >> j++ & 1) * lb;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minEnd(int n, int x) {
        n--;
        long long ans = x;
        int j = 0;
        for (long long t = ~x, lb; n >> j; t ^= lb) {
            lb = t & -t;
            ans |= (long long) (n >> j++ & 1) * lb;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minEnd(n, x int) int64 {
	n--
	j := 0
	for t, lb := ^x, 0; n>>j > 0; t ^= lb {
		lb = t & -t
		x |= n >> j & 1 * lb
		j++
	}
	return int64(x)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。循环次数只和入参 $n$ 有关。
- 空间复杂度：$\mathcal{O}(1)$。

## 更快的做法？

请看《Hacker's Delight》第 7.5 节。

## 思考题

额外输入一个 $\textit{forbidden}$ 数组，表示禁止出现在 $\textit{nums}$ 中的数。

在这种额外约束下，$\textit{nums}[n-1]$ 的最小值是多少？

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
