### 本题视频讲解

见[【双周赛 101】](https://www.bilibili.com/video/BV1Ga4y1M72A/)。

包括位运算的原理介绍。

### 思路

如果 $\textit{nums}_1$ 与 $\textit{nums}_2$ 有交集，那么答案就是交集的最小值。

否则取两个数组的最小值，记作 $x$ 和 $y$，那么答案就是 $\min(10x+y, 10y+x)$。

哈希表做法：

```py
class Solution:
    def minNumber(self, nums1: List[int], nums2: List[int]) -> int:
        s1, s2 = set(nums1), set(nums2)
        s = s1 & s2
        if s: return min(s)
        x, y = min(nums1), min(nums2)
        return min(x * 10 + y, y * 10 + x)
```

位运算做法：

```java [sol1-Java]
class Solution {
    public int minNumber(int[] nums1, int[] nums2) {
        int mask1 = 0, mask2 = 0;
        for (int x : nums1) mask1 |= 1 << x;
        for (int x : nums2) mask2 |= 1 << x;
        int m = mask1 & mask2;
        if (m > 0) return Integer.numberOfTrailingZeros(m);
        int x = Integer.numberOfTrailingZeros(mask1);
        int y = Integer.numberOfTrailingZeros(mask2);
        return Math.min(x * 10 + y, y * 10 + x);
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minNumber(vector<int> &nums1, vector<int> &nums2) {
        int mask1 = 0, mask2 = 0;
        for (int x: nums1) mask1 |= 1 << x;
        for (int x: nums2) mask2 |= 1 << x;
        int m = mask1 & mask2;
        if (m) return __builtin_ctz(m);
        int x = __builtin_ctz(mask1), y = __builtin_ctz(mask2);
        return min(x * 10 + y, y * 10 + x);
    }
};
```

```go [sol1-Go]
func minNumber(nums1, nums2 []int) int {
	var mask1, mask2 uint
	for _, x := range nums1 { mask1 |= 1 << x }
	for _, x := range nums2 { mask2 |= 1 << x }
	if m := mask1 & mask2; m > 0 {
		return bits.TrailingZeros(m)
	}
	x, y := bits.TrailingZeros(mask1), bits.TrailingZeros(mask2)
	return min(x*10+y, y*10+x)
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
