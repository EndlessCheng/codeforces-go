请看 [视频讲解](https://www.bilibili.com/video/BV1RH4y1W7DP/) 第四题。

## 提示 1

把 $0$ 看成「空位」。

第二种操作相当于把一个 $1$ 移动到它相邻的空位上，如果我们想得到一个下标在 $j$ 的 $1$，必须操作 $|\textit{index} - j|$ 次。

对于第一种操作，我们可以选择把和 $\textit{index}$ 相邻的 $0$ 变成 $1$（在此之前先移动相邻的 $1$），然后结合第二种操作，把相邻的 $1$ 移动到 $\textit{index}$，总共操作 $2$ 次就可以得到一个 $1$。

## 提示 2

看上去应当优先使用第一种操作+第二种操作（毕竟只需要操作 $2$ 次就能得到一个 $1$），在数组中有 $1$ 且 $\textit{maxChanges}$ 很大的情况下，答案会是 $(k-1)\cdot 2$ 吗？

不对，对于 $\textit{index}, \textit{index}-1, \textit{index}+1$ 这三个位置上的 $1$。我们可以用更少的操作得到：

- $\textit{index}$ 位置上的 $1$ 相当于操作 $0$ 次。
- $\textit{index}-1$ 和 $\textit{index}+1$ 位置上的 $1$ 只需操作 $1$ 次。

设 $c$ 为 $\textit{nums}$ 中的长度不超过 $3$ 的最长连续 $1$ 的长度。如果 $c>k$ 则 $c=k$。

如果 $\textit{maxChanges}\ge k-c$，我们可以先使用 $\max(c-1, 0)$ 次第二种操作，收集这连续的 $c$ 个 $1$，然后对于其余 $k-c$ 个 $1$，都可以用 $2$ 次操作得到，此时可以直接返回 $\max(c-1, 0) + (k-c)\cdot 2$。

接下来，要解决的就是 $\textit{maxChanges}$ 比较小的情况了。

想一想，如果 $\textit{maxChanges}=0$，也就是只能使用第二种操作，要如何计算答案呢？

## 提示 3

这是一个标准的「货仓选址」问题。我们首先算出所有 $1$ 的位置，记到数组 $\textit{pos}$ 中，例如示例 1 的 $\textit{nums} = [1,1,0,0,0,1,1,0,0,1]$，其 $\textit{pos}=[0, 1, 5, 6, 9]$。

如果 $\textit{maxChanges}=0$，我们可以枚举所有长为 $3$ 的子数组，例如 $[0,1,5]$，就好比在坐标轴上的 $0,1,5$ 位置上有三个货仓，把工厂建在哪里，可以使得货仓到工厂的距离和最小？

根据 [中位数贪心及其证明](https://leetcode.cn/problems/5TxKeK/solution/zhuan-huan-zhong-wei-shu-tan-xin-dui-din-7r9b/)，最优解是把工厂建在货仓的中位数上。例如 $[0,1,5]$ 中的 $1$，此时距离和等于 $|0-1|+|1-1|+|5-1| = 5$。

利用前缀和，可以 $\mathcal{O}(1)$ 算出子数组中的所有数到其中位数的距离之和，原理请看 [图解](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solution/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/)。

最后，如果 $\textit{maxChanges}>0$，我们可以先计算所有长为 $k - \textit{maxChanges}$ 的子数组的货仓选址问题，然后再通过 $\textit{maxChanges}\cdot 2$ 次操作得到 $\textit{maxChanges}$ 个 $1$。

注：也有不用前缀和的滑窗做法，具体见 2968 题的 [题解方法二](https://leetcode.cn/problems/apply-operations-to-maximize-frequency-score/solution/hua-dong-chuang-kou-zhong-wei-shu-tan-xi-nuvr/)。

```py [sol-Python3]
class Solution:
    def minimumMoves(self, nums: List[int], k: int, max_changes: int) -> int:
        pos = []
        c = 0  # nums 中连续的 1 长度
        for i, x in enumerate(nums):
            if x == 0:
                continue
            pos.append(i)  # 记录 1 的位置
            c = max(c, 1)
            if i > 0 and nums[i - 1] == 1:
                if i > 1 and nums[i - 2] == 1:
                    c = 3  # 有 3 个连续的 1
                else:
                    c = max(c, 2)  # 有 2 个连续的 1

        c = min(c, k)
        if max_changes >= k - c:
            # 其余 k-c 个 1 可以全部用两次操作得到
            return max(c - 1, 0) + (k - c) * 2

        n = len(pos)
        pre_sum = list(accumulate(pos, initial=0))

        ans = inf
        # 除了 max_changes 个数可以用两次操作得到，其余的 1 只能一步步移动到 pos[i]
        size = k - max_changes
        for right in range(size, n + 1):
            # s1+s2 是 j 在 [left, right) 中的所有 pos[j] 到 pos[(left+right)/2] 的距离之和
            left = right - size
            i = left + size // 2
            s1 = pos[i] * (i - left) - (pre_sum[i] - pre_sum[left])
            s2 = pre_sum[right] - pre_sum[i] - pos[i] * (right - i)
            ans = min(ans, s1 + s2)
        return ans + max_changes * 2
```

```java [sol-Java]
class Solution {
    public long minimumMoves(int[] nums, int k, int maxChanges) {
        List<Integer> pos = new ArrayList<>();
        int c = 0; // nums 中连续的 1 长度
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 0) continue;
            pos.add(i); // 记录 1 的位置
            c = Math.max(c, 1);
            if (i > 0 && nums[i - 1] == 1) {
                if (i > 1 && nums[i - 2] == 1) {
                    c = 3; // 有 3 个连续的 1
                } else {
                    c = Math.max(c, 2); // 有 2 个连续的 1
                }
            }
        }

        c = Math.min(c, k);
        if (maxChanges >= k - c) {
            // 其余 k-c 个 1 可以全部用两次操作得到
            return Math.max(c - 1, 0) + (k - c) * 2;
        }

        int n = pos.size();
        long[] sum = new long[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + pos.get(i);
        }

        long ans = Long.MAX_VALUE;
        // 除了 maxChanges 个数可以用两次操作得到，其余的 1 只能一步步移动到 pos[i]
        int size = k - maxChanges;
        for (int right = size; right <= n; right++) {
            // s1+s2 是 j 在 [left, right) 中的所有 pos[j] 到 index=pos[(left+right)/2] 的距离之和
            int left = right - size;
            int i = left + size / 2;
            long index = pos.get(i);
            long s1 = index * (i - left) - (sum[i] - sum[left]);
            long s2 = sum[right] - sum[i] - index * (right - i);
            ans = Math.min(ans, s1 + s2);
        }
        return ans + maxChanges * 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumMoves(vector<int> &nums, int k, int maxChanges) {
        vector<int> pos;
        int c = 0; // nums 中连续的 1 长度
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] == 0) continue;
            pos.push_back(i); // 记录 1 的位置
            c = max(c, 1);
            if (i > 0 && nums[i - 1] == 1) {
                if (i > 1 && nums[i - 2] == 1) {
                    c = 3; // 有 3 个连续的 1
                } else {
                    c = max(c, 2); // 有 2 个连续的 1
                }
            }
        }

        c = min(c, k);
        if (maxChanges >= k - c) {
            // 其余 k-c 个 1 可以全部用两次操作得到
            return max(c - 1, 0) + (k - c) * 2;
        }

        int n = pos.size();
        vector<long long> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + pos[i];
        }

        long long ans = LLONG_MAX;
        // 除了 maxChanges 个数可以用两次操作得到，其余的 1 只能一步步移动到 pos[i]
        int size = k - maxChanges;
        for (int right = size; right <= n; right++) {
            // s1+s2 是 j 在 [left, right) 中的所有 pos[j] 到 index=pos[(left+right)/2] 的距离之和
            int left = right - size;
            int i = left + size / 2;
            long long index = pos[i];
            long long s1 = index * (i - left) - (sum[i] - sum[left]);
            long long s2 = sum[right] - sum[i] - index * (right - i);
            ans = min(ans, s1 + s2);
        }
        return ans + maxChanges * 2;
    }
};
```

```go [sol-Go]
func minimumMoves(nums []int, k, maxChanges int) int64 {
	pos := []int{}
	c := 0 // nums 中连续的 1 长度
	for i, x := range nums {
		if x == 0 {
			continue
		}
		pos = append(pos, i) // 记录 1 的位置
		c = max(c, 1)
		if i > 0 && nums[i-1] == 1 {
			if i > 1 && nums[i-2] == 1 {
				c = 3 // 有 3 个连续的 1
			} else {
				c = max(c, 2) // 有 2 个连续的 1
			}
		}
	}

	c = min(c, k)
	if maxChanges >= k-c {
		// 其余 k-c 个 1 可以全部用两次操作得到
		return int64(max(c-1, 0) + (k-c)*2)
	}

	n := len(pos)
	sum := make([]int, n+1)
	for i, x := range pos {
		sum[i+1] = sum[i] + x
	}

	ans := math.MaxInt
	// 除了 maxChanges 个数可以用两次操作得到，其余的 1 只能一步步移动到 pos[i]
	size := k - maxChanges
	for right := size; right <= n; right++ {
		// s1+s2 是 j 在 [left, right) 中的所有 pos[j] 到 pos[(left+right)/2] 的距离之和
		left := right - size
		i := left + size/2
		s1 := pos[i]*(i-left) - (sum[i] - sum[left])
		s2 := sum[right] - sum[i] - pos[i]*(right-i)
		ans = min(ans, s1+s2)
	}
	return int64(ans + maxChanges*2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 题单：中位数贪心

- [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
- [2033. 获取单值网格的最小操作数](https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid/) 1672
- [2448. 使数组相等的最小开销](https://leetcode.cn/problems/minimum-cost-to-make-array-equal/) 2005
- [2607. 使子数组元素和相等](https://leetcode.cn/problems/make-k-subarray-sums-equal/) 2071
- [2967. 使数组成为等数数组的最小代价](https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic/) 2116
- [1478. 安排邮筒](https://leetcode.cn/problems/allocate-mailboxes/) 2190
- [2968. 执行操作使频率分数最大](https://leetcode.cn/problems/apply-operations-to-maximize-frequency-score/) 2444
- [1703. 得到连续 K 个 1 的最少相邻交换次数](https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/) 2467
- [LCP 24. 数字游戏](https://leetcode.cn/problems/5TxKeK/)
- [296. 最佳的碰头地点](https://leetcode.cn/problems/best-meeting-point/) 二维的情况（会员题）

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
