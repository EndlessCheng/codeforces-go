[视频讲解](https://www.bilibili.com/video/BV1Lm4y1N7mf/)

把所有满足下标的二进制中的 $1$ 的个数等于 $k$ 的 $\textit{nums}[i]$ 加起来，就是答案。

有关位运算的知识点，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def sumIndicesWithKSetBits(self, nums: List[int], k: int) -> int:
        return sum(x for i, x in enumerate(nums) if i.bit_count() == k)
```

```java [sol-Java]
class Solution {
    public int sumIndicesWithKSetBits(List<Integer> nums, int k) {
        int ans = 0, n = nums.size();
        for (int i = 0; i < n; i++) {
            if (Integer.bitCount(i) == k) {
                ans += nums.get(i);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumIndicesWithKSetBits(vector<int> &nums, int k) {
        int ans = 0, n = nums.size();
        for (int i = 0; i < n; i++) {
            if (__builtin_popcount(i) == k) {
                ans += nums[i];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumIndicesWithKSetBits(nums []int, k int) (ans int) {
	for i, x := range nums {
		if bits.OnesCount(uint(i)) == k {
			ans += x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。理由见视频。
- 空间复杂度：$\mathcal{O}(1)$。
