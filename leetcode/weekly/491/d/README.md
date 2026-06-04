**问**：某班有 $10$ 个人至少 $20$ 岁，$3$ 个人至少 $21$ 岁，那么恰好 $20$ 岁的人有多少个？

**答**：「至少 $20$ 岁」可以分成「恰好 $20$ 岁」和「至少 $21$ 岁」，所以「至少 $20$ 岁」的人数减去「至少 $21$ 岁」的人数，就是「恰好 $20$ 岁」的人数，即 $10-3=7$。

根据这个思路，本题等价于如下两个问题：

- 子数组**至少**包含 $k$ 个不同整数，且至少有 $k$ 个不同整数都至少出现 $m$ 次。
- 子数组**至少**包含 $k+1$ 个不同整数，且至少有 $k$ 个不同整数都至少出现 $m$ 次。（为什么是至少有 $k$ 个而不是 $k+1$ 个？解答见 [本题视频讲解](https://www.bilibili.com/video/BV1V4PMzrEYG/?t=30m56s)）

二者相减，所表达的含义就是**恰好**包含 $k$ 个不同整数，且至少有 $k$ 个不同整数都至少出现 $m$ 次。

对于每个问题，由于子数组越长越满足要求，越短越不满足要求，有单调性，所以可以用**滑动窗口**解决。如果你不了解滑动窗口，可以看视频[【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

如果你之前没有做过统计子数组个数的滑动窗口，推荐先完成 [2962. 统计最大元素出现至少 K 次的子数组](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/)（[我的题解](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/solutions/2560940/hua-dong-chuang-kou-fu-ti-dan-pythonjava-xvwg/)），这也是一道至少+统计个数的问题，且比本题要简单许多。在这篇题解中，我详细解释了 `ans += left` 这行代码的含义。

[本题视频讲解](https://www.bilibili.com/video/BV1V4PMzrEYG/?t=30m56s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countSubarrays(self, nums, k: int, m: int) -> int:
        def calc(distinct_limit: int) -> int:
            cnt = defaultdict(int)
            ge_m = 0  # 窗口中的出现次数 >= m 的元素个数
            ans = left = 0
            for x in nums:
                # 1. 入
                cnt[x] += 1
                if cnt[x] == m:
                    ge_m += 1

                # 2. 出
                while len(cnt) >= distinct_limit and ge_m >= k:
                    out = nums[left]
                    if cnt[out] == m:
                        ge_m -= 1
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

    private long calc(int[] nums, int distinctLimit, int k, int m) {
        Map<Integer, Integer> cnt = new HashMap<>();
        int geM = 0; // 窗口中的出现次数 >= m 的元素个数
        int left = 0;
        long ans = 0;
        for (int x : nums) {
            // 1. 入
            int c = cnt.merge(x, 1, Integer::sum); // c = ++cnt[x]
            if (c == m) {
                geM++;
            }

            // 2. 出
            while (cnt.size() >= distinctLimit && geM >= k) {
                int out = nums[left];
                c = cnt.get(out);
                if (c == m) {
                    geM--;
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
        auto calc = [&](int distinct_limit) -> long long {
            unordered_map<int, int> cnt;
            int ge_m = 0; // 窗口中的出现次数 >= m 的元素个数
            int left = 0;
            long long ans = 0;
            for (int x : nums) {
                // 1. 入
                if (++cnt[x] == m) {
                    ge_m++;
                }

                // 2. 出
                while (cnt.size() >= distinct_limit && ge_m >= k) {
                    int out = nums[left];
                    if (cnt[out] == m) {
                        ge_m--;
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
	calc := func(distinctLimit int) (ans int64) {
		cnt := map[int]int{}
		geM := 0 // 窗口中的出现次数 >= m 的元素个数
		left := 0
		for _, x := range nums {
			// 1. 入
			cnt[x]++
			if cnt[x] == m {
				geM++
			}

			// 2. 出
			for len(cnt) >= distinctLimit && geM >= k {
				out := nums[left]
				if cnt[out] == m {
					geM--
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
