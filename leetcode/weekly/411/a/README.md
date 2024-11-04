计算以 $0$ 为右端点的合法子串个数，以 $1$ 为右端点的合法子串个数，……，以 $n-1$ 为右端点的合法子串个数。

我们需要知道以 $i$ 为右端点的合法子串，其左端点最小是多少。

由于随着 $i$ 的变大，窗口内的字符数量变多，越不能满足题目要求，所以最小左端点会随着 $i$ 的增大而增大，有**单调性**，因此可以用 [滑动窗口](https://www.bilibili.com/video/BV1hd4y1r7Gq/) 计算。

设以 $i$ 为右端点的合法子串，其左端点**最小**是 $\textit{left}_i$。

那么以 $i$ 为右端点的合法子串，其左端点可以是 $\textit{left}_i,\textit{left}_i+1,\ldots,i$，一共

$$
i-\textit{left}_i+1
$$

个，累加到答案中。

**细节**：字符 $0$ 的 ASCII 值是偶数，字符 $1$ 的 ASCII 值是奇数，所以可以用 ASCII 值 $c\bmod 2$ 得到对应的数字。这也等价于和 $1$ 计算 AND。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1hH4y1c7T5/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countKConstraintSubstrings(self, s: str, k: int) -> int:
        ans = left = 0
        cnt = [0, 0]
        for i, c in enumerate(s):
            cnt[ord(c) & 1] += 1
            while cnt[0] > k and cnt[1] > k:
                cnt[ord(s[left]) & 1] -= 1
                left += 1
            ans += i - left + 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countKConstraintSubstrings(String S, int k) {
        char[] s = S.toCharArray();
        int ans = 0;
        int left = 0;
        int[] cnt = new int[2];
        for (int i = 0; i < s.length; i++) {
            cnt[s[i] & 1]++;
            while (cnt[0] > k && cnt[1] > k) {
                cnt[s[left] & 1]--;
                left++;
            }
            ans += i - left + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countKConstraintSubstrings(string s, int k) {
        int ans = 0, left = 0, cnt[2]{};
        for (int i = 0; i < s.length(); i++) {
            cnt[s[i] & 1]++;
            while (cnt[0] > k && cnt[1] > k) {
                cnt[s[left] & 1]--;
                left++;
            }
            ans += i - left + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countKConstraintSubstrings(s string, k int) (ans int) {
	cnt := [2]int{}
	left := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[left]&1]--
			left++
		}
		ans += i - left + 1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。注意 $\textit{left}$ 只会增加不会减少，所以二重循环的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面滑动窗口题单中的「**§2.3.2 越短越合法**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
