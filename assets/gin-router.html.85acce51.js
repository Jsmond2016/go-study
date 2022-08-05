import{_ as n,o as s,c as a,e as t}from"./app.6882f82d.js";const p={},e=t(`<h2 id="gin-\u8DEF\u7531" tabindex="-1"><a class="header-anchor" href="#gin-\u8DEF\u7531" aria-hidden="true">#</a> gin-\u8DEF\u7531</h2><p>\u4EE3\u7801\uFF1A</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
    <span class="token string">&quot;net/http&quot;</span>
    <span class="token string">&quot;time&quot;</span>

    <span class="token string">&quot;github.com/gin-gonic/gin&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
    router <span class="token operator">:=</span> gin<span class="token punctuation">.</span><span class="token function">Default</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

    router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/get&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;GET METHOD&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>

    router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/get-params/:name&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        username <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">Param</span><span class="token punctuation">(</span><span class="token string">&quot;name&quot;</span><span class="token punctuation">)</span>
        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> username<span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>

    router<span class="token punctuation">.</span><span class="token function">POST</span><span class="token punctuation">(</span><span class="token string">&quot;/post&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;POST METHOD&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>
    router<span class="token punctuation">.</span><span class="token function">PUT</span><span class="token punctuation">(</span><span class="token string">&quot;/put&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;PUT METHOD&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>
    router<span class="token punctuation">.</span><span class="token function">DELETE</span><span class="token punctuation">(</span><span class="token string">&quot;/delete&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;DELETE METHOD&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>
    router<span class="token punctuation">.</span><span class="token function">PATCH</span><span class="token punctuation">(</span><span class="token string">&quot;/patch&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;PATCH METHOD&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>
    router<span class="token punctuation">.</span><span class="token function">HEAD</span><span class="token punctuation">(</span><span class="token string">&quot;/head&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;HEAD METHOD&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>
    router<span class="token punctuation">.</span><span class="token function">OPTIONS</span><span class="token punctuation">(</span><span class="token string">&quot;/options&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> <span class="token string">&quot;OPTIONS METHOD&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>

    s <span class="token operator">:=</span> <span class="token operator">&amp;</span>http<span class="token punctuation">.</span>Server<span class="token punctuation">{</span>
        Addr<span class="token punctuation">:</span>           <span class="token string">&quot;:8080&quot;</span><span class="token punctuation">,</span>
        Handler<span class="token punctuation">:</span>        router<span class="token punctuation">,</span>
        ReadTimeout<span class="token punctuation">:</span>    <span class="token number">10</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">,</span>
        WriteTimeout<span class="token punctuation">:</span>   <span class="token number">10</span> <span class="token operator">*</span> time<span class="token punctuation">.</span>Second<span class="token punctuation">,</span>
        MaxHeaderBytes<span class="token punctuation">:</span> <span class="token number">1</span> <span class="token operator">&lt;&lt;</span> <span class="token number">20</span><span class="token punctuation">,</span>
    <span class="token punctuation">}</span>
    s<span class="token punctuation">.</span><span class="token function">ListenAndServe</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>\u6D4B\u8BD5\uFF1A</p><div class="language-http ext-http line-numbers-mode"><pre class="language-http"><code>@base = 127.0.0.1:8080

### 

<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/get?id=12312323&amp;num=2</span> <span class="token http-version property">HTTP/1.1</span></span>

###
<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/get-params/xiaoming</span> <span class="token http-version property">HTTP/1.1</span></span>

###
<span class="token request-line"><span class="token method property">POST</span> <span class="token request-target url">http://{{base}}/post?id=12312323&amp;num=2</span> <span class="token http-version property">HTTP/1.1</span></span>
<span class="token header"><span class="token header-name keyword">Content-Type</span><span class="token punctuation">:</span> <span class="token header-value">application/json</span></span>
<span class="token application-json">
<span class="token punctuation">{</span>
  <span class="token string-property property">&quot;name&quot;</span><span class="token operator">:</span> <span class="token string">&quot;Tony&quot;</span><span class="token punctuation">,</span>
  <span class="token string-property property">&quot;password&quot;</span><span class="token operator">:</span> <span class="token string">&quot;123456&quot;</span>
<span class="token punctuation">}</span>

###
</span><span class="token request-line"><span class="token method property">PUT</span> <span class="token request-target url">http://{{base}}/put?id=12312323&amp;num=2</span> <span class="token http-version property">HTTP/1.1</span></span>
<span class="token header"><span class="token header-name keyword">Content-Type</span><span class="token punctuation">:</span> <span class="token header-value">application/json</span></span>
<span class="token application-json">
<span class="token punctuation">{</span>
  <span class="token string-property property">&quot;name&quot;</span><span class="token operator">:</span> <span class="token string">&quot;Tony&quot;</span><span class="token punctuation">,</span>
  <span class="token string-property property">&quot;password&quot;</span><span class="token operator">:</span> <span class="token string">&quot;123456&quot;</span>
<span class="token punctuation">}</span>

###
<span class="token constant">DELETE</span>  <span class="token literal-property property">http</span><span class="token operator">:</span><span class="token operator">/</span><span class="token operator">/</span><span class="token punctuation">{</span><span class="token punctuation">{</span>base<span class="token punctuation">}</span><span class="token punctuation">}</span><span class="token operator">/</span><span class="token keyword">delete</span><span class="token operator">?</span>id<span class="token operator">=</span><span class="token number">12312323</span><span class="token operator">&amp;</span>num<span class="token operator">=</span><span class="token number">2</span> <span class="token constant">HTTP</span><span class="token operator">/</span><span class="token number">1.1</span>
Content<span class="token operator">-</span>Type<span class="token operator">:</span> application<span class="token operator">/</span>json

<span class="token punctuation">{</span>
  <span class="token string-property property">&quot;id&quot;</span><span class="token operator">:</span> <span class="token number">123</span>
<span class="token punctuation">}</span>

###
</span><span class="token request-line"><span class="token method property">PATCH</span> <span class="token request-target url">http://{{base}}/patch?id=12312323&amp;num=2</span> <span class="token http-version property">HTTP/1.1</span></span>
<span class="token header"><span class="token header-name keyword">Content-Type</span><span class="token punctuation">:</span> <span class="token header-value">application/json</span></span>
<span class="token application-json">
<span class="token punctuation">{</span>
  <span class="token string-property property">&quot;id&quot;</span><span class="token operator">:</span> <span class="token number">123</span>
<span class="token punctuation">}</span>
###
<span class="token constant">HEAD</span>  <span class="token literal-property property">http</span><span class="token operator">:</span><span class="token operator">/</span><span class="token operator">/</span><span class="token punctuation">{</span><span class="token punctuation">{</span>base<span class="token punctuation">}</span><span class="token punctuation">}</span><span class="token operator">/</span>head<span class="token operator">?</span>id<span class="token operator">=</span><span class="token number">12312323</span><span class="token operator">&amp;</span>num<span class="token operator">=</span><span class="token number">2</span> <span class="token constant">HTTP</span><span class="token operator">/</span><span class="token number">1.1</span>


###

</span><span class="token request-line"><span class="token method property">OPTIONS</span> <span class="token request-target url">http://{{base}}/options?id=12312323&amp;num=2</span> <span class="token http-version property">HTTP/1.1</span></span>
<span class="token header"><span class="token header-name keyword">Content-Type</span><span class="token punctuation">:</span> <span class="token header-value">application/json</span></span>
<span class="token application-json">
<span class="token punctuation">{</span>
  <span class="token string-property property">&quot;id&quot;</span><span class="token operator">:</span> <span class="token number">123</span>
<span class="token punctuation">}</span>
</span></code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,5),o=[e];function c(i,u){return s(),a("div",null,o)}var r=n(p,[["render",c],["__file","gin-router.html.vue"]]);export{r as default};
