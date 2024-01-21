[视频讲解](https://www.bilibili.com/video/BV1oV411D7gB/) 第二题。

**适用场景**：按照题目要求，数组会被分割成若干组，每一组的判断/处理逻辑是相同的。

**核心思想**：

- 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的工作（排序）。
- 内层循环负责遍历组，找出这一组最远在哪结束。

这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组（易错点）。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。

## 方法一：直接排序

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
        for (int i = 0; i < n; ) {
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
    bool canSortArray(vector<int> &nums) {
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

如果每一段的最小值，都大于等于上一段的最大值，那么我们就能把数组排成递增的。如果最小值小于上一段的最大值，那么无法排成递增的。

```py [sol-Python3]
class Solution:
    def canSortArray(self, nums: List[int]) -> bool:
        n = len(nums)
        i = pre_max = 0
        while i < n:
            mn = mx = nums[i]
            ones = mn.bit_count()
            i += 1
            while i < n and nums[i].bit_count() == ones:
                x = nums[i]
                if x < mn:
                    mn = x
                elif x > mx:
                    mx = x
                i += 1
            if mn < pre_max:
                return False
            pre_max = mx
        return True
```

```java [sol-Java]
class Solution {
    public boolean canSortArray(int[] nums) {
        int n = nums.length;
        int i = 0, preMax = 0;
        while (i < n) {
            int mn = nums[i], mx = mn;
            int ones = Integer.bitCount(mn);
            for (i++; i < n && Integer.bitCount(nums[i]) == ones; i++) {
                mn = Math.min(mn, nums[i]);
                mx = Math.max(mx, nums[i]);
            }
            if (mn < preMax) {
                return false;
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
    bool canSortArray(vector<int> &nums) {
        int n = nums.size();
        int i = 0, pre_max = 0;
        while (i < n) {
            int mn = nums[i], mx = mn;
            int ones = __builtin_popcount(mn);
            for (i++; i < n && __builtin_popcount(nums[i]) == ones; i++) {
                mn = min(mn, nums[i]);
                mx = max(mx, nums[i]);
            }
            if (mn < pre_max) {
                return false;
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
		mn, mx := nums[i], nums[i]
		ones := bits.OnesCount(uint(mn))
		for i++; i < n && bits.OnesCount(uint(nums[i])) == ones; i++ {
			mn = min(mn, nums[i])
			mx = max(mx, nums[i])
		}
		if mn < preMax {
			return false
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

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
