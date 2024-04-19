## 方法一：排序 + 二分

由于排序不影响答案，可以先排序。

然后枚举 $\textit{nums}[j]$，那么 $\textit{nums}[i]$ 需要满足

$$
\textit{lower} - \textit{nums}[j] \le \textit{nums}[i] \le \textit{upper} - \textit{nums}[j]
$$

并且 $0\le i < j$。

我们可以计算出 $\le \textit{upper} - \textit{nums}[j]$ 的元素个数，减去 $< \textit{lower} - \textit{nums}[j]$ 的元素个数，加入答案。

这都可以用二分查找求出，原理请看视频[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

附：[本题视频讲解](https://www.bilibili.com/video/BV1GY411i7RP/)

```py [sol-Python3]
class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        ans = 0
        nums.sort()
        for j, x in enumerate(nums):
            r = bisect_right(nums, upper - x, 0, j)  # <= upper-nums[j] 的 nums[i] 的个数
            l = bisect_left(nums, lower - x, 0, j)  # < lower-nums[j] 的 nums[i] 的个数
            ans += r - l
        return ans
```

```java [sol-Java]
class Solution {
    public long countFairPairs(int[] nums, int lower, int upper) {
        long ans = 0;
        Arrays.sort(nums);
        for (int j = 0; j < nums.length; ++j) {
            int r = lowerBound(nums, j, upper - nums[j] + 1); // <= upper-nums[j] 的 nums[i] 的个数
            int l = lowerBound(nums, j, lower - nums[j]); // < lower-nums[j] 的 nums[i] 的个数
            ans += r - l;
        }
        return ans;
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int right, int target) {
        int left = -1; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] < target
            // nums[right] >= target
            int mid = (left + right) >>> 1;
            if (nums[mid] < target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countFairPairs(vector<int> &nums, int lower, int upper) {
        long long ans = 0;
        ranges::sort(nums);
        for (int j = 0; j < nums.size(); ++j) {
            auto r = upper_bound(nums.begin(), nums.begin() + j, upper - nums[j]); // <= upper-nums[j] 的 nums[i] 的个数
            auto l = lower_bound(nums.begin(), nums.begin() + j, lower - nums[j]); // < lower-nums[j] 的 nums[i] 的个数
            ans += r - l;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countFairPairs(nums []int, lower, upper int) (ans int64) {
	sort.Ints(nums)
	for j, x := range nums {
		r := sort.SearchInts(nums[:j], upper-x+1) // <= upper-nums[j] 的 nums[i] 的个数
		l := sort.SearchInts(nums[:j], lower-x) // < lower-nums[j] 的 nums[i] 的个数
		ans += int64(r - l)
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，仅用到若干额外变量。

## 方法二：三指针

由于随着 $\textit{nums}[j]$ 的变大，$\textit{upper}-\textit{nums}[j]$ 和 $\textit{lower} - \textit{nums}[j]$ 都在变小，有单调性，可以用三指针实现。

```py [sol-Python3]
class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        ans = 0
        nums.sort()
        left = right = len(nums)
        for j, x in enumerate(nums):
            while right and nums[right - 1] > upper - x:
                right -= 1
            while left and nums[left - 1] >= lower - x:
                left -= 1
            ans += min(right, j) - min(left, j)
        return ans
```

```java [sol-Java]
class Solution {
    public long countFairPairs(int[] nums, int lower, int upper) {
        long ans = 0;
        Arrays.sort(nums);
        int left = nums.length;
        int right = nums.length;
        for (int j = 0; j < nums.length; j++) {
            while (right > 0 && nums[right - 1] > upper - nums[j]) {
                right--;
            }
            while (left > 0 && nums[left - 1] >= lower - nums[j]) {
                left--;
            }
            ans += Math.min(right, j) - Math.min(left, j);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countFairPairs(vector<int>& nums, int lower, int upper) {
        long long ans = 0;
        ranges::sort(nums);
        int left = nums.size(), right = left;
        for (int j = 0; j < nums.size(); j++) {
            while (right && nums[right - 1] > upper - nums[j]) {
                right--;
            }
            while (left && nums[left - 1] >= lower - nums[j]) {
                left--;
            }
            ans += min(right, j) - min(left, j);
        }
        return ans;
    }
};
```

```go [sol-Go]
func countFairPairs(nums []int, lower, upper int) (ans int64) {
	slices.Sort(nums)
	left, right := len(nums), len(nums)
	for j, x := range nums {
		for right > 0 && nums[right-1] > upper-x {
			right--
		}
		for left > 0 && nums[left-1] >= lower-x {
			left--
		}
		ans += int64(min(right, j)-min(left, j))
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，仅用到若干额外变量。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
