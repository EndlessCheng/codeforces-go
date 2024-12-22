设 $\textit{nums}$ 的长度为 $n$。总方案数为组合数 $\dbinom n 5$，减去不合法的方案数，即为答案。

枚举不合法子序列正中间的数 $x = \textit{nums}[i]$，分类讨论：

- 设 $x$ 左边有 $\textit{pre}_x$ 个 $x$，右边有 $\textit{suf}_x$ 个 $x$。
- 如果子序列只有一个 $x$，那么左边从不等于 $x$ 的数中选两个，右边从不等于 $x$ 的数中选两个，方案数为
  $$
  \dbinom {i - \textit{pre}_x} 2  \cdot \dbinom {n-1-i-\textit{suf}_x} 2
  $$
- 如果子序列只有两个 $x$，枚举子序列的另一个数 $y$，$y$ 至少要出现两次，子序列才是不合法的：
  - 设 $x$ 左边有 $\textit{pre}_y$ 个 $y$，右边有 $\textit{suf}_y$ 个 $y$。讨论左右两边 $y$ 的个数。
  - 左边有两个 $y$，右边有一个 $x$，并且右边另一个数不等于 $x$（但可以等于 $y$），方案数为
  $$
  \dbinom {\textit{pre}_y} 2  \cdot \textit{suf}_x \cdot (n-1-i- \textit{suf}_x)
  $$
  - 右边有两个 $y$，左边有一个 $x$，并且左边另一个数不等于 $x$（但可以等于 $y$），方案数为
  $$
  \dbinom {\textit{suf}_y} 2  \cdot \textit{pre}_x \cdot (i- \textit{pre}_x)
  $$
  - 左右各有一个 $y$，另一个 $x$ 在左边，并且左边另一个数不等于 $x$ 也不等于 $y$（不然就和上面的方案数重复了），方案数为
  $$
  \textit{pre}_y\cdot\textit{suf}_y\cdot\textit{pre}_x\cdot(n-1-i-\textit{suf}_x-\textit{suf}_y)
  $$
  - 左右各有一个 $y$，另一个 $x$ 在右边，并且右边另一个数不等于 $x$ 也不等于 $y$（不然就和上面的方案数重复了），方案数为
  $$
  \textit{pre}_y\cdot\textit{suf}_y\cdot\textit{suf}_x\cdot(i-\textit{pre}_x-\textit{pre}_y)
  $$

