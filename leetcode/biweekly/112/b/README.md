有一个更强的结论：改成 $j-i=2$，也是一样的。

既然可以交换相距为 $2$ 的字符，那么相距为 $4$ 的字符可以通过多次交换实现。例如 $x-y-z$ 变成 $y-x-z$ 变成 $y-z-x$ 变成 $z-y-z$。

依此类推，所有相距为偶数的字符都可以随意交换。

所以只需要看下标为偶数的字符个数是否都一样，下标为奇数的字符个数是否都一样。

请看 [视频讲解](https://www.bilibili.com/video/BV1um4y1M7Rv/)。

```py [sol-Python3]
class Solution:
    def checkStrings(self, s1: str, s2: str) -> bool:
        return Counter(s1[::2]) == Counter(s2[::2]) and Counter(s1[1::2]) == Counter(s2[1::2])
```

```java [sol-Java]
class Solution {
    public boolean checkStrings(String s1, String s2) {
        int[][] cnt1 = new int[2][26];
        int[][] cnt2 = new int[2][26];
        for (int i = 0; i < s1.length(); i++) {
            cnt1[i % 2][s1.charAt(i) - 'a']++;
            cnt2[i % 2][s2.charAt(i) - 'a']++;
        }
        return Arrays.deepEquals(cnt1, cnt2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool checkStrings(string s1, string s2) {
        int cnt1[2][26]{}, cnt2[2][26]{};
        for (int i = 0; i < s1.length(); i++) {
            cnt1[i % 2][s1[i] - 'a']++;
            cnt2[i % 2][s2[i] - 'a']++;
        }
        return memcmp(cnt1, cnt2, sizeof(cnt1)) == 0;
    }
};
```

```go [sol-Go]
func checkStrings(s1, s2 string) bool {
	var cnt1, cnt2 [2][26]int
	for i, c := range s1 {
		cnt1[i%2][c-'a']++
		cnt2[i%2][s2[i]-'a']++
	}
	return cnt1 == cnt2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 为 $s_1$ 的长度。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$，其中 $|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。

## 思考题

改成 $j-i=3$ 要怎么做？

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
