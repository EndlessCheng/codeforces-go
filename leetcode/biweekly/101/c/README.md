### 视频讲解

见[【双周赛 101】](https://www.bilibili.com/video/BV1Ga4y1M72A/)。

视频还讲解了如何手算二元一次不定方程！欢迎点赞！

---

为方便描述，将 $\textit{arr}$ 简记为 $a$。

### 提示 1

先来解决 $a$ 不是循环数组的情况。

根据题意，考虑从 $i$ 和 $i+1$ 开始的两个长为 $k$ 的子数组的和，如果要求这两个和相等，则有

$$
a[i]+a[i+1]+\cdots + a[i+k-1] = a[i+1]+a[i+2]+\cdots + a[i+k]
$$

化简得

$$
a[i] = a[i+k]
$$

换句话说：

- $a[0] = a[k] = a[2k] = \cdots$
- $a[1] = a[k+1] = a[2k+1] = \cdots$
- $a[2] = a[k+2] = a[2k+2] = \cdots$
- ……

### 提示 2

按照 $i\bmod k$ 的结果将 $a$ 分组，对每一组（记作 $b$），我们需要解决：

让数组 $b$ 的所有元素相等的最少运算次数。

根据**中位数贪心**，将 $b$ 的所有元素变为 $b$ 的中位数是最优的。

证明：设 $b$ 的长度为 $m$，设要将所有 $b[i]$ 变为 $x$。假设 $b$ 已经从小到大排序。首先，如果 $x$ 取在区间 $[b[0],b[m-1]]$ 之外，那么 $x$ 向区间方向移动可以使距离和变小；同时，如果 $x$ 取在区间 $[b[0],b[m-1]]$ 之内，无论如何移动 $x$，它到 $b[0]$ 和 $b[m-1]$ 的距离和都是一个定值 $b[m-1]-b[0]$，那么去掉 $b[0]$ 和 $b[m-1]$ 这两个最左最右的数，问题规模缩小。不断缩小问题规模，如果最后剩下 $1$ 个数，那么 $x$ 就取它；如果最后剩下 $2$ 个数，那么 $x$ 取这两个数之间的任意值都可以（包括这两个数）。因此 $x$ 可以取 $b[m/2]$。

### 提示 3

回到原问题。

比如 $n=6,k=4$，那么 $a[2]$ 循环后是 $a[8]$，和 $a[0]$ 在同一组，而 $a[1]$ 无论怎么循环都无法和 $a[0]$ 在同一组。（$(1+6n)\bmod 4 \ne 0$）

根据这个例子，可以猜想一个结论：

**一个循环数组如果既有周期 $n$，又有周期 $k$，则必然有周期 $\gcd(n,k)$。**

证明：根据 [裴蜀定理](https://oi-wiki.org/math/number-theory/bezouts/)，有

$$
a[i] = a[i+nx+ky] = a[i+\gcd(n,k)]
$$

这样就转换成了不是循环数组的情况。

> 注：代码中的排序可以换成快速选择，从而做到 $O(n)$ 的时间复杂度。具体见 C++ 代码。

```py [sol1-Python3]
class Solution:
    def makeSubKSumEqual(self, arr: List[int], k: int) -> int:
        k = gcd(k, len(arr))
        ans = 0
        for i in range(k):
            b = sorted(arr[i::k])
            mid = b[len(b) // 2]
            ans += sum(abs(x - mid) for x in b)
        return ans
```

```java [sol1-Java]
class Solution {
    public long makeSubKSumEqual(int[] arr, int k) {
        int n = arr.length;
        k = gcd(k, n);
        long ans = 0;
        for (int i = 0; i < k; ++i) {
            var b = new ArrayList<Integer>();
            for (int j = i; j < n; j += k)
                b.add(arr[j]);
            Collections.sort(b);
            int mid = b.get(b.size() / 2);
            for (int x : b)
                ans += Math.abs(x - mid);
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

```cpp [sol1-C++]
class Solution {
public:
    long long makeSubKSumEqual(vector<int> &arr, int k) {
        int n = arr.size();
        k = gcd(k, n);
        vector<vector<int>> g(k);
        for (int i = 0; i < n; ++i)
            g[i % k].push_back(arr[i]);

        long long ans = 0;
        for (auto &b: g) {
            nth_element(b.begin(), b.begin() + b.size() / 2, b.end());
            for (int x: b)
                ans += abs(x - b[b.size() / 2]);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func makeSubKSumEqual(arr []int, k int) (ans int64) {
	k = gcd(k, len(arr))
	g := make([][]int, k)
	for i, x := range arr {
		g[i%k] = append(g[i%k], x)
	}
	for _, b := range g {
		sort.Ints(b)
		for _, x := range b {
			ans += int64(abs(x - b[len(b)/2]))
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$ 或 $O(n)$，其中 $n$ 为 $\textit{arr}$ 的长度。采用快速选择找中位数可以做到 $O(n)$，见 C++ 代码。
- 空间复杂度：$O(n)$。

### 相似题目

- [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
