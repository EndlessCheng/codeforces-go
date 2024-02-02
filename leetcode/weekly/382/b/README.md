[视频讲解](https://www.bilibili.com/video/BV1we411J7Y8/) 第二题。

设数组的最大值为 $m$，当 $x\ge 2$ 时，我们有

$$
x^{2^p} \le m
$$

解得

$$
p \le \log_2 \log_x m
$$

在本题数据范围下，$p$ 至多为 $4$。

所以暴力枚举数组中的数，作为 $x$，然后不断看 $x^2,x^4,\cdots$ 在数组中的个数。直到个数不足 $2$ 个为止，退出循环。

注意模式的正中间的数字只取一个。如果最后 $x$ 有一个，那么个数加一，否则个数减一。

注意特判 $x=1$ 的情况。

```py [sol-Python3]
class Solution:
    def maximumLength(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        ans = cnt[1] - 1 | 1  # 奇数
        del cnt[1]
        for x in cnt:
            res = 0
            while cnt[x] > 1:
                res += 2
                x *= x
            ans = max(ans, res + (1 if x in cnt else -1))  # 保证 res 是奇数
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumLength(int[] nums) {
        HashMap<Long, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            cnt.merge((long) x, 1, Integer::sum);
        }
        Integer c1 = cnt.remove(1L);
        int ans = c1 != null ? c1 - 1 | 1 : 0;
        for (long x : cnt.keySet()) {
            int res = 0;
            for (; cnt.getOrDefault(x, 0) > 1; x *= x) {
                res += 2;
            }
            ans = Math.max(ans, res + (cnt.containsKey(x) ? 1 : -1)); // 保证 res 是奇数
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumLength(vector<int> &nums) {
        unordered_map<long long, int> cnt;
        for (int x : nums) {
            cnt[x]++;
        }
        int ans = cnt[1] - 1 | 1; // 奇数
        cnt.erase(1);
        for (auto &[num, _] : cnt) {
            int res = 0;
            long long x = num;
            for (; cnt.contains(x) && cnt[x] > 1; x *= x) {
                res += 2; 
            }
            ans = max(ans, res + (cnt.contains(x) ? 1 : -1)); // 保证 res 是奇数
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumLength(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	ans := cnt[1] - 1 | 1 // 奇数
	delete(cnt, 1)
	for x := range cnt {
		res := 0
		for ; cnt[x] > 1; x *= x {
			res += 2
		}
		res += cnt[x]
		ans = max(ans, res-1|1) // 保证 res 是奇数
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log \log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
