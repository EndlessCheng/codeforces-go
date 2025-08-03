## 方法一：分组循环

三段式子数组必须满足「严格递增 - 严格递减 - 严格递增」，一共三段，每一段**至少要有两个数**。

利用 [分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solutions/2528771/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/)，我们可以遍历所有的极大三段式子数组。极大的意思是子数组不能再往左右延长。

⚠**注意**：第三段的起点也是下一个极大三段式子数组的第一段的起点。

定义：

- 第一段的范围为 $[\textit{start},\textit{peak}]$。
- 第二段的范围为 $[\textit{peak},\textit{bottom}]$。
- 第三段从 $\textit{bottom}$ 开始。

⚠**注意**：第一二段之间的峰顶 $\textit{peak}$ 是第一二段共享的，第二三段之间的谷底 $\textit{bottom}$ 是第二三段共享的。

该三段式子数组中的最大三段式子数组和，由三部分组成：

1. 必须包含从 $\textit{peak}-1$ 到 $\textit{bottom}+1$ 的所有元素。（每一段至少有两个数）
2. 从第一段的倒数第三个数（下标 $\textit{peak}-2$）开始，往左的最大元素和。如果不存在则为 $0$。
3. 从第三段的第三个数（下标 $\textit{bottom}+2$）开始，往右的最大元素和。如果不存在则为 $0$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1BEh3zZEoM/?t=37m15s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxSumTrionic(self, nums: List[int]) -> int:
        n = len(nums)
        ans = -inf
        i = 0
        while i < n:
            # 第一段
            start = i
            i += 1
            while i < n and nums[i - 1] < nums[i]:
                i += 1
            if i == start + 1:  # 第一段至少要有两个数
                continue

            # 第二段
            peak = i - 1
            res = nums[peak - 1] + nums[peak]  # 第一段的最后两个数必选
            while i < n and nums[i - 1] > nums[i]:
                res += nums[i]  # 第二段的所有元素必选
                i += 1
            if i == peak + 1 or i == n:  # 第二段至少要有两个数，第三段至少要有两个数
                continue

            # 第三段
            bottom = i - 1
            res += nums[i]  # 第三段的前两个数必选（第一个数在上面的循环中加了）
            # 从第三段的第三个数往右，计算最大元素和
            max_s = s = 0
            i += 1
            while i < n and nums[i - 1] < nums[i]:
                s += nums[i]
                max_s = max(max_s, s)
                i += 1
            res += max_s

            # 从第一段的倒数第三个数往左，计算最大元素和
            max_s = s = 0
            for j in range(peak - 2, start - 1, -1):
                s += nums[j]
                max_s = max(max_s, s)
            res += max_s
            ans = max(ans, res)

            i = bottom  # 第三段的起点也是下一个极大三段式子数组的第一段的起点
        return ans
