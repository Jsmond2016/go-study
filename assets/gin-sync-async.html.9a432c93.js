import{_ as n,o as s,c as a,e as t}from"./app.286a3ac7.js";const p={},e=t(`<h2 id="gin-sync-async-\u5F02\u6B65\u5904\u7406" tabindex="-1"><a class="header-anchor" href="#gin-sync-async-\u5F02\u6B65\u5904\u7406" aria-hidden="true">#</a> gin-sync-async \u5F02\u6B65\u5904\u7406</h2><blockquote><p><code>gin-demo\\chapter-04\\exercise-4\\main.go</code></p></blockquote><p>\u77E5\u8BC6\u8981\u70B9\uFF1A</p><p>\u5904\u7406\u4E00\u4E9B\u8017\u65F6\u4EFB\u52A1\u65F6\uFF0C\u4F7F\u7528\u5F02\u6B65\u64CD\u4F5C</p><p>\u4EE3\u7801\uFF1A</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;log&quot;</span>
	<span class="token string">&quot;net/http&quot;</span>
	<span class="token string">&quot;time&quot;</span>

	<span class="token string">&quot;github.com/gin-gonic/gin&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	router <span class="token operator">:=</span> gin<span class="token punctuation">.</span><span class="token function">Default</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token comment">//1. \u5F02\u6B65\uFF0C\u7C7B\u4F3C setTimeout\uFF0C\u5148\u8FD4\u56DE\u7ED3\u679C\uFF0C go func \u8BED\u53E5\u540E\u53F0\u7EE7\u7EED\u6267\u884C</span>
	router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/upload-async&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token comment">// goroutine \u4E2D\u53EA\u80FD\u4F7F\u7528\u53EA\u8BFB\u7684\u4E0A\u4E0B\u6587 c.Copy()</span>
		cCp <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">Copy</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
		<span class="token keyword">go</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
			time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span><span class="token number">5</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">)</span>
			<span class="token comment">// \u6CE8\u610F\u4F7F\u7528\u53EA\u8BFB\u4E0A\u4E0B\u6587</span>
			log<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">&quot;Done! in path &quot;</span> <span class="token operator">+</span> cCp<span class="token punctuation">.</span>Request<span class="token punctuation">.</span>URL<span class="token punctuation">.</span>Path<span class="token punctuation">)</span>
		<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
		c<span class="token punctuation">.</span><span class="token function">JSON</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;\u5F02\u6B65\u8BF7\u6C42\u6210\u529F\uFF0C\u8BF7\u7A0D\u540E\u5C1D\u8BD5&quot;</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>
	<span class="token comment">//2. \u540C\u6B65\uFF0C\u8981\u7B49\u5F85 \u5904\u7406\u65F6\u95F4\uFF0C\u5904\u7406\u5B8C\u6210\u540E\u8FD4\u56DE</span>
	router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/sync&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		time<span class="token punctuation">.</span><span class="token function">Sleep</span><span class="token punctuation">(</span><span class="token number">5</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">)</span>
		<span class="token comment">// \u6CE8\u610F\u53EF\u4EE5\u4F7F\u7528\u539F\u59CB\u4E0A\u4E0B\u6587</span>
		log<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">&quot;Done! in path &quot;</span> <span class="token operator">+</span> c<span class="token punctuation">.</span>Request<span class="token punctuation">.</span>URL<span class="token punctuation">.</span>Path<span class="token punctuation">)</span>
		c<span class="token punctuation">.</span><span class="token function">JSON</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;\u540C\u6B65\u5904\u7406\u4E2D\uFF0C\u8017\u65F6 5s &quot;</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>

	<span class="token comment">// Listen and serve on 0.0.0.0:8080</span>
	router<span class="token punctuation">.</span><span class="token function">Run</span><span class="token punctuation">(</span><span class="token string">&quot;:8080&quot;</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>\u6D4B\u8BD5\uFF1A</p><div class="language-http ext-http line-numbers-mode"><pre class="language-http"><code>@base = 127.0.0.1:8080

### 
<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/upload-async</span> <span class="token http-version property">HTTP/1.1</span></span>

###
<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/sync</span> <span class="token http-version property">HTTP/1.1</span></span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,8),o=[e];function c(i,u){return s(),a("div",null,o)}var r=n(p,[["render",c],["__file","gin-sync-async.html.vue"]]);export{r as default};
