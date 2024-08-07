根据题意，修改成 $\textit{nums}$ 中的数字，可以让最小得分为 $0$。那么分数就等于最大得分。

那么从小到大排序后，我们可以：

- 修改最大的两个数为 $\textit{nums}[n-3]$，最大得分为 $\textit{nums}[n-3]-\textit{nums}[0]$；
- 修改最小的为 $\textit{nums}[1]$，最大的为 $\textit{nums}[n-2]$，最大得分为 $\textit{nums}[n-2]-\textit{nums}[1]$；
- 修改最小的两个数为 $\textit{nums}[2]$，最大得分为 $\textit{nums}[n-1]-\textit{nums}[2]$。

这样修改的理由是，修改成再更大/更小的数，不会影响最大得分了。

附：[视频讲解](https://www.bilibili.com/video/BV15D4y1G7ms/)

```py [sol-Python3]
class Solution:
    def minimizeSum(self, a: List[int]) -> int:
        a.sort()
        return min(a[-3] - a[0], a[-2] - a[1], a[-1] - a[2])
```

```java [sol-Java]
class Solution {
    public int minimizeSum(int[] a) {
        Arrays.sort(a);
        int n = a.length;
        return Math.min(Math.min(a[n - 3] - a[0], a[n - 2] - a[1]), a[n - 1] - a[2]);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizeSum(vector<int>& a) {
        ranges::sort(a);
        int n = a.size();
        return min({a[n - 3] - a[0], a[n - 2] - a[1], a[n - 1] - a[2]});
    }
};
```

```go [sol-Go]
func minimizeSum(a []int) int {
	slices.Sort(a)
	n := len(a)
	return min(a[n-3]-a[0], a[n-2]-a[1], a[n-1]-a[2])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。手动维护或者用快速选择可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序时栈的开销，仅用到若干额外变量。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
