import{_ as i,c as a,o as n,ae as h}from"./chunks/framework.DXW4JJKH.js";const g=JSON.parse('{"title":"golang 中，数字操作常用方法","description":"","frontmatter":{},"headers":[],"relativePath":"record/base/go-number.md","filePath":"record/base/go-number.md"}'),k={name:"record/base/go-number.md"};function l(t,s,p,e,E,r){return n(),a("div",null,s[0]||(s[0]=[h(`<h1 id="golang-中-数字操作常用方法" tabindex="-1">golang 中，数字操作常用方法 <a class="header-anchor" href="#golang-中-数字操作常用方法" aria-label="Permalink to &quot;golang 中，数字操作常用方法&quot;">​</a></h1><p>我来介绍一下 Go 语言中常用的数字操作方法。我将创建一个新的文档来详细说明。</p><h2 id="一、基本数字类型" tabindex="-1">一、基本数字类型 <a class="header-anchor" href="#一、基本数字类型" aria-label="Permalink to &quot;一、基本数字类型&quot;">​</a></h2><div class="language-go vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 整数类型</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">var</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> i8 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">int8</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">    // -128 到 127</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">var</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> i16 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">int16</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">  // -32768 到 32767</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">var</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> i32 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">int32</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">  // -2^31 到 2^31-1</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">var</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> i64 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">int64</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">  // -2^63 到 2^63-1</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">var</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> ui </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">uint</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">    // 无符号整数</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 浮点类型</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">var</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> f32 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">float32</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">  // 32位浮点数</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">var</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> f64 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">float64</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">  // 64位浮点数</span></span></code></pre></div><h2 id="二、数字转换" tabindex="-1">二、数字转换 <a class="header-anchor" href="#二、数字转换" aria-label="Permalink to &quot;二、数字转换&quot;">​</a></h2><div class="language-go vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 1. 类型转换</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">intNum </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;"> 42</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">float64Num </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;"> float64</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(intNum)    </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// int 转 float64</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">uint32Num </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;"> uint32</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(intNum)      </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// int 转 uint32</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 2. 字符串与数字转换</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">import</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> &quot;</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">strconv</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 字符串转数字</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">intVal, _ </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> strconv.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Atoi</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;42&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)           </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 字符串转 int</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">int64Val, _ </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> strconv.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">ParseInt</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;42&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">10</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">64</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)  </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 字符串转 int64</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">floatVal, _ </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> strconv.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">ParseFloat</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;3.14&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">64</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)  </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 字符串转 float64</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 数字转字符串</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">strVal </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> strconv.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Itoa</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">42</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)                </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// int 转字符串</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">str64 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> strconv.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">FormatInt</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">42</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">10</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)        </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// int64 转字符串</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">strFloat </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> strconv.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">FormatFloat</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">3.14</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&#39;</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">f</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&#39;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">2</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">64</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)  </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// float64 转字符串</span></span></code></pre></div><h2 id="三、数学运算" tabindex="-1">三、数学运算 <a class="header-anchor" href="#三、数学运算" aria-label="Permalink to &quot;三、数学运算&quot;">​</a></h2><div class="language-go vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">import</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> &quot;</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">math</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 1. 基本运算</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">max </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Max</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">1.0</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">2.0</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)    </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 取最大值</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">min </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Min</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">1.0</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">2.0</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)    </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 取最小值</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">abs </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Abs</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">-</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">3.14</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)       </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 绝对值</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">pow </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Pow</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">2</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">3</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)        </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 幂运算</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">sqrt </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Sqrt</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">16</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)        </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 平方根</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 2. 取整操作</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">ceil </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Ceil</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">3.14</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)      </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 向上取整</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">floor </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Floor</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">3.14</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)    </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 向下取整</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">round </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Round</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">3.14</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)    </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 四舍五入</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 3. 三角函数</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">sin </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Sin</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(math.Pi </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">/</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;"> 2</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)  </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 正弦</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">cos </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Cos</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">0</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)            </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 余弦</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">tan </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Tan</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(math.Pi </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">/</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;"> 4</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)  </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 正切</span></span></code></pre></div><h2 id="四、随机数生成" tabindex="-1">四、随机数生成 <a class="header-anchor" href="#四、随机数生成" aria-label="Permalink to &quot;四、随机数生成&quot;">​</a></h2><div class="language-go vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">import</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> (</span></span>
<span class="line"><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">    &quot;</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">math/rand</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;</span></span>
<span class="line"><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">    &quot;</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">time</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 1. 基本随机数</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">rand.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Seed</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(time.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Now</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">().</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">UnixNano</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">())  </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 设置随机种子</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">n </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> rand.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Int</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">()                   </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 生成随机整数</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">f </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> rand.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Float64</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">()               </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 生成随机浮点数</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 2. 范围随机数</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">randInt </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> rand.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Intn</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">100</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)         </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 生成 [0,100) 的随机整数</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">randFloat </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> rand.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Float64</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">() </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">*</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;"> 100</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">  // 生成 [0,100) 的随机浮点数</span></span></code></pre></div><h2 id="五、位运算" tabindex="-1">五、位运算 <a class="header-anchor" href="#五、位运算" aria-label="Permalink to &quot;五、位运算&quot;">​</a></h2><div class="language-go vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">a, b </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;"> 60</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">13</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 常用位运算</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">and </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> a </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">&amp;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> b     </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 按位与</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">or </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> a </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">|</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> b      </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 按位或</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">xor </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> a </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">^</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> b     </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 按位异或</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">not </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;"> ^</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">a        </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 按位取反</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">left </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> a </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">&lt;&lt;</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;"> 2</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">   // 左移</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">right </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> a </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">&gt;&gt;</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;"> 2</span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">  // 右移</span></span></code></pre></div><h2 id="六、数字判断" tabindex="-1">六、数字判断 <a class="header-anchor" href="#六、数字判断" aria-label="Permalink to &quot;六、数字判断&quot;">​</a></h2><div class="language-go vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">import</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> &quot;</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">math</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 1. 特殊值判断</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">isNaN </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">IsNaN</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(f)       </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 是否是 NaN</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">isInf </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">IsInf</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(f, </span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">0</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)    </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 是否是无穷大</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 2. 符号判断</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">isPositive </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Signbit</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(f)  </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 判断符号位</span></span></code></pre></div><h2 id="七、实用技巧" tabindex="-1">七、实用技巧 <a class="header-anchor" href="#七、实用技巧" aria-label="Permalink to &quot;七、实用技巧&quot;">​</a></h2><ol><li><strong>精度处理</strong>：</li></ol><div class="language-go vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 保留两位小数</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">rounded </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Round</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">3.1415926</span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">*</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">100</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">) </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">/</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;"> 100</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 金额计算（使用 decimal 包）</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">import</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> &quot;</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">github.com/shopspring/decimal</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">d1 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> decimal.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">NewFromFloat</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">2.1</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">d2 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> decimal.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">NewFromFloat</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">1.2</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">sum </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> d1.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Add</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(d2)  </span><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 精确计算</span></span></code></pre></div><ol start="2"><li><strong>溢出检查</strong>：</li></ol><div class="language-go vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">// 检查加法溢出</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">if</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> a </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">&gt;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> math.MaxInt64 </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">-</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> b {</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">    // 处理溢出情况</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">}</span></span></code></pre></div><h2 id="注意事项" tabindex="-1">注意事项 <a class="header-anchor" href="#注意事项" aria-label="Permalink to &quot;注意事项&quot;">​</a></h2><ol><li>浮点数计算可能存在精度问题，需要特别注意</li><li>金融计算建议使用 decimal 包</li><li>进行类型转换时要注意可能的数据溢出</li><li>随机数生成前必须设置随机种子</li><li>位运算要注意符号位的影响</li></ol><div class="language- vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang"></span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span></span></span>
<span class="line"><span>这个文档涵盖了 Go 语言中最常用的数字操作方法。建议根据实际场景选择合适的方法，特别是在需要高精度计算的场景下。</span></span></code></pre></div>`,22)]))}const y=i(k,[["render",l]]);export{g as __pageData,y as default};
