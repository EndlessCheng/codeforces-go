# 模板仓库

由于算法知识点非常多且杂，将自己学习到的算法**亲手**整理成一个小型 Wiki 是有必要的。

一个算法模板应当涵盖下面几点：
- 对该算法的基本介绍（核心思想、复杂度等）
- 参考链接或书籍章节（讲的比较好的资料）
- 模板代码（可以包含一些注释）
- 使用模板的注意事项
- 模板补充内容（常见题型中的额外代码、建模技巧等）
- 相关题目（模板题、经典题、建模转换题等）

注：为方便在比赛中使用，一部分函数写成了 lambda 的形式。

## 目录

- 数据结构
  - [双端队列 deque.go](deque.go)
  - [堆（优先队列）heap.go](heap.go)
  - [单调栈、单调队列 common.go 中的 monotoneCollection](common.go)
  - [并查集 union_find.go](union_find.go)
  - [ST 表 common.go 中的 stInit](common.go)
  - [树状数组 fenwick_tree.go](fenwick_tree.go)
  - [线段树 segment_tree.go](segment_tree.go)
  - [左偏树（可并堆）leftist_tree.go](leftist_tree.go)
  - [笛卡尔树 cartesian_tree.go](cartesian_tree.go)
  - [二叉搜索树公共方法 bst.go](bst.go)
  - [Treap treap.go](treap.go)
  - [伸展树 splay.go](splay.go)
  - [动态树 LCT link_cut_tree.go](link_cut_tree.go)
  - [红黑树 red_black_tree.go](red_black_tree.go)
  - [替罪羊树 scapegoat_tree.go](scapegoat_tree.go)
  - [k-d 树 kd_tree.go](kd_tree.go)
  - 珂朵莉树（ODT）
    - [数组版 odt.go](odt.go)
    - [平衡树版 odt_bst.go](odt_bst.go)
- [字符串 strings.go](strings.go)
  - Hash
  - KMP
  - 扩展 KMP（Z algorithm）
  - 最小表示法
  - Manacher
  - AC 自动机
  - 后缀数组（SA）
  - 字典树（trie）
- 数学
  - [数论 math.go](math.go)
    - 最大公因数（GCD）
    - 类欧几里得算法 ∑⌊(ai+b)/m⌋
    - Pollard-Rho
    - 线性筛
    - 欧拉函数
    - 原根
    - 扩展 GCD
    - 逆元
    - 中国剩余定理（CRT）
    - 扩展中国剩余定理（EXCRT）
    - 离散对数
      - 小步大步算法（BSGS）
      - 扩展小步大步算法（exBSGS）
      - 二次剩余
      - Jacobi 符号
    - 莫比乌斯函数
    - 数论分块
    - 杜教筛
  - [FFT math_fft.go](math_fft.go)
  - [NTT math_ntt.go](math_ntt.go)
    - 包含多项式全家桶（求逆、开方等等）
  - [矩阵、高斯消元、线性基 math_matrix.go](math_matrix.go)
  - [连分数、佩尔方程 math_continued_fraction.go](math_continued_fraction.go)
  - [计算几何 geometry.go](geometry.go)
    - 直线和点
    - 直线和直线
    - 直线和圆
    - 圆和圆
    - 凸包
    - 最近点对
    - 最远点对
  - [博弈论 math.go](math.go)
    - SG 函数
- [动态规划 dp.go](dp.go)
  - 背包
  - 线性 DP
  - 区间 DP
  - 状压 DP
  - 数位 DP
  - 树形 DP
    - 换根 DP
- [图论 graph.go](graph.go)
  - 欧拉回路
  - 割点
  - 割边（桥）
  - 双连通分量
  - 最短路
  - 最小生成树
  - 最小差值生成树
  - 二分图最大匹配
  - 带权二分图最大完美匹配
  - 拓扑排序
  - 极大强连通分量
  - 2-SAT
  - 基环树
  - 最大流
  - 最小费用最大流
  - [树上问题 graph_tree.go](graph_tree.go)
    - 直径
    - 重心
    - 最近公共祖先（LCA）
      - 树上差分
    - 树链剖分（重链剖分，HLD）
    - 树上启发式合并（DSU）
    - 树分块
- [位运算 bits.go](bits.go)
  - bitset
- 其他
  - [三分查找 sort.go](sort.go)
  - [0-1 分数规划 sort.go](sort.go)
  - [莫队算法 common.go 中的 moAlgorithm](common.go)
     - 带修莫队
     - 回滚莫队
     - 树上莫队
- [快读模板 io.go](io.go)


## 代码生成、测试及对拍

代码生成见 [template](./template) 中的 `generator.go` 以及各个文件夹下的 `generator.go`

代码测试及对拍见 [testutil.go](/main/testutil/testutil.go)

