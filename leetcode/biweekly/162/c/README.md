假设先玩陆地游乐设施。

贪心地，完成陆地游乐设施的时间越早越好，最早完成时间为

$$
\textit{minFinish} = \min_{i=0}^{n-1} \textit{landStartTime}[i] + \textit{landDuration}[i]
$$

对于水上游乐设施，每个设施的最早开始时间为

$$
\max(\textit{waterStartTime}[i],\textit{minFinish})
$$

所以完成水上游乐设施的最早时间为

$$
\min_{i=0}^{m-1} \max(\textit{waterStartTime}[i],\textit{minFinish}) + \textit{waterDuration}[i]
$$

对于先玩水上游乐设施的情况，计算方式同上。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def solve(self, landStartTime: List[int], landDuration: List[int], waterStartTime: List[int], waterDuration: List[int]) -> int:
        min_finish = min(start + duration for start, duration in zip(landStartTime, landDuration))
        return min(max(start, min_finish) + duration for start, duration in zip(waterStartTime, waterDuration))

    def earliestFinishTime(self, landStartTime: List[int], landDuration: List[int], waterStartTime: List[int], waterDuration: List[int]) -> int:
        land_water = self.solve(landStartTime, landDuration, waterStartTime, waterDuration)
        water_land = self.solve(waterStartTime, waterDuration, landStartTime, landDuration)
        return min(land_water, water_land)
```

```java [sol-Java]
class Solution {
    public int earliestFinishTime(int[] landStartTime, int[] landDuration, int[] waterStartTime, int[] waterDuration) {
        int landWater = solve(landStartTime, landDuration, waterStartTime, waterDuration);
        int waterLand = solve(waterStartTime, waterDuration, landStartTime, landDuration);
        return Math.min(landWater, waterLand);
    }

    private int solve(int[] landStartTime, int[] landDuration, int[] waterStartTime, int[] waterDuration) {
        int minFinish = Integer.MAX_VALUE;
        for (int i = 0; i < landStartTime.length; i++) {
            minFinish = Math.min(minFinish, landStartTime[i] + landDuration[i]);
        }

        int res = Integer.MAX_VALUE;
        for (int i = 0; i < waterStartTime.length; i++) {
            res = Math.min(res, Math.max(waterStartTime[i], minFinish) + waterDuration[i]);
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    int solve(vector<int>& landStartTime, vector<int>& landDuration, vector<int>& waterStartTime, vector<int>& waterDuration) {
        int min_finish = INT_MAX;
        for (int i = 0; i < landStartTime.size(); i++) {
            min_finish = min(min_finish, landStartTime[i] + landDuration[i]);
        }

        int res = INT_MAX;
        for (int i = 0; i < waterStartTime.size(); i++) {
            res = min(res, max(waterStartTime[i], min_finish) + waterDuration[i]);
        }
        return res;
    }

public:
    int earliestFinishTime(vector<int>& landStartTime, vector<int>& landDuration, vector<int>& waterStartTime, vector<int>& waterDuration) {
        int land_water = solve(landStartTime, landDuration, waterStartTime, waterDuration);
        int water_land = solve(waterStartTime, waterDuration, landStartTime, landDuration);
        return min(land_water, water_land);
    }
};
```

```go [sol-Go]
func solve(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	minFinish := math.MaxInt
	for i, start := range landStartTime {
		minFinish = min(minFinish, start+landDuration[i])
	}

	res := math.MaxInt
	for i, start := range waterStartTime {
		res = min(res, max(start, minFinish)+waterDuration[i])
	}
	return res
}

func earliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration []int) int {
	landWater := solve(landStartTime, landDuration, waterStartTime, waterDuration)
	waterLand := solve(waterStartTime, waterDuration, landStartTime, landDuration)
	return min(landWater, waterLand)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $\textit{landStartTime}$ 的长度，$m$ 是 $\textit{waterStartTime}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
