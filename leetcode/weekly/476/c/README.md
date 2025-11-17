题目让我们计算去掉 $0$ 后的不同整数个数，也就是**不含** $0$ 的整数个数。

设 $n$ 的十进制长度为 $m$。

对于长为 $1,2,\ldots,m-1$ 的不含 $0$ 整数，每一位 $1$ 到 $9$ 可以随便填，所以长为 $k$ 的不含 $0$ 整数有 $9^k$ 个。根据等比数列求和公式，累加得

$$
9^1 + 9^2 + \cdots + 9^{m-1} = \dfrac{9^m-9}{8}
$$

对于长为 $m$ 的整数 $x$，如果 $x$ 包含 $0$，那么去掉 $0$ 后长度小于 $m$，这样的整数上面已经统计过了。所以只需要统计长为 $m$ 且 $\le n$ 的不含 $0$ 整数个数。

从高到低遍历 $n$，设当前遍历的这一位为 $d$，分类讨论：

- 如果 $d\ge 1$：
  - 如果这一位填 $[1,d-1]$ 中的数，那么后面的位可以随便填 $1$ 到 $9$，方案数为 $(d-1)\cdot 9^k$，其中 $k$ 为后面的位数。
  - 如果这一位填 $d$，那么后面的位就不能随便填。我们继续遍历，计算这一位填 $d$ 的方案数（这是一个规模更小的子问题）。
- 如果 $d=0$，那么这一位只能填 $0$，不合法，跳出循环。

[本题视频讲解](https://www.bilibili.com/video/BV1ZuCQBJEjD/?t=5m35s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def countDistinct(self, n: int) -> int:
        s = str(n)
        m = len(s)

        # 计算长度小于 m 的不含 0 的整数个数
        # 9^1 + 9^2 + ... + 9^(m-1) = (9^m - 9) / 8
        pow9 = 9 ** m
        ans = (pow9 - 9) // 8

        # 计算长度恰好等于 m 的不含 0 的整数个数
        for i, d in enumerate(s):
            if d == '0':  # 只能填 0，不合法，跳出循环
                break
            # 这一位填 1 到 d-1，后面的数位可以随便填 1 到 9
            v = int(d) - 1
            if i == m - 1:
                v += 1  # 最低位可以等于 d
            pow9 //= 9
            ans += v * pow9
            # 然后，这一位填 d，继续遍历

        return ans
```

```java [sol-Java]
class Solution {
    public long countDistinct(long n) {
        char[] s = Long.toString(n).toCharArray();
        int m = s.length;

        // 计算长度小于 m 的不含 0 的整数个数
        // 9^1 + 9^2 + ... + 9^(m-1) = (9^m - 9) / 8
        long pow9 = (long) Math.pow(9, m);
        long ans = (pow9 - 9) / 8;

        // 计算长度恰好等于 m 的不含 0 的整数个数
        for (int i = 0; i < m; i++) {
            char d = s[i];
            if (d == '0') { // 只能填 0，不合法，跳出循环
                break;
            }
            // 这一位填 1 到 d-1，后面的数位可以随便填 1 到 9
            int v = d - '1';
            if (i == m - 1) {
                v++; // 最低位可以等于 d
            }
            pow9 /= 9;
            ans += v * pow9;
            // 然后，这一位填 d，继续遍历
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countDistinct(long long n) {
        string s = to_string(n);
        int m = s.size();

        // 计算长度小于 m 的不含 0 的整数个数
        // 9^1 + 9^2 + ... + 9^(m-1) = (9^m - 9) / 8
        long long pow9 = pow(9, m);
        long long ans = (pow9 - 9) / 8;

        // 计算长度恰好等于 m 的不含 0 的整数个数
        for (int i = 0; i < m; i++) {
            char d = s[i];
            if (d == '0') { // 只能填 0，不合法，跳出循环
                break;
            }
            // 这一位填 1 到 d-1，后面的数位可以随便填 1 到 9
            int v = d - '1';
            if (i == m - 1) {
                v++; // 最低位可以等于 d
            }
            pow9 /= 9;
            ans += v * pow9;
            // 然后，这一位填 d，继续遍历
        }

        return ans;
    }
};
```

```go [sol-Go]
func countDistinct(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	m := len(s)

	// 计算长度小于 m 的不含 0 的整数个数
	// 9^1 + 9^2 + ... + 9^(m-1) = (9^m - 9) / 8
	pow9 := int64(math.Pow(9, float64(m)))
	ans := (pow9 - 9) / 8

	// 计算长度恰好等于 m 的不含 0 的整数个数
	for i, d := range s {
		if d == '0' { // 只能填 0，不合法，跳出循环
			break
		}
		// 这一位填 1 到 d-1，后面的数位可以随便填 1 到 9
		v := d - '1'
		if i == m-1 {
			v++ // 最低位可以等于 d
		}
		pow9 /= 9
		ans += int64(v) * pow9
		// 然后，这一位填 d，继续遍历
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。$n$ 的十进制长度是 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(\log n)$。

## 优化

对于长为 $m$ 的不含 0 的整数个数，也可以从低到高遍历，遇到为 $0$ 的数位，就把个数置为 $0$。从而达到和上面的代码一样的效果。

```py [sol-Python3]
class Solution:
    def countDistinct(self, n: int) -> int:
        ans = 0
        pow9 = 1
        while n > 0:
            n, d = divmod(n, 10)
            if d == 0:
                ans = 0
            else:
                if pow9 > 1:
                    d -= 1
                ans += d * pow9
            pow9 *= 9
        return ans + (pow9 - 9) // 8
```

```java [sol-Java]
class Solution {
    public long countDistinct(long n) {
        long ans = 0;
        long pow9 = 1;
        for (; n > 0; n /= 10, pow9 *= 9) {
            long d = n % 10;
            if (d == 0) {
                ans = 0;
                continue;
            }
            if (pow9 > 1) {
                d--;
            }
            ans += d * pow9;
        }
        return ans + (pow9 - 9) / 8;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countDistinct(long long n) {
        long long ans = 0, pow9 = 1;
        for (; n > 0; n /= 10, pow9 *= 9) {
            long long d = n % 10;
            if (d == 0) {
                ans = 0;
                continue;
            }
            if (pow9 > 1) {
                d--;
            }
            ans += d * pow9;
        }
        return ans + (pow9 - 9) / 8;
    }
};
```

```go [sol-Go]
func countDistinct(n int64) (ans int64) {
	pow9 := int64(1)
	for ; n > 0; n /= 10 {
		d := n % 10
		if d == 0 {
			ans = 0
		} else {
			if pow9 > 1 {
				d--
			}
			ans += d * pow9
		}
		pow9 *= 9
	}
	return ans + (pow9-9)/8
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。$n$ 的十进制长度是 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
