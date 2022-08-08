import{_ as n,o as s,c as a,e as t}from"./app.50b80db3.js";const e={},p=t(`<h2 id="gin-ping-\u6D4B\u8BD5" tabindex="-1"><a class="header-anchor" href="#gin-ping-\u6D4B\u8BD5" aria-hidden="true">#</a> Gin-ping \u6D4B\u8BD5</h2><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
    <span class="token string">&quot;net/http&quot;</span>
    <span class="token string">&quot;time&quot;</span>

    <span class="token string">&quot;github.com/gin-gonic/gin&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
    router <span class="token operator">:=</span> gin<span class="token punctuation">.</span><span class="token function">Default</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
    router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/ping&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;Hello World&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>
    <span class="token comment">// r.Run() // \u76D1\u542C\u5E76\u5728 0.0.0.0:8080 \u4E0A\u542F\u52A8\u670D\u52A1</span>
    <span class="token comment">// http.ListenAndServe(&quot;:8888&quot;, router)</span>

    <span class="token comment">// \u81EA\u5B9A\u4E49HTTP\u670D\u52A1\u5668\u7684\u914D\u7F6E</span>
    s <span class="token operator">:=</span> <span class="token operator">&amp;</span>http<span class="token punctuation">.</span>Server<span class="token punctuation">{</span>
        Addr<span class="token punctuation">:</span>           <span class="token string">&quot;:8080&quot;</span><span class="token punctuation">,</span>
        Handler<span class="token punctuation">:</span>        router<span class="token punctuation">,</span>
        ReadTimeout<span class="token punctuation">:</span>    <span class="token number">10</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">,</span>
        WriteTimeout<span class="token punctuation">:</span>   <span class="token number">10</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">,</span>
        MaxHeaderBytes<span class="token punctuation">:</span> <span class="token number">1</span> <span class="token operator">&lt;&lt;</span> <span class="token number">20</span><span class="token punctuation">,</span>
    <span class="token punctuation">}</span>
    s<span class="token punctuation">.</span><span class="token function">ListenAndServe</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>\u6D4B\u8BD5\u4EE3\u7801\uFF1A</p><div class="language-http ext-http line-numbers-mode"><pre class="language-http"><code>@base = 127.0.0.1:8080

### 

<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/v1/ping</span> <span class="token http-version property">HTTP/1.1</span></span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,4),i=[p];function o(c,l){return s(),a("div",null,i)}var r=n(e,[["render",o],["__file","gin-ping.html.vue"]]);export{r as default};
