本题思路和 [735. 行星碰撞](https://leetcode.cn/problems/asteroid-collision/) 是一样的。

用列表 $\textit{toLeft}$ 维护向左的机器人，用栈 $\textit{st}$ 维护向右的机器人。

按照 $\textit{positions}[i]$ 从小到大排序。遍历机器人，如果遇到一个向左的机器人 $p$，分类讨论：

- 如果 $p$ 的健康度小于栈顶，那么栈顶的健康度减一。
- 如果 $p$ 的健康度等于栈顶，那么移除栈顶。
- 如果 $p$ 的健康度大于栈顶，那么移除栈顶，将 $p$ 的健康度减一后加入 $\textit{toLeft}$。
- 请注意，如果健康度减一，那么在减一之前，健康度一定是大于 $1$ 的，**不存在健康度减为** $0$ **的情况**。

最后剩余的就是 $\textit{toLeft}$ 和 $\textit{st}$ 中的机器人，合并，按照编号排序后，返回健康度列表。

代码实现时，也可以直接在 $\textit{healths}$ 上修改，最后返回 $\textit{healths}$ 中的正数。

视频讲解见[【周赛 351】](https://www.bilibili.com/video/BV1du41187ZN/)第四题，欢迎点赞！

```py [sol-Python3]
class Solution:
    def survivedRobotsHealths(self, positions: List[int], healths: List[int], directions: str) -> List[int]:
        n = len(positions)
        a = sorted(zip(range(n), positions, healths, directions), key=lambda p: p[1])
        to_left = []
        st = []
        for i, _, h, d in a:
            if d == 'R':  # 向右，存入栈中
                st.append([i, h])
                continue
            # 当前机器人向左，与栈中向右的机器人碰撞
            while st:
                top = st[-1]
                if top[1] > h:  # 栈顶的健康度大
                    top[1] -= 1
                    break
                if top[1] == h:  # 健康度一样大
                    st.pop()
                    break
                h -= 1  # 当前机器人的健康度大
                st.pop()  # 移除栈顶
            else:  # while 循环没有 break，说明当前机器人把栈中的全部撞掉
                to_left.append([i, h])
        to_left += st  # 合并剩余的机器人
        to_left.sort(key=lambda p: p[0])  # 按编号排序
        return [h for _, h in to_left]
```

```java [sol-Java]
class Solution {
    public List<Integer> survivedRobotsHealths(int[] positions, int[] healths, String directions) {
        int n = positions.length;
        var id = new Integer[n];
        for (int i = 0; i < n; ++i) id[i] = i;
        Arrays.sort(id, (i, j) -> positions[i] - positions[j]);

        var st = new ArrayDeque<Integer>();
        for (int i : id) {
            if (directions.charAt(i) == 'R') { // 向右，存入栈中
                st.push(i);
                continue;
            }
            // 向左，与栈中向右的机器人碰撞
            while (!st.isEmpty()) {
                int top = st.peek();
                if (healths[top] > healths[i]) { // 栈顶的健康度大
                    healths[top]--;
                    healths[i] = 0;
                    break;
                }
                if (healths[top] == healths[i]) { // 健康度一样大
                    healths[st.pop()] = 0;
                    healths[i] = 0;
                    break;
                }
                healths[st.pop()] = 0;
                healths[i]--; // 当前机器人的健康度大
            }
        }

        var ans = new ArrayList<Integer>();
        for (int h : healths) if (h > 0) ans.add(h);
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> survivedRobotsHealths(vector<int> &positions, vector<int> &healths, string directions) {
        int n = positions.size(), id[n];
        iota(id, id + n, 0);
        sort(id, id + n, [&](const int i, const int j) {
            return positions[i] < positions[j];
        });

        stack<int> st;
        for (int i: id) {
            if (directions[i] == 'R') { // 向右，存入栈中
                st.push(i);
                continue;
            }
            // 向左，与栈中向右的机器人碰撞
            while (!st.empty()) {
                int top = st.top();
                if (healths[top] > healths[i]) { // 栈顶的健康度大
                    healths[top]--;
                    healths[i] = 0;
                    break;
                }
                if (healths[top] == healths[i]) { // 健康度一样大
                    healths[top] = 0;
                    st.pop(); // 移除栈顶
                    healths[i] = 0;
                    break;
                }
                healths[top] = 0;
                st.pop(); // 移除栈顶
                healths[i]--; // 当前机器人的健康度大
            }
        }

        // 去掉 0
        healths.erase(remove(healths.begin(), healths.end(), 0), healths.end());
        return healths;
    }
};
```

```go [sol-Go]
func survivedRobotsHealths(positions, healths []int, directions string) []int {
	type data struct {
		i, p, h int
		d       byte
	}
	a := make([]data, len(positions))
	for i, p := range positions {
		a[i] = data{i, p, healths[i], directions[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].p < a[j].p })

	var toLeft, st []data
next:
	for _, p := range a {
		if p.d == 'R' { // 向右，存入栈中
			st = append(st, p)
			continue
		}
		// p 向左，与栈中向右的机器人碰撞
		for len(st) > 0 {
			top := &st[len(st)-1]
			if top.h > p.h { // 栈顶的健康度大
				top.h--
				continue next
			}
			if top.h == p.h { // 健康度一样大
				st = st[:len(st)-1]
				continue next
			}
			p.h-- // p 的健康度大
			st = st[:len(st)-1] // 移除栈顶
		}
		toLeft = append(toLeft, p) // p 把栈中的全部撞掉
	}

	// 合并剩余的机器人，按编号排序
	toLeft = append(toLeft, st...)
	sort.Slice(toLeft, func(i, j int) bool { return toLeft[i].i < toLeft[j].i })
	ans := make([]int, len(toLeft))
	for i, p := range toLeft {
		ans[i] = p.h
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{positions}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。

#### 相似题目

- [735. 行星碰撞](https://leetcode.cn/problems/asteroid-collision/)
