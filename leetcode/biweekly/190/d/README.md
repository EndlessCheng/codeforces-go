本文介绍两种方法：

1. 至多考虑 $\mathcal{O}(\log U)$ 个要删除的位置。其中 $U=\max(\textit{nums})$。
2. 至多考虑 $1$ 个要删除的位置。

## 性质一

数组 $a$ 的有效分割是把 $a$ 分割成前缀 $P$ 和后缀 $Q$，满足 

$$
\gcd(P) = \gcd(Q) = G
$$

由于

$$
\gcd(a) = \gcd(P + Q) = \gcd(\gcd(P), \gcd(Q)) = \gcd(G,G) = G
$$

所以对于 $a$ 的任意有效分割 $P+Q$，都满足

$$
\gcd(P) = \gcd(Q) = \gcd(a)
$$

## 性质二

如果我们发现一个元素 $\textit{nums}[i]$ 不改变前缀 GCD，那么删除 $\textit{nums}[i]$，相比不删除，会让得分（有效分割个数）变得更大吗？

设 $\textit{pre}[i]$ 为 $\textit{nums}$ 的前缀 $[0,i]$ 的 GCD。

设 $\textit{suf}[i]$ 为 $\textit{nums}$ 的后缀 $[i,n-1]$ 的 GCD，其中 $n$ 是 $\textit{nums}$ 的长度。

设下标 $k$ 满足 $\textit{pre}[k] = \textit{pre}[k-1]$，那么有

$$
\gcd(\textit{nums}) = \gcd(\textit{pre}[k], \textit{suf}[k+1]) = \gcd(\textit{pre}[k-1], \textit{suf}[k+1])
$$

这说明删除 $\textit{nums}[k]$ 不改变数组的 GCD。

设 $G = \gcd(\textit{nums})$。根据 GCD 的定义，$\textit{nums}[k]$ 是 $G$ 的倍数。

设删除 $\textit{nums}[k]$ 后的数组为 $a$，其长度为 $n-1$。

考虑 $a$ 的一个有效分割 $(j,j+1)$，根据性质一，我们有

$$
\gcd(a[0,j]) = \gcd(a[j+1,n-2]) = \gcd(a) = G
$$

如果不删除 $\textit{nums}[k]$，分割 $(j,j+1)$ 是否仍然为有效分割？

把 $\textit{nums}[k]$ 放回 $a$ 中：

- 如果 $\textit{nums}[k]$ 位于前缀中，由于 $\textit{nums}[k]$ 是 $G$ 的倍数，所以 $\gcd(G, \textit{nums}[k]) = G$，说明添加 $\textit{nums}[k]$ 后，前缀 GCD 仍然是 $G$。后缀没变，其 GCD 仍然是 $G$。所以前缀后缀 GCD 仍然相等，分割 $(j,j+1)$ 仍然为有效分割。
- 如果 $\textit{nums}[k]$ 位于后缀中，同理，分割 $(j,j+1)$ 仍然为有效分割。

**结论**：在 $a$ 中的每个有效分割，都能在 $\textit{nums}$ 中找到与之对应的有效分割。所以 $a$ 的得分小于等于 $\textit{nums}$ 的得分，我们无需考虑满足 $\textit{pre}[k] = \textit{pre}[k-1]$ 的下标 $k$。

由此，我们得到一个重要剪枝：

- 如果要删除，只需考虑满足 $k=0$ 或者 $\textit{pre}[k] \ne \textit{pre}[k-1]$ 的下标 $k$。

> **问**：这样的 $k$ 有多少个？换句话说，前缀 GCD 会变化多少次？
> 
> **答**：设前缀 $[0,i-1]$ 的 GCD 为 $g$，考虑前缀 $[0,i]$ 的 GCD，即 $\gcd(g, \textit{nums}[i])$，这是 $g$ 的因子。所以前缀 $[0,i]$ 的 GCD，要么不变仍然为 $g$，要么是 $g$ 的真因子，不超过 $\dfrac{g}{2}$。所以前缀越长，GCD 要么不变，要么至少减半。而一个数至多减半 $\mathcal{O}(\log U)$ 次，其中 $U=\max(\textit{nums})$。所以前缀 GCD 只会变化 $\mathcal{O}(\log U)$ 次。

## 计算得分

对于数组 $a$，如何计算其得分（有效分割个数）？

考虑前后缀分解，先倒序遍历 $a$，算出后缀 GCD $\textit{suf}$。然后正序遍历 $a$，算出前缀 GCD $\textit{pre}$。对于一个分割 $(i,i+1)$，如果发现 $\textit{pre}[i] = \textit{suf}[i+1]$，那么把得分加一。

