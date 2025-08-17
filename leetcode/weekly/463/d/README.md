## 算法一：暴力

暴力处理每个询问，把下标为 $l,l+k,l+2k,\dots$ 的数都乘以 $v$。

最坏情况每次需要 $\mathcal{O}\left(\dfrac{n}{k}\right)$ 的时间，整体 $\mathcal{O}\left(\dfrac{nq}{k}\right)$ 时间。其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。

**特点**：当 $k$ 比较大时，算法比较快。

## 算法二：差分数组（商分数组）

**前置知识**：[差分数组](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)

如果 $k=1$，我们可以用差分数组（准确来说叫**商分数组**）记录询问，然后计算商分数组的前缀积，即可得到最终的数组。

商分数组 $d$ 与差分数组的区别是，初始值每一项都是 $1$（乘法单位元）；记录询问时，$d[l]$ 乘以 $v$，$d[r+1]$ 除以 $v$，即乘以 $v$ 的**逆元**。关于逆元，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

对于其他 $k$ 呢？

比如 $k=3$。我们可以把所有询问分为 $k=3$ 组：

- 作用在下标 $0,3,6,\dots$ 上的询问。
- 作用在下标 $1,4,7,\dots$ 上的询问。
- 作用在下标 $2,5,8,\dots$ 上的询问。

比如 $l=1$，$r=9$，更新的下标是 $1,4,7$。在左端点 $1$ 处乘以 $v$，右端点 $7+k=10$ 处除以 $v$（乘以 $v$ 的逆元）。这样我们计算 $1,4,7,10,\dots$ 的前缀积，就可以正确地得到最终数组每一项要乘的数了。

这里的 $7$ 是怎么算的？我们要找 $\le r$ 的最大的 $3k+1$，或者说，要把 $r$ 减少多少。这个减少量等同于当 $l=0$，$r=8$ 时，$r$ 到 $\le r$ 的最近的 $k$ 的倍数的距离，即 $8\bmod k = 2$。

一般地，在左端点 $l$ 处乘以 $v$，右端点 $r-(r-l)\bmod k+k$ 处除以 $v$（乘以 $v$ 的逆元）。

处理每个询问只需要 $\mathcal{O}(\log M)$ 时间计算逆元，其中 $M=10^9+7$。然而，我们需要遍历 $\mathcal{O}(K)$ 个长为 $\mathcal{O}(n)$ 的商分数组，总体需要 $\mathcal{O}(nK + q\log M)$ 的时间。其中 $K$ 是 $k_i$ 的最大值。

**特点**：当 $K$ 比较小时，算法比较快。

## 「平衡」两个算法

人为规定一个阈值 $B$。当 $k\ge B$ 的时候，我们用算法一；$k<B$ 的时候，我们用算法二。

总体时间复杂度为 $\mathcal{O}\left(\dfrac{nq}{B} + nB + q\log M\right)$。根据基本不等式，当 $B=\sqrt q$ 时，取到最小值

$$
\mathcal{O}(n\sqrt q + q\log M)
$$

这足以通过本题。

