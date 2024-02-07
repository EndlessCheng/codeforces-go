# 算法竞赛模板库 by 灵茶山艾府 💭💡🎈

## 算法 Algorithm

由于算法知识点繁杂，将自己学习到的算法、做过的题目分类整理好是有必要的。

一个算法模板应当涵盖以下几点：
- 对该算法的基本介绍（核心思想、复杂度等）
- 参考链接或书籍章节（讲的比较好的资料）
- 模板代码（可以包含一些注释、使用说明）
- 模板补充内容（常见题型中的额外代码、建模技巧等）
- 相关题目链接（模板题、经典题、思维转换题等）


## 算法目录

[不了解 Go？快速入门教程](https://gobyexample-cn.github.io/)

- [集合论与位运算](https://leetcode.cn/circle/discuss/CaOJ45/)
- 数据结构
  - [单调栈 monotone_stack.go](/copypasta/monotone_stack.go)
  - [单调队列 monotone_queue.go](/copypasta/monotone_queue.go)
    - 二维单调队列
  - [双端队列 deque.go](/copypasta/deque.go)
  - [堆（优先队列）heap.go](/copypasta/heap.go)
    - 支持修改、删除指定元素的堆
    - 懒删除堆
    - 对顶维
    - 前缀中位数
    - 滑动窗口前 k 小元素和
  - [并查集 union_find.go](/copypasta/union_find.go)
    - 点权并查集
    - 边权并查集（种类并查集）
    - 可持久化并查集
    - 回滚并查集 & 动态图连通性
  - [稀疏表（ST 表）sparse_table.go](/copypasta/sparse_table.go)
  - [树状数组 fenwick_tree.go](/copypasta/fenwick_tree.go)
    - 差分树状数组（支持区间加、区间求和）
    - 树套树 & 三维偏序
  - [线段树 segment_tree.go](/copypasta/segment_tree.go)
    - 线段树二分
    - 延迟标记（懒标记）
    - 动态开点
    - 线段树合并
    - 线段树分裂
    - 持久化（主席树）
  - [0-1 线段树 segment_tree01.go](/copypasta/segment_tree01.go)
  - [左偏树（可并堆）leftist_tree.go](/copypasta/leftist_tree.go)
  - [笛卡尔树 cartesian_tree.go](/copypasta/cartesian_tree.go)
  - [二叉搜索树公共方法 bst.go](/copypasta/bst.go)
  - [Treap treap.go](/copypasta/treap.go)
    - [前 k 小元素和](/copypasta/treap_kthsum.go)
  - [伸展树 splay.go](/copypasta/splay.go)
  - [动态树 LCT link_cut_tree.go](/copypasta/link_cut_tree.go)
  - [红黑树 red_black_tree.go](/copypasta/red_black_tree.go)
  - [替罪羊树 scapegoat_tree.go](/copypasta/scapegoat_tree.go)
  - [k-d 树 kd_tree.go](/copypasta/kd_tree.go)
  - 珂朵莉树（ODT）
    - [数组版 odt.go](/copypasta/odt.go)
    - [平衡树版 odt_bst.go](/copypasta/odt_bst.go)
  - [根号分治、分块 sqrt_decomposition.go](/copypasta/sqrt_decomposition.go)
  - [莫队算法 mo.go](/copypasta/mo.go)
     - 普通莫队
     - 带修莫队
     - 回滚莫队
     - 树上莫队
- [字符串 strings.go](/copypasta/strings.go)
  - 字符串哈希
  - KMP
    - pi 函数
    - border
    - 最小循环节
    - fail 树（失配树 / border 树）
  - 扩展 KMP（Z algorithm）
  - 最小表示法
  - 最长回文子串 
    - Manacher 算法
  - [回文自动机（回文树，PAM）pam.go](/copypasta/pam.go)
  - 后缀数组（SA）
  - [后缀自动机（SAM）sam.go](/copypasta/sam.go)
  - [字典树 trie.go](/copypasta/trie.go)
    - 可持久化字典树
  - [0-1 字典树 trie01.go](/copypasta/trie01.go)
    - 最大异或和
    - 第 k 大异或和
    - 删除元素
    - 可持久化 0-1 字典树
    - 【研究】0-1 字典树上最多有多少个节点
  - [AC 自动机 acam.go](/copypasta/acam.go)
- 数学
  - [数论 math.go](/copypasta/math.go)
    - 辗转相除法（最大公因数 GCD）
    - 类欧几里得算法 ∑⌊(ai+b)/m⌋
    - Pollard-Rho 质因数分解算法
    - 埃氏筛（埃拉托斯特尼筛法）
    - 欧拉筛（线性筛）
    - 欧拉函数
    - 原根
    - 扩展 GCD
      - 二元一次不定方程
    - 逆元
      - 线性求逆元
    - 中国剩余定理（CRT）
      - 扩展中国剩余定理
    - 离散对数
    - 大步小步算法（BSGS）
      - 扩展大步小步算法
    - 二次剩余
    - Jacobi 符号
    - N 次剩余
    - 卢卡斯定理
      - 扩展卢卡斯定理
    - 卡特兰数
    - 默慈金数
    - 那罗延数
    - 斯特林数
      - 第一类斯特林数（轮换）
      - 第二类斯特林数（子集）
    - 贝尔数
    - 欧拉数
    - 数论分块（整除分块）
    - 莫比乌斯函数
    - 莫比乌斯反演
      - 互质计数问题
      - GCD 求和问题
    - 杜教筛
  - [组合数学 math_comb.go](/copypasta/math_comb.go)
    - 常见模型
    - 常用恒等式
    - 容斥原理
  - [快速傅里叶变换 FFT math_fft.go](/copypasta/math_fft.go)
  - [快速数论变换 NTT math_ntt.go](/copypasta/math_ntt.go)
    - 包含多项式全家桶（求逆、开方等等）
  - [快速沃尔什变换 FWT math_fwt.go](/copypasta/math_fwt.go)
  - [连分数、佩尔方程 math_continued_fraction.go](/copypasta/math_continued_fraction.go)
  - [线性代数 math_matrix.go](/copypasta/math_matrix.go)
    - 矩阵相关运算
    - 高斯消元
    - 行列式
    - 线性基
  - [数值分析 math_numerical_analysis.go](copypasta/math_numerical_analysis.go)
    - 自适应辛普森积分
    - 拉格朗日插值
  - [计算几何 geometry.go](/copypasta/geometry.go)
    - 线与点
    - 线与线
    - 圆与点
      - 最小圆覆盖
        - Welzl 随机增量法
      - 固定半径覆盖最多点
    - 圆与线
    - 圆与圆
    - 圆与矩形
    - 最近点对
    - 多边形与点
      - 判断点在凸多边形内 O(log n)
      - 判断点在任意多边形内
        - 转角法（统计绕数）
    - 凸包
    - 最远点对
      - 旋转卡壳
    - 半平面交
  - [博弈论 games.go](/copypasta/games.go)
    - SG 函数
- [动态规划 dp.go](/copypasta/dp.go)
  - 背包
    - 0-1 背包
    - 完全背包
    - 多重背包
      - 二进制优化
      - 单调队列优化
      - 同余前缀和优化（求方案数）
    - 分组背包
    - 树上背包（依赖背包）
    - 字典序最小方案
  - 线性 DP
    - 最大子段和
    - LCS
    - LPS
    - LIS
      - 狄尔沃斯定理
    - LCIS
    - 长度为 m 的 LIS 个数
    - 本质不同子序列个数
  - 区间 DP
  - 环形 DP
  - 博弈 DP
  - 概率 DP
  - 期望 DP
  - 状压 DP
    - 全排列 DP
    - 旅行商问题（TSP）
    - 子集 DP
    - 高维前缀和（SOS DP）
    - 插头 DP
  - 数位 DP
    - 记忆化搜索（同时跑上下界）
  - 倍增优化 DP
  - 斜率优化 DP（CHT）
  - WQS 二分优化 DP（凸优化 DP / 带权二分）
  - 树形 DP
    - 树的直径个数
    - 在任一直径上的节点个数
    - 树上最大独立集
    - 树上最小顶点覆盖
    - 树上最小支配集
    - 树上最大匹配
    - 换根 DP（二次扫描法）
- [图论 graph.go](/copypasta/graph.go)
  - 链式前向星
  - DFS 常用技巧
  - BFS 常用技巧
  - 欧拉回路和欧拉路径
    - 无向图
    - 有向图
  - 割点
  - 割边（桥）
  - 双连通分量（BCC）
    - v-BCC
    - e-BCC
  - 仙人掌 & 圆方树
  - 最短路
    - Dijkstra
    - SPFA（队列优化的 Bellman-Ford）
      - 差分约束系统
    - Floyd-Warshall
    - Johnson
    - 0-1 BFS（双端队列 BFS）
    - 字典序最小最短路
    - 同余最短路
  - 最小环
  - 最小斯坦纳树
  - 最小生成树（MST）
    - Kruskal
    - Prim
  - 单度限制最小生成树
  - 次小生成树
  - 曼哈顿距离最小生成树
  - 最小差值生成树
  - 最小树形图
    - 朱刘算法
  - 二分图判定（染色）
  - 二分图找奇环
  - 二分图最大匹配
    - 匈牙利算法
  - 带权二分图最大完美匹配
    - Kuhn–Munkres 算法
  - 拓扑排序
  - 强连通分量（SCC）
    - Kosaraju
    - Tarjan
  - 2-SAT
  - 基环树
  - 最大流
    - Dinic
    - ISAP
    - HLPP
  - 最小费用最大流
    - SPFA
    - Dijkstra
  - 三元环计数
  - 四元环计数
  - [树上问题 graph_tree.go](/copypasta/graph_tree.go)
    - 直径
    - 重心
    - 点分治
    - 点分树
    - 最近公共祖先（LCA）
      - 倍增
      - ST 表
      - Tarjan
      - 树上差分
      - 虚树
    - 重链剖分（HLD）
    - 长链剖分
    - 树上启发式合并（small to large）
      - 按大小合并
      - 轻重儿子合并
    - 树分块
    - Prufer 序列
- 其他
  - [位运算笔记 bits.go](/copypasta/bits.go)
    - bitset
    - 区间位运算 trick（含 GCD）
  - [二分 三分 sort.go](/copypasta/sort.go)
    - 二分答案
    - 0-1 分数规划
    - 整体二分
  - [搜索 search.go](/copypasta/search.go)
    - 枚举排列
    - 枚举组合
    - 生成下一个排列
    - 康托展开
    - 逆康托展开
    - 枚举子集
      - Gosper's Hack
    - 折半枚举（Meet in the middle）
      - 超大背包问题
  - [随机算法 rand.go](/copypasta/rand.go)
    - 模拟退火
  - [基础算法 common.go](/copypasta/common.go)
    - 算法思路整理
    - 滑动窗口
    - 前缀和
    - 二维前缀和
    - 二维差分
    - 离散化
  - [杂项 misc.go](/copypasta/misc.go)
- [快速输入输出模板 io.go](/copypasta/io.go)

欢迎关注 bilibili[@灵茶山艾府](https://space.bilibili.com/206214)，高质量算法教学，持续输出中！

## 如何选择题目 How to Choose Problems

### Rating < 2100

这一阶段主要目标是提高对问题的观察能力。做构造题可以针对性地训练这一点。

选择难度在自己 rating 到 rating+200 范围内的构造题 (tag: constructive algorithms)，按照过题人数降序做题，比如 [1700,1900] 区间的就是下面这个链接：

[https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=constructive+algorithms%2C1700-1900](https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=constructive+algorithms%2C1700-1900)

通过大量的构造题训练，提高观察能力，快速找到切题入口。具体见我在知乎上的这篇 [回答](https://www.zhihu.com/question/353734418/answer/2353160035)。

### Rating >= 2100（个人训练用，仅供参考）

见识更高的山、更广的海。

按人数从高到低，做 2200+ 的题目。**建议不设置难度上限**！由于按人数排序，难度分不会太高，**不设上限可以避免错过高分好题**。

- [构造题 2200+](https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=constructive+algorithms%2C2200-)：锻炼手玩能力。
- [DP 2200+](https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=dp%2C2200-)：几乎每场都有 DP。
- [数学综合：数论、组合数学、概率期望等 2200+](https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=combine-tags-by-or%2Ccombinatorics%2Cfft%2Cmatrices%2Cnumber+theory%2Cprobabilities%2Cchinese+remainder+theorem%2C2200-)：包含 6 个 tag。
- [图论综合：图论+树上问题 2200+](https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=combine-tags-by-or%2C2-sat%2Cdsu%2Cflows%2Cgraph+matchings%2Cgraphs%2Cshortest+paths%2Ctrees%2C2200-)：包含 7 个 tag。
- [字符串 2200+](https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=combine-tags-by-or%2Cstring+suffix+structures%2Cstrings%2C2200-)：数据结构题不好筛选，可以找树状数组/线段树的题单，这里只单独筛选字符串的题。
- [交互 2200+](https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=interactive%2C2200-)：偶尔做做，了解一些解题套路。
- [博弈 2000+](https://codeforces.com/problemset?order=BY_SOLVED_DESC&tags=games%2C2000-)：也适合锻炼手玩。由于题目比较少，从 2000 开始筛选。

**我的 Codeforces 账号**

[![0x3F](https://img.shields.io/badge/0x3F-MASTER%202208-orange?style=for-the-badge)](https://codeforces.com/profile/0x3F)


## 测试及对拍 Testing

编写一个 `run(io.Reader, io.Writer)` 函数来处理输入输出。这样写的理由是：

- 在 `main` 中调用 `run(os.Stdin, os.Stdout)` 来执行代码；
- 测试时，将测试数据转换成 `strings.Reader` 当作输入，并用一个 `strings.Builder` 来接收输出，将这二者传入 `run` 中，然后就能比较输出与答案了；
- 对拍时需要实现一个暴力算法 `runAC`，参数和 `run` 一样。通过 [随机数据生成器](/main/testutil/rand.go) 来生成数据，分别传入 `runAC` 和 `run`，通过比对各自的输出，来检查 `run` 中的问题。

具体可以见 Codeforces 代码仓库 [main](/main)，所有非交互题的代码及其对应测试全部按照上述框架实现。

例如：[1439C_test.go](/main/1400-1499/1439C_test.go)

交互题的写法要复杂一些，需要把涉及输入输出的地方抽象成接口，详见 [interactive_problem](/copypasta/template/interactive_problem)。

## 学习资料及题目 Resources

注：由于入门经典上选了很多区域赛的题，一部分题目可以在 GYM 上找到，这样可以就可以用 Go 编程提交了。

[算法竞赛入门经典（第二版）](https://github.com/aoapc-book/aoapc-bac2nd)

[算法竞赛入门经典训练指南](https://github.com/klb3713/aoapc-book/tree/master/TrainingGuide/bookcodes)

[算法竞赛入门经典训练指南（升级版）](https://gitee.com/sukhoeing/aoapc-training-guide2)

[算法竞赛进阶指南](https://github.com/lydrainbowcat/tedukuri)

[算法竞赛入门到进阶](https://github.com/luoyongjun999/code)

[《算法竞赛》配套题单](https://www.luogu.com.cn/training/441063)

[OI Public Library（含国家队论文）](https://github.com/enkerewpo/OI-Public-Library)

[算法竞赛 (ICPC, OI, etc) 论文，课件，文档，笔记等](https://github.com/LzyRapx/Competitive-Programming-Docs)

[算法竞赛课件分享 by hzwer](https://github.com/hzwer/shareOI)

[算法第四版 Java 源码](https://algs4.cs.princeton.edu/code/)

[数据结构和算法动态可视化](https://visualgo.net/zh)

[OI Wiki](https://oi-wiki.org/)

[CP-Algorithms](https://cp-algorithms.com/)

[The Ultimate Topic List (with Resources, Problems and Templates)](https://codeforces.com/blog/entry/95106)

[洛谷日报](https://www.craft.do/s/N0l80k2gv46Psq)

[All the good tutorials found for Competitive Programming](https://codeforces.com/blog/entry/57282)

[Codeforces Problem Topics](https://codeforces.com/blog/entry/55274)

[The Ultimate Topic List(with Tutorials, Problems, and Templates)](https://blog.shahjalalshohag.com/topic-list/)

[GeeksforGeeks 上的算法合集](https://www.geeksforgeeks.org/how-to-prepare-for-acm-icpc/)

[Pepcy 模板](http://pepcy.cf/icpc-templates/)

[F0RE1GNERS 模板](https://github.com/F0RE1GNERS/template)

https://github.com/hh2048/XCPC 含 jiangly 模板

https://www.cnblogs.com/alex-wei/p/contents.html

[【模板整合计划】目录](https://www.cnblogs.com/Xing-Ling/p/10930556.html)

[算法学习笔记（目录）](https://zhuanlan.zhihu.com/p/105467597)

[洛谷模板题（建议按难度筛选）](https://www.luogu.com.cn/problem/list?keyword=%E6%A8%A1%E6%9D%BF&page=1)

[能力全面提升综合题单](https://www.luogu.com.cn/training/9391)

[Luogu Problem List](https://github.com/SFOI-Team/luogu-problem-list/blob/master/list.md)

[洛谷原试炼场](https://www.luogu.com.cn/paste/0id3h6on)

[Links of ICPC/CCPC Contests from China](https://codeforces.com/blog/entry/84429)

[AtCoder 题目分类](https://atcoder-tags.herokuapp.com/explain)

### AtCoder 版《挑战程序设计竞赛》

[AtCoder 版！蟻本 (初級編)](https://qiita.com/drken/items/e77685614f3c6bf86f44)

[AtCoder 版！蟻本 (中級編)](https://qiita.com/drken/items/2f56925972c1d34e05d8)

[AtCoder 版！蟻本 (上級編)](https://qiita.com/drken/items/9b311d553aa434bb26e4)

[AtCoder 版！蟻本 (発展的トピック編)](https://qiita.com/drken/items/0de3d205690d92307b7c)

### 待整理

[【杂文】记一些有用的神奇网站](https://www.cnblogs.com/Xing-Ling/p/10897760.html)

[偶然在 GitHub 上发现的超长列表](https://github.com/dhs347/Dream/blob/master/%E8%AE%A1%E5%88%92/%E8%AE%A1%E5%88%92%E4%B9%A6/A%E8%AE%A1%E5%88%92_%E9%98%B6%E6%AE%B51.md)

[算法竞赛训练中较难的部分](https://blog.csdn.net/skywalkert/article/details/48924861)

[算法竞赛中可能不太会遇到的论文题](https://blog.csdn.net/skywalkert/article/details/48878925)

[[杂谈]OI/ACM中冷门算法](https://zhuanlan.zhihu.com/p/21924647)

[Things I don't know](https://codeforces.com/blog/entry/92248)

> [meme] If you know at least 3 of these things and you are not red — you are doing it wrong. Stop learning useless algorithms, go and solve some problems, learn how to use binary search.

https://blog.csdn.net/calabash_boy/article/details/79973483

https://github.com/zimpha/algorithmic-library

https://www.luogu.com.cn/blog/command-block/blog-suo-yin-zhi-ding-post

https://wcysai.github.io/

https://www.luogu.com.cn/blog/Troverld/index

[C++ @cache](https://codeforces.com/blog/entry/124683)

## 其他 Others

My GoLand `Live Templates` and `Postfix Completion` [settings](/misc/my_goland_template)

### Useful Tools

[GeoGebra](https://www.geogebra.org/classic)

[Draw Geometry](https://csacademy.com/app/geometry_widget/)

[Draw Graph](https://csacademy.com/app/graph_editor/)

[OEIS](https://oeis.org/)

[Wolfram|Alpha](https://www.wolframalpha.com/)

[UpSolve.me](https://upsolve.me/)

[Codeforces Upsolving Helper](https://codeforces-upsolving-helper.herokuapp.com/)

[Contests Filter](https://codeforceshelper.herokuapp.com/contests)

[Codeforced](http://codeforced.github.io/handle/)

[Codeforces Visualizer](https://cfviz.netlify.app/)

[Codeforces Solve Tracker](https://tom0727.github.io/cf-problems/)

[Another Codeforces Solve Tracker](https://cftracker.netlify.app/contests)

[AtCoder Problems](https://kenkoooo.com/atcoder/#/table/)

[AtCoder Companions](https://atcoder-companions.kakira.dev/)

[AtCoder-Codeforces Rating converter](https://silverfoxxxy.github.io/rating-converter)

### Rating and Difficulties

[Open Codeforces Rating System](https://codeforces.com/blog/entry/20762)

[How to Interpret Contest Ratings](https://codeforces.com/blog/entry/68288)

[Codeforces: Problem Difficulties](https://codeforces.com/blog/entry/62865)

[Elo rating system](https://en.wikipedia.org/wiki/Elo_rating_system#Theory)

### Stay Healthy

[Exercises!](https://musclewiki.org/)