[本题视频讲解](https://www.bilibili.com/video/BV12g4X68EMH/?t=9m42s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    # 返回 a 的有效分割个数
    def count_valid_splits(self, a: list[int]) -> int:
        n = len(a)
        # suf[i] 是后缀 a[i:] 的 GCD
        suf = [0] * (n + 1)
        for j in range(n - 1, -1, -1):
            suf[j] = gcd(suf[j + 1], a[j])

        cnt = pre = 0
        for j, x in enumerate(a):
            pre = gcd(pre, x)
            # 现在 pre 是前缀 a[:j+1] 的 GCD
            if pre == suf[j + 1]:
                cnt += 1
        return cnt

    def maxValidSplits(self, nums: list[int]) -> int:
        ans = self.count_valid_splits(nums)  # 不删除元素

        # count_valid_splits 只会调用 O(log max(nums)) 次
        g = 0
        for i, x in enumerate(nums):
            if g > 0 and x % g == 0:  # x 不改变前缀 GCD
                continue  # 把 x 删了 ans 也不会变大
            g = gcd(g, x)
            ans = max(ans, self.count_valid_splits(nums[:i] + nums[i + 1:]))  # 删 x

        return ans
```

```java [sol-Java]
class Solution {
    public int maxValidSplits(int[] nums) {
        int ans = countValidSplits(nums, -1); // 不删除元素

        // countValidSplits 只会调用 O(log max(nums)) 次
        int g = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (g > 0 && x % g == 0) { // x 不改变前缀 GCD
                continue; // 把 x 删了 ans 也不会变大
            }
            g = gcd(g, x);
            ans = Math.max(ans, countValidSplits(nums, i)); // 删 x
        }

        return ans;
    }

    // 返回删除 nums[skip] 后，剩余元素的有效分割个数
    private int countValidSplits(int[] nums, int skip) {
        int n = nums.length;
        // suf[i] 是后缀 [i,n-1]（除去 skip）的 GCD
        int[] suf = new int[n + 1];
        for (int j = n - 1; j >= 0; j--) {
            if (j != skip) {
                suf[j] = gcd(suf[j + 1], nums[j]);
            } else {
                suf[j] = suf[j + 1];
            }
        }

        int pre = 0;
        int cnt = 0;
        for (int j = 0; j < n; j++) {
            if (j != skip) {
                pre = gcd(pre, nums[j]);
                // 现在 pre 是前缀 [0,j]（除去 skip）的 GCD
                if (pre == suf[j + 1]) {
                    cnt++;
                }
            }
        }
        return cnt;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 返回删除 nums[skip] 后，剩余元素的有效分割个数
    int count_valid_splits(vector<int>& nums, int skip) {
        int n = nums.size();
        // suf[i] 是后缀 [i,n-1]（除去 skip）的 GCD
        vector<int> suf(n + 1);
        for (int j = n - 1; j >= 0; j--) {
            if (j != skip) {
                suf[j] = gcd(suf[j + 1], nums[j]);
            } else {
                suf[j] = suf[j + 1];
            }
        }

        int pre = 0;
        int cnt = 0;
        for (int j = 0; j < n; j++) {
            if (j != skip) {
                pre = gcd(pre, nums[j]);
                // 现在 pre 是前缀 [0,j]（除去 skip）的 GCD
                if (pre == suf[j + 1]) {
                    cnt++;
                }
            }
        }
        return cnt;
    }

public:
    int maxValidSplits(vector<int>& nums) {
        int ans = count_valid_splits(nums, -1); // 不删除元素

        // count_valid_splits 只会调用 O(log max(nums)) 次
        int g = 0;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (g > 0 && x % g == 0) { // x 不改变前缀 GCD
                continue; // 把 x 删了 ans 也不会变大
            }
            g = gcd(g, x);
            ans = max(ans, count_valid_splits(nums, i)); // 删 x
        }

        return ans;
    }
};
```

```go [sol-Go]
// 返回删除 nums[skip] 后，剩余元素的有效分割个数
func countValidSplits(nums []int, skip int) (cnt int) {
	n := len(nums)
	// suf[i] 是后缀 [i,n-1]（除去 skip）的 GCD
	suf := make([]int, n+1)
	for j := n - 1; j >= 0; j-- {
		if j != skip {
			suf[j] = gcd(suf[j+1], nums[j])
		} else {
			suf[j] = suf[j+1]
		}
	}

	pre := 0
	for j, x := range nums {
		if j != skip {
			pre = gcd(pre, x)
			// 现在 pre 是前缀 [0,j]（除去 skip）的 GCD
			if pre == suf[j+1] {
				cnt++
			}
		}
	}
	return
}

