## 方法一：状态机

对于每种字母，定义如下四种状态：

- $0$：初始状态。如果遇到小写字母，转到状态 $1$，否则不合法，转到状态 $-1$。
- $1$：如果遇到小写字母则不变，遇到大写字母则转到状态 $2$。
- $2$：如果遇到大写字母则不变，遇到小写字母则转到状态 $-1$。
- $-1$：遇到任何字母都不再变化。

![394C.png](https://pic.leetcode.cn/1713671840-HgbYWt-394C.png)

答案为状态为 $2$ 的字母种数。

附 ASCII 的性质：

- 对于大写英文字母：其二进制从右往左第 $6$ 个比特值一定是 $0$。
- 对于小写英文字母：其二进制从右往左第 $6$ 个比特值一定是 $1$。
- 对于任何英文字母：其小写字母二进制低 $5$ 位，一定和其大写字母二进制低 $5$ 位相等。

请看 [视频讲解](https://www.bilibili.com/video/BV1gu4m1F7B8/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def numberOfSpecialChars(self, word: str) -> int:
        ans = 0
        state = [0] * 27
        for c in map(ord, word):
            x = c & 31  # 转成数字 1~26
            if c & 32:  # 小写字母
                if state[x] == 0:
                    state[x] = 1
                elif state[x] == 2:
                    state[x] = -1
                    ans -= 1
            else:  # 大写字母
                if state[x] == 0:
                    state[x] = -1
                elif state[x] == 1:
                    state[x] = 2
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int numberOfSpecialChars(String word) {
        int ans = 0;
        int[] state = new int[27];
        for (char c : word.toCharArray()) {
            int x = c & 31; // 转成数字 1~26
            if ((c & 32) > 0) { // 小写字母
                if (state[x] == 0) {
                    state[x] = 1;
                } else if (state[x] == 2) {
                    state[x] = -1;
                    ans--;
                }
            } else { // 大写字母
                if (state[x] == 0) {
                    state[x] = -1;
                } else if (state[x] == 1) {
                    state[x] = 2;
                    ans++;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfSpecialChars(string word) {
        int ans = 0;
        int state[27]{};
        for (char c : word) {
            int x = c & 31; // 转成数字 1~26
            if (c & 32) { // 小写字母
                if (state[x] == 0) {
                    state[x] = 1;
                } else if (state[x] == 2) {
                    state[x] = -1;
                    ans--;
                }
            } else { // 大写字母
                if (state[x] == 0) {
                    state[x] = -1;
                } else if (state[x] == 1) {
                    state[x] = 2;
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSpecialChars(word string) (ans int) {
	state := [27]int{}
	for _, c := range word {
		x := c & 31
		if c&32 > 0 {
			if state[x] == 0 {
				state[x] = 1
			} else if state[x] == 2 {
				state[x] = -1
				ans--
			}
		} else {
			if state[x] == 0 {
				state[x] = -1
			} else if state[x] == 1 {
				state[x] = 2
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 为 $\textit{word}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为英文字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 方法二：位运算

我们可以把小写字母加到集合 $\textit{lower}$ 中，大写字母（转成小写后）加到集合 $\textit{upper}$ 中。如果遍历到小写字母 $c$ 时，发现 $c$ 也在 $\textit{upper}$ 中，说明大写字母在小写字母前面，把 $c$ 加到集合 $\textit{invalid}$ 中。

最后计算 $\textit{lower}$ 和 $\textit{upper}$ 的交集，并从交集中去掉 $\textit{invalid}$ 中的元素，剩下的元素个数即为答案。

这可以用位运算实现，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def numberOfSpecialChars(self, word: str) -> int:
        lower = upper = invalid = 0
        for c in map(ord, word):
            bit = 1 << (c & 31)
            if c & 32:  # 小写字母
                lower |= bit
                if upper & bit:  # c 也在 upper 中
                    invalid |= bit  # 不合法
            else:  # 大写字母
                upper |= bit
        # 从交集 lower & upper 中去掉不合法的字母 invalid
        return (lower & upper & ~invalid).bit_count()
```

```java [sol-Java]
class Solution {
    public int numberOfSpecialChars(String word) {
        int lower = 0, upper = 0, invalid = 0;
        for (char c : word.toCharArray()) {
            int bit = 1 << (c & 31);
            if ((c & 32) > 0) { // 小写字母
                lower |= bit;
                if ((upper & bit) > 0) { // c 也在 upper 中
                    invalid |= bit; // 不合法
                }
            } else { // 大写字母
                upper |= bit;
            }
        }
        // 从交集 lower & upper 中去掉不合法的字母 invalid
        return Integer.bitCount(lower & upper & ~invalid);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfSpecialChars(string word) {
        int lower = 0, upper = 0, invalid = 0;
        for (char c : word) {
            int bit = 1 << (c & 31);
            if (c & 32) { // 小写字母
                lower |= bit;
                if (upper & bit) { // c 也在 upper 中
                    invalid |= bit; // 不合法
                }
            } else { // 大写字母
                upper |= bit;
            }
        }
        // 从交集 lower & upper 中去掉不合法的字母 invalid
        return __builtin_popcount(lower & upper & ~invalid);
    }
};
```

```go [sol-Go]
func numberOfSpecialChars(word string) int {
	var lower, upper, invalid uint
	for _, c := range word {
		bit := uint(1) << (c & 31)
		if c&32 > 0 { // 小写字母
			lower |= bit
			if upper&bit > 0 { // c 也在 upper 中
				invalid |= bit // 不合法
			}
		} else { // 大写字母
			upper |= bit
		}
	}
	// 从交集 lower & upper 中去掉不合法的字母 invalid
	return bits.OnesCount(lower & upper &^ invalid)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
