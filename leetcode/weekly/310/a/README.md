### 视频讲解

见[【周赛 310】](https://www.bilibili.com/video/BV1it4y1L7kL)

### 思路

用哈希表 $\textit{cnt}$ 统计每个偶数的出现次数，在统计的过程中，维护出现次数最大的最小元素。

如果没有偶数，就返回 $-1$。

```py [sol1-Python3]
class Solution:
    def mostFrequentEven(self, nums: List[int]) -> int:
        ans = -1
        cnt = Counter()
        for x in nums:
            if x % 2: continue  # 跳过奇数
            cnt[x] += 1
            if cnt[x] > cnt[ans] or cnt[x] == cnt[ans] and x < ans:
                ans = x  # 出现次数最大的数中，值最小的
        return ans
```

```py [sol1-Python3 多次遍历]
class Solution:
    def mostFrequentEven(self, nums: List[int]) -> int:
        cnt = Counter(x for x in nums if x % 2 == 0)  # 统计每个偶数的出现次数
        if len(cnt) == 0: return -1  # 没有偶数
        max_cnt = max(cnt.values())  # 最大出现次数
        return min(x for x, c in cnt.items() if c == max_cnt)  # 等于最大出现次数的最小值
```

```java [sol1-Java]
class Solution {
    public int mostFrequentEven(int[] nums) {
        int ans = -1;
        var cnt = new HashMap<Integer, Integer>();
        for (var x : nums) {
            if (x % 2 > 0) continue; // 跳过奇数
            int c = cnt.merge(x, 1, Integer::sum); // ++cnt[x]
            if (ans < 0 || c > cnt.get(ans) || c == cnt.get(ans) && x < ans)
                ans = x; // 出现次数最大的数中，值最小的
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int mostFrequentEven(vector<int> &nums) {
        int ans = -1;
        unordered_map<int, int> cnt;
        for (int x: nums) {
            if (x % 2) continue; // 跳过奇数
            int c = ++cnt[x];
            if (c > cnt[ans] || c == cnt[ans] && x < ans)
                ans = x; // 出现次数最大的数中，值最小的
        }
        return ans;
    }
};
```

```go [sol1-Go]
func mostFrequentEven(nums []int) int {
	ans := -1
	cnt := map[int]int{}
	for _, x := range nums {
		if x%2 == 0 { // 统计偶数
			cnt[x]++
			if cnt[x] > cnt[ans] || cnt[x] == cnt[ans] && x < ans {
				ans = x // 出现次数最大的数中，值最小的
			}
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

---

欢迎关注[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)，高质量算法教学，持续更新中~

附：[每日一题·高质量题解精选](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
