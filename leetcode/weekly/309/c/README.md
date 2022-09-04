下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

由于所有元素对按位与均为 $0$，在优雅子数组中的从低到高的第 $i$ 个比特位上，至多有一个比特 $1$，其余均为比特 $0$。

因此在本题数据范围下，优雅子数组的长度不会超过 30。

暴力枚举即可。

代码实现时可以把在优雅子数组中的元素**按位或**起来，这样可以 $O(1)$ 判断当前元素是否与前面的元素按位与的结果为 $0$。

#### 复杂度分析

- 时间复杂度：$O(n\log\max(\textit{nums}))$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。

```py [sol1-Python3]
class Solution:
    def longestNiceSubarray(self, nums: List[int]) -> int:
        ans = 0
        for i, or_ in enumerate(nums):
            j = i - 1
            while j >= 0 and (or_ & nums[j]) == 0:
                or_ |= nums[j]
                j -= 1
            ans = max(ans, i - j)
        return ans
```

```java [sol1-Java]
class Solution {
    public int longestNiceSubarray(int[] nums) {
        int ans = 0;
        for (int i = 0; i < nums.length; ++i) {
            int or = 0, j = i;
            while (j >= 0 && (or & nums[j]) == 0)
                or |= nums[j--];
            ans = Math.max(ans, i - j);
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int longestNiceSubarray(vector<int> &nums) {
        int ans = 0;
        for (int i = 0; i < nums.size(); ++i) {
            int or_ = 0, j = i;
            while (j >= 0 && (or_ & nums[j]) == 0)
                or_ |= nums[j--];
            ans = max(ans, i - j);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func longestNiceSubarray(nums []int) (ans int) {
	for i, or := range nums {
		j := i - 1
		for ; j >= 0 && or&nums[j] == 0; j-- {
			or |= nums[j]
		}
		ans = max(ans, i-j)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

进一步地，由于优雅子数组的所有元素按位与均为 $0$（可以理解成这些二进制数对应的集合没有交集），我们可以用双指针来优化上述过程，如果当前 $\textit{or}$ 与 $\textit{nums}[\textit{right}]$ 按位与的结果不为 $0$，则从 $\textit{or}$ 中去掉 $\textit{nums}[\textit{left}]$，并将 $\textit{left}$ 右移。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干变量。

```py [sol2-Python3]
class Solution:
    def longestNiceSubarray(self, nums: List[int]) -> int:
        ans = left = or_ = 0
        for right, x in enumerate(nums):
            while or_ & x:
                or_ &= ~nums[left]
                left += 1
            or_ |= x
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol2-Java]
class Solution {
    public int longestNiceSubarray(int[] nums) {
        int ans = 0;
        for (int left = 0, right = 0, or = 0; right < nums.length; right++) {
            while ((or & nums[right]) > 0)
                or &= ~nums[left++];
            or |= nums[right];
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int longestNiceSubarray(vector<int> &nums) {
        int ans = 0;
        for (int left = 0, right = 0, or_ = 0; right < nums.size(); right++) {
            while ((or_ & nums[right]) > 0)
                or_ &= ~nums[left++];
            or_ |= nums[right];
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol2-Go]
func longestNiceSubarray(nums []int) (ans int) {
	left, or := 0, 0
	for right, x := range nums {
		for or&x > 0 {
			or &^= nums[left]
			left += 1
		}
		or |= x
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```
