[本题视频讲解](https://www.bilibili.com/video/BV1jg4y1y7PA/)

为方便描述，下文将 $\textit{nums}$ 简记为 $a$，其长度为 $n$。

首先来看能移出哪些**后缀**。

设 $a$ 的**最长严格递增前缀**的最后一个数是 $a[i]$。

这里可以特判下，如果 $i=n-1$，说明 $a$ 是严格递增数组，所有非空子数组都可以移除，我们直接返回 $\dfrac{n(n+1)}{2}$。

否则 $i<n-1$，我们可以移除这些区间中的数：

- $[i+1,n-1]$
- $[i,n-1]$
- $[i-1,n-1]$
- $\cdots$
- $[0,n-1]$

这一共有 $i+2$ 个。

例如 $[3,4,1,2]$ 的 $i=1$，我们可以移除

- $[1,2]$
- $[4,1,2]$
- $[3,4,1,2]$

这 $3$ 个后缀。

然后来看一般的情况。

移除子数组后，剩下的部分是一个前缀加一个后缀，我们从右往左枚举后缀的第一个数 $a[j]$，同时维护最长前缀的最后一个数的下标 $i$。

> 由于不能移除空数组，$i$ 与 $j$ 的中间至少要有一个数，所以必须要有 $i\le j-2$。但是 $i=j-1$ 的情况说明 $a$ 是严格递增数组，我们已经在前面判断了，所以无需判断 $i$ 和 $j-2$ 的大小关系。

我们可以像 [滑动窗口](https://www.bilibili.com/video/BV1hd4y1r7Gq/) 那样不断左移 $i$ 直到 $i<0$ 或者 $a[i]<a[j]$ 为止。

此时我们可以移除这些区间中的数：

- $[i+1,j-1]$
- $[i,j-1]$
- $[i-1,j-1]$
- $\cdots$
- $[0,j-1]$

这一共有 $i+2$ 个。注意 $i=-1$ 时只能移除前缀 $[0,j-1]$。

累加这些个数，即为答案。

注意 $j$ 枚举到 $1$ 就行，因为我们不能移除空数组。

```py [sol-Python3]
class Solution:
    def incremovableSubarrayCount(self, a: List[int]) -> int:
        n = len(a)
        i = 0
        while i < n - 1 and a[i] < a[i + 1]:
            i += 1
        if i == n - 1:  # 每个非空子数组都可以移除
            return n * (n + 1) // 2

        ans = i + 2
        j = n - 1
        while j == n - 1 or a[j] < a[j + 1]:
            while i >= 0 and a[i] >= a[j]:
                i -= 1
            ans += i + 2
            j -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public long incremovableSubarrayCount(int[] a) {
        int n = a.length;
        int i = 0;
        while (i < n - 1 && a[i] < a[i + 1]) {
            i++;
        }
        if (i == n - 1) { // 每个非空子数组都可以移除
            return (long) n * (n + 1) / 2;
        }

        long ans = i + 2;
        for (int j = n - 1; j == n - 1 || a[j] < a[j + 1]; j--) {
            while (i >= 0 && a[i] >= a[j]) {
                i--;
            }
            ans += i + 2;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long incremovableSubarrayCount(vector<int> &a) {
        int n = a.size();
        int i = 0;
        while (i < n - 1 && a[i] < a[i + 1]) {
            i++;
        }
        if (i == n - 1) { // 每个非空子数组都可以移除
            return (long long) n * (n + 1) / 2;
        }

        long long ans = i + 2;
        for (int j = n - 1; j == n - 1 || a[j] < a[j + 1]; j--) {
            while (i >= 0 && a[i] >= a[j]) {
                i--;
            }
            ans += i + 2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func incremovableSubarrayCount(a []int) int64 {
	n := len(a)
	i := 0
	for i < n-1 && a[i] < a[i+1] {
		i++
	}
	if i == n-1 { // 每个非空子数组都可以移除
		return int64(n) * int64(n+1) / 2
	}

	ans := int64(i + 2)
	for j := n - 1; j == n-1 || a[j] < a[j+1]; j-- {
		for i >= 0 && a[i] >= a[j] {
			i--
		}
		ans += int64(i + 2)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。注意二重循环中的 $i$ 只会减小不会变大，$i$ 只会减小 $\mathcal{O}(n)$ 次，所以二重循环的次数也是 $\mathcal{O}(n)$ 的。
- 空间复杂度：$\mathcal{O}(1)$。

#### 相似题目

- [1574. 删除最短的子数组使剩余数组有序](https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/)
- [【题单】滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
