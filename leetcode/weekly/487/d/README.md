为方便描述，下文把 $\textit{nums}$ 简称为 $a$。

枚举删除的是 $a[i]$，我们需要知道：

- 以 $i-1$ 结尾的最长交替子数组的长度，记作 $\textit{pre}[i-1]$。
- 以 $i+1$ 开头的最长交替子数组的长度，记作 $\textit{suf}[i+1]$。

这可以递推算出。以 $\textit{pre}$ 为例：

- 如果 $i=0$ 或者 $a[i-1] = a[i]$，那么 $\textit{pre}[i]$ 只能是 $1$。
- 否则，$\textit{pre}[i]$ 至少是 $2$。如果 $a[i-2] \ne a[i-1]$，并且 $a[i-2] < a[i-1] > a[i]$ 和 $a[i-2] > a[i-1] < a[i]$ 其中一个成立，那么可以把 $a[i]$ 拼接在以 $i-1$ 结尾的最长交替子数组的后面，即 $\textit{pre}[i] = \textit{pre}[i-1] + 1$。

然后来计算答案。

删除 $a[i]$ 后，有两种情况可以把左右两侧的交替子数组拼在一起：

- $a[i-2] < a[i-1] > a[i+1] < a[i+2]$。
- $a[i-2] > a[i-1] < a[i+1] > a[i+2]$。

此时可以拼接，拼接后的交替子数组的长度为

$$
\textit{pre}[i-1] + \textit{suf}[i+1]
$$

用上式更新答案的最大值。

还有两种特殊情况：

- 只满足 $a[i-2] < a[i-1] > a[i+1]$ 或者 $a[i-2] > a[i-1] < a[i+1]$，拼接后的交替子数组的长度为 $\textit{pre}[i-1] + 1$。
- 只满足 $a[i-1] > a[i+1] < a[i+2]$ 或者 $a[i-1] < a[i+1] > a[i+2]$，拼接后的交替子数组的长度为 $\textit{suf}[i+1] + 1$。

此外，还可以不删除元素，最长长度为 $\max(\textit{pre})$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def calc(self, a: List[int]) -> List[int]:
        f = [0] * len(a)  # f[i] 表示以 i 结尾的最长交替子数组的长度
        for i, x in enumerate(a):
            if i == 0 or a[i - 1] == x:
                f[i] = 1
            elif i > 1 and a[i - 2] != a[i - 1] and (a[i - 2] < a[i - 1]) == (a[i - 1] > x):
                f[i] = f[i - 1] + 1
            else:
                f[i] = 2
        return f

    def longestAlternating(self, nums: List[int]) -> int:
        n = len(nums)
        pre = self.calc(nums)  # pre[i] 表示以 i 结尾的最长交替子数组的长度
        suf = self.calc(nums[::-1])[::-1]  # suf[i] 表示以 i 开头的最长交替子数组的长度

        # 不删除元素时的最长交替子数组的长度
        ans = max(pre)

        # 枚举删除 nums[i]
        for i in range(1, n - 1):
            if nums[i - 1] == nums[i + 1]:  # 无法拼接
                continue

            # 计算 (i-2,i-1), (i-1,i+1), (i+1,i+2) 的大小关系
            x = (nums[i - 2] > nums[i - 1]) - (nums[i - 2] < nums[i - 1]) if i > 1 else 0
            y = (nums[i - 1] > nums[i + 1]) - (nums[i - 1] < nums[i + 1])
            z = (nums[i + 1] > nums[i + 2]) - (nums[i + 1] < nums[i + 2]) if i < n - 2 else 0

            if x == -y and x == z:  # 左右两边可以拼接
                ans = max(ans, pre[i - 1] + suf[i + 1])
            else:
                if x == -y:
                    ans = max(ans, pre[i - 1] + 1)  # 只拼接 nums[i+1]
                if z == -y:
                    ans = max(ans, suf[i + 1] + 1)  # 只拼接 nums[i-1]

        return ans
```

```java [sol-Java]
class Solution {
    public int longestAlternating(int[] nums) {
        int n = nums.length;
        int[] pre = calc(nums); // pre[i] 表示以 i 结尾的最长交替子数组的长度

        reverse(nums);
        int[] suf = calc(nums); // suf[i] 表示以 i 开头的最长交替子数组的长度
        reverse(suf);
        reverse(nums);

        // 不删除元素时的最长交替子数组的长度
        int ans = 0;
        for (int x : pre) {
            ans = Math.max(ans, x);
        }

        // 枚举删除 nums[i]
        for (int i = 1; i < n - 1; i++) {
            if (nums[i - 1] == nums[i + 1]) { // 无法拼接
                continue;
            }

            // 计算 (i-2,i-1), (i-1,i+1), (i+1,i+2) 的大小关系
            int x = i > 1 ? Integer.compare(nums[i - 2], nums[i - 1]) : 0;
            int y = Integer.compare(nums[i - 1], nums[i + 1]);
            int z = i < n - 2 ? Integer.compare(nums[i + 1], nums[i + 2]) : 0;

            if (x == -y && x == z) { // 左右两边可以拼接
                ans = Math.max(ans, pre[i - 1] + suf[i + 1]);
            } else {
                if (x == -y) {
                    ans = Math.max(ans, pre[i - 1] + 1); // 只拼接 nums[i+1]
                }
                if (z == -y) {
                    ans = Math.max(ans, suf[i + 1] + 1); // 只拼接 nums[i-1]
                }
            }
        }

        return ans;
    }

