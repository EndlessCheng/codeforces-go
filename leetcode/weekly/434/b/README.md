注意输入的 $\textit{events}$ 不保证是按时间顺序发生的，需要先排序。

按照时间戳 $\textit{timestamp}$ 从小到大排序，时间戳相同的，离线事件排在前面，因为题目要求「状态变更在所有相同时间发生的消息事件之前处理」。

然后模拟：

- **离线事件**：用一个数组 $\textit{onlineT}$ 记录用户下次在线的时间戳（$60$ 秒后）。如果 $\textit{onlineT}[i]\le$ 当前时间戳，则表示用户 $i$ 已在线。
- **消息事件**：把相应用户的提及次数加一。

[本题视频讲解](https://www.bilibili.com/video/BV15sFNewEia/?t=3m39s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countMentions(self, numberOfUsers: int, events: List[List[str]]) -> List[int]:
        # 按照时间戳从小到大排序，时间戳相同的，离线事件排在前面
        events.sort(key=lambda e: (int(e[1]), e[0][2]))

        ans = [0] * numberOfUsers
        online_t = [0] * numberOfUsers
        for type_, timestamp, mention in events:
            cur_t = int(timestamp)  # 当前时间
            if type_[0] == 'O':  # 离线
                online_t[int(mention)] = cur_t + 60  # 下次在线时间
            elif mention[0] == 'A':  # @所有人
                for i in range(numberOfUsers):
                    ans[i] += 1
            elif mention[0] == 'H':  # @所有在线用户
                for i, t in enumerate(online_t):
                    if t <= cur_t:  # 在线
                        ans[i] += 1
            else:  # @id
                for s in mention.split():
                    ans[int(s[2:])] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int[] countMentions(int numberOfUsers, List<List<String>> events) {
        // 按照时间戳从小到大排序，时间戳相同的，离线事件排在前面
        events.sort((a, b) -> {
            int ta = Integer.parseInt(a.get(1));
            int tb = Integer.parseInt(b.get(1));
            return ta != tb ? ta - tb : b.get(0).charAt(0) - a.get(0).charAt(0);
        });

        int[] ans = new int[numberOfUsers];
        int[] onlineT = new int[numberOfUsers];
        for (List<String> e : events) {
            int curT = Integer.parseInt(e.get(1)); // 当前时间
            String mention = e.get(2);
            if (e.get(0).charAt(0) == 'O') { // 离线
                onlineT[Integer.parseInt(mention)] = curT + 60; // 下次在线时间
            } else if (mention.charAt(0) == 'A') { // @所有人
                for (int i = 0; i < numberOfUsers; i++) {
                    ans[i]++;
                }
            } else if (mention.charAt(0) == 'H') { // @所有在线用户
                for (int i = 0; i < numberOfUsers; i++) {
                    if (onlineT[i] <= curT) { // 在线
                        ans[i]++;
                    }
                }
            } else { // @id
                for (String s : mention.split(" ")) {
                    int i = Integer.parseInt(s.substring(2));
                    ans[i]++;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
#include<ranges>
class Solution {
public:
    vector<int> countMentions(int numberOfUsers, vector<vector<string>>& events) {
        // 按照时间戳从小到大排序，时间戳相同的，离线事件排在前面
        ranges::sort(events, {}, [](auto& e) {
            return pair(stoi(e[1]), e[0][2]);
        });

        vector<int> ans(numberOfUsers);
        vector<int> online_t(numberOfUsers);
        for (auto& e : events) {
            int cur_t = stoi(e[1]); // 当前时间
            string& mention = e[2];
            if (e[0][0] == 'O') { // 离线
                online_t[stoi(mention)] = cur_t + 60; // 下次在线时间
            } else if (mention[0] == 'A') { // @所有人
                for (int& v : ans) {
                    v++;
                }
            } else if (mention[0] == 'H') { // @所有在线用户
                for (int i = 0; i < numberOfUsers; i++) {
                    if (online_t[i] <= cur_t) { // 在线
                        ans[i]++;
                    }
                }
            } else { // @id
                for (const auto& part : mention | ranges::views::split(' ')) {
                    string s(part.begin() + 2, part.end());
                    ans[stoi(s)]++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func countMentions(numberOfUsers int, events [][]string) []int {
	// 按照时间戳从小到大排序，时间戳相同的，离线事件排在前面
	slices.SortFunc(events, func(a, b []string) int {
		ta, _ := strconv.Atoi(a[1])
		tb, _ := strconv.Atoi(b[1])
		return cmp.Or(ta-tb, int(b[0][0])-int(a[0][0]))
	})

	ans := make([]int, numberOfUsers)
	onlineT := make([]int, numberOfUsers)
	for _, e := range events {
		curT, _ := strconv.Atoi(e[1]) // 当前时间
		mention := e[2]
		if e[0][0] == 'O' { // 离线
			i, _ := strconv.Atoi(mention)
			onlineT[i] = curT + 60 // 下次在线时间
		} else if mention[0] == 'A' { // @所有人
			for i := range ans {
				ans[i]++
			}
		} else if mention[0] == 'H' { // @所有在线用户
			for i, t := range onlineT {
				if t <= curT { // 在线
					ans[i]++
				}
			}
		} else { // @id
			for _, s := range strings.Split(mention, " ") {
				i, _ := strconv.Atoi(s[2:])
				ans[i]++
			}
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var countMentions = function(numberOfUsers, events) {
    // 按照时间戳从小到大排序，时间戳相同的，离线事件排在前面
    events.sort((a, b) => parseInt(a[1]) - parseInt(b[1]) || b[0][0].charCodeAt(0) - a[0][0].charCodeAt(0));

    const ans = Array(numberOfUsers).fill(0);
    const onlineT = Array(numberOfUsers).fill(0);
    for (const [type, timestamp, mention] of events) {
        const curT = parseInt(timestamp); // 当前时间
        if (type[0] === 'O') { // 离线
            onlineT[parseInt(mention)] = curT + 60; // 下次在线时间
        } else if (mention[0] === 'A') { // @所有人
            for (let i = 0; i < numberOfUsers; i++) {
                ans[i]++;
            }
        } else if (mention[0] === 'H') { // @所有在线用户
            for (let i = 0; i < numberOfUsers; i++) {
                if (onlineT[i] <= curT) { // 在线
                    ans[i]++;
                }
            }
        } else { // @id
            for (const s of mention.split(" ")) {
                ans[parseInt(s.slice(2))] += 1;
            }
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_mentions(number_of_users: i32, mut events: Vec<Vec<String>>) -> Vec<i32> {
        // 按照时间戳从小到大排序，时间戳相同的，离线事件排在前面
        events.sort_unstable_by_key(|e| (e[1].parse::<i32>().unwrap(), e[0].as_bytes()[2]));

        let n = number_of_users as usize;
        let mut ans = vec![0; n];
        let mut online_t = vec![0; n];
        for e in events {
            let cur_t = e[1].parse().unwrap(); // 当前时间
            let mention = &e[2];
            if e[0].as_bytes()[0] == b'O' { // 离线
                online_t[mention.parse::<usize>().unwrap()] = cur_t + 60; // 下次在线时间
            } else if mention.as_bytes()[0] == b'A' { // @所有人
                for cnt in ans.iter_mut() {
                    *cnt += 1;
                }
            } else if mention.as_bytes()[0] == b'H' { // @所有在线用户
                for (&t, cnt) in online_t.iter().zip(ans.iter_mut()) {
                    if t <= cur_t { // 在线
                        *cnt += 1;
                    }
                }
            } else { // @id
                for s in mention.split(' ') {
                    ans[s[2..].parse::<usize>().unwrap()] += 1;
                }
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn + m\log m\log U + L)$，其中 $m$ 是 $\textit{events}$ 的长度，$n$ 是 $\textit{numberOfUsers}$，$U\le 10^5$ 是时间戳的最大值，$L$ 是所有 `mentions_string` 的长度之和。排序需要 $\mathcal{O}(m\log m)$ 次比较，每次比较需要 $\mathcal{O}(\log U)$ 的时间把长为 $\mathcal{O}(\log U)$ 的字符串时间戳转成整数。注：如果预处理这个转换，可以把排序的过程优化至 $\mathcal{O}(m\log m)$。
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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
