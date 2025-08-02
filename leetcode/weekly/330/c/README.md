### 提示 1

问题相当于把 $\textit{weights}$ 划分成 $k$ 个连续子数组，分数等于每个子数组的两端的值之和。

### 提示 2

$\textit{weights}[0]$ 和 $\textit{weights}[n-1]$ 一定在分数中，最大分数和最小分数相减，抵消了。

上一个子数组的末尾和下一个子数组的开头一定**同时**在分数中。

### 提示 3

把所有 $n-1$ 个 $\textit{weights}[i]+\textit{weights}[i+1]$ 算出来，排序，那么最大的 $k-1$ 个数和最小的 $k-1$ 个数相减，即为答案。

附：[视频讲解](https://www.bilibili.com/video/BV1mD4y1E7QK/)

```py [sol-Python3]
class Solution:
    def putMarbles(self, weights: List[int], k: int) -> int:
        for i in range(len(weights) - 1):
            weights[i] += weights[i + 1]
        weights.pop()
        weights.sort()
        return sum(weights[len(weights) - k + 1:]) - sum(weights[:k - 1])
```

```java [sol-Java]
class Solution {
    public long putMarbles(int[] weights, int k) {
        int n = weights.length;
        for (int i = 0; i < n - 1; i++) {
            weights[i] += weights[i + 1];
        }
        Arrays.sort(weights, 0, n - 1); // 去掉最后一个数

        long ans = 0;
        for (int i = 0; i < k - 1; i++) {
            ans += weights[n - 2 - i] - weights[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long putMarbles(vector<int>& weights, int k) {
        int n = weights.size();
        for (int i = 0; i < n - 1; i++) {
            weights[i] += weights[i + 1];
        }
        sort(weights.begin(), weights.end() - 1); // 去掉最后一个数

        long long ans = 0;
        for (int i = 0; i < k - 1; i++) {
            ans += weights[n - 2 - i] - weights[i];
        }
        return ans;
    }
};
```

```cpp [sol-C++ 快速选择]
class Solution {
public:
    long long putMarbles(vector<int>& weights, int k) {
        k--; // 注意这里减一了
        if (k == 0) {
            return 0;
        }
        int n = weights.size() - 1;
        for (int i = 0; i < n; i++) {
            weights[i] += weights[i + 1];
        }
        weights.pop_back();

        long ans = 0;
        ranges::nth_element(weights, weights.begin() + k);
        for (int i = 0; i < k; i++) {
            ans -= weights[i];
        }
        ranges::nth_element(weights, weights.end() - k);
        for (int i = 0; i < k; i++) {
            ans += weights[n - 1 - i];
        }
        return ans;
    }
};
```

```go [sol-Go]
func putMarbles(weights []int, k int) (ans int64) {
	for i, w := range weights[1:] {
		weights[i] += w
	}
	weights = weights[:len(weights)-1]
	slices.Sort(weights)
	for _, w := range weights[len(weights)-k+1:] {
		ans += int64(w)
	}
	for _, w := range weights[:k-1] {
		ans -= int64(w)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。用快速选择算法可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈空间。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
