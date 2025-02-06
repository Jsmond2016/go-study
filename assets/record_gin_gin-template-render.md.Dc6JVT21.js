import{_ as i,c as a,o as n,ae as t}from"./chunks/framework.9_tlr4WW.js";const g=JSON.parse('{"title":"","description":"","frontmatter":{},"headers":[],"relativePath":"record/gin/gin-template-render.md","filePath":"record/gin/gin-template-render.md"}'),p={name:"record/gin/gin-template-render.md"};function l(h,s,e,k,E,d){return n(),a("div",null,s[0]||(s[0]=[t(`<h2 id="gin-template-render-模板渲染" tabindex="-1">gin-template-render 模板渲染 <a class="header-anchor" href="#gin-template-render-模板渲染" aria-label="Permalink to &quot;gin-template-render 模板渲染&quot;">​</a></h2><p>知识要点：</p><p>操作步骤：</p><ul><li><p>本地新建 <code>templates</code> 目录，添加以下文件</p><ul><li><code>templates/index.tmpl</code> 代码为：</li></ul><div class="language-tmpl vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">tmpl</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span>&lt;html&gt;</span></span>
<span class="line"><span>    &lt;h1&gt;</span></span>
<span class="line"><span>       Page-1 {{ .title }}</span></span>
<span class="line"><span>    &lt;/h1&gt;</span></span>
<span class="line"><span>&lt;/html&gt;</span></span></code></pre></div><ul><li><code>templates/posts/index.tmpl</code></li></ul><div class="language-tmpl vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">tmpl</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span>{{ define &quot;posts/index.tmpl&quot; }}</span></span>
<span class="line"><span>&lt;html&gt;</span></span>
<span class="line"><span>    &lt;h1&gt;</span></span>
<span class="line"><span>       posts {{ .title }}</span></span>
<span class="line"><span>    &lt;/h1&gt;</span></span>
<span class="line"><span>&lt;/html&gt;</span></span>
<span class="line"><span>{{ end }}</span></span></code></pre></div><ul><li><code>templates/user/index.tmpl</code></li></ul><div class="language-tmpl vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">tmpl</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span>{{ define &quot;user/index.tmpl&quot; }}</span></span>
<span class="line"><span>&lt;html&gt;</span></span>
<span class="line"><span>    &lt;h1&gt;</span></span>
<span class="line"><span>       user {{ .title }}</span></span>
<span class="line"><span>    &lt;/h1&gt;</span></span>
<span class="line"><span>&lt;/html&gt;</span></span>
<span class="line"><span>{{ end }}</span></span></code></pre></div></li></ul><p>代码：</p><div class="language-go vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">go</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">package</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;"> main</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">import</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> (</span></span>
<span class="line"><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">	&quot;</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">net/http</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">	&quot;</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">github.com/gin-gonic/gin</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">func</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;"> main</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">() {</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">	router </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">:=</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> gin.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Default</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">()</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	//--------- 单个 index 模板 -----------------</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// router.LoadHTMLGlob(&quot;templates&quot;)</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// router.GET(&quot;/index&quot;, func(c *gin.Context) {</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// 	//根据完整文件名渲染模板，并传递参数</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// 	c.HTML(http.StatusOK, &quot;index.tmpl&quot;, gin.H{</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// 		&quot;title&quot;: &quot;Hello, World!&quot;,</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// 	})</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// })</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// -------- 多个 index 模板 -----------------</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// 多个 index 模板，使用 * 匹配多个层级，index 同名也没关系</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">	router.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">LoadHTMLGlob</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;templates/**/*&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">	router.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">GET</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;/posts/index&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">func</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#E36209;--shiki-dark:#FFAB70;">c</span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;"> *</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">gin</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Context</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">) {</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">		//根据完整文件名渲染模板，并传递参数</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">		c.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">HTML</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(http.StatusOK, </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;posts/index.tmpl&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">gin</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">H</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">{</span></span>
<span class="line"><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">			&quot;title&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;Hello, Posts!&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">		})</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">	})</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">	router.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">GET</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;/user/index&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">func</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#E36209;--shiki-dark:#FFAB70;">c</span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;"> *</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">gin</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Context</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">) {</span></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">		//根据完整文件名渲染模板，并传递参数</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">		c.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">HTML</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(http.StatusOK, </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;user/index.tmpl&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">, </span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">gin</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">H</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">{</span></span>
<span class="line"><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">			&quot;title&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">: </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;Hello, User!&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">,</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">		})</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">	})</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">	// Listen and serve on 0.0.0.0:8080</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">	router.</span><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">Run</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">(</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">&quot;:8080&quot;</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">)</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">}</span></span></code></pre></div><p>测试：</p><div class="language-http vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">http</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">@</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">base = </span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;">127.0.0.1:8080</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">### c.json 响应 gin.H 方式返回</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">GET</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> http://{{base}}/json-gin </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">HTTP</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">/</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">1.1</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">### c.json 响应  自定义对象的方式返回</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">GET</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> http://{{base}}/json </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">HTTP</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">/</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">1.1</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">### c.xml 响应 的方式返回</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">GET</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> http://{{base}}/json-xml </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">HTTP</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">/</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">1.1</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">### c.yml 响应 的方式返回</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">GET</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> http://{{base}}/json-yml </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">HTTP</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">/</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">1.1</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#6A737D;--shiki-dark:#6A737D;">### c.protobuf 响应 的方式返回</span></span>
<span class="line"><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">GET</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;"> http://{{base}}/protobuf </span><span style="--shiki-light:#D73A49;--shiki-dark:#F97583;">HTTP</span><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">/</span><span style="--shiki-light:#005CC5;--shiki-dark:#79B8FF;">1.1</span></span></code></pre></div>`,8)]))}const c=i(p,[["render",l]]);export{g as __pageData,c as default};
