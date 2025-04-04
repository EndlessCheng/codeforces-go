定义 $f[i]$ 表示删除后缀 $s[i:]$ 所需的最大操作数。

根据题意，我们可以枚举删除字母的长度 $j$，如果 $s[i:i+j] = s[i+j:i+2j]$，那么可以删除，此时有转移 $f[i] = f[i+j] + 1$。如果不存在两个子串相等的情况，则 $f[i] = 1$。$f[i]$ 取所有情况的最大值。

倒着计算 $f[i]$，答案为 $f[0]$。

最后，我们需要快速判断两个子串是否相同。这可以用 $\mathcal{O}(n^2)$ 的 DP 预处理出来，具体见代码。

[视频讲解](https://www.bilibili.com/video/BV1kd4y1q7fC)

```py [sol-Python3]
class Solution:
    def deleteString(self, s: str) -> int:
        n = len(s)
        if len(set(s)) == 1:  # 特判全部相同的情况
            return n

        lcp = [[0] * (n + 1) for _ in range(n + 1)]  # lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
        for i in range(n - 1, -1, -1):
            for j in range(n - 1, i, -1):
                if s[i] == s[j]:
                    lcp[i][j] = lcp[i + 1][j + 1] + 1

        f = [0] * n
        for i in range(n - 1, -1, -1):
            for j in range(1, (n - i) // 2 + 1):
                if lcp[i][i + j] >= j:  # 说明 s[i:i+j] == s[i+j:i+2*j]
                    f[i] = max(f[i], f[i + j])  # 这里改成手动 if 更新会更快
            f[i] += 1
        return f[0]
```

```java [sol-Java]
class Solution {
    public int deleteString(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        if (allEqual(s)) { // 特判全部相同的情况
            return n;
        }

        int[][] lcp = new int[n + 1][n + 1]; // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
        for (int i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j > i; j--) {
                if (s[i] == s[j]) {
                    lcp[i][j] = lcp[i + 1][j + 1] + 1;
                }
            }
        }

        int[] f = new int[n];
        for (int i = n - 1; i >= 0; i--) {
            for (int j = 1; i + j * 2 <= n; j++) {
                if (lcp[i][i + j] >= j) { // 说明 s[i:i+j] == s[i+j:i+j*2]
                    f[i] = Math.max(f[i], f[i + j]);
                }
            }
            f[i]++;
        }
        return f[0];
    }

    private boolean allEqual(char[] s) {
        for (int i = 1; i < s.length; i++) {
            if (s[i] != s[0]) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int deleteString(string s) {
        int n = s.size();
        if (equal(s.begin() + 1, s.end(), s.begin())) { // 特判全部相同的情况
            return n;
        }

        vector lcp(n + 1, vector<int>(n + 1)); // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
        for (int i = n - 1; i >= 0; i--) {
            for (int j = n - 1; j > i; j--) {
                if (s[i] == s[j]) {
                    lcp[i][j] = lcp[i + 1][j + 1] + 1;
                }
            }
        }

        vector<int> f(n);
        for (int i = n - 1; i >= 0; i--) {
            for (int j = 1; i + j * 2 <= n; j++) {
                if (lcp[i][i + j] >= j) { // 说明 s[i:i+j] == s[i+j:i+j*2]
                    f[i] = max(f[i], f[i + j]);
                }
            }
            f[i]++;
        }
        return f[0];
    }
};
```

```go [sol-Go]
func deleteString(s string) int {
	n := len(s)
	if allEqual(s) { // 特判全部相同的情况
		return n
	}

	lcp := make([][]int, n+1) // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
	lcp[n] = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		lcp[i] = make([]int, n+1)
		for j := n - 1; j > i; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

	f := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := 1; i+j*2 <= n; j++ {
			if lcp[i][i+j] >= j { // 说明 s[i:i+j] == s[i+j:i+j*2]
				f[i] = max(f[i], f[i+j])
			}
		}
		f[i]++
	}
	return f[0]
}

func allEqual(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

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
