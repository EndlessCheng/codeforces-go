[本题视频讲解](https://www.bilibili.com/video/BV1jg4y1y7PA/)

把数组排序，枚举 $\textit{nums}[i]$ 作为最长的那条边，那么 $\textit{nums}[i]$ 左边的数之和越大越好，这样才能尽可能地组成多边形，并且周长尽量长。

所以周长就是 $\textit{nums}$ 的某个前缀和。从大到小枚举 $\textit{nums}[i]$，如果满足

$$
\textit{nums}[0] + \textit{nums}[1] + \cdots + \textit{nums}[i-1] > \textit{nums}[i]
$$

那么答案就是

$$
\textit{nums}[0] + \textit{nums}[1] + \cdots + \textit{nums}[i-1] + \textit{nums}[i]
$$

```py [sol-Python3]
class Solution:
    def largestPerimeter(self, nums: List[int]) -> int:
        nums.sort()
        s = sum(nums)
        for i in range(len(nums) - 1, 1, -1):
            x = nums[i]
            if s > x * 2:  # s-x > x
                return s
            s -= x
        return -1
```

```java [sol-Java]
class Solution {
    public long largestPerimeter(int[] nums) {
        Arrays.sort(nums);
        long s = 0;
        for (int x : nums) {
            s += x;
        }
        for (int i = nums.length - 1; i > 1; i--) {
            int x = nums[i];
            if (s > x * 2) { // s-x > x
                return s;
            }
            s -= x;
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long largestPerimeter(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        long long s = accumulate(nums.begin(), nums.end(), 0LL);
        for (int i = nums.size() - 1; i > 1; --i) {
            int x = nums[i];
            if (s > x * 2) { // s-x > x
                return s;
            }
            s -= x;
        }
        return -1;
    }
};
```

```go [sol-Go]
func largestPerimeter(nums []int) int64 {
	slices.Sort(nums)
	s := 0
	for _, x := range nums {
		s += x
	}
	for i := len(nums) - 1; i > 1; i-- {
		x := nums[i]
		if s > x*2 { // s-x > x
			return int64(s)
		}
		s -= x
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。不计入排序的栈开销。
