import{_ as n,o as s,c as a,e as t}from"./app.286a3ac7.js";const e={},i=t(`<h2 id="gin-template-render-\u6A21\u677F\u6E32\u67D3" tabindex="-1"><a class="header-anchor" href="#gin-template-render-\u6A21\u677F\u6E32\u67D3" aria-hidden="true">#</a> gin-template-render \u6A21\u677F\u6E32\u67D3</h2><p>\u77E5\u8BC6\u8981\u70B9\uFF1A</p><p>\u64CD\u4F5C\u6B65\u9AA4\uFF1A</p><ul><li><p>\u672C\u5730\u65B0\u5EFA <code>templates</code> \u76EE\u5F55\uFF0C\u6DFB\u52A0\u4EE5\u4E0B\u6587\u4EF6</p><ul><li><code>templates/index.tmpl</code> \u4EE3\u7801\u4E3A\uFF1A</li></ul><div class="language-tmpl ext-tmpl line-numbers-mode"><pre class="language-tmpl"><code>&lt;html&gt;
    &lt;h1&gt;
       Page-1 {{ .title }}
    &lt;/h1&gt;
&lt;/html&gt;
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><ul><li><code>templates/posts/index.tmpl</code></li></ul><div class="language-tmpl ext-tmpl line-numbers-mode"><pre class="language-tmpl"><code>{{ define &quot;posts/index.tmpl&quot; }}
&lt;html&gt;
    &lt;h1&gt;
       posts {{ .title }}
    &lt;/h1&gt;
&lt;/html&gt;
{{ end }}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><ul><li><code>templates/user/index.tmpl</code></li></ul><div class="language-tmpl ext-tmpl line-numbers-mode"><pre class="language-tmpl"><code>{{ define &quot;user/index.tmpl&quot; }}
&lt;html&gt;
    &lt;h1&gt;
       user {{ .title }}
    &lt;/h1&gt;
&lt;/html&gt;
{{ end }}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div></li></ul><p>\u4EE3\u7801\uFF1A</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;net/http&quot;</span>

	<span class="token string">&quot;github.com/gin-gonic/gin&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	router <span class="token operator">:=</span> gin<span class="token punctuation">.</span><span class="token function">Default</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	<span class="token comment">//--------- \u5355\u4E2A index \u6A21\u677F -----------------</span>
	<span class="token comment">// router.LoadHTMLGlob(&quot;templates&quot;)</span>
	<span class="token comment">// router.GET(&quot;/index&quot;, func(c *gin.Context) {</span>
	<span class="token comment">// 	//\u6839\u636E\u5B8C\u6574\u6587\u4EF6\u540D\u6E32\u67D3\u6A21\u677F\uFF0C\u5E76\u4F20\u9012\u53C2\u6570</span>
	<span class="token comment">// 	c.HTML(http.StatusOK, &quot;index.tmpl&quot;, gin.H{</span>
	<span class="token comment">// 		&quot;title&quot;: &quot;Hello, World!&quot;,</span>
	<span class="token comment">// 	})</span>
	<span class="token comment">// })</span>

	<span class="token comment">// -------- \u591A\u4E2A index \u6A21\u677F -----------------</span>
	<span class="token comment">// \u591A\u4E2A index \u6A21\u677F\uFF0C\u4F7F\u7528 * \u5339\u914D\u591A\u4E2A\u5C42\u7EA7\uFF0Cindex \u540C\u540D\u4E5F\u6CA1\u5173\u7CFB</span>
	router<span class="token punctuation">.</span><span class="token function">LoadHTMLGlob</span><span class="token punctuation">(</span><span class="token string">&quot;templates/**/*&quot;</span><span class="token punctuation">)</span>

	router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/posts/index&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token comment">//\u6839\u636E\u5B8C\u6574\u6587\u4EF6\u540D\u6E32\u67D3\u6A21\u677F\uFF0C\u5E76\u4F20\u9012\u53C2\u6570</span>
		c<span class="token punctuation">.</span><span class="token function">HTML</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;posts/index.tmpl&quot;</span><span class="token punctuation">,</span> gin<span class="token punctuation">.</span>H<span class="token punctuation">{</span>
			<span class="token string">&quot;title&quot;</span><span class="token punctuation">:</span> <span class="token string">&quot;Hello, Posts!&quot;</span><span class="token punctuation">,</span>
		<span class="token punctuation">}</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>

	router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/user/index&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token comment">//\u6839\u636E\u5B8C\u6574\u6587\u4EF6\u540D\u6E32\u67D3\u6A21\u677F\uFF0C\u5E76\u4F20\u9012\u53C2\u6570</span>
		c<span class="token punctuation">.</span><span class="token function">HTML</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;user/index.tmpl&quot;</span><span class="token punctuation">,</span> gin<span class="token punctuation">.</span>H<span class="token punctuation">{</span>
			<span class="token string">&quot;title&quot;</span><span class="token punctuation">:</span> <span class="token string">&quot;Hello, User!&quot;</span><span class="token punctuation">,</span>
		<span class="token punctuation">}</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>

	<span class="token comment">// Listen and serve on 0.0.0.0:8080</span>
	router<span class="token punctuation">.</span><span class="token function">Run</span><span class="token punctuation">(</span><span class="token string">&quot;:8080&quot;</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>\u6D4B\u8BD5\uFF1A</p><div class="language-http ext-http line-numbers-mode"><pre class="language-http"><code>
@base = 127.0.0.1:8080

### c.json \u54CD\u5E94 gin.H \u65B9\u5F0F\u8FD4\u56DE
<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/json-gin</span> <span class="token http-version property">HTTP/1.1</span></span>

### c.json \u54CD\u5E94  \u81EA\u5B9A\u4E49\u5BF9\u8C61\u7684\u65B9\u5F0F\u8FD4\u56DE
<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/json</span> <span class="token http-version property">HTTP/1.1</span></span>

### c.xml \u54CD\u5E94 \u7684\u65B9\u5F0F\u8FD4\u56DE
<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/json-xml</span> <span class="token http-version property">HTTP/1.1</span></span>

### c.yml \u54CD\u5E94 \u7684\u65B9\u5F0F\u8FD4\u56DE
<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/json-yml</span> <span class="token http-version property">HTTP/1.1</span></span>

### c.protobuf \u54CD\u5E94 \u7684\u65B9\u5F0F\u8FD4\u56DE
<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/protobuf</span> <span class="token http-version property">HTTP/1.1</span></span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,8),l=[i];function p(o,c){return s(),a("div",null,l)}var d=n(e,[["render",p],["__file","gin-template-render.html.vue"]]);export{d as default};
