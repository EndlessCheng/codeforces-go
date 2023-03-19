### 视频讲解

两种方法都讲了。见[【周赛 337】](https://www.bilibili.com/video/BV1EL411C7YU/)。

# 方法一：回溯

### 前置知识：子集型回溯

见[【基础算法精讲 14】](https://www.bilibili.com/video/BV1mG4y1A7Gu/)。

### 思路

在枚举 [78. 子集](https://leetcode.cn/problems/subsets/) 的基础上加个判断。

在选择 $x=\textit{nums}[i]$ 的时候，如果之前选过 $x-k$ 或 $x+k$，则不能选，否则可以选。

代码实现时，可以用哈希表或者数组来记录选过的数，从而 $O(1)$ 判断 $x-k$ 和 $x+k$ 是否选过。

### 写法一：选或不选

```py [sol1-Python3]
class Solution:
    def beautifulSubsets(self, nums: List[int], k: int) -> int:
        ans = -1  # 去掉空集
        cnt = [0] * (max(nums) + k * 2)  # 用数组实现比哈希表更快
        def dfs(i: int) -> None:
            if i == len(nums):
                nonlocal ans
                ans += 1
                return
            dfs(i + 1)  # 不选
            x = nums[i]
            if cnt[x - k] == 0 and cnt[x + k] == 0:
                cnt[x] += 1  # 选
                dfs(i + 1)
                cnt[x] -= 1  # 恢复现场
        dfs(0)
        return ans
```

```java [sol1-Java]
class Solution {
    private int[] nums, cnt;
    private int k, ans = -1; // 去掉空集

    public int beautifulSubsets(int[] nums, int k) {
        this.nums = nums;
        this.k = k;
        cnt = new int[k * 2 + 1001]; // 用数组实现比哈希表更快
        dfs(0);
        return ans;
    }

    private void dfs(int i) {
        if (i == nums.length) {
            ans++;
            return;
        }
        dfs(i + 1); // 不选
        int x = nums[i] + k; // 避免负数下标
        if (cnt[x - k] == 0 && cnt[x + k] == 0) {
            ++cnt[x]; // 选
            dfs(i + 1);
            --cnt[x]; // 恢复现场
        }
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int beautifulSubsets(vector<int> &nums, int k) {
        int ans = -1; // 去掉空集
        int cnt[3001]{}; // 用数组实现比哈希表更快
        function<void(int)> dfs = [&](int i) {
            if (i == nums.size()) {
                ans++;
                return;
            }
            dfs(i + 1); // 不选
            int x = nums[i] + k; // 避免负数下标
            if (cnt[x - k] == 0 && cnt[x + k] == 0) {
                ++cnt[x]; // 选
                dfs(i + 1);
                --cnt[x]; // 恢复现场
            }
        };
        dfs(0);
        return ans;
    }
};
```

```go [sol1-Go]
func beautifulSubsets(nums []int, k int) int {
	ans := -1 // 去掉空集
	cnt := make([]int, 1001+k*2) // 用数组实现比哈希表更快
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			ans++
			return
		}
		dfs(i + 1) // 不选
		x := nums[i] + k // 避免负数下标
		if cnt[x-k] == 0 && cnt[x+k] == 0 {
			cnt[x]++ // 选
			dfs(i + 1)
			cnt[x]-- // 恢复现场
		}
	}
	dfs(0)
	return ans
}
```

### 写法二：枚举选哪个

```py [sol2-Python3]
class Solution:
    def beautifulSubsets(self, nums: List[int], k: int) -> int:
        ans = -1  # 去掉空集
        cnt = [0] * (max(nums) + k * 2)  # 用数组实现比哈希表更快
        def dfs(i: int) -> None:  # 从 i 开始选
            nonlocal ans
            ans += 1
            if i == len(nums):
                return
            for j in range(i, len(nums)):  # 枚举选哪个
                x = nums[j]
                if cnt[x - k] == 0 and cnt[x + k] == 0:
                    cnt[x] += 1  # 选
                    dfs(j + 1)
                    cnt[x] -= 1  # 恢复现场
        dfs(0)
        return ans
```

```java [sol2-Java]
class Solution {
    private int[] nums, cnt;
    private int k, ans = -1; // 去掉空集

    public int beautifulSubsets(int[] nums, int k) {
        this.nums = nums;
        this.k = k;
        cnt = new int[k * 2 + 1001]; // 用数组实现比哈希表更快
        dfs(0);
        return ans;
    }

    // 从 i 开始选
    private void dfs(int i) {
        ++ans; // 合法子集
        if (i == nums.length)
            return;
        for (int j = i; j < nums.length; ++j) { // 枚举选哪个
            int x = nums[j] + k; // 避免负数下标
            if (cnt[x - k] == 0 && cnt[x + k] == 0) {
                ++cnt[x]; // 选
                dfs(j + 1);
                --cnt[x]; // 恢复现场
            }
        }
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int beautifulSubsets(vector<int> &nums, int k) {
        int ans = -1; // 去掉空集
        int cnt[3001]{}; // 用数组实现比哈希表更快
        function<void(int)> dfs = [&](int i) { // 从 i 开始选
            ++ans; // 合法子集
            if (i == nums.size())
                return;
            for (int j = i; j < nums.size(); ++j) { // 枚举选哪个
                int x = nums[j] + k; // 避免负数下标
                if (cnt[x - k] == 0 && cnt[x + k] == 0) {
                    ++cnt[x]; // 选
                    dfs(j + 1);
                    --cnt[x]; // 恢复现场
                }
            }
        };
        dfs(0);
        return ans;
    }
};
```

```go [sol2-Go]
func beautifulSubsets(nums []int, k int) int {
	ans := -1 // 去掉空集
	cnt := make([]int, 1001+k*2)
	var dfs func(int)
	dfs = func(i int) { // 从 i 开始选
		ans++ // 合法子集
		if i == len(nums) {
			return
		}
		for j := i; j < len(nums); j++ { // 枚举选哪个
			x := nums[j] + k // 避免负数下标
			if cnt[x-k] == 0 && cnt[x+k] == 0 {
				cnt[x]++ // 选
				dfs(j + 1)
				cnt[x]-- // 恢复现场
			}
		}
	}
	dfs(0)
	return ans
}
```

### 复杂度分析

- 时间复杂度：$O(2^n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。用哈希表实现是 $O(n)$。（数组需要一些额外空间，但是更快。）

### 相似题目

- [78. 子集](https://leetcode.cn/problems/subsets/)
- [784. 字母大小写全排列](https://leetcode.cn/problems/letter-case-permutation/)
- [1601. 最多可达成的换楼请求数目](https://leetcode.cn/problems/maximum-number-of-achievable-transfer-requests/)
- [2397. 被列覆盖的最多行数](https://leetcode.cn/problems/maximum-rows-covered-by-columns/)

# 方法二：动态规划

### 前置知识：动态规划基础（打家劫舍）

见[【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。

### 前置知识：乘法原理

见 [乘法原理](https://baike.baidu.com/item/%E4%B9%98%E6%B3%95%E5%8E%9F%E7%90%86/7538447)。

### 前置知识：同余

两个数 $x$ 和 $y$，如果 $(x-y)\bmod k = 0$，则称 $x$ 与 $y$ 对模 $k$ 同余，记作

$$
x\equiv y \pmod k
$$

例如 $42\equiv 12 \pmod {10}$，$-17\equiv 3 \pmod {10}$。

### 思路

如果两个数模 $k$ **不同余**，那么必然无法相差 $k$。

所以我们可以按照模 $k$ 的结果分组，每一组用哈希表/有序集合统计元素及其出现次数。

每一组怎么思考呢？

按照 key 从小到大排序后（设这些 key 组成了数组 $g$），相邻的 key 如果相差 $k$，那么不能同时选（类似 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)）。

为什么不考虑非相邻的 key？因为这个组里面的 key 都是模 $k$ 同余的，非相邻的 key 相差大于 $k$。

设 $g$ 的大小为 $m$。考虑最大的数 $g[m-1]$「选或不选」：

- 如果不选 $g[m-1]$，那么问题变成一个 $m-1$ 个数的子问题。
- 如果选 $g[m-1]$：
  - 这有 $2^c-1$ 种方案，这里 $c$ 为 $g[m-1]$ 的出现次数；
  - 如果 $g[m-1]-g[m-2] = k$，那么 $g[m-2]$ 绝对不能选，问题变成一个 $m-2$ 个数的子问题。
  - 如果 $g[m-1]-g[m-2] \ne k$，那么 $g[m-2]$ 可选可不选，问题变成一个 $m-1$ 个数的子问题。

定义 $f[i]$ 表示考虑前 $i$ 个 key 的方案数，可以得到一个类似打家劫舍的转移方程：

- 如果 $g[i]-g[i-1]=k$，那么 $f[i]=f[i-1]+f[i-2] \cdot( 2^{c_i}-1)$。
- 如果 $g[i]-g[i-1]\ne k$，那么 $f[i]=f[i-1]\cdot 2^{c_i}$。

其中 $c_i$ 为 $g[i]$ 的出现次数。

代码实现时，为了避免负数下标，需要偏移一位。

每组的初始值为 $f[0]=1$，$f[1] = 2^{c_0}$。

每组的答案为 $f[m]$（因为偏移了一位）。

根据乘法原理，最终答案为每组的答案的乘积。注意把空集去掉。

```py [sol2-Python3]
class Solution:
    def beautifulSubsets(self, nums: List[int], k: int) -> int:
        groups = defaultdict(Counter)
        for x in nums:
            groups[x % k][x] += 1
        ans = 1
        for cnt in groups.values():
            g = sorted(cnt.items())
            m = len(g)
            f = [0] * (m + 1)
            f[0] = 1
            f[1] = 1 << g[0][1]
            for i in range(1, m):
                if g[i][0] - g[i - 1][0] == k:
                    f[i + 1] = f[i] + f[i - 1] * ((1 << g[i][1]) - 1)
                else:
                    f[i + 1] = f[i] << g[i][1]
            ans *= f[m]  # 乘法原理
        return ans - 1  # -1 去掉空集
```

```java [sol2-Java]
class Solution {
    public int beautifulSubsets(int[] nums, int k) {
        var groups = new HashMap<Integer, TreeMap<Integer, Integer>>();
        for (int x : nums)
            groups.computeIfAbsent((x % k), key -> new TreeMap<>()).merge(x, 1, Integer::sum);
        int ans = 1;
        for (var g : groups.values()) {
            int m = g.size();
            var f = new int[m + 1];
            f[0] = 1;
            int i = 1, pre = 0;
            for (var e : g.entrySet()) {
                int cur = e.getKey();
                if (i > 1 && cur - pre == k)
                    f[i] = f[i - 1] + f[i - 2] * ((1 << e.getValue()) - 1);
                else
                    f[i] = f[i - 1] << e.getValue();
                pre = cur;
                ++i;
            }
            ans *= f[m]; // 乘法原理
        }
        return ans - 1; // -1 去掉空集
    }
}
```

```cpp [sol2-C++]
class Solution {
public:
    int beautifulSubsets(vector<int> &nums, int k) {
        unordered_map<int, map<int, int>> groups;
        for (int x : nums)
            ++groups[x % k][x];
        int ans = 1;
        for (auto&[_, g]: groups) {
            int m = g.size(), f[m + 1];
            auto it = g.begin();
            f[0] = 1;
            f[1] = 1 << it++->second;
            for (int i = 2; it != g.end(); ++it, ++i)
                if (it->first - prev(it)->first == k)
                    f[i] = f[i - 1] + f[i - 2] * ((1 << it->second) - 1);
                else
                    f[i] = f[i - 1] << it->second;
            ans *= f[m]; // 乘法原理
        }
        return ans - 1; // -1 去掉空集
    }
};
```

```go [sol2-Go]
func beautifulSubsets(nums []int, k int) int {
	groups := map[int]map[int]int{}
	for _, x := range nums {
		if groups[x%k] == nil {
			groups[x%k] = map[int]int{}
		}
		groups[x%k][x]++
	}
	ans := 1
	for _, cnt := range groups {
		m := len(cnt)
		type pair struct{ x, c int }
		g := make([]pair, 0, m)
		for x, c := range cnt {
			g = append(g, pair{x, c})
		}
		sort.Slice(g, func(i, j int) bool { return g[i].x < g[j].x })
		f := make([]int, m+1)
		f[0] = 1
		f[1] = 1 << g[0].c
		for i := 1; i < m; i++ {
			if g[i].x-g[i-1].x == k {
				f[i+1] = f[i] + f[i-1]*(1<<g[i].c-1)
			} else {
				f[i+1] = f[i] << g[i].c
			}
		}
		ans *= f[m] // 乘法原理
	}
	return ans - 1 // -1 去掉空集
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
