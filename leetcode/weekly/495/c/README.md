## 分析

枚举 $n$ 的因子 $k$，考虑对于一个固定的 $k$，能否满足要求。

要让 $\textit{nums}$ 是递增的，那么首先要满足：

- 分割后，长为 $k$ 的子数组都是递增的。
- 每个子数组的最小值 $\ge $ 上一个子数组的最大值。

对于一个子数组 $a$，如果能通过循环移动变成递增的，那么反过来，一个递增数组，通过循环移动，可以变成什么样的数组？

- 子数组 $a$ 本来就是递增的。
- 或者，子数组 $a$ 由两个递增段组成，且第一段的最小值 $\ge $ 第二段的最大值。 

## 预处理

为了快速判断子数组 $a$ 是否满足要求，我们可以预处理：

- 对于每个下标 $i$，计算下一个递减的位置 $j$。具体地，$j$ 是最小的满足 $j\ge i$ 且 $\textit{nums}[j] > \textit{nums}[j+1]$ 的下标。如果不存在，那么 $j = n$。

把 $j$ 记在数组 $\textit{nextDec}$ 中。

有了 $\textit{nextDec}$，就可以快速找到下一个递增段的位置，从而判断子数组 $a$ 是否满足要求。可以根据 $a$ 是递增的还是两段递增分类讨论，具体见代码。