    private int[] calc(int[] a) {
        int n = a.length;
        int[] f = new int[n]; // f[i] 表示以 i 结尾的最长交替子数组的长度
        for (int i = 0; i < n; i++) {
            if (i == 0 || a[i - 1] == a[i]) {
                f[i] = 1;
            } else if (i > 1 && a[i - 2] != a[i - 1] && (a[i - 2] < a[i - 1]) == (a[i - 1] > a[i])) {
                f[i] = f[i - 1] + 1;
            } else {
                f[i] = 2;
            }
        }
        return f;
    }

    private void reverse(int[] a) {
        int i = 0;
        int j = a.length - 1;
        while (i < j) {
            int tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
            i++;
            j--;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> calc(const vector<int>& a) {
        int n = a.size();
        vector<int> f(n); // f[i] 表示以 i 结尾的最长交替子数组的长度
        for (int i = 0; i < n; i++) {
            if (i == 0 || a[i - 1] == a[i]) {
                f[i] = 1;
            } else if (i > 1 && a[i - 2] != a[i - 1] && (a[i - 2] < a[i - 1]) == (a[i - 1] > a[i])) {
                f[i] = f[i - 1] + 1;
            } else {
                f[i] = 2;
            }
        }
        return f;
    }

public:
    int longestAlternating(vector<int>& nums) {
        int n = nums.size();
        vector<int> pre = calc(nums); // pre[i] 表示以 i 结尾的最长交替子数组的长度

        ranges::reverse(nums);
        vector<int> suf = calc(nums); // suf[i] 表示以 i 开头的最长交替子数组的长度
        ranges::reverse(suf);
        ranges::reverse(nums);

        // 不删除元素时的最长交替子数组的长度
        int ans = ranges::max(pre);

        // 枚举删除 nums[i]
        for (int i = 1; i < n - 1; i++) {
            if (nums[i - 1] == nums[i + 1]) { // 无法拼接
                continue;
            }

            // 计算 (i-2,i-1), (i-1,i+1), (i+1,i+2) 的大小关系
            int x = i > 1 ? (nums[i - 2] > nums[i - 1]) - (nums[i - 2] < nums[i - 1]) : 0;
            int y = (nums[i - 1] > nums[i + 1]) - (nums[i - 1] < nums[i + 1]);
            int z = i < n - 2 ? (nums[i + 1] > nums[i + 2]) - (nums[i + 1] < nums[i + 2]) : 0;

            if (x == -y && x == z) { // 左右两边可以拼接
                ans = max(ans, pre[i - 1] + suf[i + 1]);
            } else {
                if (x == -y) {
                    ans = max(ans, pre[i - 1] + 1); // 只拼接 nums[i+1]
                }
                if (z == -y) {
                    ans = max(ans, suf[i + 1] + 1); // 只拼接 nums[i-1]
                }
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func calc(a []int) []int {
	f := make([]int, len(a)) // f[i] 表示以 i 结尾的最长交替子数组的长度
	for i, x := range a {
		if i == 0 || a[i-1] == x {
			f[i] = 1
		} else if i > 1 && a[i-2] != a[i-1] && (a[i-2] < a[i-1]) == (a[i-1] > x) {
			f[i] = f[i-1] + 1
		} else {
			f[i] = 2
		}
	}
	return f
}

func longestAlternating(nums []int) int {
	n := len(nums)
	pre := calc(nums) // pre[i] 表示以 i 结尾的最长交替子数组的长度

	slices.Reverse(nums)
	suf := calc(nums) // suf[i] 表示以 i 开头的最长交替子数组的长度
	slices.Reverse(suf)
	slices.Reverse(nums)

	// 不删除元素时的最长交替子数组的长度
	ans := slices.Max(pre)

	// 枚举删除 nums[i]
	for i := 1; i < n-1; i++ {
		if nums[i-1] == nums[i+1] { // 无法拼接
			continue
		}

		// 计算 (i-2,i-1), (i-1,i+1), (i+1,i+2) 的大小关系
		x := 0
		if i > 1 {
			x = cmp.Compare(nums[i-2], nums[i-1])
		}

		y := cmp.Compare(nums[i-1], nums[i+1])

		z := 0
		if i < n-2 {
			z = cmp.Compare(nums[i+1], nums[i+2])
		}

		if x == -y && x == z { // 左右两边可以拼接
			ans = max(ans, pre[i-1]+suf[i+1])
		} else {
			if x == -y {
				ans = max(ans, pre[i-1]+1) // 只拼接 nums[i+1] 
			}
			if z == -y {
				ans = max(ans, suf[i+1]+1) // 只拼接 nums[i-1] 
			}
		}
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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
