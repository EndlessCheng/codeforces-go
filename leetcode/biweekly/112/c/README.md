请看 [视频讲解](https://www.bilibili.com/video/BV1um4y1M7Rv/) 第三题。

看到「长度固定的子数组」就要想到滑动窗口！

维护窗口内的元素出现次数 $\textit{cnt}$，以及元素和 $\textit{sum}$。

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: List[int], m: int, k: int) -> int:
        ans = 0
        s = sum(nums[:k - 1])  # 先统计 k-1 个数
        cnt = Counter(nums[:k - 1])
        for out, in_ in zip(nums, nums[k - 1:]):
            s += in_  # 再添加一个数就是 k 个数了
            cnt[in_] += 1
            if len(cnt) >= m:
                ans = max(ans, s)
                
            s -= out  # 下一个子数组不包含 out，移出窗口
            cnt[out] -= 1
            if cnt[out] == 0:
                del cnt[out]
        return ans
```

```java [sol-Java]
class Solution {
    public long maxSum(List<Integer> nums, int m, int k) {
        var a = nums.stream().mapToInt(i -> i).toArray();
        long ans = 0, sum = 0;
        var cnt = new HashMap<Integer, Integer>();
        for (int i = 0; i < k - 1; i++) { // 先统计 k-1 个数
            sum += a[i];
            cnt.merge(a[i], 1, Integer::sum); // cnt[a[i]]++
        }
        for (int i = k - 1; i < nums.size(); i++) {
            sum += a[i]; // 再添加一个数就是 k 个数了
            cnt.merge(a[i], 1, Integer::sum); // cnt[a[i]]++
            if (cnt.size() >= m)
                ans = Math.max(ans, sum);

            int out = a[i - k + 1];
            sum -= out; // 下一个子数组不包含 out，移出窗口
            if (cnt.merge(out, -1, Integer::sum) == 0) // --cnt[out] == 0
                cnt.remove(out);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSum(vector<int> &nums, int m, int k) {
        long long ans = 0, sum = 0;
        unordered_map<int, int> cnt;
        for (int i = 0; i < k - 1; i++) { // 先统计 k-1 个数
            sum += nums[i];
            cnt[nums[i]]++;
        }
        for (int i = k - 1; i < nums.size(); i++) {
            sum += nums[i]; // 再添加一个数就是 k 个数了
            cnt[nums[i]]++;
            if (cnt.size() >= m)
                ans = max(ans, sum);

            int out = nums[i - k + 1];
            sum -= out; // 下一个子数组不包含 out，移出窗口
            if (--cnt[out] == 0)
                cnt.erase(out);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSum(nums []int, m, k int) (ans int64) {
	sum := int64(0)
	cnt := map[int]int{}
	for _, x := range nums[:k-1] { // 先统计 k-1 个数
		sum += int64(x)
		cnt[x]++
	}
	for i, in := range nums[k-1:] {
		sum += int64(in) // 再添加一个数就是 k 个数了
		cnt[in]++
		if len(cnt) >= m && sum > ans {
			ans = sum
		}

		out := nums[i]
		sum -= int64(out) // 下一个子数组不包含 out，移出窗口
		cnt[out]--
		if cnt[out] == 0 {
			delete(cnt, out)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。哈希表的大小不会超过窗口长度，即 $k$。

## 练习：滑动窗口

- [2461. 长度为 K 子数组中的最大和](https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/)
- [1343. 大小为 K 且平均值大于等于阈值的子数组数目](https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/)
- [2379. 得到 K 个黑块的最少涂色次数](https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/)
- [2653. 滑动子数组的美丽值](https://leetcode.cn/problems/sliding-subarray-beauty/)
- [995. K 连续位的最小翻转次数](https://leetcode.cn/problems/minimum-number-of-k-consecutive-bit-flips/)