[本题视频讲解](https://www.bilibili.com/video/BV11UXSB7EGz/?t=44m54s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def sortableIntegers(self, nums: List[int]) -> int:
        n = len(nums)
        next_dec = [0] * n  # nums[next_dec[i]] > nums[next_dec[i] + 1]
        next_dec[-1] = p = n
        # 对于每个 i，记录下一个递减的位置
        for i in range(n - 2, -1, -1):
            if nums[i] > nums[i + 1]:
                p = i
            next_dec[i] = p

        def solve(k: int) -> None:
            last_max = 0  # 上一段的最大值
            for r in range(k - 1, n, k):
                l = r - k + 1
                m = next_dec[l]
                if m >= r:
                    # [l, r] 是递增的，最小值为 nums[l]，最大值为 nums[r]
                    # 最小值必须 >= 上一段的最大值
                    if nums[l] < last_max:
                        return
                    last_max = nums[r]
                else:
                    # [l, m] 是第一段，[m+1, r] 是第二段
                    # 第二段必须是递增的，且第二段的最小值必须 >= 上一段的最大值，且第二段的最大值必须 <= 第一段的最小值
                    if next_dec[m + 1] < r or nums[m + 1] < last_max or nums[r] > nums[l]:
                        return
                    last_max = nums[m]
            nonlocal ans
            ans += k  # 满足要求

        ans = 0
        # 枚举 n 的因子 k
        for k in range(1, isqrt(n) + 1):
            if n % k == 0:
                solve(k)
                if k * k < n:
                    solve(n // k)
        return ans
```

```java [sol-Java]
class Solution {
    public int sortableIntegers(int[] nums) {
        int n = nums.length;
        int[] nextDec = new int[n]; // nums[nextDec[i]] > nums[nextDec[i] + 1]
        nextDec[n - 1] = n;
        int p = n;
        // 对于每个 i，记录下一个递减的位置
        for (int i = n - 2; i >= 0; i--) {
            if (nums[i] > nums[i + 1]) {
                p = i;
            }
            nextDec[i] = p;
        }

        int ans = 0;
        // 枚举 n 的因子 k
        for (int k = 1; k * k <= n; k++) {
            if (n % k == 0) {
                ans += solve(k, nums, nextDec);
                if (k * k < n) {
                    ans += solve(n / k, nums, nextDec);
                }
            }
        }
        return ans;
    }

    private int solve(int k, int[] nums, int[] nextDec) {
        int lastMax = 0; // 上一段的最大值
        for (int r = k - 1; r < nums.length; r += k) {
            int l = r - k + 1;
            int m = nextDec[l];
            if (m >= r) {
                // [l, r] 是递增的，最小值为 nums[l]，最大值为 nums[r]
                // 最小值必须 >= 上一段的最大值
                if (nums[l] < lastMax) {
                    return 0;
                }
                lastMax = nums[r];
            } else {
                // [l, m] 是第一段，[m+1, r] 是第二段
                // 第二段必须是递增的，且第二段的最小值必须 >= 上一段的最大值，且第二段的最大值必须 <= 第一段的最小值
                if (nextDec[m + 1] < r || nums[m + 1] < lastMax || nums[r] > nums[l]) {
                    return 0;
                }
                lastMax = nums[m];
            }
        }
        return k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sortableIntegers(vector<int>& nums) {
        int n = nums.size();
        vector<int> next_dec(n); // nums[next_dec[i]] > nums[next_dec[i] + 1]
        next_dec[n - 1] = n;
        int p = n;
        // 对于每个 i，记录下一个递减的位置
        for (int i = n - 2; i >= 0; i--) {
            if (nums[i] > nums[i + 1]) {
                p = i;
            }
            next_dec[i] = p;
        }

        int ans = 0;

        auto solve = [&](int k) -> void {
            int last_max = 0; // 上一段的最大值
            for (int r = k - 1; r < n; r += k) {
                int l = r - k + 1;
                int m = next_dec[l];
                if (m >= r) {
                    // [l, r] 是递增的，最小值为 nums[l]，最大值为 nums[r]
                    // 最小值必须 >= 上一段的最大值
                    if (nums[l] < last_max) {
                        return;
                    }
                    last_max = nums[r];
                } else {
                    // [l, m] 是第一段，[m+1, r] 是第二段
                    // 第二段必须是递增的，且第二段的最小值必须 >= 上一段的最大值，且第二段的最大值必须 <= 第一段的最小值
                    if (next_dec[m + 1] < r || nums[m + 1] < last_max || nums[r] > nums[l]) {
                        return;
                    }
                    last_max = nums[m];
                }
            }
            ans += k; // 满足要求
        };

        // 枚举 n 的因子 k
        for (int k = 1; k * k <= n; k++) {
            if (n % k == 0) {
                solve(k);
                if (k * k < n) {
                    solve(n / k);
                }
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func sortableIntegers(nums []int) (ans int) {
	n := len(nums)
	nextDec := make([]int, n) // nums[nextDec[i]] > nums[nextDec[i] + 1]
	nextDec[n-1] = n
	p := n
	// 对于每个 i，记录下一个递减的位置
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			p = i
		}
		nextDec[i] = p
	}

	solve := func(k int) {
		lastMax := 0 // 上一段的最大值
		for r := k - 1; r < n; r += k {
			l := r - k + 1
			m := nextDec[l]
			if m >= r {
				// [l, r] 是递增的，最小值为 nums[l]，最大值为 nums[r]
				// 最小值必须 >= 上一段的最大值
				if nums[l] < lastMax {
					return
				}
				lastMax = nums[r]
			} else {
				// [l, m] 是第一段，[m+1, r] 是第二段
				// 第二段必须是递增的，且第二段的最小值必须 >= 上一段的最大值，且第二段的最大值必须 <= 第一段的最小值
				if nextDec[m+1] < r || nums[m+1] < lastMax || nums[r] > nums[l] {
					return
				}
				lastMax = nums[m]
			}
		}
		ans += k // 满足要求
	}

	// 枚举 n 的因子 k
	for k := 1; k*k <= n; k++ {
		if n%k == 0 {
			solve(k)
			if k*k < n {
				solve(n / k)
			}
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log \log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。对于 $n$ 的因子 $k$，我们跑了一个 $\mathcal{O}\left(\dfrac{n}{k}\right)$ 的循环，所以总的循环次数为 $\mathcal{O}\left(\sum\limits_{k|n}\dfrac{n}{k}\right)$。由于 $n$ 的因子的倒数和是 $\mathcal{O}(\log \log n)$ 级别的，所以时间复杂度为 $\mathcal{O}(n\log \log n)$。
- 空间复杂度：$\mathcal{O}(n)$。

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
