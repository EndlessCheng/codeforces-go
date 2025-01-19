本题 [视频讲解](https://www.bilibili.com/video/BV1RY4y157nW) 已出炉，额外介绍了单调栈的原理，欢迎点赞三连~

---

#### 提示 1-1

枚举每位巫师，假设他是最弱的巫师，那么他能在哪些子数组中？

#### 提示 1-2

左右边界最远能到哪？具体地，这些子数组的左边界的最小值是多少，右边界的最大值是多少？

#### 提示 1-3

用**单调栈**来计算左右边界，原理请看 [单调栈【基础算法精讲 26】](https://www.bilibili.com/video/BV1VN411J7S7/)。

#### 提示 1-4

注意本题是可能有**重复元素**的，这会对最终答案的计算产生什么影响？

#### 提示 1-5

设左右边界为 $[L,R]$。

为了避免重复计算，我们可以考虑左侧**严格小于**当前元素的最近元素位置 $L-1$，以及右侧**小于等于**当前元素的最近元素位置 $R+1$。

以示例 1 中的数组 $[1,3,1,2]$ 为例，如果左右两侧都是找严格小于，那么第一个 $1$ 和第二个 $1$ 算出来的边界范围都是一样的（都是整个数组），这就重复统计了，为了避免这种情况，可以把某一侧改为小于等于，比如把右侧改成小于等于，那么第一个 $1$ 算出来的右边界不会触及或越过第二个 $1$，这样就能避免重复统计同一个子数组。

---

#### 提示 2-1

设当前枚举的巫师的能力值为 $v$，那么他对答案产生的贡献是 $v$ 乘上在左右边界 $[L,R]$ 内的所有包含 $v$ 的子数组的元素和的和。

#### 提示 2-2

如何计算子数组的元素和？

用前缀和来计算。

#### 提示 2-3

如何计算子数组的元素和**的和**？

不妨将子数组的右端点固定，子数组左端点的范围是多少？

对于多个不同的右端点，其对应的左端点的范围是否均相同？

#### 提示 2-4

设子数组左端点为 $l$，右端点为 $r$，当前枚举的元素下标为 $i$，那么有 $L\le l\le i \le r\le R$。

设 $\textit{strength}$ 数组的前缀和为 $s$，其中 $s[i]=\sum\limits_{j=0}^{i-1} \textit{strength}[j]$，因此子数组 $[l,r]$ 的元素和可以表示为

$$
s[r+1]-s[l]
$$

在范围 $[L,R]$ 内的所有子数组的元素和的和可以表示为

$$
\begin{aligned}
&\sum_{r=i+1}^{R+1}\sum_{l=L}^{i} (s[r]-s[l]) \\
=&\left(\sum_{r=i+1}^{R+1}\sum_{l=L}^{i} s[r]\right)-\left(\sum_{r=i+1}^{R+1}\sum_{l=L}^{i} s[l]\right) \\
=&(i-L+1)\cdot \sum_{r=i+1}^{R+1}s[r] -(R-i+1)\cdot \sum_{l=L}^{i} s[l]
\end{aligned}
$$

因此我们还需要计算出前缀和 $s$ 的前缀和 $\textit{ss}$，其中 $\textit{ss}[i]=\sum\limits_{j=0}^{i-1}s[j]$，上式即为

$$
(i-L+1)\cdot (\textit{ss}[R+2]-\textit{ss}[i+1]) - (R-i+1)\cdot (\textit{ss}[i+1]-\textit{ss}[L])
$$

再乘上 $v$ 即为当前巫师的贡献，累加所有贡献即为答案。

关于取模的技巧，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
class Solution:
    def totalStrength(self, strength: List[int]) -> int:
        n = len(strength)
        # left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
        left = [-1] * n
        # right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
        right = [n] * n
        st = []
        for i, v in enumerate(strength):
            while st and strength[st[-1]] >= v:
                right[st.pop()] = i
            if st:
                left[i] = st[-1]
            st.append(i)

        ss = list(accumulate(accumulate(strength, initial=0), initial=0))  # 前缀和的前缀和

        ans = 0
        for i, v in enumerate(strength):
            l, r = left[i] + 1, right[i] - 1  # [l, r]  左闭右闭
            tot = (i - l + 1) * (ss[r + 2] - ss[i + 1]) - (r - i + 1) * (ss[i + 1] - ss[l])
            ans += v * tot  # 累加贡献
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int totalStrength(int[] strength) {
        final int MOD = 1_000_000_007;

        int n = strength.length;
        int[] left = new int[n];  // left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
        int[] right = new int[n]; // right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
        Arrays.fill(right, n);
        Deque<Integer> st = new ArrayDeque<>();
        st.push(-1); // 哨兵，方便计算 left[i]
        for (int i = 0; i < n; i++) {
            while (st.size() > 1 && strength[st.peek()] >= strength[i]) {
                right[st.pop()] = i;
            }
            left[i] = st.peek();
            st.push(i);
        }

        long s = 0; // 前缀和
        int[] ss = new int[n + 2]; // 前缀和的前缀和
        for (int i = 1; i <= n; i++) {
            s += strength[i - 1];
            ss[i + 1] = (int) ((ss[i] + s) % MOD); // 注意取模后，下面计算两个 ss 相减，结果可能为负
        }

        long ans = 0;
        for (int i = 0; i < n; i++) {
            int l = left[i] + 1;
            int r = right[i] - 1; // [l,r] 左闭右闭
            long tot = ((long) (i - l + 1) * (ss[r + 2] - ss[i + 1]) - (long) (r - i + 1) * (ss[i + 1] - ss[l])) % MOD;
            ans = (ans + strength[i] * tot) % MOD; // 累加贡献
        }
        return (int) (ans + MOD) % MOD; // 防止算出负数
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int totalStrength(vector<int>& strength) {
        const int MOD = 1'000'000'007;

        int n = strength.size();
        vector<int> left(n, -1); // left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
        vector<int> right(n, n); // right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
        stack<int> st;
        st.push(-1); // 哨兵，方便计算 left[i]
        for (int i = 0; i < n; i++) {
            while (st.size() > 1 && strength[st.top()] >= strength[i]) {
                right[st.top()] = i;
                st.pop();
            }
            left[i] = st.top();
            st.push(i);
        }

        long long s = 0; // 前缀和
        vector<int> ss(n + 2); // 前缀和的前缀和
        for (int i = 1; i <= n; i++) {
            s += strength[i - 1];
            ss[i + 1] = (ss[i] + s) % MOD; // 注意取模后，下面计算两个 ss 相减，结果可能为负
        }

        long long ans = 0;
        for (int i = 0; i < n; i++) {
            long long l = left[i] + 1, r = right[i] - 1; // [l,r] 左闭右闭
            long long tot = ((i - l + 1) * (ss[r + 2] - ss[i + 1]) - (r - i + 1) * (ss[i + 1] - ss[l])) % MOD;
            ans = (ans + strength[i] * tot) % MOD; // 累加贡献
        }
        return (ans + MOD) % MOD; // 防止算出负数
    }
};
```

```go [sol-Go]
func totalStrength(strength []int) (ans int) {
	const mod = 1_000_000_007

	n := len(strength)
	left := make([]int, n)  // left[i] 为左侧严格小于 strength[i] 的最近元素位置（不存在时为 -1）
	right := make([]int, n) // right[i] 为右侧小于等于 strength[i] 的最近元素位置（不存在时为 n）
	for i := range right {
		right[i] = n
	}
	st := []int{-1} // 哨兵，方便计算 left[i]
	for i, v := range strength {
		for len(st) > 1 && strength[st[len(st)-1]] >= v {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	s := 0 // 前缀和
	ss := make([]int, n+2) // 前缀和的前缀和
	for i, v := range strength {
		s += v
		ss[i+2] = (ss[i+1] + s) % mod // 注意取模后，下面计算两个 ss 相减，结果可能为负
	}
	for i, v := range strength {
		l, r := left[i]+1, right[i]-1 // [l,r] 左闭右闭
		tot := ((i-l+1)*(ss[r+2]-ss[i+1]) - (r-i+1)*(ss[i+1]-ss[l])) % mod
		ans = (ans + v*tot) % mod // 累加贡献
	}
	return (ans + mod) % mod // 防止算出负数
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. 【本题相关】[单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
