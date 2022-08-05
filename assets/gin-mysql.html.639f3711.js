import{_ as n,o as s,c as a,e as t}from"./app.6882f82d.js";const p={},o=t(`<p>gin-mysql</p><p>\u77E5\u8BC6\u8981\u70B9\uFF1A\u5BF9 MySQL \u7684 crud \u64CD\u4F5C</p><p>Code\uFF1A</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">package</span> main

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;database/sql&quot;</span>
	<span class="token string">&quot;fmt&quot;</span>
	<span class="token string">&quot;log&quot;</span>
	<span class="token string">&quot;net/http&quot;</span>
	<span class="token string">&quot;strconv&quot;</span>

	<span class="token string">&quot;github.com/gin-gonic/gin&quot;</span>
	<span class="token boolean">_</span> <span class="token string">&quot;github.com/go-sql-driver/mysql&quot;</span>
<span class="token punctuation">)</span>

<span class="token comment">//\u5B9A\u4E49User\u7C7B\u578B\u7ED3\u6784</span>
<span class="token keyword">type</span> User <span class="token keyword">struct</span> <span class="token punctuation">{</span>
	Id       <span class="token builtin">int</span>    <span class="token string">\`json:&quot;id&quot;\`</span>
	Username <span class="token builtin">string</span> <span class="token string">\`json:&quot;username&quot;\`</span>
	Password <span class="token builtin">string</span> <span class="token string">\`json:&quot;password&quot;\`</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">initDB</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token operator">*</span>sql<span class="token punctuation">.</span>DB <span class="token punctuation">{</span>
	db<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> sql<span class="token punctuation">.</span><span class="token function">Open</span><span class="token punctuation">(</span><span class="token string">&quot;mysql&quot;</span><span class="token punctuation">,</span> <span class="token string">&quot;root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&quot;</span><span class="token punctuation">)</span>
	<span class="token keyword">return</span> db
<span class="token punctuation">}</span>

<span class="token comment">//\u5B9A\u4E49\u4E00\u4E2AgetALL\u51FD\u6570\u7528\u4E8E\u56DE\u53BB\u5168\u90E8\u7684\u4FE1\u606F</span>
<span class="token keyword">func</span> <span class="token function">getAll</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">(</span>users <span class="token punctuation">[</span><span class="token punctuation">]</span>User<span class="token punctuation">,</span> err <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	<span class="token comment">//1.\u64CD\u4F5C\u6570\u636E\u5E93</span>
	db <span class="token operator">:=</span> <span class="token function">initDB</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token comment">//\u9519\u8BEF\u68C0\u67E5</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatal</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u63A8\u8FDF\u6570\u636E\u5E93\u8FDE\u63A5\u7684\u5173\u95ED</span>
	<span class="token keyword">defer</span> db<span class="token punctuation">.</span><span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	<span class="token comment">//2.\u67E5\u8BE2</span>
	rows<span class="token punctuation">,</span> err <span class="token operator">:=</span> db<span class="token punctuation">.</span><span class="token function">Query</span><span class="token punctuation">(</span><span class="token string">&quot;SELECT id, username, password FROM user_info&quot;</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatal</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">for</span> rows<span class="token punctuation">.</span><span class="token function">Next</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">var</span> user User
		<span class="token comment">//\u904D\u5386\u8868\u4E2D\u6240\u6709\u884C\u7684\u4FE1\u606F</span>
		rows<span class="token punctuation">.</span><span class="token function">Scan</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>user<span class="token punctuation">.</span>Id<span class="token punctuation">,</span> <span class="token operator">&amp;</span>user<span class="token punctuation">.</span>Username<span class="token punctuation">,</span> <span class="token operator">&amp;</span>user<span class="token punctuation">.</span>Password<span class="token punctuation">)</span>
		<span class="token comment">//\u5C06user\u6DFB\u52A0\u5230users\u4E2D</span>
		users <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>users<span class="token punctuation">,</span> user<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u6700\u540E\u5173\u95ED\u8FDE\u63A5</span>
	<span class="token keyword">defer</span> rows<span class="token punctuation">.</span><span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">return</span>
<span class="token punctuation">}</span>

<span class="token comment">//\u63D2\u5165\u6570\u636E</span>
<span class="token keyword">func</span> <span class="token function">add</span><span class="token punctuation">(</span>user User<span class="token punctuation">)</span> <span class="token punctuation">(</span>Id <span class="token builtin">int</span><span class="token punctuation">,</span> err <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	<span class="token comment">//1.\u64CD\u4F5C\u6570\u636E\u5E93</span>
	db <span class="token operator">:=</span> <span class="token function">initDB</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token comment">//\u9519\u8BEF\u68C0\u67E5</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatal</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u63A8\u8FDF\u6570\u636E\u5E93\u8FDE\u63A5\u7684\u5173\u95ED</span>
	<span class="token keyword">defer</span> db<span class="token punctuation">.</span><span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	stmt<span class="token punctuation">,</span> err <span class="token operator">:=</span> db<span class="token punctuation">.</span><span class="token function">Prepare</span><span class="token punctuation">(</span><span class="token string">&quot;INSERT INTO user_info(username, password) VALUES (?, ?)&quot;</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u6267\u884C\u63D2\u5165\u64CD\u4F5C</span>
	rs<span class="token punctuation">,</span> err <span class="token operator">:=</span> stmt<span class="token punctuation">.</span><span class="token function">Exec</span><span class="token punctuation">(</span>user<span class="token punctuation">.</span>Username<span class="token punctuation">,</span> user<span class="token punctuation">.</span>Password<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u8FD4\u56DE\u63D2\u5165\u7684id</span>
	id<span class="token punctuation">,</span> err <span class="token operator">:=</span> rs<span class="token punctuation">.</span><span class="token function">LastInsertId</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalln</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u5C06id\u7C7B\u578B\u8F6C\u6362</span>
	Id <span class="token operator">=</span> <span class="token function">int</span><span class="token punctuation">(</span>id<span class="token punctuation">)</span>
	<span class="token keyword">defer</span> stmt<span class="token punctuation">.</span><span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">return</span>
<span class="token punctuation">}</span>

<span class="token comment">//\u4FEE\u6539\u6570\u636E</span>
<span class="token keyword">func</span> <span class="token function">update</span><span class="token punctuation">(</span>user User<span class="token punctuation">)</span> <span class="token punctuation">(</span>rowsAffected <span class="token builtin">int64</span><span class="token punctuation">,</span> err <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

	<span class="token comment">//1.\u64CD\u4F5C\u6570\u636E\u5E93</span>
	db <span class="token operator">:=</span> <span class="token function">initDB</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token comment">//\u9519\u8BEF\u68C0\u67E5</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatal</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u63A8\u8FDF\u6570\u636E\u5E93\u8FDE\u63A5\u7684\u5173\u95ED</span>
	<span class="token keyword">defer</span> db<span class="token punctuation">.</span><span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	stmt<span class="token punctuation">,</span> err <span class="token operator">:=</span> db<span class="token punctuation">.</span><span class="token function">Prepare</span><span class="token punctuation">(</span><span class="token string">&quot;UPDATE  user_info SET username=?, password=? WHERE id=?&quot;</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u6267\u884C\u4FEE\u6539\u64CD\u4F5C</span>
	rs<span class="token punctuation">,</span> err <span class="token operator">:=</span> stmt<span class="token punctuation">.</span><span class="token function">Exec</span><span class="token punctuation">(</span>user<span class="token punctuation">.</span>Username<span class="token punctuation">,</span> user<span class="token punctuation">.</span>Password<span class="token punctuation">,</span> user<span class="token punctuation">.</span>Id<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u8FD4\u56DE\u63D2\u5165\u7684id</span>
	rowsAffected<span class="token punctuation">,</span> err <span class="token operator">=</span> rs<span class="token punctuation">.</span><span class="token function">RowsAffected</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalln</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token keyword">defer</span> stmt<span class="token punctuation">.</span><span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">return</span>
<span class="token punctuation">}</span>

<span class="token comment">//\u901A\u8FC7id\u5220\u9664</span>
<span class="token keyword">func</span> <span class="token function">del</span><span class="token punctuation">(</span>id <span class="token builtin">int</span><span class="token punctuation">)</span> <span class="token punctuation">(</span>rows <span class="token builtin">int</span><span class="token punctuation">,</span> err <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token comment">// 1.\u64CD\u4F5C\u6570\u636E\u5E93</span>
	db <span class="token operator">:=</span> <span class="token function">initDB</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token comment">//\u9519\u8BEF\u68C0\u67E5</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatal</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u63A8\u8FDF\u6570\u636E\u5E93\u8FDE\u63A5\u7684\u5173\u95ED</span>
	<span class="token keyword">defer</span> db<span class="token punctuation">.</span><span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	stmt<span class="token punctuation">,</span> err <span class="token operator">:=</span> db<span class="token punctuation">.</span><span class="token function">Prepare</span><span class="token punctuation">(</span><span class="token string">&quot;DELETE FROM user_info WHERE id=?&quot;</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalln</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	rs<span class="token punctuation">,</span> err <span class="token operator">:=</span> stmt<span class="token punctuation">.</span><span class="token function">Exec</span><span class="token punctuation">(</span>id<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalln</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token comment">//\u5220\u9664\u7684\u884C\u6570</span>
	row<span class="token punctuation">,</span> err <span class="token operator">:=</span> rs<span class="token punctuation">.</span><span class="token function">RowsAffected</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		log<span class="token punctuation">.</span><span class="token function">Fatalln</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token keyword">defer</span> stmt<span class="token punctuation">.</span><span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	rows <span class="token operator">=</span> <span class="token function">int</span><span class="token punctuation">(</span>row<span class="token punctuation">)</span>
	<span class="token keyword">return</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">main</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token comment">//\u521B\u5EFA\u4E00\u4E2A\u8DEF\u7531Handler</span>
	router <span class="token operator">:=</span> gin<span class="token punctuation">.</span><span class="token function">Default</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	<span class="token comment">//get\u65B9\u6CD5\u7684\u67E5\u8BE2</span>
	router<span class="token punctuation">.</span><span class="token function">GET</span><span class="token punctuation">(</span><span class="token string">&quot;/user&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		users<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">getAll</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatal</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		<span class="token comment">//H is a shortcut for map[string]interface{}</span>
		c<span class="token punctuation">.</span><span class="token function">JSON</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> gin<span class="token punctuation">.</span>H<span class="token punctuation">{</span>
			<span class="token string">&quot;result&quot;</span><span class="token punctuation">:</span> users<span class="token punctuation">,</span>
			<span class="token string">&quot;count&quot;</span><span class="token punctuation">:</span>  <span class="token function">len</span><span class="token punctuation">(</span>users<span class="token punctuation">)</span><span class="token punctuation">,</span>
		<span class="token punctuation">}</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>

	router<span class="token punctuation">.</span><span class="token function">POST</span><span class="token punctuation">(</span><span class="token string">&quot;/add&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">var</span> u User
		err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">Bind</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>u<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatal</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		Id<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">add</span><span class="token punctuation">(</span>u<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatal</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		fmt<span class="token punctuation">.</span><span class="token function">Print</span><span class="token punctuation">(</span><span class="token string">&quot;id=&quot;</span><span class="token punctuation">,</span> Id<span class="token punctuation">)</span>
		fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">&quot;===========================================&quot;</span><span class="token punctuation">)</span>
		c<span class="token punctuation">.</span><span class="token function">JSON</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> gin<span class="token punctuation">.</span>H<span class="token punctuation">{</span>
			<span class="token string">&quot;message&quot;</span><span class="token punctuation">:</span> fmt<span class="token punctuation">.</span><span class="token function">Sprintf</span><span class="token punctuation">(</span><span class="token string">&quot;%s \u63D2\u5165\u6210\u529F, id \u4E3A %d&quot;</span><span class="token punctuation">,</span> u<span class="token punctuation">.</span>Username<span class="token punctuation">,</span> Id<span class="token punctuation">)</span><span class="token punctuation">,</span>
		<span class="token punctuation">}</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>

	<span class="token comment">// \u5229\u7528put\u65B9\u6CD5\u4FEE\u6539\u6570\u636E</span>
	router<span class="token punctuation">.</span><span class="token function">PUT</span><span class="token punctuation">(</span><span class="token string">&quot;/update&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">var</span> u User
		err <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">Bind</span><span class="token punctuation">(</span><span class="token operator">&amp;</span>u<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatal</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		num<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> <span class="token function">update</span><span class="token punctuation">(</span>u<span class="token punctuation">)</span>
		fmt<span class="token punctuation">.</span><span class="token function">Print</span><span class="token punctuation">(</span><span class="token string">&quot;num=&quot;</span><span class="token punctuation">,</span> num<span class="token punctuation">)</span>
		c<span class="token punctuation">.</span><span class="token function">JSON</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> gin<span class="token punctuation">.</span>H<span class="token punctuation">{</span>
			<span class="token string">&quot;message&quot;</span><span class="token punctuation">:</span> fmt<span class="token punctuation">.</span><span class="token function">Sprintf</span><span class="token punctuation">(</span><span class="token string">&quot;\u4FEE\u6539id: %d \u6210\u529F&quot;</span><span class="token punctuation">,</span> u<span class="token punctuation">.</span>Id<span class="token punctuation">)</span><span class="token punctuation">,</span>
		<span class="token punctuation">}</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>

	<span class="token comment">//\u5229\u7528DELETE\u8BF7\u6C42\u65B9\u6CD5\u901A\u8FC7id\u5220\u9664</span>
	router<span class="token punctuation">.</span><span class="token function">DELETE</span><span class="token punctuation">(</span><span class="token string">&quot;/delete/:id&quot;</span><span class="token punctuation">,</span> <span class="token keyword">func</span><span class="token punctuation">(</span>c <span class="token operator">*</span>gin<span class="token punctuation">.</span>Context<span class="token punctuation">)</span> <span class="token punctuation">{</span>
		id <span class="token operator">:=</span> c<span class="token punctuation">.</span><span class="token function">Param</span><span class="token punctuation">(</span><span class="token string">&quot;id&quot;</span><span class="token punctuation">)</span>

		Id<span class="token punctuation">,</span> err <span class="token operator">:=</span> strconv<span class="token punctuation">.</span><span class="token function">Atoi</span><span class="token punctuation">(</span>id<span class="token punctuation">)</span>

		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalln</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		rows<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">del</span><span class="token punctuation">(</span>Id<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			log<span class="token punctuation">.</span><span class="token function">Fatalln</span><span class="token punctuation">(</span>err<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		fmt<span class="token punctuation">.</span><span class="token function">Println</span><span class="token punctuation">(</span><span class="token string">&quot;delete rows &quot;</span><span class="token punctuation">,</span> rows<span class="token punctuation">)</span>

		c<span class="token punctuation">.</span><span class="token function">JSON</span><span class="token punctuation">(</span>http<span class="token punctuation">.</span>StatusOK<span class="token punctuation">,</span> gin<span class="token punctuation">.</span>H<span class="token punctuation">{</span>
			<span class="token string">&quot;message&quot;</span><span class="token punctuation">:</span> fmt<span class="token punctuation">.</span><span class="token function">Sprintf</span><span class="token punctuation">(</span><span class="token string">&quot;Successfully deleted user: %s&quot;</span><span class="token punctuation">,</span> id<span class="token punctuation">)</span><span class="token punctuation">,</span>
		<span class="token punctuation">}</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">)</span>

	router<span class="token punctuation">.</span><span class="token function">Run</span><span class="token punctuation">(</span><span class="token string">&quot;:8080&quot;</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>\u6D4B\u8BD5\uFF1A</p><div class="language-http ext-http line-numbers-mode"><pre class="language-http"><code>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div>`,6),e=[o];function c(i,l){return s(),a("div",null,e)}var k=n(p,[["render",c],["__file","gin-mysql.html.vue"]]);export{k as default};
