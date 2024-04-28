## 提示 1：二分答案

套路，见[【题单】二分答案](https://leetcode.cn/circle/discuss/SqopEo/) 中的「第 K 小/大」。

一共有 $m = \dfrac{n(n+1)}{2}$ 个非空连续子数组，中位数是其中第 $k = \left\lceil\dfrac{m}{2}\right\rceil$ 个。

二分中位数 $\textit{upper}$，问题变成：

- $\texttt{distinct}$ 值 $\le \textit{upper}$ 的子数组有多少个？

设子数组的个数为 $\textit{cnt}$，如果 $\textit{cnt} < k$ 说明二分的 $\textit{upper}$ 小了，更新二分左边界 $\textit{left}$，否则更新二分右边界 $\textit{right}$。

## 提示 2：滑动窗口

怎么计算 $\texttt{distinct}$ 值 $\le \textit{upper}$ 的子数组个数？

这又是一个套路，见[【题单】滑动窗口](https://leetcode.cn/circle/discuss/0viNMK/) 中的「不定长滑动窗口（求子数组个数）」，类似 [713. 乘积小于 K 的子数组](https://leetcode.cn/problems/subarray-product-less-than-k/)。

由于子数组越长，不同元素个数（$\texttt{distinct}$ 值）不会变小，这样的**单调性**可以让我们滑窗。

用一个哈希表 $\textit{freq}$ 统计窗口（子数组）内的元素及其出现次数。

枚举窗口右端点 $r$，把 $\textit{nums}[r]$ 加入 $\textit{freq}$。如果发现 $\textit{freq}$ 的大小超过 $\textit{upper}$，就不断增大窗口左端点 $l$，直到 $\textit{freq}$ 的大小 $\le \textit{upper}$ 为止。

此时右端点为 $r$，左端点为 $l,l+1,l+2,\cdots,r$ 的子数组都是满足要求的（$\texttt{distinct}$ 值 $\le \textit{upper}$），一共有 $r-l+1$ 个，加入统计的子数组个数 $\textit{cnt}$。

## 其它细节

开区间二分左边界：$0$，一定不满足要求，因为没有子数组的 $\texttt{distinct}$ 值是 $0$。

开区间二分右边界：$n$ 或者 $\textit{nums}$ 中的不同元素个数，一定满足要求，因为所有子数组的 $\texttt{distinct}$ 值不超过 $n$。

## 答疑

**问**：为什么二分出来的答案，一定是某个子数组的 $\texttt{distinct}$ 值？有没有可能，二分出来的答案不是任何子数组的 $\texttt{distinct}$ 值？

**答**：反证法。如果答案 $d$ 不是任何子数组的 $\texttt{distinct}$ 值，那么 $\texttt{distinct}$ 值 $\le d$ 和 $\le d-1$ 算出来的子数组个数是一样的。也就是说 $d-1$ 同样满足要求，即 `check(d - 1) == true`，这与循环不变量相矛盾。

关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)

关于滑动窗口的原理，请看 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

[本题视频讲解](https://www.bilibili.com/video/BV1Pw4m1C79N/)（第四题），欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def medianOfUniquenessArray(self, nums: List[int]) -> int:
        n = len(nums)
        k = (n * (n + 1) // 2 + 1) // 2

        def check(upper: int) -> bool:
            cnt = l = 0
            freq = Counter()
            for r, in_ in enumerate(nums):
                freq[in_] += 1
                while len(freq) > upper:
                    out = nums[l]
                    freq[out] -= 1
                    if freq[out] == 0:
                        del freq[out]
                    l += 1
                cnt += r - l + 1
                if cnt >= k:
                    return True
            return False

        return bisect_left(range(len(set(nums))), True, 1, key=check)
```

```java [sol-Java]
class Solution {
    public int medianOfUniquenessArray(int[] nums) {
        int n = nums.length;
        long k = ((long) n * (n + 1) / 2 + 1) / 2;
        int left = 0;
        int right = n;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            if (check(nums, mid, k)) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }

    private boolean check(int[] nums, int upper, long k) {
        long cnt = 0;
        int l = 0;
        HashMap<Integer, Integer> freq = new HashMap<>();
        for (int r = 0; r < nums.length; r++) {
            freq.merge(nums[r], 1, Integer::sum);
            while (freq.size() > upper) {
                int out = nums[l++];
                if (freq.merge(out, -1, Integer::sum) == 0) {
                    freq.remove(out);
                }
            }
            cnt += r - l + 1;
            if (cnt >= k) {
                return true;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int medianOfUniquenessArray(vector<int>& nums) {
        int n = nums.size();
        long long k = ((long long) n * (n + 1) / 2 + 1) / 2;

        auto check = [&](int upper) {
            long long cnt = 0;
            int l = 0;
            unordered_map<int, int> freq;
            for (int r = 0; r < n; r++) {
                freq[nums[r]]++;
                while (freq.size() > upper) {
                    int out = nums[l++];
                    if (--freq[out] == 0) {
                        freq.erase(out);
                    }
                }
                cnt += r - l + 1;
                if (cnt >= k) {
                    return true;
                }
            }
            return false;
        };

        int left = 0, right = n;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? right : left) = mid;
        }
        return right;
    }
};
```

```go [sol-Go]
func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	k := (n*(n+1)/2 + 1) / 2
	ans := 1 + sort.Search(n-1, func(upper int) bool {
		upper++
		cnt := 0
		l := 0
		freq := map[int]int{}
		for r, in := range nums {
			freq[in]++
			for len(freq) > upper {
				out := nums[l]
				freq[out]--
				if freq[out] == 0 {
					delete(freq, out)
				}
				l++
			}
			cnt += r - l + 1
			if cnt >= k {
				return true
			}
		}
		return false
	})
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。二分 $\mathcal{O}(\log n)$ 次，每次会跑一个 $\mathcal{O}(n)$ 的滑动窗口。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
