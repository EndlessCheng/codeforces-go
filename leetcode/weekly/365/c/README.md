![lc2875-c.png](https://pic.leetcode.cn/1745215015-YlWiNO-lc2875-c.png)

关于滑动窗口的原理，请看视频 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

## 答疑

**问**：去掉中间 $k$ 个 $\textit{nums}$ 后，剩余元素的个数是否会 $\ge n$？

**答**：不会。如果 $\ge n$，说明剩余元素包含了 $\textit{nums}$ 中的每个数，所以剩余元素之和 $\textit{rem}\ge \textit{total}$，这与 $\textit{rem}=\textit{target}\bmod \textit{total} < \textit{total}$ 相矛盾。这也解释了为什么只需要在 $\textit{nums}+\textit{nums}$ 中滑窗，而不需要在 $\textit{nums}+\textit{nums}+\textit{nums}$ 这样更长的数组中滑窗。

**问**：是否需要特判 $\textit{rem} = 0$ 这种情况？

**答**：无需特判。如果 $\textit{rem} = 0$，那么和为 $0$ 的最短子数组，就是空数组，长度为 $0$。此时答案为 $\left\lfloor\dfrac{target}{total}\right\rfloor\cdot n$。

```py [sol-Python3]
class Solution:
    def minSizeSubarray(self, nums: List[int], target: int) -> int:
        total = sum(nums)
        n = len(nums)
        ans = inf
        left = s = 0
        rem = target % total
        for right in range(n * 2):
            s += nums[right % n]
            while s > rem:
                s -= nums[left % n]
                left += 1
            if s == rem:
                ans = min(ans, right - left + 1)
        return ans + target // total * n if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minSizeSubarray(int[] nums, int target) {
        long total = 0;
        for (int x : nums) {
            total += x;
        }

        int n = nums.length;
        int ans = Integer.MAX_VALUE;
        long sum = 0;
        int left = 0;
        for (int right = 0; right < n * 2; right++) {
            sum += nums[right % n];
            while (sum > target % total) {
                sum -= nums[left % n];
                left++;
            }
            if (sum == target % total) {
                ans = Math.min(ans, right - left + 1);
            }
        }

        return ans == Integer.MAX_VALUE ? -1 : ans + (int) (target / total) * n;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minSizeSubarray(vector<int>& nums, int target) {
        long long total = reduce(nums.begin(), nums.end(), 0LL);
        int n = nums.size();
        int ans = INT_MAX;
        long long sum = 0;
        int left = 0;
        for (int right = 0; right < n * 2; right++) {
            sum += nums[right % n];
            while (sum > target % total) {
                sum -= nums[left % n];
                left++;
            }
            if (sum == target % total) {
                ans = min(ans, right - left + 1);
            }
        }
        return ans == INT_MAX ? -1 : ans + target / total * n;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int minSizeSubarray(int* nums, int numsSize, int target) {
    long long total = 0;
    for (int i = 0; i < numsSize; i++) {
        total += nums[i];
    }

    int n = numsSize;
    int ans = INT_MAX;
    long long sum = 0;
    int left = 0;
    for (int right = 0; right < n * 2; right++) {
        sum += nums[right % n];
        while (sum > target % total) {
            sum -= nums[left % n];
            left++;
        }
        if (sum == target % total) {
            ans = MIN(ans, right - left + 1);
        }
    }

    return ans == INT_MAX ? -1 : ans + target / total * n;
}
```

```go [sol-Go]
func minSizeSubarray(nums []int, target int) int {
	total := 0
	for _, x := range nums {
		total += x
	}

	ans := math.MaxInt
	left, sum, n := 0, 0, len(nums)
	for right := 0; right < n*2; right++ {
		sum += nums[right%n]
		for sum > target%total {
			sum -= nums[left%n]
			left++
		}
		if sum == target%total {
			ans = min(ans, right-left+1)
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans + target/total*n
}
```

```js [sol-JavaScript]
var minSizeSubarray = function (nums, target) {
    const total = _.sum(nums);
    const n = nums.length;
    let ans = Infinity;
    let left = 0, sum = 0;
    for (let right = 0; right < n * 2; right++) {
        sum += nums[right % n];
        while (sum > target % total) {
            sum -= nums[left % n];
            left++;
        }
        if (sum === target % total) {
            ans = Math.min(ans, right - left + 1);
        }
    }
    return ans === Infinity ? -1 : ans + Math.floor(target / total) * n;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_size_subarray(nums: Vec<i32>, target: i32) -> i32 {
        let target = target as i64;
        let total = nums.iter().map(|&x| x as i64).sum::<i64>();
        let n = nums.len();
        let mut ans = usize::MAX;
        let mut sum = 0;
        let mut left = 0;
        for right in 0..n * 2 {
            sum += nums[right % n];
            while sum > (target % total) as i32 {
                sum -= nums[left % n];
                left += 1;
            }
            if sum == (target % total) as i32 {
                ans = ans.min(right - left + 1);
            }
        }
        if ans < usize::MAX {
            ans as i32 + (target / total) as i32 * n as i32
        } else {
            -1
        }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
