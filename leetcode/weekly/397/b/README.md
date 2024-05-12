我们要计算的是一个类似后缀和的概念，例如 $k=1$ 时，计算的就是后缀和的最大值，$k=2$ 时，计算的是间隔一个数的后缀和。

枚举终点为 $i$，我们要计算的是下标为

$$
i,i-k,i-2k,\cdots
$$

的元素之和的最小值。

一边遍历，一边累加元素和，取元素和的最大值，即为答案。

$i$ 可以从 $n-k$ 开始枚举到 $n-1$。

请看 [视频讲解](https://www.bilibili.com/video/BV1bx4y1i7rP/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maximumEnergy(self, energy: List[int], k: int) -> int:
        n = len(energy)
        ans = -inf
        for i in range(n - k, n):
            ans = max(ans, max(accumulate(energy[j] for j in range(i, -1, -k))))
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumEnergy(int[] energy, int k) {
        int n = energy.length;
        int ans = Integer.MIN_VALUE;
        for (int i = n - k; i < n; i++) {
            int s = 0;
            for (int j = i; j >= 0; j -= k) {
                s += energy[j];
                ans = Math.max(ans, s);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumEnergy(vector<int>& energy, int k) {
        int n = energy.size(), ans = INT_MIN;
        for (int i = n - k; i < n; i++) {
            int s = 0;
            for (int j = i; j >= 0; j -= k) {
                s += energy[j];
                ans = max(ans, s);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	ans := math.MinInt
	for i := n - k; i < n; i++ {
		s := 0
		for j := i; j >= 0; j -= k {
			s += energy[j]
			ans = max(ans, s)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{energy}$ 的长度。每个元素恰好访问一次。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
