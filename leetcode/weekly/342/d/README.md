## 本题视频讲解

见[【周赛 342】](https://www.bilibili.com/video/BV1Bs4y1A7Wa/)第四题。

## 方法一：计算最短的 GCD 等于 1 的子数组

### 提示 1

首先，如果所有数的 GCD（最大公约数）大于 $1$，那么无论如何都无法操作出 $1$，我们返回 $-1$。

如果 $\textit{nums}$ 中有一个 $1$，那么从 $1$ 向左向右不断替换就能把所有数变成 $1$。

例如 $[2,2,1,2,2]\rightarrow[2,\underline{1},1,2,2]\rightarrow[\underline{1},1,1,2,2]\rightarrow[1,1,1,\underline{1},2]\rightarrow[1,1,1,1,\underline{1}]$，一共 $n-1=5-1=4$ 次操作。

如果有多个 $1$，那么每个 $1$ 只需要向左修改，最后一个 $1$ 向右修改剩余的数字。

例如 $[2,1,2,1,2]\rightarrow[\underline{1},1,2,1,2]\rightarrow[1,1,\underline{1},1,2]\rightarrow[1,1,1,1,\underline{1}]$，一共 $n-\textit{cnt}_1=5-2=3$ 次操作。这里 $\textit{cnt}_1$ 表示 $\textit{nums}$ 中 $1$ 的个数。

所以如果 $\textit{nums}$ 中有 $1$，答案为

$$
n-\textit{cnt}_1
$$

如果 $\textit{nums}$ 中没有 $1$ 呢？

### 提示 2

如果 $\textit{nums}$ 中没有 $1$，想办法花费尽量少的操作得出一个 $1$。

由于只能操作相邻的数，所以这个 $1$ 必然是一个连续子数组的 GCD。（如果在不连续的情况下得到了 $1$，那么这个 $1$ 只能属于其中某个连续子数组，其余的操作是多余的。）

那么找到最短的 GCD 为 $1$ 的子数组，设其长度为 $\textit{minSize}$，那么我们需要操作 $\textit{minSize}-1$ 次得到 $1$。

例如 $[2,6,3,4]$ 中的 $[3,4]$ 可以操作 $2-1=1$ 次得到 $1$。

然后就转化成提示 1 中的情况了，最终答案为

$$
(\textit{minSize}-1) + (n-1) = \textit{minSize}+n-2
$$

```py [sol1-Python3]
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        if gcd(*nums) > 1:
            return -1
        n = len(nums)
        cnt1 = sum(x == 1 for x in nums)
        if cnt1:
            return n - cnt1

        min_size = n
        for i in range(n):
            g = 0
            for j in range(i, n):
                g = gcd(g, nums[j])
                if g == 1:
                    # 这里本来是 j-i+1，把 +1 提出来合并到 return 中
                    min_size = min(min_size, j - i)
                    break
        return min_size + n - 1
```

```java [sol1-Java]
class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length, gcdAll = 0, cnt1 = 0;
        for (int x : nums) {
            gcdAll = gcd(gcdAll, x);
            if (x == 1) ++cnt1;
        }
        if (gcdAll > 1) return -1;
        if (cnt1 > 0) return n - cnt1;

        int minSize = n;
        for (int i = 0; i < n; ++i) {
            int g = 0;
            for (int j = i; j < n; ++j) {
                g = gcd(g, nums[j]);
                if (g == 1) {
                    // 这里本来是 j-i+1，把 +1 提出来合并到 return 中
                    minSize = Math.min(minSize, j - i);
                    break;
                }
            }
        }
        return minSize + n - 1;
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

```cpp [sol1-C++]
class Solution {
public:
    int minOperations(vector<int> &nums) {
        int n = nums.size(), gcd_all = 0, cnt1 = 0;
        for (int x: nums) {
            gcd_all = gcd(gcd_all, x);
            cnt1 += x == 1;
        }
        if (gcd_all > 1) return -1;
        if (cnt1) return n - cnt1;

        int min_size = n;
        for (int i = 0; i < n; ++i) {
            int g = 0;
            for (int j = i; j < n; ++j) {
                g = gcd(g, nums[j]);
                if (g == 1) {
                    // 这里本来是 j-i+1，把 +1 提出来合并到 return 中
                    min_size = min(min_size, j - i);
                    break;
                }
            }
        }
        return min_size + n - 1;
    }
};
```

```go [sol1-Go]
func minOperations(nums []int) int {
	n, gcdAll, cnt1 := len(nums), 0, 0
	for _, x := range nums {
		gcdAll = gcd(gcdAll, x)
		if x == 1 {
			cnt1++
		}
	}
	if gcdAll > 1 {
		return -1
	}
	if cnt1 > 0 {
		return n - cnt1
	}

	minSize := n
	for i := range nums {
		g := 0
		for j, x := range nums[i:] {
			g = gcd(g, x)
			if g == 1 {
				// 这里本来是 j+1，把 +1 提出来合并到 return 中
				minSize = min(minSize, j)
				break
			}
		}
	}
	return minSize + n - 1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+\log U))$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。外层循环时，单看 $g=\textit{nums}[i]$，它因为求 GCD 减半的次数是 $\mathcal{O}(\log U)$ 次，因此内层循环的时间复杂度为 $\mathcal{O}(n+\log U)$，所以总的时间复杂度为 $\mathcal{O}(n(n+\log U))$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 方法二：利用 GCD 的性质

