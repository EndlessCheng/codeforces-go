## 方法一：前后缀分解

枚举替换的元素是 $\textit{nums}[i]$，我们需要知道什么？

- 以 $\textit{nums}[i-1]$ 结尾的最长非递减子数组的长度，记作 $\textit{pre}[i-1]$。
- 以 $\textit{nums}[i+1]$ 开头的最长非递减子数组的长度，记作 $\textit{suf}[i+1]$。
- 如果 $\textit{nums}[i-1] \le \textit{nums}[i+1]$，那么把 $\textit{nums}[i]$ 替换成区间 $[\textit{nums}[i-1],\textit{nums}[i+1]]$ 中的任意整数，就可以得到一个长为 $\textit{pre}[i-1] + 1 + \textit{suf}[i+1]$ 的非递减子数组。
- 也可以把 $\textit{nums}[i]$ 替换成 $\textit{nums}[i-1]$，拼在 $\textit{pre}[i-1]$ 的后面，得到一个长为 $\textit{pre}[i-1] + 1$ 的非递减子数组。
- 也可以把 $\textit{nums}[i]$ 替换成 $\textit{nums}[i+1]$，拼在 $\textit{suf}[i+1]$ 的前面，得到一个长为 $\textit{suf}[i+1] + 1$ 的非递减子数组。
- 所有情况取最大值。

对于 $\textit{suf}$，我们可以倒着遍历 $\textit{nums}$，如果 $\textit{nums}[i]\le \textit{nums}[i+1]$，那么 $\textit{nums}[i]$ 可以拼在 $\textit{suf}[i+1]$ 的前面，所以 $\textit{suf}[i] = \textit{suf}[i+1]+1$；否则 $\textit{suf}[i] = 1$。

对于 $\textit{pre}$，我们可以正着遍历 $\textit{nums}$，如果 $\textit{nums}[i-1]\le \textit{nums}[i]$，那么 $\textit{nums}[i]$ 可以拼在 $\textit{pre}[i-1]$ 的后面，所以 $\textit{pre}[i] = \textit{pre}[i-1]+1$；否则 $\textit{pre}[i] = 1$。

代码实现时，可以在遍历 $\textit{nums}$ 的同时计算 $\textit{pre}$，所以 $\textit{pre}$ 可以简化成一个变量。

