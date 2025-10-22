## 方法一：差分

**前置知识**：[差分原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)。

设 $x = \textit{nums}[i]$。根据题意，$x$ 可以变成 $[x-k,x+k]$ 中的整数。

题目让我们计算操作后，最多有多少个数相同。

例如 $\textit{nums}=[2,4]$，$k=1$，$\textit{numOperations}=2$。$2$ 可以变成 $[1,3]$ 中的整数，$4$ 可以变成 $[3,5]$ 中的整数。$2$ 和 $4$ 都可以变成 $3$，所以答案是 $2$。

一般地，$x$ 可以变成 $[x-k,x+k]$ 中的整数，我们可以把 $[x-k,x+k]$ 中的每个整数的出现次数都加一，然后统计出现次数的最大值。这可以用差分实现。

计算差分的前缀和。设有 $\textit{sumD}$ 个数可以变成 $y$。

如果 $y$ 不在 $\textit{nums}$ 中，那么 $y$ 的最大出现次数不能超过 $\textit{numOperations}$，与 $\textit{sumD}$ 取最小值，得 $\min(\textit{sumD}, \textit{numOperations})$。

如果 $y$ 在 $\textit{nums}$ 中，且出现了 $\textit{cnt}$ 次，那么有 $\textit{sumD}-\textit{cnt}$ 个其他元素（不等于 $y$ 的数）可以变成 $y$，但这不能超过 $\textit{numOperations}$，所以有

$$
\min(\textit{sumD}-\textit{cnt}, \textit{numOperations})
$$

个其他元素可以变成 $y$，再加上 $y$ 自身出现的次数 $\textit{cnt}$，得到 $y$ 的最大频率为

$$
\textit{cnt} + \min(\textit{sumD}-\textit{cnt}, \textit{numOperations}) = \min(\textit{sumD}, \textit{cnt}+\textit{numOperations})
$$

注意上式兼容 $y$ 不在 $\textit{nums}$ 中的情况，此时 $\textit{cnt}=0$。

