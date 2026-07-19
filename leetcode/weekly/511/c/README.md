> 注意我们排的是子序列，不需要连续。

对于字符串的子序列的排序，可以拆分成若干次如下原子操作：

- 选择子序列中一左一右的两个字符 $x$ 和 $y$，如果 $x > y$，则交换 $x$ 和 $y$。

由于 $[x,y]$ 也是子序列，所以原问题的操作，等价于：

- 选择 $s[i]=\texttt{1}$ 和 $s[j]=\texttt{0}$（$i<j$），交换 $s[i]$ 和 $s[j]$。

这意味着，**$s$ 中的 $\texttt{0}$ 只能向左移动**。

所以，为了尽量能把 $s$ 变成 $t = \textit{strs}[i]$，应当把 $t$ 中靠左的 $\texttt{?}$ 变成 $\texttt{0}$，靠右的 $\texttt{?}$ 变成 $\texttt{1}$，并保证 $s$ 和 $t$ 中的 $\texttt{0}$ 的个数相等。注：由于字符串长度一样，$\texttt{0}$ 的个数相等了，$\texttt{1}$ 的个数也就相等了。

然后，设 $s$ 中的 $\texttt{0}$ 的下标列表为 $P$，$t$ 中的 $\texttt{0}$ 的下标列表为 $Q$。由于 $s$ 中的 $\texttt{0}$ 只能向左移动，所以每个 $P[j]$ 都要 $\ge Q[j]$，才能保证我们能把 $s$ 变成 $t$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 写法一（优化前）

```py [sol-Python3]
class Solution:
    def transformStr(self, s: str, strs: list[str]) -> list[bool]:
        total0 = s.count('0')

        def check(t: str) -> bool:
            cnt0 = t.count('0')
            cnt_q = t.count('?')
            # t 中的 '0' 的个数在闭区间 [cnt0, cnt0+cnt_q] 中，total0 必须在这个范围内
            if not cnt0 <= total0 <= cnt0 + cnt_q:
                return False

            # 把前 total0-cnt0 个 '?' 改成 '0'
            t = list(t)
            for i, ch in enumerate(t):
                if cnt0 == total0:
                    break
                if ch == '?':
                    t[i] = '0'
                    cnt0 += 1

            # 判断能否把 s 变成 t
            i = j = 0
            for _ in range(total0):
                # 找下一个 s[i] = '0'
                while s[i] != '0':
                    i += 1

                # 找下一个 t[j] = '0'
                while t[j] != '0':
                    j += 1

                # s 中的 '0' 无法右移，所以无法把 s 变成 t
                if i < j:
                    return False

                i += 1
                j += 1

            return True

        return list(map(check, strs))
```

