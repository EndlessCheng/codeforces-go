统计每个字母的出现次数，按照出现次数从大到小排序。

根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728)，出现次数前 $8$ 大的字母，只需要按一次；出现次数前 $9$ 到 $16$ 大的字母，需要按两次；依此类推。

把出现次数和对应的按键次数相乘再相加，得到的按键次数之和就是最小的。

请看 [视频讲解](https://www.bilibili.com/video/BV1Q5411C7mN/) 第三题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minimumPushes(self, word: str) -> int:
        a = sorted(Counter(word).values(), reverse=True)
        return sum(c * (i // 8 + 1) for i, c in enumerate(a))
```

```java [sol-Java]
class Solution {
    public int minimumPushes(String word) {
        int[] cnt = new int[26];
        for (char b : word.toCharArray()) {
            cnt[b - 'a']++;
        }
        Arrays.sort(cnt);

        int ans = 0;
        for (int i = 0; i < 26; i++) {
            ans += cnt[25 - i] * (i / 8 + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumPushes(string word) {
        int cnt[26]{};
        for (char b: word) {
            cnt[b - 'a']++;
        }
        ranges::sort(cnt, greater());

        int ans = 0;
        for (int i = 0; i < 26; i++) {
            ans += cnt[i] * (i / 8 + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumPushes(word string) (ans int) {
	cnt := [26]int{}
	for _, b := range word {
		cnt[b-'a']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cnt[:])))

	for i, c := range cnt {
		ans += c * (i/8 + 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|\log |\Sigma|)$，其中 $n$ 为 $\textit{word}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$O(|\Sigma|)$。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
