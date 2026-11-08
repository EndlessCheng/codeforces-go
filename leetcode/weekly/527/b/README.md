## 问题分解

如果把 $\textit{nums}$ 的后缀 $[i,n-1]$ 递减排序后，$\textit{nums}$ 等于 $\textit{target}$，那么必须满足：

1. **大前提**：$\textit{nums}$ 的每个数的出现次数恰好等于 $\textit{target}$ 的每个数的出现次数。
2. $\textit{nums}$ 的前 $i$ 个数 $[0,i-1]$ 必须等于 $\textit{target}$ 的前 $i$ 个数，我们需要知道 $i$ **最大是多少**。
3. $\textit{target}$ 的后 $n-i$ 个数 $[i,n-1]$ 必须是递减的（允许相邻相等），我们需要知道 $n-i$ 最大是多少，即 $i$ **最小是多少**。

如果 1 和 2 都成立，那么 $\textit{target}$ 的后 $n-i$ 个数的出现次数等于 $\textit{nums}$ 的后 $n-i$ 个数的出现次数，只要保证这两部分都是递减的，它们就一定相等。

这样就把原问题拆分成了三个小问题，分别解决。

## 具体思路

第一个问题，用数组（或者哈希表）统计元素出现次数。

第二个问题，从左到右遍历数组，找到第一个 $\textit{nums}[i]\ne \textit{target}[i]$ 的位置，即为 $i$ 的最大值 $\textit{maxI}$。如果不存在则 $\textit{maxI}=n-1$。

第三个问题，设 $\textit{minI}$ 是最大的满足 $\textit{target}[i-1] < \textit{target}[i]$ 下标 $i$。如果不存在则 $\textit{minI} = 0$。那么 $i$ 必须 $\ge \textit{minI}$，才能保证 $\textit{target}$ 的后缀 $[i,n-1]$ 不存在递增的情况（即递减的）。

所以满足要求的 $i$ 的范围是

$$
[\textit{minI}, \textit{maxI}]
$$

答案为

$$
\max(\textit{maxI} - \textit{minI} + 1, 0)
$$

注意最后要和 $0$ 取最大值，以保证结果非负。

[本题视频讲解](https://www.bilibili.com/video/BV1LyW76zEvX/?t=11m4s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countSuffix(self, nums: list[int], target: list[int]) -> int:
        # 大前提：每种元素的出现次数都相等
        if Counter(nums) != Counter(target):
            return 0

        n = len(nums)
        max_i = n - 1
        for i, (x, y) in enumerate(zip(nums, target)):
            if x != y:
                max_i = i
                break

        min_i = 0
        for i in range(n - 1, 0, -1):
            if target[i - 1] < target[i]:
                min_i = i
                break

        return max(max_i - min_i + 1, 0)
```

```java [sol-Java]
class Solution {
    public int countSuffix(int[] nums, int[] target) {
        int n = nums.length;
        int mx = 0;
        for (int x : nums) mx = Math.max(mx, x);
        for (int x : target) mx = Math.max(mx, x);

        // 大前提：每种元素的出现次数都相等
        int[] diff = new int[mx + 1]; // int[] 比 HashMap 快
        for (int i = 0; i < n; i++) {
            diff[nums[i]]++;
            diff[target[i]]--;
        }
        for (int d : diff) {
            if (d != 0) {
                return 0;
            }
        }

        int maxI = n - 1;
        for (int i = 0; i < n; i++) {
            if (nums[i] != target[i]) {
                maxI = i;
                break;
            }
        }

        int minI = 0;
        for (int i = n - 1; i > 0; i--) {
            if (target[i - 1] < target[i]) {
                minI = i;
                break;
            }
        }

        return Math.max(maxI - minI + 1, 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countSuffix(vector<int>& nums, vector<int>& target) {
        int n = nums.size();
        int mx = max(ranges::max(nums), ranges::max(target));

        // 大前提：每种元素的出现次数都相等
        vector<int> diff(mx + 1); // vector 比 unordered_map 快
        for (int i = 0; i < n; i++) {
            diff[nums[i]]++;
            diff[target[i]]--;
        }
        for (int d : diff) {
            if (d) {
                return 0;
            }
        }

        int max_i = n - 1;
        for (int i = 0; i < n; i++) {
            if (nums[i] != target[i]) {
                max_i = i;
                break;
            }
        }

        int min_i = 0;
        for (int i = n - 1; i > 0; i--) {
            if (target[i - 1] < target[i]) {
                min_i = i;
                break;
            }
        }

        return max(max_i - min_i + 1, 0);
    }
};
```

```go [sol-Go]
func countSuffix(nums []int, target []int) int {
	// 大前提：每种元素的出现次数都相等
	mx := max(slices.Max(nums), slices.Max(target))
	diff := make([]int, mx+1) // slice 比 map 快
	for i, x := range nums {
		diff[x]++
		diff[target[i]]--
	}
	for _, d := range diff {
		if d != 0 {
			return 0
		}
	}

	n := len(nums)
	maxI := n - 1
	for i, x := range nums {
		if x != target[i] {
			maxI = i
			break
		}
	}

	minI := 0
	for i := n - 1; i > 0; i-- {
		if target[i-1] < target[i] {
			minI = i
			break
		}
	}

	return max(maxI-minI+1, 0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\max(\textit{nums}),\max(\textit{target}))$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+U)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
