## 方法一：正向思维

整体是个 BFS 的框架。

设 $v=\textit{nums}[i]$。

- 如果 $v$ 不是质数，只能往 $i-1$ 和 $i+1$ 走一步。
- 如果 $v$ 是质数，除了能往 $i-1$ 和 $i+1$ 走一步，还能往 $v$ 在 $\textit{nums}$ 中的倍数那走一步。

怎么知道 $v$ 的倍数在哪？

预处理。在执行 BFS 之前，对于 $\textit{nums}$ 中的每个数 $x=\textit{nums}[i]$，把 $x$ 的的质因子 $p$ 和下标 $i$ 插入哈希表。哈希表的 key 是质数，value 是下标列表。

预处理后，从哈希表中可以直接获取到 $v$ 的倍数的下标。

注意遍历完下标列表后，要把列表从哈希表中删除（或者清空），避免反复遍历列表。比如一个有很多质数 $2$ 的数组，我们把这些 $2$ 的下标入队，出队的时候，不能重复遍历 $2$ 的倍数的下标列表。

代码实现时：

1. 预处理每个数的质因子列表，思路和**埃氏筛**是一样的。
2. 用双列表实现 BFS，原理请看[【基础算法精讲 13】](https://www.bilibili.com/video/BV1hG4y1277i/)。当然，用队列也可以。

```py [sol-Python3]
# 预处理每个数的质因子列表，思路同埃氏筛
MX = 1_000_001
prime_factors = [[] for _ in range(MX)]
for i in range(2, MX):
    if not prime_factors[i]:  # i 是质数
        for j in range(i, MX, i):  # i 的倍数有质因子 i
            prime_factors[j].append(i)

class Solution:
    def minJumps(self, nums: List[int]) -> int:
        n = len(nums)
        groups = defaultdict(list)
        for i, x in enumerate(nums):
            for p in prime_factors[x]:
                groups[p].append(i)  # 对于质数 p，可以跳到下标 i

        ans = 0
        vis = [False] * n
        vis[0] = True
        q = [0]

        while True:
            tmp = q
            q = []
            for i in tmp:
                if i == n - 1:
                    return ans
                idx = groups[nums[i]]
                idx.append(i + 1)
                if i:
                    idx.append(i - 1)
                for j in idx:  # 可以从 i 跳到 j
                    if not vis[j]:
                        vis[j] = True
                        q.append(j)
                idx.clear()  # 避免重复访问下标列表
            ans += 1
```

```java [sol-Java]
class Solution {
    private static final int MX = 1_000_001;
    private static final List<Integer>[] primeFactors = new ArrayList[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.setAll(primeFactors, _ -> new ArrayList<>());
        // 预处理每个数的质因子列表，思路同埃氏筛
        for (int i = 2; i < MX; i++) {
            if (primeFactors[i].isEmpty()) { // i 是质数
                for (int j = i; j < MX; j += i) { // i 的倍数有质因子 i
                    primeFactors[j].add(i);
                }
            }
        }
    }

    public int minJumps(int[] nums) {
        init();

        int n = nums.length;
        Map<Integer, List<Integer>> groups = new HashMap<>();
        for (int i = 0; i < n; i++) {
            for (int p : primeFactors[nums[i]]) {
                // 对于质数 p，可以跳到下标 i
                groups.computeIfAbsent(p, _ -> new ArrayList<>()).add(i);
            }
        }

        int ans = 0;
        boolean[] vis = new boolean[n];
        vis[0] = true;
        List<Integer> q = List.of(0);

        while (true) {
            List<Integer> tmp = q;
            q = new ArrayList<>();
            for (int i : tmp) {
                if (i == n - 1) {
                    return ans;
                }
                List<Integer> idx = groups.computeIfAbsent(nums[i], _ -> new ArrayList<>());
                idx.add(i + 1);
                if (i > 0) {
                    idx.add(i - 1);
                }
                for (int j : idx) { // 可以从 i 跳到 j
                    if (!vis[j]) {
                        vis[j] = true;
                        q.add(j);
                    }
                }
                idx.clear(); // 避免重复访问下标列表
            }
            ans++;
        }
    }
}
```

```cpp [sol-C++]
const int MX = 1'000'001;
vector<int> prime_factors[MX];

int init = [] {
    // 预处理每个数的质因子列表，思路同埃氏筛
    for (int i = 2; i < MX; i++) {
        if (prime_factors[i].empty()) { // i 是质数
            for (int j = i; j < MX; j += i) { // i 的倍数有质因子 i
                prime_factors[j].push_back(i);
            }
        }
    }
    return 0;
}();

class Solution {
public:
    int minJumps(vector<int>& nums) {
        int n = nums.size();
        unordered_map<int, vector<int>> groups;
        for (int i = 0; i < n; i++) {
            for (int p : prime_factors[nums[i]]) {
                groups[p].push_back(i); // 对于质数 p，可以跳到下标 i
            }
        }

        int ans = 0;
        vector<int8_t> vis(n);
        vis[0] = true;
        vector<int> q = {0};

        while (true) {
            auto tmp = q;
            q.clear();
            for (int i : tmp) {
                if (i == n - 1) {
                    return ans;
                }
                auto& idx = groups[nums[i]];
                idx.push_back(i + 1);
                if (i > 0) {
                    idx.push_back(i - 1);
                }
                for (int j : idx) { // 可以从 i 跳到 j
                    if (!vis[j]) {
                        vis[j] = true;
                        q.push_back(j);
                    }
                }
                idx.clear(); // 避免重复访问下标列表
            }
            ans++;
        }
    }
};
```

```go [sol-Go]
const mx = 1_000_001

var primeFactors = [mx][]int{}

func init() {
	// 预处理每个数的质因子列表，思路同埃氏筛
	for i := 2; i < mx; i++ {
		if primeFactors[i] == nil { // i 是质数
			for j := i; j < mx; j += i { // i 的倍数有质因子 i
				primeFactors[j] = append(primeFactors[j], i)
			}
		}
	}
}

func minJumps(nums []int) (ans int) {
	n := len(nums)
	groups := map[int][]int{}
	for i, x := range nums {
		for _, p := range primeFactors[x] {
			groups[p] = append(groups[p], i) // 对于质数 p，可以跳到下标 i
		}
	}

	vis := make([]bool, n)
	vis[0] = true
	q := []int{0}
	for {
		tmp := q
		q = nil
		for _, i := range tmp {
			if i == n-1 {
				return
			}
			idx := groups[nums[i]]
			idx = append(idx, i+1)
			if i > 0 {
				idx = append(idx, i-1)
			}
			for _, j := range idx { // 可以从 i 跳到 j
				if !vis[j] {
					vis[j] = true
					q = append(q, j)
				}
			}
			delete(groups, nums[i]) // 避免重复访问下标列表
		}
		ans++
	}
}
```

#### 复杂度分析

预处理的时间和空间不计入。

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。循环次数取决于下标列表的总长度（这决定了 BFS 的循环次数）。在最坏情况下，构造这样一个 $\textit{nums}$ 数组，它含有 $\mathcal{O}(\log U)$ 个不同的质数，以及 $\mathcal{O}(n-\log U)$ 个有 $\mathcal{O}(\log U)$ 个质因子的数。此时下标列表的总长度为 $\mathcal{O}(n\log U)$，且 BFS 的循环次数也为 $\mathcal{O}(n\log U)$。
- 空间复杂度：$\mathcal{O}(n\log U)$。

## 方法二：逆向思维

从终点 $n-1$ 跳到起点 $0$。

跳跃规则反过来，从 $i$ 跳到 $\textit{nums}[i]$ 的质因子 $p$ 的下标。

在执行 BFS 之前，对于 $\textit{nums}$ 中的每个数 $x=\textit{nums}[i]$，如果 $x$ 是质数，把 $x$ 和下标 $i$ 插入哈希表。哈希表的 key 是质数，value 是下标列表。

```py [sol-Python3]
# 预处理每个数的质因子列表，思路同埃氏筛
MX = 1_000_001
prime_factors = [[] for _ in range(MX)]
for i in range(2, MX):
    if not prime_factors[i]:  # i 是质数
        for j in range(i, MX, i):  # i 的倍数有质因子 i
            prime_factors[j].append(i)

class Solution:
    def minJumps(self, nums: List[int]) -> int:
        n = len(nums)
        groups = defaultdict(list)
        for i, x in enumerate(nums):
            if len(prime_factors[x]) == 1:  # x 是质数
                groups[x].append(i)

        ans = 0
        vis = [False] * n
        vis[-1] = True
        q = [n - 1]

        while True:
            tmp = q
            q = []
            for i in tmp:
                if i == 0:
                    return ans
                if not vis[i - 1]:
                    vis[i - 1] = True
                    q.append(i - 1)
                if i < n - 1 and not vis[i + 1]:
                    vis[i + 1] = True
                    q.append(i + 1)
                # 逆向思维：从 i 倒着跳到 nums[i] 的质因子 p 的下标 j
                for p in prime_factors[nums[i]]:
                    idx = groups[p]
                    for j in idx:
                        if not vis[j]:
                            vis[j] = True
                            q.append(j)
                    idx.clear()  # 避免重复访问下标列表
            ans += 1
```

```java [sol-Java]
class Solution {
    private static final int MX = 1_000_001;
    private static final List<Integer>[] primeFactors = new ArrayList[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.setAll(primeFactors, _ -> new ArrayList<>());
        // 预处理每个数的质因子列表，思路同埃氏筛
        for (int i = 2; i < MX; i++) {
            if (primeFactors[i].isEmpty()) { // i 是质数
                for (int j = i; j < MX; j += i) { // i 的倍数有质因子 i
                    primeFactors[j].add(i);
                }
            }
        }
    }

    public int minJumps(int[] nums) {
        init();

        int n = nums.length;
        Map<Integer, List<Integer>> groups = new HashMap<>();
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (primeFactors[x].size() == 1) { // x 是质数
                groups.computeIfAbsent(x, _ -> new ArrayList<>()).add(i);
            }
        }

        int ans = 0;
        boolean[] vis = new boolean[n];
        vis[n - 1] = true;
        List<Integer> q = List.of(n - 1);

        while (true) {
            List<Integer> tmp = q;
            q = new ArrayList<>();
            for (int i : tmp) {
                if (i == 0) {
                    return ans;
                }
                if (!vis[i - 1]) {
                    vis[i - 1] = true;
                    q.add(i - 1);
                }
                if (i + 1 < n && !vis[i + 1]) {
                    vis[i + 1] = true;
                    q.add(i + 1);
                }
                // 逆向思维：从 i 倒着跳到 nums[i] 的质因子 p 的下标 j
                for (int p : primeFactors[nums[i]]) {
                    List<Integer> idx = groups.remove(p); // 避免重复访问下标列表
                    if (idx != null) {
                        for (int j : idx) {
                            if (!vis[j]) {
                                vis[j] = true;
                                q.add(j);
                            }
                        }
                    }
                }
            }
            ans++;
        }
    }
}
```

```cpp [sol-C++]
const int MX = 1'000'001;
vector<int> prime_factors[MX];

int init = [] {
    // 预处理每个数的质因子列表，思路同埃氏筛
    for (int i = 2; i < MX; i++) {
        if (prime_factors[i].empty()) { // i 是质数
            for (int j = i; j < MX; j += i) { // i 的倍数有质因子 i
                prime_factors[j].push_back(i);
            }
        }
    }
    return 0;
}();

class Solution {
public:
    int minJumps(vector<int>& nums) {
        int n = nums.size();
        unordered_map<int, vector<int>> groups;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (prime_factors[x].size() == 1) { // x 是质数
                groups[x].push_back(i);
            }
        }

        int ans = 0;
        vector<int8_t> vis(n);
        vis[n - 1] = true;
        vector<int> q = {n - 1};

        while (true) {
            auto tmp = q;
            q.clear();
            for (int i : tmp) {
                if (i == 0) {
                    return ans;
                }
                if (!vis[i - 1]) {
                    vis[i - 1] = true;
                    q.push_back(i - 1);
                }
                if (i < n - 1 && !vis[i + 1]) {
                    vis[i + 1] = true;
                    q.push_back(i + 1);
                }
                // 逆向思维：从 i 倒着跳到 nums[i] 的质因子 p 的下标 j
                for (int p : prime_factors[nums[i]]) {
                    auto it = groups.find(p);
                    if (it != groups.end()) {
                        for (int j : it->second) {
                            if (!vis[j]) {
                                vis[j] = true;
                                q.push_back(j);
                            }
                        }
                        groups.erase(it); // 避免重复访问下标列表
                    }
                }
            }
            ans++;
        }
    }
};
```

```go [sol-Go]
const mx = 1_000_001

var primeFactors = [mx][]int{}

func init() {
	// 预处理每个数的质因子列表，思路同埃氏筛
	for i := 2; i < mx; i++ {
		if primeFactors[i] == nil { // i 是质数
			for j := i; j < mx; j += i { // i 的倍数有质因子 i
				primeFactors[j] = append(primeFactors[j], i)
			}
		}
	}
}

func minJumps(nums []int) (ans int) {
	n := len(nums)
	groups := map[int][]int{}
	for i, x := range nums {
		if len(primeFactors[x]) == 1 { // x 是质数
			groups[x] = append(groups[x], i)
		}
	}

	vis := make([]bool, n)
	vis[n-1] = true
	q := []int{n - 1}
	for {
		tmp := q
		q = nil
		for _, i := range tmp {
			if i == 0 {
				return
			}
			if !vis[i-1] {
				vis[i-1] = true
				q = append(q, i-1)
			}
			if i < n-1 && !vis[i+1] {
				vis[i+1] = true
				q = append(q, i+1)
			}
			// 逆向思维：从 i 倒着跳到 nums[i] 的质因子 p 的下标 j
			for _, p := range primeFactors[nums[i]] {
				for _, j := range groups[p] {
					if !vis[j] {
						vis[j] = true
						q = append(q, j)
					}
				}
				delete(groups, p) // 避免重复访问下标列表
			}
		}
		ans++
	}
}
```

#### 复杂度分析

预处理的时间和空间不计入。

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见数学题单的「**§1.3 质因数分解**」。

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
