![](https://pic.leetcode.cn/1696131371-UYYmoV-w365c-c.png)

关于滑动窗口的原理，请看视频：[滑动窗口【基础算法精讲 03】](https://b23.tv/pRbxHhG)

## 答疑

**问**：剩余元素的个数会不会大于等于 $n$？

**答**：不会，如果大于等于 $n$，那么数组中的每个数至少出现一次，这意味着剩余元素之和至少为 $\textit{total}$，这与 $\textit{target}\bmod \textit{total} < \textit{total}$ 相矛盾。这也解释了为什么只需要在 $\textit{nums}+\textit{nums}$ 中滑窗，而不需要在 $\textit{nums}+\textit{nums}+\textit{nums}$ 这样更长的数组中滑窗。

```py [sol-Python3]
class Solution:
    def minSizeSubarray(self, nums: List[int], target: int) -> int:
        total = sum(nums)
        n = len(nums)
        ans = inf
        left = s = 0
        for right in range(n * 2):
            s += nums[right % n]
            while s > target % total:
                s -= nums[left % n]
                left += 1
            if s == target % total:
                ans = min(ans, right - left + 1)
        return ans + target // total * n if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minSizeSubarray(int[] nums, int target) {
        long total = 0;
        for (int x : nums) total += x;
        int n = nums.length;
        int ans = Integer.MAX_VALUE;
        int left = 0;
        long sum = 0;
        for (int right = 0; right < n * 2; right++) {
            sum += nums[right % n];
            while (sum > target % total) {
                sum -= nums[left++ % n];
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
    int minSizeSubarray(vector<int> &nums, int target) {
        long long total = accumulate(nums.begin(), nums.end(), 0LL);
        int n = nums.size();
        int ans = INT_MAX;
        int left = 0;
        long long sum = 0;
        for (int right = 0; right < n * 2; right++) {
            sum += nums[right % n];
            while (sum > target % total) {
                sum -= nums[left++ % n];
            }
            if (sum == target % total) {
                ans = min(ans, right - left + 1);
            }
        }
        return ans == INT_MAX ? -1 : ans + target / total * n;
    }
};
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

func min(a, b int) int { if b < a { return b }; return a }
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
            sum -= nums[left++ % n];
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
        let total: i64 = nums.iter().map(|&x| x as i64).sum();
        let n = nums.len();
        let mut ans = usize::MAX;
        let mut left = 0;
        let mut sum = 0;
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
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 题单（右边的数字是题目难度）

#### 定长滑动窗口

- [1456. 定长子串中元音的最大数目](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/) 1263
- [2269. 找到一个数字的 K 美丽值](https://leetcode.cn/problems/find-the-k-beauty-of-a-number/) 1280
- [1984. 学生分数的最小差值](https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/) 1306
- [643. 子数组最大平均数 I](https://leetcode.cn/problems/maximum-average-subarray-i/)
- [1343. 大小为 K 且平均值大于等于阈值的子数组数目](https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/) 1317
- [2090. 半径为 k 的子数组平均值](https://leetcode.cn/problems/k-radius-subarray-averages/) 1358
- [2379. 得到 K 个黑块的最少涂色次数](https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/) 1360
- [1052. 爱生气的书店老板](https://leetcode.cn/problems/grumpy-bookstore-owner/) 1418
- [2841. 几乎唯一子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/) 1546
- [2461. 长度为 K 子数组中的最大和](https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/) 1553
- [1423. 可获得的最大点数](https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/) 1574
- [2134. 最少交换次数来组合所有的 1 II](https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together-ii/) 1748
- [2653. 滑动子数组的美丽值](https://leetcode.cn/problems/sliding-subarray-beauty/) 1786
- [567. 字符串的排列](https://leetcode.cn/problems/permutation-in-string/)
- [438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)
- [2156. 查找给定哈希值的子串](https://leetcode.cn/problems/find-substring-with-given-hash-value/) 2063
- [346. 数据流中的移动平均值](https://leetcode.cn/problems/moving-average-from-data-stream/)（会员题）
- [1100. 长度为 K 的无重复字符子串](https://leetcode.cn/problems/find-k-length-substrings-with-no-repeated-characters/)（会员题）

#### 不定长滑动窗口（求最长/最大）

- [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
- [1493. 删掉一个元素以后全为 1 的最长子数组](https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/) 1423
- [904. 水果成篮](https://leetcode.cn/problems/fruit-into-baskets/) 1516
- [1695. 删除子数组的最大得分](https://leetcode.cn/problems/maximum-erasure-value/) 1529
- [2841. 几乎唯一子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/) 1546
- [2024. 考试的最大困扰度](https://leetcode.cn/problems/maximize-the-confusion-of-an-exam/) 1643
- [1004. 最大连续1的个数 III](https://leetcode.cn/problems/max-consecutive-ones-iii/) 1656
- [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/) 1672
- [2401. 最长优雅子数组](https://leetcode.cn/problems/longest-nice-subarray/) 1750
- [1658. 将 x 减到 0 的最小操作数](https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/) 1817
- [1838. 最高频元素的频数](https://leetcode.cn/problems/frequency-of-the-most-frequent-element/) 1876
- [2831. 找出最长等值子数组](https://leetcode.cn/problems/find-the-longest-equal-subarray/) 1976
- [2106. 摘水果](https://leetcode.cn/problems/maximum-fruits-harvested-after-at-most-k-steps/) 2062
- [1610. 可见点的最大数目](https://leetcode.cn/problems/maximum-number-of-visible-points/) 2147
- [159. 至多包含两个不同字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters/)（会员题）
- [340. 至多包含 K 个不同字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters/)（会员题）

#### 不定长滑动窗口（求最短/最小）

- [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)
- [1234. 替换子串得到平衡字符串](https://leetcode.cn/problems/replace-the-substring-for-balanced-string/) 1878
- [1574. 删除最短的子数组使剩余数组有序](https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/) 1932
- [76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/)

#### 不定长滑动窗口（求子数组个数）

- [2799. 统计完全子数组的数目](https://leetcode.cn/problems/count-complete-subarrays-in-an-array/) 1398
- [713. 乘积小于 K 的子数组](https://leetcode.cn/problems/subarray-product-less-than-k/)
- [1358. 包含所有三种字符的子字符串数目](https://leetcode.cn/problems/number-of-substrings-containing-all-three-characters/) 1646
- [2302. 统计得分小于 K 的子数组数目](https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/) 1808
- [2537. 统计好子数组的数目](https://leetcode.cn/problems/count-the-number-of-good-subarrays/) 1892
- [2762. 不间断子数组](https://leetcode.cn/problems/continuous-subarrays/) 1940

#### 多指针滑动窗口

- [930. 和相同的二元子数组](https://leetcode.cn/problems/binary-subarrays-with-sum/) 1592
- [1248. 统计「优美子数组」](https://leetcode.cn/problems/count-number-of-nice-subarrays/) 1624
- [1712. 将数组分成三个子数组的方案数](https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/) 2079
- [2444. 统计定界子数组的数目](https://leetcode.cn/problems/count-subarrays-with-fixed-bounds/) 2093
- [992. K 个不同整数的子数组](https://leetcode.cn/problems/subarrays-with-k-different-integers/) 2210
