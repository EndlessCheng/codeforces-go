按照相同字母分组，每一组累加代价。

贪心地，保留代价和最大的那一组，其余组全部删除。

也就是用总代价，减去最大代价和，即为答案。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minCost(self, s: str, cost: List[int]) -> int:
        sum_ = defaultdict(int)
        for ch, x in zip(s, cost):
            sum_[ch] += x
        return sum(sum_.values()) - max(sum_.values())
```

```java [sol-Java]
class Solution {
    public long minCost(String s, int[] cost) {
        long total = 0;
        long[] sum = new long[26];
        long maxSum = 0;
        for (int i = 0; i < cost.length; i++) {
            total += cost[i];
            int idx = s.charAt(i) - 'a';
            sum[idx] += cost[i];
            maxSum = Math.max(maxSum, sum[idx]);
        }
        return total - maxSum;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(string s, vector<int>& cost) {
        long long total = 0;
        long long sum[26]{};
        for (int i = 0; i < cost.size(); i++) {
            total += cost[i];
            sum[s[i] - 'a'] += cost[i];
        }
        return total - ranges::max(sum);
    }
};
```

```go [sol-Go]
func minCost(s string, cost []int) int64 {
	total := 0
	sum := [26]int{}
	for i, x := range cost {
		total += x
		sum[s[i]-'a'] += x
	}
	return int64(total - slices.Max(sum[:]))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
