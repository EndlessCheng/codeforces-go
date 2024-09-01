考察编程基本功，按照题目要求写代码即可。

```py [sol-Python3]
class Solution:
    def stringHash(self, s: str, k: int) -> str:
        ans = []
        for i in range(0, len(s), k):
            total = sum(ord(c) for c in s[i: i + k]) - ord('a') * k
            ans.append(ascii_lowercase[total % 26])
        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String stringHash(String s, int k) {
        int n = s.length();
        char[] ans = new char[n / k];
        for (int i = 0; i < n; i += k) {
            int sum = 0;
            for (int j = i; j < i + k; j++) {
                sum += s.charAt(j) - 'a';
            }
            ans[i / k] = (char) ('a' + sum % 26);
        }
        return new String(ans);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string stringHash(string s, int k) {
        int n = s.length();
        string ans(n / k, 'a');
        for (int i = 0; i < n; i += k) {
            int sum = 0;
            for (int j = i; j < i + k; j++) {
                sum += s[j] - 'a';
            }
            ans[i / k] += sum % 26;
        }
        return ans;
    }
};
```

```go [sol-Go]
func stringHash(s string, k int) string {
	n := len(s)
	ans := make([]byte, n/k)
	for i := 0; i < n; i += k {
		sum := 0
		for _, b := range s[i : i+k] {
			sum += int(b - 'a')
		}
		ans[i/k] = 'a' + byte(sum%26)
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 时间复杂度：$\mathcal{O}(1)$。返回值不计入。Python 忽略切片空间。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