如果 $n=10^5$ 要怎么做？

原理见我之前写的 [这篇题解的方法二](https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/solution/by-endlesscheng-zai1/)，或者看开头贴的视频链接。

```py [sol2-Python3]
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        if gcd(*nums) > 1:
            return -1
        n = len(nums)
        cnt1 = sum(x == 1 for x in nums)
        if cnt1:
            return n - cnt1

        min_size = n
        a = []  # [GCD，相同 GCD 闭区间的右端点]
        for i, x in enumerate(nums):
            a.append([x, i])

            # 原地去重，因为相同的 GCD 都相邻在一起
            j = 0
            for p in a:
                p[0] = gcd(p[0], x)
                if a[j][0] != p[0]:
                    j += 1
                    a[j] = p
                else:
                    a[j][1] = p[1]
            del a[j + 1:]

            if a[0][0] == 1:
                # 这里本来是 i-a[0][1]+1，把 +1 提出来合并到 return 中
                min_size = min(min_size, i - a[0][1])
        return min_size + n - 1
```

```java [sol2-Java]
class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length, gcdAll = 0, cnt1 = 0;
        for (int x : nums) {
            gcdAll = gcd(gcdAll, x);
            if (x == 1) ++cnt1;
        }
        if (gcdAll > 1) return -1;
        if (cnt1 > 0) return n - cnt1;

        int minSize = n;
        var g = new ArrayList<int[]>(); // [GCD，相同 GCD 闭区间的右端点]
        for (int i = 0; i < n; ++i) {
            g.add(new int[]{nums[i], i});
            // 原地去重，因为相同的 GCD 都相邻在一起
            var j = 0;
            for (var p : g) {
                p[0] = gcd(p[0], nums[i]);
                if (g.get(j)[0] == p[0])
                    g.get(j)[1] = p[1]; // 合并相同值，下标取最小的
                else g.set(++j, p);
            }
            g.subList(j + 1, g.size()).clear();
            if (g.get(0)[0] == 1)
                // 这里本来是 i-g.get(0)[1]+1，把 +1 提出来合并到 return 中
                minSize = Math.min(minSize, i - g.get(0)[1]);
        }
        return minSize + n - 1;
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

```cpp [sol2-C++]
class Solution {
public:
    int minOperations(vector<int> &nums) {
        int n = nums.size(), gcd_all = 0, cnt1 = 0;
        for (int x: nums) {
            gcd_all = gcd(gcd_all, x);
            cnt1 += x == 1;
        }
        if (gcd_all > 1) return -1;
        if (cnt1) return n - cnt1;

        int min_size = n;
        vector<pair<int, int>> g; // {GCD，相同 GCD 闭区间的右端点}
        for (int i = 0; i < n; ++i) {
            g.emplace_back(nums[i], i);
            // 原地去重，因为相同的 GCD 都相邻在一起
            int j = 0;
            for (auto &p: g) {
                p.first = gcd(p.first, nums[i]);
                if (g[j].first == p.first)
                    g[j].second = p.second;
                else g[++j] = move(p);
            }
            g.resize(j + 1);
            if (g[0].first == 1)
                // 这里本来是 i-g[0].second+1，把 +1 提出来合并到 return 中
                min_size = min(min_size, i - g[0].second);
        }
        return min_size + n - 1;
    }
};
```

```go [sol2-Go]
func minOperations(nums []int) int {
	n, gcdAll, cnt1 := len(nums), 0, 0
	for _, x := range nums {
		gcdAll = gcd(gcdAll, x)
		if x == 1 {
			cnt1++
		}
	}
	if gcdAll > 1 {
		return -1
	}
	if cnt1 > 0 {
		return n - cnt1
	}

	minSize := n
	type result struct{ gcd, i int }
	a := []result{}
	for i, x := range nums {
		for j, r := range a {
			a[j].gcd = gcd(r.gcd, x)
		}
		a = append(a, result{x, i})

		// 去重
		j := 0
		for _, q := range a[1:] {
			if a[j].gcd != q.gcd {
				j++
				a[j] = q
			} else {
				a[j].i = q.i // 相同 gcd 保存最右边的位置
			}
		}
		a = a[:j+1]

		if a[0].gcd == 1 {
			// 这里本来是 i-a[0].i+1，把 +1 提出来合并到 return 中
			minSize = min(minSize, i-a[0].i)
		}
	}
	return minSize + n - 1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。单看每个元素，它因为求 GCD 减半的次数是 $\mathcal{O}(\log U)$ 次，并且每次去重的时间复杂度也为 $\mathcal{O}(\log U)$，因此时间复杂度为 $\mathcal{O}(n\log U)$。
- 空间复杂度：$\mathcal{O}(\log U)$。

> 注：由于本题数据范围小，这两种做法的运行时间区别并不明显。

### 可以用该模板秒杀的题目

按位或：

- [898. 子数组按位或操作](https://leetcode.cn/problems/bitwise-ors-of-subarrays/)

按位与：

- [1521. 找到最接近目标值的函数值](https://leetcode.cn/problems/find-a-value-of-a-mysterious-function-closest-to-target/)

GCD：

- [Codeforces 475D. CGCDSSQ](https://codeforces.com/problemset/problem/475/D)
- [Codeforces 1632D. New Year Concert](https://codeforces.com/problemset/problem/1632/D)

乘法：

- [蓝桥杯 2021 年第十二届国赛真题 - 和与乘积](https://www.dotcpp.com/oj/problem2622.html)
