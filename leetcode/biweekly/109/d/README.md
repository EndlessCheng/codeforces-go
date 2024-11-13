把 $n$ 看成背包容量，$1^x,2^x,3^x,\cdots$ 看成物品，本题就是一个 **0-1 背包**模板题，具体请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

代码实现时，由于 $n=300,x=1$ 算出来的结果不超过 $64$ 位整数的最大值，可以在计算结束后再取模。

```py [sol-Python3]
class Solution:
    def numberOfWays(self, n: int, x: int) -> int:
        f = [1] + [0] * n
        for i in range(1, n + 1):
            v = i ** x
            if v > n:
                break
            for s in range(n, v - 1, -1):
                f[s] += f[s - v]
        return f[n] % 1_000_000_007
```

```py [sol-Python3 预处理 0ms]
# 预处理所有答案
MX_N, MX_X = 300, 5
f = [[1] + [0] * MX_N for _ in range(MX_X + 1)]
for x in range(1, MX_X + 1):
    for i in count(1):
        v = i ** x
        if v > MX_N:
            break
        for s in range(MX_N, v - 1, -1):
            f[x][s] += f[x][s - v]

class Solution:
    def numberOfWays(self, n: int, x: int) -> int:
        return f[x][n] % 1_000_000_007
```

```java [sol-Java]
class Solution {
    int numberOfWays(int n, int x) {
        long[] f = new long[n + 1];
        f[0] = 1;
        // 本题数据范围小，Math.pow 的计算结果一定准确
        for (int i = 1; Math.pow(i, x) <= n; i++) {
            int v = (int) Math.pow(i, x);
            for (int s = n; s >= v; s--) {
                f[s] += f[s - v];
            }
        }
        return (int) (f[n] % 1_000_000_007);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfWays(int n, int x) {
        vector<long long> f(n + 1);
        f[0] = 1;
        // 本题数据范围小，pow 的计算结果一定准确
        for (int i = 1; pow(i, x) <= n; i++) {
            int v = pow(i, x);
            for (int s = n; s >= v; s--) {
                f[s] += f[s - v];
            }
        }
        return f[n] % 1'000'000'007;
    }
};
```

```go [sol-Go]
func numberOfWays(n, x int) int {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; pow(i, x) <= n; i++ {
		v := pow(i, x)
		for s := n; s >= v; s-- {
			f[s] += f[s-v]
		}
	}
	return f[n] % 1_000_000_007
}

// 本题数据范围小，math.Pow 的计算结果一定准确
func pow(i, x int) int {
	return int(math.Pow(float64(i), float64(x)))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt[x]n)$。
- 空间复杂度：$\mathcal{O}(n)$。

注：也可以使用快速幂计算 $\texttt{pow}$，原理见[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
