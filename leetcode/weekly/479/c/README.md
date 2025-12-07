**提示**：如果问题没有让我们分别计算每个 $\text{score}(j)$，而是计算 $\text{score}(j)$ 的总和，通常可以用**贡献法**解决。

横看成岭侧成峰，考虑每个房间对总得分的贡献：

- 有多少个起点 $j$，可以在房间 $i$ 得到 $1$ 分？ 

设 $\textit{damage}$ 的**前缀和**数组为 $s$。关于 $s$ 数组的定义，请看 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

从起点 $j\ (j\le i)$ 一步步走到房间 $i$，一共扣除了 $s[i+1] - s[j]$ 的血量。所以当前剩余血量为 $\textit{hp} - (s[i+1] - s[j])$。

如果剩余血量至少为 $\textit{requirement}[i]$，即

$$
\textit{hp} - (s[i+1] - s[j]) \ge \textit{requirement}[i]
$$

那么从 $j$ 出发，可以在房间 $i$ 得到 $1$ 分。

移项，得

$$
s[j] \ge s[i+1] + \textit{requirement}[i] - \textit{hp}
$$

由于题目保证 $\textit{damage}[i]$ 非负，所以 $s$ 是递增数组。我们可以在 $s$ 的 $[0,i]$ 中二分查找第一个 $\ge s[i+1] + \textit{requirement}[i] - \textit{hp}$ 的元素下标 $j$（如果不存在则 $j=i+1$），那么 $[j,i]$ 中的整数都可以作为起点，在房间 $i$ 得到 $1$ 分。

> 注：如果 $\textit{damage}[i] < 0$，可以用有序集合或者值域树状数组，同样可以快速计算大于等于一个数的元素个数。

所以房间 $i$ 对总得分的贡献为 

$$
i-j+1
$$

累加即为答案。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def totalScore(self, hp: int, damage: List[int], requirement: List[int]) -> int:
        s = [0] * (len(damage) + 1)
        ans = 0
        for i, (dmg, req) in enumerate(zip(damage, requirement)):
            s[i + 1] = s[i] + dmg
            low = s[i + 1] + req - hp
            j = bisect_left(s, low, 0, i + 1)  # 在 [0, i] 中二分
            ans += i - j + 1
        return ans
```

```java [sol-Java]
class Solution {
    public long totalScore(int hp, int[] damage, int[] requirement) {
        int n = damage.length;
        int[] sum = new int[n + 1];
        long ans = 0;
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + damage[i];
            int low = sum[i + 1] + requirement[i] - hp;
            // 本题 sum 是严格递增的，没有重复元素，可以用 Arrays.binarySearch
            int j = Arrays.binarySearch(sum, 0, i + 1, low); // 在 [0, i] 中二分
            if (j < 0) j = ~j;
            ans += i - j + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long totalScore(int hp, vector<int>& damage, vector<int>& requirement) {
        int n = damage.size();
        vector<int> sum(n + 1);
        long long ans = 0;
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + damage[i];
            int low = sum[i + 1] + requirement[i] - hp;
            int j = lower_bound(sum.begin(), sum.begin() + i + 1, low) - sum.begin();
            ans += i - j + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func totalScore(hp int, damage []int, requirement []int) (ans int64) {
	sum := make([]int, len(damage)+1)
	for i, req := range requirement {
		sum[i+1] = sum[i] + damage[i]
		low := sum[i+1] + req - hp
		j := sort.SearchInts(sum[:i+1], low)
		ans += int64(i - j + 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{damage}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

力扣上的贡献法题目主要集中在单调栈中，见单调栈题单的「**三、贡献法**」。

[979. 在二叉树中分配硬币](https://leetcode.cn/problems/distribute-coins-in-binary-tree/) 这题也是贡献法。

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
