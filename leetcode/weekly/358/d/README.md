请看 [视频讲解](https://www.bilibili.com/video/BV1wh4y1Q7XW/) 第四题。

本文接着 [贡献法+单调栈](https://leetcode.cn/problems/sum-of-subarray-minimums/solution/gong-xian-fa-dan-diao-zhan-san-chong-shi-gxa5/) 继续讲，请先阅读这篇题解，因为核心逻辑是一样的。

贪心地说，先考虑 $\textit{nums}$ 中最大的元素 $x$，我们需要知道：有多少个子数组可以取 $x$ 作为质数分数最高的元素。

我们可以先把 $[1,10^5]$ 中的每个数的质数分数（不同质因子的数目）预处理出来，记作数组 $\textit{omega}$。

然后用单调栈求出每个 $i$ 左侧质数分数【大于等于】$\textit{omega}[\textit{nums}[i]]$ 的最近的数的下标 $\textit{left}[i]$（不存在则为 $-1$），以及右侧质数分数【大于】$\textit{omega}[\textit{nums}[i]]$ 的最近的数的下标 $\textit{right}[i]$（不存在则为 $n$）。

请注意，我们不能在 $i$ 左侧包含质数分数和 $\textit{omega}[\textit{nums}[i]]$ 一样的数，否则不满足题目「如果多个元素质数分数相同且最高，选择下标最小的一个」的要求。所以对于左侧用【大于等于】。

那么子数组的左端点可以在 $(\textit{left}[i],i]$ 内，子数组的右端点可以在 $[i,\textit{right}[i])$ 内。

根据 [乘法原理](https://baike.baidu.com/item/%E4%B9%98%E6%B3%95%E5%8E%9F%E7%90%86/7538447)，一共有

$$
\textit{tot} = (i-\textit{left}[i])\cdot (\textit{right}[i]-i)
$$

个子数组以 $\textit{nums}[i]$ 作为这个子数组的质数分数。

设 $k'=\min(k, \textit{tot})$，则 $\textit{nums}[i]$ 对答案的贡献为

$$
\textit{nums}[i] ^ {k'}
$$

上式可以用快速幂计算，具体请看 [50. Pow(x, n)](https://leetcode.cn/problems/powx-n/)。

根据上式，按照 $\textit{nums}$ 从大到小的顺序计算贡献，一边计算一边更新剩余操作次数 $k$。

```py [sol-Python3]
MOD = 10 ** 9 + 7
MX = 10 ** 5 + 1
omega = [0] * MX
for i in range(2, MX):  # 预处理 omega
    if omega[i] == 0:  # i 是质数
        for j in range(i, MX, i):
            omega[j] += 1  # i 是 j 的一个质因子

class Solution:
    def maximumScore(self, nums: List[int], k: int) -> int:
        n = len(nums)
        left = [-1] * n  # 质数分数 >= omega[nums[i]] 的左侧最近元素下标
        right = [n] * n  # 质数分数 >  omega[nums[i]] 的右侧最近元素下标
        st = []
        for i, v in enumerate(nums):
            while st and omega[nums[st[-1]]] < omega[v]:
                right[st.pop()] = i
            if st: left[i] = st[-1]
            st.append(i)

        ans = 1
        for i, v, l, r in sorted(zip(range(n), nums, left, right), key=lambda z: -z[1]):
            tot = (i - l) * (r - i)
            if tot >= k:
                ans = ans * pow(v, k, MOD) % MOD
                break
            ans = ans * pow(v, tot, MOD) % MOD
            k -= tot  # 更新剩余操作次数
        return ans
```

```java [sol-Java]
class Solution {
    private static final long MOD = (long) 1e9 + 7;
    private static final int MX = (int) 1e5 + 1;
    private static final int[] omega = new int[MX];

    static {
        for (int i = 2; i < MX; i++)
            if (omega[i] == 0) // i 是质数
                for (int j = i; j < MX; j += i)
                    omega[j]++; // i 是 j 的一个质因子
    }

    public int maximumScore(List<Integer> nums, int k) {
        var a = nums.stream().mapToInt(i -> i).toArray();
        int n = a.length;
        var left = new int[n]; // 质数分数 >= omega[nums[i]] 的左侧最近元素下标
        var right = new int[n];// 质数分数 >  omega[nums[i]] 的右侧最近元素下标
        Arrays.fill(right, n);
        var st = new ArrayDeque<Integer>();
        st.push(-1); // 方便赋值 left
        for (int i = 0; i < n; i++) {
            while (st.size() > 1 && omega[a[st.peek()]] < omega[a[i]])
                right[st.pop()] = i;
            left[i] = st.peek();
            st.push(i);
        }

        var ids = new Integer[n];
        for (int i = 0; i < n; i++) ids[i] = i;
        Arrays.sort(ids, (i, j) -> a[j] - a[i]);

        long ans = 1;
        for (int i : ids) {
            long tot = (long) (i - left[i]) * (right[i] - i);
            if (tot >= k) {
                ans = ans * pow(a[i], k) % MOD;
                break;
            }
            ans = ans * pow(a[i], (int) tot) % MOD;
            k -= tot; // 更新剩余操作次数
        }
        return (int) ans;
    }

    private long pow(long x, int n) {
        var res = 1L;
        for (; n > 0; n /= 2) {
            if (n % 2 > 0) res = res * x % MOD;
            x = x * x % MOD;
        }
        return res;
    }
}
```

```cpp [sol-C++]
const int MX = 1e5 + 1;
int omega[MX];
int init = []() {
    for (int i = 2; i < MX; i++)
        if (omega[i] == 0) // i 是质数
            for (int j = i; j < MX; j += i)
                omega[j]++; // i 是 j 的一个质因子
    return 0;
}();

class Solution {
    const long long MOD = 1e9 + 7;

    long long pow(long long x, int n) {
        long long res = 1;
        for (; n; n /= 2) {
            if (n % 2) res = res * x % MOD;
            x = x * x % MOD;
        }
        return res;
    }

public:
    int maximumScore(vector<int> &nums, int k) {
        int n = nums.size();
        vector<int> left(n, -1); // 质数分数 >= omega[nums[i]] 的左侧最近元素下标
        vector<int> right(n, n); // 质数分数 >  omega[nums[i]] 的右侧最近元素下标
        stack<int> st;
        for (int i = 0; i < n; i++) {
            while (!st.empty() && omega[nums[st.top()]] < omega[nums[i]]) {
                right[st.top()] = i;
                st.pop();
            }
            if (!st.empty()) left[i] = st.top();
            st.push(i);
        }

        vector<int> id(n);
        iota(id.begin(), id.end(), 0);
        sort(id.begin(), id.end(), [&](const int i, const int j) {
            return nums[i] > nums[j];
        });

        long long ans = 1;
        for (int i: id) {
            long long tot = (long long) (i - left[i]) * (right[i] - i);
            if (tot >= k) {
                ans = ans * pow(nums[i], k) % MOD;
                break;
            }
            ans = ans * pow(nums[i], tot) % MOD;
            k -= tot; // 更新剩余操作次数
        }
        return ans;
    }
};
```

```go [sol-Go]
const mod int = 1e9 + 7

// 预处理 omega
const mx int = 1e5
var omega [mx + 1]int8
func init() {
	for i := 2; i <= mx; i++ {
		if omega[i] == 0 { // i 是质数
			for j := i; j <= mx; j += i {
				omega[j]++ // i 是 j 的一个质因子
			}
		}
	}
}

func maximumScore(nums []int, k int) int {
	n := len(nums)
	left := make([]int, n)  // 质数分数 >= omega[nums[i]] 的左侧最近元素下标
	right := make([]int, n) // 质数分数 >  omega[nums[i]] 的右侧最近元素下标
	for i := range right {
		right[i] = n
	}
	st := []int{-1}
	for i, v := range nums {
		for len(st) > 1 && omega[nums[st[len(st)-1]]] < omega[v] {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return nums[id[i]] > nums[id[j]] })

	ans := 1
	for _, i := range id {
		tot := (i - left[i]) * (right[i] - i)
		if tot >= k {
			ans = ans * pow(nums[i], k) % mod
			break
		}
		ans = ans * pow(nums[i], tot) % mod
		k -= tot // 更新剩余操作次数
	}
	return ans
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

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。这里忽略预处理 $\textit{omega}$ 的时间和空间。
- 空间复杂度：$\mathcal{O}(n)$。

#### 相似题目：单调栈+贡献法

- [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/)
- [1856. 子数组最小乘积的最大值](https://leetcode.cn/problems/maximum-subarray-min-product/)
- [2104. 子数组范围和](https://leetcode.cn/problems/sum-of-subarray-ranges/)
- [2281. 巫师的总力量和](https://leetcode.cn/problems/sum-of-total-strength-of-wizards/)
