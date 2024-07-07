枚举 $[0,2^n-1]$ 中的 $i$，如果 $i$ 的长为 $n$ 的二进制中，没有相邻的 $0$，那么将其二进制字符串加入答案。

怎么判断二进制中是否有相邻的 $0$？

我们可以把 $i$ 取反（保留低 $n$ 位），记作 $x$。问题变成：判断 $x$ 中是否有相邻的 $1$。

这可以用 `x & (x >> 1)` 来判断，如果这个值等于 $0$，则说明 $x$ 中没有相邻的 $1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Ry411q71f/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def validStrings(self, n: int) -> List[str]:
        ans = []
        mask = (1 << n) - 1
        for i in range(1 << n):
            x = mask ^ i
            if (x >> 1) & x == 0:
                ans.append(f"{i:0{n}b}")
        return ans
```

```java [sol-Java]
class Solution {
    public List<String> validStrings(int n) {
        List<String> ans = new ArrayList<>();
        int mask = (1 << n) - 1;
        for (int i = 0; i < (1 << n); i++) {
            int x = mask ^ i;
            if (((x >> 1) & x) == 0) {
                ans.add(Integer.toBinaryString((1 << n) | i).substring(1));
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> validStrings(int n) {
        vector<string> ans;
        int mask = (1 << n) - 1;
        for (int i = 0; i < (1 << n); i++) {
            int x = mask ^ i;
            if (((x >> 1) & x) == 0) {
                ans.push_back(bitset<18>(i).to_string().substr(18 - n));
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func validStrings(n int) (ans []string) {
	for i := 0; i < 1<<n; i++ {
		x := 1<<n - 1 ^ i
		if x>>1&x == 0 {
			ans = append(ans, fmt.Sprintf("%0*b", n, i))
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(2^n)$。更细致的分析表明（见视频讲解），答案的长度是**斐波那契数**，所以枚举 $i$ 的复杂度更高，时间复杂度为 $\mathcal{O}(2^n)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

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
