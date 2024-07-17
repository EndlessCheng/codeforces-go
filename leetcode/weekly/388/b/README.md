首先，应当选 $\textit{happiness}$ 中最大的 $k$ 的数。

这些数要按照什么顺序选呢？

由于小的数减成 $0$ 就不再减少了，优先选大的更好。

比如 $2,1,1$，如果按照 $1,1,2$ 的顺序选，答案为 $1+0+0=1$。但按照 $2,1,1$ 的顺序选，答案为 $2+0+0=2$ 更优。

请看 [视频讲解](https://www.bilibili.com/video/BV1Xr421J77b/) 第二题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumHappinessSum(self, happiness: List[int], k: int) -> int:
        happiness.sort(reverse=True)
        ans = 0
        for i, x in enumerate(happiness[:k]):
            if x <= i:
                break
            ans += x - i
        return ans
```

```java [sol-Java]
class Solution {
    public long maximumHappinessSum(int[] happiness, int k) {
        Arrays.sort(happiness);
        int n = happiness.length;
        long ans = 0;
        for (int i = n - 1; i >= n - k && happiness[i] > n - 1 - i; i--) {
            ans += happiness[i] - (n - 1 - i);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maximumHappinessSum(vector<int>& happiness, int k) {
        ranges::sort(happiness, greater());
        long long ans = 0;
        for (int i = 0; i < k && happiness[i] > i; i++) {
            ans += happiness[i] - i;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumHappinessSum(happiness []int, k int) (ans int64) {
	slices.SortFunc(happiness, func(a, b int) int { return b - a })
	for i, x := range happiness[:k] {
		if x <= i {
			break
		}
		ans += int64(x - i)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{happiness}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序和切片的开销。

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
