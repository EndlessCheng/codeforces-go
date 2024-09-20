从左到右遍历 $\textit{word}$。初始化 $x=0$，每遇到一个数字 $d$，就把 $x$ 更新为 $(10x+d)\bmod m$。

为什么可以在计算中取模？原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[视频讲解](https://www.bilibili.com/video/BV1wj411G7sH/)

```py [sol-Python3]
class Solution:
    def divisibilityArray(self, word: str, m: int) -> List[int]:
        ans = []
        x = 0
        for d in map(int, word):
            x = (x * 10 + d) % m
            ans.append(0 if x else 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int[] divisibilityArray(String word, int m) {
        char[] s = word.toCharArray();
        int[] ans = new int[s.length];
        long x = 0;
        for (int i = 0; i < s.length; i++) {
            x = (x * 10 + (s[i] - '0')) % m;
            if (x == 0) {
                ans[i] = 1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> divisibilityArray(string word, int m) {
        vector<int> ans(word.length());
        long long x = 0;
        for (int i = 0; i < word.length(); i++) {
            x = (x * 10 + (word[i] - '0')) % m;
            ans[i] = x == 0;
        }
        return ans;
    }
};
```

```go [sol-Go]
func divisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))
	x := 0
	for i, c := range word {
		x = (x*10 + int(c-'0')) % m
		if x == 0 {
			ans[i] = 1
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var divisibilityArray = function(word, m) {
    const ans = Array(word.length);
    let x = 0;
    for (let i = 0; i < word.length; i++) {
        x = (x * 10 + (word.charCodeAt(i) - '0'.charCodeAt(0))) % m;
        ans[i] = x === 0 ? 1 : 0;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn divisibility_array(word: String, m: i32) -> Vec<i32> {
        let mut ans = vec![0; word.len()];
        let mut x = 0i64;
        for (i, c) in word.bytes().enumerate() {
            x = (x * 10 + (c - b'0') as i64) % m as i64;
            if x == 0 {
                ans[i] = 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值的空间不计入。Java 忽略 $s$ 的空间。

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
