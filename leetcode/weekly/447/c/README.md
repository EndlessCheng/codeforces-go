首先读清楚题目，要计算的是字典序最小的**数组**，不是拼接后的数。

为了让字典序最小，我们要从小到大枚举。把 $\textit{nums}$ 从小到大排序，然后写一个类似 [46. 全排列](https://leetcode.cn/problems/permutations/) 的爆搜：

- 枚举答案的第一个位置填 $\textit{nums}[0],\textit{nums}[1],\ldots,\textit{nums}[n-1]$。
- 枚举答案的第二个位置填 $\textit{nums}[0],\textit{nums}[1],\ldots,\textit{nums}[n-1]$。但不能填之前填过的数字。
- 依此类推。

在枚举的过程中，维护拼接的数字模 $k$ 的结果。为什么可以在中途取模，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

一旦我们找到了答案（拼接的 $n$ 个数模 $k$ 等于 $0$），就立刻返回 $\texttt{true}$，不再继续递归搜索。

为了避免重复访问同样的状态，用一个 $\textit{vis}$ 数组记录访问过的状态。由于我们找到了答案就立刻返回 $\texttt{true}$，如果重复访问同样的状态，那么继续递归一定无法找到答案，应该返回 $\texttt{false}$。

## 细节

用集合 $S$ 表示剩余未填（没有选过的）数字的**下标**，$x$ 表示当前拼接的数字模 $k$ 的结果。

枚举 $S$ 中的下标 $i$，填入 $\textit{nums}[i]$，那么：

- $S$ 变成 $S\setminus \{i\}$。
- $x$ 变成 $(x\cdot 10^L + \textit{nums}[i])\bmod k$，其中 $L$ 是 $\textit{nums}[i]$ 的十进制长度。

递归入口：$\textit{dfs}(U,0)$，其中全集 $U=\{0,1,2,\ldots,n-1\}$。

递归边界：$\textit{dfs}(\varnothing,0)=\texttt{true}$，其余 $\textit{dfs}(\varnothing,x)=\texttt{false}$。

**代码实现时，用二进制表示集合，用位运算实现集合操作，具体请看** [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def concatenatedDivisibility(self, nums: List[int], k: int) -> List[int]:
        nums.sort()
        pow10 = [10 ** len(str(x)) for x in nums]

        ans = []
        @cache  # 充当 vis
        def dfs(s: int, x: int) -> bool:
            if s == 0:
                return x == 0
            # 枚举在 s 中的下标 i
            for i, (p10, v) in enumerate(zip(pow10, nums)):
                if s & (1 << i) and dfs(s ^ (1 << i), (x * p10 + v) % k):
                    ans.append(v)
                    return True
            return False

        if not dfs((1 << len(nums)) - 1, 0):
            return []
        ans.reverse()  # nums[i] 是倒序加入答案的，所以要反转
        return ans
```

```java [sol-Java]
class Solution {
    public int[] concatenatedDivisibility(int[] nums, int k) {
        Arrays.sort(nums);
        int n = nums.length;
        int[] pow10 = new int[n];
        for (int i = 0; i < n; i++) {
            pow10[i] = (int) Math.pow(10, Integer.toString(nums[i]).length());
        }

        int[] ans = new int[n];
        boolean[][] vis = new boolean[1 << n][k];
        if (!dfs((1 << n) - 1, 0, nums, pow10, k, vis, ans)) {
            return new int[]{};
        }
        return ans;
    }

    private boolean dfs(int s, int x, int[] nums, int[] pow10, int k, boolean[][] vis, int[] ans) {
        if (s == 0) {
            return x == 0;
        }
        if (vis[s][x]) {
            return false;
        }
        vis[s][x] = true;
        // 枚举在 s 中的下标 i
        for (int i = 0; i < nums.length; i++) {
            if ((s & (1 << i)) > 0 && dfs(s ^ (1 << i), (x * pow10[i] + nums[i]) % k, nums, pow10, k, vis, ans)) {
                ans[nums.length - Integer.bitCount(s)] = nums[i];
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
    vector<int> concatenatedDivisibility(vector<int>& nums, int k) {
        ranges::sort(nums);
        int n = nums.size();
        vector<int> pow10(n);
        for (int i = 0; i < n; i++) {
            pow10[i] = pow(10, to_string(nums[i]).size());
        }

        vector<int> ans;
        vector vis(1 << n, vector<uint8_t>(k));
        auto dfs = [&](this auto&& dfs, int s, int x) -> bool {
            if (s == 0) {
                return x == 0;
            }
            if (vis[s][x]) {
                return false;
            }
            vis[s][x] = true;
            // 枚举在 s 中的下标 i
            for (int i = 0; i < n; i++) {
                if (s & (1 << i) && dfs(s ^ (1 << i), (x * pow10[i] + nums[i]) % k)) {
                    ans.push_back(nums[i]);
                    return true;
                }
            }
            return false;
        };
        if (!dfs((1 << n) - 1, 0)) {
            return {};
        }
        ranges::reverse(ans); // nums[i] 是倒序加入答案的，所以要反转
        return ans;
    }
};
```

```go [sol-Go]
func concatenatedDivisibility(nums []int, k int) []int {
	slices.Sort(nums)
	n := len(nums)
	pow10 := make([]int, n)
	for i, x := range nums {
		pow10[i] = int(math.Pow10(len(strconv.Itoa(x))))
	}

	ans := make([]int, 0, n)
	vis := make([][]bool, 1<<n)
	for i := range vis {
		vis[i] = make([]bool, k)
	}
	var dfs func(int, int) bool
	dfs = func(s, x int) bool {
		if s == 0 {
			return x == 0
		}
		if vis[s][x] {
			return false
		}
		vis[s][x] = true
		// 枚举在 s 中的下标 i
		for t := uint(s); t > 0; t &= t - 1 {
			i := bits.TrailingZeros(t)
			if dfs(s^1<<i, (x*pow10[i]+nums[i])%k) {
				ans = append(ans, nums[i])
				return true
			}
		}
		return false
	}
	if !dfs(1<<n-1, 0) {
		return nil
	}
	slices.Reverse(ans) // nums[i] 是倒序加入答案的，所以要反转
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk2^n)$，其中 $n$ 是 $\textit{nums}$ 的长度。有 $\mathcal{O}(k2^n)$ 个状态，每个状态至多访问一次，每个状态需要花费 $\mathcal{O}(n)$ 的时间枚举下标 $i$。
- 空间复杂度：$\mathcal{O}(k2^n)$。

更多相似题目，见下面动态规划题单的「**§9.1 排列型 ① 相邻无关**」。虽然本题不是 DP，但思路是类似的。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
