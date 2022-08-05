import{_ as n,o as s,c as a,e as t}from"./app.286a3ac7.js";const e={},i=t(`<h2 id="gin-static" tabindex="-1"><a class="header-anchor" href="#gin-static" aria-hidden="true">#</a> gin-static</h2><blockquote><p>\u9759\u6001\u6587\u4EF6\u670D\u52A1\uFF0C<code>gin_demo\\chapter-04\\exercise-3</code></p></blockquote><p>\u77E5\u8BC6\u8981\u70B9\uFF1A</p><p>\u51C6\u5907\uFF1A</p><p>2\u4E2A\u6587\u4EF6\u5939\u548C2\u4E2A\u56FE\u7247</p><ul><li><p><code>assets/img.png</code></p></li><li><p><code>images/img.png</code></p></li></ul><p>\u4EE3\u7801\uFF1A</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;net/http&quot;</span>

	<span class="token string">&quot;github.com/gin-gonic/gin&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	router <span class="token operator">:=</span> gin<span class="token punctuation">.</span><span class="token function">Default</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token comment">// \u4E0B\u9762\u6D4B\u8BD5\u9759\u6001\u6587\u4EF6\u670D\u52A1</span>

	<span class="token comment">// router.StaticFS(&quot;/files&quot;, http.Dir(&quot;../&quot;))</span>

	<span class="token comment">// Static \u63D0\u4F9B\u7ED9\u5B9A\u6587\u4EF6\u7CFB\u7EDF\u6839\u76EE\u5F55\u4E2D\u7684\u6587\u4EF6\u3002</span>
	router<span class="token punctuation">.</span><span class="token function">Static</span><span class="token punctuation">(</span><span class="token string">&quot;/assets&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;./assets&quot;</span><span class="token punctuation">)</span> <span class="token comment">// \u8BF7\u6C42\u65B9\u5F0F\uFF1Ahttp://localhost:8080/assets/img.png</span>
	<span class="token comment">// \u663E\u793A\u5F53\u524D\u6587\u4EF6\u5939\u4E0B\u7684\u6240\u6709\u6587\u4EF6/\u6216\u8005\u6307\u5B9A\u6587\u4EF6</span>
	router<span class="token punctuation">.</span><span class="token function">StaticFS</span><span class="token punctuation">(</span><span class="token string">&quot;/show-dir&quot;</span><span class="token punctuation">,</span> http<span class="token punctuation">.</span><span class="token function">Dir</span><span class="token punctuation">(</span><span class="token string">&quot;.&quot;</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	router<span class="token punctuation">.</span><span class="token function">StaticFile</span><span class="token punctuation">(</span><span class="token string">&quot;/image&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;./images/img.png&quot;</span><span class="token punctuation">)</span>

	router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/redirect&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token comment">//\u652F\u6301\u5185\u90E8\u548C\u5916\u90E8\u7684\u91CD\u5B9A\u5411</span>
		c<span class="token punctuation">.</span><span class="token function">Redirect</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusMovedPermanently<span class="token punctuation">,</span> <span class="token string">&quot;http://www.baidu.com/&quot;</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>

	router<span class="token punctuation">.</span><span class="token function">Run</span><span class="token punctuation">(</span><span class="token string">&quot;:8080&quot;</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>\u6D4B\u8BD5\uFF1A</p><p>\u6D4F\u89C8\u5668\u8BBF\u95EE\uFF1A</p><div class="language-bash ext-sh line-numbers-mode"><pre class="language-bash"><code>http://localhost:8080/assets/img.png


http://localhost:8080/image/img.png
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,11),p=[i];function c(o,u){return s(),a("div",null,p)}var d=n(e,[["render",c],["__file","gin-static.html.vue"]]);export{d as default};