```java [sol-Java]
class Solution {
    public boolean[] transformStr(String S, String[] strs) {
        char[] s = S.toCharArray();
        int total0 = 0;
        for (char ch : s) {
            total0 += '1' - ch; // 统计 '0' 的个数
        }

        boolean[] ans = new boolean[strs.length];
        next:
        for (int idx = 0; idx < strs.length; idx++) {
            char[] t = strs[idx].toCharArray();
            int cnt0 = 0;
            int cntQ = 0;
            for (char ch : t) {
                if (ch == '0') {
                    cnt0++;
                } else if (ch == '?') {
                    cntQ++;
                }
            }

            // str 中的 '0' 的个数在闭区间 [cnt0, cnt0+cntQ] 中，total0 必须在这个范围内
            if (total0 < cnt0 || total0 > cnt0 + cntQ) {
                continue;
            }

            // 把前 total0-cnt0 个 '?' 改成 '0'
            for (int i = 0; i < t.length && cnt0 < total0; i++) {
                if (t[i] == '?') {
                    t[i] = '0';
                    cnt0++;
                }
            }

            // 判断能否把 s 变成 t
            int i = 0;
            int j = 0;
            while (cnt0-- > 0) {
                // 找下一个 s[i] = '0'
                while (s[i] != '0') {
                    i++;
                }

                // 找下一个 t[j] = '0'
                while (t[j] != '0') {
                    j++;
                }

                // s 中的 '0' 无法右移，所以无法把 s 变成 t
                if (i < j) {
                    continue next;
                }

                i++;
                j++;
            }

            ans[idx] = true;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> transformStr(string s, vector<string>& strs) {
        int total0 = ranges::count(s, '0');
        vector<bool> ans(strs.size());

        for (int idx = 0; idx < strs.size(); idx++) {
            auto& t = strs[idx];
            int cnt0 = ranges::count(t, '0');
            int cnt_q = ranges::count(t, '?');
            // str 中的 '0' 的个数在闭区间 [cnt0, cnt0+cnt_q] 中，total0 必须在这个范围内
            if (total0 < cnt0 || total0 > cnt0 + cnt_q) {
                continue;
            }

            // 把前 total0-cnt0 个 '?' 改成 '0'
            for (char& ch : t) {
                if (cnt0 == total0) {
                    break;
                }
                if (ch == '?') {
                    ch = '0';
                    cnt0++;
                }
            }

            // 判断能否把 s 变成 t
            bool ok = true;
            int i = 0, j = 0;
            while (cnt0--) {
                // 找下一个 s[i] = '0'
                while (s[i] != '0') {
                    i++;
                }

                // 找下一个 t[j] = '0'
                while (t[j] != '0') {
                    j++;
                }

                // s 中的 '0' 无法右移，所以无法把 s 变成 t
                if (i < j) {
                    ok = false;
                    break;
                }

                i++;
                j++;
            }

            ans[idx] = ok;
        }

        return ans;
    }
};
```

```go [sol-Go]
func transformStr(s string, strs []string) []bool {
	total0 := strings.Count(s, "0")
	ans := make([]bool, len(strs))

next:
	for idx, str := range strs {
		cnt0 := strings.Count(str, "0")
		cntQ := strings.Count(str, "?")
		// str 中的 '0' 的个数在闭区间 [cnt0, cnt0+cntQ] 中，total0 必须在这个范围内
		if total0 < cnt0 || total0 > cnt0+cntQ {
			continue
		}

		// 把前 total0-cnt0 个 '?' 改成 '0'
		t := []byte(str)
		for i, ch := range t {
			if cnt0 == total0 {
				break
			}
			if ch == '?' {
				t[i] = '0'
				cnt0++
			}
		}

		// 判断能否把 s 变成 t
		i, j := 0, 0
		for range total0 {
			// 找下一个 s[i] = '0'
			for s[i] != '0' {
				i++
			}

			// 找下一个 t[j] = '0'
			for t[j] != '0' {
				j++
			}

			// s 中的 '0' 无法右移，所以无法把 s 变成 t
			if i < j {
				continue next
			}

			i++
			j++
		}

		ans[idx] = true
	}

	return ans
}
```

## 写法二（优化）

把处理 $\texttt{?}$ 的逻辑合并到双指针中，这样无需修改字符串。

```py [sol-Python3]
class Solution:
    def transformStr(self, s: str, strs: list[str]) -> list[bool]:
        total0 = s.count('0')

        def check(t: str) -> bool:
            cnt0 = t.count('0')
            cnt_q = t.count('?')
            # t 中的 '0' 的个数在闭区间 [cnt0, cnt0+cnt_q] 中，total0 必须在这个范围内
            if not cnt0 <= total0 <= cnt0 + cnt_q:
                return False

            # 判断能否把 s 变成 t
            i = j = 0
            for _ in range(total0):
                # 找下一个 s[i] = '0'
                while s[i] != '0':
                    i += 1

                # 找下一个 t[j] = '0'
                while t[j] == '1' or t[j] == '?' and cnt0 == total0:
                    j += 1

                # s 中的 '0' 无法右移，所以无法把 s 变成 t
                if i < j:
                    return False

                if t[j] == '?':
                    cnt0 += 1

                i += 1
                j += 1

            return True

        return list(map(check, strs))
```

