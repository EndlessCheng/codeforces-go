// ==UserScript==
// @name         力扣清爽版
// @namespace    http://tampermonkey.net/
// @version      0.1
// @author       endlesscheng
// @match        https://leetcode.cn/problems/*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=leetcode.cn
// @grant        GM_addStyle
// @run-at       document-start
// ==/UserScript==

;(function () {
    'use strict'
    GM_addStyle(`
/* 笔记 */
#editor > .text-label-r,
/* 返回旧版 */
#editor > div.absolute.right-\\[25px\\].bottom-\\[84px\\].z-overlay,
/* 反馈 */
#editor > div.absolute.right-\\[25px\\].bottom-\\[20px\\].z-overlay,
#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > div.px-5.pt-3,
#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > div.px-5.py-3.pt-\\[38px\\],
#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > div.px-5.py-3:nth-of-type(5),
#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > div.px-5.py-3:nth-of-type(6),
#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > div.px-5.py-3:nth-of-type(7),
#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > div.px-5.py-3:nth-of-type(8),
#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > hr.border-divider-3,
#qd-content > div.h-full.flex-col.ssg__qd-splitter-primary-w > div > div > div > div.flex.h-full.w-full.overflow-y-auto > div > div > div.mt-auto.px-5.pt-8.pb-2\\.5 {
  display: none !important;
}

/* 答题区域 */
#__next > div > div > div > div:nth-of-type(1) {
  height: 100vh !important;
  width: 100vw !important;
  position: fixed !important;
  top: 0 !important;
  left: 0 !important;
}

/* 导航区域 */
#__next .h-full > nav {
  z-index: 0 !important;
  display: none !important;
}
`)
})()
