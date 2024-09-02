考虑枚举所有长为 $n$ 的回文数。

首先，知道了回文数的左半边，就知道了回文数的右半边，所以可以枚举回文数的左半边。

设 $m = \left\lfloor\dfrac{n-1}{2}\right\rfloor$，设 $\textit{base} = 10^m$。

在 $[\textit{base}, 10\cdot\textit{base})$ 范围内枚举所有长为 $n$ 的回文数的左半边。

如果回文数 $x$ 能被 $k$ 整除，那么接下来需要解决的问题有两个：

1. 计算 $x$ 有多少个不同的排列。
2. 不能重复统计。

为了保证不重复统计，可以像 [49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/solutions/2718519/ha-xi-biao-fen-zu-jian-ji-xie-fa-pythonj-1ukv/) 那样，把 $x$ 的十进制字符串 $s$ **排序**，如果之前遇到过同样的字符串 $t$，那么 $s$ 生成的所有排列，$t$ 也能生成。用哈希表记录排序后的字符串，如果 $s$ 排序后在哈希表中，那么就跳过。

下面是组合数学时间。

本质上计算的是「**有重复元素的排列个数**」。

统计 $s$ 中的每个数字的出现次数 $\textit{cnt}$。

先填最高位。由于不能有前导零，最高位可以填的数有 $n-\textit{cnt}_0$ 个。其余 $n-1$ 个数随便排，有 $(n-1)!$ 种方案。

当然，这里面有重复的，例如 $x=34543$，其中两个 $3$ 和两个 $4$ 的排列就是重复的，由于这两个 $3$ **无法区分**，两个 $4$ **无法区分**，方案数要除以 $2!2!$。为什么是除法，可以看本题视频讲解。

综上，排列个数为

$$
\dfrac{(n-\textit{cnt}_0)\cdot (n-1)!}{\prod\limits_{i=0}^{9}\textit{cnt}_i!}
$$

加入答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ajHYeoEG5/) 第三题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countGoodIntegers(self, n: int, k: int) -> int:
        fac = [factorial(i) for i in range(n + 1)]
        ans = 0
        vis = set()
        base = 10 ** ((n - 1) // 2)
        for i in range(base, base * 10):  # 枚举回文数左半边
            s = str(i)
            s += s[::-1][n % 2:]
            if int(s) % k:  # 回文数不能被 k 整除
                continue

            sorted_s = ''.join(sorted(s))
            if sorted_s in vis:  # 不能重复统计
                continue
            vis.add(sorted_s)

            cnt = Counter(sorted_s)
            res = (n - cnt['0']) * fac[n - 1]
            for c in cnt.values():
                res //= fac[c]
            ans += res
        return ans
```

```java [sol-Java]
class Solution {
    public long countGoodIntegers(int n, int k) {
        int[] factorial = new int[n + 1];
        factorial[0] = 1;
        for (int i = 1; i <= n; i++) {
            factorial[i] = factorial[i - 1] * i;
        }

        long ans = 0;
        Set<String> vis = new HashSet<>();
        int base = (int) Math.pow(10, (n - 1) / 2);
        for (int i = base; i < base * 10; i++) { // 枚举回文数左半边
            String s = Integer.toString(i);
            s += new StringBuilder(s).reverse().substring(n % 2);
            if (Long.parseLong(s) % k > 0) { // 回文数不能被 k 整除
                continue;
            }

            char[] sortedS = s.toCharArray();
            Arrays.sort(sortedS);
            if (!vis.add(new String(sortedS))) { // 不能重复统计
                continue;
            }

            int[] cnt = new int[10];
            for (char c : sortedS) {
                cnt[c - '0']++;
            }
            int res = (n - cnt[0]) * factorial[n - 1];
            for (int c : cnt) {
                res /= factorial[c];
            }
            ans += res;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countGoodIntegers(int n, int k) {
        vector<int> factorial(n + 1);
        factorial[0] = 1;
        for (int i = 1; i <= n; i++) {
            factorial[i] = factorial[i - 1] * i;
        }

        long long ans = 0;
        unordered_set<string> vis;
        int base = pow(10, (n - 1) / 2);
        for (int i = base; i < base * 10; i++) { // 枚举回文数左半边
            string s = to_string(i);
            s += string(s.rbegin() + (n % 2), s.rend());
            if (stoll(s) % k) { // 回文数不能被 k 整除
                continue;
            }

            ranges::sort(s);
            if (!vis.insert(s).second) { // 不能重复统计
                continue;
            }

            int cnt[10]{};
            for (char c : s) {
                cnt[c - '0']++;
            }
            int res = (n - cnt[0]) * factorial[n - 1];
            for (int c : cnt) {
                res /= factorial[c];
            }
            ans += res;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countGoodIntegers(n, k int) (ans int64) {
	factorial := make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	vis := map[string]bool{}
	base := int(math.Pow10((n - 1) / 2))
	for i := base; i < base*10; i++ { // 枚举回文数左半边
		x := i
		t := i
		if n%2 > 0 {
			t /= 10
		}
		for ; t > 0; t /= 10 {
			x = x*10 + t%10
		}
		if x%k > 0 { // 回文数不能被 k 整除
			continue
		}

		bs := []byte(strconv.Itoa(x))
		slices.Sort(bs)
		s := string(bs)
		if vis[s] { // 不能重复统计
			continue
		}
		vis[s] = true

		cnt := [10]int{}
		for _, c := range bs {
			cnt[c-'0']++
		}
		res := (n - cnt[0]) * factorial[n-1]
		for _, c := range cnt {
			res /= factorial[c]
		}
		ans += int64(res)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(10^m\cdot n\log n)$，其中 $m = \left\lfloor\dfrac{n-1}{2}\right\rfloor$。
- 空间复杂度：$\mathcal{O}(10^m\cdot n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
