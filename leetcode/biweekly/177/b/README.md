这个「消除」的规则类似力扣常见的消除问题（见我数据结构题单的 §3.3 邻项消除）。

用栈保存未被消除的字符，如果

$$
新遍历到的字符下标 - 该字符在栈中的最大下标 > k
$$

那么把新遍历到的字符入栈，否则忽略新遍历到的字符。

由于未被消除的字符都在栈中，所以新遍历到的字符的下标就是栈的大小。

为了快速获取字符在栈中的最大下标，可以用一个数组（或者哈希表）保存每个字符在栈中的最新位置。

[本题视频讲解](https://www.bilibili.com/video/BV1VvABz9EGz/?t=4m52s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def mergeCharacters(self, s: str, k: int) -> str:
        st = []
        last = defaultdict(lambda: -inf)
        for ch in s:
            if len(st) - last[ch] > k:
                last[ch] = len(st)
                st.append(ch)
        return ''.join(st)
```

```java [sol-Java]
class Solution {
    public String mergeCharacters(String s, int k) {
        StringBuilder st = new StringBuilder();
        int[] last = new int[26];
        Arrays.fill(last, Integer.MIN_VALUE / 2);
        for (char ch : s.toCharArray()) {
            if (st.length() - last[ch - 'a'] > k) {
                last[ch - 'a'] = st.length();
                st.append(ch);
            }
        }
        return st.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string mergeCharacters(string s, int k) {
        string st;
        vector<int> last(26, INT_MIN / 2);
        for (char ch : s) {
            if (st.size() - last[ch - 'a'] > k) {
                last[ch - 'a'] = st.size();
                st += ch;
            }
        }
        return st;
    }
};
```

```go [sol-Go]
func mergeCharacters(s string, k int) string {
	st := []byte{}
	last := [26]int{}
	for i := range last {
		last[i] = math.MinInt / 2
	}
	for _, ch := range s {
		if len(st)-last[ch-'a'] > k {
			last[ch-'a'] = len(st)
			st = append(st, byte(ch))
		}
	}
	return string(st)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。注意创建大小为 $|\Sigma|$ 的数组需要 $\mathcal{O}(|\Sigma|)$ 的时间。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。返回值不计入。

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
