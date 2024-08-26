本题相当于把 [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/) 中的「出现次数不超过 $1$」改成「出现次数不超过 $2$」，做法请看视频 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/?t=12m)，把代码第 $8$ 行的 $1$ 改成 $2$ 即可。

```py [sol-Python3]
class Solution:
    def maximumLengthSubstring(self, s: str) -> int:
        ans = left = 0
        cnt = defaultdict(int)
        for i, c in enumerate(s):
            cnt[c] += 1
            while cnt[c] > 2:
                cnt[s[left]] -= 1
                left += 1
            ans = max(ans, i - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumLengthSubstring(String S) {
        char[] s = S.toCharArray();
        int ans = 0;
        int left = 0;
        int[] cnt = new int[26];
        for (int i = 0; i < s.length; i++) {
            int b = s[i] - 'a';
            cnt[b]++;
            while (cnt[b] > 2) {
                cnt[s[left++] - 'a']--;
            }
            ans = Math.max(ans, i - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumLengthSubstring(string s) {
        int ans = 0, left = 0, cnt[26]{};
        for (int i = 0; i < s.length(); i++) {
            int b = s[i] - 'a';
            cnt[b]++;
            while (cnt[b] > 2) {
                cnt[s[left++] - 'a']--;
            }
            ans = max(ans, i - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumLengthSubstring(s string) (ans int) {
	cnt := [26]int{}
	left := 0
	for i, b := range s {
		b -= 'a'
		cnt[b]++
		for cnt[b] > 2 {
			cnt[s[left]-'a']--
			left++
		}
		ans = max(ans, i-left+1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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
