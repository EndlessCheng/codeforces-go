## 前置知识

[背包问题 & 空间压缩【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)

## 哪种背包？

- 如果区分相同数字，就是 0-1 背包。
- 如果不区分相同数字，就是多重背包。

## 朴素 DP

用哈希表统计每个数的出现次数，记在 $\textit{cnt}$ 中。

定义 $f[i][j]$ 表示从 $\textit{cnt}$ 的前 $i$ 种数中选择一些数，元素和**恰好**为 $j$ 的方案数。

设第 $i$ 种数的值为 $x$。

枚举第 $i$ 种数选了 $k=0,1,2,\cdots,\textit{cnt}[x]$ 个，选完之后，剩下要解决的子问题，就是从 $\textit{cnt}$ 的前 $i-1$ 种数中选择一些数，元素和恰好为 $j-kx$ 的方案数，即 $f[i-1][j-kx]$。

根据加法原理，累加这些方案数，即

$$
f[i][j] = \sum_{k=0}^{\textit{cnt}[x]} f[i-1][j-kx]
$$

注意必须满足 $j-kx\ge 0$。

初始值 $f[0][0] = \textit{cnt}[0] + 1$，表示「在什么也不选的情况下，元素和为 $0$」有 $\textit{cnt}[0] + 1$ 个方案。

答案为 

$$
\sum_{i=l}^r f[m][i]
$$

其中 $m$ 是 $\textit{cnt}$ 的大小（不包括 $0$）。

## 优化方法一：式子变形+滚动数组

举例说明，假设 $x=2$ 并且 $\textit{cnt}[x]=3$，那么选 $0,1,2,3$ 个 $x$ 都是可以的，即

$$
f[i][j] = f[i-1][j] + f[i-1][j-2] + f[i-1][j-4] + f[i-1][j-6]
$$

我们把 $f[i][j-2]$ 的递推式也列出来看看：

$$
f[i][j-2] = f[i-1][j-2] + f[i-1][j-4] + f[i-1][j-6] + f[i-1][j-8]
$$

注意到，中间的 $f[i-1][j-2] + f[i-1][j-4] + f[i-1][j-6]$ 算了两遍，这可以优化掉：

$$
\begin{aligned}
f[i][j] =\ & f[i-1][j] + f[i-1][j-2] + f[i-1][j-4] + f[i-1][j-6]\\
=\ &f[i-1][j] + (f[i][j-2] - f[i-1][j-8])\\
=\ &f[i][j-2] + f[i-1][j] - f[i-1][j-8]\\
\end{aligned}
$$

一般地，我们有

$$
f[i][j] = f[i][j-x] + f[i-1][j] - f[i-1][j-(\textit{cnt}[x]+1)\cdot x]
$$

如果 $j-(\textit{cnt}[x]+1)\cdot x < 0$，由于无法选一些数组成负数，对应的方案数为 $0$，所以

$$
f[i][j] = f[i][j-x] + f[i-1][j]
$$

这样就可以 $\mathcal{O}(1)$ 地算出每个 $f[i][j]$ 了。

代码实现时，可以用滚动数组优化，具体请看上面贴的背包问题的视频。

> 关于取模的原理，见文末的「算法小课堂」。

```py [sol-Python3]
class Solution:
    def countSubMultisets(self, nums: List[int], l: int, r: int) -> int:
        MOD = 10 ** 9 + 7
        total = sum(nums)
        if l > total:
            return 0

        r = min(r, total)
        cnt = Counter(nums)
        f = [cnt[0] + 1] + [0] * r
        del cnt[0]

        s = 0
        for x, c in cnt.items():
            new_f = f.copy()
            s = min(s + x * c, r)  # 到目前为止，能选的元素和至多为 s
            for j in range(x, s + 1):  # 把循环上界从 r 改成 s，能快一倍
                new_f[j] += new_f[j - x]
                if j >= (c + 1) * x:
                    new_f[j] -= f[j - (c + 1) * x]
                new_f[j] %= MOD
            f = new_f
        return sum(f[l:]) % MOD
```

