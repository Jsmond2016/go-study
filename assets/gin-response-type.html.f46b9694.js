import{_ as n,o as s,c as a,e as t}from"./app.50b80db3.js";const p={},e=t(`<h2 id="gin-response-type" tabindex="-1"><a class="header-anchor" href="#gin-response-type" aria-hidden="true">#</a> gin-response-type</h2><p><strong>\u77E5\u8BC6\u8981\u70B9\uFF1A</strong></p><ul><li><p><code>c.JSON</code></p></li><li><p><code>c.XML</code></p></li><li><p><code>c.YAML</code></p></li><li><p><code>c.ProtoBuf</code></p></li></ul><p>Code:</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
    <span class="token string">&quot;net/http&quot;</span>

    <span class="token string">&quot;github.com/gin-gonic/gin&quot;</span>
    <span class="token string">&quot;github.com/gin-gonic/gin/testdata/protoexample&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
    router <span class="token operator">:=</span> gin<span class="token punctuation">.</span><span class="token function">Default</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

    <span class="token comment">// gin.H is a shortcut for map[string]interface{}</span>
    router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/json-gin&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        <span class="token comment">// c.JSON(http.StatusOK, gin.H{&quot;message&quot;: &quot;hey&quot;, &quot;status&quot;: http.StatusOK})</span>
        c<span class="token punctuation">.</span><span class="token function">JSON</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> gin<span class="token punctuation">.</span>H<span class="token punctuation">{</span><span class="token string">&quot;code&quot;</span><span class="token punctuation">:</span> <span class="token number">0</span><span class="token punctuation">,</span> <span class="token string">&quot;data&quot;</span><span class="token punctuation">:</span> <span class="token boolean">nil</span><span class="token punctuation">}</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>

    router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/json&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        <span class="token comment">// You also can use a struct</span>
        <span class="token keyword">var</span> msg <span class="token keyword">struct</span> <span class="token punctuation">{</span>
            Name    <span class="token builtin">string</span> <span class="token string">\`json:&quot;userName&quot;\`</span>
            Message <span class="token builtin">string</span> <span class="token string">\`json:&quot;msg&quot;\`</span>
            Number  <span class="token builtin">int</span>    <span class="token string">\`json:&quot;num&quot;\`</span>
        <span class="token punctuation">}</span>
        msg<span class="token punctuation">.</span>Name <span class="token operator">=</span> <span class="token string">&quot;Tony&quot;</span>
        msg<span class="token punctuation">.</span>Message <span class="token operator">=</span> <span class="token string">&quot;\u7406\u53D1&quot;</span>
        msg<span class="token punctuation">.</span>Number <span class="token operator">=</span> <span class="token number">100</span>
        <span class="token comment">// \u6CE8\u610F msg.Name \u53D8\u6210\u4E86 &quot;user&quot; \u5B57\u6BB5</span>
        <span class="token comment">// \u4EE5\u4E0B\u65B9\u5F0F\u90FD\u4F1A\u8F93\u51FA :   {&quot;userName&quot;: &quot;Tony&quot;, &quot;msg&quot;: &quot;\u7406\u53D1&quot;, &quot;num&quot;: 100}</span>
        c<span class="token punctuation">.</span><span class="token function">JSON</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> msg<span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>

    router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/json-xml&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">XML</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> gin<span class="token punctuation">.</span>H<span class="token punctuation">{</span><span class="token string">&quot;user&quot;</span><span class="token punctuation">:</span> <span class="token string">&quot;Tony&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;message&quot;</span><span class="token punctuation">:</span> <span class="token string">&quot;hey&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;status&quot;</span><span class="token punctuation">:</span> http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">}</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>

    router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/json-yml&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        c<span class="token punctuation">.</span><span class="token function">YAML</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> gin<span class="token punctuation">.</span>H<span class="token punctuation">{</span><span class="token string">&quot;message&quot;</span><span class="token punctuation">:</span> <span class="token string">&quot;Tony&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;status&quot;</span><span class="token punctuation">:</span> http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">}</span><span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>

    router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/protobuf&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
        reps <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int64</span><span class="token punctuation">{</span><span class="token function">int64</span><span class="token punctuation">(</span><span class="token number">1</span><span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token function">int64</span><span class="token punctuation">(</span><span class="token number">2</span><span class="token punctuation">)</span><span class="token punctuation">}</span>
        label <span class="token operator">:=</span> <span class="token string">&quot;test&quot;</span>
        <span class="token comment">// The specific definition of protobuf is written in the testdata/protoexample file.</span>
        data <span class="token operator">:=</span> <span class="token operator">&amp;</span>protoexample<span class="token punctuation">.</span>Test<span class="token punctuation">{</span>
            Label<span class="token punctuation">:</span> <span class="token operator">&amp;</span>label<span class="token punctuation">,</span>
            Reps<span class="token punctuation">:</span>  reps<span class="token punctuation">,</span>
        <span class="token punctuation">}</span>
        <span class="token comment">// Note that data becomes binary data in the response</span>
        <span class="token comment">// Will output protoexample.Test protobuf serialized data</span>
        c<span class="token punctuation">.</span><span class="token function">ProtoBuf</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> data<span class="token punctuation">)</span>
    <span class="token punctuation">}</span><span class="token punctuation">)</span>

    <span class="token comment">// Listen and serve on 0.0.0.0:8080</span>
    router<span class="token punctuation">.</span><span class="token function">Run</span><span class="token punctuation">(</span><span class="token string">&quot;:8080&quot;</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>\u6D4B\u8BD5</p><div class="language-http ext-http line-numbers-mode"><pre class="language-http"><code>@base = 127.0.0.1:8080

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
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,7),o=[e];function c(u,i){return s(),a("div",null,o)}var r=n(p,[["render",c],["__file","gin-response-type.html.vue"]]);export{r as default};
