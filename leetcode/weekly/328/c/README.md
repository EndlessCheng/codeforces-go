[视频讲解](https://www.bilibili.com/video/BV1QT41127kJ/) 已出炉，欢迎点赞三连~

---

子数组统计问题，常用双指针（不定长滑动窗口）实现，具体原理可以看我的[【同向双指针+简洁模板】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)，看完你就掌握双指针啦~

本题的做法：

用一个哈希表 $\textit{cnt}$ 统计窗口内每个元素的出现次数。

枚举子数组右端点 $\textit{right}$，那么答案增加了 $\textit{cnt}[\textit{nums}[\textit{right}]]$；然后看左端点 $\textit{left}$ 最大可以是多少，如果去掉左端点，答案没有小于 $k$，就可以移动左端点。

由于左端点及其左边的都可以是好子数组的左端点，所以每个右端点对应的答案个数为 $\textit{left}+1$。

```py [sol1-Python3]
class Solution:
    def countGood(self, nums: List[int], k: int) -> int:
        cnt = Counter()
        ans = left = pairs = 0
        for x in nums:
            pairs += cnt[x]
            cnt[x] += 1  # 移入右端点
            while pairs - cnt[nums[left]] + 1 >= k:
                cnt[nums[left]] -= 1  # 移出左端点
                pairs -= cnt[nums[left]]
                left += 1
            if pairs >= k:
                ans += left + 1
        return ans
```

```java [sol1-Java]
class Solution {
    public long countGood(int[] nums, int k) {
        var cnt = new HashMap<Integer, Integer>();
        long ans = 0;
        int left = 0, pairs = 0;
        for (int x : nums) {
            pairs += cnt.getOrDefault(x, 0);
            cnt.merge(x, 1, Integer::sum); // 移入右端点
            while (pairs - cnt.get(nums[left]) + 1 >= k)
                pairs -= cnt.merge(nums[left++], -1, Integer::sum); // 移出左端点
            if (pairs >= k) ans += left + 1;
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long countGood(vector<int> &nums, int k) {
        unordered_map<int, int> cnt;
        long ans = 0;
        int left = 0, pairs = 0;
        for (int x : nums) {
            pairs += cnt[x]++; // 移入右端点
            while (pairs - cnt[nums[left]] + 1 >= k)
                pairs -= --cnt[nums[left++]]; // 移出左端点
            if (pairs >= k) ans += left + 1;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func countGood(nums []int, k int) (ans int64) {
	cnt := map[int]int{}
	left, pairs := 0, 0
	for _, x := range nums {
		pairs += cnt[x]
		cnt[x]++ // 移入右端点
		for pairs-cnt[nums[left]]+1 >= k {
			cnt[nums[left]]-- // 移出左端点
			pairs -= cnt[nums[left]]
			left++
		}
		if pairs >= k {
			ans += int64(left + 1)
		}
	}
	return
}
````

还可以把更新 $\textit{ans}$ 的逻辑写在循环内部。

你更喜欢哪种写法呢？

```py [sol2-Python3]
class Solution:
    def countGood(self, nums: List[int], k: int) -> int:
        cnt = Counter()
        ans = left = pairs = 0
        for x in nums:
            pairs += cnt[x]
            cnt[x] += 1  # 移入右端点
            ans += left
            while pairs >= k:
                ans += 1
                cnt[nums[left]] -= 1  # 移出左端点
                pairs -= cnt[nums[left]]
                left += 1
        return ans
```

```java [sol2-Java]
class Solution {
    public long countGood(int[] nums, int k) {
        var cnt = new HashMap<Integer, Integer>();
        long ans = 0;
        int left = 0, pairs = 0;
        for (int x : nums) {
            pairs += cnt.getOrDefault(x, 0);
            cnt.merge(x, 1, Integer::sum); // 移入右端点
            ans += left;
            while (pairs >= k) {
                pairs -= cnt.merge(nums[left++], -1, Integer::sum); // 移出左端点
                ++ans;
            }
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    long long countGood(vector<int> &nums, int k) {
        unordered_map<int, int> cnt;
        long ans = 0;
        int left = 0, pairs = 0;
        for (int x : nums) {
            pairs += cnt[x]++; // 移入右端点
            ans += left;
            while (pairs >= k) {
                pairs -= --cnt[nums[left++]]; // 移出左端点
                ++ans;
            }
        }
        return ans;
    }
};
```

```go [sol2-Go]
func countGood(nums []int, k int) (ans int64) {
	cnt := map[int]int{}
	left, pairs := 0, 0
	for _, x := range nums {
		pairs += cnt[x]
		cnt[x]++ // 移入右端点
		ans += int64(left)
		for pairs >= k {
			ans++
			cnt[nums[left]]-- // 移出左端点
			pairs -= cnt[nums[left]]
			left++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

#### 相似题目（同向双指针）

- [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)，[题解](https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/1959540/xia-biao-zong-suan-cuo-qing-kan-zhe-by-e-iaks/)
- [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)，[题解](https://leetcode.cn/problems/minimum-size-subarray-sum/solutions/1959532/biao-ti-xia-biao-zong-suan-cuo-qing-kan-k81nh/)
- [713. 乘积小于 K 的子数组](https://leetcode.cn/problems/subarray-product-less-than-k/)，[题解](https://leetcode.cn/problems/subarray-product-less-than-k/solutions/1959538/xia-biao-zong-suan-cuo-qing-kan-zhe-by-e-jebq/)
