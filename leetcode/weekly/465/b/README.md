**前置知识**：回溯（搜索），见[【基础算法精讲 14】](https://www.bilibili.com/video/BV1mG4y1A7Gu/)。

题目要求把 $n$ 拆分成 $k$ 个数的乘积，这可以暴搜枚举所有拆分方案。

用一个长为 $k$ 的数组 $\textit{path}$ 记录搜索过程中拆分出的数字。

定义 $\textit{dfs}(i,n,\textit{mn},\textit{mx})$ 表示构造下标在 $[i,k-1]$ 中的数字：

- 枚举 $n$ 的因子 $d$，填到 $\textit{path}[i]$ 中。
- $n$ 更新为 $\dfrac{n}{d}$。
- $\textit{mn}$ 更新为 $\min(\textit{mn},d)$。
- $\textit{mx}$ 更新为 $\max(\textit{mx},d)$。

递归到 $\textit{dfs}(i+1,n/d,\min(\textit{mn},d),\max(\textit{mx},d))$，继续搜索。

递归边界：

- 当 $i=k-1$ 时，填入 $\textit{path}[i] = n$。
- $\textit{mn}$ 更新为 $\min(\textit{mn},n)$。
- $\textit{mx}$ 更新为 $\max(\textit{mx},n)$。
- 如果 $\textit{mx} - \textit{mn}$ 小于目前的最小差值，更新最小差值，更新答案为 $\textit{path}$ 的拷贝。

递归入口：$\textit{dfs}(0,n,\infty,0)$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1SMaGz7EXe/?t=20m55s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
# 预处理每个数的因子
MX = 100_001
divisors = [[] for _ in range(MX)]
for i in range(1, MX):
    for j in range(i, MX, i):  # 枚举 i 的倍数 j
        divisors[j].append(i)  # i 是 j 的因子

class Solution:
    def minDifference(self, n: int, k: int) -> List[int]:
        min_diff = inf
        path = [0] * k
        ans = None

        def dfs(i: int, n: int, mn: int, mx: int) -> None:
            if i == k - 1:
                nonlocal min_diff, ans
                d = max(mx, n) - min(mn, n)  # 最后一个数是 n
                if d < min_diff:
                    min_diff = d
                    path[i] = n
                    ans = path.copy()  # path[:]
                return
            for d in divisors[n]:  # 枚举 n 的因子 d
                path[i] = d  # 直接覆盖，无需恢复现场
                dfs(i + 1, n // d, min(mn, d), max(mx, d))

        dfs(0, n, inf, 0)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] minDifference(int n, int k) {
        int[] path = new int[k];
        dfs(0, n, Integer.MAX_VALUE, 0, path);
        return ans;
    }

    private int minDiff = Integer.MAX_VALUE;
    private int[] ans;

    private void dfs(int i, int n, int mn, int mx, int[] path) {
        if (i == path.length - 1) {
            int d = Math.max(mx, n) - Math.min(mn, n); // 最后一个数是 n
            if (d < minDiff) {
                minDiff = d;
                path[i] = n;
                ans = path.clone();
            }
            return;
        }
        for (int d = 1; d <= n; d++) { // 枚举 n 的因子 d
            if (n % d == 0) {
                path[i] = d; // 直接覆盖，无需恢复现场
                dfs(i + 1, n / d, Math.min(mn, d), Math.max(mx, d), path);
            }
        }
    }
}
```

```java [sol-Java 预处理]
class Solution {
    private static final int MX = 100_001;
    private static final List<Integer>[] divisors = new ArrayList[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理每个数的因子
        Arrays.setAll(divisors, _ -> new ArrayList<>());
        for (int i = 1; i < MX; i++) {
            for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
                divisors[j].add(i); // i 是 j 的因子
            }
        }
    }

    public int[] minDifference(int n, int k) {
        init();
        int[] path = new int[k];
        dfs(0, n, Integer.MAX_VALUE, 0, path);
        return ans;
    }

    private int minDiff = Integer.MAX_VALUE;
    private int[] ans;

    private void dfs(int i, int n, int mn, int mx, int[] path) {
        if (i == path.length - 1) {
            int d = Math.max(mx, n) - Math.min(mn, n); // 最后一个数是 n
            if (d < minDiff) {
                minDiff = d;
                path[i] = n;
                ans = path.clone();
            }
            return;
        }
        for (int d : divisors[n]) { // 枚举 n 的因子 d
            path[i] = d; // 直接覆盖，无需恢复现场
            dfs(i + 1, n / d, Math.min(mn, d), Math.max(mx, d), path);
        }
    }
}
```

```cpp [sol-C++]
const int MX = 100'001;
vector<int> divisors[MX];

