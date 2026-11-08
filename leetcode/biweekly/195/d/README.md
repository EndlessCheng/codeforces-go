设 $A = (a * x),\ B=(b * y)$。注：更常用的记号是 $a^x$。

暴力匹配的想法是，枚举 $B$ 在 $A$ 中的匹配起点 $\textit{start}=0,1,2,\ldots,|a|-1$，判断 $B$ 是否等于 $A$ 的从 $\textit{start}$ 开始的长为 $|B|$ 的子串。

但 $|B|$ 太大了，这样做会超时。

例如 $|a|=4,\ |b|=6,\ \textit{start}=0$，我们从 $i=0$ 开始，判断 $a[i\bmod 4] = b[i\bmod 6]$ 是否成立。如果前 $12$ 个字母都匹配，那么继续匹配相当于又从 $a[12\bmod 4]=a[0]$ 和 $b[12\bmod 6]=b[0]$ 开始匹配了，这两个字母（以及后续字母）我们已经知道它们是匹配的。所以**只要前 $12$ 个字母都匹配，那么后面所有字母都匹配**。

一般地，只需比较 $A[\textit{start}:]$ 和 $B$ 的前

$$
\text{lcm}(|a|,|b|)
$$

个字母是否匹配。

代码实现时：

1. 如果 $B$ 比 $A$ 还长，那么 $B$ 一定不是 $A$ 的子串。
2. 枚举 $\textit{start}$ 还需要保证 $|A[\textit{start}:]| \ge |B|$，即 $|A| - \textit{start}\ge |B|$，即 $\textit{start} \le |A| - |B|$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def repeatedStringMatch(self, a: str, b: str, x: int, y: int) -> bool:
        n, m = len(a), len(b)
        len_a, len_b = n * x, m * y
        if len_b > len_a:
            return False

        firsts = min(lcm(n, m), len_b)
        for start in range(min(n, len_a - len_b + 1)):
            # 判断 B 是否为 A[start:] 的前缀
            # 只需比较前 firsts 个字母
            if all(a[(start + i) % n] == b[i % m] for i in range(firsts)):
                return True
        return False
```

```java [sol-Java]
class Solution {
    public boolean repeatedStringMatch(String a, String b, int x, int y) {
        char[] s = a.toCharArray();
        char[] t = b.toCharArray();
        int n = s.length;
        int m = t.length;
        long lenA = (long) n * x;
        long lenB = (long) m * y;
        if (lenB > lenA) {
            return false;
        }

        int firsts = (int) Math.min(lcm(n, m), lenB);
        next:
        for (int start = 0; start < Math.min(n, lenA - lenB + 1); start++) {
            // 判断 B 是否为 A[start,lenA) 的前缀
            for (int i = 0; i < firsts; i++) { // 只需比较前 firsts 个字母
                if (s[(start + i) % n] != t[i % m]) {
                    continue next;
                }
            }
            return true;
        }
        return false;
    }

    private int lcm(int a, int b) {
        return a / gcd(a, b) * b;
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

```cpp [sol-C++]
class Solution {
public:
    bool repeatedStringMatch(string a, string b, int x, int y) {
        int n = a.size(), m = b.size();
        long long len_a = 1LL * n * x, len_b = 1LL * m * y;
        if (len_b > len_a) {
            return false;
        }

        int firsts = min(1LL * lcm(n, m), len_b);
        int mx = min(1LL * n, len_a - len_b + 1);
        for (int start = 0; start < mx; start++) {
            // 判断 B 是否为 A[start,len_a) 的前缀
            bool ok = true;
            for (int i = 0; i < firsts; i++) { // 只需比较前 firsts 个字母
                if (a[(start + i) % n] != b[i % m]) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return true;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func repeatedStringMatch(a, b string, x, y int) bool {
	n, m := len(a), len(b)
	lenA, lenB := n*x, m*y
	if lenB > lenA {
		return false
	}

	firsts := min(lcm(n, m), lenB)
next:
	for start := range min(n, lenA-lenB+1) {
		// 判断 B 是否为 A[start:] 的前缀
		for i := range firsts { // 只需比较前 firsts 个字母
			if a[(start+i)%n] != b[i%m] {
				continue next
			}
		}
		return true
	}
	return false
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2m)$，其中 $n$ 是 $a$ 的长度，$m$ 是 $\textit{b}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [686. 重复叠加字符串匹配](https://leetcode.cn/problems/repeated-string-match/)
- [466. 统计重复个数](https://leetcode.cn/problems/count-the-repetitions/)

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
