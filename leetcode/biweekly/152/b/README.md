创建一个哈希表，直接把整个 $\textit{cell}$ 字符串作为哈希表的键，无需解析行号列号。

- $\texttt{setCell}$：把键值对 $\textit{cell}$ 和 $\textit{value}$ 插入哈希表。
- $\texttt{resetCell}$：把 $\textit{cell}$ 从哈希表中删除。
- $\texttt{getValue}$：去掉第一个字符（$\texttt{=}$），用 $\texttt{+}$ 分割字符串，得到两个字符串，分别转成整数再相加：
   - 如果字符串的第一个字符是大写字母，那么查找哈希表，得到对应的 $\textit{value}$（没有就是 $0$）。
   - 否则，把字符串转成整数。

```py [sol-Python3]
class Spreadsheet:
    def __init__(self, _):
        self.data = {}

    def setCell(self, cell: str, value: int) -> None:
        self.data[cell] = value

    def resetCell(self, cell: str) -> None:
        self.data.pop(cell, None)

    def getValue(self, formula: str) -> int:
        ans = 0
        for cell in formula[1:].split("+"):
            # 注：如果用 defaultdict(int)，访问 self.data[cell] 也会把 cell 插入哈希表，增加空间复杂度
            ans += self.data.get(cell, 0) if cell[0].isupper() else int(cell)
        return ans
```

```java [sol-Java]
class Spreadsheet {
    private final Map<String, Integer> data = new HashMap<>();

    public Spreadsheet(int rows) {
    }

    public void setCell(String cell, int value) {
        data.put(cell, value);
    }

    public void resetCell(String cell) {
        data.remove(cell);
    }

    public int getValue(String formula) {
        int ans = 0;
        for (String cell : formula.substring(1).split("\\+")) {
            if (Character.isUpperCase(cell.charAt(0))) {
                ans += data.getOrDefault(cell, 0);
            } else {
                ans += Integer.parseInt(cell);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Spreadsheet {
    unordered_map<string, int> data;

public:
    Spreadsheet(int) {}

    void setCell(string cell, int value) {
        data[cell] = value;
    }

    void resetCell(string cell) {
        data.erase(cell);
    }

    int getValue(string formula) {
        int i = formula.find('+');
        string s = formula.substr(1, i - 1);
        string t = formula.substr(i + 1);
        // 注意 s 不在 data 中的时候，data[s] 会把 s 插入 data，这里从简没有判断
        return (isupper(s[0]) ? data[s] : stoi(s)) +
               (isupper(t[0]) ? data[t] : stoi(t));
    }
};
```

```go [sol-Go]
type Spreadsheet map[string]int

func Constructor(int) Spreadsheet {
	return Spreadsheet{}
}

func (s Spreadsheet) SetCell(cell string, value int) {
	s[cell] = value
}

func (s Spreadsheet) ResetCell(cell string) {
	delete(s, cell)
}

func (s Spreadsheet) GetValue(formula string) (ans int) {
	for _, cell := range strings.Split(formula[1:], "+") {
		if unicode.IsUpper(rune(cell[0])) {
			ans += s[cell]
		} else {
			x, _ := strconv.Atoi(cell)
			ans += x
		}
	}
	return
}
```

```js [sol-JavaScript]
class Spreadsheet {
    constructor(rows) {
        this.data = new Map();
    }

    setCell(cell, value) {
        this.data.set(cell, value);
    }

    resetCell(cell) {
        this.data.delete(cell);
    }

    getValue(formula) {
        let ans = 0;
        for (const cell of formula.slice(1).split("+")) {
            if ("A" <= cell[0] && cell[0] <= "Z") {
                ans += this.data.get(cell) ?? 0;
            } else {
                ans += parseInt(cell);
            }
        }
        return ans;
    }
}
```

```rust [sol-Rust]
use std::collections::HashMap;

struct Spreadsheet {
    data: HashMap<String, i32>,
}

impl Spreadsheet {
    fn new(_: i32) -> Self {
        Spreadsheet {
            data: HashMap::new(),
        }
    }

    fn set_cell(&mut self, cell: String, value: i32) {
        self.data.insert(cell, value);
    }

    fn reset_cell(&mut self, cell: String) {
        self.data.remove(&cell);
    }

    fn get_value(&self, formula: String) -> i32 {
        formula[1..]
            .split('+')
            .map(|cell| {
                if cell.as_bytes()[0].is_ascii_uppercase() {
                    *self.data.get(cell).unwrap_or(&0)
                } else {
                    cell.parse::<i32>().unwrap()
                }
            })
            .sum()
    }
}
```

#### 复杂度分析

- 时间复杂度：初始化为 $\mathcal{O}(1)$，其余为 $\mathcal{O}(L)$，其中 $L$ 是 $\textit{cell}$（或者 $\textit{formula}$）的长度。
- 空间复杂度：$\mathcal{O}(qL)$。其中 $q$ 为 $\texttt{setCell}$ 的调用次数。

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
