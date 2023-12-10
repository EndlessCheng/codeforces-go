请先看 [没有思路？一个视频讲透滑动窗口！【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

本题算法如下：

1. 设 $\textit{mx} = \max(\textit{nums})$。
2. 右端点 $\textit{right}$ 从左到右遍历 $\textit{nums}$。遍历到元素 $x=\textit{nums}[\textit{right}]$ 时，如果 $x=\textit{mx}$，就把计数器 $\textit{cntMx}$ 加一。
3. 如果此时 $\textit{cntMx}=k$，则不断右移左指针 $\textit{left}$，直到窗口内的 $\textit{mx}$ 的出现次数**小于** $k$ 为止。此时，对于右端点为 $\textit{right}$ 且左端点小于 $\textit{left}$ 的子数组，$\textit{mx}$ 的出现次数都至少为 $k$，把答案增加 $\textit{left}$。

例如示例 1，当右端点移到第二个 $3$ 时，左端点移到 $2$，此时 $[1,3,2,3]$ 和 $[3,2,3]$ 是满足要求的。当右端点移到第三个 $3$ 时，左端点也移到第三个 $3$，此时 $[1,3,2,3,3], [3,2,3,3], [2,3,3], [3,3]$ 都是满足要求的。所以答案为 $2+4=6$。

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        mx = max(nums)
        ans = cnt_mx = left = 0
        for x in nums:
            if x == mx:
                cnt_mx += 1
            while cnt_mx == k:
                if nums[left] == mx:
                    cnt_mx -= 1
                left += 1
            ans += left
        return ans
```

```java [sol-Java]
public class Solution {
    public long countSubarrays(int[] nums, int k) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        long ans = 0;
        int cntMx = 0, left = 0;
        for (int x : nums) {
            if (x == mx) {
                cntMx++;
            }
            while (cntMx == k) {
                if (nums[left++] == mx) {
                    cntMx--;
                }
            }
            ans += left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubarrays(vector<int> &nums, int k) {
        int mx = *max_element(nums.begin(), nums.end());
        long long ans = 0;
        int cnt_mx = 0, left = 0;
        for (int x : nums) {
            cnt_mx += x == mx;
            while (cnt_mx == k) {
                cnt_mx -= nums[left++] == mx;
            }
            ans += left;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubarrays(nums []int, k int) (ans int64) {
	mx := slices.Max(nums)
	cntMx, left := 0, 0
	for _, x := range nums {
		if x == mx {
			cntMx++
		}
		for cntMx == k {
			if nums[left] == mx {
				cntMx--
			}
			left++
		}
		ans += int64(left)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

把问题改成：子数组的最大值至少出现 $k$ 次，要怎么做？

用**单调栈**做，思路可以参考 [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/)

## 题单

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
- [395. 至少有 K 个重复字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/)
- [1763. 最长的美好子字符串](https://leetcode.cn/problems/longest-nice-substring/)
- [2953. 统计完全子字符串](https://leetcode.cn/problems/count-complete-substrings/)
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
- [2743. 计算没有重复字符的子字符串数量](https://leetcode.cn/problems/count-substrings-without-repeating-character/)（会员题）

#### 多指针滑动窗口

- [930. 和相同的二元子数组](https://leetcode.cn/problems/binary-subarrays-with-sum/) 1592
- [1248. 统计「优美子数组」](https://leetcode.cn/problems/count-number-of-nice-subarrays/) 1624
- [2563. 统计公平数对的数目](https://leetcode.cn/problems/count-the-number-of-fair-pairs/) 1721
- [1712. 将数组分成三个子数组的方案数](https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/) 2079
- [2444. 统计定界子数组的数目](https://leetcode.cn/problems/count-subarrays-with-fixed-bounds/) 2093
- [992. K 个不同整数的子数组](https://leetcode.cn/problems/subarrays-with-k-different-integers/) 2210
