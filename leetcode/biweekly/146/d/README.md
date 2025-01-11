## 方法一：枚举

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
  - 左右各一个 $y$，左边还有一个 $x$，右边另一个数不等于 $x$ 也不等于 $y$（不然就和上面的方案数重复了），方案数为
  $$
  \textit{pre}_y\cdot\textit{suf}_y\cdot\textit{pre}_x\cdot(n-1-i-\textit{suf}_x-\textit{suf}_y)
  $$
  - 左右各一个 $y$，右边还有一个 $x$，左边另一个数不等于 $x$ 也不等于 $y$（不然就和上面的方案数重复了），方案数为
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

### 小技巧：用数组代替哈希表

把 $\textit{nums}$ 离散化，比如把 $[0,10,20,30]$ 压缩成 $[0,1,2,3]$。

这样可以用数组去统计元素个数，比哈希表更快。

> 注：这个优化对 Python 并不明显。

```java [sol-Java]
class Solution {
    public int subsequencesWithMiddleMode(int[] nums) {
        int n = nums.length;
        long ans = (long) n * (n - 1) * (n - 2) * (n - 3) * (n - 4) / 120; // 所有方案数

        int[] a = nums.clone();
        Arrays.sort(a);
        int m = 1;
        for (int i = 1; i < n; i++) {
            if (a[i] != a[i - 1]) {
                a[m++] = a[i]; // 原地去重
            }
        }
        for (int i = 0; i < n; i++) {
            nums[i] = Arrays.binarySearch(a, 0, m, nums[i]);
        }

        int[] suf = new int[m];
        for (int x : nums) {
            suf[x]++;
        }
        int[] pre = new int[m];
        // 枚举 x，作为子序列正中间的数
        for (int left = 0; left < n - 2; left++) {
            int x = nums[left];
            suf[x]--;
            if (left > 1) {
                int right = n - 1 - left;
                int preX = pre[x];
                int sufX = suf[x];
                // 不合法：只有一个 x
                ans -= (long) comb2(left - preX) * comb2(right - sufX);
                // 不合法：只有两个 x，且至少有两个 y（y != x）
                for (int y = 0; y < m; y++) {
                    if (y == x) {
                        continue;
                    }
                    int preY = pre[y];
                    int sufY = suf[y];
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
            pre[x]++;
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

        vector<int> a(nums);
        ranges::sort(a);
        a.erase(unique(a.begin(), a.end()), a.end()); // 去重
        for (int &x : nums) {
            x = ranges::lower_bound(a, x) - a.begin();
        }

        vector<int> pre(a.size()), suf(a.size());
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
                for (int y = 0; y < a.size(); y++) {
                    if (y == x) {
                        continue;
                    }
                    int pre_y = pre[y], suf_y = suf[y];
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

	a := slices.Clone(nums)
	slices.Sort(a)
	a = slices.Compact(a)
	for i, x := range nums {
		nums[i] = sort.SearchInts(a, x)
	}

	suf := make([]int, len(a))
	for _, x := range nums {
		suf[x]++
	}
	pre := make([]int, len(a))
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

## 方法二：式子变形

考虑把方法一的内层循环优化成 $\mathcal{O}(1)$。

内层循环有四条公式，我们一条一条分析。

### 1)

对于

$$
\dbinom {\textit{pre}_y} 2  \cdot \textit{suf}_x \cdot (n-1-i- \textit{suf}_x)
$$

内层循环的总和为

$$
\begin{aligned}
    & \sum_{y\ne x}\dbinom {\textit{pre}_y} 2  \cdot \textit{suf}_x \cdot (n-1-i- \textit{suf}_x)      \\
={} & \textit{suf}_x \cdot (n-1-i- \textit{suf}_x)\cdot \sum_{y\ne x}\dbinom {\textit{pre}_y} 2        \\
\end{aligned}
$$

其中

$$
\sum_{y\ne x}\dbinom {\textit{pre}_y} 2 = \left(\sum_{y}\dbinom {\textit{pre}_y} 2\right) - \dbinom {\textit{pre}_x} 2
$$

所以核心是维护

$$
\sum\limits_{y}\dbinom {\textit{pre}_y} 2
$$

如果 $\textit{pre}_y$ 增加了 $1$，那么上式增加了 

$$
\dbinom {\textit{pre}_y+1} 2 - \dbinom {\textit{pre}_y} 2 = \textit{pre}_y
$$

### 2)

对于

$$
\dbinom {\textit{suf}_y} 2  \cdot \textit{pre}_x \cdot (i- \textit{pre}_x)
$$

同理，核心是维护

$$
\sum\limits_{y}\dbinom {\textit{suf}_y} 2
$$

### 3)

对于

$$
\textit{pre}_y\cdot\textit{suf}_y\cdot\textit{pre}_x\cdot(n-1-i-\textit{suf}_x-\textit{suf}_y)
$$

设 $\textit{right}=n-1-i$，内层循环的总和为

$$
\begin{aligned}
& \sum_{y\ne x}\textit{pre}_y\cdot\textit{suf}_y\cdot\textit{pre}_x\cdot(\textit{right}-\textit{suf}_x-\textit{suf}_y)      \\
={} & \textit{pre}_x\cdot \sum_{y\ne x}\textit{pre}_y\cdot\textit{suf}_y\cdot(\textit{right}-\textit{suf}_x-\textit{suf}_y)       \\
={} & \textit{pre}_x\cdot \sum_{y\ne x}\textit{pre}_y\cdot\textit{suf}_y\cdot(\textit{right}-\textit{suf}_x)-\textit{pre}_y\cdot\textit{suf}_y^{\,2}       \\
={} & \textit{pre}_x\cdot \left((\textit{right}-\textit{suf}_x)\cdot \sum_{y\ne x}\textit{pre}_y\cdot\textit{suf}_y-\sum_{y\ne x}\textit{pre}_y\cdot\textit{suf}_y^{\,2}\right)       \\
\end{aligned}
$$

其中

$$
\sum_{y\ne x}\textit{pre}_y\cdot\textit{suf}_y = \left(\sum_{y}\textit{pre}_y\cdot\textit{suf}_y\right) - \textit{pre}_x\cdot\textit{suf}_x
$$

$$
\sum_{y\ne x}\textit{pre}_y\cdot\textit{suf}_y^{\,2} = \left(\sum_{y}\textit{pre}_y\cdot\textit{suf}_y^{\,2}\right) - \textit{pre}_x\cdot\textit{suf}_x^{\,2}
$$

所以核心是维护

$$
\sum_{y}\textit{pre}_y\cdot\textit{suf}_y
$$

和

$$
\sum_{y}\textit{pre}_y\cdot\textit{suf}_y^{\,2}
$$

如果 $\textit{pre}_y$ 增加了 $1$，那么上式增加了 $\textit{suf}_y^{\,2}$。

如果 $\textit{suf}_y$ 减少了 $1$，设 $\textit{sy}$ 为减少后的值，那么上式减少了 $\textit{pre}_y\cdot ((\textit{sy}+1)^2 - \textit{sy}^2) =\textit{pre}_y\cdot(\textit{sy}\cdot 2 + 1)$。

### 4)

对于

$$
\textit{pre}_y\cdot\textit{suf}_y\cdot\textit{suf}_x\cdot(i-\textit{pre}_x-\textit{pre}_y)
$$

同上，需要额外维护

$$
\sum_{y}\textit{pre}_y^2\cdot\textit{suf}_y
$$

遍历 $\textit{nums}$，更新 $\textit{pre}$ 和 $\textit{suf}$，同时 $\mathcal{O}(1)$ 维护上述和式，就可以把内层循环优化至 $\mathcal{O}(1)$。

|  变量名 |  含义  |
|---|---|
|  $\textit{px}$ | $\textit{pre}_x$  |
|  $\textit{sx}$ | $\textit{suf}_x$  |
|  $\textit{cp}$ | $\sum\limits_{y}\dbinom {\textit{pre}_y} 2$  |
|  $\textit{cs}$ | $\sum\limits_{y}\dbinom {\textit{suf}_y} 2$  |
|  $\textit{ps}$ | $\sum\limits_{y}\textit{pre}_y\cdot\textit{suf}_y$  |
|  $p2s$ |  $\sum\limits_{y}\textit{pre}_y^2\cdot\textit{suf}_y$ |
|  $ps2$ |  $\sum\limits_{y}\textit{pre}_y\cdot\textit{suf}_y^{\,2}$ |

```py [sol-Python3]
class Solution:
    def subsequencesWithMiddleMode(self, nums: List[int]) -> int:
        n = len(nums)
        suf = Counter(nums)
        pre = defaultdict(int)
        ans = comb(n, 5)  # 所有方案数

        cp = ps = p2s = ps2 = 0
        cs = sum(comb(c, 2) for c in suf.values())

        # 枚举 x，作为子序列正中间的数
        for left, x in enumerate(nums[:-2]):
            suf[x] -= 1

            px = pre[x]
            sx = suf[x]

            cs -= sx
            ps -= px
            p2s -= px * px
            ps2 -= (sx * 2 + 1) * px

            right = n - 1 - left
            ans -= comb(left - px, 2) * comb(right - sx, 2)
            ans -= (cp - comb(px, 2)) * sx * (right - sx)
            ans -= (cs - comb(sx, 2)) * px * (left - px)
            ans -= ((ps - px * sx) * (right - sx) - (ps2 - px * sx * sx)) * px
            ans -= ((ps - px * sx) * (left - px)  - (p2s - px * px * sx)) * sx

            cp += px
            ps += sx
            ps2 += sx * sx
            p2s += (px * 2 + 1) * sx

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

        int cp = 0, cs = 0, ps = 0, p2s = 0, ps2 = 0;
        for (int c : suf.values()) {
            cs += comb2(c);
        }

        // 枚举 x，作为子序列正中间的数
        for (int left = 0; left < n - 2; left++) {
            int x = nums[left];
            suf.merge(x, -1, Integer::sum); // suf[x]--

            int px = pre.getOrDefault(x, 0);
            int sx = suf.get(x);

            cs -= sx;
            ps -= px;
            p2s -= px * px;
            ps2 -= (sx * 2 + 1) * px;

            int right = n - 1 - left;
            ans -= (long) comb2(left - px) * comb2(right - sx);
            ans -= (long) (cp - comb2(px)) * sx * (right - sx);
            ans -= (long) (cs - comb2(sx)) * px * (left - px);
            ans -= (long) ((ps - px * sx) * (right - sx) - (ps2 - px * sx * sx)) * px;
            ans -= (long) ((ps - px * sx) * (left - px) - (p2s - px * px * sx)) * sx;

            cp += px;
            ps += sx;
            ps2 += sx * sx;
            p2s += (px * 2 + 1) * sx;

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

        int cp = 0, cs = 0, ps = 0, p2s = 0, ps2 = 0;
        for (auto& [_, c] : suf) {
            cs += comb2(c);
        }

        // 枚举 x，作为子序列正中间的数
        for (int left = 0; left < n - 2; left++) {
            int x = nums[left];
            suf[x]--;

            int px = pre[x];
            int sx = suf[x];

            cs -= sx;
            ps -= px;
            p2s -= px * px;
            ps2 -= (sx * 2 + 1) * px;

            int right = n - 1 - left;
            ans -= 1LL * comb2(left - px) * comb2(right - sx);
            ans -= 1LL * (cp - comb2(px)) * sx * (right - sx);
            ans -= 1LL * (cs - comb2(sx)) * px * (left - px);
            ans -= 1LL * ((ps - px * sx) * (right - sx) - (ps2 - px * sx * sx)) * px;
            ans -= 1LL * ((ps - px * sx) * (left - px)  - (p2s - px * px * sx)) * sx;

            cp += px;
            ps += sx;
            ps2 += sx * sx;
            p2s += (px * 2 + 1) * sx;

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
	for _, num := range nums {
		suf[num]++
	}
	pre := make(map[int]int, len(suf)) // 预分配空间

	var cp, cs, ps, p2s, ps2 int
	for _, c := range suf {
		cs += comb2(c)
	}

	// 枚举 x，作为子序列正中间的数
	for left, x := range nums[:n-2] {
		suf[x]--

		px := pre[x]
		sx := suf[x]

		cs -= sx
		ps -= px
		p2s -= px * px
		ps2 -= (sx*2 + 1) * px

		right := n - 1 - left
		ans -= comb2(left-px) * comb2(right-sx)
		ans -= (cp - comb2(px)) * sx * (right - sx)
		ans -= (cs - comb2(sx)) * px * (left - px)
		ans -= ((ps-px*sx)*(right-sx) - (ps2 - px*sx*sx)) * px
		ans -= ((ps-px*sx)*(left-px)  - (p2s - px*px*sx)) * sx

		cp += px
		ps += sx
		ps2 += sx * sx
		p2s += (px*2 + 1) * sx

		pre[x]++
	}
	return ans % 1_000_000_007
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

1. 改成 $\textit{seq}[0]$ 是众数。
2. 改成 $\textit{seq}[1]$ 是众数。
3. 把 $5$ 改成 $k=3,4,6,7$ 这些数呢？规定正中间的数的下标是 $\left\lfloor\dfrac{k}{2}\right\rfloor$。
4. 如果可以修改元素呢？每改一个数，就问你此时的答案。

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
