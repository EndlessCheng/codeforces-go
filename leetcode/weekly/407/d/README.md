## 方法一：差分数组

**请先阅读**：[差分数组原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)，推荐和[【图解】从一维差分到二维差分](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/) 一起看。

「子数组内的每个元素的值增加或减少 $1$」，这个操作可以转换成修改**差分数组**两个位置上的数。

设 $d_1$ 为 $\textit{nums}$ 的差分数组，$d_2$ 为 $\textit{target}$ 的差分数组。

由于差分数组和原数组是一一对应的，所以问题等价于：

- 把 $d_1$ 变成 $d_2$。
- 每次操作，可以选择两个下标 $i$ 和 $j$（或者只选一个下标 $i$，对应操作 $\textit{nums}$ 后缀的情况），把 $d_1[i]$ 加一（或减一），把 $d_1[j]$ 减一（或加一）。

从左到右遍历 $d_1$ 和 $d_2$，同时维护一个变量 $s$，表示对 $d_1[i]$ 增大/减少的累积量：

- 如果把 $d_1[i]$ 增大了 $k$，那么后面可以把 $d_1[j]$ **免费减少**，至多免费减少 $k$ 次。
- 如果把 $d_1[i]$ 减少了 $k$，那么后面可以把 $d_1[j]$ **免费增大**，至多免费增大 $k$ 次。

设 $k = d_2[i] - d_1[i]$，分类讨论：

- 如果 $k > 0$ 且 $s\ge 0$，那么必须通过操作，把 $d_1[i]$ 增大到 $d_2[i]$，操作 $k$ 次。
- 如果 $k > 0$ 且 $s < 0$，那么可以免费增大至多 $-s$ 次，如果 $k \le -s$ 则无需额外操作，否则要**额外操作** $k+s$ 次。综合一下，就是额外操作 $\max(s+k,0)$ 次。
- 如果 $k \le 0$ 且 $s\le 0$，那么必须通过操作，把 $d_1[i]$ 减少到 $d_2[i]$，操作 $-k$ 次。
- 如果 $k \le 0$ 且 $s > 0$，那么可以免费减少至多 $s$ 次，如果 $-k \le s$ 则无需额外操作，否则要**额外操作** $-k-s$ 次。综合一下，就是额外操作 $-\min(s+k,0)$ 次。
- 最后把 $k$ 加到 $s$ 中。

代码实现时，可以单独计算 $i=0$ 的情况，方便在计算差分数组的同时计算答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV16Z421N7P2/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int], target: List[int]) -> int:
        s = target[0] - nums[0]
        ans = abs(s)
        for (a, b), (c, d) in pairwise(zip(nums, target)):
            k = (d - c) - (b - a)
            if k > 0:
                ans += k if s >= 0 else max(k + s, 0)
            else:
                ans -= k if s <= 0 else min(k + s, 0)
            s += k
        return ans