[本题视频讲解](https://www.bilibili.com/video/BV1cgmBYqEhu/?t=3m37s)，欢迎点赞关注~

### 答疑

**问**：为什么代码只考虑在 $\textit{diff}$ 和 $\textit{nums}$ 中的数字？

**答**：比如 $x$ 在 $\textit{diff}$ 中，$x+1,x+2,\ldots$ 不在 $\textit{diff}$ 中，那么 $x+1,x+2,\ldots$ 的 $\textit{sumD}$ 和 $\textit{x}$ 的是一样的，无需重复计算。此外，要想算出比 $\min(\textit{sumD}, \textit{cnt}+\textit{numOperations})$ 更大的数，要么 $\textit{sumD}$ 变大，要么 $\textit{cnt}$ 变大。「变大」时的 $x$ 必然在 $\textit{diff}$ 或 $\textit{nums}$ 中。

```py [sol-Python3]
class Solution:
    def maxFrequency(self, nums: List[int], k: int, numOperations: int) -> int:
        cnt = defaultdict(int)
        diff = defaultdict(int)
        for x in nums:
            cnt[x] += 1
            diff[x]  # 把 x 插入 diff，以保证下面能遍历到 x
            diff[x - k] += 1  # 把 [x-k,x+k] 中的每个整数的出现次数都加一
            diff[x + k + 1] -= 1

        ans = sum_d = 0
        for x, d in sorted(diff.items()):
            sum_d += d
            ans = max(ans, min(sum_d, cnt[x] + numOperations))
        return ans
```

```java [sol-Java]
class Solution {
    int maxFrequency(int[] nums, int k, int numOperations) {
        Map<Integer, Integer> cnt = new HashMap<>();
        Map<Integer, Integer> diff = new TreeMap<>();
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum); // cnt[x]++
            diff.putIfAbsent(x, 0); // 把 x 插入 diff，以保证下面能遍历到 x
            // 把 [x-k, x+k] 中的每个整数的出现次数都加一
            diff.merge(x - k, 1, Integer::sum); // diff[x-k]++
            diff.merge(x + k + 1, -1, Integer::sum); // diff[x+k+1]--
        }

        int ans = 0;
        int sumD = 0;
        for (Map.Entry<Integer, Integer> e : diff.entrySet()) {
            sumD += e.getValue();
            ans = Math.max(ans, Math.min(sumD, cnt.getOrDefault(e.getKey(), 0) + numOperations));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFrequency(vector<int>& nums, int k, int numOperations) {
        unordered_map<int, int> cnt;
        map<int, int> diff;
        for (int x : nums) {
            cnt[x]++;
            diff[x]; // 把 x 插入 diff，以保证下面能遍历到 x
            diff[x - k]++; // 把 [x-k, x+k] 中的每个整数的出现次数都加一
            diff[x + k + 1]--;
        }

        int ans = 0, sum_d = 0;
        for (auto& [x, d] : diff) {
            sum_d += d;
            ans = max(ans, min(sum_d, cnt[x] + numOperations));
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxFrequency(nums []int, k, numOperations int) (ans int) {
	cnt := map[int]int{}
	diff := map[int]int{}
	for _, x := range nums {
		cnt[x]++
		diff[x] += 0 // 把 x 插入 diff，以保证下面能遍历到 x
		diff[x-k]++  // 把 [x-k, x+k] 中的每个整数的出现次数都加一
		diff[x+k+1]--
	}

	sumD := 0
	for _, x := range slices.Sorted(maps.Keys(diff)) {
		sumD += diff[x]
		ans = max(ans, min(sumD, cnt[x]+numOperations))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：同向三指针 + 同向双指针

### 核心思路

1. 计算有多少个数能变成 $x$，其中 $x = \textit{nums}[i]$。用**同向三指针**实现。
2. 计算有多少个数能变成 $x$，其中 $x$ 不一定在 $\textit{nums}$ 中。用**同向双指针**实现。

### 同向三指针

把 $\textit{nums}$ 从小到大排序。

遍历 $\textit{nums}$。设 $x=\textit{nums}[i]$，计算元素值在 $[x-k,x+k]$ 中的元素个数，这些元素都可以变成 $x$。

遍历的同时，维护左指针 $\textit{left}$，它是最小的满足

$$
\textit{nums}[\textit{left}] \ge x - k
$$

的下标。

遍历的同时，维护右指针 $\textit{right}$，它是最小的满足

$$
\textit{nums}[\textit{right}] > x + k
$$

的下标。如果不存在，则 $\textit{right}=n$。

下标在左闭右开区间 $[\textit{left},\textit{right})$ 中的元素，都可以变成 $x$。这有

$$
\textit{sumD} = \textit{right} - \textit{left}
$$

个。

遍历的同时，求出 $x$ 有 $\textit{cnt}$ 个。然后用方法一的公式，更新答案的最大值。

### 同向双指针

同向三指针没有考虑「变成不在 $\textit{nums}$ 中的数」这种情况。

然而，不在 $\textit{nums}$ 中的数太多了！怎么减少计算量？

假设都变成 $y$，那么只有 $[y-k,y+k]$ 中的数能变成 $y$。

把 $\textit{nums}[i]$ 视作一维数轴上的点，想象一个窗口 $[y-k,y+k]$ 在不断地向右滑动，什么时候窗口内的元素个数才会变大？

- 如果我们从 $[y-k-1,y+k-1]$ 移动到 $[y-k,y+k]$，且 $y+k$ 不在 $\textit{nums}$ 中，此时窗口内的元素个数不会变大，甚至因为左端点收缩了，元素个数可能会变小。
- 所以，只有当 $y+k$ 恰好在 $\textit{nums}$ 中时，窗口内的元素个数才可能会变大。
- **结论**：我们只需考虑 $y+k$ 在 $\textit{nums}$ 中时的 $y$！

于是，枚举 $x=\textit{nums}[\textit{right}]$，计算元素值在 $[x-2k,x]$ 中的元素个数，这些元素都可以变成同一个数 $y=x-k$。

左指针 $\textit{left}$ 是最小的满足

$$
\textit{nums}[\textit{left}] \ge x-2k
$$

的下标。

计算好 $\textit{left}$ 后，下标在 $[\textit{left}, \textit{right}]$ 中的数可以变成一样的，这有

$$
\textit{right} - \textit{left} + 1
$$

个。注意上式不能超过 $\textit{numOperations}$。

### 小优化

由于同向双指针算出的结果不超过 $\textit{numOperations}$，所以当同向三指针计算完毕后，如果发现答案已经 $\ge \textit{numOperations}$，那么无需计算同向双指针。

```py [sol-Python3]
class Solution:
    def maxFrequency(self, nums: List[int], k: int, numOperations: int) -> int:
        nums.sort()

        n = len(nums)
        ans = cnt = left = right = 0
        for i, x in enumerate(nums):
            cnt += 1
            # 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
            if i < n - 1 and x == nums[i + 1]:
                continue
            while nums[left] < x - k:
                left += 1
            while right < n and nums[right] <= x + k:
                right += 1
            ans = max(ans, min(right - left, cnt + numOperations))
            cnt = 0

        if ans >= numOperations:
            return ans

        left = 0
        for right, x in enumerate(nums):
            while nums[left] < x - k * 2:
                left += 1
            ans = max(ans, right - left + 1)
        return min(ans, numOperations)  # 最后和 numOperations 取最小值
```

```java [sol-Java]
class Solution {
    public int maxFrequency(int[] nums, int k, int numOperations) {
        Arrays.sort(nums);

        int n = nums.length;
        int ans = 0;
        int cnt = 0;
        int left = 0;
        int right = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            cnt++;
            // 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
            if (i < n - 1 && x == nums[i + 1]) {
                continue;
            }
            while (nums[left] < x - k) {
                left++;
            }
            while (right < n && nums[right] <= x + k) {
                right++;
            }
            ans = Math.max(ans, Math.min(right - left, cnt + numOperations));
            cnt = 0;
        }

        if (ans >= numOperations) {
            return ans;
        }

        left = 0;
        for (right = 0; right < n; right++) {
            int x = nums[right];
            while (nums[left] < x - k * 2) {
                left++;
            }
            ans = Math.max(ans, right - left + 1);
        }
        return Math.min(ans, numOperations); // 最后和 numOperations 取最小值
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFrequency(vector<int>& nums, int k, int numOperations) {
        ranges::sort(nums);

        int n = nums.size();
        int ans = 0, cnt = 0, left = 0, right = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            cnt++;
            // 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
            if (i < n - 1 && x == nums[i + 1]) {
                continue;
            }
            while (nums[left] < x - k) {
                left++;
            }
            while (right < n && nums[right] <= x + k) {
                right++;
            }
            ans = max(ans, min(right - left, cnt + numOperations));
            cnt = 0;
        }

        if (ans >= numOperations) {
            return ans;
        }

        left = 0;
        for (int right = 0; right < n; right++) {
            int x = nums[right];
            while (nums[left] < x - k * 2) {
                left++;
            }
            ans = max(ans, right - left + 1);
        }
        return min(ans, numOperations); // 最后和 numOperations 取最小值
    }
};
```

```go [sol-Go]
func maxFrequency(nums []int, k, numOperations int) (ans int) {
	slices.Sort(nums)

	n := len(nums)
	var cnt, left, right int
	for i, x := range nums {
		cnt++
		// 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
		if i < n-1 && x == nums[i+1] {
			continue
		}
		for nums[left] < x-k {
			left++
		}
		for right < n && nums[right] <= x+k {
			right++
		}
		ans = max(ans, min(right-left, cnt+numOperations))
		cnt = 0
	}

	if ans >= numOperations {
		return ans
	}

	left = 0
	for right, x := range nums {
		for nums[left] < x-k*2 {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return min(ans, numOperations) // 最后和 numOperations 取最小值
}
```

也可以把两个 for 循环合起来。

```py [sol-Python3]
class Solution:
    def maxFrequency(self, nums: List[int], k: int, numOperations: int) -> int:
        nums.sort()

        n = len(nums)
        ans = cnt = left = right = left2 = 0
        for i, x in enumerate(nums):
            while nums[left2] < x - k * 2:
                left2 += 1
            ans = max(ans, min(i - left2 + 1, numOperations))

            cnt += 1
            # 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
            if i < n - 1 and x == nums[i + 1]:
                continue
            while nums[left] < x - k:
                left += 1
            while right < n and nums[right] <= x + k:
                right += 1
            ans = max(ans, min(right - left, cnt + numOperations))
            cnt = 0

        return ans
```

```java [sol-Java]
class Solution {
    public int maxFrequency(int[] nums, int k, int numOperations) {
        Arrays.sort(nums);

        int n = nums.length;
        int ans = 0;
        int cnt = 0;
        int left = 0;
        int right = 0;
        int left2 = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            while (nums[left2] < x - k * 2) {
                left2++;
            }
            ans = Math.max(ans, Math.min(i - left2 + 1, numOperations));

            cnt++;
            // 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
            if (i < n - 1 && x == nums[i + 1]) {
                continue;
            }
            while (nums[left] < x - k) {
                left++;
            }
            while (right < n && nums[right] <= x + k) {
                right++;
            }
            ans = Math.max(ans, Math.min(right - left, cnt + numOperations));
            cnt = 0;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFrequency(vector<int>& nums, int k, int numOperations) {
        ranges::sort(nums);

        int n = nums.size();
        int ans = 0, cnt = 0, left = 0, right = 0, left2 = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            while (nums[left2] < x - k * 2) {
                left2++;
            }
            ans = max(ans, min(i - left2 + 1, numOperations));

            cnt++;
            // 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
            if (i < n - 1 && x == nums[i + 1]) {
                continue;
            }
            while (nums[left] < x - k) {
                left++;
            }
            while (right < n && nums[right] <= x + k) {
                right++;
            }
            ans = max(ans, min(right - left, cnt + numOperations));
            cnt = 0;
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxFrequency(nums []int, k, numOperations int) (ans int) {
	slices.Sort(nums)

	n := len(nums)
	var cnt, left, right, left2 int
	for i, x := range nums {
		for nums[left2] < x-k*2 {
			left2++
		}
		ans = max(ans, min(i-left2+1, numOperations))

		cnt++
		// 循环直到连续相同段的末尾，这样可以统计出 x 的出现次数
		if i < n-1 && x == nums[i+1] {
			continue
		}
		for nums[left] < x-k {
			left++
		}
		for right < n && nums[right] <= x+k {
			right++
		}
		ans = max(ans, min(right-left, cnt+numOperations))
		cnt = 0
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

1. 下面数据结构题单的「**§2.1 一维差分**」。
2. 下面双指针题单的「**§3.2 同向双指针**」和「**五、三指针**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
