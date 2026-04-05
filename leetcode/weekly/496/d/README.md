## 前置题目

[213. 打家劫舍 II](https://leetcode.cn/problems/house-robber-ii/)，[我的题解](https://leetcode.cn/problems/house-robber-ii/solutions/2445622/jian-ji-xie-fa-zhi-jie-diao-yong-198-ti-qhvri/)。

## 分析

由于长为 $n$ 的**环形**数组至多有 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个峰值，所以当 $k > \left\lfloor\dfrac{n}{2}\right\rfloor$ 时无解。

如果能把环形数组变成普通数组（非环形），问题就更容易解决。

**普通数组的第一个数和最后一个数不能是峰值**。

分类讨论：

- 如果 $\textit{nums}[0]$ 是峰值，那么 $\textit{nums}[n-1]$ 不能是峰值，可以放在普通数组的最左边。于是构造数组 $a = [\textit{nums}[n-1]] + \textit{nums}$，问题变成使 $a$ 包含 $k$ 个峰值的最小操作次数。
- 如果 $\textit{nums}[0]$ 不是峰值，那么 $\textit{nums}[0]$ 可以放在普通数组的最右边。于是构造数组 $a = \textit{nums} + [\textit{nums}[0]]$，问题变成使 $a$ 包含 $k$ 个峰值的最小操作次数。

## 寻找子问题

设普通数组 $a$ 的长度为 $n$。用「选或不选」讨论是否把 $a[n-2]$ 变成峰值：

- 不把 $a[n-2]$ 变成峰值，问题变成使 $a$ 的前缀 $[0,n-2]$ 包含 $k$ 个峰值的最小操作次数。
- 把 $a[n-2]$ 变成峰值，那么 $a[n-3]$ 不能是峰值，问题变成使 $a$ 的前缀 $[0,n-3]$ 包含 $k-1$ 个峰值的最小操作次数。

由于选或不选都会把原问题变成一个**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> 注：从右往左思考，主要是方便把递归翻译成递推。从左往右思考也是可以的。

## 状态定义与状态转移方程

根据上面的讨论，定义 $\textit{dfs}(\textit{left},i)$，表示使 $a$ 的前缀 $[0,i+1]$ 包含 $\textit{left}$ 个峰值的最小操作次数。

用「选或不选」讨论是否把 $a[i]$ 变成峰值：

- 不把 $a[i]$ 变成峰值，问题变成使 $[0,i]$ 包含 $k$ 个峰值的最小操作次数，即 $\textit{dfs}(\textit{left},i)$。
- 把 $a[i]$ 变成峰值，那么 $a[i-1]$ 不能是峰值，问题变成使 $[0,i-1]$ 包含 $\textit{left}-1$ 个峰值的最小操作次数，即 $\textit{dfs}(\textit{left}-1,i-2)$。

这两种情况取最小值，就得到了 $\textit{dfs}(\textit{left},i)$，即

$$
\textit{dfs}(\textit{left},i) = \min(\textit{dfs}(\textit{left},i), \textit{dfs}(\textit{left}-1,i-2) + \textit{op}_i)
$$

其中 $\textit{op}_i$ 是把 $a[i]$ 变成峰值的操作次数，见周赛第三题。

**递归边界**：

- $\textit{dfs}(0,i)=0$。任务完成。
- $\textit{dfs}(\textit{left},0)=\textit{dfs}(\textit{left},-1) = \infty\ (\textit{left}\ge 1)$。当 $i\le 0$ 时，如果还剩下要包含的峰值，不合法。返回 $\infty$，这样上面公式中的 $\min$ 不会取到不合法的情况。
- 优化：由于 $[0,i+1]$ 至多有 $\left\lfloor\dfrac{i+1}{2}\right\rfloor$ 个峰值，所以当 $\textit{left} > \left\lfloor\dfrac{i+1}{2}\right\rfloor$ 时，不合法，返回 $\infty$。

**递归入口**：$\textit{dfs}(k,n-2)$，这是原问题，也是答案。

⚠**巨大优化**：如果 $\textit{nums}$ 已经有至少 $k$ 个峰值，无需操作，直接返回 $0$。

关于记忆化搜索的原理，请看视频讲解 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 记忆化搜索

```py [sol-Python3]
class Solution:
    # 非环形版本
    def solve(self, a: List[int], k: int) -> int:
        # 返回使 [0,i+1] 包含 left 个峰值的最小操作次数
        @cache
        def dfs(left: int, i: int) -> int:
            if left == 0:
                return 0
            if left > (i + 1) // 2:  # [0,i+1] 至多有 (i+1)//2 个峰值
                return inf

            # 选或不选
            not_choose = dfs(left, i - 1)
            choose = dfs(left - 1, i - 2) + max(max(a[i - 1], a[i + 1]) - a[i] + 1, 0)
            return min(not_choose, choose)

        ans = dfs(k, len(a) - 2)
        dfs.cache_clear()
        return ans

    def minOperations(self, nums: List[int], k: int) -> int:
        n = len(nums)
        if k > n // 2:
            return -1

        cnt = 0
        for i in range(n):
            if nums[i - 1] < nums[i] > nums[(i + 1) % n]:
                cnt += 1
        if cnt >= k:  # 优化：已经有至少 k 个峰值了，无需操作
            return 0

        # 如果 nums[0] 是峰顶，那么 nums[-1] 不是峰顶
        ans1 = self.solve([nums[-1]] + nums, k)
        # 如果 nums[0] 不是峰顶
        ans2 = self.solve(nums + [nums[0]], k)
        return min(ans1, ans2)
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int n = nums.length;
        if (k > n / 2) {
            return -1;
        }

        int cnt = 0;
        for (int i = 0; i < n; i++) {
            if (nums[(i - 1 + n) % n] < nums[i] && nums[i] > nums[(i + 1) % n]) {
                cnt++;
            }
        }
        if (cnt >= k) { // 优化：已经有至少 k 个峰值了，无需操作
            return 0;
        }

        // 如果 nums[0] 是峰顶，那么 nums[n-1] 不是峰顶
        int[] a = new int[n + 1];
        a[0] = nums[n - 1];
        System.arraycopy(nums, 0, a, 1, n);
        int ans1 = solve(a, k);

        // 如果 nums[0] 不是峰顶
        int[] b = new int[n + 1];
        System.arraycopy(nums, 0, b, 0, n);
        b[n] = nums[0];
        int ans2 = solve(b, k);

        return Math.min(ans1, ans2);
    }

    // 非环形版本
    public int solve(int[] a, int k) {
        int n = a.length;
        int[][] memo = new int[k + 1][n - 1];
        for (int[] row : memo) {
            Arrays.fill(row, -1); // -1 表示没有计算过
        }
        return dfs(k, n - 2, a, memo);
    }

    // 返回使 [0,i+1] 包含 left 个峰值的最小操作次数
    private int dfs(int left, int i, int[] a, int[][] memo) {
        if (left == 0) {
            return 0;
        }
        if (left > (i + 1) / 2) { // [0,i+1] 至多有 (i+1)/2 个峰值
            return Integer.MAX_VALUE / 2; // 防止加法溢出
        }

        if (memo[left][i] != -1) { // 之前计算过
            return memo[left][i];
        }

        // 选或不选
        int notChoose = dfs(left, i - 1, a, memo);
        int choose = dfs(left - 1, i - 2, a, memo) + Math.max(Math.max(a[i - 1], a[i + 1]) - a[i] + 1, 0);
        int res = Math.min(notChoose, choose);

        memo[left][i] = res; // 记忆化
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 非环形版本
    int solve(vector<int>& a, int k) {
        int n = a.size();
        vector memo(k + 1, vector<int>(n - 1, -1)); // -1 表示没有计算过

        // 返回使 [0,i+1] 包含 left 个峰值的最小操作次数
        auto dfs = [&](this auto&& dfs, int left, int i) -> int {
            if (left == 0) {
                return 0;
            }
            if (left > (i + 1) / 2) { // [0,i+1] 至多有 (i+1)/2 个峰值
                return INT_MAX / 2; // 防止加法溢出
            }

            int& res = memo[left][i]; // 注意这里是引用
            if (res != -1) { // 之前计算过
                return res;
            }

            // 选或不选
            int not_choose = dfs(left, i - 1);
            int choose = dfs(left - 1, i - 2) + max(max(a[i - 1], a[i + 1]) - a[i] + 1, 0);
            res = min(not_choose, choose);

            return res;
        };

        return dfs(k, n - 2);
    }

public:
    int minOperations(vector<int>& nums, int k) {
        int n = nums.size();
        if (k > n / 2) {
            return -1;
        }

        int cnt = 0;
        for (int i = 0; i < n; i++) {
            if (nums[(i - 1 + n) % n] < nums[i] && nums[i] > nums[(i + 1) % n]) {
                cnt++;
            }
        }
        if (cnt >= k) { // 优化：已经有至少 k 个峰值了，无需操作
            return 0;
        }

        // 如果 nums[0] 是峰顶，那么 nums[n-1] 不是峰顶
        vector<int> a = {nums.back()};
        a.insert(a.end(), nums.begin(), nums.end());
        int ans1 = solve(a, k);

        // 如果 nums[0] 不是峰顶
        nums.push_back(nums[0]);
        int ans2 = solve(nums, k);

        return min(ans1, ans2);
    }
};
```

```go [sol-Go]
// 非环形版本
func solve(a []int, k int) int {
	n := len(a)
	memo := make([][]int, k+1)
	for i := range memo {
		memo[i] = make([]int, n-1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}

	// 返回使 [0,i+1] 包含 left 个峰值的最小操作次数
	var dfs func(int, int) int
	dfs = func(left, i int) int {
		if left == 0 {
			return 0
		}
		if left > (i+1)/2 { // [0,i+1] 至多有 (i+1)/2 个峰值
			return math.MaxInt / 2 // 防止加法溢出
		}

		p := &memo[left][i]
		if *p != -1 { // 之前计算过
			return *p
		}

		// 选或不选
		notChoose := dfs(left, i-1)
		choose := dfs(left-1, i-2) + max(max(a[i-1], a[i+1])-a[i]+1, 0)
		res := min(notChoose, choose)

		*p = res // 记忆化
		return res
	}

	return dfs(k, n-2)
}

func minOperations(nums []int, k int) int {
	n := len(nums)
	if k > n/2 {
		return -1
	}

	cnt := 0
	for i, x := range nums {
		if nums[(i-1+n)%n] < x && x > nums[(i+1)%n] {
			cnt++
		}
	}
	if cnt >= k { // 优化：已经有至少 k 个峰值了，无需操作
		return 0
	}

	// 如果 nums[0] 是峰顶，那么 nums[n-1] 不是峰顶
	ans1 := solve(append([]int{nums[n-1]}, nums...), k)
	// 如果 nums[0] 不是峰顶
	ans2 := solve(append(nums, nums[0]), k)
	return min(ans1, ans2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nk)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(nk)$。
- 空间复杂度：$\mathcal{O}(nk)$。保存多少状态，就需要多少空间。

## 1:1 翻译成递推

```py [sol-Python3]
class Solution:
    # 非环形版本
    def solve(self, a: List[int], k: int) -> int:
        n = len(a)
        f = [[0] * n for _ in range(k + 1)]
        for i in range(1, k + 1):
            f[i][0] = f[i][1] = inf

        ops = [max(max(a[i - 1], a[i + 1]) - a[i] + 1, 0) for i in range(1, n - 1)]
        for left in range(1, k + 1):
            for i in range(1, n - 1):
                # 选或不选
                not_choose = f[left][i]
                choose = f[left - 1][i - 1] + ops[i - 1]
                f[left][i + 1] = min(not_choose, choose)

        return f[k][n - 1]

    def minOperations(self, nums: List[int], k: int) -> int:
        n = len(nums)
        if k > n // 2:
            return -1

        cnt = 0
        for i in range(n):
            if nums[i - 1] < nums[i] > nums[(i + 1) % n]:
                cnt += 1
        if cnt >= k:  # 优化：已经有至少 k 个峰值了，无需操作
            return 0

        # 如果 nums[0] 是峰顶，那么 nums[-1] 不是峰顶
        ans1 = self.solve([nums[-1]] + nums, k)
        # 如果 nums[0] 不是峰顶
        ans2 = self.solve(nums + [nums[0]], k)
        return min(ans1, ans2)
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int n = nums.length;
        if (k > n / 2) {
            return -1;
        }

        int cnt = 0;
        for (int i = 0; i < n; i++) {
            if (nums[(i - 1 + n) % n] < nums[i] && nums[i] > nums[(i + 1) % n]) {
                cnt++;
            }
        }
        if (cnt >= k) { // 优化：已经有至少 k 个峰值了，无需操作
            return 0;
        }

        // 如果 nums[0] 是峰顶，那么 nums[n-1] 不是峰顶
        int[] a = new int[n + 1];
        a[0] = nums[n - 1];
        System.arraycopy(nums, 0, a, 1, n);
        int ans1 = solve(a, k);

        // 如果 nums[0] 不是峰顶
        int[] b = new int[n + 1];
        System.arraycopy(nums, 0, b, 0, n);
        b[n] = nums[0];
        int ans2 = solve(b, k);

        return Math.min(ans1, ans2);
    }

    // 非环形版本
    public int solve(int[] a, int k) {
        int n = a.length;
        int[][] f = new int[k + 1][n];
        for (int i = 1; i <= k; i++) {
            f[i][0] = f[i][1] = Integer.MAX_VALUE / 2;
        }

        for (int left = 1; left <= k; left++) {
            for (int i = 1; i < n - 1; i++) {
                // 选或不选
                int notChoose = f[left][i];
                int choose = f[left - 1][i - 1] + Math.max(Math.max(a[i - 1], a[i + 1]) - a[i] + 1, 0);
                f[left][i + 1] = Math.min(notChoose, choose);
            }
        }

        return f[k][n - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
    // 非环形版本
    int solve(vector<int>& a, int k) {
        int n = a.size();
        vector f(k + 1, vector<int>(n));
        for (int i = 1; i <= k; i++) {
            f[i][0] = f[i][1] = INT_MAX / 2;
        }

        for (int left = 1; left <= k; left++) {
            for (int i = 1; i < n - 1; i++) {
                // 选或不选
                int not_choose = f[left][i];
                int choose = f[left - 1][i - 1] + max(max(a[i - 1], a[i + 1]) - a[i] + 1, 0);
                f[left][i + 1] = min(not_choose, choose);
            }
        }

        return f[k][n - 1];
    }

public:
    int minOperations(vector<int>& nums, int k) {
        int n = nums.size();
        if (k > n / 2) {
            return -1;
        }

        int cnt = 0;
        for (int i = 0; i < n; i++) {
            if (nums[(i - 1 + n) % n] < nums[i] && nums[i] > nums[(i + 1) % n]) {
                cnt++;
            }
        }
        if (cnt >= k) { // 优化：已经有至少 k 个峰值了，无需操作
            return 0;
        }

        // 如果 nums[0] 是峰顶，那么 nums[n-1] 不是峰顶
        vector<int> a = {nums.back()};
        a.insert(a.end(), nums.begin(), nums.end());
        int ans1 = solve(a, k);

        // 如果 nums[0] 不是峰顶
        nums.push_back(nums[0]);
        int ans2 = solve(nums, k);

        return min(ans1, ans2);
    }
};
```

```go [sol-Go]
// 非环形版本
func solve(a []int, k int) int {
	n := len(a)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n)
		if i > 0 {
			f[i][0] = math.MaxInt / 2
			f[i][1] = math.MaxInt / 2
		}
	}

	for left := 1; left <= k; left++ {
		for i := 1; i < n-1; i++ {
			// 选或不选
			notChoose := f[left][i]
			choose := f[left-1][i-1] + max(max(a[i-1], a[i+1])-a[i]+1, 0)
			f[left][i+1] = min(notChoose, choose)
		}
	}

	return f[k][n-1]
}

func minOperations(nums []int, k int) int {
	n := len(nums)
	if k > n/2 {
		return -1
	}

	cnt := 0
	for i, x := range nums {
		if nums[(i-1+n)%n] < x && x > nums[(i+1)%n] {
			cnt++
		}
	}
	if cnt >= k { // 优化：已经有至少 k 个峰值了，无需操作
		return 0
	}

	// 如果 nums[0] 是峰顶，那么 nums[n-1] 不是峰顶
	ans1 := solve(append([]int{nums[n-1]}, nums...), k)
	// 如果 nums[0] 不是峰顶
	ans2 := solve(append(nums, nums[0]), k)
	return min(ans1, ans2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。

**注**：用滚动数组可以把空间复杂度优化到 $\mathcal{O}(n)$。

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
