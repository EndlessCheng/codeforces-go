在遍历 $\textit{nums}$ 的过程中，维护元素 $x$ 上一次出现的位置 $\textit{last}[x]$。

- 遍历到 $\textit{nums}[i] = 1$ 时，我们需要知道上一个 $2$ 的位置 $j = \textit{last}[2]$，用 $i-j$ 更新答案的最小值。
- 遍历到 $\textit{nums}[i] = 2$ 时，我们需要知道上一个 $1$ 的位置 $j = \textit{last}[1]$，用 $i-j$ 更新答案的最小值。

> 注：无需找 $i$ 右边的 $j$，因为我们会继续遍历，遍历到 $j$ 那个位置时，它会去找左边的 $i$。

[本题视频讲解](https://www.bilibili.com/video/BV1dxXSBAE6F/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minAbsoluteDifference(self, nums: List[int]) -> int:
        ans = inf
        last = [-inf] * 2  # last[x] 表示 x+1 上一次出现的位置

        for i, x in enumerate(nums):
            if x > 0:
                # 如果 x 是 1，那么找上一个 2 的位置
                # 如果 x 是 2，那么找上一个 1 的位置
                x -= 1
                ans = min(ans, i - last[x ^ 1])
                last[x] = i

        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minAbsoluteDifference(int[] nums) {
        int n = nums.length;
        int ans = n;
        // last[x] 表示 x+1 上一次出现的位置
        int[] last = {-n, -n}; // i - (-n) >= n，不会让 ans 变小

        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (x > 0) {
                // 如果 x 是 1，那么找上一个 2 的位置
                // 如果 x 是 2，那么找上一个 1 的位置
                x--;
                ans = Math.min(ans, i - last[x ^ 1]);
                last[x] = i;
            }
        }

        return ans == n ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minAbsoluteDifference(vector<int>& nums) {
        int n = nums.size();
        int ans = n;
        // last[x] 表示 x+1 上一次出现的位置
        int last[2] = {-n, -n}; // i - (-n) >= n，不会让 ans 变小

        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (x > 0) {
                // 如果 x 是 1，那么找上一个 2 的位置
                // 如果 x 是 2，那么找上一个 1 的位置
                x--;
                ans = min(ans, i - last[x ^ 1]);
                last[x] = i;
            }
        }

        return ans == n ? -1 : ans;
    }
};
```

```go [sol-Go]
func minAbsoluteDifference(nums []int) int {
	n := len(nums)
	ans := n
	// last[x] 表示 x+1 上一次出现的位置
	last := [2]int{-n, -n} // i - (-n) >= n，不会让 ans 变小

	for i, x := range nums {
		if x > 0 {
			// 如果 x 是 1，那么找上一个 2 的位置
			// 如果 x 是 2，那么找上一个 1 的位置
			x--
			ans = min(ans, i-last[x^1])
			last[x] = i
		}
	}

	if ans == n {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面数据结构题单的「**§0.1 枚举右，维护左**」。

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
