把两个数组从小到大排序。

由于只移除两个元素，我们可以枚举 $\textit{nums}_1$ 中**保留**的最小元素是 $\textit{nums}_1[0]$ 还是 $\textit{nums}_1[1]$ 还是 $\textit{nums}_1[2]$。

设 $\textit{diff} = \textit{nums}_2[0] - \textit{nums}_1[i]$，其中 $\textit{nums}_1[i]$ 是保留的最小元素。

问题变成：

- 判断 $\textit{nums}_2$ 是否为序列 $\textit{nums}_1[i] + \textit{diff}$ 的子序列。

做法同 [392. 判断子序列](https://leetcode.cn/problems/is-subsequence/)，可以用**同向双指针**解决。

请看 [视频讲解](https://www.bilibili.com/video/BV1Pw4m1C79N/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumAddedInteger(self, nums1: List[int], nums2: List[int]) -> int:
        nums1.sort()
        nums2.sort()
        # 枚举保留 nums1[2] 或者 nums1[1] 或者 nums1[0]
        # 倒着枚举是因为 nums1[i] 越大答案越小，第一个满足的就是答案
        for i in range(2, 0, -1):
            diff = nums2[0] - nums1[i]
            # 在 {nums1[i] + diff} 中找子序列 nums2
            j = 0
            for v in nums1[i:]:
                if nums2[j] == v + diff:
                    j += 1
                    # nums2 是 {nums1[i] + diff} 的子序列
                    if j == len(nums2):
                        return diff
        # 题目保证答案一定存在
        return nums2[0] - nums1[0]
```

```java [sol-Java]
class Solution {
    public int minimumAddedInteger(int[] nums1, int[] nums2) {
        Arrays.sort(nums1);
        Arrays.sort(nums2);
        // 枚举保留 nums1[2] 或者 nums1[1] 或者 nums1[0]
        // 倒着枚举是因为 nums1[i] 越大答案越小，第一个满足的就是答案
        for (int i = 2; i > 0; i--) {
            int diff = nums2[0] - nums1[i];
            // 在 {nums1[i] + diff} 中找子序列 nums2
            int j = 0;
            for (int k = i; k < nums1.length; k++) {
                if (nums2[j] == nums1[k] + diff && ++j == nums2.length) {
                    // nums2 是 {nums1[i] + diff} 的子序列
                    return diff;
                }
            }
        }
        // 题目保证答案一定存在
        return nums2[0] - nums1[0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumAddedInteger(vector<int>& nums1, vector<int>& nums2) {
        ranges::sort(nums1);
        ranges::sort(nums2);
        // 枚举保留 nums1[2] 或者 nums1[1] 或者 nums1[0]
        // 倒着枚举是因为 nums1[i] 越大答案越小，第一个满足的就是答案
        for (int i = 2; i; i--) {
            int diff = nums2[0] - nums1[i];
            // 在 {nums1[i] + diff} 中找子序列 nums2
            int j = 0;
            for (int k = i; k < nums1.size(); k++) {
                if (nums2[j] == nums1[k] + diff && ++j == nums2.size()) {
                    // nums2 是 {nums1[i] + diff} 的子序列
                    return diff;
                }
            }
        }
        // 题目保证答案一定存在
        return nums2[0] - nums1[0];
    }
};
```

```go [sol-Go]
func minimumAddedInteger(nums1, nums2 []int) int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	// 枚举保留 nums1[2] 或者 nums1[1] 或者 nums1[0]
	// 倒着枚举是因为 nums1[i] 越大答案越小，第一个满足的就是答案
	for i := 2; i > 0; i-- {
		diff := nums2[0] - nums1[i]
		// 在 {nums1[i] + diff} 中找子序列 nums2
		j := 0
		for _, v := range nums1[i:] {
			if nums2[j] == v+diff {
				j++
				// nums2 是 {nums1[i] + diff} 的子序列
				if j == len(nums2) {
					return diff
				}
			}
		}
	}
	// 题目保证答案一定存在
	return nums2[0] - nums1[0]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

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