```

```java [sol-Java]
class Solution {
    public long maxSumTrionic(int[] nums) {
        int n = nums.length;
        long ans = Long.MIN_VALUE;
        for (int i = 0; i < n;) {
            // 第一段
            int start = i;
            for (i++; i < n && nums[i - 1] < nums[i]; i++);
            if (i == start + 1) { // 第一段至少要有两个数
                continue;
            }

            // 第二段
            int peak = i - 1;
            long res = nums[peak - 1] + nums[peak]; // 第一段的最后两个数必选
            for (; i < n && nums[i - 1] > nums[i]; i++) {
                res += nums[i]; // 第二段的所有元素必选
            }
            if (i == peak + 1 || i == n) { // 第二段至少要有两个数，第三段至少要有两个数
                continue;
            }

            // 第三段
            int bottom = i - 1;
            res += nums[i]; // 第三段的前两个数必选（第一个数在上面的循环中加了）
            // 从第三段的第三个数往右，计算最大元素和
            long maxS = 0;
            long s = 0;
            for (i++; i < n && nums[i - 1] < nums[i]; i++) {
                s += nums[i];
                maxS = Math.max(maxS, s);
            }
            res += maxS;

            // 从第一段的倒数第三个数往左，计算最大元素和
            maxS = 0;
            s = 0;
            for (int j = peak - 2; j >= start; j--) {
                s += nums[j];
                maxS = Math.max(maxS, s);
            }
            res += maxS;
            ans = Math.max(ans, res);

            i = bottom; // 第三段的起点也是下一个极大三段式子数组的第一段的起点
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSumTrionic(vector<int>& nums) {
        int n = nums.size();
        long long ans = LLONG_MIN;
        for (int i = 0; i < n;) {
            // 第一段
            int start = i;
            for (i++; i < n && nums[i - 1] < nums[i]; i++);
            if (i == start + 1) { // 第一段至少要有两个数
                continue;
            }

            // 第二段
            int peak = i - 1;
            long long res = nums[peak - 1] + nums[peak]; // 第一段的最后两个数必选
            for (; i < n && nums[i - 1] > nums[i]; i++) {
                res += nums[i]; // 第二段的所有元素必选
            }
            if (i == peak + 1 || i == n) { // 第二段至少要有两个数，第三段至少要有两个数
                continue;
            }

            // 第三段
            int bottom = i - 1;
            res += nums[i]; // 第三段的前两个数必选（第一个数在上面的循环中加了）
            // 从第三段的第三个数往右，计算最大元素和
            long long max_s = 0, s = 0;
            for (i++; i < n && nums[i - 1] < nums[i]; i++) {
                s += nums[i];
                max_s = max(max_s, s);
            }
            res += max_s;

            // 从第一段的倒数第三个数往左，计算最大元素和
            max_s = 0; s = 0;
            for (int j = peak - 2; j >= start; j--) {
                s += nums[j];
                max_s = max(max_s, s);
            }
            res += max_s;
            ans = max(ans, res);

            i = bottom; // 第三段的起点也是下一个极大三段式子数组的第一段的起点
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSumTrionic(nums []int) int64 {
	n := len(nums)
	ans := math.MinInt
	for i := 0; i < n; {
		// 第一段
		start := i
		for i++; i < n && nums[i-1] < nums[i]; i++ {
		}
		if i == start+1 { // 第一段至少要有两个数
			continue
		}

		// 第二段
		peak := i - 1
		res := nums[peak-1] + nums[peak] // 第一段的最后两个数必选
		for ; i < n && nums[i-1] > nums[i]; i++ {
			res += nums[i] // 第二段的所有元素必选
		}
		if i == peak+1 || i == n { // 第二段至少要有两个数，第三段至少要有两个数
			continue
		}

		// 第三段
		bottom := i - 1
		res += nums[i] // 第三段的前两个数必选（第一个数在上面的循环中加了）
		// 从第三段的第三个数往右，计算最大元素和
		maxS, s := 0, 0
		for i++; i < n && nums[i-1] < nums[i]; i++ {
			s += nums[i]
			maxS = max(maxS, s)
		}
		res += maxS

		// 从第一段的倒数第三个数往左，计算最大元素和
		maxS, s = 0, 0
		for j := peak - 2; j >= start; j-- {
			s += nums[j]
			maxS = max(maxS, s)
		}
		res += maxS
		ans = max(ans, res)

		i = bottom // 第三段的起点也是下一个极大三段式子数组的第一段的起点
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。见 [分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solutions/2528771/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/)。对于本题，同一个元素可以在两个相交的极大三段式子数组中各遍历一次。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：状态机 DP

### 寻找子问题

假设 $\textit{nums}[i]$ 是第三段的最后一个数，那么上一个数呢？

- $\textit{nums}[i-1]$ 是第三段的倒数第二个数。问题变成 $\textit{nums}[i-1]$ 作为第三段的最后一个数时的子数组最大和。
- $\textit{nums}[i-1]$ 是谷底，即第二段的最后一个数。问题变成 $\textit{nums}[i-1]$ 作为第二段的最后一个数时的子数组最大和。

假设 $\textit{nums}[i]$ 是第二段的最后一个数（谷底），那么上一个数呢？

- $\textit{nums}[i-1]$ 是第二段的倒数第二个数。问题变成 $\textit{nums}[i-1]$ 作为第二段的最后一个数时的子数组最大和。
- $\textit{nums}[i-1]$ 是峰顶，即第一段的最后一个数。问题变成 $\textit{nums}[i-1]$ 作为第一段的最后一个数时的子数组最大和。

假设 $\textit{nums}[i]$ 是第一段的最后一个数（峰顶），那么上一个数呢？

- $\textit{nums}[i-1]$ 是第一段的倒数第二个数。问题变成 $\textit{nums}[i-1]$ 作为第一段的最后一个数时的子数组最大和。
- $\textit{nums}[i-1]$ 是第一段的第一个数。

### 状态设计与状态转移方程

根据上述讨论，我们需要知道两个关键信息：

- 子数组最后一个数的下标 $i$。
- $\textit{nums}[i]$ 属于第 $j$ 段。其中 $j\in \{1,2,3\}$。

定义 $f[i][j]$ 表示 $\textit{nums}[i]$ 作为第 $j$ 段的最后一个数时的子数组最大和。

根据上述讨论，状态转移方程为

$$
\begin{aligned}
f[i][3] &= \max(f[i-1][3], f[i-1][2]) + \textit{nums}[i]  & (\textit{nums}[i-1] < \textit{nums}[i])     \\
f[i][2] &= \max(f[i-1][2], f[i-1][1]) + \textit{nums}[i]  & (\textit{nums}[i-1] > \textit{nums}[i])     \\
f[i][1] &= \max(f[i-1][1], \textit{nums}[i-1]) + \textit{nums}[i]  & (\textit{nums}[i-1] > \textit{nums}[i])     \\
\end{aligned}
$$

如果不满足括号中的要求，则 $f[i][j] = -\infty$。

初始值：$f[0][j] = -\infty$。

答案：$f[i][3]$ 的最大值。

**注**：从 $i=1$ 开始算。由于 $f[0][1] = -\infty$，所以 $f[1][1] = \textit{nums}[0] + \textit{nums}[1]$，这也保证了第一段至少有两个数。对于第二段来说，第一段的最后一个数也是第二段的第一个数，所以第二段也至少有两个数。对于第三段来说，第二段的最后一个数也是第三段的第一个数，所以第三段也至少有两个数。（如果 $f[i][3]\ne -\infty$，说明至少发生了一次 $f[i][3] = f[i-1][2] + \textit{nums}[i]$ 的转移。）

代码实现时，第一个维度可以优化掉。

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def maxSumTrionic(self, nums: List[int]) -> int:
        ans = f1 = f2 = f3 = -inf
        for x, y in pairwise(nums):
            f3 = max(f3, f2) + y if x < y else -inf
            f2 = max(f2, f1) + y if x > y else -inf
            f1 = max(f1, x)  + y if x < y else -inf
            ans = max(ans, f3)
        return ans
```

```java [sol-Java]
class Solution {
    public long maxSumTrionic(int[] nums) {
        final long NEG_INF = Long.MIN_VALUE / 2; // 除 2 防止下面加法（和负数相加）溢出
        long ans = NEG_INF;
        long f1 = NEG_INF;
        long f2 = NEG_INF;
        long f3 = NEG_INF;
        for (int i = 1; i < nums.length; i++) {
            int x = nums[i - 1];
            int y = nums[i];
            f3 = x < y ? Math.max(f3, f2) + y : NEG_INF;
            f2 = x > y ? Math.max(f2, f1) + y : NEG_INF;
            f1 = x < y ? Math.max(f1, x) + y : NEG_INF;
            ans = Math.max(ans, f3);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSumTrionic(vector<int>& nums) {
        const long long neg_inf = LLONG_MIN / 2; // 除 2 防止下面加法（和负数相加）溢出
        long long ans = neg_inf, f1 = neg_inf, f2 = neg_inf, f3 = neg_inf;
        for (int i = 1; i < nums.size(); i++) {
            long long x = nums[i - 1], y = nums[i];
            f3 = x < y ? max(f3, f2) + y : neg_inf;
            f2 = x > y ? max(f2, f1) + y : neg_inf;
            f1 = x < y ? max(f1, x) + y : neg_inf;
            ans = max(ans, f3);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxSumTrionic(nums []int) int64 {
	const negInf = math.MinInt / 2 // 除 2 防止下面加法（和负数相加）溢出
	ans, f1, f2, f3 := negInf, negInf, negInf, negInf
	for i := 1; i < len(nums); i++ {
		x, y := nums[i-1], nums[i]
		if x < y { // 第一段或者第三段
			f3 = max(f3, f2) + y
			ans = max(ans, f3)
			f2 = negInf
			f1 = max(f1, x) + y
		} else if x > y { // 第二段
			f2 = max(f2, f1) + y
			f1, f3 = negInf, negInf
		} else {
			f1, f2, f3 = negInf, negInf, negInf
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

1. 滑动窗口与双指针题单的「**六、分组循环**」。
2. 动态规划题单的「**六、状态机 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
