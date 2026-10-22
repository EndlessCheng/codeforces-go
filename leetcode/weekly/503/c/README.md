## 分析

操作 2 意味着 $\textit{nums}$ 可以视作一个环形数组。操作 1 只反转了数组，没有改变元素的相邻关系，操作之前相邻的元素，操作之后仍然是相邻的。想象你有一串念珠，操作 1 是把念珠翻个面，操作 2 是旋转念珠，都不会拆开念珠。 

所以从最终结果 $[0,1,2,\ldots,n-1]$ 倒推，$\textit{nums}$：

- 要么是由一个递增排列旋转得到，例如 $\textit{nums}=[3,4,0,1,2]$。
- 要么是由一个递减排列旋转得到，例如 $\textit{nums}=[2,1,0,4,3]$。

## nums 由一个递增排列旋转得到

设 $i$ 是满足 $\textit{nums}[i-1] > \textit{nums}[i]$ 的下标。

如果不存在这样的 $i$，那么 $\textit{nums}$ 已经递增，无需操作。

如果只有一个这样的下标，还得满足 $\textit{nums}[0] > \textit{nums}[n-1]$，才能把 $\textit{nums}$ 左旋成递增的。此时有两种方法：

1. 只左旋，操作 $i$ 次。
2. 先反转一次，然后左旋 $n-i$ 次，把 $\textit{nums}$ 变成递减的，然后再反转一次。一共操作 $n-i+2$ 次。

## nums 由一个递减排列旋转得到

设 $i$ 是满足 $\textit{nums}[i-1] < \textit{nums}[i]$ 的下标。

如果不存在这样的 $i$，那么 $\textit{nums}$ 是递减的，只需反转一次。

如果只有一个这样的下标，还得满足 $\textit{nums}[0] < \textit{nums}[n-1]$，才能把 $\textit{nums}$ 左旋成递减的。此时有两种方法：

1. 左旋 $i$ 次，把 $\textit{nums}$ 变成递减的，然后反转一次。一共操作 $i+1$ 次。
2. 先反转一次，然后左旋 $n-i$ 次。一共操作 $n-i+1$ 次。

一共有四种情况，取最小值即为答案。

如果两类方法都没有满足的 $i$，返回 $-1$。

