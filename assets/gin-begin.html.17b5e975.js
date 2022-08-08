import{_ as n,o as s,c as a,e}from"./app.50b80db3.js";const i={},l=e(`<h1 id="gin-\u5B89\u88C5\u548C\u521D\u59CB\u5316" tabindex="-1"><a class="header-anchor" href="#gin-\u5B89\u88C5\u548C\u521D\u59CB\u5316" aria-hidden="true">#</a> Gin \u5B89\u88C5\u548C\u521D\u59CB\u5316</h1><ul><li>\u521B\u5EFA\u76EE\u5F55\u548C\u521D\u59CB\u5316 go mod</li></ul><div class="language-bash ext-sh line-numbers-mode"><pre class="language-bash"><code><span class="token function">mkdir</span> gin_demo/chapter-01


<span class="token builtin class-name">cd</span> gin_demo


go mod init <span class="token builtin class-name">test</span>


<span class="token builtin class-name">cd</span> chapter-01


<span class="token function">touch</span> main.go
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><ul><li>\u5B89\u88C5 gin</li></ul><div class="language-bash ext-sh line-numbers-mode"><pre class="language-bash"><code>go get -u github.com/gin-gonic/gin
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><ul><li>\u5B89\u88C5 gowatch \u8FDB\u884C\u70ED\u7F16\u8BD1</li></ul><div class="language-bash ext-sh line-numbers-mode"><pre class="language-bash"><code>go get github.com/silenceper/gowatch



<span class="token comment"># \u4F7F\u7528</span>



gowatch main.go
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><ul><li>Hello, Gin</li></ul><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main



<span class="token keyword">import</span> <span class="token punctuation">(</span>

    <span class="token string">&quot;github.com/gin-gonic/gin&quot;</span>

    <span class="token string">&quot;net/http&quot;</span>

<span class="token punctuation">)</span>



<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

    router <span class="token operator">:=</span> gin<span class="token punctuation">.</span><span class="token function">Default</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

    router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>

        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;Hello World&quot;</span><span class="token punctuation">)</span>

    <span class="token punctuation">}</span><span class="token punctuation">)</span>

    router<span class="token punctuation">.</span><span class="token function">Run</span><span class="token punctuation">(</span><span class="token string">&quot;:8000&quot;</span><span class="token punctuation">)</span>

<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>\u8FD0\u884C\uFF1A</p><div class="language-bash ext-sh line-numbers-mode"><pre class="language-bash"><code>gowatch main.go
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>\u6D4B\u8BD5\uFF1A\u65B0\u5EFA <code>test.http</code> \u7ED3\u5408 vs code \u7684 rest-client \u63D2\u4EF6 \u8FDB\u884C\u6D4B\u8BD5</p><div class="language-text ext-text line-numbers-mode"><pre class="language-text"><code>GET http://localhost:8000/ HTTP/1.1
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div>`,13),c=[l];function t(d,u){return s(),a("div",null,c)}var p=n(i,[["render",t],["__file","gin-begin.html.vue"]]);export{p as default};
