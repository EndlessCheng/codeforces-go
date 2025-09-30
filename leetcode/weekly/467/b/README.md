题意：选出最大的 $\min(k,n)$ 个不同的数。

1. 把 $\textit{nums}$ 从大到小排序。
2. 把 $\textit{nums}$ 去重。
3. 如果 $\textit{nums}$ 的长度大于 $k$，返回 $\textit{nums}$ 的前 $k$ 个数；否则返回 $\textit{nums}$。

[本题视频讲解](https://www.bilibili.com/video/BV1TBpczdE8P/?t=1m30s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxKDistinct(self, nums: list[int], k: int) -> list[int]:
        return nlargest(k, set(nums))
```

```py [sol-Python3 写法二]
class Solution:
    def maxKDistinct(self, nums: list[int], k: int) -> list[int]:
        a = sorted(set(nums), reverse=True)  # 去重+排序
        return a[:k]
```

```java [sol-Java]
class Solution {
    public int[] maxKDistinct(int[] nums, int k) {
        Arrays.sort(nums);

        int uniques = removeDuplicates(nums);
        int size = Math.min(uniques, k);

        int[] ans = new int[size];
        for (int i = 0; i < size; i++) {
            ans[i] = nums[uniques - 1 - i]; // 题目要求从大到小
        }
        return ans;
    }

    // 26. 删除有序数组中的重复项
    private int removeDuplicates(int[] nums) {
        int k = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != nums[i - 1]) { // nums[i] 不是重复项
                nums[k++] = nums[i]; // 保留 nums[i]
            }
        }
        return k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maxKDistinct(vector<int>& nums, int k) {
        ranges::sort(nums, greater());
        nums.erase(ranges::unique(nums).begin(), nums.end()); // 原地去重
        if (nums.size() > k) {
            nums.resize(k);
        }
        return nums;
    }
};
```

```go [sol-Go]
func maxKDistinct(nums []int, k int) []int {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	nums = slices.Compact(nums) // 原地去重
	return nums[:min(k, len(nums))]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n\log k)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。原地去重则空间复杂度为 $\mathcal{O}(1)$，忽略排序的栈开销，返回值不计入。

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
