请看 [视频讲解](https://www.bilibili.com/video/BV1Dt4y1j7qh) 第三题。

## 方法一：暴力枚举

子数组中任意两个数按位与均为 $0$，意味着任意两个数对应的集合的**交集为空**（见 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)）。

这意味着子数组中的从低到高的第 $i$ 个比特位上，至多有一个比特 $1$，其余均为比特 $0$。例如子数组不可能有两个奇数（最低位为 $1$)，否则这两个数按位与是大于 $0$ 的。

根据鸽巢原理（抽屉原理），在本题数据范围下，优雅子数组的长度不会超过 $30$。例如子数组为 $[2^0,2^1,2^2,\cdots,2^{29}]$，我们无法再加入一个数 $x$，使 $x$ 与子数组中的任何一个数按位与均为 $0$。

既然长度不会超过 $30$，直接暴力枚举子数组的右端点 $i$ 即可。

代码实现时，可以把在子数组中的元素**按位或**起来（求并集），这样可以 $\mathcal{O}(1)$ 判断当前元素是否与前面的元素按位与的结果为 $0$（交集为空）。

```py [sol1-Python3]
class Solution:
    def longestNiceSubarray(self, nums: List[int]) -> int:
        ans = 0
        for i, or_ in enumerate(nums):  # 枚举子数组右端点 i
            j = i - 1
            while j >= 0 and (or_ & nums[j]) == 0:  # nums[j] 与子数组中的任意元素 AND 均为 0
                or_ |= nums[j]  # 加到子数组中
                j -= 1  # 向左扩展
            ans = max(ans, i - j)
        return ans
```

```java [sol1-Java]
class Solution {
    public int longestNiceSubarray(int[] nums) {
        int ans = 0;
        for (int i = 0; i < nums.length; i++) { // 枚举子数组右端点 i
            int or = 0, j = i;
            while (j >= 0 && (or & nums[j]) == 0) // nums[j] 与子数组中的任意元素 AND 均为 0
                or |= nums[j--]; // 加到子数组中
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
        for (int i = 0; i < nums.size(); i++) { // 枚举子数组右端点 i
            int or_ = 0, j = i;
            while (j >= 0 && (or_ & nums[j]) == 0) // nums[j] 与子数组中的任意元素 AND 均为 0
                or_ |= nums[j--]; // 加到子数组中
            ans = max(ans, i - j);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func longestNiceSubarray(nums []int) (ans int) {
	for i, or := range nums { // 枚举子数组右端点 i
		j := i - 1
		for ; j >= 0 && or&nums[j] == 0; j-- { // nums[j] 与子数组中的任意元素 AND 均为 0
			or |= nums[j] // 加到子数组中
		}
		ans = max(ans, i-j)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。

## 方法二：滑动窗口

不了解滑动窗口的同学请看 [基础算法精讲](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

进一步地，既然这些数对应的集合的交集为空，我们可以用**滑动窗口**优化上述过程，右边加入 $\textit{nums}[\textit{right}]$，左边移出 $\textit{nums}[\textit{left}]$。如果 $\textit{or}$ 与新加入的 $\textit{nums}[\textit{right}]$ 有交集，则不断从 $\textit{or}$ 中去掉集合 $\textit{nums}[\textit{left}]$，直到 $\textit{or}$ 与 $\textit{nums}[\textit{right}]$ 交集为空。

如何把集合语言翻译成位运算代码，见 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

```py [sol2-Python3]
class Solution:
    def longestNiceSubarray(self, nums: List[int]) -> int:
        ans = left = or_ = 0
        for right, x in enumerate(nums):
            while or_ & x:  # 有交集
                or_ ^= nums[left]  # 从 or_ 中去掉集合 nums[left]
                left += 1
            or_ |= x  # 把集合 x 并入 or_ 中
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol2-Java]
class Solution {
    public int longestNiceSubarray(int[] nums) {
        int ans = 0;
        for (int left = 0, right = 0, or = 0; right < nums.length; right++) {
            while ((or & nums[right]) > 0) // 有交集
                or ^= nums[left++]; // 从 or 中去掉集合 nums[left]
            or |= nums[right]; // 把集合 nums[right] 并入 or 中
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
            while (or_ & nums[right]) // 有交集
                or_ ^= nums[left++]; // 从 or 中去掉集合 nums[left]
            or_ |= nums[right]; // 把集合 nums[right] 并入 or 中
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
		for or&x > 0 { // 有交集
			or ^= nums[left] // 从 or 中去掉集合 nums[left]
			left += 1
		}
		or |= x // 把集合 x 并入 or 中
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$，仅用到若干变量。
