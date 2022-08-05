import{_ as n,o as s,c as a,e as t}from"./app.6882f82d.js";const p={},o=t(`<h2 id="gin-router-group" tabindex="-1"><a class="header-anchor" href="#gin-router-group" aria-hidden="true">#</a> gin-router-group</h2><p>Code:</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
    <span class="token string">&quot;fmt&quot;</span>
    <span class="token string">&quot;net/http&quot;</span>

    <span class="token string">&quot;github.com/gin-gonic/gin&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
    router <span class="token operator">:=</span> gin<span class="token punctuation">.</span><span class="token function">Default</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

    <span class="token comment">// Simple group: v1</span>
    v1 <span class="token operator">:=</span> router<span class="token punctuation">.</span><span class="token function">Group</span><span class="token punctuation">(</span><span class="token string">&quot;/v1&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">{</span>
        v1<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/login&quot;</span><span class="token punctuation">,</span> loginEndpoint<span class="token punctuation">)</span>
        v1<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/submit&quot;</span><span class="token punctuation">,</span> submitEndpoint<span class="token punctuation">)</span>
        v1<span class="token punctuation">.</span><span class="token function">POST</span><span class="token punctuation">(</span><span class="token string">&quot;/read&quot;</span><span class="token punctuation">,</span> readEndpoint<span class="token punctuation">)</span>
    <span class="token punctuation">}</span>

    <span class="token comment">// Simple group: v2</span>
    v2 <span class="token operator">:=</span> router<span class="token punctuation">.</span><span class="token function">Group</span><span class="token punctuation">(</span><span class="token string">&quot;/v2&quot;</span><span class="token punctuation">)</span>
    <span class="token punctuation">{</span>
        v2<span class="token punctuation">.</span><span class="token function">POST</span><span class="token punctuation">(</span><span class="token string">&quot;/login&quot;</span><span class="token punctuation">,</span> loginEndpoint<span class="token punctuation">)</span>
        v2<span class="token punctuation">.</span><span class="token function">POST</span><span class="token punctuation">(</span><span class="token string">&quot;/submit&quot;</span><span class="token punctuation">,</span> submitEndpoint<span class="token punctuation">)</span>
        v2<span class="token punctuation">.</span><span class="token function">POST</span><span class="token punctuation">(</span><span class="token string">&quot;/read&quot;</span><span class="token punctuation">,</span> readEndpoint<span class="token punctuation">)</span>
    <span class="token punctuation">}</span>

    router<span class="token punctuation">.</span><span class="token function">Run</span><span class="token punctuation">(</span><span class="token string">&quot;:8080&quot;</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
<span class="token keyword">func</span> <span class="token function">loginEndpoint</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
    name <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">DefaultQuery</span><span class="token punctuation">(</span><span class="token string">&quot;name&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;Guest&quot;</span><span class="token punctuation">)</span> <span class="token comment">//\u53EF\u8BBE\u7F6E\u9ED8\u8BA4\u503C</span>
    c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> fmt<span class="token punctuation">.</span><span class="token function">Sprintf</span><span class="token punctuation">(</span><span class="token string">&quot;Hello %s \\n&quot;</span><span class="token punctuation">,</span> name<span class="token punctuation">)</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">submitEndpoint</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
    name <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">DefaultQuery</span><span class="token punctuation">(</span><span class="token string">&quot;name&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;Guest&quot;</span><span class="token punctuation">)</span> <span class="token comment">//\u53EF\u8BBE\u7F6E\u9ED8\u8BA4\u503C</span>
    c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> fmt<span class="token punctuation">.</span><span class="token function">Sprintf</span><span class="token punctuation">(</span><span class="token string">&quot;Hello %s \\n&quot;</span><span class="token punctuation">,</span> name<span class="token punctuation">)</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readEndpoint</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
    name <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">DefaultQuery</span><span class="token punctuation">(</span><span class="token string">&quot;name&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;Guest&quot;</span><span class="token punctuation">)</span> <span class="token comment">//\u53EF\u8BBE\u7F6E\u9ED8\u8BA4\u503C</span>
    c<span class="token punctuation">.</span><span class="token function">String</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> fmt<span class="token punctuation">.</span><span class="token function">Sprintf</span><span class="token punctuation">(</span><span class="token string">&quot;Hello %s \\n&quot;</span><span class="token punctuation">,</span> name<span class="token punctuation">)</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>\u6D4B\u8BD5\uFF1A</p><div class="language-http ext-http line-numbers-mode"><pre class="language-http"><code>@base = 127.0.0.1:8080

### 

<span class="token request-line"><span class="token method property">GET</span> <span class="token request-target url">http://{{base}}/v1/login?name=hanru</span> <span class="token http-version property">HTTP/1.1</span></span>

###
<span class="token request-line"><span class="token method property">POST</span> <span class="token request-target url">http://{{base}}/form?name=\u6D4B\u8BD5\u4EBA\u5458&amp;num=2</span> <span class="token http-version property">HTTP/1.1</span></span>
<span class="token header"><span class="token header-name keyword">Content-Type</span><span class="token punctuation">:</span> <span class="token header-value">application/x-www-form-urlencoded</span></span>

username=Tony
&amp;password=123455
&amp;hobby=&quot;basketball&quot;
&amp;hobby=&quot;running&quot;
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,5),e=[o];function c(u,i){return s(),a("div",null,e)}var r=n(p,[["render",c],["__file","gin-router-group.html.vue"]]);export{r as default};