func maxValidSplits(nums []int) int {
	ans := countValidSplits(nums, -1) // 不删除元素

	// countValidSplits 只会调用 O(log max(nums)) 次
	g := 0
	for i, x := range nums {
		if g > 0 && x%g == 0 { // x 不改变前缀 GCD
			continue // 把 x 删了 ans 也不会变大
		}
		g = gcd(g, x)
		ans = max(ans, countValidSplits(nums, i)) // 删 x
	}

	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

## 常数优化

设删除 $\textit{nums}[k]$ 后的数组为 $a$，其前缀 GCD 数组为 $\textit{pre}$，后缀 GCD 数组为 $\textit{suf}$。

设 $G = \gcd(a)$。

由于 $\textit{pre}$ 是递减的，设 $\textit{pre}$ 中的第一个 $G$ 的下标是 $p$，那么 $\textit{pre}$ 的后缀 $[p,|a|-1]$ 都是 $G$。

由于 $\textit{suf}$ 是递增的，设 $\textit{suf}$ 中的最后一个 $G$ 的下标是 $q$，那么 $\textit{suf}$ 的前缀 $[0,q]$ 都是 $G$。

有效分割 $(i,i+1)$ 要满足 $\textit{pre}[i] = \textit{suf}[i+1] = G$，所以有 $i\ge p$ 且 $i+1\le q$，这样的 $i$ 一共有

$$
\max(q-p, 0)
$$

个。

所以找到 $p$ 和 $q$，就能算出有效分割个数。这一结论可以让我们在计算前后缀 GCD 时，提前跳出循环，减少计算量。

```py [sol-Python3]
class Solution:
    def maxValidSplits(self, nums: list[int]) -> int:
        n = len(nums)
        pre_gcd = [0] * n
        g = 0
        for i, x in enumerate(nums):
            g = gcd(g, x)
            pre_gcd[i] = g

        suf_gcd = [0] * (n + 1)
        for i in range(n - 1, -1, -1):
            suf_gcd[i] = gcd(suf_gcd[i + 1], nums[i])

        # 不删任何数
        all_gcd = suf_gcd[0]
        p = pre_gcd.index(all_gcd)  # [p,n-1] 都是 all_gcd
        q = bisect_right(suf_gcd, all_gcd, 0, n) - 1  # [0,q] 都是 all_gcd
        ans = max(q - p, 0)  # 满足 i >= p 且 i+1 <= q 的 i 的个数

        for i in range(n):
            if i > 0 and pre_gcd[i] == pre_gcd[i - 1]:
                continue

            # 删除 nums[i]
            new_g = gcd(pre_gcd[i - 1], suf_gcd[i + 1]) if i else suf_gcd[i + 1]

            g = 0
            for j, x in enumerate(nums):
                if j == i:
                    continue
                g = gcd(g, x)
                if g == new_g:
                    p = j
                    break

            g = 0
            for j in range(n - 1, -1, -1):
                if j == i:
                    continue
                g = gcd(g, nums[j])
                if g == new_g:
                    q = j
                    break

            res = q - p
            if p <= i < q:
                res -= 1  # 因为删除了 nums[i]，少一个有效分割
            ans = max(ans, res)

        return ans
```

```java [sol-Java]
class Solution {
    public int maxValidSplits(int[] nums) {
        int n = nums.length;
        int[] preGcd = new int[n];
        int g = 0;
        for (int i = 0; i < n; i++) {
            g = gcd(g, nums[i]);
            preGcd[i] = g;
        }

        int[] sufGcd = new int[n + 1];
        for (int i = n - 1; i >= 0; i--) {
            sufGcd[i] = gcd(sufGcd[i + 1], nums[i]);
        }

        // 不删任何数
        int allGcd = sufGcd[0];
        int p = 0;
        while (preGcd[p] != allGcd) {
            p++;
        }
        int q = n - 1;
        while (sufGcd[q] != allGcd) {
            q--;
        }
        int ans = Math.max(q - p, 0); // 满足 i >= p 且 i+1 <= q 的 i 的个数

        for (int i = 0; i < n; i++) {
            if (i > 0 && preGcd[i] == preGcd[i - 1]) {
                continue;
            }

            // 删除 nums[i]
            int newG = i > 0 ? gcd(preGcd[i - 1], sufGcd[i + 1]) : sufGcd[i + 1];

            g = 0;
            for (int j = 0; j < n; j++) {
                if (j == i) {
                    continue;
                }
                g = gcd(g, nums[j]);
                if (g == newG) {
                    p = j;
                    break;
                }
            }

            g = 0;
            for (int j = n - 1; j >= 0; j--) {
                if (j == i) {
                    continue;
                }
                g = gcd(g, nums[j]);
                if (g == newG) {
                    q = j;
                    break;
                }
            }

            int res = q - p;
            if (p <= i && i < q) {
                res--; // 因为删除了 nums[i]，少一个有效分割
            }
            ans = Math.max(ans, res);
        }

        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxValidSplits(vector<int>& nums) {
        int n = nums.size();
        vector<int> pre_gcd(n);
        int g = 0;
        for (int i = 0; i < n; i++) {
            g = __gcd(g, nums[i]); // __gcd 比 gcd 快
            pre_gcd[i] = g;
        }

        vector<int> suf_gcd(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            suf_gcd[i] = __gcd(suf_gcd[i + 1], nums[i]);
        }

        // 不删任何数
        int all_gcd = suf_gcd[0];
        int p = ranges::find(pre_gcd, all_gcd) - pre_gcd.begin(); // [p,n-1] 都是 all_gcd
        int q = upper_bound(suf_gcd.begin(), suf_gcd.begin() + n, all_gcd) - suf_gcd.begin() - 1; // [0,q] 都是 all_gcd
        int ans = max(q - p, 0); // 满足 i >= p 且 i+1 <= q 的 i 的个数

        for (int i = 0; i < n; i++) {
            if (i > 0 && pre_gcd[i] == pre_gcd[i - 1]) {
                continue;
            }

            // 删除 nums[i]
            int new_g = i ? __gcd(pre_gcd[i - 1], suf_gcd[i + 1]) : suf_gcd[i + 1];

            g = 0;
            for (int j = 0; j < n; j++) {
                if (j == i) {
                    continue;
                }
                g = __gcd(g, nums[j]);
                if (g == new_g) {
                    p = j;
                    break;
                }
            }

            g = 0;
            for (int j = n - 1; j >= 0; j--) {
                if (j == i) {
                    continue;
                }
                g = __gcd(g, nums[j]);
                if (g == new_g) {
                    q = j;
                    break;
                }
            }

            int res = q - p;
            if (p <= i && i < q) {
                res--; // 因为删除了 nums[i]，少一个有效分割
            }
            ans = max(ans, res);
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxValidSplits(nums []int) int {
	n := len(nums)
	preGcd := make([]int, n)
	g := 0
	for i, x := range nums {
		g = gcd(g, x)
		preGcd[i] = g
	}

	sufGcd := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sufGcd[i] = gcd(sufGcd[i+1], nums[i])
	}

	// 不删任何数
	allGcd := sufGcd[0]
	p := sort.Search(n, func(i int) bool { return preGcd[i] == allGcd }) // [p,n-1] 都是 allGcd
	q := sort.SearchInts(sufGcd[:n], allGcd+1) - 1 // [0,q] 都是 allGcd
	ans := max(q-p, 0) // 满足 i >= p 且 i+1 <= q 的 i 的个数

	for i := range n {
		if i > 0 && preGcd[i] == preGcd[i-1] {
			continue
		}

		// 删除 nums[i]
		newG := sufGcd[i+1]
		if i > 0 {
			newG = gcd(preGcd[i-1], sufGcd[i+1])
		}

		g = 0
		for j, x := range nums {
			if j == i {
				continue
			}
			g = gcd(g, x)
			if g == newG {
				p = j
				break
			}
		}

		g = 0
		for j := n - 1; j >= 0; j-- {
			if j == i {
				continue
			}
			g = gcd(g, nums[j])
			if g == newG {
				q = j
				break
			}
		}

		res := q - p
		if p <= i && i < q {
			res-- // 因为删除了 nums[i]，少一个有效分割
		}
		ans = max(ans, res)
	}

	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+\log U)\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。我们枚举了 $\mathcal{O}(\log U)$ 个删除位置，每次需要 $\mathcal{O}(n+\log U)$ 的时间计算前后缀 GCD。为什么是 $\mathcal{O}(n+\log U)$？由于前缀（后缀）越长，GCD 要么不变，要么至少减半。一个数至多减半 $\mathcal{O}(\log U)$ 次，所以 $\texttt{countValidSplits}$ 的循环次数是 $\mathcal{O}(n + \log U)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 性质三

在性质二中，我们证明了，如果删除 $\textit{nums}[k]$ 不改变数组的 GCD，那么无需考虑删除 $\textit{nums}[k]$ 的情况。

反之，如果删除 $\textit{nums}[k]$ 导致数组 GCD 增大呢？

设 $\textit{nums}[i]$ 的质因子分解中，质因子 $p$ 的指数（出现次数）为 $\nu_p(\textit{nums}[i])$。从质因子分解的角度来说，GCD 的本质是对于每个质数 $p$，计算 $\min\limits_{i=0}^{n-1} \nu_p(\textit{nums}[i])$。

对于 $\textit{nums}[k]$，如果其存在某个质因子 $p$ 使得 $\nu_p(\textit{nums}[k])$ 是所有 $\nu_p(\textit{nums}[i])$ 中的**唯一最小值**，那么删除 $\textit{nums}[k]$ 会导致 GCD 中的 $p$ 的指数增大，从而导致 GCD 增大。反之，删除 $\textit{nums}[k]$ 不会改变数组 GCD。

此外，如果 $\textit{nums}[k]$ 有「唯一最小值」性质，那么一定要考虑删除。这是因为如果不删，那么对于 $\textit{nums}$ 的任何分割，不包含 $\textit{nums}[k]$ 的那一侧，GCD 中的 $p$ 的指数比包含 $\textit{nums}[k]$ 那一侧的更大，所以两侧的 GCD 一定不同，所以 $\textit{nums}$ 没有任何有效分割。

分类讨论：

- 如果有多个会导致 GCD 增大的数，那么即使我们删除了其中一个数，仍然存在有「唯一最小值」性质的数，剩余 $n-1$ 个数没有任何有效分割。
- 如果只有一个会导致 GCD 增大的数，那么删除这个数。
- 如果没有会导致 GCD 增大的数，那么只需计算 $\textit{nums}$ 的有效分割个数。

**结论**：至多只需考虑一个要删除的数。

```py [sol-Python3]
class Solution:
    def maxValidSplits(self, nums: list[int]) -> int:
        n = len(nums)
        pre_gcd = [0] * n
        g = 0
        for i, x in enumerate(nums):
            g = gcd(g, x)
            pre_gcd[i] = g

        suf_gcd = [0] * (n + 1)
        for i in range(n - 1, -1, -1):
            suf_gcd[i] = gcd(suf_gcd[i + 1], nums[i])

        # 不删任何数
        all_gcd = suf_gcd[0]
        p = pre_gcd.index(all_gcd)  # [p,n-1] 都是 all_gcd
        q = bisect_right(suf_gcd, all_gcd, 0, n) - 1  # [0,q] 都是 all_gcd
        ans = max(q - p, 0)  # 满足 i >= p 且 i+1 <= q 的 i 的个数

        for i in range(n):
            if i > 0 and pre_gcd[i] == pre_gcd[i - 1]:
                continue

            # 删除 nums[i]
            new_g = gcd(pre_gcd[i - 1], suf_gcd[i + 1]) if i else suf_gcd[i + 1]
            if new_g == all_gcd:
                continue

            g = 0
            for j, x in enumerate(nums):
                if j == i:
                    continue
                g = gcd(g, x)
                if g == new_g:
                    p = j
                    break

            g = 0
            for j in range(n - 1, -1, -1):
                if j == i:
                    continue
                g = gcd(g, nums[j])
                if g == new_g:
                    q = j
                    break

            res = q - p
            if p <= i < q:
                res -= 1  # 因为删除了 nums[i]，少一个有效分割
            ans = max(ans, res)
            break

        return ans
```

```java [sol-Java]
class Solution {
    public int maxValidSplits(int[] nums) {
        int n = nums.length;
        int[] preGcd = new int[n];
        int g = 0;
        for (int i = 0; i < n; i++) {
            g = gcd(g, nums[i]);
            preGcd[i] = g;
        }

        int[] sufGcd = new int[n + 1];
        for (int i = n - 1; i >= 0; i--) {
            sufGcd[i] = gcd(sufGcd[i + 1], nums[i]);
        }

        // 不删任何数
        int allGcd = sufGcd[0];
        int p = 0;
        while (preGcd[p] != allGcd) {
            p++;
        }
        int q = n - 1;
        while (sufGcd[q] != allGcd) {
            q--;
        }
        int ans = Math.max(q - p, 0); // 满足 i >= p 且 i+1 <= q 的 i 的个数

        for (int i = 0; i < n; i++) {
            if (i > 0 && preGcd[i] == preGcd[i - 1]) {
                continue;
            }

            // 删除 nums[i]
            int newG = i > 0 ? gcd(preGcd[i - 1], sufGcd[i + 1]) : sufGcd[i + 1];
            if (newG == allGcd) {
                continue;
            }

            g = 0;
            for (int j = 0; j < n; j++) {
                if (j == i) {
                    continue;
                }
                g = gcd(g, nums[j]);
                if (g == newG) {
                    p = j;
                    break;
                }
            }

            g = 0;
            for (int j = n - 1; j >= 0; j--) {
                if (j == i) {
                    continue;
                }
                g = gcd(g, nums[j]);
                if (g == newG) {
                    q = j;
                    break;
                }
            }

            int res = q - p;
            if (p <= i && i < q) {
                res--; // 因为删除了 nums[i]，少一个有效分割
            }
            ans = Math.max(ans, res);
            break;
        }

        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxValidSplits(vector<int>& nums) {
        int n = nums.size();
        vector<int> pre_gcd(n);
        int g = 0;
        for (int i = 0; i < n; i++) {
            g = __gcd(g, nums[i]); // __gcd 比 gcd 快
            pre_gcd[i] = g;
        }

        vector<int> suf_gcd(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            suf_gcd[i] = __gcd(suf_gcd[i + 1], nums[i]);
        }

        // 不删任何数
        int all_gcd = suf_gcd[0];
        int p = ranges::find(pre_gcd, all_gcd) - pre_gcd.begin(); // [p,n-1] 都是 all_gcd
        int q = upper_bound(suf_gcd.begin(), suf_gcd.begin() + n, all_gcd) - suf_gcd.begin() - 1; // [0,q] 都是 all_gcd
        int ans = max(q - p, 0); // 满足 i >= p 且 i+1 <= q 的 i 的个数

        for (int i = 0; i < n; i++) {
            if (i > 0 && pre_gcd[i] == pre_gcd[i - 1]) {
                continue;
            }

            // 删除 nums[i]
            int new_g = i ? __gcd(pre_gcd[i - 1], suf_gcd[i + 1]) : suf_gcd[i + 1];
            if (new_g == all_gcd) {
                continue;
            }

            g = 0;
            for (int j = 0; j < n; j++) {
                if (j == i) {
                    continue;
                }
                g = __gcd(g, nums[j]);
                if (g == new_g) {
                    p = j;
                    break;
                }
            }

            g = 0;
            for (int j = n - 1; j >= 0; j--) {
                if (j == i) {
                    continue;
                }
                g = __gcd(g, nums[j]);
                if (g == new_g) {
                    q = j;
                    break;
                }
            }

            int res = q - p;
            if (p <= i && i < q) {
                res--; // 因为删除了 nums[i]，少一个有效分割
            }
            ans = max(ans, res);
            break;
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxValidSplits(nums []int) int {
	n := len(nums)
	preGcd := make([]int, n)
	g := 0
	for i, x := range nums {
		g = gcd(g, x)
		preGcd[i] = g
	}

	sufGcd := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sufGcd[i] = gcd(sufGcd[i+1], nums[i])
	}

	// 不删任何数
	allGcd := sufGcd[0]
	p := sort.Search(n, func(i int) bool { return preGcd[i] == allGcd }) // [p,n-1] 都是 allGcd
	q := sort.SearchInts(sufGcd[:n], allGcd+1) - 1 // [0,q] 都是 allGcd
	ans := max(q-p, 0) // 满足 i >= p 且 i+1 <= q 的 i 的个数

	for i := range n {
		if i > 0 && preGcd[i] == preGcd[i-1] {
			continue
		}

		// 删除 nums[i]
		newG := sufGcd[i+1]
		if i > 0 {
			newG = gcd(preGcd[i-1], sufGcd[i+1])
		}

		if newG == allGcd {
			continue
		}

		g = 0
		for j, x := range nums {
			if j == i {
				continue
			}
			g = gcd(g, x)
			if g == newG {
				p = j
				break
			}
		}

		g = 0
		for j := n - 1; j >= 0; j-- {
			if j == i {
				continue
			}
			g = gcd(g, nums[j])
			if g == newG {
				q = j
				break
			}
		}

		res := q - p
		if p <= i && i < q {
			res-- // 因为删除了 nums[i]，少一个有效分割
		}
		ans = max(ans, res)
		break
	}

	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+\log^2 U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。注意我们计算了 $\mathcal{O}(\log U)$ 次删除后的 GCD，每次算 GCD 需要 $\mathcal{O}(\log U)$ 的时间，所以这部分需要 $\mathcal{O}(\log^2 U)$ 的时间。其余同上面的复杂度分析，需要 $\mathcal{O}(n + \log U)$ 的时间。
- 空间复杂度：$\mathcal{O}(n)$。

## 性质四

设 $\textit{nums}[k]$ 左侧元素的 GCD 为 $L$，右侧元素的 GCD 为 $R$。

删除 $\textit{nums}[k]$ 后，设剩余 $n-1$ 个数组成了数组 $a$，那么 $\gcd(a) = \gcd(L,R) = G$。

如果 $a$ 中存在有效分割，可以推出什么结论？

设有效分割为 $(j,j+1)$。根据性质一，$\gcd(a[0,j]) = G$。

分类讨论：

- 如果 $j < k$，由于前缀 $a[0,j]$ 比 $a[0,k-1]$ 短（或相同），所以 $\gcd(a[0,j])\ge \gcd(a[0,k-1]) = L$，即 $G\ge L$。又由于 $\gcd(L,R) = G$，所以 $G\le L$。所以 $G = L$。结合 $\gcd(L,R) = G = L$ 可知，$R$ 是 $L$ 的倍数。
- 如果 $j\ge k$，用后缀分析，同理可得，$L$ 是 $R$ 的倍数。

**结论**：如果 $a$ 中存在有效分割，那么 $L$ 是 $R$ 的倍数，或者 $R$ 是 $L$ 的倍数。由该命题的逆否命题可知，如果 $L$ 和 $R$ 不满足倍数关系，那么不存在有效分割。

所以只需在 $L$ 和 $R$ 一个是另一个的倍数的情况下计算 $\textit{newG}$。如果 $L$ 是 $R$ 的倍数，那么 $\textit{newG} = \gcd(L,R) = R$；如果 $R$ 是 $L$ 的倍数，那么 $\textit{newG} = \gcd(L,R) = L$。

```py [sol-Python3]
class Solution:
    def maxValidSplits(self, nums: list[int]) -> int:
        n = len(nums)
        pre_gcd = [0] * n
        g = 0
        for i, x in enumerate(nums):
            g = gcd(g, x)
            pre_gcd[i] = g

        suf_gcd = [0] * (n + 1)
        for i in range(n - 1, -1, -1):
            suf_gcd[i] = gcd(suf_gcd[i + 1], nums[i])

        # 不删任何数
        all_gcd = suf_gcd[0]
        p = pre_gcd.index(all_gcd)  # [p,n-1] 都是 all_gcd
        q = bisect_right(suf_gcd, all_gcd, 0, n) - 1  # [0,q] 都是 all_gcd
        ans = max(q - p, 0)  # 满足 i >= p 且 i+1 <= q 的 i 的个数

        for i in range(n):
            if i > 0 and pre_gcd[i] == pre_gcd[i - 1]:
                continue

            # 删除 nums[i]
            pre_g = pre_gcd[i - 1] if i > 0 else 0
            suf_g = suf_gcd[i + 1] if i < n - 1 else 0
            if suf_g > all_gcd and pre_g % suf_g == 0:  # 兼容 pre_g 为 0 的情况
                new_g = suf_g
            elif pre_g > all_gcd and suf_g % pre_g == 0:  # 兼容 suf_g 为 0 的情况
                new_g = pre_g
            else:
                continue

            g = 0
            for j, x in enumerate(nums):
                if j == i:
                    continue
                g = gcd(g, x)
                if g == new_g:
                    p = j
                    break

            g = 0
            for j in range(n - 1, -1, -1):
                if j == i:
                    continue
                g = gcd(g, nums[j])
                if g == new_g:
                    q = j
                    break

            res = q - p
            if p <= i < q:
                res -= 1  # 因为删除了 nums[i]，少一个有效分割
            ans = max(ans, res)
            break

        return ans
```

```java [sol-Java]
class Solution {
    public int maxValidSplits(int[] nums) {
        int n = nums.length;
        int[] preGcd = new int[n];
        int g = 0;
        for (int i = 0; i < n; i++) {
            g = gcd(g, nums[i]);
            preGcd[i] = g;
        }

        int[] sufGcd = new int[n + 1];
        for (int i = n - 1; i >= 0; i--) {
            sufGcd[i] = gcd(sufGcd[i + 1], nums[i]);
        }

        // 不删任何数
        int allGcd = sufGcd[0];
        int p = 0;
        while (preGcd[p] != allGcd) {
            p++;
        }
        int q = n - 1;
        while (sufGcd[q] != allGcd) {
            q--;
        }
        int ans = Math.max(q - p, 0); // 满足 i >= p 且 i+1 <= q 的 i 的个数

        for (int i = 0; i < n; i++) {
            if (i > 0 && preGcd[i] == preGcd[i - 1]) {
                continue;
            }

            // 删除 nums[i]
            int preG = i > 0 ? preGcd[i - 1] : 0;
            int sufG = i < n - 1 ? sufGcd[i + 1] : 0;
            int newG;
            if (sufG > allGcd && preG % sufG == 0) { // 兼容 preG 为 0 的情况
                newG = sufG;
            } else if (preG > allGcd && sufG % preG == 0) { // 兼容 sufG 为 0 的情况
                newG = preG;
            } else {
                continue;
            }

            g = 0;
            for (int j = 0; j < n; j++) {
                if (j == i) {
                    continue;
                }
                g = gcd(g, nums[j]);
                if (g == newG) {
                    p = j;
                    break;
                }
            }

            g = 0;
            for (int j = n - 1; j >= 0; j--) {
                if (j == i) {
                    continue;
                }
                g = gcd(g, nums[j]);
                if (g == newG) {
                    q = j;
                    break;
                }
            }

            int res = q - p;
            if (p <= i && i < q) {
                res--; // 因为删除了 nums[i]，少一个有效分割
            }
            ans = Math.max(ans, res);
            break;
        }

        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxValidSplits(vector<int>& nums) {
        int n = nums.size();
        vector<int> pre_gcd(n);
        int g = 0;
        for (int i = 0; i < n; i++) {
            g = __gcd(g, nums[i]); // __gcd 比 gcd 快
            pre_gcd[i] = g;
        }

        vector<int> suf_gcd(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            suf_gcd[i] = __gcd(suf_gcd[i + 1], nums[i]);
        }

        // 不删任何数
        int all_gcd = suf_gcd[0];
        int p = ranges::find(pre_gcd, all_gcd) - pre_gcd.begin(); // [p,n-1] 都是 all_gcd
        int q = upper_bound(suf_gcd.begin(), suf_gcd.begin() + n, all_gcd) - suf_gcd.begin() - 1; // [0,q] 都是 all_gcd
        int ans = max(q - p, 0); // 满足 i >= p 且 i+1 <= q 的 i 的个数

        for (int i = 0; i < n; i++) {
            if (i > 0 && pre_gcd[i] == pre_gcd[i - 1]) {
                continue;
            }

            // 删除 nums[i]
            int pre_g = i > 0 ? pre_gcd[i - 1] : 0;
            int suf_g = i < n - 1 ? suf_gcd[i + 1] : 0;
            int new_g;
            if (suf_g > all_gcd && pre_g % suf_g == 0) { // 兼容 pre_g 为 0 的情况
                new_g = suf_g;
            } else if (pre_g > all_gcd && suf_g % pre_g == 0) { // 兼容 suf_g 为 0 的情况
                new_g = pre_g;
            } else {
                continue;
            }

            g = 0;
            for (int j = 0; j < n; j++) {
                if (j == i) {
                    continue;
                }
                g = __gcd(g, nums[j]);
                if (g == new_g) {
                    p = j;
                    break;
                }
            }

            g = 0;
            for (int j = n - 1; j >= 0; j--) {
                if (j == i) {
                    continue;
                }
                g = __gcd(g, nums[j]);
                if (g == new_g) {
                    q = j;
                    break;
                }
            }

            int res = q - p;
            if (p <= i && i < q) {
                res--; // 因为删除了 nums[i]，少一个有效分割
            }
            ans = max(ans, res);
            break;
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxValidSplits(nums []int) int {
	n := len(nums)
	preGcd := make([]int, n)
	g := 0
	for i, x := range nums {
		g = gcd(g, x)
		preGcd[i] = g
	}

	sufGcd := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sufGcd[i] = gcd(sufGcd[i+1], nums[i])
	}

	// 不删任何数
	allGcd := sufGcd[0]
	p := sort.Search(n, func(i int) bool { return preGcd[i] == allGcd }) // [p,n-1] 都是 allGcd
	q := sort.SearchInts(sufGcd[:n], allGcd+1) - 1 // [0,q] 都是 allGcd
	ans := max(q-p, 0) // 满足 i >= p 且 i+1 <= q 的 i 的个数

	for i := range n {
		if i > 0 && preGcd[i] == preGcd[i-1] {
			continue
		}

		// 删除 nums[i]
		preG := 0
		if i > 0 {
			preG = preGcd[i-1]
		}
		sufG := 0
		if i < n-1 {
			sufG = sufGcd[i+1]
		}
		newG := 0
		if sufG > allGcd && preG%sufG == 0 { // 兼容 preG 为 0 的情况
			newG = sufG
		} else if preG > allGcd && sufG%preG == 0 { // 兼容 sufG 为 0 的情况
			newG = preG
		} else {
			continue
		}

		g = 0
		for j, x := range nums {
			if j == i {
				continue
			}
			g = gcd(g, x)
			if g == newG {
				p = j
				break
			}
		}

		g = 0
		for j := n - 1; j >= 0; j-- {
			if j == i {
				continue
			}
			g = gcd(g, nums[j])
			if g == newG {
				q = j
				break
			}
		}

		res := q - p
		if p <= i && i < q {
			res-- // 因为删除了 nums[i]，少一个有效分割
		}
		ans = max(ans, res)
		break
	}

	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。见前文的复杂度分析。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

1. 动态规划题单的「**专题：前后缀分解**」。
2. 数学题单的「**§1.6 最大公约数**」。
3. 位运算题单的「**GCD LogTrick**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