[本题视频讲解](https://www.bilibili.com/video/BV19bkQBkEhG/?t=2m48s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 1:
            return 1

        suf = [0] * n
        suf[-1] = 1
        ans = 2
        for i in range(n - 2, 0, -1):
            if nums[i] <= nums[i + 1]:
                suf[i] = suf[i + 1] + 1
                ans = max(ans, suf[i] + 1)  # 把 nums[i-1] 拼在 suf[i] 前面
            else:
                suf[i] = 1

        pre = 1
        for i in range(1, n - 1):
            if nums[i - 1] <= nums[i + 1]:
                ans = max(ans, pre + 1 + suf[i + 1])  # 替换 nums[i]
            if nums[i - 1] <= nums[i]:
                pre += 1
                ans = max(ans, pre + 1)  # 把 nums[i+1] 拼在 pre 后面
            else:
                pre = 1
        return ans
```

```java [sol-Java]
class Solution {
    public int longestSubarray(int[] nums) {
        int n = nums.length;
        if (n == 1) {
            return 1;
        }

        int[] suf = new int[n];
        suf[n - 1] = 1;
        int ans = 2;
        for (int i = n - 2; i > 0; i--) {
            if (nums[i] <= nums[i + 1]) {
                suf[i] = suf[i + 1] + 1;
                ans = Math.max(ans, suf[i] + 1); // 把 nums[i-1] 拼在 suf[i] 前面
            } else {
                suf[i] = 1;
            }
        }

        int pre = 1;
        for (int i = 1; i < n - 1; i++) {
            if (nums[i - 1] <= nums[i + 1]) {
                ans = Math.max(ans, pre + 1 + suf[i + 1]); // 替换 nums[i]
            }
            if (nums[i - 1] <= nums[i]) {
                pre++;
                ans = Math.max(ans, pre + 1); // 把 nums[i+1] 拼在 pre 后面
            } else {
                pre = 1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubarray(vector<int>& nums) {
        int n = nums.size();
        if (n == 1) {
            return 1;
        }

        vector<int> suf(n);
        suf[n - 1] = 1;
        int ans = 2;
        for (int i = n - 2; i > 0; i--) {
            if (nums[i] <= nums[i + 1]) {
                suf[i] = suf[i + 1] + 1;
                ans = max(ans, suf[i] + 1); // 把 nums[i-1] 拼在 suf[i] 前面
            } else {
                suf[i] = 1;
            }
        }

        int pre = 1;
        for (int i = 1; i < n - 1; i++) {
            if (nums[i - 1] <= nums[i + 1]) {
                ans = max(ans, pre + 1 + suf[i + 1]); // 替换 nums[i]
            }
            if (nums[i - 1] <= nums[i]) {
                pre++;
                ans = max(ans, pre + 1); // 把 nums[i+1] 拼在 pre 后面
            } else {
                pre = 1;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	suf := make([]int, n)
	suf[n-1] = 1
	ans := 2
	for i := n - 2; i > 0; i-- {
		if nums[i] <= nums[i+1] {
			suf[i] = suf[i+1] + 1
			ans = max(ans, suf[i]+1) // 把 nums[i-1] 拼在 suf[i] 前面
		} else {
			suf[i] = 1
		}
	}

	pre := 1
	for i := 1; i < n-1; i++ {
		if nums[i-1] <= nums[i+1] {
			ans = max(ans, pre+1+suf[i+1]) // 替换 nums[i]
		}
		if nums[i-1] <= nums[i] {
			pre++
			ans = max(ans, pre+1) // 把 nums[i+1] 拼在 pre 后面
		} else {
			pre = 1
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：状态机 DP

做法类似 [1186. 删除一次得到子数组最大和](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/)，[我的题解](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/solutions/2321829/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-hzz6/)。

仿照 1186 题，定义 $f[i][j]$ 表示非递减子数组右端点的下标是 $i$，在不能/可以替换元素的情况下，子数组的最长长度。为保证比大小时元素未被替换，这里规定 $\textit{nums}[i]$ 不能替换。

- 如果 $j=0$（不能替换）：
  - 如果 $\textit{nums}[i-1]\le \textit{nums}[i]$，我们可以拼在 $i-1$ 后面，问题变成非递减子数组右端点的下标是 $i-1$，在不能替换元素的情况下，子数组的最长长度，即 $f[i-1][0]$，得到 $f[i][0] = f[i-1][0] + 1$。
  - 否则只能 $\textit{nums}[i]$ 单独一个数，$f[i][0] = 1$。
  - **注**：$f[i][0]$ 等同于方法一的 $\textit{pre}[i]$。
- 如果 $j=1$（可以替换）：
  - 情况一：同上，如果 $\textit{nums}[i-1]\le \textit{nums}[i]$，那么 $f[i][1] = f[i-1][1] + 1$。
  - 情况二：如果 $\textit{nums}[i-2]\le \textit{nums}[i]$，那么可以替换 $\textit{nums}[i-1]$，问题变成非递减子数组右端点的下标是 $i-2$，在不能替换元素的情况下，子数组的最长长度，即 $f[i-2][0]$，得到 $f[i][1] = f[i-2][0] + 2$。否则不能拼接，只能替换 $\textit{nums}[i-1]$，得到 $f[i][1] = 2$。
  - 两种情况取最大值。

初始值：$f[0][j] = 1$。

最后计算 $f[i][1]$ 的最大值。

但这没有考虑子数组右端点的下标是 $i$，且替换 $\textit{nums}[i]$ 的情况。

我们替换 $\textit{nums}[i]$，拼在 $f[i-1][0]$ 的后面，得到 $f[i-1][0]+1$。所以还需要计算 $f[i-1][0]+1$ 的最大值。

最终答案为 $f[i][1]$ 的最大值，$f[i-1][0]+1$ 的最大值，二者取最大值。

### 优化前

```py [sol-Python3]
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        n = len(nums)
        f = [[0, 0] for _ in range(n)]
        f[0] = [1, 1]

        ans = 1  # 以 nums[0] 结尾的子数组长度
        for i in range(1, n):
            if nums[i - 1] <= nums[i]:
                f[i][0] = f[i - 1][0] + 1
                f[i][1] = f[i - 1][1] + 1
            else:
                f[i][0] = 1
                # 不需要写 f[i][1] = 1，因为下面算出来的值至少是 2

            if i >= 2 and nums[i - 2] <= nums[i]:
                f[i][1] = max(f[i][1], f[i - 2][0] + 2)
            else:
                f[i][1] = max(f[i][1], 2)

            ans = max(ans, f[i - 1][0] + 1, f[i][1])
        return ans
```

```java [sol-Java]
class Solution {
    public int longestSubarray(int[] nums) {
        int n = nums.length;
        int[][] f = new int[n][2];
        f[0][0] = f[0][1] = 1;

        int ans = 1; // 以 nums[0] 结尾的子数组长度
        for (int i = 1; i < n; i++) {
            if (nums[i - 1] <= nums[i]) {
                f[i][0] = f[i - 1][0] + 1;
                f[i][1] = f[i - 1][1] + 1;
            } else {
                f[i][0] = 1;
                // 不需要写 f[i][1] = 1，因为下面算出来的值至少是 2
            }

            if (i >= 2 && nums[i - 2] <= nums[i]) {
                f[i][1] = Math.max(f[i][1], f[i - 2][0] + 2);
            } else {
                f[i][1] = Math.max(f[i][1], 2);
            }

            ans = Math.max(ans, Math.max(f[i - 1][0] + 1, f[i][1]));
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubarray(vector<int>& nums) {
        int n = nums.size();
        vector<array<int, 2>> f(n);
        f[0] = {1, 1};

        int ans = 1; // 以 nums[0] 结尾的子数组长度
        for (int i = 1; i < n; i++) {
            if (nums[i - 1] <= nums[i]) {
                f[i][0] = f[i - 1][0] + 1;
                f[i][1] = f[i - 1][1] + 1;
            } else {
                f[i][0] = 1;
                // 不需要写 f[i][1] = 1，因为下面算出来的值至少是 2
            }

            if (i >= 2 && nums[i - 2] <= nums[i]) {
                f[i][1] = max(f[i][1], f[i - 2][0] + 2);
            } else {
                f[i][1] = max(f[i][1], 2);
            }

            // ans = max({ans, f[i - 1][0] + 1, f[i][1]}); 这种写法比下面的慢
            ans = max(ans, max(f[i - 1][0] + 1, f[i][1]));
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestSubarray(nums []int) int {
	n := len(nums)
	f := make([][2]int, n)
	f[0] = [2]int{1, 1}

	ans := 1 // 以 nums[0] 结尾的子数组长度
	for i := 1; i < n; i++ {
		if nums[i-1] <= nums[i] {
			f[i][0] = f[i-1][0] + 1
			f[i][1] = f[i-1][1] + 1
		} else {
			f[i][0] = 1
			// 不需要写 f[i][1] = 1，因为下面算出来的值至少是 2
		}

		if i >= 2 && nums[i-2] <= nums[i] {
			f[i][1] = max(f[i][1], f[i-2][0]+2)
		} else {
			f[i][1] = max(f[i][1], 2)
		}

		ans = max(ans, f[i-1][0]+1, f[i][1])
	}
	return ans
}
```

### 空间优化

```py [sol-Python3]
# 更快的写法见【Python3 写法二】
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        pre0, f0, f1 = 0, 1, 1

        ans = 1  # 以 nums[0] 结尾的子数组长度
        for i in range(1, len(nums)):
            tmp = f0
            if nums[i - 1] <= nums[i]:
                f0 += 1
                f1 += 1
            else:
                f0 = 1
                f1 = 0  # 清除旧数据

            if i >= 2 and nums[i - 2] <= nums[i]:
                f1 = max(f1, pre0 + 2)
            else:
                f1 = max(f1, 2)

            ans = max(ans, tmp + 1, f1)
            pre0 = tmp
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        pre0, f0, f1 = 0, 1, 1

        ans = 1  # 以 nums[0] 结尾的子数组长度
        for i in range(1, len(nums)):
            tmp = f0
            if nums[i - 1] <= nums[i]:
                f0 += 1
                f1 += 1
            else:
                f0 = 1
                f1 = 0  # 清除旧数据

            if i >= 2 and nums[i - 2] <= nums[i]:
                if pre0 + 2 > f1: f1 = pre0 + 2
            else:
                if 2 > f1: f1 = 2

            if f1 > ans: ans = f1
            if tmp + 1 > ans: ans = tmp + 1
            pre0 = tmp
        return ans
```

```java [sol-Java]
class Solution {
    public int longestSubarray(int[] nums) {
        int pre0 = 0, f0 = 1, f1 = 1;

        int ans = 1; // 以 nums[0] 结尾的子数组长度
        for (int i = 1; i < nums.length; i++) {
            int tmp = f0;
            if (nums[i - 1] <= nums[i]) {
                f0++;
                f1++;
            } else {
                f0 = 1;
                f1 = 0; // 清除旧数据
            }

            if (i >= 2 && nums[i - 2] <= nums[i]) {
                f1 = Math.max(f1, pre0 + 2);
            } else {
                f1 = Math.max(f1, 2);
            }

            ans = Math.max(ans, Math.max(tmp + 1, f1));
            pre0 = tmp;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSubarray(vector<int>& nums) {
        int pre0 = 0, f0 = 1, f1 = 1;

        int ans = 1; // 以 nums[0] 结尾的子数组长度
        for (int i = 1; i < nums.size(); i++) {
            int tmp = f0;
            if (nums[i - 1] <= nums[i]) {
                f0++;
                f1++;
            } else {
                f0 = 1;
                f1 = 0; // 清除旧数据
            }

            if (i >= 2 && nums[i - 2] <= nums[i]) {
                f1 = max(f1, pre0 + 2);
            } else {
                f1 = max(f1, 2);
            }

            ans = max(ans, max(tmp + 1, f1));
            pre0 = tmp;
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestSubarray(nums []int) int {
	pre0, f0, f1 := 0, 1, 1

	ans := 1 // 以 nums[0] 结尾的子数组长度
	for i := 1; i < len(nums); i++ {
		tmp := f0
		if nums[i-1] <= nums[i] {
			f0++
			f1++
		} else {
			f0 = 1
			f1 = 0 // 清除旧数据
		}

		if i >= 2 && nums[i-2] <= nums[i] {
			f1 = max(f1, pre0+2)
		} else {
			f1 = max(f1, 2)
		}

		ans = max(ans, tmp+1, f1)
		pre0 = tmp
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

1. 把「非递减」改成「严格递增」怎么做？
2. 把「子数组」改成「子序列」怎么做？解决非递减和严格递增两种情况。

欢迎在评论区分享你的思路/代码。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」「**§7.3 子数组 DP**」和「**六、状态机 DP**」。

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