$\textit{pre}$ 和 $\textit{suf}$ 可以用两个哈希表分别维护。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ifkqYjEvc/?t=17m53s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def subsequencesWithMiddleMode(self, nums: List[int]) -> int:
        n = len(nums)
        suf = Counter(nums)
        pre = defaultdict(int)
        ans = comb(n, 5)  # 所有方案数
        # 枚举 x，作为子序列正中间的数
        for left, x in enumerate(nums[:-2]):
            suf[x] -= 1
            if left > 1:
                right = n - 1 - left
                pre_x, suf_x = pre[x], suf[x]
                # 不合法：只有一个 x
                ans -= comb(left - pre_x, 2) * comb(right - suf_x, 2)
                # 不合法：只有两个 x，且至少有两个 y（y != x）
                for y, suf_y in suf.items():  # 注意 suf_y 可能是 0
                    if y == x:
                        continue
                    pre_y = pre[y]
                    # 左边有两个 y，右边有一个 x，即 yy x xz（z 可以等于 y）
                    ans -= comb(pre_y, 2) * suf_x * (right - suf_x)
                    # 右边有两个 y，左边有一个 x，即 zx x yy（z 可以等于 y）
                    ans -= comb(suf_y, 2) * pre_x * (left - pre_x)
                    # 左右各有一个 y，另一个 x 在左边，即 xy x yz（z != y）
                    ans -= pre_y * suf_y * pre_x * (right - suf_x - suf_y)
                    # 左右各有一个 y，另一个 x 在右边，即 zy x xy（z != y）
                    ans -= pre_y * suf_y * suf_x * (left - pre_x - pre_y)
            pre[x] += 1
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int subsequencesWithMiddleMode(int[] nums) {
        int n = nums.length;
        long ans = (long) n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120; // 所有方案数
        Map<Integer, Integer> suf = new HashMap<>();
        for (int x : nums) {
            suf.merge(x, 1, Integer::sum); // suf[x]++
        }
        Map<Integer, Integer> pre = new HashMap<>(suf.size()); // 预分配空间
        // 枚举 x，作为子序列正中间的数
        for (int left = 0; left < n - 2; left++) {
            int x = nums[left];
            suf.merge(x, -1, Integer::sum); // suf[x]--
            if (left > 1) {
                int right = n - 1 - left;
                int preX = pre.getOrDefault(x, 0);
                int sufX = suf.get(x);
                // 不合法：只有一个 x
                ans -= (long) comb2(left - preX) * comb2(right - sufX);
                // 不合法：只有两个 x，且至少有两个 y（y != x）
                for (Map.Entry<Integer, Integer> e : suf.entrySet()) {
                    int y = e.getKey();
                    if (y == x) {
                        continue;
                    }
                    int sufY = e.getValue(); // 注意 sufY 可能是 0
                    int preY = pre.getOrDefault(y, 0);
                    // 左边有两个 y，右边有一个 x，即 yy x xz（z 可以等于 y）
                    ans -= (long) comb2(preY) * sufX * (right - sufX);
                    // 右边有两个 y，左边有一个 x，即 zx x yy（z 可以等于 y）
                    ans -= (long) comb2(sufY) * preX * (left - preX);
                    // 左右各有一个 y，另一个 x 在左边，即 xy x yz（z != y）
                    ans -= (long) preY * sufY * preX * (right - sufX - sufY);
                    // 左右各有一个 y，另一个 x 在右边，即 zy x xy（z != y）
                    ans -= (long) preY * sufY * sufX * (left - preX - preY);
                }
            }
            pre.merge(x, 1, Integer::sum); // pre[x]++
        }
        return (int) (ans % 1_000_000_007);
    }

    private int comb2(int num) {
        return num * (num - 1) / 2;
    }
}
```

```cpp [sol-C++]
class Solution {
    int comb2(int num) {
        return num * (num - 1) / 2;
    }

public:
    int subsequencesWithMiddleMode(vector<int>& nums) {
        int n = nums.size();
        long long ans = 1LL * n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120; // 所有方案数
        unordered_map<int, int> pre, suf;
        for (int x : nums) {
            suf[x]++;
        }
        // 枚举 x，作为子序列正中间的数
        for (int left = 0; left < n - 2; left++) {
            int x = nums[left];
            suf[x]--;
            if (left > 1) {
                int right = n - 1 - left;
                int pre_x = pre[x], suf_x = suf[x];
                // 不合法：只有一个 x
                ans -= 1LL * comb2(left - pre_x) * comb2(right - suf_x);
                // 不合法：只有两个 x，且至少有两个 y（y != x）
                for (auto& [y, suf_y] : suf) { // 注意 suf_y 可能是 0
                    if (y == x) {
                        continue;
                    }
                    int pre_y = pre[y];
                    // 左边有两个 y，右边有一个 x，即 yy x xz（z 可以等于 y）
                    ans -= 1LL * comb2(pre_y) * suf_x * (right - suf_x);
                    // 右边有两个 y，左边有一个 x，即 zx x yy（z 可以等于 y）
                    ans -= 1LL * comb2(suf_y) * pre_x * (left - pre_x);
                    // 左右各有一个 y，另一个 x 在左边，即 xy x yz（z != y）
                    ans -= 1LL * pre_y * suf_y * pre_x * (right - suf_x - suf_y);
                    // 左右各有一个 y，另一个 x 在右边，即 zy x xy（z != y）
                    ans -= 1LL * pre_y * suf_y * suf_x * (left - pre_x - pre_y);
                }
            }
            pre[x]++;
        }
        return ans % 1'000'000'007;
    }
};
```

```go [sol-Go]
func comb2(num int) int {
	return num * (num - 1) / 2
}

func subsequencesWithMiddleMode(nums []int) int {
	n := len(nums)
	ans := n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120 // 所有方案数
	suf := map[int]int{}
	for _, x := range nums {
		suf[x]++
	}
	pre := make(map[int]int, len(suf)) // 预分配空间
	// 枚举 x，作为子序列正中间的数
	for left, x := range nums[:n-2] {
		suf[x]--
		if left > 1 {
			right := n - 1 - left
			preX, sufX := pre[x], suf[x]
			// 不合法：只有一个 x
			ans -= comb2(left-preX) * comb2(right-sufX)
			// 不合法：只有两个 x，且至少有两个 y（y != x）
			for y, sufY := range suf { // 注意 sufY 可能是 0
				if y == x {
					continue
				}
				preY := pre[y]
				// 左边有两个 y，右边有一个 x，即 yy x xz（z 可以等于 y）
				ans -= comb2(preY) * sufX * (right - sufX)
				// 右边有两个 y，左边有一个 x，即 zx x yy（z 可以等于 y）
				ans -= comb2(sufY) * preX * (left - preX)
				// 左右各有一个 y，另一个 x 在左边，即 xy x yz（z != y）
				ans -= preY * sufY * preX * (right - sufX - sufY)
				// 左右各有一个 y，另一个 x 在右边，即 zy x xy（z != y）
				ans -= preY * sufY * sufX * (left - preX - preY)
			}
		}
		pre[x]++
	}
	return ans % 1_000_000_007
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

1. 如果可以修改元素呢？每改一个数，就问你此时的答案。
2. 把 $5$ 改成 $k=3,4,6,7$ 这些数呢？规定正中间的数的下标是 $\left\lfloor\dfrac{k}{2}\right\rfloor$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
