**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

从小到大枚举子串右端点 $\textit{right}$，如果子串符合要求，则右移左端点 $\textit{left}$。

滑动窗口的内层循环结束时，右端点**固定**在 $\textit{right}$，左端点在 $0,1,2,\cdots,\textit{left}-1$ 的所有子串都是合法的，这一共有 $\textit{left}$ 个，加入答案。

具体请看 [周赛视频讲解](https://www.bilibili.com/video/BV1UcyYY4EnQ/) 第二题，欢迎点赞关注~

## 答疑

**问**：为什么只需判断 `cnt[c - 'a'] >= k`？

**答**：因为只有 `cnt[c - 'a']++` 导致了 `cnt[c - 'a']` 达到 $k$，所以其余字母的出现次数必然小于 $k$，无需判断。当然，这里的 `>=` 写成 `==` 也是可以的。

```py [sol-Python3]
class Solution:
    def numberOfSubstrings(self, s: str, k: int) -> int:
        ans = left = 0
        cnt = defaultdict(int)
        for c in s:
            cnt[c] += 1
            while cnt[c] >= k:
                cnt[s[left]] -= 1
                left += 1
            ans += left
        return ans
```

```java [sol-Java]
class Solution {
    int numberOfSubstrings(String S, int k) {
        char[] s = S.toCharArray();
        int ans = 0;
        int left = 0;
        int[] cnt = new int[26];
        for (char c : s) {
            cnt[c - 'a']++;
            while (cnt[c - 'a'] >= k) {
                cnt[s[left] - 'a']--;
                left++;
            }
            ans += left;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfSubstrings(string s, int k) {
        int ans = 0, left = 0, cnt[26]{};
        for (char c : s) {
            cnt[c - 'a']++;
            while (cnt[c - 'a'] >= k) {
                cnt[s[left] - 'a']--;
                left++;
            }
            ans += left;
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSubstrings(s string, k int) (ans int) {
	cnt := [26]int{}
	left := 0
	for _, c := range s {
		cnt[c-'a']++
		for cnt[c-'a'] >= k {
			cnt[s[left]-'a']--
			left++
		}
		ans += left
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以滑动窗口的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

更多相似题目，见下面滑动窗口题单中的「**§2.3.1 越长越合法**」。

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
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
