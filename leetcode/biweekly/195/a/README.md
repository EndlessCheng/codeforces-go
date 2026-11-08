如果子串 $[L,R]$ 是有效子串，那么更短的 $[L+1,R],[L+2,R],\ldots,[R,R]$ 也都是有效子串。

如果子串 $[L-1,R]$ 不是有效子串，那么更长的 $[L-2,R],[L-3,R],\ldots,[0,R]$ 也都不是有效子串。

于是，枚举有效子串的右端点 $R = 0,1,2,\ldots,n-1$，同时维护合法左端点的最小值 $L$，那么对于一个**固定**的 $R$，有 $R-L+1$ 个有效子串，加到答案中。

[本题视频讲解](https://www.bilibili.com/video/BV152W76MEX3/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countValidSubstrings(self, s: str) -> int:
        left = -1  # left+1 是有效子串的最小左端点
        ans = 0
        for i in range(len(s)):
            # s[i - 2: i + 1] == "abc"
            if i >= 2 and s[i] == 'c' and s[i - 1] == 'b' and s[i - 2] == 'a':
                left = i - 2
            ans += i - left
        return ans
```

```java [sol-Java]
class Solution {
    public long countValidSubstrings(String S) {
        char[] s = S.toCharArray();
        int left = -1; // left+1 是有效子串的最小左端点
        long ans = 0;
        for (int i = 0; i < s.length; i++) {
            if (i >= 2 && s[i] == 'c' && s[i - 1] == 'b' && s[i - 2] == 'a') {
                left = i - 2;
            }
            ans += i - left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countValidSubstrings(string s) {
        int left = -1; // left+1 是有效子串的最小左端点
        long long ans = 0;
        for (int i = 0; i < s.size(); i++) {
            // s.substr(i - 2, 3) == "abc"
            if (i >= 2 && s[i] == 'c' && s[i - 1] == 'b' && s[i - 2] == 'a') {
                left = i - 2;
            }
            ans += i - left;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countValidSubstrings(s string) (ans int64) {
	left := -1 // left+1 是有效子串的最小左端点
	for i := range s {
		// s[i-2:i+1] == "abc" 
		if i >= 2 && s[i] == 'c' && s[i-1] == 'b' && s[i-2] == 'a' {
			left = i - 2
		}
		ans += int64(i - left)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nK)$，其中 $n$ 是 $s$ 的长度，$K=3$ 是字符串 $\texttt{abc}$ 的长度。**注**：对于 $K$ 更大的情况，用 [KMP 算法](https://www.zhihu.com/question/21923021/answer/37475572) 可以优化至 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
