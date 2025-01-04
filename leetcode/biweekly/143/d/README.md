## 方法一：从左到右爆搜

首先，由于数位乘积中的质因子只有 $2,3,5,7$，如果 $t$ 包含其他质因子，那么直接返回 $-1$。如果 $t$ 只包含质因子 $2,3,5,7$，那么答案一定存在。

下文把 $\textit{num}$ 简记为 $s$。其长度为 $n$。

例如 $s=123$，并假设答案的长度也为 $3$。仿照 [数位 DP](https://www.bilibili.com/video/BV1rS4y1s721/?t=20m05s) 的思路，写一个爆搜：

- 最高位如果填 $1$，那么第二位不能填 $1$（不然小于 $s$ 了），至少要填 $2$。
- 最高位如果填的数大于 $1$，那么第二位，以及第三位，填的数字不受到 $s$ 的约束，可以填 $[1,9]$ 中的任意数字。

这启发我们也像数位 DP 那样，在递归的过程中，用一个参数 $\textit{isLimit}$ 表示「是否受到 $s$ 的约束」。

如何判断所填数位之积是 $t$ 的倍数呢？

比如 $t=10$：

- 如果最高位填 $2$，那么后面两位所填数字，只需满足乘积是 $t/2=5$ 的倍数。
- 如果最高位填 $5$，那么后面两位所填数字，只需满足乘积是 $t/5=2$ 的倍数。
- 如果最高位填 $6$，由于 $6$ 中有因子 $2$，那么后面两位所填数字，只需满足乘积是 $t/2=5$ 的倍数。

一般地，如果填的数字是 $d$，那么余下的数位，需要满足乘积是 $\dfrac{t}{\text{GCD}(t,d)}$ 的倍数。

综上所述，写一个带 $\textit{vis}$ 的爆搜，参数有：

- $i$：表示当前填到 $s$ 的第 $i$ 个数位了。
- $t$：表示 $[i,n-1]$ 所填数位，需要满足乘积是 $t$ 的倍数。
- $\textit{isLimit}$：表示是否受到 $s$ 的约束。如果为 $\texttt{false}$，那么当前数位可以填 $[1,9]$ 中的数；如果为 $\texttt{true}$，那么当前数位只能填 $[\max(s[i],1),9]$ 中的数。这里和 $1$ 取最大值是因为 $s[i]$ 可能为 $0$，但我们不能填 $0$。在受到 $s$ 约束的情况下，如果填的数字为 $s[i]$，那么后面仍然会受到 $s$ 的约束。

递归边界：当 $i=n$ 时，如果 $t=1$，说明当前填法是符合要求的，返回 $\texttt{true}$，否则返回 $\texttt{false}$。

递归入口：$\textit{dfs}(0,t,\texttt{true})$。从最高位开始填，一开始受到 $s$ 的约束。

如果下面的 DFS 返回 $\texttt{true}$，说明找到了答案，也直接返回 $\texttt{true}$。

### 细节

如果 $s$ 很短但 $t$ 很大，答案是会比 $s$ 还长的。

为了兼顾这种情况，我们可以往 $s$ 的前面添加

$$
\max(\textit{cnt}-n+1,1)
$$

个前导 $0$。其中 $\textit{cnt}$ 是 $t$ 的质因子个数。

注意至少要添加 $1$ 个前导零，因为可能有 $s=999$ 这种情况，即使 $t=2$，答案（$1112$）长度也比 $s$ 要长。

注意添加前导零会影响可以填入的数字，当 $\textit{isLimit}=\texttt{true}$ 且 $i < \textit{cnt}$ 时，我们可以填入 $0$。这和数位 DP 的「跳过不填数字」是一样的。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1cgmBYqEhu/?t=28m31s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def smallestNumber(self, s: str, t: int) -> str:
        tmp = t
        cnt = 0
        for p in 2, 3, 5, 7:
            while tmp % p == 0:
                tmp //= p
                cnt += 1
        if tmp > 1:  # t 包含其他质因子
            return "-1"

        # 补前导零（至少一个）
        cnt = max(cnt - len(s) + 1, 1)
        s = '0' * cnt + s

        n = len(s)
        ans = ['0'] * n

        @cache  # 仅仅作为 vis 标记使用
        def dfs(i: int, t: int, is_limit: bool) -> bool:
            if i == n:
                return t == 1

            if is_limit and i < cnt and dfs(i + 1, t, True):  # 填 0（跳过）
                return True

            low = int(s[i]) if is_limit else 0
            for d in range(max(low, 1), 10):
                if dfs(i + 1, t // gcd(t, d), is_limit and d == low):
                    ans[i] = str(d)
                    return True
            return False

        dfs(0, t, True)
        dfs.cache_clear()  # 防止爆内存
        return ''.join(ans).lstrip('0')  # 去掉前导零
```

```py [sol-Python3 写法二]
class Solution:
    def smallestNumber(self, s: str, t: int) -> str:
        tmp = t
        cnt = 0
        for p in 2, 3, 5, 7:
            while tmp % p == 0:
                tmp //= p
                cnt += 1
        if tmp > 1:  # t 包含其他质因子
            return "-1"

        # 补前导零（至少一个）
        cnt = max(cnt - len(s) + 1, 1)
        s = '0' * cnt + s

        n = len(s)
        ans = ['0'] * n
        vis = [set() for _ in range(n)]

        def dfs(i: int, t: int, is_limit: bool) -> bool:
            if i == n:
                return t == 1

            if not is_limit:
                if t in vis[i]:
                    return False
                vis[i].add(t)

            if is_limit and i < cnt and dfs(i + 1, t, True):  # 填 0（跳过）
                return True

            low = int(s[i]) if is_limit else 0
            for d in range(max(low, 1), 10):
                if dfs(i + 1, t // gcd(t, d), is_limit and d == low):
                    ans[i] = str(d)
                    return True
            return False

        dfs(0, t, True)
        return ''.join(ans).lstrip('0')  # 去掉前导零
```

```java [sol-Java]
class Solution {
    public String smallestNumber(String s, long t) {
        long tmp = t;
        int cnt = 0;
        for (int p : new int[]{2, 3, 5, 7}) {
            while (tmp % p == 0) {
                tmp /= p;
                cnt++;
            }
        }
        if (tmp > 1) { // t 包含其他质因子
            return "-1";
        }

        // 补前导零（至少一个）
        cnt = Math.max(cnt - s.length() + 1, 1);
        s = "0".repeat(cnt) + s;

        int n = s.length();
        char[] ans = new char[n];
        Arrays.fill(ans, '0');

        Set<Long>[] vis = new HashSet[n];
        Arrays.setAll(vis, i -> new HashSet<>());

        dfs(0, t, true, cnt, s.toCharArray(), ans, vis);
        for (int i = 0; ; i++) {
            if (ans[i] != '0') {
                return new String(ans, i, n - i); // 去掉前导零
            }
        }
    }

    private boolean dfs(int i, long t, boolean isLimit, int cnt, char[] s, char[] ans, Set<Long>[] vis) {
        if (i == s.length) {
            return t == 1;
        }
        if (!isLimit && !vis[i].add(t)) {
            return false;
        }

        if (isLimit && i < cnt && dfs(i + 1, t, true, cnt, s, ans, vis)) { // 填 0（跳过）
            return true;
        }

        int low = isLimit ? s[i] - '0' : 0;
        for (int d = Math.max(low, 1); d <= 9; d++) {
            if (dfs(i + 1, t / gcd(t, d), isLimit && d == low, cnt, s, ans, vis)) {
                ans[i] = (char) ('0' + d);
                return true;
            }
        }
        return false;
    }

    private long gcd(long a, long b) {
        while (a != 0) {
            long tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
// 超时了，看下面的方法二吧
class Solution {
public:
    string smallestNumber(string s, long long t) {
        long long tmp = t;
        int cnt = 0;
        for (int p : {2, 3, 5, 7}) {
            while (tmp % p == 0) {
                tmp /= p;
                cnt++;
            }
        }
        if (tmp > 1) {
            return "-1";
        }

        cnt = max(cnt - (int) s.length() + 1, 1);
        s = string(cnt, '0') + s;

        int n = s.length();
        string ans(n, '0');
        vector<unordered_set<long long>> vis(n);
        auto dfs = [&](auto&& dfs, int i, long long t, bool is_limit) -> bool {
            if (i == n) {
                return t == 1;
            }
            if (!is_limit && !vis[i].insert(t).second) {
                return false;
            }

            if (is_limit && i < cnt && dfs(dfs, i + 1, t, true)) { // 填 0（跳过）
                return true;
            }

            int low = is_limit ? s[i] - '0' : 0;
            for (int d = max(low, 1); d <= 9; d++) {
                if (dfs(dfs, i + 1, t / gcd(t, d), is_limit && d == low)) {
                    ans[i] = '0' + d;
                    return true;
                }
            }
            return false;
        };
        dfs(dfs, 0, t, true);

        auto it = ranges::find_if(ans, [](char c) { return c != '0'; });
        return string(it, ans.end()); // 去掉前导零
    }
};
```

```go [sol-Go]
func smallestNumber(s string, t int64) string {
	tmp, cnt := int(t), 0
	for _, p := range []int{2, 3, 5, 7} {
		for tmp%p == 0 {
			tmp /= p
			cnt++
		}
	}
	if tmp > 1 { // t 包含其他质因子
		return "-1"
	}

	// 补前导零（至少一个）
	cnt = max(cnt-len(s)+1, 1)
	s = strings.Repeat("0", cnt) + s

	n := len(s)
	ans := bytes.Repeat([]byte{'0'}, n)
	type pair struct{ i, t int }
	vis := map[pair]bool{}

	var dfs func(int, int, bool) bool
	dfs = func(i, t int, isLimit bool) bool {
		if i == n {
			return t == 1
		}
		if !isLimit {
			p := pair{i, t}
			if vis[p] {
				return false
			}
			vis[p] = true
		}

		if isLimit && i < cnt && dfs(i+1, t, true) { // 填 0（跳过）
			return true
		}

		low := 0
		if isLimit {
			low = int(s[i] - '0')
		}
		for d := max(low, 1); d <= 9; d++ {
			if dfs(i+1, t/gcd(t, d), isLimit && d == low) {
				ans[i] = '0' + byte(d)
				return true
			}
		}
		return false
	}
	dfs(0, int(t), true)

	i := bytes.LastIndexByte(ans, '0')
	return string(ans[i+1:])
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + D\log^5 t)$，其中 $n$ 是 $s$ 的长度，$D=9$。在 DFS 中，当 $i$ 小于 $\mathcal{O}(n-\log t)$ 时，只要某个 $s[i]$ 增加了 $1$，后面的数字可以随便填，一定可以找到答案。所以主要分析最后 $\mathcal{O}(\log t)$ 个 $i$。此时对应的 $t$，由于其质因子 $2,3,5,7$ 可以取到各自的 $\mathcal{O}(\log t)$ 个不同的幂次，根据乘法原理，总共有 $\mathcal{O}(\log^4 t)$ 个不同的 $t$，再算上枚举 $d$ 的 $D$ 次循环，总的时间复杂度为 $\mathcal{O}(n + D\log^5 t)$。
- 空间复杂度：$\mathcal{O}(n + \log^5 t)$。

## 方法二：从右到左枚举

首先，在不修改数字的情况下，把上面递归中的 $t$ 值，记在 $\textit{leftT}$ 中。

定义 $\textit{leftT}[0]=t$，递推式

$$
\textit{leftT}[i+1] = \dfrac{\textit{leftT}[i]}{\text{GCD}(\textit{leftT}[i], s[i])}
$$

这样方便我们在从右到左枚举时，知道当前的 $t$ 值。

如果 $\textit{leftT}[n]=1$，说明无需修改，直接返回 $s$。

如果 $s$ 不包含 $0$，那么从 $i=n-1$ 开始倒着枚举：

- 尝试增加 $s[n-1]$，并计算出新的 $t$ 值为 $\dfrac{\textit{leftT}[n-1]}{\text{GCD}(\textit{leftT}[n-1], s[n-1])}$。如果 $t=1$，返回 $s$。
- 如果 $s[n-1]$ 增加到 $10$，结束增加，继续枚举 $i=n-2$。
- 尝试增加 $s[n-2]$，并计算出新的 $t$ 值为 $\dfrac{\textit{leftT}[n-2]}{\text{GCD}(\textit{leftT}[n-2], s[n-2])}$。此时 $s[n-1]$ 是可以随便填的，我们可以从 $9$ 枚举到 $1$，如果发现 $s[n-1]$ 可以整除 $t$，则填入数字，并把 $t$ 变成 $\dfrac{t}{s[n-1]}$。如果 $t=1$，返回 $s$。

一般地，倒着枚举 $i$，尝试增加 $s[i]$，并计算出新的 $t$ 值为 $\dfrac{\textit{leftT}[i]}{\text{GCD}(\textit{leftT}[i], s[i])}$。

然后在 $j=n-1,n-2,n-3,\ldots,i+1$ 中填入数字，填入的方式是枚举 $s[j]=9,8,7,\ldots,1$，如果发现 $s[j]$ 可以整除 $t$，则更新 $t$ 为 $\dfrac{t}{s[j]}$，然后继续枚举下一个 $s[j]$。这种填法可以保证我们把大的数字填在了右边，从而使答案尽量小。

填完后如果 $t=1$，返回 $s$。

如果枚举到 $i=0$，仍然没有找到答案，说明答案一定比 $s$ 长，直接按照上面的填法生成答案。

### 细节

如果 $s$ 中有 $0$，那么这个 $0$ 必须要修改。设 $s$ 最左边的 $0$ 的下标为 $i_0$，上面的算法不能从 $i=n-1$ 开始枚举，而是从 $i=i_0$ 开始枚举。

```py [sol-Python3]
class Solution:
    def smallestNumber(self, s: str, t: int) -> str:
        tmp = t
        for i in range(9, 1, -1):
            while tmp % i == 0:
                tmp //= i
        if tmp > 1:  # t 包含大于 7 的质因子
            return "-1"

        n = len(s)
        left_t = [0] * (n + 1)
        left_t[0] = t
        i0 = n - 1
        for i, c in enumerate(s):
            if c == '0':
                i0 = i
                break
            left_t[i + 1] = left_t[i] // gcd(left_t[i], int(c))
        if left_t[n] == 1:  # s 的数位之积是 t 的倍数
            return s

        # 假设答案和 s 一样长
        s = list(map(int, s))
        for i in range(i0, -1, -1):
            for s[i] in range(s[i] + 1, 10):
                tt = left_t[i] // gcd(left_t[i], s[i])
                k = 9
                for j in range(n - 1, i, -1):
                    while tt % k:
                        k -= 1
                    tt //= k
                    s[j] = k
                if tt == 1:
                    return ''.join(map(str, s))

        # 答案一定比 s 长
        ans = []
        for i in range(9, 1, -1):
            while t % i == 0:
                ans.append(str(i))
                t //= i
        return ''.join(ans[::-1]).rjust(n + 1, '1')  # 前面补 1
```

```java [sol-Java]
class Solution {
    public String smallestNumber(String num, long t) {
        long tmp = t;
        for (int i = 9; i > 1; i--) {
            while (tmp % i == 0) {
                tmp /= i;
            }
        }
        if (tmp > 1) { // t 包含大于 7 的质因子
            return "-1";
        }

        char[] s = num.toCharArray();
        int n = s.length;
        long[] leftT = new long[n + 1];
        leftT[0] = t;
        int i0 = n - 1;
        for (int i = 0; i < n; i++) {
            if (s[i] == '0') {
                i0 = i;
                break;
            }
            leftT[i + 1] = leftT[i] / gcd(leftT[i], s[i] - '0');
        }
        if (leftT[n] == 1) { // num 的数位之积是 t 的倍数
            return num;
        }

        // 假设答案和 num 一样长
        for (int i = i0; i >= 0; i--) {
            while (++s[i] <= '9') {
                long tt = leftT[i] / gcd(leftT[i], s[i] - '0');
                int k = 9;
                for (int j = n - 1; j > i; j--) {
                    while (tt % k != 0) {
                        k--;
                    }
                    tt /= k;
                    s[j] = (char) ('0' + k);
                }
                if (tt == 1) {
                    return new String(s);
                }
            }
        }

        // 答案一定比 num 长
        StringBuilder ans = new StringBuilder();
        for (int i = 9; i > 1; i--) {
            while (t % i == 0) {
                ans.append((char) ('0' + i));
                t /= i;
            }
        }
        while (ans.length() <= n) {
            ans.append('1');
        }
        return ans.reverse().toString();
    }

    private long gcd(long a, long b) {
        while (a != 0) {
            long tmp = a;
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
    string smallestNumber(string s, long long t) {
        long long tmp = t;
        for (int i = 9; i > 1; i--) {
            while (tmp % i == 0) {
                tmp /= i;
            }
        }
        if (tmp > 1) { // t 包含大于 7 的质因子
            return "-1";
        }

        int n = s.length();
        vector<long long> left_t(n + 1);
        left_t[0] = t;
        int i0 = n - 1;
        for (int i = 0; i < n; i++) {
            if (s[i] == '0') {
                i0 = i;
                break;
            }
            left_t[i + 1] = left_t[i] / gcd(left_t[i], s[i] - '0');
        }
        if (left_t[n] == 1) { // s 的数位之积是 t 的倍数
            return s;
        }

        // 假设答案和 s 一样长
        for (int i = i0; i >= 0; i--) {
            while (++s[i] <= '9') {
                long long tt = left_t[i] / gcd(left_t[i], s[i] - '0');
                int k = 9;
                for (int j = n - 1; j > i; j--) {
                    while (tt % k) {
                        k--;
                    }
                    tt /= k;
                    s[j] = '0' + k;
                }
                if (tt == 1) {
                    return s;
                }
            }
        }

        // 答案一定比 s 长
        string ans;
        for (int i = 9; i > 1; i--) {
            while (t % i == 0) {
                ans += '0' + i;
                t /= i;
            }
        }
        ans += string(max(n + 1 - (int) ans.length(), 0), '1');
        ranges::reverse(ans);
        return ans;
    }
};
```

```go [sol-Go]
func smallestNumber(num string, t int64) string {
	tmp := int(t)
	for i := 9; i > 1; i-- {
		for tmp%i == 0 {
			tmp /= i
		}
	}
	if tmp > 1 { // t 包含大于 7 的质因子
		return "-1"
	}

	n := len(num)
	leftT := make([]int, n+1)
	leftT[0] = int(t)
	i0 := n - 1
	for i, c := range num {
		if c == '0' {
			i0 = i
			break
		}
		leftT[i+1] = leftT[i] / gcd(leftT[i], int(c-'0'))
	}
	if leftT[n] == 1 { // num 的数位之积是 t 的倍数
		return num
	}

	// 假设答案和 num 一样长
	s := []byte(num)
	for i := i0; i >= 0; i-- {
		for s[i]++; s[i] <= '9'; s[i]++ {
			tt := leftT[i] / gcd(leftT[i], int(s[i]-'0'))
			k := 9
			for j := n - 1; j > i; j-- {
				for tt%k > 0 {
					k--
				}
				tt /= k
				s[j] = '0' + byte(k)
			}
			if tt == 1 {
				return string(s)
			}
		}
	}

	// 答案一定比 num 长
	ans := []byte{}
	for i := int64(9); i > 1; i-- {
		for t%i == 0 {
			ans = append(ans, '0'+byte(i))
			t /= i
		}
	}
	for len(ans) <= n {
		ans = append(ans, '1')
	}
	slices.Reverse(ans)
	return string(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + D\log^2 t)$，其中 $n$ 是 $s$ 的长度，$D=9$。分析四重循环的循环次数，如果从 $i=n-1$ 开始循环，$i$ 至多减少 $\mathcal{O}(\log t)$ 次，就一定能在右边填入 $\mathcal{O}(\log t)$ 个数字，所以 $j$ 的循环次数是 $\mathcal{O}(\log t)$。而如果 $i$ 远小于 $n-1$，则一定能填入数字，$j$ 的循环次数是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [3260. 找出最大的 N 位 K 回文数](https://leetcode.cn/problems/find-the-largest-palindrome-divisible-by-k/) 2370

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
