单独考虑一侧的房子，定义 $f[i]$ 表示前 $i$ 个地块的放置方案数，其中第 $i$ 个地块可以放房子，也可以不放房子。

考虑第 $i$ 个地块：

- 若不放房子，那么第 $i-1$ 个地块可放可不放，则有 $f[i] = f[i-1]$；
- 若放房子，那么第 $i-1$ 个地块无法放房子，第 $i-2$ 个地块可放可不放，则有 $f[i] = f[i-2]$。

因此

$$
f[i] = f[i-1] + f[i-2]
$$

边界为

- $f[0]=1$，空只有一种选择，就是不放房子。
- $f[1]=2$，放与不放两种方案。

由于两侧的房屋互相独立，根据乘法原理，答案为 $f[n]^2$。

[本题视频讲解](https://www.bilibili.com/video/BV1pW4y1r7xs)

```py [sol-Python3]
MOD = 1_000_000_007
MX = 10_001
f = [1, 2]
while len(f) < MX:
    f.append((f[-1] + f[-2]) % MOD)

class Solution:
    def countHousePlacements(self, n: int) -> int:
        return f[n] ** 2 % MOD
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;
    private static final int MX = 10_001;
    private static final int[] f = new int[MX];

    static {
        f[0] = 1;
        f[1] = 2;
        for (int i = 2; i < MX; i++) {
            f[i] = (f[i - 1] + f[i - 2]) % MOD;
        }
    }

    public int countHousePlacements(int n) {
        return (int) ((long) f[n] * f[n] % MOD);
    }
}
```

```cpp [sol-C++]
const int MOD = 1'000'000'007;
const int MX = 10'001;
int f[MX] = {1, 2};

int init = []() {
    for (int i = 2; i < MX; i++) {
        f[i] = (f[i - 1] + f[i - 2]) % MOD;
    }
    return 0;
}();

class Solution {
public:
    int countHousePlacements(int n) {
        return 1LL * f[n] * f[n] % MOD;
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

var f = [10_001]int{1, 2}

func init() {
	for i := 2; i < len(f); i++ {
		f[i] = (f[i-1] + f[i-2]) % mod
	}
}

func countHousePlacements(n int) int {
	return f[n] * f[n] % mod
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

添加一个约束：同一列不能都盖房子，要怎么做？

欢迎在评论区分享你的思路/代码。

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
