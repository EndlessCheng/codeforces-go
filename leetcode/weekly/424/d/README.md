从特殊到一般。

## 没有连续 -1 的情况

下文用下划线 $\_$ 表示 $-1$（空位）。

先来看这个例子：

$$
[1,\_,1,2,3,4,\_,4,5,\_,5,6,7,8,9,\_,9,10,11,\_,11,12,\_,12]
$$

如果**只能填一种数字**，应该填什么？

考虑到填入的数字要和 $1,4,5,9,11,12$ 计算绝对差，为了让最大绝对差最小，可以都填 $6$ 或者 $7$，最大绝对差为 $6$。填其他数字只会让最大绝对差更大。

如果能填两种数字呢？

考虑把空位分成两组：$1,4,5$ 和 $9,11,12$。

- $1,4,5$ 靠近最小的空位 $1$，可以都填 $3$，最大绝对差为 $2$。
- $9,11,12$ 靠近最大的空位 $12$，可以都填 $10$ 或者 $11$，最大绝对差也为 $2$。

如果改成（更加一般的情况）

$$
[1,\_,4,4,\_,5,6,7,8,9,\_,11,11,\_,12]
$$

同样地：

- 左边两个空位所填数字，绝对差和 $1,4,5$ 有关，可以都填 $3$，最大绝对差为 $2$。
- 右边两个空位所填数字，绝对差和 $9,11,12$ 有关，可以都填 $10$ 或者 $11$，最大绝对差也为 $2$。

一般地，在没有连续 $-1$ 的情况下，我们可以：

1. 先找到**和空位相邻**的最小数字 $\textit{minL}$，最大数字 $\textit{maxR}$。
2. 对于任意空位，设其左右两侧数字为 $l$ 和 $r$（$l\le r$），如果 $r$ 离 $\textit{minL}$ 比 $l$ 离 $\textit{maxR}$ 更近，即 $r-\textit{minL} < \textit{maxR}-l$，那么绝对差和 $\textit{minL},r$ 有关，取二者中间值，绝对差为 $\left\lceil\dfrac{r-\textit{minL}}{2}\right\rceil$；否则，绝对差和 $\textit{maxR},l$ 有关，取二者中间值，绝对差为 $\left\lceil\dfrac{\textit{maxR}-\textit{L}}{2}\right\rceil$。

两种情况可以整合为

$$
\left\lceil\dfrac{\min(r-\textit{minL},\textit{maxR}-\textit{L})}{2}\right\rceil
$$

用上式更新答案的最大值。

## 有连续 -1 的情况

如果连续空位都填 $x$ 或者都填 $y$，那么做法同上。例如数组中有一段 $10,\_,\_,20$，上面定义的 $l$ 和 $r$ 就分别是 $10$ 和 $20$。按照上面给出的公式计算最大绝对差。

但是，还可以在连续空位中填入两个不同的数 $x$ 和 $y$，这可能得到更小的绝对差。

**一旦我们填了两个不同的数，就不能再填其他数字了**。

设我们填入的两个数分别为 $x$ 和 $y$（$x\le y$）。

由于 $\textit{minL},\textit{maxR}$ 都和空位相邻，这两个数必须和 $x,y$ 计算绝对差。

谁和 $\textit{minL}$ 计算绝对差？谁和 $\textit{maxR}$ 计算绝对差？

由于 $x\le y$，为了让答案尽量小，$x$ 和 $\textit{minL}$ 计算绝对差 $d_1$，$y$ 和 $\textit{maxR}$ 计算绝对差 $d_2$。

我们的目标是让 $\max(d_1,d_2)$ 尽量小。

例如 $\textit{minL}=1,\textit{maxR}=30$，那么令 $x=10,y=20$，得 $d_1=9,d_2=10$，所以 $\max(d_1,d_2)=10$。

一般地，取 $\textit{minL}$ 和 $\textit{maxR}$ 之间的三等分点，可以得到

$$
\max(d_1,d_2) = \left\lceil\dfrac{\textit{maxR} - \textit{minL}}{3}\right\rceil = \left\lfloor\dfrac{\textit{maxR} - \textit{minL}+2}{3}\right\rfloor
$$

> 也可以这样理解，在不存在孤立 $-1$ 的情况下，上式相当于答案的上界。

**总结**：

- 连续空位都填 $x$ 或者都填 $y$，最大绝对差为 $\left\lceil\dfrac{\min(r-\textit{minL},\textit{maxR}-\textit{L})}{2}\right\rceil$。
- 连续空位填 $x$ 以及 $y$，最大绝对差为 $\left\lceil\dfrac{\textit{maxR} - \textit{minL}}{3}\right\rceil$。

二者取最小值，得

$$
\min\left(\left\lceil\dfrac{\min(r-\textit{minL},\textit{maxR}-\textit{L})}{2}\right\rceil,\left\lceil\dfrac{\textit{maxR} - \textit{minL}}{3}\right\rceil\right)
$$

用上式更新答案的最大值。

