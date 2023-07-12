[视频讲解（第三题）](https://www.bilibili.com/video/BV18u411Y7Gt/)

### 提示 1

题目最后要求机器人之间的距离，此时把任意两个机器人的位置交换，并不会对答案产生影响。

> 假设 $d$ 秒后机器人的位置数组为 $[1,2,3]$，那么交换成 $[2,1,3]$，所有机器人之间两两距离之和保持不变。

既然如此，那么可以把机器人都看成是**完全一样的，无法区分**。

### 提示 2

相撞等价于**机器人互相穿过对方**，因为我们无法区分机器人。

所以可以无视相撞的规则，把每个机器人都看成是独立运动的。

类似的思路在 [1503. 所有蚂蚁掉下来前的最后一刻](https://leetcode.cn/problems/last-moment-before-all-ants-fall-out-of-a-plank/) 中出现过。

### 提示 3

设 $d$ 秒后机器人的位置数组为 $a$，根据提示 1，可以把数组 $a$ 从小到大排序，再计算所有机器人之间两两距离之和。

从小到大枚举 $a[i]$，此时左边有 $i$ 个数都不超过 $a[i]$，那么 $a[i]$ 与其左侧机器人的距离之和为

$$
\begin{aligned}
&(a[i] - a[0])+ (a[i] - a[1]) + \cdots + (a[i] - a[i-1])\\
=&\ i\cdot a[i] - (a[0] + a[1] + \cdots + a[i-1])
\end{aligned}
$$

其中 $a[0] + a[1] + \cdots + a[i-1]$ 可以一边遍历 $a$，一边计算出来。

计算时，为了避免溢出，需要取模。这样做的正确性见下面的「算法小课堂：模运算」。

### 答疑

**问**：为什么不能对 $a[i]$ 取模？

**答**：注意 $a$ 中第 $i$ 小的数要乘上 $i$，取模后每个元素的大小关系就乱了，原来第 $i$ 小的数要乘的就不一定是 $i$ 了，所以会算出错误的结果。

```py [sol-Python3]
class Solution:
    def sumDistance(self, nums: List[int], s: str, d: int) -> int:
        MOD = 10 ** 9 + 7
        for i, c in enumerate(s):
            nums[i] += d if c == 'R' else -d
        nums.sort()
        ans = s = 0
        for i, x in enumerate(nums):
            ans += i * x - s
            s += x
        return ans % MOD
```

```java [sol-Java]
class Solution {
    public int sumDistance(int[] nums, String s, int d) {
        final long MOD = (long) 1e9 + 7;
        int n = nums.length;
        var a = new long[n];
        for (int i = 0; i < n; i++) // 注意 2e9+1e9 溢出了
            a[i] = (long) nums[i] + d * ((s.charAt(i) & 2) - 1); // L=-1, R=1
        long ans = 0, sum = 0;
        Arrays.sort(a);
        for (int i = 0; i < n; i++) {
            ans = (ans + i * a[i] - sum) % MOD;
            sum += a[i];
        }
        return (int) ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumDistance(vector<int> &nums, string s, int d) {
        const int MOD = 1e9 + 7;
        int n = nums.size();
        long long a[n], ans = 0, sum = 0;
        for (int i = 0; i < n; i++) // 注意 2e9+1e9 溢出了
            a[i] = (long long) nums[i] + d * ((s[i] & 2) - 1); // L=-1, R=1
        sort(a, a + n);
        for (int i = 0; i < n; i++) {
            ans = (ans + i * a[i] - sum) % MOD;
            sum += a[i];
        }
        return ans;
    }
};
```

```go [sol-Go]
func sumDistance(nums []int, s string, d int) (ans int) {
	const mod int = 1e9 + 7
	for i, c := range s {
		nums[i] += d * int(c&2-1) // L=-1, R=1
	}
	sort.Ints(nums)
	sum := 0
	for i, x := range nums {
		ans = (ans + i*x - sum) % mod
		sum += x
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。如果需要用一个新的 $\textit{nums}$ 数组记录则需要 $\mathcal{O}(n)$ 空间，否则为 $\mathcal{O}(1)$。

### 算法小课堂：模运算

如果让你计算 $1234\cdot 6789$ 的**个位数**，你会如何计算？

由于只有个位数会影响到乘积的个位数，那么 $4\cdot 9=36$ 的个位数 $6$ 就是答案。

对于 $1234+6789$ 的个位数，同理，$4+9=13$ 的个位数 $3$ 就是答案。

你能把这个结论抽象成数学等式吗？

一般地，涉及到取模的题目，通常会用到如下等式（上面计算的是 $m=10$）：

$$
(a+b)\bmod m = ((a\bmod m) + (b\bmod m)) \bmod m
$$

$$
(a\cdot b) \bmod m=((a\bmod m)\cdot  (b\bmod m)) \bmod m
$$

证明：根据**带余除法**，任意整数 $a$ 都可以表示为 $a=km+r$，这里 $r$ 相当于 $a\bmod m$。那么设 $a=k_1m+r_1,\ b=k_2m+r_2$。

第一个等式：

$$
\begin{aligned}
&\ (a+b) \bmod m\\
=&\ ((k_1+k_2) m+r_1+r_2)\bmod m\\
=&\ (r_1+r_2)\bmod m\\
=&\ ((a\bmod m) + (b\bmod m)) \bmod m
\end{aligned}
$$

第二个等式：

$$
\begin{aligned}
&\ (a\cdot b) \bmod m\\
=&\ (k_1k_2m^2+(k_1r_2+k_2r_1)m+r_1r_2)\bmod m\\
=&\ (r_1r_2)\bmod m\\
=&\ ((a\bmod m)\cdot  (b\bmod m)) \bmod m
\end{aligned}
$$

根据这两个恒等式，可以随意地对代码中的加法和乘法的结果取模。
