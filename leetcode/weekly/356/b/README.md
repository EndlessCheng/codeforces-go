下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

经典滑窗，视频讲解可以看[【基础算法精讲 01】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

```py [sol-Python3]
class Solution:
    def countCompleteSubarrays(self, nums: List[int]) -> int:
        m = len(set(nums))
        cnt = Counter()
        ans = left = 0
        for v in nums:  # 枚举子数组右端点 v=nums[i]
            cnt[v] += 1
            while len(cnt) == m:
                x = nums[left]
                cnt[x] -= 1
                if cnt[x] == 0:
                    del cnt[x]
                left += 1
            ans += left  # 子数组左端点 < left 的都是合法的
        return ans
```

```java [sol-Java]
class Solution {
    public int countCompleteSubarrays(int[] nums) {
        var set = new HashSet<Integer>();
        for (int x : nums) set.add(x);
        int m = set.size();
        var cnt = new HashMap<Integer, Integer>();
        int ans = 0, left = 0;
        for (int v : nums) { // 枚举子数组右端点 v=nums[i]
            cnt.merge(v, 1, Integer::sum);
            while (cnt.size() == m) {
                int x = nums[left++];
                if (cnt.merge(x, -1, Integer::sum) == 0)
                    cnt.remove(x);
            }
            ans += left; // 子数组左端点 < left 的都是合法的
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countCompleteSubarrays(vector<int> &nums) {
        int m = unordered_set<int>(nums.begin(), nums.end()).size();
        unordered_map<int, int> cnt;
        int ans = 0, left = 0;
        for (int v: nums) { // 枚举子数组右端点 v=nums[i]
            cnt[v]++;
            while (cnt.size() == m) {
                int x = nums[left++];
                if (--cnt[x] == 0)
                    cnt.erase(x);
            }
            ans += left; // 子数组左端点 < left 的都是合法的
        }
        return ans;
    }
};
```

```go [sol-Go]
func countCompleteSubarrays(nums []int) (ans int) {
	set := map[int]struct{}{}
	for _, v := range nums {
		set[v] = struct{}{}
	}
	m := len(set)

	cnt := map[int]int{}
	left := 0
	for _, v := range nums { // 枚举子数组右端点 v=nums[i]
		cnt[v]++
		for len(cnt) == m {
			x := nums[left]
			cnt[x]--
			if cnt[x] == 0 {
				delete(cnt, x)
			}
			left++
		}
		ans += left // 子数组左端点 < left 的都是合法的
	}
	return
}
```

```js [sol-JavaScript]
var countCompleteSubarrays = function (nums) {
    const m = new Set(nums).size;
    let cnt = new Map();
    let ans = 0, left = 0;
    for (const v of nums) { // 枚举子数组右端点 v=nums[i]
        cnt.set(v, (cnt.get(v) ?? 0) + 1);
        while (cnt.size === m) {
            const x = nums[left++];
            cnt.set(x, cnt.get(x) - 1);
            if (cnt.get(x) === 0)
                cnt.delete(x);
        }
        ans += left; // 子数组左端点 < left 的都是合法的
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

#### 相似题目

- [992. K 个不同整数的子数组](https://leetcode.cn/problems/subarrays-with-k-different-integers/)
