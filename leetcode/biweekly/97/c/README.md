### 前置知识：滑动窗口

见 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

### 思路

我们可以强制让第二条线段的右端点恰好落在奖品上，设**第二条**线段右端点在 $\textit{prizePositions}[\textit{right}]$ 时，左端点最远覆盖了 $\textit{prizePositions}[\textit{left}]$，我们需要知道在 $\textit{prizePositions}[\textit{left}]$ 左侧的第一条线段最多可以覆盖多少个奖品。

那么，先想想只有一条线段要怎么做。

使用双指针，设线段右端点在 $\textit{prizePositions}[\textit{right}]$ 时，左端点最远覆盖了 $\textit{prizePositions}[\textit{left}]$，那么当前覆盖的奖品个数为 $\textit{right} - \textit{left} + 1$。

同时，用一个数组 $\textit{pre}[\textit{right}+1]$ 记录线段右端点**不超过** $\textit{prizePositions}[\textit{right}]$ 时最多可以覆盖多少个奖品。下标错开一位是为了方便下面计算。

初始 $\textit{pre}[0]=0$。根据 $\textit{pre}$ 的定义，有

$$
\textit{pre}[\textit{right}+1] = \max(\textit{pre}[\textit{right}],\textit{right} - \textit{left} + 1)
$$

回到第二条线段的计算，根据开头说的，此时最多可以覆盖的奖品数为

$$
\textit{right}-\textit{left}+1+\textit{pre}[\textit{left}]
$$

这里 $\textit{pre}[\textit{left}]$ 表示**第一条**线段右端点**不超过** $\textit{prizePositions}[\textit{left}-1]$ 时最多可以覆盖多少个奖品。

遍历过程中取上式的最大值，即为答案。

由于我们遍历了所有的奖品作为第二条线段的右端点，且通过 $\textit{pre}[\textit{left}]$ 保证第一条线段与第二条线段没有任何交点，且第一条线段覆盖了第二条线段左侧的最多奖品。那么这样遍历后，算出的答案就一定是所有情况中的最大值。

如果脑中没有一幅直观的图像，可以看看 [视频讲解【双周赛 97】](https://www.bilibili.com/video/BV1rM4y1X7z9/)的第三题。

**小优化**：如果 $2k+1\ge \textit{prizePositions}[n-1] - \textit{prizePositions}[0]$，说明所有奖品都可以获得，返回 $n$。

```py [sol-Python3]
class Solution:
    def maximizeWin(self, prizePositions: List[int], k: int) -> int:
        n = len(prizePositions)
        if k * 2 + 1 >= prizePositions[-1] - prizePositions[0]:
            return n
        pre = [0] * (n + 1)
        ans = left = 0
        for right, p in enumerate(prizePositions):
            while p - prizePositions[left] > k:
                left += 1
            ans = max(ans, right - left + 1 + pre[left])
            pre[right + 1] = max(pre[right], right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int maximizeWin(int[] prizePositions, int k) {
        int n = prizePositions.length;
        if (k * 2 + 1 >= prizePositions[n - 1] - prizePositions[0]) {
            return n;
        }
        int ans = 0;
        int left = 0;
        int[] pre = new int[n + 1];
        for (int right = 0; right < n; right++) {
            while (prizePositions[right] - prizePositions[left] > k) {
                left++;
            }
            ans = Math.max(ans, right - left + 1 + pre[left]);
            pre[right + 1] = Math.max(pre[right], right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximizeWin(vector<int>& prizePositions, int k) {
        int n = prizePositions.size();
        if (k * 2 + 1 >= prizePositions[n - 1] - prizePositions[0]) {
            return n;
        }
        int ans = 0, left = 0;
        vector<int> pre(n + 1);
        for (int right = 0; right < n; right++) {
            while (prizePositions[right] - prizePositions[left] > k) {
                left++;
            }
            ans = max(ans, right - left + 1 + pre[left]);
            pre[right + 1] = max(pre[right], right - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximizeWin(prizePositions []int, k int) (ans int) {
	n := len(prizePositions)
	if k*2+1 >= prizePositions[n-1]-prizePositions[0] {
		return n
	}
	pre := make([]int, n+1)
	left := 0
	for right, p := range prizePositions {
		for p-prizePositions[left] > k {
			left++
		}
		ans = max(ans, right-left+1+pre[left])
		pre[right+1] = max(pre[right], right-left+1)
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{prizePositions}$ 的长度。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

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
