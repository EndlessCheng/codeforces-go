本题有三个要求：

1. $0\le i < j < n$。这可以用**枚举右维护左**，枚举 $j$，维护左边的 $i$。具体怎么维护下面细讲。
2. $\textit{nums}[i] < \textit{nums}[j]$。这可以用**值域分治**。例如元素范围 $[1,100]$，我们以 $50$ 为界，把问题分成：
    - $1 \le \textit{nums}[i] < \textit{nums}[j]\le 50$ 的影子对个数，这是个规模更小的子问题，可以递归（分治）解决。
    - $51 \le \textit{nums}[i] < \textit{nums}[j]\le 100$ 的影子对个数，这是个规模更小的子问题，可以递归（分治）解决。
    - $1\le \textit{nums}[i] \le 50 < \textit{nums}[j]\le 100$ 的影子对个数。下面细讲。
3. 在 $[i+1,j-1]$ 内不能有下标 $k$ 满足 $\textit{nums}[i] < \textit{nums}[k] < \textit{nums}[j]$。

如何计算 $1\le \textit{nums}[i] \le 50 < \textit{nums}[j]\le 100$ 的影子对个数？

把 $[1,50]$ 叫做下部，$[51,100]$ 叫做上部。$\textit{nums}[i]$ 在下部，$\textit{nums}[j]$ 在上部。

从左到右遍历 $\textit{nums}$，对于一个固定的 $\textit{nums}[j]$，什么样的 $\textit{nums}[i]$ 满足要求？

分类讨论：

- 如果 $\textit{nums}[k]$ 在下部。用一个栈 $\textit{lowSt}$ 维护**下部**遍历过的元素，如果当前元素 $x$ 比栈顶大，那么栈顶永远不能作为 $\textit{nums}[i]$（因为 $\textit{nums}[i] < x < \textit{nums}[j]$），弹出栈顶。弹出这些元素后，$\textit{lowSt}$ 从栈底到栈顶是递减的（可以相等），没有干扰我们的 $\textit{nums}[k]$，栈中每个数都适合作为 $\textit{nums}[i]$。
- 如果 $\textit{nums}[k]$ 在上部。设上部中的 $\textit{nums}[p]$ 是 $\textit{nums}[j]$ 左侧最近的小于 $\textit{nums}[j]$ 的数（这是单调栈的标准应用）。令 $k=p$，那么 $i<p$ 的 $\textit{nums}[i]$ 不满足题目的第三个要求。于是，只有 $\textit{lowSt}$ 中的下标 $\ge p$ 的 $\textit{nums}[i]$，才能与 $\textit{nums}[j]$ 组成影子对。$\textit{lowSt}$ 保存下标，我们在 $\textit{lowSt}$ 中二分 $p$，即可求出满足要求的下标 $i$ 的个数。

