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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
