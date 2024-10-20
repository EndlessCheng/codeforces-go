注意我们只能添加字母，不能删除字母，也不能修改最后一个字母左边的字母，所以只有把第一个字母改成和 $\textit{target}[0]$ 一样后，才能继续向后添加/修改字母。

比如 $\textit{target}=\texttt{ccc}$，那么操作过程中产生的字符串一定是 

$$
\texttt{a},\texttt{b},\texttt{c},\texttt{ca},\texttt{cb},\texttt{cc},\texttt{cca},\texttt{ccb},\texttt{ccc}
$$

所以操作方式是**唯一**的，模拟即可。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1UcyYY4EnQ/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def stringSequence(self, target: str) -> List[str]:
        ans = []
        s = []
        for c in target:
            s.append('')  # 占位
            for s[-1] in ascii_lowercase[:ord(c) - ord('a') + 1]:
                ans.append(''.join(s))
        return ans
```

```java [sol-Java]
class Solution {
    List<String> stringSequence(String target) {
        List<String> ans = new ArrayList<>();
        StringBuilder s = new StringBuilder();
        for (int c : target.toCharArray()) {
            s.append('a'); // 占位
            for (char j = 'a'; j <= c; j++) {
                s.setCharAt(s.length() - 1, j);
                ans.add(s.toString());
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> stringSequence(string target) {
        vector<string> ans;
        string s;
        for (int c : target) {
            s += 'a'; // 占位
            for (char j = 'a'; j <= c; j++) {
                s.back() = j;
                ans.push_back(s);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func stringSequence(target string) (ans []string) {
	s := make([]byte, len(target))
	for i, c := range target {
		for j := byte('a'); j <= byte(c); j++ {
			s[i] = j
			ans = append(ans, string(s[:i+1]))
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(|\Sigma|n^2)$，其中 $n$ 是 $\textit{target}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

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