```

```java [sol-Java]
class Solution {
    public long minimumOperations(int[] nums, int[] target) {
        long s = target[0] - nums[0];
        long ans = Math.abs(s);
        for (int i = 1; i < nums.length; i++) {
            int k = (target[i] - target[i - 1]) - (nums[i] - nums[i - 1]);
            if (k > 0) {
                ans += s >= 0 ? k : Math.max(k + s, 0);
            } else {
                ans -= s <= 0 ? k : Math.min(k + s, 0);
            }
            s += k;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumOperations(vector<int>& nums, vector<int>& target) {
        long long s = target[0] - nums[0];
        long long ans = abs(s);
        for (int i = 1; i < nums.size(); i++) {
            int k = (target[i] - target[i - 1]) - (nums[i] - nums[i - 1]);
            if (k > 0) {
                ans += s >= 0 ? k : max(k + s, 0LL);
            } else {
                ans -= s <= 0 ? k : min(k + s, 0LL);
            }
            s += k;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumOperations(nums, target []int) int64 {
	s := target[0] - nums[0]
	ans := abs(s)
	for i := 1; i < len(nums); i++ {
		k := (target[i] - target[i-1]) - (nums[i] - nums[i-1])
		if k > 0 {
			if s >= 0 {
				ans += k
			} else {
				ans += max(k+s, 0)
			}
		} else {
			if s <= 0 {
				ans -= k
			} else {
				ans -= min(k+s, 0)
			}
		}
		s += k
	}
	return int64(ans)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：进一步分析差分数组的性质

创建一个数组 $a$，其中 $a[i] = \textit{target}[i]-\textit{nums}[i]$。

如果我们能把 $a$ 变成全为 $0$ 的数组（或者把全为 $0$ 的数组变成 $a$），那么就能把 $\textit{nums}$ 变成 $\textit{target}$。最小操作次数是多少呢？

计算 $a$ 的长为 $n+1$ 的差分数组 $d$。例如示例 2 的 $a=[1,-2,2]$，对应的 $d=[1,-3,4,-2]$，注意这里的 $d[n]=-a[n-1]$。

为什么差分数组的长度是 $n+1$ 而不是 $n$？如果区间操作是 $a$ 的后缀 $a[i]$ 到 $a[n-1]$，那么我们会修改差分数组的 $d[i]$ 和 $d[n]$。一般来说 $d[n]$ 是可以忽略的，但对于本题，我们接下来的分析，需要用到 $d[n]$。

我们把 $d$ 看成是从一个全 $0$ 的差分数组开始，通过若干次操作得到的差分数组，并且每次操作都会**恰好**修改两个数 $d[i]$ 和 $d[j]$，即使操作的是 $a$ 的后缀。

仍然以 $d=[1,-3,4,-2]$ 为例说明。想一想，**差分数组中的正数，对应着什么样的数**？$d[0]=1$，那么一定有一个对应的操作，在把 $d[0]$ 增加 $1$ 的同时，把 $d[j]$ 减少了 $1$。这说明我们可以**把差分数组中的正数和负数一一对应起来**。

![LC3229-cut.png](https://pic.leetcode.cn/1721618787-jmBcJH-LC3229-cut.png)

在 $d=[1,-3,4,-2]$ 这个例子中，$d[0]=1$ 和 $d[2]=4$ 对应着 $5$ 个 $+1$ 操作。由于我们是从全 $0$ 的差分数组开始的，所以一定有 $5$ 个对应的 $-1$ 操作，这正好是 $d[1]=-3$ 和 $d[3]=-2$。

如果同一个 $d[i]$ 上面既有 $+1$，又有 $-1$，那这肯定不如只有 $+1$ 或者只有 $-1$ 优。所以最优操作中，没有 $+1$ 和 $-1$「互相抵消」的情况。

由于每次操作会产生一个 $+1$ 和一个 $-1$，所以**操作次数就等于所有正数** $d[i]$ **之和**。

```py [sol-Python3]
class Solution:
    def minimumOperations(self, nums: List[int], target: List[int]) -> int:
        a = [0] + [t - x for x, t in zip(nums, target)] + [0]  # 前后加 0，方便计算
        return sum(max(y - x, 0) for x, y in pairwise(a))
```

```java [sol-Java]
class Solution {
    public long minimumOperations(int[] nums, int[] target) {
        int n = nums.length;
        long ans = Math.max(target[0] - nums[0], 0);
        for (int i = 1; i < n; i++) {
            ans += Math.max((target[i] - nums[i]) - (target[i - 1] - nums[i - 1]), 0);
        }
        ans += Math.max(-(target[n - 1] - nums[n - 1]), 0);
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumOperations(vector<int>& nums, vector<int>& target) {
        int n = nums.size();
        long long ans = max(target[0] - nums[0], 0);
        for (int i = 1; i < n; i++) {
            ans += max((target[i] - nums[i]) - (target[i - 1] - nums[i - 1]), 0);
        }
        ans += max(-(target[n - 1] - nums[n - 1]), 0);
        return ans;
    }
};
```

```go [sol-Go]
func minimumOperations(nums, target []int) int64 {
	n := len(nums)
	ans := max(target[0]-nums[0], 0)
	for i := 1; i < n; i++ {
		ans += max((target[i]-nums[i])-(target[i-1]-nums[i-1]), 0)
	}
	ans += max(-(target[n-1] - nums[n-1]), 0)
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。

## 相似题目

- [1526. 形成目标数组的子数组最少增加次数](https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/)

更多差分题目，见下面的数据结构题单中的「**§2.1 一维差分**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
