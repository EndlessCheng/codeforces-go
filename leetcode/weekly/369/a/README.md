请看 [视频讲解](https://www.bilibili.com/video/BV1tw411q7VZ/)。

枚举 $0$ 到 $30$ 的每个比特位 $i$（因为 $\textit{nums}[i]<2^{31}$）。

遍历数组，如果第 $i$ 个比特位上的 $1$ 的个数 $\ge k$，则把 $2^i$ 加到答案中。

关于位运算的技巧，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def findKOr(self, nums: List[int], k: int) -> int:
        ans = 0
        for i in range(31):
            cnt1 = sum(x >> i & 1 for x in nums)
            if cnt1 >= k:
                ans |= 1 << i
        return ans
```

```java [sol-Java]
class Solution {
    public int findKOr(int[] nums, int k) {
        int ans = 0;
        for (int i = 0; i < 31; i++) {
            int cnt1 = 0;
            for (int x : nums) {
                cnt1 += (x >> i) & 1;
            }
            if (cnt1 >= k) {
                ans |= 1 << i;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findKOr(vector<int>& nums, int k) {
        int ans = 0;
        for (int i = 0; i < 31; i++) {
            int cnt1 = 0;
            for (int x : nums) {
                cnt1 += (x >> i) & 1;
            }
            if (cnt1 >= k) {
                ans |= 1 << i;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findKOr(nums []int, k int) (ans int) {
	for i := 0; i < 31; i++ {
		cnt1 := 0
		for _, x := range nums {
			cnt1 += x >> i & 1
		}
		if cnt1 >= k {
			ans |= 1 << i
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。
