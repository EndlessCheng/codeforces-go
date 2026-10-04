## 方法一：前后缀分解

枚举 $k=0,1,2,\ldots,n-1$，总旋转次数是如下三部分之和：

- $s[0]$ 开始拨号，一直拨到 $s[k-1]$ 的最少总旋转次数。记作 $\textit{pre}[k-1]$。
- 指针从数字 $s[k-1]$ 旋转到数字 $s[n-1]$ 的最少旋转次数。
- 从 $s[n-1]$ 开始倒着拨号，一直拨到 $s[k]$ 的最少总旋转次数。记作 $\textit{suf}[k]$。

我们可以先倒着遍历 $s$，预处理所有 $\textit{suf}$。

然后正着遍历 $s$，计算 $\textit{pre}$ 的同时，用三部分之和更新答案的最小值。单独计算 $k=0$ 的情况。

对于数字 $x$ 和 $y$，设 $d = |x-y|$，那么：

- 指针不经过 $0\text{-}9$ 的旋转次数为 $d$。
- 指针经过 $0\text{-}9$ 的旋转次数为 $10-d$。

二者取最小值，得到指针从数字 $x$ 旋转到数字 $y$ 的最少旋转次数为

$$
\min(d, 10-d)
$$

[本题视频讲解](https://www.bilibili.com/video/BV1FiHj6uESs/?t=3m31s)，欢迎点赞关注~

```py [sol-Python3]
# 指针从数字 x 旋转到数字 y 的最少旋转次数
def dis(x: int, y: int) -> int:
    d = abs(x - y)
    return min(d, 10 - d)

class Solution:
    def minRotations(self, n: int, s: str) -> int:
        s = list(map(ord, s))  # 避免下面反复调用 ord

        suf = 0
        for x, y in pairwise(s):
            suf += dis(x, y)

        ans = dis(ord('0'), s[-1]) + suf  # k=0 的情况

        pre = dis(ord('0'), s[0])
        for x, y in pairwise(s):
            op = dis(x, y)
            suf -= op  # 撤销
            ans = min(ans, pre + dis(x, s[-1]) + suf)
            pre += op
        return ans
```

```java [sol-Java]
class Solution {
    public int minRotations(int n, String S) {
        char[] s = S.toCharArray();

        int suf = 0;
        for (int i = 1; i < n; i++) {
            suf += dis(s[i - 1], s[i]);
        }

        int ans = dis('0', s[n - 1]) + suf; // k=0 的情况

        int pre = dis('0', s[0]);
        for (int k = 1; k < n; k++) {
            int op = dis(s[k - 1], s[k]);
            suf -= op; // 撤销
            ans = Math.min(ans, pre + dis(s[k - 1], s[n - 1]) + suf);
            pre += op;
        }
        return ans;
    }

    // 指针从数字 x 旋转到数字 y 的最少旋转次数
    private int dis(char x, char y) {
        int d = Math.abs(x - y);
        return Math.min(d, 10 - d);
    }
}
```

```cpp [sol-C++]
class Solution {
    // 指针从数字 x 旋转到数字 y 的最少旋转次数
    int dis(char x, char y) {
        int d = abs(x - y);
        return min(d, 10 - d);
    }

public:
    int minRotations(int n, string s) {
        int suf = 0;
        for (int i = 1; i < n; i++) {
            suf += dis(s[i - 1], s[i]);
        }

        int ans = dis('0', s[n - 1]) + suf; // k=0 的情况

        int pre = dis('0', s[0]);
        for (int k = 1; k < n; k++) {
            int op = dis(s[k - 1], s[k]);
            suf -= op; // 撤销
            ans = min(ans, pre + dis(s[k - 1], s[n - 1]) + suf);
            pre += op;
        }
        return ans;
    }
};
```

```go [sol-Go]
// 指针从数字 x 旋转到数字 y 的最少旋转次数
func dis(x, y byte) int {
	d := abs(int(x) - int(y))
	return min(d, 10-d)
}

func minRotations(n int, s string) int {
	suf := 0
	for i := 1; i < n; i++ {
		suf += dis(s[i-1], s[i])
	}

	ans := dis('0', s[n-1]) + suf // k=0 的情况

	pre := dis('0', s[0])
	for k := 1; k < n; k++ {
		op := dis(s[k-1], s[k])
		suf -= op // 撤销
		ans = min(ans, pre+dis(s[k-1], s[n-1])+suf)
		pre += op
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：计算增量

设 $\textit{base}$ 为不操作时的最少总旋转次数。

反转后缀 $[k,n-1]$ 后，只有 $s[k-1]$ 到 $s[k]$ 发生了变化：

- 减去从 $s[k-1]$ 到 $s[k]$ 的旋转次数。
- 加上从 $s[k-1]$ 到 $s[n-1]$ 的旋转次数。

计算这个增量的最小值，加上 $\textit{base}$，即为答案。

```py [sol-Python3]
# 指针从数字 x 旋转到数字 y 的最少旋转次数
def dis(x: int, y: int) -> int:
    d = abs(x - y)
    return min(d, 10 - d)

class Solution:
    def minRotations(self, n: int, s: str) -> int:
        s = "0" + s
        base = mn = 0
        end = ord(s[-1])
        for x, y in pairwise(map(ord, s)):
            op = dis(x, y)
            base += op
            mn = min(mn, dis(x, end) - op)
        return base + mn
```

```py [sol-Python3 写法二]
# 指针从数字 x 旋转到数字 y 的最少旋转次数
def dis(x: int, y: int) -> int:
    d = abs(x - y)
    return min(d, 10 - d)

class Solution:
    def minRotations(self, n: int, s: str) -> int:
        base = mn = 0
        end = ord(s[-1])
        pre = ord('0')
        for ch in s:
            cur = ord(ch)
            op = dis(pre, cur)
            base += op
            mn = min(mn, dis(pre, end) - op)
            pre = cur
        return base + mn
```

```java [sol-Java]
class Solution {
    public int minRotations(int n, String s) {
        int base = 0;
        int mn = 0;
        char pre = '0';
        char end = s.charAt(n - 1);
        for (char cur : s.toCharArray()) {
            int op = dis(pre, cur);
            base += op;
            mn = Math.min(mn, dis(pre, end) - op);
            pre = cur;
        }
        return base + mn;
    }

    // 指针从数字 x 旋转到数字 y 的最少旋转次数
    private int dis(char x, char y) {
        int d = Math.abs(x - y);
        return Math.min(d, 10 - d);
    }
}
```

```cpp [sol-C++]
class Solution {
    // 指针从数字 x 旋转到数字 y 的最少旋转次数
    int dis(char x, char y) {
        int d = abs(x - y);
        return min(d, 10 - d);
    }

public:
    int minRotations(int n, string s) {
        int base = 0, mn = 0;
        char pre = '0';
        for (char cur : s) {
            int op = dis(pre, cur);
            base += op;
            mn = min(mn, dis(pre, s[n - 1]) - op);
            pre = cur;
        }
        return base + mn;
    }
};
```

```go [sol-Go]
// 指针从数字 x 旋转到数字 y 的最少旋转次数
func dis(x, y byte) int {
	d := abs(int(x) - int(y))
	return min(d, 10-d)
}

func minRotations(n int, s string) int {
	base, mn := 0, 0
	pre := byte('0')
	for _, cur := range s {
		op := dis(pre, byte(cur))
		base += op
		mn = min(mn, dis(pre, s[n-1])-op)
		pre = byte(cur)
	}
	return base + mn
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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