具体例子请看 [本题视频讲解](https://www.bilibili.com/video/BV1k7Yv6WE3i/?t=21m30s) 中画的图。欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def shadowPairs(self, nums: list[int]) -> int:
        def solve(a: list[int], low: int, high: int) -> None:
            n = len(a)
            if n <= 1 or low == high:
                return

            nonlocal ans
            mid = (low + high) // 2
            low_st = []
            high_st = []
            b = []
            c = []

            for i, x in enumerate(a):
                if x <= mid:  # x 在下部，作为 nums[i]
                    while low_st and a[low_st[-1]] < x:
                        low_st.pop()  # 因为 x 的出现，栈顶不能作为 nums[i]
                    low_st.append(i)
                    b.append(x)
                else:  # x 在上部，作为 nums[j]
                    # 找到 x 左侧第一个小于 x 的最近元素，作为 nums[k]
                    while high_st and a[high_st[-1]] >= x:
                        high_st.pop()
                    ans += len(low_st)
                    if high_st:
                        # low_st 中 < high_st[-1] 的下标不能作为 nums[i]
                        ans -= bisect_left(low_st, high_st[-1])
                    high_st.append(i)
                    c.append(x)

            solve(b, low, mid)
            solve(c, mid + 1, high)

        sorted_nums = sorted(set(nums))
        for i, x in enumerate(nums):
            nums[i] = bisect_left(sorted_nums, x)

        ans = 0
        solve(nums, 0, len(sorted_nums) - 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int shadowPairs(int[] nums) {
        int n = nums.length;
        int[] sorted = nums.clone();
        Arrays.sort(sorted);

        List<Integer> a = new ArrayList<>(n);
        for (int x : nums) {
            a.add(Arrays.binarySearch(sorted, x));
        }

        return solve(a, 0, n - 1);
    }

    private int solve(List<Integer> a, int low, int high) {
        int n = a.size();
        if (n <= 1 || low == high) {
            return 0;
        }

        List<Integer> lowSt = new ArrayList<>();
        List<Integer> highSt = new ArrayList<>();
        List<Integer> b = new ArrayList<>();
        List<Integer> c = new ArrayList<>();
        int mid = (low + high) / 2;
        int res = 0;

        for (int i = 0; i < n; i++) {
            int x = a.get(i);
            if (x <= mid) { // x 在下部，作为 nums[i]
                while (!lowSt.isEmpty() && a.get(lowSt.getLast()) < x) {
                    lowSt.removeLast(); // 因为 x 的出现，栈顶不能作为 nums[i]
                }
                lowSt.add(i);
                b.add(x);
            } else { // x 在上部，作为 nums[j]
                // 找到 x 左侧第一个小于 x 的最近元素，作为 nums[k]
                while (!highSt.isEmpty() && a.get(highSt.getLast()) >= x) {
                    highSt.removeLast();
                }
                res += lowSt.size();
                if (!highSt.isEmpty()) {
                    // lowSt 中 < highSt.getLast() 的下标不能作为 nums[i]
                    int p = Collections.binarySearch(lowSt, highSt.getLast());
                    if (p < 0) {
                        p = ~p; // 见 Collections.binarySearch 源码
                    }
                    res -= p;
                }
                highSt.add(i);
                c.add(x);
            }
        }

        return res + solve(b, low, mid) + solve(c, mid + 1, high);
    }
}
```

```cpp [sol-C++]
class Solution {
    int solve(vector<int>& a, int low, int high) {
        int n = a.size();
        if (n <= 1 || low == high) {
            return 0;
        }

        vector<int> low_st, high_st, b, c;
        int mid = (low + high) / 2;
        int res = 0;

        for (int i = 0; i < n; i++) {
            int x = a[i];
            if (x <= mid) { // x 在下部，作为 nums[i]
                while (!low_st.empty() && a[low_st.back()] < x) {
                    low_st.pop_back(); // 因为 x 的出现，栈顶不能作为 nums[i]
                }
                low_st.push_back(i);
                b.push_back(x);
            } else { // x 在上部，作为 nums[j]
                // 找到 x 左侧第一个小于 x 的最近元素，作为 nums[k]
                while (!high_st.empty() && a[high_st.back()] >= x) {
                    high_st.pop_back();
                }
                res += low_st.size();
                if (!high_st.empty()) {
                    // low_st 中 < high_st.back() 的下标不能作为 nums[i]
                    res -= ranges::lower_bound(low_st, high_st.back()) - low_st.begin();
                }
                high_st.push_back(i);
                c.push_back(x);
            }
        }

        return res + solve(b, low, mid) + solve(c, mid + 1, high);
    }

public:
    int shadowPairs(vector<int>& nums) {
        auto sorted_nums = nums;
        ranges::sort(sorted_nums);
        sorted_nums.erase(ranges::unique(sorted_nums).begin(), sorted_nums.end());
        for (int& x : nums) {
            x = ranges::lower_bound(sorted_nums, x) - sorted_nums.begin();
        }

        return solve(nums, 0, sorted_nums.size() - 1);
    }
};
```

```go [sol-Go]
func solve(a []int, low, high int) (res int) {
	n := len(a)
	if n <= 1 || low == high {
		return
	}

	var lowSt, highSt, b, c []int
	mid := (low + high) / 2

	for i, x := range a {
		if x <= mid { // x 在下部，作为 nums[i]
			for len(lowSt) > 0 && a[lowSt[len(lowSt)-1]] < x {
				lowSt = lowSt[:len(lowSt)-1] // 因为 x 的出现，栈顶不能作为 nums[i]
			}
			lowSt = append(lowSt, i)
			b = append(b, x)
		} else { // x 在上部，作为 nums[j]
			// 找到 x 左侧第一个小于 x 的最近元素，作为 nums[k]
			for len(highSt) > 0 && a[highSt[len(highSt)-1]] >= x {
				highSt = highSt[:len(highSt)-1]
			}
			res += len(lowSt)
			if len(highSt) > 0 {
				// lowSt 中 < highSt[len(highSt)-1] 的下标不能作为 nums[i]
				res -= sort.SearchInts(lowSt, highSt[len(highSt)-1])
			}
			highSt = append(highSt, i)
			c = append(c, x)
		}
	}

	return res + solve(b, low, mid) + solve(c, mid+1, high)
}

func shadowPairs(nums []int) int {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	for i, x := range nums {
		nums[i] = sort.SearchInts(sorted, x)
	}

	return solve(nums, 0, len(sorted)-1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log^2 n)$，其中 $n$ 是 $\textit{nums}$ 的长度。分治有 $\mathcal{O}(\log n)$ 层，每层有 $n$ 个数，所以我们一共做了 $\mathcal{O}(n\log n)$ 次二分查找，总的时间复杂度为 $\mathcal{O}(n\log^2 n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相关内容

力扣又引入新的知识点了，难道后面要出 **CDQ 分治**了？

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
