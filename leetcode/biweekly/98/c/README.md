### 提示 1

暴力做法是从小到大枚举答案，但这显然会超时。

如何利用「$x$ 不是答案」这一信息呢？

### 提示 2

如果 $1$ 不是答案，说明 $1$ 在 $\textit{nums}$ 中，因为 $1$ 只有一个比特是 $1$（下同）。

继续枚举，如果 $2$ 不是答案，说明 $2$ 在 $\textit{nums}$ 中。

那么 $3$ 肯定不是答案，因为 $1$ 和 $2$ 都在 $\textit{nums}$ 中，且 $1|2=3$。

继续枚举，如果 $4$ 不是答案，说明 $4$ 在 $\textit{nums}$ 中。

那么 $5,6,7$ 肯定不是答案，因为 $1,2,4$ 都在 $\textit{nums}$ 中，它们可以通过或运算组成 $1$ 到 $7$ 中的任意数字。

### 提示 3

因此，我们只需要从小到大挨个判断 $2^i$ 是否在 $\textit{nums}$ 中，第一个不在 $\textit{nums}$ 中的就是答案。

代码实现时，可以用哈希表可以加速这个判断过程。

附：[视频讲解](https://www.bilibili.com/video/BV15D4y1G7ms/)。

```py [sol1-Python3]
class Solution:
    def minImpossibleOR(self, nums: List[int]) -> int:
        s = set(nums)
        return next(1 << i for i in count() if 1 << i not in s)
```

```java [sol1-Java]
class Solution {
    public int minImpossibleOR(int[] nums) {
        var s = new HashSet<Integer>();
        for (int x : nums) s.add(x);
        for (int i = 1; ; i <<= 1)
            if (!s.contains(i))
                return i;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minImpossibleOR(vector<int> &nums) {
        unordered_set s(nums.begin(), nums.end());
        for (int i = 1;; i <<= 1)
            if (!s.count(i))
                return i;
    }
};
```

```go [sol1-Go]
func minImpossibleOR(a []int) (ans int) {
	has := map[int]bool{}
	for _, v := range a {
		has[v] = true
	}
	for i := 1; ; i <<= 1 {
		if !has[i] {
			return i
		}
	}
}
```

### 复杂度分析

- 时间复杂度：$O(n+\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$O(n)$。

### 进一步优化

由于只需要看 $2$ 的幂次，我们可以用一个 $\textit{mask}$ 记录 $\textit{nums}$ 中是 $2$ 的幂次的数。

那么答案就是 $\textit{mask}$ 中的最低比特 $0$。这可以取反后，用 $\textit{lowbit}$ 得到，具体见 [视频讲解](https://www.bilibili.com/video/BV15D4y1G7ms/)。

```py [sol2-Python3]
class Solution:
    def minImpossibleOR(self, nums: List[int]) -> int:
        mask = 0
        for x in nums:
            if (x & (x - 1)) == 0:  # x 是 2 的幂次
                mask |= x
        mask = ~mask
        return mask & -mask  # lowbit
```

```java [sol2-Java]
class Solution {
    public int minImpossibleOR(int[] nums) {
        int mask = 0;
        for (int x : nums)
            if ((x & (x - 1)) == 0) // x 是 2 的幂次
                mask |= x;
        mask = ~mask;
        return mask & -mask; // lowbit
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int minImpossibleOR(vector<int> &nums) {
        int mask = 0;
        for (int x : nums)
            if ((x & (x - 1)) == 0) // x 是 2 的幂次
                mask |= x;
        mask = ~mask;
        return mask & -mask; // lowbit
    }
};
```

```go [sol2-Go]
func minImpossibleOR(nums []int) int {
	mask := 0
	for _, x := range nums {
		if x&(x-1) == 0 { // x 是 2 的幂次
			mask |= x
		}
	}
	mask = ^mask
	return mask & -mask // lowbit
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

### 可以用到这个优化技巧的题目

- [2154. 将找到的值乘以 2](https://leetcode.cn/problems/keep-multiplying-found-values-by-two/)

---

如果你觉得自己的思维能力有些薄弱，可以做做 [从周赛中学算法 - 2022 年周赛题目总结（下篇）](https://leetcode.cn/circle/discuss/WR1MJP/) 中的「思维题」这节，所有题目我都写了题解。
