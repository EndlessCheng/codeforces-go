#### 提示 1-1

枚举每位巫师，假设他是最弱的巫师，那么他能在哪些子数组中？

#### 提示 1-2

左右边界最远能到哪？具体地，这些子数组的左边界的最小值是多少，右边界的最大值是多少？

#### 提示 1-3

用**单调栈**来计算左右边界。

不了解单调栈的同学可以看一下 [496. 下一个更大元素 I](https://leetcode-cn.com/problems/next-greater-element-i/)。对于本题，我们需要求的是下一个更小元素。

#### 提示 1-4

注意本题是可能有**重复元素**的，这会对最终答案的计算产生什么影响？

#### 提示 1-5

为了避免重复计算，我们可以考虑左侧**严格小于**当前元素的最近元素位置，以及右侧**小于等于**当前元素的最近元素位置。

以示例 1 中的数组 $[1,3,1,2]$ 为例，如果左右两侧都是找严格小于，那么第一个 $1$ 和第二个 $1$ 算出来的边界范围都是一样的（都是整个数组），这就重复统计了，为了避免这种情况，可以把某一侧改为小于等于，比如把右侧改成小于等于，那么第一个 $1$ 算出来的右边界不会触及或越过第二个 $1$，这样就能避免重复统计同一个子数组。

---

#### 提示 2-1

设当前枚举的巫师的能力值为 $v$，上面算出来的左右边界为 $[L,R]$，那么他对答案产生的贡献是 $v$ 乘上 $[L,R]$ 内所有包含 $v$ 的子数组的元素和的和。

#### 提示 2-2

如何计算子数组的元素和？

用前缀和来计算。

#### 提示 2-3

如何计算子数组的元素和**的和**？

不妨将子数组的右端点固定，子数组左端点的范围是多少？

对于多个不同的右端点，其对应的左端点的范围是否均相同？

#### 提示 2-4

设子数组右端点为 $r$，左端点为 $l$，当前枚举的元素下标为 $i$，那么有 $l\le i \le r$。

设 $\textit{strength}$ 数组的前缀和为 $s$，在范围 $[L,R]$ 内的所有子数组的元素和的和为

$$
\begin{aligned}
&\sum_{r=i+1}^{R+1}\sum_{l=L}^{i} s[r]-s[l] \\
=&\sum_{r=i+1}^{R+1}\left((i-L+1)\cdot s[r] - \sum_{l=L}^{i} s[l]\right) \\
=&(i-L+1)\cdot \sum_{r=i+1}^{R+1}s[r] -(R-i+1)\cdot \sum_{l=L}^{i} s[l]
\end{aligned}
$$

因此我们还需要计算出前缀和 $s$ 的前缀和 $\textit{ss}$，上式即为

$$
(i-L+1)\cdot (\textit{ss}[R+2]-\textit{ss}[i+1]) - (R-i+1)\cdot (\textit{ss}[i+1]-\textit{ss}[L]) 
$$

累加所有贡献即为答案。

#### 相似题目 

- [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/)
- [1856. 子数组最小乘积的最大值](https://leetcode.cn/problems/maximum-subarray-min-product/)
- [2104. 子数组范围和](https://leetcode.cn/problems/sum-of-subarray-ranges/)

#### 复杂度分析

- 时间复杂度：$O(n)$。
- 空间复杂度：$O(n)$。

```Python [sol1-Python3]
class Solution:
    def totalStrength(self, strength: List[int]) -> int:
        n = len(strength)
        left, st = [-1] * n, []  # left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
        for i, v in enumerate(strength):
            while st and strength[st[-1]] >= v: st.pop()
            if st: left[i] = st[-1]
            st.append(i)

        right, st = [n] * n, []  # right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
        for i in range(n - 1, -1, -1):
            while st and strength[st[-1]] > strength[i]: st.pop()
            if st: right[i] = st[-1]
            st.append(i)

        ss = list(accumulate(accumulate(strength, initial=0), initial=0))  # 前缀和的前缀和

        ans = 0
        for i, v in enumerate(strength):
            l, r = left[i] + 1, right[i] - 1  # [l, r]  左闭右闭
            tot = (i - l + 1) * (ss[r + 2] - ss[i + 1]) - (r - i + 1) * (ss[i + 1] - ss[l])
            ans += v * tot  # 累加贡献
        return ans % (10 ** 9 + 7)
```

```java [sol1-Java]
class Solution {
    public int totalStrength(int[] strength) {
        final int mod = (int) 1e9 + 7;

        var n = strength.length;
        var left = new int[n]; // left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
        var st = new ArrayDeque<Integer>();
        for (var i = 0; i < n; i++) {
            while (!st.isEmpty() && strength[st.peek()] >= strength[i]) st.pop();
            left[i] = st.isEmpty() ? -1 : st.peek();
            st.push(i);
        }

        var right = new int[n]; // right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
        st = new ArrayDeque<>();
        for (var i = n - 1; i >= 0; i--) {
            while (!st.isEmpty() && strength[st.peek()] > strength[i]) st.pop();
            right[i] = st.isEmpty() ? n : st.peek();
            st.push(i);
        }

        var s = new int[n + 1]; // 前缀和
        for (var i = 0; i < n; ++i) s[i + 1] = (s[i] + strength[i]) % mod;
        var ss = new int[n + 2]; // 前缀和的前缀和
        for (var i = 0; i <= n; ++i) ss[i + 1] = (ss[i] + s[i]) % mod;

        long ans = 0;
        for (var i = 0; i < n; ++i) {
            int l = left[i] + 1, r = right[i] - 1; // [l,r] 左闭右闭
            var tot = ((long) (i - l + 1) * (ss[r + 2] - ss[i + 1]) - (long) (r - i + 1) * (ss[i + 1] - ss[l])) % mod;
            ans = (ans + strength[i] * tot) % mod; // 累加贡献
        }
        return (int) (ans + mod) % mod; // 防止算出负数（因为上面算 tot 有个减法）
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    int totalStrength(vector<int> &strength) {
        const int mod = 1e9 + 7;

        int n = strength.size();
        // left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
        // right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
        vector<int> left(n, -1), right(n, n);
        stack<int> st;
        for (int i = 0; i < n; ++i) {
            while (!st.empty() && strength[st.top()] >= strength[i])
                st.pop();
            if (!st.empty()) left[i] = st.top();
            st.push(i);
        }

        while (!st.empty()) st.pop();
        for (int i = n - 1; i >= 0; --i) {
            while (!st.empty() && strength[st.top()] > strength[i])
                st.pop();
            if (!st.empty()) right[i] = st.top();
            st.push(i);
        }

        vector<int> s(n + 1), ss(n + 2);
        for (int i = 0; i < n; ++i) s[i + 1] = (s[i] + strength[i]) % mod; // 前缀和
        for (int i = 0; i <= n; ++i) ss[i + 1] = (ss[i] + s[i]) % mod; // 前缀和的前缀和

        int ans = 0;
        for (int i = 0; i < n; ++i) {
            long l = left[i] + 1, r = right[i] - 1; // [l,r] 左闭右闭
            long tot = ((i - l + 1) * (ss[r + 2] - ss[i + 1]) - (r - i + 1) * (ss[i + 1] - ss[l])) % mod;
            ans = (ans + strength[i] * tot) % mod; // 累加贡献
        }
        return (ans + mod) % mod; // 防止算出负数（因为上面算 tot 有个减法）
    }
};
```

```go [sol1-Go]
func totalStrength(strength []int) (ans int) {
	const mod int = 1e9 + 7

	n := len(strength)
	left := make([]int, n) // left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
	st := []int{}
	for i, v := range strength {
		for len(st) > 0 && strength[st[len(st)-1]] >= v {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			left[i] = st[len(st)-1]
		} else {
			left[i] = -1
		}
		st = append(st, i)
	}

	right := make([]int, n) // right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
	st = []int{}
	for i := n - 1; i >= 0; i-- {
		v := strength[i]
		for len(st) > 0 && strength[st[len(st)-1]] > v {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			right[i] = st[len(st)-1]
		} else {
			right[i] = n
		}
		st = append(st, i)
	}

	s := make([]int, n+1) // 前缀和
	for i, v := range strength {
		s[i+1] = (s[i] + v) % mod
	}
	ss := make([]int, n+2) // 前缀和的前缀和
	for i, v := range s {
		ss[i+1] = (ss[i] + v) % mod
	}
	for i, v := range strength {
		l, r := left[i]+1, right[i]-1 // [l,r] 左闭右闭
		tot := ((i-l+1)*(ss[r+2]-ss[i+1]) - (r-i+1)*(ss[i+1]-ss[l])) % mod
		ans = (ans + v*tot) % mod // 累加贡献
	}
	return (ans + mod) % mod // 防止算出负数（因为上面算 tot 有个减法）
}
```
