[本题视频讲解](https://www.bilibili.com/video/BV1994y1A7oo/)

## 前置知识：中位数贪心

为方便描述，将 $\textit{nums}$ 简记为 $a$。

**定理**：将 $a$ 的所有元素变为 $a$ 的**中位数**是最优的。

**证明**：设 $a$ 的长度为 $n$，设要将所有 $a[i]$ 变为 $x$。假设 $a$ 已经从小到大排序。首先，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之外，那么 $x$ 向区间方向移动可以使距离和变小；同时，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之内，无论如何移动 $x$，它到 $a[0]$ 和 $a[n-1]$ 的距离和都是一个定值 $a[n-1]-a[0]$，那么去掉 $a[0]$ 和 $a[n-1]$ 这两个最左最右的数，问题规模缩小。不断缩小问题规模，如果最后剩下 $1$ 个数，那么 $x$ 就取它；如果最后剩下 $2$ 个数，那么 $x$ 取这两个数之间的任意值都可以（包括这两个数）。因此 $x$ 可以取 $a[n/2]$。

本题回文数可能取不到中位数，我们可以找离中位数最近的数。

## 前置知识：二分查找

请看视频[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)

## 思路

首先预处理出 $10^9$ 内的回文数，这可以通过枚举回文数的左半部分得到。

> 注：也可以多预处理一个 $10^9+1$，不过考虑到 $10^9-1$ 也是回文数，这一步可以省略。

然后二分找离 $\textit{nums}$ 的中位数最近的回文数（中位数左右两侧都要找），作为要变成的数字。具体见代码。

```py [sol-Python3]
# 严格按顺序从小到大生成所有回文数（不用字符串转换）
pal = []
base = 1
while base <= 10000:
    # 生成奇数长度回文数
    for i in range(base, base * 10):
        x = i
        t = i // 10
        while t:
            x = x * 10 + t % 10
            t //= 10
        pal.append(x)
    # 生成偶数长度回文数
    if base <= 1000:
        for i in range(base, base * 10):
            x = t = i
            while t:
                x = x * 10 + t % 10
                t //= 10
            pal.append(x)
    base *= 10
pal.append(1_000_000_001)  # 哨兵，防止下面代码中的 i 下标越界

class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        # 注：排序只是为了找中位数，如果用快速选择算法，可以做到 O(n)
        nums.sort()

        # 返回 nums 中的所有数变成 pal[i] 的总代价
        def cost(i: int) -> int:
            target = pal[i]
            return sum(abs(x - target) for x in nums)

        n = len(nums)
        i = bisect_left(pal, nums[(n - 1) // 2])  # 二分找中位数右侧最近的回文数
        if pal[i] <= nums[n // 2]:  # 回文数在中位数范围内
            return cost(i)  # 直接变成 pal[i]
        return min(cost(i - 1), cost(i))  # 枚举离中位数最近的两个回文数 pal[i-1] 和 pal[i]
```

```java [sol-Java]
class Solution {
    private static final int[] pal = new int[109999];

    static {
        // 严格按顺序从小到大生成所有回文数（不用字符串转换）
        int palIdx = 0;
        for (int base = 1; base <= 10000; base *= 10) {
            // 生成奇数长度回文数
            for (int i = base; i < base * 10; i++) {
                int x = i;
                for (int t = i / 10; t > 0; t /= 10) {
                    x = x * 10 + t % 10;
                }
                pal[palIdx++] = x;
            }
            // 生成偶数长度回文数
            if (base <= 1000) {
                for (int i = base; i < base * 10; i++) {
                    int x = i;
                    for (int t = i; t > 0; t /= 10) {
                        x = x * 10 + t % 10;
                    }
                    pal[palIdx++] = x;
                }
            }
        }
        pal[palIdx++] = 1_000_000_001; // 哨兵，防止下面代码中的 i 下标越界
    }

    public long minimumCost(int[] nums) {
        // 注：排序只是为了找中位数，如果用快速选择算法，可以做到 O(n)
        Arrays.sort(nums);
        int n = nums.length;

        // 二分找中位数右侧最近的回文数
        int i = lowerBound(nums[(n - 1) / 2]);

        // 回文数在中位数范围内
        if (pal[i] <= nums[n / 2]) {
            return cost(nums, i); // 直接变成 pal[i]
        }

        // 枚举离中位数最近的两个回文数 pal[i-1] 和 pal[i]
        return Math.min(cost(nums, i - 1), cost(nums, i));
    }

    // 返回 nums 中的所有数变成 pal[i] 的总代价
    private long cost(int[] nums, int i) {
        int target = pal[i];
        long sum = 0;
        for (int x : nums) {
            sum += Math.abs(x - target);
        }
        return sum;
    }

    // 开区间写法
    private int lowerBound(int target) {
        int left = -1, right = pal.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // pal[left] < target
            // pal[right] >= target
            int mid = left + (right - left) / 2;
            if (pal[mid] < target)
                left = mid; // 范围缩小到 (mid, right)
            else
                right = mid; // 范围缩小到 (left, mid)
        }
        return right; // 或者 left+1
    }
}
```

```cpp [sol-C++]
vector<int> pal;

auto init = [] {
    // 严格按顺序从小到大生成所有回文数（不用字符串转换）
    for (int base = 1; base <= 10000; base *= 10) {
        // 生成奇数长度回文数
        for (int i = base; i < base * 10; i++) {
            int x = i;
            for (int t = i / 10; t; t /= 10) {
                x = x * 10 + t % 10;
            }
            pal.push_back(x);
        }
        // 生成偶数长度回文数
        if (base <= 1000) {
            for (int i = base; i < base * 10; i++) {
                int x = i;
                for (int t = i; t; t /= 10) {
                    x = x * 10 + t % 10;
                }
                pal.push_back(x);
            }
        }
    }
    pal.push_back(1'000'000'001); // 哨兵，防止下面代码中的 i 下标越界
    return 0;
}();

class Solution {
public:
    long long minimumCost(vector<int> &nums) {
        // 注：排序只是为了找中位数，如果用快速选择算法，可以做到 O(n)
        sort(nums.begin(), nums.end());

        // 返回 nums 中的所有数变成 pal[i] 的总代价
        auto cost = [&](int i) -> long long {
            int target = pal[i];
            long long sum = 0;
            for (int x: nums) {
                sum += abs(x - target);
            }
            return sum;
        };

        int n = nums.size();
        // 二分找中位数右侧最近的回文数
        int i = lower_bound(pal.begin(), pal.end(), nums[(n - 1) / 2]) - pal.begin();
        if (pal[i] <= nums[n / 2]) { // 回文数在中位数范围内
            return cost(i); // 直接变成 pal[i]
        }
        return min(cost(i - 1), cost(i)); // 枚举离中位数最近的两个回文数 pal[i-1] 和 pal[i]
    }
};
```

```cpp [sol-C++ 快速选择]
vector<int> pal;

auto init = [] {
    pal.push_back(0); // 哨兵，防止下面代码中的 i 下标越界
    // 严格按顺序从小到大生成所有回文数（不用字符串转换）
    for (int base = 1; base <= 10000; base *= 10) {
        // 生成奇数长度回文数
        for (int i = base; i < base * 10; i++) {
            int x = i;
            for (int t = i / 10; t; t /= 10) {
                x = x * 10 + t % 10;
            }
            pal.push_back(x);
        }
        // 生成偶数长度回文数
        if (base <= 1000) {
            for (int i = base; i < base * 10; i++) {
                int x = i;
                for (int t = i; t; t /= 10) {
                    x = x * 10 + t % 10;
                }
                pal.push_back(x);
            }
        }
    }
    pal.push_back(1'000'000'001); // 哨兵，防止下面代码中的 i 下标越界
    return 0;
}();

class Solution {
public:
    long long minimumCost(vector<int> &nums) {
        // 返回 nums 中的所有数变成 pal[i] 的总代价
        auto cost = [&](int i) -> long long {
            int target = pal[i];
            long long sum = 0;
            for (int x: nums) {
                sum += abs(x - target);
            }
            return sum;
        };

        int m = (nums.size() - 1) / 2;
        nth_element(nums.begin(), nums.begin() + m, nums.end()); // 快速选择
        int mid = nums[m]; // 中位数

        // 二分找中位数右侧最近的回文数
        int i = lower_bound(pal.begin(), pal.end(), mid) - pal.begin();

        // 枚举离中位数最近的两个回文数 pal[i-1] 和 pal[i]
        return min(cost(i - 1), cost(i));
    }
};
```

```go [sol-Go]
var pal = make([]int, 0, 109999)

func init() {
	// 按顺序从小到大生成所有回文数
	for base := 1; base <= 10000; base *= 10 {
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			pal = append(pal, x)
		}
		if base <= 1000 {
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				pal = append(pal, x)
			}
		}
	}
	pal = append(pal, 1_000_000_001) // 哨兵，防止下标越界
}

func minimumCost(nums []int) int64 {
	// 注：排序只是为了找中位数，如果用快速选择算法，可以做到 O(n)
	slices.Sort(nums)

	// 返回所有 nums[i] 变成 pal[i] 的总代价
	cost := func(i int) (res int64) {
		target := pal[i]
		for _, x := range nums {
			res += int64(abs(x - target))
		}
		return
	}

	n := len(nums)
	i := sort.SearchInts(pal, nums[(n-1)/2]) // 二分找中位数右侧最近的回文数
	if pal[i] <= nums[n/2] { // 回文数在中位数范围内
		return cost(i) // 直接变成 pal[i]
	}
	return min(cost(i-1), cost(i)) // 枚举离中位数最近的两个回文数 pal[i-1] 和 pal[i]
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + \log U)$ 或 $\mathcal{O}(n + \log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U$ 为 $10^9+1$ 内的回文数个数。忽略预处理的时间。
- 空间复杂度：$\mathcal{O}(1)$。忽略预处理的空间。

## 思考题

能否直接找离某个数最近的回文数？

这题是 [564. 寻找最近的回文数](https://leetcode.cn/problems/find-the-closest-palindrome/)，注意数据范围。

本题也可以从中位数出发，向左右暴力找回文数。在本题的数据范围下，最大的回文数间隔只有 $10011001 - 10000001 = 11000$。

## 生成回文数

- [2081. k 镜像数字的和](https://leetcode.cn/problems/sum-of-k-mirror-numbers/)

## 中位数贪心（右边数字为难度分）

- [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
- [2033. 获取单值网格的最小操作数](https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid/) 1672
- [2448. 使数组相等的最小开销](https://leetcode.cn/problems/minimum-cost-to-make-array-equal/) 2005
- [2607. 使子数组元素和相等](https://leetcode.cn/problems/make-k-subarray-sums-equal/) 2071
- [1703. 得到连续 K 个 1 的最少相邻交换次数](https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/) 2467