编写一个 `run(io.Reader, io.Writer)` 函数来处理输入输出。这样写的理由是：

- 在 `main()` 中调用 `run(os.Stdin, os.Stdout)` 来执行代码；
- 测试时，将测试数据转换成 `strings.Reader` 当作输入，并用一个 `strings.Builder` 来接收输出，将这二者传入 `run` 中，然后就能比较输出与答案了。
- 对拍时需要实现一个暴力算法 `runAC`，参数和 `run` 一样。通过[随机数据生成器](/main/testutil/rand.go)来生成数据，分别传入 `runAC` 和 `run`，通过比对各自的输出，来检查 `run` 中的问题。

[main](/main) 中的所有非交互题的 CF 代码及其对应测试全部按照上述框架实现。

## 学习资料及题目

注：由于入门经典上选了很多区域赛的题，一部分题目可以在 GYM 上找到，这样可以就可以用 Go 编程提交了。

[算法竞赛入门经典（第二版）](https://github.com/aoapc-book/aoapc-bac2nd)

[算法竞赛入门经典训练指南](https://github.com/klb3713/aoapc-book/tree/master/TrainingGuide/bookcodes)

[算法竞赛入门经典训练指南（升级版）](https://gitee.com/sukhoeing/aoapc-training-guide2)

[算法竞赛进阶指南](https://github.com/lydrainbowcat/tedukuri)

[算法竞赛入门到进阶](https://github.com/luoyongjun999/code)

[OI Public Library（含国家队论文）](https://github.com/enkerewpo/OI-Public-Library)

[算法竞赛 (ICPC, OI, etc) 论文，课件，文档，笔记等](https://github.com/LzyRapx/Competitive-Programming-Docs)

[算法竞赛课件分享 by hzwer](https://github.com/hzwer/shareOI)

[算法第四版 Java 源码](https://algs4.cs.princeton.edu/code/)

[数据结构和算法动态可视化](https://visualgo.net/zh)

[OI Wiki](https://oi-wiki.org/)

[CP-Algorithms](https://cp-algorithms.com/)

[All the good tutorials found for Competitive Programming](https://codeforces.com/blog/entry/57282)

[GeeksforGeeks 上的算法合集](https://www.geeksforgeeks.org/how-to-prepare-for-acm-icpc/)

[Pepcy 模板](http://pepcy.cf/icpc-templates/)

[F0RE1GNERS 模板](https://github.com/F0RE1GNERS/template)

[算法学习笔记（目录）](https://zhuanlan.zhihu.com/p/105467597)

[洛谷模板题（建议按难度筛选）](https://www.luogu.com.cn/problem/list?keyword=%E6%A8%A1%E6%9D%BF&page=1)

[Codeforces Problem Topics](https://codeforces.com/blog/entry/55274)

[Luogu Problem List](https://github.com/SFOI-Team/luogu-problem-list/blob/master/list.md)

### 洛谷日报

[2021 年洛谷日报索引](https://www.luogu.com.cn/discuss/show/287888)

[2020 年洛谷日报索引](https://www.luogu.com.cn/discuss/show/179788)

[2019 年洛谷日报索引](https://www.luogu.com.cn/discuss/show/92685)

[2018 年洛谷日报索引](https://www.luogu.com.cn/discuss/show/48491)

### 高级竞赛算法

[算法进阶课](https://www.acwing.com/activity/content/32/)

### AtCoder 版《挑战程序设计竞赛》

[AtCoder 版！蟻本 (初級編)](https://qiita.com/drken/items/e77685614f3c6bf86f44)

[AtCoder 版！蟻本 (中級編)](https://qiita.com/drken/items/2f56925972c1d34e05d8)

[AtCoder 版！蟻本 (上級編)](https://qiita.com/drken/items/9b311d553aa434bb26e4)

[AtCoder 版！蟻本 (発展的トピック編)](https://qiita.com/drken/items/0de3d205690d92307b7c)

### 待整理

[偶然在 GitHub 上发现的超长列表](https://github.com/dhs347/Dream/blob/master/%E8%AE%A1%E5%88%92/%E8%AE%A1%E5%88%92%E4%B9%A6/A%E8%AE%A1%E5%88%92_%E9%98%B6%E6%AE%B51.md)

[算法竞赛训练中较难的部分](https://blog.csdn.net/skywalkert/article/details/48924861)

[算法竞赛中可能不太会遇到的论文题](https://blog.csdn.net/skywalkert/article/details/48878925)

https://blog.csdn.net/calabash_boy/article/details/79973483

https://github.com/zimpha/algorithmic-library

https://www.luogu.com.cn/blog/command-block/blog-suo-yin-zhi-ding-post

https://wcysai.github.io/
