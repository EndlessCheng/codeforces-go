根据题意，如果 $\textit{nums}[i]$ 的出现次数超过 $2$，则无法分割，否则可以分割。

```py [sol-Python3]
class Solution:
    def isPossibleToSplit(self, nums: List[int]) -> bool:
        return Counter(nums).most_common(1)[0][1] <= 2
```

```java [sol-Java]
class Solution {
    public boolean isPossibleToSplit(int[] nums) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            if (cnt.merge(x, 1, Integer::sum) > 2) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isPossibleToSplit(vector<int> &nums) {
        unordered_map<int, int> cnt;
        for (int x : nums) {
            if (++cnt[x] > 2) {
                return false;
            }
        }
        return true;
    }
};
```

```go [sol-Go]
func isPossibleToSplit(nums []int) bool {
	cnt := map[int]int{}
	for _, x := range nums {
		if cnt[x] == 2 {
			return false
		}
		cnt[x]++
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
