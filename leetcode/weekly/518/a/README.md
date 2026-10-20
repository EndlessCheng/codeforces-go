## 方法一：定长滑动窗口

任意循环移位后的字符串，都是 $s+s$ 的子串。

我们可以在 $s+s$ 上跑定长滑动窗口，维护窗口的得分。原理请看[【套路】教你解决定长滑窗！适用于所有定长滑窗题目！](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)

此外，$s+s$ 的第一个长为 $n$ 的子串和最后一个长为 $n$ 的子串是同一个，为避免重复统计，不考虑最后一个长为 $n$ 的子串。

[本题视频讲解](https://www.bilibili.com/video/BV18sbp6xEeE/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countRotations(self, s: str, k: int) -> int:
        n = len(s)
        ans = same = 0
        for i in range(n * 2 - 2):
            # 1. 入
            if s[i % n] == s[(i + 1) % n]:
                same += 1

            # 注意：一个窗口内有 n-1 对字母，对应 n-1 次循环，所以窗口大小是 n-1
            left = i - n + 2
            if left < 0:
                continue

            # 2. 更新答案
            if same == k:
                ans += 1

            # 3. 出
            if s[left] == s[(left + 1) % n]:
                same -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countRotations(String S, int k) {
        char[] s = S.toCharArray();
        int n = s.length;
        int same = 0;
        int ans = 0;
        for (int i = 0; i < n * 2 - 2; i++) {
            // 1. 入
            if (s[i % n] == s[(i + 1) % n]) {
                same++;
            }

            // 注意：一个窗口内有 n-1 对字母，对应 n-1 次循环，所以窗口大小是 n-1
            int left = i - n + 2;
            if (left < 0) {
                continue;
            }

            // 2. 更新答案
            if (same == k) {
                ans++;
            }

            // 3. 出
            if (s[left] == s[(left + 1) % n]) {
                same--;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countRotations(string s, int k) {
        int n = s.size();
        int same = 0;
        int ans = 0;
        for (int i = 0; i < n * 2 - 2; i++) {
            // 1. 入
            if (s[i % n] == s[(i + 1) % n]) {
                same++;
            }

            // 注意：一个窗口内有 n-1 对字母，对应 n-1 次循环，所以窗口大小是 n-1
            int left = i - n + 2;
            if (left < 0) {
                continue;
            }

            // 2. 更新答案
            if (same == k) {
                ans++;
            }

            // 3. 出
            if (s[left] == s[(left + 1) % n]) {
                same--;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countRotations(s string, k int) (ans int) {
	n := len(s)
	same := 0
	for i := range n*2 - 2 {
		// 1. 入
		if s[i%n] == s[(i+1)%n] {
			same++
		}

		// 注意：一个窗口内有 n-1 对字母，对应 n-1 次循环，所以窗口大小是 n-1
		left := i - n + 2
		if left < 0 {
			continue
		}

		// 2. 更新答案
		if same == k {
			ans++
		}

		// 3. 出
		if s[left] == s[(left+1)%n] {
			same--
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：分类讨论

把 $s$ 视作一个环形字符串，$s$ 的循环移位是环形字符串在某处断开后的结果。

设环形字符串 $s$ 有 $c$ 对相邻相同字母（注意 $s[n-1]$ 与 $s[0]$ 是相邻的）。

分类讨论：

- 如果 $c=k$，那么在相邻字母不同的位置断开，字符串的得分恰好为 $k$。这样的断开位置有 $n-c$ 个。
- 如果 $c=k+1$，那么在相邻字母相同的位置断开，字符串的得分恰好为 $k$。这样的断开位置有 $c$ 个。
- 其余情况，无法使字符串的得分恰好为 $k$。

```py [sol-Python3]
class Solution:
    def countRotations(self, s: str, k: int) -> int:
        n = len(s)
        c = 1 if s[0] == s[-1] else 0
        for x, y in pairwise(s):
            if x == y:
                c += 1

        if c == k:
            return n - c
        if c == k + 1:
            return c
        return 0
```

```java [sol-Java]
class Solution {
    public int countRotations(String S, int k) {
        char[] s = S.toCharArray();
        int n = s.length;
        int c = s[0] == s[n - 1] ? 1 : 0;
        for (int i = 1; i < n; i++) {
            if (s[i - 1] == s[i]) {
                c++;
            }
        }

        if (c == k) {
            return n - c;
        }
        if (c == k + 1) {
            return c;
        }
        return 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countRotations(string s, int k) {
        int n = s.size();
        int c = s[0] == s[n - 1];
        for (int i = 1; i < n; i++) {
            c += s[i - 1] == s[i];
        }

        if (c == k) {
            return n - c;
        }
        if (c == k + 1) {
            return c;
        }
        return 0;
    }
};
```

```go [sol-Go]
func countRotations(s string, k int) int {
	n := len(s)
	c := 0
	if s[0] == s[n-1] {
		c = 1
	}
	for i := 1; i < n; i++ {
		if s[i-1] == s[i] {
			c++
		}
	}

	if c == k {
		return n - c
	}
	if c == k+1 {
		return c
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面滑动窗口题单的「**一、定长滑动窗口**」。

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
