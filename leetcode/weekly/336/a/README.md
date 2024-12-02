遍历下标在 $[\textit{left},\textit{right}]$ 内的字符串 $s=\textit{words}[i]$，如果 $s$ 的第一个字母和最后一个字母都是元音，把答案加一。

```py [sol-Python3]
class Solution:
    def vowelStrings(self, words: List[str], left: int, right: int) -> int:
        return sum(s[0] in "aeiou" and s[-1] in "aeiou" for s in words[left:right+1])
```

```java [sol-Java]
class Solution {
    private static final String VOWEL = "aeiou";

    public int vowelStrings(String[] words, int left, int right) {
        int ans = 0;
        for (int i = left; i <= right; i++) {
            String s = words[i];
            if (VOWEL.indexOf(s.charAt(0)) != -1 &&
                VOWEL.indexOf(s.charAt(s.length() - 1)) != -1) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int vowelStrings(vector<string>& words, int left, int right) {
        const string vowel = "aeiou";
        int ans = 0;
        for (int i = left; i <= right; i++) {
            string& s = words[i];
            ans += vowel.find(s[0]) != string::npos &&
                   vowel.find(s.back()) != string::npos;
        }
        return ans;
    }
};
```

```go [sol-Go]
func vowelStrings(words []string, left, right int) (ans int) {
	for _, s := range words[left : right+1] {
		if strings.Contains("aeiou", s[:1]) && strings.Contains("aeiou", s[len(s)-1:]) {
			ans++
		}
	}
	return
}
```

```js [sol-JavaScript]
var vowelStrings = function(words, left, right) {
    let ans = 0;
    for (let i = left; i <= right; i++) {
        const s = words[i];
        if ("aeiou".includes(s[0]) && "aeiou".includes(s[s.length - 1])) {
            ans++;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn vowel_strings(words: Vec<String>, left: i32, right: i32) -> i32 {
        words[left as usize..=right as usize]
            .iter()
            .filter(|s|
                "aeiou".contains(s.chars().next().unwrap()) &&
                "aeiou".contains(s.chars().last().unwrap())
            ).count() as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{right}-\textit{left})$。注意每个字符串只取第一个和最后一个字母，不会遍历整个字符串，所以处理每个字符串的时间是 $\mathcal{O}(1)$ 的。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