int init = [] {
    // 预处理每个数的因子
    for (int i = 1; i < MX; i++) {
        for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
            divisors[j].push_back(i); // i 是 j 的因子
        }
    }
    return 0;
}();

class Solution {
public:
    vector<int> minDifference(int n, int k) {
        int min_diff = INT_MAX;
        vector<int> path(k), ans;

        auto dfs = [&](this auto&& dfs, int i, int n, int mn, int mx) -> void {
            if (i == k - 1) {
                int d = max(mx, n) - min(mn, n); // 最后一个数是 n
                if (d < min_diff) {
                    min_diff = d;
                    path[i] = n;
                    ans = path;
                }
                return;
            }
            for (int d : divisors[n]) { // 枚举 n 的因子 d
                path[i] = d; // 直接覆盖，无需恢复现场
                dfs(i + 1, n / d, min(mn, d), max(mx, d));
            }
        };

        dfs(0, n, INT_MAX, 0);
        return ans;
    }
};
```

```go [sol-Go]
const mx = 100_001
var divisors [mx][]int

func init() {
	// 预处理每个数的因子
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}
}

func minDifference(n, k int) (ans []int) {
	minDiff := math.MaxInt
	path := make([]int, k)
	var dfs func(int, int, int, int)
	dfs = func(i, n, mn, mx int) {
		if i == k-1 {
			d := max(mx, n) - min(mn, n) // 最后一个数是 n
			if d < minDiff {
				minDiff = d
				path[i] = n
				ans = slices.Clone(path)
			}
			return
		}
		for _, d := range divisors[n] { // 枚举 n 的因子 d
			path[i] = d // 直接覆盖，无需恢复现场
			dfs(i+1, n/d, min(mn, d), max(mx, d))
		}
	}
	dfs(0, n, math.MaxInt, 0)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(D^k)$，其中 $D\le 128$ 是因子个数的最大值。搜索树是一棵 $D$ 叉树，高度为 $k$，一共有 $\mathcal{O}(D^k)$ 个节点，遍历这棵搜索树需要 $\mathcal{O}(D^k)$ 的时间。由于越往下分叉越少，实际节点个数远小于 $\mathcal{O}(D^k)$。测试表明，当 $n=90720$，$k=5$ 时节点个数达到最大，为 $254816$。
- 空间复杂度：$\mathcal{O}(k)$。

## 优化

1. **排除等效冗余**。比如先填 $2$ 再填 $3$，和先填 $3$ 再填 $2$，二者是等效的。不妨规定所填数字必须大于等于上一个填的数字。
2. 根据 1，如果因子 $d^2 >n$，那么下一个填的数必然小于 $d$，不满足要求，此时直接退出循环。
3. 根据 1，$\textit{path}$ 是递增的，所以最小值就是 $\textit{path}[0]$，最大值就是 $\textit{path}[k-1]$，无需在递归过程中维护最小值和最大值。

```py [sol-Python3]
# 预处理每个数的因子
MX = 100_001
divisors = [[] for _ in range(MX)]
for i in range(1, MX):
    for j in range(i, MX, i):  # 枚举 i 的倍数 j
        divisors[j].append(i)  # i 是 j 的因子

class Solution:
    def minDifference(self, n: int, k: int) -> List[int]:
        min_diff = inf
        path = [0] * k
        ans = None

        def dfs(i: int, n: int) -> None:
            if i == k - 1:
                nonlocal min_diff, ans
                # path[0] 最小，n 最大
                if n - path[0] < min_diff:
                    min_diff = n - path[0]
                    path[i] = n
                    ans = path.copy()  # path[:]
                return
            for d in divisors[n]:  # 枚举 n 的因子 d
                if d * d > n:
                    break
                if i == 0 or d >= path[i - 1]:
                    path[i] = d  # 直接覆盖，无需恢复现场
                    dfs(i + 1, n // d)

        dfs(0, n)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] minDifference(int n, int k) {
        int[] path = new int[k];
        dfs(0, n, path);
        return ans;
    }

    private int minDiff = Integer.MAX_VALUE;
    private int[] ans;

    private void dfs(int i, int n, int[] path) {
        if (i == path.length - 1) {
            // path[0] 最小，n 最大
            if (n - path[0] < minDiff) {
                minDiff = n - path[0];
                path[i] = n;
                ans = path.clone();
            }
            return;
        }
        for (int d = i == 0 ? 1 : path[i - 1]; d * d <= n; d++) { // 枚举 n 的因子 d
            if (n % d == 0) {
                path[i] = d; // 直接覆盖，无需恢复现场
                dfs(i + 1, n / d, path);
            }
        }
    }
}
```

```java [sol-Java 预处理]
class Solution {
    private static final int MX = 100_001;
    private static final List<Integer>[] divisors = new ArrayList[MX];
    private static boolean initialized = false;

    // 这样写比 static block 更快
    private void init() {
        if (initialized) {
            return;
        }
        initialized = true;

        // 预处理每个数的因子
        Arrays.setAll(divisors, _ -> new ArrayList<>());
        for (int i = 1; i < MX; i++) {
            for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
                divisors[j].add(i); // i 是 j 的因子
            }
        }
    }

    public int[] minDifference(int n, int k) {
        init();
        int[] path = new int[k];
        dfs(0, n, path);
        return ans;
    }

    private int minDiff = Integer.MAX_VALUE;
    private int[] ans;

    private void dfs(int i, int n, int[] path) {
        if (i == path.length - 1) {
            // path[0] 最小，n 最大
            if (n - path[0] < minDiff) {
                minDiff = n - path[0];
                path[i] = n;
                ans = path.clone();
            }
            return;
        }
        int maxD = (int) Math.sqrt(n);
        for (int d : divisors[n]) { // 枚举 n 的因子 d
            if (d > maxD) {
                break;
            }
            if (i == 0 || d >= path[i - 1]) {
                path[i] = d; // 直接覆盖，无需恢复现场
                dfs(i + 1, n / d, path);
            }
        }
    }
}
```

```cpp [sol-C++]
const int MX = 100'001;
vector<int> divisors[MX];

int init = [] {
    // 预处理每个数的因子
    for (int i = 1; i < MX; i++) {
        for (int j = i; j < MX; j += i) { // 枚举 i 的倍数 j
            divisors[j].push_back(i); // i 是 j 的因子
        }
    }
    return 0;
}();

class Solution {
public:
    vector<int> minDifference(int n, int k) {
        int min_diff = INT_MAX;
        vector<int> path(k), ans;

        auto dfs = [&](this auto&& dfs, int i, int n) -> void {
            if (i == k - 1) {
                // path[0] 最小，n 最大
                if (n - path[0] < min_diff) {
                    min_diff = n - path[0];
                    path[i] = n;
                    ans = path;
                }
                return;
            }
            int max_d = sqrt(n);
            for (int d : divisors[n]) { // 枚举 n 的因子 d
                if (d > max_d) {
                    break;
                }
                if (i == 0 || d >= path[i - 1]) {
                    path[i] = d; // 直接覆盖，无需恢复现场
                    dfs(i + 1, n / d);
                }
            }
        };

        dfs(0, n);
        return ans;
    }
};
```

```go [sol-Go]
const mx = 100_001
var divisors [mx][]int

func init() {
	// 预处理每个数的因子
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisors[j] = append(divisors[j], i) // i 是 j 的因子
		}
	}
}

func minDifference(n, k int) (ans []int) {
	minDiff := math.MaxInt
	path := make([]int, k)
	var dfs func(int, int)
	dfs = func(i, n int) {
		if i == k-1 {
			// path[0] 最小，n 最大
			if n-path[0] < minDiff {
				minDiff = n - path[0]
				path[i] = n
				ans = slices.Clone(path)
			}
			return
		}
		for _, d := range divisors[n] { // 枚举 n 的因子 d
			if d*d > n {
				break
			}
			if i == 0 || d >= path[i-1] {
				path[i] = d // 直接覆盖，无需恢复现场
				dfs(i+1, n/d)
			}
		}
	}
	dfs(0, n)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(D^{k/2})$，其中 $D\le 128$ 是因子个数的最大值。搜索树是一棵 $\sqrt{D}$ 叉树，高度为 $k$，一共有 $\mathcal{O}(D^{k/2})$ 个节点，遍历这棵搜索树需要 $\mathcal{O}(D^{k/2})$ 的时间。测试表明，当 $n=90720$，$k=5$ 时节点个数达到最大，为 $4400$。
- 空间复杂度：$\mathcal{O}(k)$。

**注**：如果把枚举 $d$ 的上界由 $\left\lfloor\sqrt n\right\rfloor$ 改成 $\left\lfloor\sqrt[k-i] n\right\rfloor$，则由调和级数可知，节点个数降至 $\mathcal{O}(D^{\log k})$。测试表明，当 $n=90720$，$k=5$ 时节点个数达到最大，为 $2864$。

## 专题训练

见下面回溯题单的「**§4.7 搜索**」。

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
