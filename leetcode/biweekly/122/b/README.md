**适用场景**：按照题目要求，数组会被分割成若干组，每一组的判断/处理逻辑是相同的。

**核心思想**：

- **外层循环**负责遍历组之前的准备工作（记录开始位置），和遍历组之后的工作（排序）。
- **内层循环**负责遍历组，找出这一组最远在哪结束。

这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组（易错点）。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。

## 方法一：直接排序

内层循环不断向后遍历，直到到达数组末尾，或者遇到二进制中的 $1$ 的个数不同的元素。

组内元素可以相邻交换，根据冒泡排序的思想，组内元素是可以从小到大排序的。

对于每一组，都从小到大排序。如果最后数组是有序的，返回 $\texttt{true}$，否则返回 $\texttt{false}$。

```py [sol-Python3]
class Solution:
    def canSortArray(self, nums: List[int]) -> bool:
        n = len(nums)
        i = 0
        while i < n:
            start = i
            ones = nums[i].bit_count()
            i += 1
            while i < n and nums[i].bit_count() == ones:
                i += 1
            nums[start:i] = sorted(nums[start:i])
        return all(x <= y for x, y in pairwise(nums))
```

```java [sol-Java]
class Solution {
    public boolean canSortArray(int[] nums) {
        int n = nums.length;
        for (int i = 0; i < n;) {
            int start = i;
            int ones = Integer.bitCount(nums[i++]);
            while (i < n && Integer.bitCount(nums[i]) == ones) {
                i++;
            }
            Arrays.sort(nums, start, i);
        }
        for (int i = 1; i < n; i++) {
            if (nums[i] < nums[i - 1]) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canSortArray(vector<int>& nums) {
        for (int i = 0, n = nums.size(); i < n;) {
            int start = i;
            int ones = __builtin_popcount(nums[i++]);
            while (i < n && __builtin_popcount(nums[i]) == ones) {
                i++;
            }
            sort(nums.begin() + start, nums.begin() + i);
        }
        return ranges::is_sorted(nums);
    }
};
```

```go [sol-Go]
func canSortArray(nums []int) bool {
	for i, n := 0, len(nums); i < n; {
		start := i
		ones := bits.OnesCount(uint(nums[i]))
		i++
		for i < n && bits.OnesCount(uint(nums[i])) == ones {
			i++
		}
		slices.Sort(nums[start:i])
	}
	return slices.IsSorted(nums)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 方法二：记录每一段的最小值和最大值

对于每一组，如果这一组的每个数，都大于等于上一组的最大值 $\textit{preMax}$，那么我们就能把数组排成递增的，否则不行。

由于题目保证 $\textit{nums}[i] > 0$，我们可以把 $\textit{preMax}$ 和本组最大值 $\textit{mx}$ 都初始化成 $0$。

```py [sol-Python3]
class Solution:
    def canSortArray(self, nums: List[int]) -> bool:
        n = len(nums)
        i = pre_max = 0
        while i < n:
            mx = 0
            ones = nums[i].bit_count()
            while i < n and nums[i].bit_count() == ones:
                x = nums[i]
                if x < pre_max:  # 无法排成有序的
                    return False
                mx = max(mx, x)  # 更新本组最大值
                i += 1
            pre_max = mx
        return True
```

```java [sol-Java]
class Solution {
    public boolean canSortArray(int[] nums) {
        int n = nums.length;
        int preMax = 0;
        for (int i = 0; i < n;) {
            int mx = 0;
            int ones = Integer.bitCount(nums[i]);
            while (i < n && Integer.bitCount(nums[i]) == ones) {
                if (nums[i] < preMax) { // 无法排成有序的
                    return false;
                }
                mx = Math.max(mx, nums[i++]); // 更新本组最大值
            }
            preMax = mx;
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canSortArray(vector<int>& nums) {
        int n = nums.size();
        int pre_max = 0;
        for (int i = 0; i < n;) {
            int mx = 0;
            int ones = __builtin_popcount(nums[i]);
            while (i < n && __builtin_popcount(nums[i]) == ones) {
                if (nums[i] < pre_max) { // 无法排成有序的
                    return false;
                }
                mx = max(mx, nums[i++]); // 更新本组最大值
            }
            pre_max = mx;
        }
        return true;
    }
};
```

```go [sol-Go]
func canSortArray(nums []int) bool {
	preMax := 0
	for i, n := 0, len(nums); i < n; {
		mx := 0
		ones := bits.OnesCount(uint(nums[i]))
		for ; i < n && bits.OnesCount(uint(nums[i])) == ones; i++ {
			if nums[i] < preMax { // 无法排成有序的
				return false
			}
			mx = max(mx, nums[i]) // 更新本组最大值
		}
		preMax = mx
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 练习

一般来说，分组循环的模板如下（可根据题目调整）：

```py
n = len(nums)
i = 0
while i < n:
    start = i
    while i < n and ...:
        i += 1
    # 从 start 到 i-1 是一组
    # 下一组从 i 开始，无需 i += 1
```

学会一个模板是远远不够的，需要大量练习才能灵活运用。

- [1446. 连续字符](https://leetcode.cn/problems/consecutive-characters/) 1165
- [1869. 哪种连续子字符串更长](https://leetcode.cn/problems/longer-contiguous-segments-of-ones-than-zeros/) 1205
- [1957. 删除字符使字符串变好](https://leetcode.cn/problems/delete-characters-to-make-fancy-string/) 1358
- [2110. 股票平滑下跌阶段的数目](https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/) 1408
- [228. 汇总区间](https://leetcode.cn/problems/summary-ranges/)
- [2760. 最长奇偶子数组](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/) 1420
- [1887. 使数组元素相等的减少操作次数](https://leetcode.cn/problems/reduction-operations-to-make-the-array-elements-equal/) 1428
- [2038. 如果相邻两个颜色均相同则删除当前颜色](https://leetcode.cn/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/) 1468
- [1759. 统计同质子字符串的数目](https://leetcode.cn/problems/count-number-of-homogenous-substrings/) 1491
- [1578. 使绳子变成彩色的最短时间](https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/) 1574
- [1839. 所有元音按顺序排布的最长子字符串](https://leetcode.cn/problems/longest-substring-of-all-vowels-in-order/) 1580
- [2765. 最长交替子序列](https://leetcode.cn/problems/longest-alternating-subarray/) 1581

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
