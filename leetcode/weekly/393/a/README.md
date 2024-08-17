## 方法一：枚举

请看 [视频讲解](https://www.bilibili.com/video/BV1dJ4m1V7hK/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def findLatestTime(self, s: str) -> str:
        for h in range(11, -1, -1):
            for m in range(59, -1, -1):
                t = f"{h:02d}:{m:02d}"  # f-string
                if all(x == '?' or x == y for x, y in zip(s, t)):
                    return t
```

```java [sol-Java]
class Solution {
    public String findLatestTime(String S) {
        char[] s = S.toCharArray();
        for (int h = 11; ; h--) {
            if (s[0] != '?' && s[0] - '0' != h / 10 || s[1] != '?' && s[1] - '0' != h % 10) {
                continue;
            }
            for (int m = 59; m >= 0; m--) {
                if (s[3] != '?' && s[3] - '0' != m / 10 || s[4] != '?' && s[4] - '0' != m % 10) {
                    continue;
                }
                return String.format("%02d:%02d", h, m);
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string findLatestTime(string s) {
        for (int h = 11; ; h--) {
            if (s[0] != '?' && s[0] - '0' != h / 10 || s[1] != '?' && s[1] - '0' != h % 10) {
                continue;
            }
            for (int m = 59; m >= 0; m--) {
                if (s[3] != '?' && s[3] - '0' != m / 10 || s[4] != '?' && s[4] - '0' != m % 10) {
                    continue;
                }
                char ans[6];
                sprintf(ans, "%02d:%02d", h, m);
                return string(ans);
            }
        }
    }
};
```

```go [sol-Go]
func findLatestTime(s string) string {
	for h := 11; ; h-- {
		if s[0] != '?' && s[0]-'0' != byte(h/10) || s[1] != '?' && s[1]-'0' != byte(h%10) {
			continue
		}
		for m := 59; m >= 0; m-- {
			if s[3] != '?' && s[3]-'0' != byte(m/10) || s[4] != '?' && s[4]-'0' != byte(m%10) {
				continue
			}
			return fmt.Sprintf("%02d:%02d", h, m)
		}
	}
}
```

## 方法二：直接判断

```py [sol-Python3]
class Solution:
    def findLatestTime(self, s: str) -> str:
        s = list(s)
        if s[0] == '?':
            s[0] = '1' if s[1] == '?' or s[1] <= '1' else '0'
        if s[1] == '?':
            s[1] = '1' if s[0] == '1' else '9'
        if s[3] == '?':
            s[3] = '5'
        if s[4] == '?':
            s[4] = '9'
        return ''.join(s)
```

```java [sol-Java]
class Solution {
    public String findLatestTime(String s) {
        char[] t = s.toCharArray();
        if (t[0] == '?') {
            t[0] = t[1] == '?' || t[1] <= '1' ? '1' : '0';
        }
        if (t[1] == '?') {
            t[1] = t[0] == '1' ? '1' : '9';
        }
        if (t[3] == '?') {
            t[3] = '5';
        }
        if (t[4] == '?') {
            t[4] = '9';
        }
        return new String(t);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string findLatestTime(string s) {
        if (s[0] == '?') {
            s[0] = s[1] == '?' || s[1] <= '1' ? '1' : '0';
        }
        if (s[1] == '?') {
            s[1] = s[0] == '1' ? '1' : '9';
        }
        if (s[3] == '?') {
            s[3] = '5';
        }
        if (s[4] == '?') {
            s[4] = '9';
        }
        return s;
    }
};
```

```go [sol-Go]
func findLatestTime(s string) string {
	t := []byte(s)
	if t[0] == '?' {
		if t[1] == '?' || t[1] <= '1' {
			t[0] = '1'
		} else {
			t[0] = '0'
		}
	}
	if t[1] == '?' {
		if t[0] == '1' {
			t[1] = '1'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

- [1736. 替换隐藏数字得到的最晚时间](https://leetcode.cn/problems/latest-time-by-replacing-hidden-digits/)

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
