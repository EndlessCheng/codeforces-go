请先看 [没有思路？一个视频讲透滑动窗口！【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

本题算法如下：

1. 设 $\textit{mx} = \max(\textit{nums})$。
2. 右端点 $\textit{right}$ 从左到右遍历 $\textit{nums}$。遍历到元素 $x=\textit{nums}[\textit{right}]$ 时，如果 $x=\textit{mx}$，就把计数器 $\textit{cntMx}$ 加一。
3. 如果此时 $\textit{cntMx}=k$，则不断右移左指针 $\textit{left}$，直到窗口内的 $\textit{mx}$ 的出现次数**小于** $k$ 为止。此时，对于右端点为 $\textit{right}$ 且左端点小于 $\textit{left}$ 的子数组，$\textit{mx}$ 的出现次数都至少为 $k$，把答案增加 $\textit{left}$。

例如示例 1，当右端点移到第二个 $3$ 时，左端点移到 $2$，此时 $[1,3,2,3]$ 和 $[3,2,3]$ 是满足要求的。当右端点移到第三个 $3$ 时，左端点也移到第三个 $3$，此时 $[1,3,2,3,3], [3,2,3,3], [2,3,3], [3,3]$ 都是满足要求的。所以答案为 $2+4=6$。

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        mx = max(nums)
        ans = cnt_mx = left = 0
        for x in nums:
            if x == mx:
                cnt_mx += 1
            while cnt_mx == k:
                if nums[left] == mx:
                    cnt_mx -= 1
                left += 1
            ans += left
        return ans
```

```java [sol-Java]
class Solution {
    public long countSubarrays(int[] nums, int k) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        long ans = 0;
        int cntMx = 0, left = 0;
        for (int x : nums) {
            if (x == mx) {
                cntMx++;
            }
            while (cntMx == k) {
                if (nums[left++] == mx) {
                    cntMx--;
                }
            }
            ans += left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubarrays(vector<int>& nums, int k) {
        int mx = ranges::max(nums);
        long long ans = 0;
        int cnt_mx = 0, left = 0;
        for (int x : nums) {
            cnt_mx += x == mx;
            while (cnt_mx == k) {
                cnt_mx -= nums[left++] == mx;
            }
            ans += left;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubarrays(nums []int, k int) (ans int64) {
	mx := slices.Max(nums)
	cntMx, left := 0, 0
	for _, x := range nums {
		if x == mx {
			cntMx++
		}
		for cntMx == k {
			if nums[left] == mx {
				cntMx--
			}
			left++
		}
		ans += int64(left)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

把问题改成：子数组的最大值在子数组中至少出现 $k$ 次，要怎么做？（原题是整个数组的最大值，这里是子数组的最大值）

用**单调栈**做，思路可以参考 [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/)

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
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
