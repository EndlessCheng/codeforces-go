遍历 $\textit{words}$，判断 $x$ 是否在 $\textit{words}[i]$ 中，如果是则把 $i$ 加入答案。

```py [sol-Python3]
class Solution:
    def findWordsContaining(self, words: List[str], x: str) -> List[int]:
        return [i for i, s in enumerate(words) if x in s]
```

```java [sol-Java]
class Solution {
    public List<Integer> findWordsContaining(String[] words, char x) {
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < words.length; i++) {
            if (words[i].indexOf(x) >= 0) {
                ans.add(i);
            }
        }
        return ans;
    }
}
```

```java [sol-Java Stream]
class Solution {
    public List<Integer> findWordsContaining(String[] words, char x) {
        return IntStream.range(0, words.length)
                        .filter(i -> words[i].indexOf(x) >= 0)
                        .boxed()
                        .toList();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findWordsContaining(vector<string>& words, char x) {
        vector<int> ans;
        for (int i = 0; i < words.size(); i++) {
            if (words[i].contains(x)) {
                ans.push_back(i);
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int* findWordsContaining(char** words, int wordsSize, char x, int* returnSize) {
    int* ans = malloc(sizeof(int) * wordsSize);
    int k = 0;
    for (int i = 0; i < wordsSize; i++) {
        if (strchr(words[i], x)) {
            ans[k++] = i;
        }
    }
    *returnSize = k;
    return ans;
}
```

```go [sol-Go]
func findWordsContaining(words []string, x byte) (ans []int) {
	for i, s := range words {
		if strings.IndexByte(s, x) >= 0 {
			ans = append(ans, i)
		}
	}
	return
}
```

```js [sol-JS]
var findWordsContaining = function(words, x) {
    const ans = [];
    for (let i = 0; i < words.length; i++) {
        if (words[i].includes(x)) {
            ans.push(i);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_words_containing(words: Vec<String>, x: char) -> Vec<i32> {
        words.into_iter()
             .enumerate()
             .filter_map(|(i, s)| s.contains(x).then_some(i as i32))
             .collect()
    }
}
```

```rust [sol-Rust 写法二]
impl Solution {
    pub fn find_words_containing(words: Vec<String>, x: char) -> Vec<i32> {
        let mut ans = vec![];
        for (i, s) in words.into_iter().enumerate() {
            if s.contains(x) {
                ans.push(i as i32);
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 为所有 $\textit{words}[i]$ 的长度之和。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