[本题视频讲解](https://www.bilibili.com/video/BV16FG76JEQo/?t=20m29s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: list[int]) -> int:
        n = len(nums)
        cnt = p = 0
        for i in range(1, n):
            if nums[i - 1] > nums[i]:
                cnt += 1
                if cnt > 1:
                    break
                p = i

        if cnt == 0:  # 已是递增
            return 0
        if cnt == 1 and nums[0] > nums[-1]:  # 两个递增段
            ans = min(p, n - p + 2)
        else:
            ans = inf

        cnt = p = 0
        for i in range(1, n):
            if nums[i - 1] < nums[i]:
                cnt += 1
                if cnt > 1:
                    break
                p = i

        if cnt == 0:  # 已是递减
            return 1
        if cnt == 1 and nums[0] < nums[-1]:  # 两个递减段
            ans = min(ans, p + 1, n - p + 1)

        return -1 if ans == inf else ans
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;
        int cnt = 0;
        int p = 0;
        for (int i = 1; i < n && cnt < 2; i++) {
            if (nums[i - 1] > nums[i]) {
                cnt++;
                p = i;
            }
        }

        if (cnt == 0) { // 已是递增
            return 0;
        }
        int ans = Integer.MAX_VALUE;
        if (cnt == 1 && nums[0] > nums[n - 1]) { // 两个递增段
            ans = Math.min(p, n - p + 2);
        }

        cnt = p = 0;
        for (int i = 1; i < n && cnt < 2; i++) {
            if (nums[i - 1] < nums[i]) {
                cnt++;
                p = i;
            }
        }

        if (cnt == 0) { // 已是递减
            return 1;
        }
        if (cnt == 1 && nums[0] < nums[n - 1]) { // 两个递减段
            ans = Math.min(ans, Math.min(p, n - p) + 1);
        }

        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums) {
        int n = nums.size();
        int cnt = 0;
        int p = 0;
        for (int i = 1; i < n && cnt < 2; i++) {
            if (nums[i - 1] > nums[i]) {
                cnt++;
                p = i;
            }
        }

        if (cnt == 0) { // 已是递增
            return 0;
        }
        int ans = INT_MAX;
        if (cnt == 1 && nums[0] > nums[n - 1]) { // 两个递增段
            ans = min(p, n - p + 2);
        }

        cnt = p = 0;
        for (int i = 1; i < n && cnt < 2; i++) {
            if (nums[i - 1] < nums[i]) {
                cnt++;
                p = i;
            }
        }

        if (cnt == 0) { // 已是递减
            return 1;
        }
        if (cnt == 1 && nums[0] < nums[n - 1]) { // 两个递减段
            ans = min(ans, min(p, n - p) + 1);
        }

        return ans == INT_MAX ? -1 : ans;
    }
};
```

```go [sol-Go]
func minOperations(nums []int) int {
	n := len(nums)
	cnt := 0
	p := 0
	for i := 1; i < n && cnt < 2; i++ {
		if nums[i-1] > nums[i] {
			cnt++
			p = i
		}
	}

	if cnt == 0 { // 已是递增
		return 0
	}
	ans := math.MaxInt
	if cnt == 1 && nums[0] > nums[n-1] { // 两个递增段
		ans = min(p, n-p+2)
	}

	cnt = 0
	p = 0
	for i := 1; i < n && cnt < 2; i++ {
		if nums[i-1] < nums[i] {
			cnt++
			p = i
		}
	}

	if cnt == 0 { // 已是递减
		return 1
	}
	if cnt == 1 && nums[0] < nums[n-1] { // 两个递减段
		ans = min(ans, p+1, n-p+1)
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

## 写法二：合二为一

写法一分别讨论了两种情况，写了两个循环。下面把这两个循环合为一个循环。

讨论相邻元素的差值 $\textit{nums}[i] - \textit{nums}[i-1]$：

- 如果 $\textit{nums}$ 是由一个递增排列旋转得到（例如 $\textit{nums}=[3,4,0,1,2]$），那么 $\textit{nums}$ 的相邻元素的差值几乎都是 $1$，除了从 $n-1$ 跳到 $0$ 的那一处是 $-(n-1)$。
- 如果 $\textit{nums}$ 是由一个递减排列旋转得到（例如 $\textit{nums}=[2,1,0,4,3]$），那么 $\textit{nums}$ 的相邻元素的差值几乎都是 $-1$，除了从 $0$ 跳到 $n-1$ 的那一处是 $n-1$。

于是，先求出相邻元素的主要差值 $d$ 是 $1$ 还是 $-1$，然后找到第一个满足 $\textit{nums}[i] - \textit{nums}[i-1]\ne d$ 的位置（记录此时的 $i$ 为 $p$），根据写法一的公式计算答案。如果发现不止一个 $i$ 满足 $\textit{nums}[i] - \textit{nums}[i-1]\ne d$，则返回 $-1$。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: list[int]) -> int:
        n = len(nums)
        if n == 1:
            return 0

        d = nums[1] - nums[0]
        if abs(d) != 1:
            d = nums[2] - nums[1]
            if abs(d) != 1:  # 不能有多个跳变位置
                return -1

        p = 0
        for i in range(1, n):
            if nums[i] - nums[i - 1] != d:
                # i 是跳变位置
                if p > 0:  # 不能有多个跳变位置
                    return -1
                p = i

        if d == 1:
            return min(p, n - p + 2)  # 递增
        return min(p, n - p) + 1  # 递减
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;
        if (n == 1) {
            return 0;
        }

        int d = nums[1] - nums[0];
        if (Math.abs(d) != 1) {
            d = nums[2] - nums[1];
            if (Math.abs(d) != 1) { // 不能有多个跳变位置
                return -1;
            }
        }

        int p = 0;
        for (int i = 1; i < n; i++) {
            if (nums[i] - nums[i - 1] != d) {
                // i 是跳变位置
                if (p > 0) { // 不能有多个跳变位置
                    return -1;
                }
                p = i;
            }
        }

        if (d == 1) {
            return Math.min(p, n - p + 2); // 递增
        }
        return Math.min(p, n - p) + 1; // 递减
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums) {
        int n = nums.size();
        if (n == 1) {
            return 0;
        }

        int d = nums[1] - nums[0];
        if (abs(d) != 1) {
            d = nums[2] - nums[1];
            if (abs(d) != 1) { // 不能有多个跳变位置
                return -1;
            }
        }

        int p = 0;
        for (int i = 1; i < n; i++) {
            if (nums[i] - nums[i - 1] != d) {
                // i 是跳变位置
                if (p > 0) { // 不能有多个跳变位置
                    return -1;
                }
                p = i;
            }
        }

        if (d == 1) {
            return min(p, n - p + 2); // 递增
        }
        return min(p, n - p) + 1; // 递减
    }
};
```

```go [sol-Go]
func minOperations(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	d := nums[1] - nums[0]
	if abs(d) != 1 {
		d = nums[2] - nums[1]
		if abs(d) != 1 { // 不能有多个跳变位置
			return -1
		}
	}

	p := 0
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] != d {
			// i 是跳变位置
			if p > 0 { // 不能有多个跳变位置
				return -1
			}
			p = i
		}
	}

	if d == 1 {
		return min(p, n-p+2) // 递增
	}
	return min(p, n-p) + 1 // 递减
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面思维题单的「**§5.7 分类讨论**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
