// ==UserScript==
// @name         高亮关键字 - 其它
// @namespace    https://github.com/EndlessCheng
// @version      0.1
// @description  highlight some important words
// @author       灵茶山艾府
// @match        https://leetcode.cn/contest/*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=leetcode.cn
// ==/UserScript==

(function () {
    'use strict';

    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions#Escaping
    function escapeRegExp(s) {
        return s.replace(/[.*+\-?^${}()|[\]\\]/g, '\\$&'); // $& means the whole matched string
    }

    String.prototype.replaceAll = function (oldSubstr, newSubstr) {
        return this.replace(new RegExp(escapeRegExp(oldSubstr), "g"), newSubstr);
    };

    String.prototype.highlight = function (substr, color) {
        return this.replaceAll(substr, "<span style='color: " + color + "'>" + substr + "</span>");
    };

    // 关键词高亮
    const color = "#49cc54"; // 49cc54  008dde
    // 高亮的词，一部分类似 Python 的关键字，另一部分是一些术语、修饰词之类
    // 注意若两个词有包含关系，把长的放前面
    const words = [
        // 描述
        "Initially", "initially", "guaranteed", "No matter", "no matter",
        "Not ", " not ", "don't", "didn't", "doesn't", "can't", "isn't", "aren't", "No ", " no ", "Non-", "non-", "without", "forbidden", "invalid", " nothing",
        "if and only if", "as long as",
        "Each ", " each ", // each 是个挺有趣的词，高亮它能帮助快速定位后面所描述的对象
        "every", " both ", " other ",
        "Also", "also",
        // " all ", "All ",
        "any number of", "Any number of",
        "equally", "equal", "same", "duplicate", "different", "unique", "distinct", "strictly", "overlapping", "overlap",
        "Only", "only", "just", "Exactly", "exactly", "always", "indeed", "precisely",
        "pairwise", "adjacent", "neighbour", "in a row", "consecutive", "continuous", "contiguous", "one after another", "disjoint", "as possible",
        "more than", "less than", "greater than",
        "except",
        // "must",

        // 求解
        "choose",
        "minimize", "maximize", "minimum", "maximum", "minimal", "maximal", "smallest", "largest", "shortest", "longest", "cheapest", "fastest",
        // " small", " big", " large", " few",
        "At least", "at least", "At most", "at most",

        // 特殊描述
        "substring", "subarray", // "subsequence",
        "Empty", " empty",
        "leading zero", "zero", "positive", "negative",
        "decreasing", "descending", "increasing", "ascending", "sorted",
        "permutation",
        "lowercase", "lower case", "uppercase", "upper case",
        "lexicographical", "palindrome",
        "undirected", "directed", "bidirectional", "direct",
        "independent",
        "expected value",
        " circle", " ring",
        "counterclockwise", "counter-clockwise", "clockwise",
        // "origin",
        "initial",
        "infinite",
        "leaf",
        "even integer",

       
        
        // "没有", "不是", "不同", "不会", "互不",
         "相同",
        // "独一无二", "唯一", "只有", "两个",
        // "所有", /*"每",*/ "任何", "任意", "或", "之一", "反之", /*"必须",*/ "仅", "其他",
        "之和",
        "最小", "最大", "最少", "最多", "最短", "最长", "最早", "最晚", "最高", // todo regex 最...
        // // "最后", "第一", 
        // "至少", "至多", "恰好", "刚好",
        // "非空", "连续",
        // // "子数组", "子区间", "区间", "子序列", "子字符串", "子串",
        // "严格", /*"递增",*/ "递减", "升序", "降序", "字典序",
        // "重复", "重新", "相邻",
        // "小写", "大写", "回文",
        // // "排列",
        // "叶子", "叶节点",
        // "单向", "双向",
        // "本身",
        // "独立",
        // // "返回",
        // // "计算",
        
        
    ];

    const tags = ['p', 'li'];
    for (let ti = 0; ti < tags.length; ti++) {
        let pNodes = document.getElementsByTagName(tags[ti]);
        for (let i = 0; i < pNodes.length; i++) {
            let text = pNodes[i].innerText.trim();
            if (text === "") {
                continue;
            }

            text = pNodes[i].innerHTML;

            for (let j = 0; j < words.length; j++) {
                text = text.highlight(words[j], color);
            }

            // 额外高亮
            const colorRed = "#ff0000";
            text = text
//                .highlight("取模", colorRed)
//                .highlight("取余", colorRed)
                .highlight("重复边", colorRed)
                .highlight("重边", colorRed)
                .highlight("最后", colorRed);
            // .highlight("第一", colorRed);

            text = text.replaceAll("Mr. ", "Mr.")
                .replaceAll("mr. ", "mr.")
                .replaceAll("Mrs. ", "Mrs.")
                .replaceAll("Ms. ", "Ms.")
                .replaceAll("Dr. ", "Dr.")
                .replaceAll("Co. ", "Co.")
                .replaceAll("Ltd. ", "Ltd.")
                .replaceAll("i. e. ", "i.e.")
                .replaceAll("i. e. ", "i.e.") // see https://codeforces.com/contest/1535/problem/A
                .replaceAll("i.e. ", "i.e.")
                .replaceAll("I.e. ", "I.e.")
                .replaceAll("E. g. ", "E.g.") // see https://codeforces.com/contest/1551/problem/E
                .replaceAll("E.g. ", "E.g.")
                .replaceAll(". $", ".$") // 防止数学公式异常
                .replaceAll(". \\", ".\\") // 防止数学公式异常
                .replaceAll("...", "⋯") // 特殊处理一些句点，这些是不需要换行处理的
                // So you decided to hold a contest on Codeforces.
                // The maximum size of an array is $$$k$$$.
                .replaceAll(". ", ".</p><p>") // 加个换行（英文）
                .replaceAll(".\n", ".</p><p>") // 加个换行（英文）
                .replaceAll("。", "。</p><p>") // 加个换行（中文）
            // .replaceAll("\\dots", "~.~.~.~") // 替换掉省略号
            // .replaceAll("\\ldots", "~.~.~.~"); // 替换掉省略号

            // .replace(/(\$\$\$.+?\$\$\$)/g, "‘$1’"); // 教训：不应该加这个，看似优化实则是帮倒忙

            pNodes[i].innerHTML = text;
        }
    }
})();
