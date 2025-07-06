下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
business_line_to_category = {
    "electronics": 0,
    "grocery": 1,
    "pharmacy": 2,
    "restaurant": 3,
}

class Solution:
    def validateCoupons(self, code: List[str], businessLine: List[str], isActive: List[bool]) -> List[str]:
        groups = [[] for _ in range(4)]
        for s, bus, active in zip(code, businessLine, isActive):
            category = business_line_to_category.get(bus, -1)
            if s and category >= 0 and active and \
               all(c == '_' or c.isalnum() for c in s):
                groups[category].append(s)

        ans = []
        for g in groups:
            g.sort()
            ans += g
        return ans
```

```java [sol-Java]
class Solution {
    private static final Map<String, Integer> BUSINESS_LINE_TO_CATEGORY = Map.of(
        "electronics", 0,
        "grocery", 1,
        "pharmacy", 2,
        "restaurant", 3
    );

    public List<String> validateCoupons(String[] code, String[] businessLine, boolean[] isActive) {
        List<String>[] groups = new ArrayList[4];
        Arrays.setAll(groups, i -> new ArrayList<>());
        for (int i = 0; i < code.length; i++) {
            String s = code[i];
            Integer category = BUSINESS_LINE_TO_CATEGORY.get(businessLine[i]);
            if (!s.isEmpty() && category != null && isActive[i] && isValid(s)) {
                groups[category].add(s);
            }
        }

        List<String> ans = new ArrayList<>();
        for (List<String> g : groups) {
            Collections.sort(g);
            ans.addAll(g);
        }
        return ans;
    }

    // 判断是否只包含下划线或字母数字
    private boolean isValid(String s) {
        for (char c : s.toCharArray()) {
            if (c != '_' && !Character.isLetterOrDigit(c)) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
unordered_map<string, int> BUSINESS_LINE_TO_CATEGORY = {
    {"electronics", 0},
    {"grocery", 1},
    {"pharmacy", 2},
    {"restaurant", 3},
};

class Solution {
    // 检查字符串是否只包含字母、数字或下划线
    bool is_valid(const string& s) {
        for (char c : s) {
            if (c != '_' && !isalnum(c)) {
                return false;
            }
        }
        return true;
    }

public:
    vector<string> validateCoupons(vector<string>& code, vector<string>& businessLine, vector<bool>& isActive) {
        vector<string> groups[4];
        for (int i = 0; i < code.size(); i++) {
            string& s = code[i];
            auto it = BUSINESS_LINE_TO_CATEGORY.find(businessLine[i]);
            if (!s.empty() && it != BUSINESS_LINE_TO_CATEGORY.end() && isActive[i] && is_valid(s)) {
                groups[it->second].push_back(s);
            }
        }

        vector<string> ans;
        for (auto& g : groups) {
            ranges::sort(g);
            ans.insert(ans.end(), g.begin(), g.end());
        }
        return ans;
    }
};
```

```go [sol-Go]
var businessLineToCategory = map[string]int{
	"electronics": 0,
	"grocery":     1,
	"pharmacy":    2,
	"restaurant":  3,
}

func isValid(s string) bool {
	for _, c := range s {
		if c != '_' && !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return s != ""
}

func validateCoupons(code []string, businessLine []string, isActive []bool) (ans []string) {
	groups := [4][]string{}
	for i, s := range code {
		category, ok := businessLineToCategory[businessLine[i]]
		if ok && isActive[i] && isValid(s) {
			groups[category] = append(groups[category], s)
		}
	}

	for _, g := range groups {
		slices.Sort(g)
		ans = append(ans, g...)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L + n\log L)$，其中 $n$ 是 $\textit{code}$ 的长度，$L$ 是 $\textit{code}[i]$ 的长度之和。
- 空间复杂度：$\mathcal{O}(n)$。

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
