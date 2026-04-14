对于每个 $q = \textit{queries}[i]$，遍历 $\textit{dictionary}$ 中的字符串 $s$，判断 $q$ 和 $s$ 是否至多有两个位置上的字母不同。

```py [sol-Python3]
class Solution:
    def twoEditWords(self, queries: List[str], dictionary: List[str]) -> List[str]:
        ans = []
        for q in queries:
            for s in dictionary:
                if sum(x != y for x, y in zip(q, s)) <= 2:
                    ans.append(q)
                    break
        return ans
```

```java [sol-Java]
class Solution {
    public List<String> twoEditWords(String[] queries, String[] dictionary) {
        List<String> ans = new ArrayList<>();
        for (String q : queries) {
            for (String s : dictionary) {
                int cnt = 0;
                for (int i = 0; i < s.length() && cnt <= 2; i++) {
                    if (q.charAt(i) != s.charAt(i)) {
                        cnt++;
                    }
                }
                if (cnt <= 2) {
                    ans.add(q);
                    break;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<string> twoEditWords(vector<string>& queries, vector<string>& dictionary) {
        vector<string> ans;
        for (auto& q : queries) {
            for (auto& s : dictionary) {
                int cnt = 0;
                for (int i = 0; i < s.size() && cnt <= 2; i++) {
                    if (q[i] != s[i]) {
                        cnt++;
                    }
                }
                if (cnt <= 2) {
                    ans.push_back(q);
                    break;
                }
            }
        }
        return ans;
    }
};
```

```c [sol-C]
char** twoEditWords(char** queries, int queriesSize, char** dictionary, int dictionarySize, int* returnSize) {
    char** ans = malloc(queriesSize * sizeof(char*));
    *returnSize = 0;

    for (int i = 0; i < queriesSize; i++) {
        char* q = queries[i];
        for (int j = 0; j < dictionarySize; j++) {
            char* s = dictionary[j];
            int cnt = 0;
            for (int k = 0; s[k] && cnt <= 2; k++) {
                if (q[k] != s[k]) {
                    cnt++;
                }
            }
            if (cnt <= 2) {
                ans[(*returnSize)++] = q;
                break;
            }
        }
    }

    return ans;
}
```

```go [sol-Go]
func twoEditWords(queries, dictionary []string) (ans []string) {
	for _, q := range queries {
	next:
		for _, s := range dictionary {
			cnt := 0
			for i := range s {
				if q[i] != s[i] {
					cnt++
					if cnt > 2 {
						continue next
					}
				}
			}
			ans = append(ans, q)
			break
		}
	}
	return
}
```

```js [sol-JavaScript]
var twoEditWords = function(queries, dictionary) {
    const ans = [];
    for (const q of queries) {
        for (const s of dictionary) {
            let cnt = 0;
            for (let i = 0; i < s.length && cnt <= 2; i++) {
                if (q[i] !== s[i]) {
                    cnt++;
                }
            }
            if (cnt <= 2) {
                ans.push(q);
                break;
            }
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn two_edit_words(queries: Vec<String>, dictionary: Vec<String>) -> Vec<String> {
        let mut ans = vec![];
        for q in queries {
            for s in &dictionary {
                let mut cnt = 0;
                for (a, b) in q.bytes().zip(s.bytes()) {
                    if a != b {
                        cnt += 1;
                        if cnt > 2 {
                            break;
                        }
                    }
                }
                if cnt <= 2 {
                    ans.push(q);
                    break;
                }
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(qdn)$，其中 $q$ 是 $\textit{queries}$ 的长度，$d$ 是 $\textit{dictionary}$ 的长度，$n$ 是 $\textit{queries}[i]$ 的长度。题目保证所有字符串长度相等。
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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
