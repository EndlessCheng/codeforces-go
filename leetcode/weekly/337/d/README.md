### 视频讲解

见[【周赛 337】](https://www.bilibili.com/video/BV1EL411C7YU/)。

### 前置知识：同余

两个数 $x$ 和 $y$，如果 $(x-y)\bmod m = 0$，则称 $x$ 与 $y$ 对模 $m$ 同余，记作

$$
x\equiv y \pmod m
$$

例如 $42\equiv 12 \pmod {10}$，$-17\equiv 3 \pmod {10}$。

### 前置知识：处理取模的小技巧

如果 $x$ 和 $y$ 均为非负数，则 $x\equiv y \pmod m$ 相当于

$$
x\bmod m = y\bmod m
$$

如果 $x<0$，$y\ge 0$，则 $x\equiv y \pmod m$ 相当于

$$
x\bmod m + m = y\bmod m
$$

例如 $-17\bmod 10 +10 = -7+10=3$。

为了避免判断 $x$ 是否为负数，等号左边可以写成

$$
(x\bmod m + m) \bmod m
$$

这样无论 $x$ 是否为负数，运算结果都会落在区间 $[0,m)$ 中。

> 注：Python 用户可以忽略，取模运算会保证结果非负。

### 提示 1

下文记 $m=\textit{value}$。

由于同一个数可以操作任意多次，因此每个 $x=\textit{nums}[i]$ 都可以通过操作，变成 $y$，满足

$$
x\equiv y \pmod m
$$

### 提示 2

枚举 $\textit{mex}$。

- 有没有对 $0$ 模 $m$ 同余的数？如果有，把这个数通过操作变成 $0$；否则答案就是 $0$。
- 有没有对 $1$ 模 $m$ 同余的数？如果有，把这个数通过操作变成 $1$；否则答案就是 $1$。
- 有没有对 $2$ 模 $m$ 同余的数？如果有，把这个数通过操作变成 $2$；否则答案就是 $2$。
- ……

### 提示 3

为了方便寻找和维护同余的数字，可以一个哈希表 $\textit{cnt}$ 统计 $(\textit{nums}[i]\bmod m + m) \bmod m$ 的个数。

- $\textit{cnt}[0\bmod m] > 0$？如果不是，则无法得到 $0$，答案是 $0$；如果是，将 $\textit{cnt}[0\bmod m]$ 减一，枚举下一个。
- $\textit{cnt}[1\bmod m] > 0$？如果不是，则无法得到 $1$，答案是 $1$；如果是，将 $\textit{cnt}[1\bmod m]$ 减一，枚举下一个。
- $\textit{cnt}[2\bmod m] > 0$？如果不是，则无法得到 $2$，答案是 $2$；如果是，将 $\textit{cnt}[2\bmod m]$ 减一，枚举下一个。
- ……

```py [sol1-Python3]
class Solution:
    def findSmallestInteger(self, nums: List[int], m: int) -> int:
        cnt = Counter(x % m for x in nums)
        mex = 0
        while cnt[mex % m]:
            cnt[mex % m] -= 1
            mex += 1
        return mex
```

```java [sol1-Java]
class Solution {
    public int findSmallestInteger(int[] nums, int m) {
        var cnt = new HashMap<Integer, Integer>();
        for (int x : nums)
            cnt.merge((x % m + m) % m, 1, Integer::sum); // cnt[(x%m+m)%m]++
        int mex = 0;
        while (cnt.merge(mex % m, -1, Integer::sum) >= 0) // cnt[mex%m]-1 >= 0
            ++mex;
        return mex;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int findSmallestInteger(vector<int> &nums, int m) {
        unordered_map<int, int> cnt;
        for (int x : nums)
            ++cnt[(x % m + m) % m];
        int mex = 0;
        while (cnt[mex % m]--)
            ++mex;
        return mex;
    }
};
```

```go [sol1-Go]
func findSmallestInteger(nums []int, m int) (mex int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[(x%m+m)%m]++
	}
	for cnt[mex%m] > 0 {
		cnt[mex%m]--
		mex++
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。第二个循环至多循环 $O(n)$ 次，因为 `cnt[mex%m]--` 至多执行 $O(n)$ 次（加多少个数，就只能减多少个数）。
- 空间复杂度：$\min(O(n),O(\textit{value}))$。哈希表中至多有 $\min(O(n),O(\textit{value}))$ 个元素。