**优化**：比如，没有 $k=3$ 的询问，那么对于 $k=3$ 的商分数组，我们既不创建，也不遍历。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1kTYyzwEDD/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def xorAfterQueries(self, nums: List[int], queries: List[List[int]]) -> int:
        MOD = 1_000_000_007
        n = len(nums)
        B = isqrt(len(queries))
        diff = [None] * B

        for l, r, k, v in queries:
            if k < B:
                # 懒初始化
                if not diff[k]:
                    diff[k] = [1] * (n + k)
                diff[k][l] = diff[k][l] * v % MOD
                r = r - (r - l) % k + k
                diff[k][r] = diff[k][r] * pow(v, -1, MOD) % MOD
            else:
                for i in range(l, r + 1, k):
                    nums[i] = nums[i] * v % MOD

        for k, d in enumerate(diff):
            if not d:
                continue
            for start in range(k):
                mul_d = 1
                for i in range(start, n, k):
                    mul_d = mul_d * d[i] % MOD
                    nums[i] = nums[i] * mul_d % MOD

        return reduce(xor, nums)
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int xorAfterQueries(int[] nums, int[][] queries) {
        int n = nums.length;
        int B = (int) Math.sqrt(queries.length);
        int[][] diff = new int[B][];

        for (int[] q : queries) {
            int l = q[0], r = q[1], k = q[2];
            long v = q[3];
            if (k < B) {
                // 懒初始化
                if (diff[k] == null) {
                    diff[k] = new int[n + k];
                    Arrays.fill(diff[k], 1);
                }
                diff[k][l] = (int) (diff[k][l] * v % MOD);
                r = r - (r - l) % k + k;
                diff[k][r] = (int) (diff[k][r] * pow(v, MOD - 2) % MOD);
            } else {
                for (int i = l; i <= r; i += k) {
                    nums[i] = (int) (nums[i] * v % MOD);
                }
            }
        }

        for (int k = 0; k < B; k++) {
            int[] d = diff[k];
            if (d == null) {
                continue;
            }
            for (int start = 0; start < k; start++) {
                long mulD = 1;
                for (int i = start; i < n; i += k) {
                    mulD = mulD * d[i] % MOD;
                    nums[i] = (int) (nums[i] * mulD % MOD);
                }
            }
        }

        int ans = 0;
        for (int x : nums) {
            ans ^= x;
        }
        return ans;
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;

    long long pow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

public:
    int xorAfterQueries(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        int B = sqrt(queries.size());
        vector<vector<int>> diff(B);

        for (auto& q : queries) {
            int l = q[0], r = q[1], k = q[2];
            long long v = q[3];
            if (k < B) {
                // 懒初始化
                if (diff[k].empty()) {
                    diff[k].resize(n + k, 1);
                }
                diff[k][l] = diff[k][l] * v % MOD;
                r = r - (r - l) % k + k;
                diff[k][r] = diff[k][r] * pow(v, MOD - 2) % MOD;
            } else {
                for (int i = l; i <= r; i += k) {
                    nums[i] = nums[i] * v % MOD;
                }
            }
        }

        for (int k = 1; k < B; k++) {
            auto& d = diff[k];
            if (d.empty()) {
                continue;
            }
            for (int start = 0; start < k; start++) {
                long long mul_d = 1;
                for (int i = start; i < n; i += k) {
                    mul_d = mul_d * d[i] % MOD;
                    nums[i] = nums[i] * mul_d % MOD;
                }
            }
        }

        return reduce(nums.begin(), nums.end(), 0, bit_xor());
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func xorAfterQueries(nums []int, queries [][]int) (ans int) {
	n := len(nums)
	B := int(math.Sqrt(float64(len(queries))))
	diff := make([][]int, B)

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		if k < B {
			// 懒初始化
			if diff[k] == nil {
				diff[k] = make([]int, n+k)
				for j := range diff[k] {
					diff[k][j] = 1
				}
			}
			diff[k][l] = diff[k][l] * v % mod
			r = r - (r-l)%k + k
			diff[k][r] = diff[k][r] * pow(v, mod-2) % mod
		} else {
			for i := l; i <= r; i += k {
				nums[i] = nums[i] * v % mod
			}
		}
	}

	for k, d := range diff {
		if d == nil {
			continue
		}
		for start := range k {
			mulD := 1
			for i := start; i < n; i += k {
				mulD = mulD * d[i] % mod
				nums[i] = nums[i] * mulD % mod
			}
		}
	}

	for _, x := range nums {
		ans ^= x
	}
	return
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

## 进一步优化

把懒初始化的想法进一步扩展。比如 $k=3$ 时，没有遇到 $l\bmod k=2$ 的组，那么这一组的商分数组全为 $1$，无需遍历。

用二维布尔数组记录询问是否有 $(k,l\bmod k)$。

```py [sol-Python3]
class Solution:
    def xorAfterQueries(self, nums: List[int], queries: List[List[int]]) -> int:
        MOD = 1_000_000_007
        n = len(nums)
        B = isqrt(len(queries))
        diff = [None] * B
        has = [None] * B

        for l, r, k, v in queries:
            if k < B:
                # 懒初始化
                if not diff[k]:
                    diff[k] = [1] * (n + k)
                    has[k] = [False] * k
                has[k][l % k] = True
                diff[k][l] = diff[k][l] * v % MOD
                r = r - (r - l) % k + k
                diff[k][r] = diff[k][r] * pow(v, -1, MOD) % MOD
            else:
                for i in range(l, r + 1, k):
                    nums[i] = nums[i] * v % MOD

        for k, d in enumerate(diff):
            if not d:
                continue
            for start, b in enumerate(has[k]):
                if not b:
                    continue
                mul_d = 1
                for i in range(start, n, k):
                    mul_d = mul_d * d[i] % MOD
                    nums[i] = nums[i] * mul_d % MOD

        return reduce(xor, nums)
```

```java [sol-Java]
class Solution {
    private static final int MOD = 1_000_000_007;

    public int xorAfterQueries(int[] nums, int[][] queries) {
        int n = nums.length;
        int B = (int) Math.sqrt(queries.length);
        int[][] diff = new int[B][];
        boolean[][] has = new boolean[B][];

        for (int[] q : queries) {
            int l = q[0], r = q[1], k = q[2];
            long v = q[3];
            if (k < B) {
                // 懒初始化
                if (diff[k] == null) {
                    diff[k] = new int[n + k];
                    Arrays.fill(diff[k], 1);
                    has[k] = new boolean[k];
                }
                has[k][l % k] = true;
                diff[k][l] = (int) (diff[k][l] * v % MOD);
                r = r - (r - l) % k + k;
                diff[k][r] = (int) (diff[k][r] * pow(v, MOD - 2) % MOD);
            } else {
                for (int i = l; i <= r; i += k) {
                    nums[i] = (int) (nums[i] * v % MOD);
                }
            }
        }

        for (int k = 0; k < B; k++) {
            int[] d = diff[k];
            if (d == null) {
                continue;
            }
            for (int start = 0; start < k; start++) {
                if (!has[k][start]) {
                    continue;
                }
                long mulD = 1;
                for (int i = start; i < n; i += k) {
                    mulD = mulD * d[i] % MOD;
                    nums[i] = (int) (nums[i] * mulD % MOD);
                }
            }
        }

        int ans = 0;
        for (int x : nums) {
            ans ^= x;
        }
        return ans;
    }

    private long pow(long x, int n) {
        long res = 1;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    const int MOD = 1'000'000'007;

    long long pow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) {
                res = res * x % MOD;
            }
            x = x * x % MOD;
        }
        return res;
    }

public:
    int xorAfterQueries(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        int B = sqrt(queries.size());
        vector<vector<int>> diff(B);
        vector<vector<int8_t>> has(B);

        for (auto& q : queries) {
            int l = q[0], r = q[1], k = q[2];
            long long v = q[3];
            if (k < B) {
                // 懒初始化
                if (diff[k].empty()) {
                    diff[k].resize(n + k, 1);
                    has[k].resize(k);
                }
                has[k][l % k] = true;
                diff[k][l] = diff[k][l] * v % MOD;
                r = r - (r - l) % k + k;
                diff[k][r] = diff[k][r] * pow(v, MOD - 2) % MOD;
            } else {
                for (int i = l; i <= r; i += k) {
                    nums[i] = nums[i] * v % MOD;
                }
            }
        }

        for (int k = 1; k < B; k++) {
            auto& d = diff[k];
            if (d.empty()) {
                continue;
            }
            for (int start = 0; start < k; start++) {
                if (!has[k][start]) {
                    continue;
                }
                long long mul_d = 1;
                for (int i = start; i < n; i += k) {
                    mul_d = mul_d * d[i] % MOD;
                    nums[i] = nums[i] * mul_d % MOD;
                }
            }
        }

        return reduce(nums.begin(), nums.end(), 0, bit_xor());
    }
};
```

```go [sol-Go]
const mod = 1_000_000_007

func xorAfterQueries(nums []int, queries [][]int) (ans int) {
	n := len(nums)
	B := int(math.Sqrt(float64(len(queries))))
	diff := make([][]int, B)
	has := make([][]bool, B)

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		if k < B {
			// 懒初始化
			if diff[k] == nil {
				diff[k] = make([]int, n+k)
				for j := range diff[k] {
					diff[k][j] = 1
				}
				has[k] = make([]bool, k)
			}
			has[k][l%k] = true
			diff[k][l] = diff[k][l] * v % mod
			r = r - (r-l)%k + k
			diff[k][r] = diff[k][r] * pow(v, mod-2) % mod
		} else {
			for i := l; i <= r; i += k {
				nums[i] = nums[i] * v % mod
			}
		}
	}

	for k, d := range diff {
		if d == nil {
			continue
		}
		for start, b := range has[k] {
			if !b {
				continue
			}
			mulD := 1
			for i := start; i < n; i += k {
				mulD = mulD * d[i] % mod
				nums[i] = nums[i] * mulD % mod
			}
		}
	}

	for _, x := range nums {
		ans ^= x
	}
	return
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt q + q\log M)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n\sqrt q)$。

## 相似题目

[3590. 第 K 小的路径异或和](https://leetcode.cn/problems/kth-smallest-path-xor-sum/)

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
