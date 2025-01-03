定义 $f[i][j]$ 表示用前 $i$ 种题目恰好组成 $j$ 分的方案数。

对于第 $i$ 种题目，枚举做 $k$ 道题目，则子问题为「前 $i-1$ 种题目恰好组成 $j-k\cdot \textit{marks}_i$ 分的方案数」，因此有

$$
f[i][j] = \sum\limits_{k=0} f[i-1][j-k\cdot \textit{marks}_i]
$$

注意 $k$ 不能超过 $\textit{count}_i$，且 $j-k\cdot \textit{marks}_i\ge 0$。

代码实现时可以像 0-1 背包那样，压缩成一维，具体可以看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

> 注：滚动优化后，$k=0$ 就是 $f[j]$，无需计算。

- [本题视频讲解](https://www.bilibili.com/video/BV1SN411c7eD/)
- 更快的做法请看 [2902. 和带限制的子多重集合的数目](https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/solution/duo-zhong-bei-bao-fang-an-shu-cong-po-su-f5ay/)

```py [sol-Python3]
class Solution:
    def waysToReachTarget(self, target: int, types: List[List[int]]) -> int:
        MOD = 1_000_000_007
        f = [1] + [0] * target
        for count, marks in types:
            for j in range(target, 0, -1):
                for k in range(1, min(count, j // marks) + 1):
                    f[j] += f[j - k * marks]
                f[j] %= MOD
        return f[-1]
```

```java [sol-Java]
class Solution {
    public int waysToReachTarget(int target, int[][] types) {
        final int MOD = 1_000_000_007;
        int[] f = new int[target + 1];
        f[0] = 1;
        for (int[] p : types) {
            int count = p[0];
            int marks = p[1];
            for (int j = target; j > 0; j--) {
                for (int k = 1; k <= Math.min(count, j / marks); k++) {
                    f[j] = (f[j] + f[j - k * marks]) % MOD;
                }
            }
        }
        return f[target];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int waysToReachTarget(int target, vector<vector<int>>& types) {
        const int MOD = 1e9 + 7;
        vector<int> f(target + 1);
        f[0] = 1;
        for (auto& p : types) {
            int count = p[0], marks = p[1];
            for (int j = target; j > 0; j--) {
                for (int k = 1; k <= min(count, j / marks); k++) {
                    f[j] = (f[j] + f[j - k * marks]) % MOD;
                }
            }
        }
        return f[target];
    }
};
```

```go [sol-Go]
func waysToReachTarget(target int, types [][]int) int {
	const mod = 1_000_000_007
	f := make([]int, target+1)
	f[0] = 1
	for _, p := range types {
		count, marks := p[0], p[1]
		for j := target; j > 0; j-- {
			for k := 1; k <= min(count, j/marks); k++ {
				f[j] += f[j-k*marks]
			}
			f[j] %= mod
		}
	}
	return f[target]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{target}\cdot S)$，其中 $S$ 为所有 $\textit{count}_i$ 之和。
- 空间复杂度：$\mathcal{O}(\textit{target})$。

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
