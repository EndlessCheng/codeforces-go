**问**：某班有 $10$ 个人至少 $20$ 岁，$3$ 个人至少 $21$ 岁，那么恰好 $20$ 岁的人有多少个？

**答**：「至少 $20$ 岁」可以分成「恰好 $20$ 岁」和「至少 $21$ 岁」，所以「至少 $20$ 岁」的人数减去「至少 $21$ 岁」的人数，就是「恰好 $20$ 岁」的人数，即 $10-3=7$。

把「出现至少 $m$ 次的整数」叫做**好整数**。题目要求子数组恰好有 $k$ 个不同的整数，其中恰好有 $k$ 个不同的好整数。

一个**错误的想法**是，把问题拆分为：

- 子数组至少有 $k$ 个不同的整数，其中至少有 $k$ 个不同的好整数（下图的三个矩形）。
- 子数组至少有 $k+1$ 个不同的整数，其中至少有 $k+1$ 个不同的好整数（下图右上角的矩形）。
- 二者相减，即为答案。

**错误原因**：相减后，我们把「子数组至少有 $k+1$ 个不同的整数，其中恰好有 $k$ 个不同的好整数」也算进来了（下图右下角的矩形）。

![lc3859-c.png](https://pic.leetcode.cn/1788955503-pWfOlg-lc3859-c.png){:width=700px}

> 注：左上空缺的是「子数组恰有 $k$ 个不同的整数，其中至少有 $k+1$ 个不同的好整数」，不存在这样的子数组。

由上图可知，问题应当拆分为：

- 子数组至少有 $k$ 个不同的整数，其中至少有 $k$ 个不同的好整数（上图的三个矩形）。
- 子数组至少有 $k+1$ 个不同的整数，其中至少有 $k$ 个不同的好整数（上图右边的两个矩形）。
- 二者相减，即为答案。

对于每个问题，由于子数组越长越满足要求，越短越不满足要求，有单调性，所以可以用**滑动窗口**解决。如果你不了解滑动窗口，可以看视频[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

如果你之前没有做过统计子数组个数的滑动窗口，推荐先完成 [2962. 统计最大元素出现至少 K 次的子数组](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/)（[我的题解](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/solutions/2560940/hua-dong-chuang-kou-fu-ti-dan-pythonjava-xvwg/)），这也是一道至少+统计个数的问题，且比本题要简单许多。在这篇题解中，我详细解释了 `ans += left` 这行代码的含义。

[本题视频讲解](https://www.bilibili.com/video/BV1V4PMzrEYG/?t=30m56s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums, k: int, m: int) -> int:
        # 子数组至少有 distinct_limit 个不同的整数，其中至少有 k 个不同的好整数
        def calc(distinct_limit: int) -> int:
            cnt = defaultdict(int)  # 用来统计不同元素个数
            good_numbers = 0  # 窗口中的好整数（出现至少 m 次的整数）的个数
            ans = left = 0
            for x in nums:
                # 1. 入
                cnt[x] += 1
                if cnt[x] == m:
                    good_numbers += 1

                # 2. 出
                while len(cnt) >= distinct_limit and good_numbers >= k:
                    out = nums[left]
                    if cnt[out] == m:
                        good_numbers -= 1
                    cnt[out] -= 1
                    if cnt[out] == 0:
                        del cnt[out]
                    left += 1

                # 3. 更新答案
                ans += left
            return ans

        return calc(k) - calc(k + 1)
```

```java [sol-Java]
class Solution {
    public long countSubarrays(int[] nums, int k, int m) {
        return calc(nums, k, k, m) - calc(nums, k + 1, k, m);
    }

    // 子数组至少有 distinctLimit 个不同的整数，其中至少有 k 个不同的好整数
    private long calc(int[] nums, int distinctLimit, int k, int m) {
        Map<Integer, Integer> cnt = new HashMap<>(); // 用来统计不同元素个数
        int goodNumbers = 0; // 窗口中的好整数（出现至少 m 次的整数）的个数
        int left = 0;
        long ans = 0;
        for (int x : nums) {
            // 1. 入
            int c = cnt.merge(x, 1, Integer::sum); // c = ++cnt[x]
            if (c == m) {
                goodNumbers++;
            }

            // 2. 出
            while (cnt.size() >= distinctLimit && goodNumbers >= k) {
                int out = nums[left];
                c = cnt.get(out);
                if (c == m) {
                    goodNumbers--;
                }
                if (c == 1) {
                    cnt.remove(out);
                } else {
                    cnt.put(out, c - 1);
                }
                left++;
            }

            // 3. 更新答案
            ans += left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubarrays(vector<int>& nums, int k, int m) {
        // 子数组至少有 distinct_limit 个不同的整数，其中至少有 k 个不同的好整数
        auto calc = [&](int distinct_limit) -> long long {
            unordered_map<int, int> cnt; // 用来统计不同元素个数
            int good_numbers = 0; // 窗口中的好整数（出现至少 m 次的整数）的个数
            int left = 0;
            long long ans = 0;
            for (int x : nums) {
                // 1. 入
                if (++cnt[x] == m) {
                    good_numbers++;
                }

                // 2. 出
                while (cnt.size() >= distinct_limit && good_numbers >= k) {
                    int out = nums[left];
                    if (cnt[out] == m) {
                        good_numbers--;
                    }
                    if (--cnt[out] == 0) {
                        cnt.erase(out);
                    }
                    left++;
                }

                // 3. 更新答案
                ans += left;
            }
            return ans;
        };

        return calc(k) - calc(k + 1);
    }
};
```

```go [sol-Go]
func countSubarrays(nums []int, k, m int) int64 {
	// 子数组至少有 distinctLimit 个不同的整数，其中至少有 k 个不同的好整数
	calc := func(distinctLimit int) (ans int64) {
		cnt := map[int]int{} // 用来统计不同元素个数
		goodNumbers := 0 // 窗口中的好整数（出现至少 m 次的整数）的个数
		left := 0
		for _, x := range nums {
			// 1. 入
			cnt[x]++
			if cnt[x] == m {
				goodNumbers++
			}

			// 2. 出
			for len(cnt) >= distinctLimit && goodNumbers >= k {
				out := nums[left]
				if cnt[out] == m {
					goodNumbers--
				}
				cnt[out]--
				if cnt[out] == 0 {
					delete(cnt, out)
				}
				left++
			}

			// 3. 更新答案
			ans += int64(left)
		}
		return
	}

	return calc(k) - calc(k+1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面滑动窗口题单的「**§2.3.3 恰好型滑动窗口**」。

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
