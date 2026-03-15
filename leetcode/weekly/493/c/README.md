推荐先完成本题的简单版本：[3738. 替换至多一个元素后最长非递减子数组](https://leetcode.cn/problems/longest-non-decreasing-subarray-after-replacing-at-most-one-element/)。

为方便描述，下文把 $\textit{nums}$ 简称为 $a$。

假设修改的是 $a[i]$，我们需要知道什么？

- 以 $i-1$ 结尾的最长等差子数组的长度，记作 $\textit{pre}[i-1]$。
- 以 $i+1$ 开头的最长等差子数组的长度，记作 $\textit{suf}[i+1]$。

这可以递推算出。以 $\textit{pre}$ 为例：

- $\textit{pre}[0] = 1$。
- 如果 $i = 1$ 或者 $a[i] - a[i-1] \ne a[i-1] - a[i-2]$，那么 $\textit{pre}[i]$ 只能是 $2$，也就是把 $a[i-1]$ 和 $a[i]$ 组成长为 $2$ 等差子数组。
- 否则可以把 $a[i]$ 拼接在以 $i-1$ 结尾的最长等差子数组的后面，即 $\textit{pre}[i] = \textit{pre}[i-1] + 1$。

然后来计算答案。

枚举修改的元素是 $a[i]$。修改后，如果满足如下条件，那么可以把左右两侧的等差子数组拼在一起：

- 设 $d = \dfrac{a[i+1] - a[i-1]}{2}$，这必须是个整数（题目要求修改后的数是整数）。把 $a[i]$ 改成 $d$。
- $a[i-1] - a[i-2] = d = a[i+2] - a[i+1]$。

拼接后的等差子数组的长度为

$$
\textit{pre}[i-1] + 1 + \textit{suf}[i+1]
$$

用上式更新答案的最大值。

还有四种特殊情况：

- 只满足 $a[i-1] - a[i-2] = d$，拼接后的等差子数组的长度为 $\textit{pre}[i-1] + 2$。
- 只满足 $a[i+2] - a[i+1] = d$，拼接后的等差子数组的长度为 $\textit{suf}[i+1] + 2$。
- 修改 $a[i]$，拼在 $\textit{pre}[i-1]$ 的后面，拼接后的等差子数组的长度为 $\textit{pre}[i-1] + 1$。
- 修改 $a[i]$，拼在 $\textit{suf}[i+1]$ 的前面，拼接后的等差子数组的长度为 $\textit{suf}[i+1] + 1$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def calc(self, nums: List[int]) -> List[int]:
        n = len(nums)
        pre = [0] * n
        pre[0] = 1
        pre[1] = 2
        for i in range(2, n):
            if nums[i - 2] + nums[i] == 2 * nums[i - 1]:  # 三个数等差
                pre[i] = pre[i - 1] + 1
            else:
                pre[i] = 2
        return pre

    def longestArithmetic(self, nums: List[int]) -> int:
        n = len(nums)
        pre = self.calc(nums)
        ans = max(pre) + 1
        if ans >= n:  # 整个数组是等差的，或者修改端点元素后是等差的
            return n

        suf = self.calc(nums[::-1])[::-1]
        # 注意 max(pre) == max(suf)，无需重复计算

        for i in range(1, n - 1):
            # 把 nums[i] 改成 d2 / 2
            d2 = nums[i + 1] - nums[i - 1]
            if d2 % 2:
                # d2 / 2 必须是整数
                continue

            ok_left = i > 1 and nums[i - 1] - nums[i - 2] == d2 // 2
            ok_right = i + 2 < n and nums[i + 2] - nums[i + 1] == d2 // 2

            if ok_left and ok_right:
                ans = max(ans, pre[i - 1] + 1 + suf[i + 1])
            elif ok_left:
                ans = max(ans, pre[i - 1] + 2)
            elif ok_right:
                ans = max(ans, suf[i + 1] + 2)

        return ans
```

```java [sol-Java]
class Solution {
    public int longestArithmetic(int[] nums) {
        int[] pre = calc(nums);
        int ans = 0;
        for (int x : pre) {
            ans = Math.max(ans, x);
        }
        ans++;

        int n = nums.length;
        if (ans >= n) { // 整个数组是等差的，或者修改端点元素后是等差的
            return n;
        }

        reverse(nums);
        int[] suf = calc(nums);
        reverse(suf);
        reverse(nums);
        // 注意 max(pre) == max(suf)，无需重复计算

        for (int i = 1; i < n - 1; i++) {
            // 把 nums[i] 改成 d2 / 2
            int d2 = nums[i + 1] - nums[i - 1];
            if (d2 % 2 != 0) { // d2 / 2 必须是整数
                continue;
            }

            boolean okLeft = i > 1 && nums[i - 1] - nums[i - 2] == d2 / 2;
            boolean okRight = i + 2 < n && nums[i + 2] - nums[i + 1] == d2 / 2;

            if (okLeft && okRight) {
                ans = Math.max(ans, pre[i - 1] + 1 + suf[i + 1]);
            } else if (okLeft) {
                ans = Math.max(ans, pre[i - 1] + 2);
            } else if (okRight) {
                ans = Math.max(ans, suf[i + 1] + 2);
            }
        }

        return ans;
    }

    private int[] calc(int[] nums) {
        int n = nums.length;
        int[] pre = new int[n];
        pre[0] = 1;
        pre[1] = 2;
        for (int i = 2; i < n; i++) {
            if (nums[i - 2] + nums[i] == 2 * nums[i - 1]) { // 三个数等差
                pre[i] = pre[i - 1] + 1;
            } else {
                pre[i] = 2;
            }
        }
        return pre;
    }

    private void reverse(int[] a) {
        for (int i = 0, j = a.length - 1; i < j; i++, j--) {
            int tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
    vector<int> calc(vector<int>& nums) {
        int n = nums.size();
        vector<int> pre(n);
        pre[0] = 1;
        pre[1] = 2;
        for (int i = 2; i < n; i++) {
            if (nums[i - 2] + nums[i] == 2 * nums[i - 1]) { // 三个数等差
                pre[i] = pre[i - 1] + 1;
            } else {
                pre[i] = 2;
            }
        }
        return pre;
    }

public:
    int longestArithmetic(vector<int>& nums) {
        int n = nums.size();
        vector<int> pre = calc(nums);
        int ans = ranges::max(pre) + 1;
        if (ans >= n) { // 整个数组是等差的，或者修改端点元素后是等差的
            return n;
        }

        ranges::reverse(nums);
        vector<int> suf = calc(nums);
        ranges::reverse(suf);
        ranges::reverse(nums);
        // 注意 max(pre) == max(suf)，无需重复计算

        for (int i = 1; i < n - 1; i++) {
            // 把 nums[i] 改成 d2 / 2
            int d2 = nums[i + 1] - nums[i - 1];
            if (d2 % 2) { // d2 / 2 必须是整数
                continue;
            }

            bool ok_left = i > 1 && nums[i - 1] - nums[i - 2] == d2 / 2;
            bool ok_right = i + 2 < n && nums[i + 2] - nums[i + 1] == d2 / 2;

            if (ok_left && ok_right) {
                ans = max(ans, pre[i - 1] + 1 + suf[i + 1]);
            } else if (ok_left) {
                ans = max(ans, pre[i - 1] + 2);
            } else if (ok_right) {
                ans = max(ans, suf[i + 1] + 2);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func calc(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	pre[0] = 1
	pre[1] = 2
	for i := 2; i < n; i++ {
		if nums[i-2]+nums[i] == nums[i-1]*2 { // 三个数等差
			pre[i] = pre[i-1] + 1
		} else {
			pre[i] = 2
		}
	}
	return pre
}

func longestArithmetic(nums []int) (ans int) {
	n := len(nums)
	pre := calc(nums)
	ans = slices.Max(pre) + 1
	if ans >= n { // 整个数组是等差的，或者修改端点元素后是等差的
		return n
	}

	slices.Reverse(nums)
	suf := calc(nums)
	slices.Reverse(suf)
	slices.Reverse(nums)
	// 注意 max(pre) == max(suf)，无需重复计算

	for i := 1; i < n-1; i++ {
		// 把 nums[i] 改成 d2/2
		d2 := nums[i+1] - nums[i-1]
		if d2%2 != 0 { // d2/2 必须是整数
			continue
		}

		okLeft := i > 1 && nums[i-1]-nums[i-2] == d2/2
		okRight := i+2 < n && nums[i+2]-nums[i+1] == d2/2

		if okLeft && okRight {
			ans = max(ans, pre[i-1]+1+suf[i+1])
		} else if okLeft {
			ans = max(ans, pre[i-1]+2)
		} else if okRight {
			ans = max(ans, suf[i+1]+2)
		}
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

[3830. 移除至多一个元素后的最长交替子数组](https://leetcode.cn/problems/longest-alternating-subarray-after-removing-at-most-one-element/)

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
