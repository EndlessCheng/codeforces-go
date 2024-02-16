[视频讲解](https://www.bilibili.com/video/BV1zt4y1R7Tc/)

遍历 $\textit{nums}$，同时用哈希表统计每个元素的出现次数，并维护出现次数的最大值 $\textit{maxCnt}$：

- 如果出现次数 $c > \textit{maxCnt}$，那么更新 $\textit{maxCnt}=c$，答案 $\textit{ans} = c$。
- 如果出现次数 $c = \textit{maxCnt}$，那么答案增加 $c$。

```py [sol-Python3]
class Solution:
    def maxFrequencyElements(self, nums: List[int]) -> int:
        ans = max_cnt = 0
        cnt = Counter()
        for x in nums:
            cnt[x] += 1
            c = cnt[x]
            if c > max_cnt:
                max_cnt = ans = c
            elif c == max_cnt:
                ans += c
        return ans
```

```java [sol-Java]
class Solution {
    public int maxFrequencyElements(int[] nums) {
        int ans = 0, maxCnt = 0;
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            int c = cnt.merge(x, 1, Integer::sum);
            if (c > maxCnt) {
                maxCnt = ans = c;
            } else if (c == maxCnt) {
                ans += c;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxFrequencyElements(vector<int> &nums) {
        int ans = 0, maxCnt = 0;
        unordered_map<int, int> cnt;
        for (int x : nums) {
            int c = ++cnt[x];
            if (c > maxCnt) {
                maxCnt = ans = c;
            } else if (c == maxCnt) {
                ans += c;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxFrequencyElements(nums []int) (ans int) {
	maxCnt := 0
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
		c := cnt[x]
		if c > maxCnt {
			maxCnt = c
			ans = c
		} else if c == maxCnt {
			ans += c
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
