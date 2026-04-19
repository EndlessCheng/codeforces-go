正序遍历 $\textit{nums}$，可以求出每个前缀 $[0,i]$ 的最大值 $\textit{preMax}[i]$。

倒序遍历 $\textit{nums}$，可以求出每个后缀 $[i,n-1]$ 的最小值 $\textit{sufMin}[i]$。

答案为第一个满足 $\textit{preMax}[i] - \textit{sufMin}[i] \le k$ 的下标 $i$。如果不存在这样的下标，返回 $-1$。

代码实现时，可以在寻找答案的同时，计算 $\textit{preMax}$。

[本题视频讲解](https://www.bilibili.com/video/BV1agddBJEnQ/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def firstStableIndex(self, nums: list[int], k: int) -> int:
        n = len(nums)
        suf_min = [0] * n  # 后缀最小值
        suf_min[-1] = nums[-1]
        for i in range(n - 2, -1, -1):
            suf_min[i] = min(suf_min[i + 1], nums[i])

        pre_max = 0  # 前缀最大值
        for i, x in enumerate(nums):
            pre_max = max(pre_max, x)
            if pre_max - suf_min[i] <= k:
                return i
        return -1
```

```java [sol-Java]
class Solution {
    public int firstStableIndex(int[] nums, int k) {
        int n = nums.length;
        int[] sufMin = new int[n]; // 后缀最小值
        sufMin[n - 1] = nums[n - 1];
        for (int i = n - 2; i >= 0; i--) {
            sufMin[i] = Math.min(sufMin[i + 1], nums[i]);
        }

        int preMax = 0; // 前缀最大值
        for (int i = 0; i < n; i++) {
            preMax = Math.max(preMax, nums[i]);
            if (preMax - sufMin[i] <= k) {
                return i;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int firstStableIndex(vector<int>& nums, int k) {
        int n = nums.size();
        vector<int> suf_min(n); // 后缀最小值
        suf_min[n - 1] = nums[n - 1];
        for (int i = n - 2; i >= 0; i--) {
            suf_min[i] = min(suf_min[i + 1], nums[i]);
        }

        int pre_max = 0; // 前缀最大值
        for (int i = 0; i < n; i++) {
            pre_max = max(pre_max, nums[i]);
            if (pre_max - suf_min[i] <= k) {
                return i;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	sufMin := make([]int, n) // 后缀最小值
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i])
	}

	preMax := 0 // 前缀最大值
	for i, x := range nums {
		preMax = max(preMax, x)
		if preMax-sufMin[i] <= k {
			return i
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
