由于所有小于 $k$ 的数都要被删除，所以统计 $\textit{nums}$ 中的小于 $k$ 的元素个数即可。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        return sum(x < k for x in nums)
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int ans = 0;
        for (int x : nums) {
            if (x < k) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int> &nums, int k) {
        int ans = 0;
        for (int x : nums) {
            ans += x < k;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int, k int) (ans int) {
	for _, x := range nums {
		if x < k {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