```py [sol-Python3]
class Solution:
    def minDifference(self, nums: List[int]) -> int:
        n = len(nums)
        # 和空位相邻的最小数字 min_l 和最大数字 max_r
        min_l, max_r = inf, 0
        for i, v in enumerate(nums):
            if v != -1 and (i > 0 and nums[i - 1] == -1 or i < n - 1 and nums[i + 1] == -1):
                min_l = min(min_l, v)
                max_r = max(max_r, v)

        def calc_diff(l: int, r: int, big: bool) -> int:
            d = (min(r - min_l, max_r - l) + 1) // 2
            if big:
                d = min(d, (max_r - min_l + 2) // 3)  # d 不能超过上界
            return d

        ans = 0
        pre_i = -1
        for i, v in enumerate(nums):
            if v == -1:
                continue
            if pre_i >= 0:
                if i - pre_i == 1:
                    ans = max(ans, abs(v - nums[pre_i]))
                else:
                    ans = max(ans, calc_diff(min(nums[pre_i], v), max(nums[pre_i], v), i - pre_i > 2))
            elif i > 0:
                ans = max(ans, calc_diff(v, v, False))
            pre_i = i
        if 0 <= pre_i < n - 1:
            ans = max(ans, calc_diff(nums[pre_i], nums[pre_i], False))
        return ans
```

```java [sol-Java]
class Solution {
    public int minDifference(int[] nums) {
        int n = nums.length;
        // 和空位相邻的最小数字 minL 和最大数字 maxR
        int minL = Integer.MAX_VALUE;
        int maxR = 0;
        for (int i = 0; i < n; i++) {
            if (nums[i] != -1 && (i > 0 && nums[i - 1] == -1 || i < n - 1 && nums[i + 1] == -1)) {
                minL = Math.min(minL, nums[i]);
                maxR = Math.max(maxR, nums[i]);
            }
        }

        int preI = -1;
        for (int i = 0; i < n; i++) {
            if (nums[i] == -1) {
                continue;
            }
            if (preI >= 0) {
                if (i - preI == 1) {
                    ans = Math.max(ans, Math.abs(nums[i] - nums[preI]));
                } else {
                    updateAns(Math.min(nums[preI], nums[i]), Math.max(nums[preI], nums[i]), i - preI > 2, minL, maxR);
                }
            } else if (i > 0) {
                updateAns(nums[i], nums[i], false, minL, maxR);
            }
            preI = i;
        }
        if (0 <= preI && preI < n - 1) {
            updateAns(nums[preI], nums[preI], false, minL, maxR);
        }
        return ans;
    }

    private int ans;

    private void updateAns(int l, int r, boolean big, int minL, int maxR) {
        int d = (Math.min(r - minL, maxR - l) + 1) / 2;
        if (big) {
            d = Math.min(d, (maxR - minL + 2) / 3); // d 不能超过上界
        }
        ans = Math.max(ans, d);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minDifference(vector<int>& nums) {
        int n = nums.size();
        // 和空位相邻的最小数字 min_l 和最大数字 max_r
        int min_l = INT_MAX, max_r = 0;
        for (int i = 0; i < n; i++) {
            if (nums[i] != -1 && (i > 0 && nums[i - 1] == -1 || i < n - 1 && nums[i + 1] == -1)) {
                min_l = min(min_l, nums[i]);
                max_r = max(max_r, nums[i]);
            }
        }

        int ans = 0;
        auto update_ans = [&](int l, int r, bool big) {
            int d = (min(r - min_l, max_r - l) + 1) / 2;
            if (big) {
                d = min(d, (max_r - min_l + 2) / 3); // d 不能超过上界
            }
            ans = max(ans, d);
        };

        int pre_i = -1;
        for (int i = 0; i < n; i++) {
            if (nums[i] == -1) {
                continue;
            }
            if (pre_i >= 0) {
                if (i - pre_i == 1) {
                    ans = max(ans, abs(nums[i] - nums[pre_i]));
                } else {
                    update_ans(min(nums[pre_i], nums[i]), max(nums[pre_i], nums[i]), i - pre_i > 2);
                }
            } else if (i > 0) {
                update_ans(nums[i], nums[i], false);
            }
            pre_i = i;
        }
        if (0 <= pre_i && pre_i < n - 1) {
            update_ans(nums[pre_i], nums[pre_i], false);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minDifference(nums []int) (ans int) {
	n := len(nums)
	// 和空位相邻的最小数字 minL 和最大数字 maxR
	minL, maxR := math.MaxInt, 0
	for i, v := range nums {
		if v != -1 && (i > 0 && nums[i-1] == -1 || i < n-1 && nums[i+1] == -1) {
			minL = min(minL, v)
			maxR = max(maxR, v)
		}
	}

	updateAns := func(l, r int, big bool) {
		d := (min(r-minL, maxR-l) + 1) / 2
		if big {
			d = min(d, (maxR-minL+2)/3) // d 不能超过上界
		}
		ans = max(ans, d)
	}

	preI := -1
	for i, v := range nums {
		if v == -1 {
			continue
		}
		if preI >= 0 {
			if i-preI == 1 {
				ans = max(ans, abs(v-nums[preI]))
			} else {
				updateAns(min(nums[preI], v), max(nums[preI], v), i-preI > 2)
			}
		} else if i > 0 {
			updateAns(v, v, false)
		}
		preI = i
	}
	if 0 <= preI && preI < n-1 {
		updateAns(nums[preI], nums[preI], false)
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果可以把 $-1$ 改成三个数 $(x,y,z)$ 中的一个呢？可以改成 $k$ 个数中的一个呢？

更多相似题目，见下面贪心题单中的「**区间贪心**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
