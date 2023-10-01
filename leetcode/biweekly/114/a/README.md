倒序遍历。由于元素范围在 $[1,50]$，我们可以用一个 $64$ 位整数表示集合，只要集合中有 $1$ 到 $k$ 就立刻返回。

原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        u = (2 << k) - 2  # 1~k
        s, n = 0, len(nums)
        for i in range(n - 1, -1, -1):
            s |= 1 << nums[i]
            if (s & u) == u:
                return n - i
```

```java [sol-Java]
class Solution {
    public int minOperations(List<Integer> nums, int k) {
        int n = nums.size();
        long u = (2L << k) - 2; // 1~k
        long s = 0;
        for (int i = n - 1; ; --i) {
            s |= 1L << nums.get(i);
            if ((s & u) == u) {
                return n - i;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(std::vector<int>& nums, int k) {
        int n = nums.size();
        long long u = (2LL << k) - 2; // 1~k
        long long s = 0;
        for (int i = n - 1; ; --i) {
            s |= 1LL << nums[i];
            if ((s & u) == u) {
                return n - i;
            }
        }
    }
};
```

```go [sol-Go]
func minOperations(nums []int, k int) int {
	all := 2<<k - 2 // 1~k
	set := 0
	for i := len(nums) - 1; ; i-- {
		set |= 1 << nums[i]
		if set&all == all {
			return len(nums) - i
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
