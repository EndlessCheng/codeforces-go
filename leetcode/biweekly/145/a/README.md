操作相当于：

- 把 $\textit{nums}$ 中的最大值都变小，但不能低于次大值。

那么最优策略是变成次大值。

继续操作，次大值再变成第三大的值，依此类推。

分类讨论：

- 如果 $k < \min(nums)$，操作次数为 $\textit{nums}$ 中的不同元素个数。
- 如果 $k = \min(nums)$，操作次数为 $\textit{nums}$ 中的不同元素个数减一。
- 如果 $k > \min(nums)$，无法操作。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1YeqHYSEXv/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        mn = min(nums)
        if k > mn:
            return -1
        return len(set(nums)) - (k == mn)
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int mn = Arrays.stream(nums).min().getAsInt();
        if (k > mn) {
            return -1;
        }
        Set<Integer> set = new HashSet<>();
        for (int x : nums) {
            set.add(x);
        }
        return set.size() - (k == mn ? 1 : 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        int mn = ranges::min(nums);
        if (k > mn) {
            return -1;
        }
        return unordered_set<int>(nums.begin(), nums.end()).size() - (k == mn);
    }
};
```

```go [sol-Go]
func minOperations(nums []int, k int) int {
	mn := slices.Min(nums)
	if k > mn {
		return -1
	}
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
	}
	if k == mn {
		return len(set) - 1
	}
	return len(set)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
