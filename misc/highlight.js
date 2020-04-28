// ==UserScript==
// @name         Codeforces Highlight
// @namespace    EndlessCheng
// @version      0.1
// @description  highlight some important words
// @author       EndlessCheng
// @match        https://atcoder.jp/*
// @match        https://codeforces.com/*
// @match        https://codeforces.ml/*
// @match        https://codingcompetitions.withgoogle.com/*
// @match        https://leetcode-cn.com/*
// ==/UserScript==

(function () {
    'use strict';

    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions#Escaping
    function escapeRegExp(s) {
        return s.replace(/[.*+\-?^${}()|[\]\\]/g, '\\$&'); // $& means the whole matched string
    }

    String.prototype.replaceAll = function (oldStr, newStr) {
        return this.replace(new RegExp(escapeRegExp(oldStr), "g"), newStr);
    };

    const color = "#f25e6b";
    const words = [
        // 注意前者要包含后者
        " not ", "don't", "didn't", "doesn't", "can't", "n't", " no ",
        " and all", " or all",
        " and ", " or ",
        " any", " all ", "every", "both ",
        "exactly", "always",
        "unique", "distinct",
        "must", "only",
        "same", "different",
        "more",

        // "Note", "note",
        "minimize", "maximize", "minimum", "maximum", "minimal", "maximal", "smallest", "largest",
        " small ", " big ",
        "at least", "at most",
        "non-zero", "positive", "integers", "an integer", "integer", "pairwise",
        "permutations", "permutation",
        "lowercase", "uppercase",
        "lexicographically", "lexicographical",
        "expected value",

        "Initially", "initially", "guaranteed",
        "modulo",

        "operations", "Operations", "operation", "Operation",

        // 高亮一些术语
        "子数组", "子序列", "子字符串", "子串",
        "升序", "降序",
        "返回",
    ];

    const tags = ['p', 'li'];
    for (let ti = 0; ti < tags.length; ti++) {
        let pNodes = document.getElementsByTagName(tags[ti]);
        for (let i = 0; i < pNodes.length; i++) {
            let text = pNodes[i].innerHTML;

            for (let j = 0; j < words.length; j++) {
                text = text.replaceAll(words[j], "<span style='color: " + color + "'>" + words[j] + "</span>");
            }

            text = text.replaceAll("Mr. ", "Mr.")
                .replaceAll("mr. ", "mr.")
                .replaceAll("i.e. ", "i.e.")
                .replaceAll("i. e. ", "i.e.")
                .replaceAll("...", "⋯") // 特殊处理一些句点+空格，这些是不需要换行处理的
                // So you decided to hold a contest on Codeforces.
                // The maximum size of an array is $$$k$$$.
                .replaceAll(". ", ".</p><p>") // 加个换行
                .replaceAll("\\dots", "~.~.~.~") // 替换掉省略号
                .replaceAll("\\ldots", "~.~.~.~"); // 替换掉省略号

            // .replace(/(\$\$\$.+?\$\$\$)/g, "‘$1’"); // 教训：不应该加这个，看似优化实则是帮倒忙

            pNodes[i].innerHTML = text;
        }
    }
})();
