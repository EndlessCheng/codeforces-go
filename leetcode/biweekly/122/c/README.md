[视频讲解](https://www.bilibili.com/video/BV1oV411D7gB/) 第三题。

考虑这个例子：$\textit{nums}=[2,3,4,5,6]$，每次操作都可以选择 $2$ 和另一个数字 $x$，由于 $x>2$，所以 $2\bmod x = 2$，于是操作等价于：

- 移除 $x$。

所以最后必定只会剩下 $2$。

所以，如果数组中的最小值只有一个，我们可以操作成只剩下一个数，返回 $1$。

但如果最小值不止一个呢？如果能构造出一个小于 $m = \min(\textit{nums})$ 的正整数，那么也可以返回 $1$。

**结论**：当且仅当 $\textit{nums}$ 中有不是 $m$ 的倍数的数，我们才能构造出一个小于 $m$ 的正整数。

**证明**：如果有不是 $m$ 的倍数的数 $x$，那么 $0 < x\bmod m < m$，构造成功。如果所有数都是 $m$ 的倍数，那么任意两个数的模都是 $m$ 的倍数，我们无法得到一个在 $[1,m-1]$ 内的数。

如果所有数都是 $m$ 的倍数，我们可以先用 $m$ 把大于 $m$ 的数都移除，然后剩下的 $\textit{cnt}$ 个 $m$ 两两一对操作，最后剩下 $\left\lceil\dfrac{\textit{cnt}}{2}\right\rceil$ 个数。

```py [sol-Python3]
class Solution:
    def minimumArrayLength(self, nums: List[int]) -> int:
        m = min(nums)
        for x in nums:
            if x % m:
                return 1
        return (nums.count(m) + 1) // 2
```

```java [sol-Java]
class Solution {
    public int minimumArrayLength(int[] nums) {
        int m = Integer.MAX_VALUE;
        for (int x : nums) {
            m = Math.min(m, x);
        }

        for (int x : nums) {
            if (x % m > 0) {
                return 1;
            }
        }

        int cnt = 0;
        for (int x : nums) {
            if (x == m) {
                cnt++;
            }
        }
        return (cnt + 1) / 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumArrayLength(vector<int> &nums) {
        int m = ranges::min(nums);
        for (int x : nums) {
            if (x % m) {
                return 1;
            }
        }
        return (ranges::count(nums, m) + 1) / 2;
    }
};
```

```go [sol-Go]
func minimumArrayLength(nums []int) int {
	m := slices.Min(nums)
	for _, x := range nums {
		if x%m > 0 {
			return 1
		}
	}
	cnt := 0
	for _, x := range nums {
		if x == m {
			cnt++
		}
	}
	return (cnt + 1) / 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
