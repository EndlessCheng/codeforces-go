遍历 $\textit{password}$，判断 $\textit{password}[i]$ 属于哪一类字符，累加相应得分。

注意同一个字符只能算一次得分。可以用一个 $\textit{vis}$ 哈希集合（或者布尔数组）标记遍历过的字符。

```py [sol-Python3]
class Solution:
    def passwordStrength(self, password: str) -> int:
        vis = set()
        ans = 0
        for ch in password:
            if ch in vis:
                continue
            vis.add(ch)
            if ch.islower():
                ans += 1
            elif ch.isupper():
                ans += 2
            elif ch.isdigit():
                ans += 3
            else:
                ans += 5
        return ans
```

```java [sol-Java]
class Solution {
    public int passwordStrength(String password) {
        boolean[] vis = new boolean[128];
        int ans = 0;
        for (char ch : password.toCharArray()) {
            if (vis[ch]) {
                continue;
            }
            vis[ch] = true;
            if (Character.isLowerCase(ch)) {
                ans++;
            } else if (Character.isUpperCase(ch)) {
                ans += 2;
            } else if (Character.isDigit(ch)) {
                ans += 3;
            } else {
                ans += 5;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int passwordStrength(string password) {
        bool vis[128]{};
        int ans = 0;
        for (char ch : password) {
            if (vis[ch]) {
                continue;
            }
            vis[ch] = true;
            if (islower(ch)) {
                ans++;
            } else if (isupper(ch)) {
                ans += 2;
            } else if (isdigit(ch)) {
                ans += 3;
            } else {
                ans += 5;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func passwordStrength(password string) (ans int) {
	vis := [128]bool{}
	for _, ch := range password {
		if vis[ch] {
			continue
		}
		vis[ch] = true
		if unicode.IsLower(ch) {
			ans++
		} else if unicode.IsUpper(ch) {
			ans += 2
		} else if unicode.IsDigit(ch) {
			ans += 3
		} else {
			ans += 5
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $\textit{password}$ 的长度，$|\Sigma|=128$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(\min(n,|\Sigma|))$ 或 $\mathcal{O}(|\Sigma|)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