```java [sol-Java]
class Solution {
    public int countSubMultisets(List<Integer> nums, int l, int r) {
        final int MOD = 1_000_000_007;
        int total = 0;
        var cnt = new HashMap<Integer, Integer>();
        for (int x : nums) {
            total += x;
            cnt.merge(x, 1, Integer::sum);
        }
        if (l > total) {
            return 0;
        }

        r = Math.min(r, total);
        int[] f = new int[r + 1];
        f[0] = cnt.getOrDefault(0, 0) + 1;
        cnt.remove(0);

        int sum = 0;
        for (var e : cnt.entrySet()) {
            int x = e.getKey(), c = e.getValue();
            int[] newF = f.clone();
            sum = Math.min(sum + x * c, r); // 到目前为止，能选的元素和至多为 sum
            for (int j = x; j <= sum; j++) { // 把循环上界从 r 改成 sum 可以快不少
                newF[j] = (newF[j] + newF[j - x]) % MOD;
                if (j >= (c + 1) * x) {
                    newF[j] = (newF[j] - f[j - (c + 1) * x] + MOD) % MOD; // 避免减法产生负数
                }
            }
            f = newF;
        }

        int ans = 0;
        for (int i = l; i <= r; ++i) {
            ans = (ans + f[i]) % MOD;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countSubMultisets(vector<int> &nums, int l, int r) {
        const int MOD = 1e9 + 7;
        int total = 0;
        unordered_map<int, int> cnt;
        for (int x: nums) {
            total += x;
            cnt[x]++;
        }
        if (l > total) {
            return 0;
        }

        r = min(r, total);
        vector<int> f(r + 1);
        f[0] = cnt[0] + 1;
        cnt.erase(0);

        int sum = 0;
        for (auto [x, c]: cnt) {
            auto new_f = f;
            sum = min(sum + x * c, r); // 到目前为止，能选的元素和至多为 sum
            for (int j = x; j <= sum; j++) { // 把循环上界从 r 改成 sum 可以快不少
                new_f[j] = (new_f[j] + new_f[j - x]) % MOD;
                if (j >= (c + 1) * x) {
                    new_f[j] = (new_f[j] - f[j - (c + 1) * x] + MOD) % MOD; // 避免减法产生负数
                }
            }
            f = move(new_f);
        }

        int ans = 0;
        for (int i = l; i <= r; i++) {
            ans = (ans + f[i]) % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubMultisets(nums []int, l, r int) (ans int) {
	const mod = 1_000_000_007
	total := 0
	cnt := map[int]int{}
	for _, x := range nums {
		total += x
		cnt[x]++
	}
	if l > total {
		return
	}

	r = min(r, total)
	f := make([]int, r+1)
	f[0] = cnt[0] + 1
	delete(cnt, 0)

	sum := 0
	for x, c := range cnt {
		newF := append([]int{}, f...)
		sum = min(sum+x*c, r) // 到目前为止，能选的元素和至多为 sum
		for j := x; j <= sum; j++ { // 把循环上界从 r 改成 sum 可以快不少
			newF[j] += newF[j-x]
			if j >= (c+1)*x {
				newF[j] -= f[j-(c+1)*x] // 注意这里有减法，可能产生负数
			}
			newF[j] %= mod
		}
		f = newF
	}

	for _, v := range f[l:] {
		ans += v
	}
	return (ans%mod + mod) % mod // 调整成非负数
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(S\min(\sqrt{S},n))$，其中 $n$ 为 $\textit{nums}$ 的长度，$S$ 为 $\textit{nums}$ 的元素和。（下面的分析不考虑 $0$）如果哈希表中只有一个 key，那么花费 $\mathcal{O}(S)$ 的时间就算完了。想一想，哈希表最多能有多少个 key？是 $n$ 个吗？不对，如果有 $n$ 个 key，由于这些 key 互不相同，这些 key 的元素之和最少是 $1+2+\cdots+n=\dfrac{n(n+1)}{2}$，但是数组的元素和为 $S$，所以 $\dfrac{n(n+1)}{2}\le S$，所以哈希表至多约有 $\sqrt{2S}$ 个 key，在本题的约束下大约是 $200$ 个。当然，哈希表的大小不会超过 $n$，所以哈希表至多有 $\mathcal{O}(\min(\sqrt{S},n))$ 个 key。
- 空间复杂度：$\mathcal{O}(S)$。

## 优化方法二：同余前缀和优化

再看一眼这个式子

$$
f[i][j] = \sum_{k=0}^{\textit{cnt}[x]} f[i-1][j-kx]
$$

如果求出 $f[i-1]$ 的**同余前缀和**，那么 $f[i][j]$ 就可以转换成两个同余前缀和的差了，这样也可以 $\mathcal{O}(1)$ 地算出每个 $f[i][j]$。

同余前缀和是指 

$$
s[p] = f[i-1][p] + f[i-1][p-x] + f[i-1][p-2x] +\cdots
$$

这可以递推算出来，即

$$
s[p] = f[i-1][p] + s[p-x]
$$

代码实现时，可以只用一个一维数组。

```py [sol-Python3]
class Solution:
    def countSubMultisets(self, nums: List[int], l: int, r: int) -> int:
        MOD = 10 ** 9 + 7
        total = sum(nums)
        if l > total:
            return 0

        r = min(r, total)
        cnt = Counter(nums)
        f = [cnt[0] + 1] + [0] * r
        del cnt[0]

        s = 0
        for x, c in cnt.items():
            s = min(s + x * c, r)
            for j in range(x, s + 1):
                f[j] = (f[j] + f[j - x]) % MOD  # 原地计算同余前缀和
            t = (c + 1) * x
            for j in range(s, t - 1, -1):
                f[j] = (f[j] - f[j - t]) % MOD  # 两个同余前缀和的差
        return sum(f[l:]) % MOD
