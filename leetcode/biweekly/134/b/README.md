设 $\textit{enemyEnergies}$ 中的最小值为 $\textit{mn}$，元素和为 $s$。

如果 $\textit{currentEnergy} < \textit{mn}$，那么操作一无法执行，无法得到任何分数，所以也无法执行操作二，返回 $0$。

否则，操作顺序如下：

1. 对 $\textit{mn}$ 执行操作一，得到 $1$ 分。
2. 对除了 $\textit{mn}$ 以外的敌人执行操作二，得到 $s - \textit{mn}$ 的能量。
3. 对 $\textit{mn}$ 执行操作一，直到能量不足。

也可以理解为，先得到 $s - \textit{mn}$ 的能量，再不断对 $\textit{mn}$ 执行操作一，所以得分为 

$$
\left\lfloor\dfrac{\textit{currentEnergy} + s - \textit{mn}}{\textit{mn}}\right\rfloor
$$

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Yz421q7dD/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maximumPoints(self, enemyEnergies: List[int], currentEnergy: int) -> int:
        mn = min(enemyEnergies)
        if currentEnergy < mn:
            return 0
        return (currentEnergy + sum(enemyEnergies) - mn) // mn
```

```java [sol-Java]
class Solution {
    public long maximumPoints(int[] enemyEnergies, int currentEnergy) {
        int mn = Integer.MAX_VALUE;
        long s = 0;
        for (int e : enemyEnergies) {
            mn = Math.min(mn, e);
            s += e;
        }
        if (currentEnergy < mn) {
            return 0;
        }
        return (currentEnergy + s - mn) / mn;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumPoints(vector<int>& enemyEnergies, int currentEnergy) {
        int mn = ranges::min(enemyEnergies);
        if (currentEnergy < mn) {
            return 0;
        }
        long long s = reduce(enemyEnergies.begin(), enemyEnergies.end(), 0LL);
        return (currentEnergy + s - mn) / mn;
    }
};
```

```go [sol-Go]
func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	mn, s := math.MaxInt, 0
	for _, e := range enemyEnergies {
		mn = min(mn, e)
		s += e
	}
	if currentEnergy < mn {
		return 0
	}
	return int64((currentEnergy + s - mn) / mn)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{enemyEnergies}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