```java [sol-Java]
class Solution {
    public boolean[] transformStr(String S, String[] strs) {
        char[] s = S.toCharArray();
        int total0 = 0;
        for (char ch : s) {
            total0 += '1' - ch; // 统计 '0' 的个数
        }

        boolean[] ans = new boolean[strs.length];
        next:
        for (int idx = 0; idx < strs.length; idx++) {
            char[] t = strs[idx].toCharArray();
            int cnt0 = 0;
            int cntQ = 0;
            for (char ch : t) {
                if (ch == '0') {
                    cnt0++;
                } else if (ch == '?') {
                    cntQ++;
                }
            }

            // str 中的 '0' 的个数在闭区间 [cnt0, cnt0+cntQ] 中，total0 必须在这个范围内
            if (total0 < cnt0 || total0 > cnt0 + cntQ) {
                continue;
            }

            // 判断能否把 s 变成 t
            int i = 0;
            int j = 0;
            for (int k = 0; k < total0; k++) {
                // 找下一个 s[i] = '0'
                while (s[i] != '0') {
                    i++;
                }

                // 找下一个 t[j] = '0'
                while (t[j] == '1' || t[j] == '?' && cnt0 == total0) {
                    j++;
                }

                // s 中的 '0' 无法右移，所以无法把 s 变成 t
                if (i < j) {
                    continue next;
                }

                if (t[j] == '?') {
                    cnt0++;
                }

                i++;
                j++;
            }

            ans[idx] = true;
        }

        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> transformStr(string s, vector<string>& strs) {
        int total0 = ranges::count(s, '0');
        vector<bool> ans(strs.size());

        for (int idx = 0; idx < strs.size(); idx++) {
            auto& t = strs[idx];
            int cnt0 = ranges::count(t, '0');
            int cnt_q = ranges::count(t, '?');
            // str 中的 '0' 的个数在闭区间 [cnt0, cnt0+cnt_q] 中，total0 必须在这个范围内
            if (total0 < cnt0 || total0 > cnt0 + cnt_q) {
                continue;
            }

            // 判断能否把 s 变成 t
            bool ok = true;
            int i = 0, j = 0;
            for (int k = 0; k < total0; k++) {
                // 找下一个 s[i] = '0'
                while (s[i] != '0') {
                    i++;
                }

                // 找下一个 t[j] = '0'
                while (t[j] == '1' || t[j] == '?' && cnt0 == total0) {
                    j++;
                }

                // s 中的 '0' 无法右移，所以无法把 s 变成 t
                if (i < j) {
                    ok = false;
                    break;
                }

                if (t[j] == '?') {
                    cnt0++;
                }

                i++;
                j++;
            }

            ans[idx] = ok;
        }

        return ans;
    }
};
```

```go [sol-Go]
func transformStr(s string, strs []string) []bool {
	total0 := strings.Count(s, "0")
	ans := make([]bool, len(strs))

next:
	for idx, t := range strs {
		cnt0 := strings.Count(t, "0")
		cntQ := strings.Count(t, "?")
		// str 中的 '0' 的个数在闭区间 [cnt0, cnt0+cntQ] 中，total0 必须在这个范围内
		if total0 < cnt0 || total0 > cnt0+cntQ {
			continue
		}

		// 判断能否把 s 变成 t
		i, j := 0, 0
		for range total0 {
			// 找下一个 s[i] = '0'
			for s[i] != '0' {
				i++
			}

			// 找下一个 t[j] = '0'
			for t[j] == '1' || t[j] == '?' && cnt0 == total0 {
				j++
			}

			// s 中的 '0' 无法右移，所以无法把 s 变成 t
			if i < j {
				continue next
			}

			if t[j] == '?' {
				cnt0++
			}

			i++
			j++
		}

		ans[idx] = true
	}

	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $\textit{strs}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 专题训练

见下面双指针题单的「**四、双序列双指针**」。

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
