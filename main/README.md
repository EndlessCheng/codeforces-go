## Codeforces 视频/文字题解精选

[我的洛谷博客](https://www.luogu.com.cn/blog/endlesscheng/)

### 手玩题（构造、思维）

|  题目 |  题解 | 难度  | 技能点  |
|---|---|---|---|
| [CF1922E Increasing Subsequences](https://codeforces.com/contest/1922/problem/E)| [视频](https://www.bilibili.com/video/BV1RK4y1q7qb/)| 1800 | 反向构造+递增子序列  |
| [CF1914F Programming Competition](https://codeforces.com/contest/1914/problem/F)| [视频](https://www.bilibili.com/video/BV1d94y1c7fb/)| 1900 | 树上贪心 |
| [CF1054D Changing Array](https://codeforces.com/problemset/problem/1054/D)  |  [文字](https://www.luogu.com.cn/blog/endlesscheng/solution-cf1054d) |  1900 | 正难则反 |
| [CF1665E MinimizOR](https://codeforces.com/problemset/problem/1665/E)  |  [文字](https://www.luogu.com.cn/blog/endlesscheng/solution-cf1665e) |  2500 | 从特殊到一般  |

### 动态规划

|  题目 |  题解 | 难度  | 技能点  |
|---|---|---|---|
| [CF452D Washer, Dryer, Folder](https://codeforces.com/problemset/problem/452/D)  |  [文字](https://www.luogu.com.cn/blog/endlesscheng/solution-cf452d) |  1900 | 状态设计  |
| [CF1733D2 Zero-One](https://codeforces.com/problemset/problem/1733/D2)  |  [文字](https://www.luogu.com.cn/blog/endlesscheng/solution-cf1733d2) |  2000 | 状态设计  |
| [CF1913D Array Collapse](https://codeforces.com/contest/1913/problem/D)  |  [视频](https://www.bilibili.com/video/BV1v64y1p7pi/) | 2100  | 单调栈+前缀和优化 DP |
| [CF1327F AND Segments](https://codeforces.com/problemset/problem/1327/F)  |  [文字](https://www.luogu.com.cn/blog/endlesscheng/cf1327f-and-segments-ti-xie) |  2500 | 状态设计+滑窗优化  |
| [CF1237F Balanced Domino Placements](https://codeforces.com/problemset/problem/1237/F)  |  [文字](https://www.luogu.com.cn/article/tekpxn7n) |  2600 | 计数 DP  |

### 数据结构

|  题目 |  题解 | 难度  | 技能点  |
|---|---|---|---|
| [CF1921F Sum of Progression](https://codeforces.com/contest/1921/problem/F)  |  [视频](https://www.bilibili.com/video/BV1hQ4y1L7Tk/) | 1900  | 带权前缀和+根号算法 |
| [CF1730E Maximums and Minimums](https://codeforces.com/contest/1730/problem/E)  |  [文字](https://www.luogu.com.cn/blog/endlesscheng/solution-cf1730e) | 2700  | 单调栈 |

### 数学综合

|  题目 |  题解 | 难度  | 技能点  |
|---|---|---|---|
| [CF1790E Vlad and a Pair of Numbers](https://codeforces.com/contest/1790/problem/E)  |  [文字](https://www.luogu.com.cn/blog/endlesscheng/solution-cf1790e) | 1400   | 位运算恒等式 |

---

## 代码框架

编写一个 `run(io.Reader, io.Writer)` 函数来处理输入输出。这样写的理由是：

- 在 `main` 中调用 `run(os.Stdin, os.Stdout)` 来执行代码；
- 测试时，将测试数据转换成 `strings.Reader` 当作输入，并用一个 `strings.Builder` 来接收输出，将这二者传入 `run` 中，然后就能比较输出与答案了；
- 对拍时需要实现一个暴力算法 `runAC`，参数和 `run` 一样。通过[随机数据生成器](/main/testutil/rand.go)来生成数据，分别传入 `runAC` 和 `run`，通过比对各自的输出，来检查 `run` 中的问题。

例如：[1439C_test.go](./1400-1499/1439C_test.go)

交互题的写法要复杂一些，为方便测试，需要把涉及输入输出的地方抽象成接口，详见 [interactive_problem](/copypasta/template/interactive_problem)。