```

```java [sol-Java]
class Solution {
    public int countSubMultisets(List<Integer> nums, int l, int r) {
        final int MOD = 1_000_000_007;
        int total = 0;
        var cnt = new HashMap<Integer, Integer>();
        for (int x : nums) {
            total += x;
            cnt.merge(x, 1, Integer::sum);
        }
        if (l > total) {
            return 0;
        }

        r = Math.min(r, total);
        int[] f = new int[r + 1];
        f[0] = cnt.getOrDefault(0, 0) + 1;
        cnt.remove(0);

        int sum = 0;
        for (var e : cnt.entrySet()) {
            int x = e.getKey(), c = e.getValue();
            sum = Math.min(sum + x * c, r);
            for (int j = x; j <= sum; j++) {
                f[j] = (f[j] + f[j - x]) % MOD; // 原地计算同余前缀和
            }
            for (int j = sum; j >= x * (c + 1); j--) {
                f[j] = (f[j] - f[j - x * (c + 1)] + MOD) % MOD; // 两个同余前缀和的差
            }
        }

        int ans = 0;
        for (int i = l; i <= r; ++i) {
            ans = (ans + f[i]) % MOD;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countSubMultisets(vector<int> &nums, int l, int r) {
        const int MOD = 1e9 + 7;
        int total = 0;
        unordered_map<int, int> cnt;
        for (int x: nums) {
            total += x;
            cnt[x]++;
        }
        if (l > total) {
            return 0;
        }

        r = min(r, total);
        vector<int> f(r + 1);
        f[0] = cnt[0] + 1;
        cnt.erase(0);

        int sum = 0;
        for (auto [x, c]: cnt) {
            sum = min(sum + x * c, r);
            for (int j = x; j <= sum; j++) {
                f[j] = (f[j] + f[j - x]) % MOD; // 原地计算同余前缀和
            }
            for (int j = sum; j >= x * (c + 1); j--) {
                f[j] = (f[j] - f[j - x * (c + 1)] + MOD) % MOD; // 两个同余前缀和的差
            }
        }

        int ans = 0;
        for (int i = l; i <= r; i++) {
            ans = (ans + f[i]) % MOD;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubMultisets(nums []int, l, r int) (ans int) {
	const mod = 1_000_000_007
	total := 0
	cnt := map[int]int{}
	for _, x := range nums {
		total += x
		cnt[x]++
	}
	if l > total {
		return
	}

	r = min(r, total)
	f := make([]int, r+1)
	f[0] = cnt[0] + 1
	delete(cnt, 0)

	sum := 0
	for x, c := range cnt {
		sum = min(sum+x*c, r)
		for j := x; j <= sum; j++ {
			f[j] = (f[j] + f[j-x]) % mod // 原地计算同余前缀和
		}
		for j := sum; j >= x*(c+1); j-- {
			f[j] = (f[j] - f[j-x*(c+1)]) % mod // 两个同余前缀和的差
		}
	}

	for _, v := range f[l:] {
		ans += v
	}
	return (ans%mod + mod) % mod // 调整成非负数
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析

同上。

## 算法小课堂：模运算

如果让你计算 $1234\cdot 6789$ 的**个位数**，你会如何计算？

由于只有个位数会影响到乘积的个位数，那么 $4\cdot 9=36$ 的个位数 $6$ 就是答案。

对于 $1234+6789$ 的个位数，同理，$4+9=13$ 的个位数 $3$ 就是答案。

你能把这个结论抽象成数学等式吗？

一般地，涉及到取模的题目，通常会用到如下等式（上面计算的是 $m=10$）：

$$
(a+b)\bmod m = ((a\bmod m) + (b\bmod m)) \bmod m
$$

$$
(a\cdot b) \bmod m=((a\bmod m)\cdot  (b\bmod m)) \bmod m
$$

证明：根据**带余除法**，任意整数 $a$ 都可以表示为 $a=km+r$，这里 $r$ 相当于 $a\bmod m$。那么设 $a=k_1m+r_1,\ b=k_2m+r_2$。

第一个等式：

$$
\begin{aligned}
&\ (a+b) \bmod m\\
=&\ ((k_1+k_2) m+r_1+r_2)\bmod m\\
=&\ (r_1+r_2)\bmod m\\
=&\ ((a\bmod m) + (b\bmod m)) \bmod m
\end{aligned}
$$

第二个等式：

$$
\begin{aligned}
&\ (a\cdot b) \bmod m\\
=&\ (k_1k_2m^2+(k_1r_2+k_2r_1)m+r_1r_2)\bmod m\\
=&\ (r_1r_2)\bmod m\\
=&\ ((a\bmod m)\cdot  (b\bmod m)) \bmod m
\end{aligned}
$$

**根据这两个恒等式，可以随意地对代码中的加法和乘法的结果取模**。
