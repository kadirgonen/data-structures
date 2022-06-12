<h1 class="mt-0 text-4xl font-extrabold text-neutral-900 dark:text-neutral">
Golang Datastructures: Trees

Table of Contents
</summary>
<div class="py-2 border-dotted ltr:border-l rtl:border-r rtl:pr-5 ltr:pl-5 ltr:-ml-5 rtl:-mr-5 border-neutral-300 dark:border-neutral-600">
<nav id="TableOfContents">
<ul>
<li><a href="#a-touch-of-theory">A touch of theory</a></li>
<li><a href="#modelling-an-html-document">Modelling an HTML document</a></li>
<li><a href="#building-mydom---a-drop-in-replacement-for-the-dom-">Building MyDOM - a drop-in replacement for the DOM 😂</a></li>
<li><a href="#implementing-node-lookup-">Implementing Node Lookup 🔎</a>
<ul>
<li><a href="#breadth-first-search-">Breadth-first search ⬅➡</a></li>
<li><a href="#depth-first-search-">Depth-first search ⬇</a></li>
</ul>
</li>
<li><a href="#finding-by-class-name-">Finding by class name 🔎</a></li>
<li><a href="#deleting-nodes-">Deleting nodes 🗑</a></li>
<li><a href="#where-to-next">Where to next?</a></li>
</ul>
</nav>
</div>
</details>
</div>
</div>
<div class="min-w-0 min-h-0 max-w-prose">
<p>You can spend quite a bit of your programming career without working with trees,
or just by simply avoiding them if you don’t understand them (which is what I
had been doing for a while).</p>
<figure class="imagecaption">
<img class="caption" src="/golang-datastructures-trees/tree_swing.png" caption="" alt="Person sitting on a swing on a tree">
<span class="caption-text" style="font-size: 0.75em;display: table;margin: 0 auto 1.5em;"></span>
</figure>
<p>Now, don&rsquo;t get me wrong - arrays, lists, stacks and queues are quite powerful
data structures and can take you pretty far, but there is a limit to their
capabilities, how you can use them and how efficient that usage can be. When you
throw in hash tables to that mix, you can solve quite some problems, but for
many of the problems out there trees are a powerful (and maybe the only) tool
if you have them under your belt.</p>
<p>So, let&rsquo;s look at trees and then we can try to use them in a small exercise.</p>
<h2 id="a-touch-of-theory" class="relative group">A touch of theory <span class="absolute top-0 w-6 transition-opacity opacity-0 ltr:-left-6 rtl:-right-6 not-prose group-hover:opacity-100"><a class="group-hover:text-primary-300 dark:group-hover:text-neutral-700" style="text-decoration-line: none !important;" href="#a-touch-of-theory" aria-label="Anchor">#</a></span></h2>
<p>Arrays, lists, queues, stacks store data in a collection that has a start and an
end, hence they are called &ldquo;linear&rdquo;. But when it comes to trees and graphs,
things can get confusing since the data is not stored in a linear fashion.</p>
<p>Trees are called nonlinear data structures. In fact, you can also say that trees
are hierarchical data structures since the data is stored in a hierarchical way.</p>
<p>For your reading pleasure, Wikipedia’s definition of trees:</p>
<blockquote>
<p>A tree is a data structure made up of nodes or vertices and edges without
having any cycle. The tree with no nodes is called the null or empty tree. A
tree that is not empty consists of a root node and potentially many levels of
additional nodes that form a hierarchy.</p>
</blockquote>
<p>What the definition states are that a tree is just a combination of nodes (or
vertices) and edges (or links between the nodes) without having a cycle.</p>
<figure class="imagecaption">
<img class="caption" src="/golang-datastructures-trees/invalid-tree.png" caption="" alt="Graph data structure graphic">
<span class="caption-text" style="font-size: 0.75em;display: table;margin: 0 auto 1.5em;"></span>
</figure>
<p>For example, the data structure represented on the diagram is a combination of
nodes, named from A to F, with six edges. Although all of its elements look
like they construct a tree, the nodes A, D, E and F have a cycle, therefore
this structure is not a tree.</p>
<p>If we would break the edge between nodes F and E and add a new node called G
with an edge between F and G, we would end up with something like this:</p>
<figure class="imagecaption">
<img class="caption" src="/golang-datastructures-trees/valid-tree.png" caption="" alt="Tree data structure graphic">
<span class="caption-text" style="font-size: 0.75em;display: table;margin: 0 auto 1.5em;"></span>
</figure>
<p>Now, since we eliminated the cycle in this graph, we can say that we have a
valid tree. It has a <strong>root</strong> with the name A, with a total of 7 <strong>nodes</strong>.
Node A has 3 <strong>children</strong> (B, D &amp; F) and those have 3 children (C, E &amp; G
respectively). Therefore, node A has 6 <strong>descendants</strong>. Also, this tree has 3
leaf nodes (C, E &amp; G) or nodes that have no children.</p>
<p>What do B, D &amp; F have in common? They are <strong>siblings</strong> because they have the
same parent (node A). They all reside on <strong>level</strong> 1 because to get from each
of them to the root we need to take only one step. For example, node G has
level 2, because the <strong>path</strong> from G to A is: G -&gt; F -&gt; A, hence we need to
follow two edges to get to A.</p>
<p>Now that we know a bit of theory about trees, let’s see how we can solve some
problems.</p>
<figure class="imagecaption">
<img class="caption" src="/golang-datastructures-trees/web_development.png" caption="" alt="Web developer">
<span class="caption-text" style="font-size: 0.75em;display: table;margin: 0 auto 1.5em;"></span>
</figure>
<h2 id="modelling-an-html-document" class="relative group">Modelling an HTML document <span class="absolute top-0 w-6 transition-opacity opacity-0 ltr:-left-6 rtl:-right-6 not-prose group-hover:opacity-100"><a class="group-hover:text-primary-300 dark:group-hover:text-neutral-700" style="text-decoration-line: none !important;" href="#modelling-an-html-document" aria-label="Anchor">#</a></span></h2>
<p>If you are a software developer that has never written any HTML, I will just
assume that you have seen (or have an idea) what HTML looks like. If you have
not, then I encourage you to right click on the page that you are reading this
and click on &lsquo;View Source&rsquo;.</p>
<p>Seriously, go for it, I&rsquo;ll wait&hellip;</p>
<p>Browsers have this thing baked in, called the DOM - a cross-platform and
language-independent application programming interface, which treats internet
documents as a tree structure wherein each node is an object representing a part
of the document. This means that when the browser reads your document&rsquo;s HTML
code it will load it and create a DOM out of it.</p>
<p>So, let’s imagine for a second we are developers working on a browser, like
Chrome or Firefox and we need to model the DOM. Well, to make this exercise
easier, let’s see a tiny HTML document:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-html" data-lang="html"><span class="p">&lt;</span><span class="nt">html</span><span class="p">&gt;</span>
  <span class="p">&lt;</span><span class="nt">h1</span><span class="p">&gt;</span>Hello, World!<span class="p">&lt;/</span><span class="nt">h1</span><span class="p">&gt;</span>
  <span class="p">&lt;</span><span class="nt">p</span><span class="p">&gt;</span>This is a simple HTML document.<span class="p">&lt;/</span><span class="nt">p</span><span class="p">&gt;</span>
<span class="p">&lt;/</span><span class="nt">html</span><span class="p">&gt;</span></code></pre></div>
<p>So, if we would model this document as a tree, it would look something like
this:</p>
<figure class="imagecaption">
<img class="caption" src="/golang-datastructures-trees/html-document-tree.png" caption="" alt="DOM node tree">
<span class="caption-text" style="font-size: 0.75em;display: table;margin: 0 auto 1.5em;"></span>
</figure>
<p>Now, we could treat the text nodes as separate <code>Node</code>s, but we can make our
lives simpler by assuming that any HTML element can have text in it.</p>
<p>The <code>html</code> node will have two children, <code>h1</code> and <code>p</code>, which will have <code>tag</code>,
<code>text</code> and <code>children</code> as fields. Let’s put this into code:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="kd">type</span> <span class="nx">Node</span> <span class="kd">struct</span> <span class="p">{</span>
    <span class="nx">tag</span>      <span class="kt">string</span>
    <span class="nx">text</span>     <span class="kt">string</span>
    <span class="nx">children</span> <span class="p">[]</span><span class="o">*</span><span class="nx">Node</span>
<span class="p">}</span></code></pre></div>
<p>A <code>Node</code> will have only the tag name and children optionally. Let’s try to
create the HTML document we saw above as a tree of <code>Node</code>s by hand:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="kd">func</span> <span class="nf">main</span><span class="p">()</span> <span class="p">{</span>
        <span class="nx">p</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
                <span class="nx">tag</span><span class="p">:</span>  <span class="s">&#34;p&#34;</span><span class="p">,</span>
                <span class="nx">text</span><span class="p">:</span> <span class="s">&#34;This is a simple HTML document.&#34;</span><span class="p">,</span>
                <span class="nx">id</span><span class="p">:</span>   <span class="s">&#34;foo&#34;</span><span class="p">,</span>
        <span class="p">}</span>

<span class="nx">h1</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
                <span class="nx">tag</span><span class="p">:</span>  <span class="s">&#34;h1&#34;</span><span class="p">,</span>
                <span class="nx">text</span><span class="p">:</span> <span class="s">&#34;Hello, World!&#34;</span><span class="p">,</span>
        <span class="p">}</span>

<span class="nx">html</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
                <span class="nx">tag</span><span class="p">:</span>      <span class="s">&#34;html&#34;</span><span class="p">,</span>
                <span class="nx">children</span><span class="p">:</span> <span class="p">[]</span><span class="o">*</span><span class="nx">Node</span><span class="p">{</span><span class="o">&amp;</span><span class="nx">p</span><span class="p">,</span> <span class="o">&amp;</span><span class="nx">h1</span><span class="p">},</span>
        <span class="p">}</span>
<span class="p">}</span></code></pre></div>
<p>That looks okay, we have a basic tree up and running now.</p>
<h2 id="building-mydom---a-drop-in-replacement-for-the-dom-" class="relative group">Building MyDOM - a drop-in replacement for the DOM 😂 <span class="absolute top-0 w-6 transition-opacity opacity-0 ltr:-left-6 rtl:-right-6 not-prose group-hover:opacity-100"><a class="group-hover:text-primary-300 dark:group-hover:text-neutral-700" style="text-decoration-line: none !important;" href="#building-mydom---a-drop-in-replacement-for-the-dom-" aria-label="Anchor">#</a></span></h2>
<p>Now that we have some tree structure in place, let&rsquo;s take a step back and see
what kind of functionality would a DOM have. For example, if MyDOM
<sup>(TM)</sup> would be a drop-in replacement of a real DOM, then with
JavaScript we should be able to access nodes and modify them.</p>
<p>The simplest way to do this with JavaScript would be to use</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-javascript" data-lang="javascript"><span class="nb">document</span><span class="p">.</span><span class="nx">getElementById</span><span class="p">(</span><span class="s1">&#39;foo&#39;</span><span class="p">)</span></code></pre></div>
<p>This function would lookup in the <code>document</code> tree to find the node whose ID is
<code>foo</code>. Let&rsquo;s update our <code>Node</code> struct to have more attributes and then work on
writing a lookup function for our tree:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="kd">type</span> <span class="nx">Node</span> <span class="kd">struct</span> <span class="p">{</span>
  <span class="nx">tag</span>      <span class="kt">string</span>
  <span class="nx">id</span>       <span class="kt">string</span>
  <span class="nx">class</span>    <span class="kt">string</span>
  <span class="nx">children</span> <span class="p">[]</span><span class="o">*</span><span class="nx">Node</span>
<span class="p">}</span></code></pre></div>
<p>Now, each of our <code>Node</code> structs will have a <code>tag</code>, <code>children</code> which is a slice
of pointers to the children of that <code>Node</code>, <code>id</code> which is the ID of that DOM
node and <code>class</code> which is the classes that can be applied to this DOM node.</p>
<p>Now, back to our <code>getElementById</code> lookup function. Let&rsquo;s see how we could
implement it. First, let&rsquo;s build an example tree that we can use for our lookup
algorithm:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-html" data-lang="html"><span class="p">&lt;</span><span class="nt">html</span><span class="p">&gt;</span>
  <span class="p">&lt;</span><span class="nt">body</span><span class="p">&gt;</span>
    <span class="p">&lt;</span><span class="nt">h1</span><span class="p">&gt;</span>This is a H1<span class="p">&lt;/</span><span class="nt">h1</span><span class="p">&gt;</span>
    <span class="p">&lt;</span><span class="nt">p</span><span class="p">&gt;</span>
      And this is some text in a paragraph. And next to it there&#39;s an image.
      <span class="p">&lt;</span><span class="nt">img</span> <span class="na">src</span><span class="o">=</span><span class="s">&#34;http://example.com/logo.svg&#34;</span> <span class="na">alt</span><span class="o">=</span><span class="s">&#34;Example&#39;s Logo&#34;</span><span class="p">/&gt;</span>
    <span class="p">&lt;/</span><span class="nt">p</span><span class="p">&gt;</span>
    <span class="p">&lt;</span><span class="nt">div</span> <span class="na">class</span><span class="o">=</span><span class="s">&#39;footer&#39;</span><span class="p">&gt;</span>
      This is the footer of the page.
      <span class="p">&lt;</span><span class="nt">span</span> <span class="na">id</span><span class="o">=</span><span class="s">&#39;copyright&#39;</span><span class="p">&gt;</span>2019 <span class="ni">&amp;copy;</span> Ilija Eftimov<span class="p">&lt;/</span><span class="nt">span</span><span class="p">&gt;</span>
    <span class="p">&lt;/</span><span class="nt">div</span><span class="p">&gt;</span>
  <span class="p">&lt;/</span><span class="nt">body</span><span class="p">&gt;</span>
<span class="p">&lt;/</span><span class="nt">html</span><span class="p">&gt;</span></code></pre></div>
<p>This is a quite complicated HTML document. Let&rsquo;s sketch out its structure in Go
using the <code>Node</code> struct as a building block:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="nx">image</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
        <span class="nx">tag</span><span class="p">:</span> <span class="s">&#34;img&#34;</span><span class="p">,</span>
        <span class="nx">src</span><span class="p">:</span> <span class="s">&#34;http://example.com/logo.svg&#34;</span><span class="p">,</span>
        <span class="nx">alt</span><span class="p">:</span> <span class="s">&#34;Example&#39;s Logo&#34;</span><span class="p">,</span>
<span class="p">}</span>

<span class="nx">p</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
        <span class="nx">tag</span><span class="p">:</span>      <span class="s">&#34;p&#34;</span><span class="p">,</span>
        <span class="nx">text</span><span class="p">:</span>     <span class="s">&#34;And this is some text in a paragraph. And next to it there&#39;s an image.&#34;</span><span class="p">,</span>
        <span class="nx">children</span><span class="p">:</span> <span class="p">[]</span><span class="o">*</span><span class="nx">Node</span><span class="p">{</span><span class="o">&amp;</span><span class="nx">image</span><span class="p">},</span>
<span class="p">}</span>

<span class="nx">span</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
        <span class="nx">tag</span><span class="p">:</span>  <span class="s">&#34;span&#34;</span><span class="p">,</span>
        <span class="nx">id</span><span class="p">:</span>   <span class="s">&#34;copyright&#34;</span><span class="p">,</span>
        <span class="nx">text</span><span class="p">:</span> <span class="s">&#34;2019 &amp;copy; Ilija Eftimov&#34;</span><span class="p">,</span>
<span class="p">}</span>

<span class="nx">div</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
        <span class="nx">tag</span><span class="p">:</span>      <span class="s">&#34;div&#34;</span><span class="p">,</span>
        <span class="nx">class</span><span class="p">:</span>    <span class="s">&#34;footer&#34;</span><span class="p">,</span>
        <span class="nx">text</span><span class="p">:</span>     <span class="s">&#34;This is the footer of the page.&#34;</span><span class="p">,</span>
        <span class="nx">children</span><span class="p">:</span> <span class="p">[]</span><span class="o">*</span><span class="nx">Node</span><span class="p">{</span><span class="o">&amp;</span><span class="nx">span</span><span class="p">},</span>
<span class="p">}</span>

<span class="nx">h1</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
        <span class="nx">tag</span><span class="p">:</span>  <span class="s">&#34;h1&#34;</span><span class="p">,</span>
        <span class="nx">text</span><span class="p">:</span> <span class="s">&#34;This is a H1&#34;</span><span class="p">,</span>
<span class="p">}</span>

<span class="nx">body</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
        <span class="nx">tag</span><span class="p">:</span>      <span class="s">&#34;body&#34;</span><span class="p">,</span>
        <span class="nx">children</span><span class="p">:</span> <span class="p">[]</span><span class="o">*</span><span class="nx">Node</span><span class="p">{</span><span class="o">&amp;</span><span class="nx">h1</span><span class="p">,</span> <span class="o">&amp;</span><span class="nx">p</span><span class="p">,</span> <span class="o">&amp;</span><span class="nx">div</span><span class="p">},</span>
<span class="p">}</span>

<span class="nx">html</span> <span class="o">:=</span> <span class="nx">Node</span><span class="p">{</span>
        <span class="nx">tag</span><span class="p">:</span>      <span class="s">&#34;html&#34;</span><span class="p">,</span>
        <span class="nx">children</span><span class="p">:</span> <span class="p">[]</span><span class="o">*</span><span class="nx">Node</span><span class="p">{</span><span class="o">&amp;</span><span class="nx">body</span><span class="p">},</span>
<span class="p">}</span></code></pre></div>
<p>We start building this tree bottom - up. That means we create structs from the
most deeply nested structs and working up towards <code>body</code> and <code>html</code>. Let&rsquo;s look
at a graphic of our tree:</p>
<figure class="imagecaption">
<img class="caption" src="/golang-datastructures-trees/mydom-tree.png" caption="" alt="DOM node tree">
<span class="caption-text" style="font-size: 0.75em;display: table;margin: 0 auto 1.5em;"></span>
</figure>
<h2 id="implementing-node-lookup-" class="relative group">Implementing Node Lookup 🔎 <span class="absolute top-0 w-6 transition-opacity opacity-0 ltr:-left-6 rtl:-right-6 not-prose group-hover:opacity-100"><a class="group-hover:text-primary-300 dark:group-hover:text-neutral-700" style="text-decoration-line: none !important;" href="#implementing-node-lookup-" aria-label="Anchor">#</a></span></h2>
<p>So, let&rsquo;s continue with what we were up to - allow JavaScript to call
<code>getElementById</code> on our <code>document</code> and find the <code>Node</code> that it&rsquo;s looking for.</p>
<p>To do this, we have to implement a tree searching algorithm. The most popular
approaches to searching (or traversal) of graphs and trees are Breadth First
Search (BFS) and Depth First Search (DFS).</p>
<h3 id="breadth-first-search-" class="relative group">Breadth-first search ⬅➡ <span class="absolute top-0 w-6 transition-opacity opacity-0 ltr:-left-6 rtl:-right-6 not-prose group-hover:opacity-100"><a class="group-hover:text-primary-300 dark:group-hover:text-neutral-700" style="text-decoration-line: none !important;" href="#breadth-first-search-" aria-label="Anchor">#</a></span></h3>
<p>BFS, as its name suggests, takes an approach to traversal where it explores
nodes in &ldquo;width&rdquo; first before it goes in &ldquo;depth&rdquo;. Here&rsquo;s a visualisation of the
steps a BFS algorithm would take to traverse the whole tree:</p>
<p><img src="/golang-datastructures-trees/mydom-tree-bfs-steps.png" alt="" />
</p>
<p>As you can see, the algorithm will take two steps in depth (over <code>html</code> and
<code>body</code>), but then it will visit all of the <code>body</code>&rsquo;s children nodes before it
proceeds to explore in depth and visit the <code>span</code> and <code>img</code> nodes.</p>
<p>If you would like to have a step-by-step playbook, it would be:</p>
<ol>
<li>We start at the root, the <code>html</code> node</li>
<li>We push it on the <code>queue</code></li>
<li>We kick off a loop where we loop while the <code>queue</code> is not empty</li>
<li>We check the next element in the <code>queue</code> for a match. If a match is found,
we return the match and we&rsquo;re done.</li>
<li>When a match is not found, we take all of the children of the
node-under-inspection and we add them to the queue, so they can be inspected</li>
<li><code>GOTO</code> 4</li>
</ol>
<p>Let&rsquo;s see a simple implementation of the algorithm in Go and I&rsquo;ll share some
tips on how you can remember the algorithm easily.</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="kd">func</span> <span class="nf">findById</span><span class="p">(</span><span class="nx">root</span> <span class="o">*</span><span class="nx">Node</span><span class="p">,</span> <span class="nx">id</span> <span class="kt">string</span><span class="p">)</span> <span class="o">*</span><span class="nx">Node</span> <span class="p">{</span>
        <span class="nx">queue</span> <span class="o">:=</span> <span class="nb">make</span><span class="p">([]</span><span class="o">*</span><span class="nx">Node</span><span class="p">,</span> <span class="mi">0</span><span class="p">)</span>
        <span class="nx">queue</span> <span class="p">=</span> <span class="nb">append</span><span class="p">(</span><span class="nx">queue</span><span class="p">,</span> <span class="nx">root</span><span class="p">)</span>
        <span class="k">for</span> <span class="nb">len</span><span class="p">(</span><span class="nx">queue</span><span class="p">)</span> <span class="p">&gt;</span> <span class="mi">0</span> <span class="p">{</span>
                <span class="nx">nextUp</span> <span class="o">:=</span> <span class="nx">queue</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span>
                <span class="nx">queue</span> <span class="p">=</span> <span class="nx">queue</span><span class="p">[</span><span class="mi">1</span><span class="p">:]</span>
                <span class="k">if</span> <span class="nx">nextUp</span><span class="p">.</span><span class="nx">id</span> <span class="o">==</span> <span class="nx">id</span> <span class="p">{</span>
                        <span class="k">return</span> <span class="nx">nextUp</span>
                <span class="p">}</span>
                <span class="k">if</span> <span class="nb">len</span><span class="p">(</span><span class="nx">nextUp</span><span class="p">.</span><span class="nx">children</span><span class="p">)</span> <span class="p">&gt;</span> <span class="mi">0</span> <span class="p">{</span>
                        <span class="k">for</span> <span class="nx">_</span><span class="p">,</span> <span class="nx">child</span> <span class="o">:=</span> <span class="k">range</span> <span class="nx">nextUp</span><span class="p">.</span><span class="nx">children</span> <span class="p">{</span>
                                <span class="nx">queue</span> <span class="p">=</span> <span class="nb">append</span><span class="p">(</span><span class="nx">queue</span><span class="p">,</span> <span class="nx">child</span><span class="p">)</span>
                        <span class="p">}</span>
                <span class="p">}</span>
        <span class="p">}</span>
        <span class="k">return</span> <span class="kc">nil</span>
<span class="p">}</span></code></pre></div>
<p>The algorithm has three key points:</p>
<ol>
<li>The <code>queue</code> - it will contain all of the nodes that the algorithm visits</li>
<li>Taking the first element of the <code>queue</code>, checking it for a match, and
proceeding with the next nodes if no match is found</li>
<li><code>Queue</code>ing up all of the children nodes for a node before moving on in the
<code>queue</code></li>
</ol>
<p>Essentially, the whole algorithm revolves around pushing children nodes on a
queue and inspecting the nodes that are queued up. Of course, if a match is not
found at the end we return <code>nil</code> instead of a pointer to a <code>Node</code>.</p>
<h3 id="depth-first-search-" class="relative group">Depth-first search ⬇ <span class="absolute top-0 w-6 transition-opacity opacity-0 ltr:-left-6 rtl:-right-6 not-prose group-hover:opacity-100"><a class="group-hover:text-primary-300 dark:group-hover:text-neutral-700" style="text-decoration-line: none !important;" href="#depth-first-search-" aria-label="Anchor">#</a></span></h3>
<p>For completeness sake, let&rsquo;s also see how DFS would work.</p>
<p>As we stated earlier, the depth-first search will go first in depth by visiting
as many nodes as possible until it reaches a leaf. When then happens, it will
backtrack and find another branch on the tree to drill down on.</p>
<p>Let&rsquo;s see what that means visually:</p>
<p><img src="/golang-datastructures-trees/mydom-tree-dfs-steps.png" alt="" />
</p>
<p>If this is confusing to you, worry not - I&rsquo;ve added a bit more granularity in
the steps to aid my explanation.</p>
<p>The algorithm starts off just like BFS - it walks down from <code>html</code> to <code>body</code>
and to <code>div</code>. Then, instead of continuing to <code>h1</code>, it takes another step to the
leaf <code>span</code>. Once it figures out that <code>span</code> is a leaf, it will move back up to
<code>div</code> to find other branches to explore. Since it won&rsquo;t find any, it will move
back to <code>body</code> to find new branches proceeding to visit <code>h1</code>. Then, it will do
the same exercise again - go back to <code>body</code> and find that there&rsquo;s another
branch to explore - ultimately visiting <code>p</code> and the <code>img</code> nodes.</p>
<p>If you&rsquo;re wondering something along the lines of &ldquo;how can we go back up to the
parent without having a pointer to it&rdquo;, then you&rsquo;re forgetting one of the
oldest tricks in the book - recursion. Let&rsquo;s see a simple recursive Go
implementation of the algorithm:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="kd">func</span> <span class="nf">findByIdDFS</span><span class="p">(</span><span class="nx">node</span> <span class="o">*</span><span class="nx">Node</span><span class="p">,</span> <span class="nx">id</span> <span class="kt">string</span><span class="p">)</span> <span class="o">*</span><span class="nx">Node</span> <span class="p">{</span>
        <span class="k">if</span> <span class="nx">node</span><span class="p">.</span><span class="nx">id</span> <span class="o">==</span> <span class="nx">id</span> <span class="p">{</span>
                <span class="k">return</span> <span class="nx">node</span>
        <span class="p">}</span>

<span class="k">if</span> <span class="nb">len</span><span class="p">(</span><span class="nx">node</span><span class="p">.</span><span class="nx">children</span><span class="p">)</span> <span class="p">&gt;</span> <span class="mi">0</span> <span class="p">{</span>
                <span class="k">for</span> <span class="nx">_</span><span class="p">,</span> <span class="nx">child</span> <span class="o">:=</span> <span class="k">range</span> <span class="nx">node</span><span class="p">.</span><span class="nx">children</span> <span class="p">{</span>
                        <span class="nf">findByIdDFS</span><span class="p">(</span><span class="nx">child</span><span class="p">,</span> <span class="nx">id</span><span class="p">)</span>
                <span class="p">}</span>
        <span class="p">}</span>
        <span class="k">return</span> <span class="kc">nil</span>
<span class="p">}</span></code></pre></div>
<figure class="imagecaption">
<img class="caption" src="/golang-datastructures-trees/people_search.png" caption="" alt="People search graphic">
<span class="caption-text" style="font-size: 0.75em;display: table;margin: 0 auto 1.5em;"></span>
</figure>
<h2 id="finding-by-class-name-" class="relative group">Finding by class name 🔎 <span class="absolute top-0 w-6 transition-opacity opacity-0 ltr:-left-6 rtl:-right-6 not-prose group-hover:opacity-100"><a class="group-hover:text-primary-300 dark:group-hover:text-neutral-700" style="text-decoration-line: none !important;" href="#finding-by-class-name-" aria-label="Anchor">#</a></span></h2>
<p>Another functionality MyDOM <sup>(TM)</sup> should have is the ability to find
nodes by a class name. Essentially, when a JavaScript script executes
<code>getElementsByClassName</code>, MyDOM should know how to collect all nodes with a
certain class.</p>
<p>As you can imagine, this is also an algorithm that would have to explore the
whole MyDOM <sup>(TM)</sup> tree and pick up the nodes that satisfy certain
conditions.</p>
<p>To make our lives easier, let&rsquo;s first implement a function that a <code>Node</code> can
receive, called <code>hasClass</code>:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="kd">func</span> <span class="p">(</span><span class="nx">n</span> <span class="o">*</span><span class="nx">Node</span><span class="p">)</span> <span class="nf">hasClass</span><span class="p">(</span><span class="nx">className</span> <span class="kt">string</span><span class="p">)</span> <span class="kt">bool</span> <span class="p">{</span>
        <span class="nx">classes</span> <span class="o">:=</span> <span class="nx">strings</span><span class="p">.</span><span class="nf">Fields</span><span class="p">(</span><span class="nx">n</span><span class="p">.</span><span class="nx">classes</span><span class="p">)</span>
        <span class="k">for</span> <span class="nx">_</span><span class="p">,</span> <span class="nx">class</span> <span class="o">:=</span> <span class="k">range</span> <span class="nx">classes</span> <span class="p">{</span>
                <span class="k">if</span> <span class="nx">class</span> <span class="o">==</span> <span class="nx">className</span> <span class="p">{</span>
                        <span class="k">return</span> <span class="kc">true</span>
                <span class="p">}</span>
        <span class="p">}</span>
        <span class="k">return</span> <span class="kc">false</span>
<span class="p">}</span></code></pre></div>
<p><code>hasClass</code> takes a <code>Node</code>&rsquo;s classes field, splits them on each space character
and then loops the slice of classes and tries to find the class name that we are
interested in. Let&rsquo;s write a couple of tests that will test this function:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="kd">type</span> <span class="nx">testcase</span> <span class="kd">struct</span> <span class="p">{</span>
        <span class="nx">className</span>      <span class="kt">string</span>
        <span class="nx">node</span>           <span class="nx">Node</span>
        <span class="nx">expectedResult</span> <span class="kt">bool</span>
<span class="p">}</span>

<span class="kd">func</span> <span class="nf">TestHasClass</span><span class="p">(</span><span class="nx">t</span> <span class="o">*</span><span class="nx">testing</span><span class="p">.</span><span class="nx">T</span><span class="p">)</span> <span class="p">{</span>
        <span class="nx">cases</span> <span class="o">:=</span> <span class="p">[]</span><span class="nx">testcase</span><span class="p">{</span>
                <span class="nx">testcase</span><span class="p">{</span>
                        <span class="nx">className</span><span class="p">:</span>      <span class="s">&#34;foo&#34;</span><span class="p">,</span>
                        <span class="nx">node</span><span class="p">:</span>           <span class="nx">Node</span><span class="p">{</span><span class="nx">classes</span><span class="p">:</span> <span class="s">&#34;foo bar&#34;</span><span class="p">},</span>
                        <span class="nx">expectedResult</span><span class="p">:</span> <span class="kc">true</span><span class="p">,</span>
                <span class="p">},</span>
                <span class="nx">testcase</span><span class="p">{</span>
                        <span class="nx">className</span><span class="p">:</span>      <span class="s">&#34;foo&#34;</span><span class="p">,</span>
                        <span class="nx">node</span><span class="p">:</span>           <span class="nx">Node</span><span class="p">{</span><span class="nx">classes</span><span class="p">:</span> <span class="s">&#34;bar baz qux&#34;</span><span class="p">},</span>
                        <span class="nx">expectedResult</span><span class="p">:</span> <span class="kc">false</span><span class="p">,</span>
                <span class="p">},</span>
                <span class="nx">testcase</span><span class="p">{</span>
                        <span class="nx">className</span><span class="p">:</span>      <span class="s">&#34;bar&#34;</span><span class="p">,</span>
                        <span class="nx">node</span><span class="p">:</span>           <span class="nx">Node</span><span class="p">{</span><span class="nx">classes</span><span class="p">:</span> <span class="s">&#34;&#34;</span><span class="p">},</span>
                        <span class="nx">expectedResult</span><span class="p">:</span> <span class="kc">false</span><span class="p">,</span>
                <span class="p">},</span>
        <span class="p">}</span>

<span class="k">for</span> <span class="nx">_</span><span class="p">,</span> <span class="k">case</span> <span class="o">:=</span> <span class="k">range</span> <span class="nx">cases</span> <span class="p">{</span>
                <span class="nx">result</span> <span class="o">:=</span> <span class="k">case</span><span class="p">.</span><span class="nx">node</span><span class="p">.</span><span class="nf">hasClass</span><span class="p">(</span><span class="nx">test</span><span class="p">.</span><span class="nx">className</span><span class="p">)</span>
                <span class="k">if</span> <span class="nx">result</span> <span class="o">!=</span> <span class="k">case</span><span class="p">.</span><span class="nx">expectedResult</span> <span class="p">{</span>
                        <span class="nx">t</span><span class="p">.</span><span class="nf">Error</span><span class="p">(</span>
                                <span class="s">&#34;For node&#34;</span><span class="p">,</span> <span class="k">case</span><span class="p">.</span><span class="nx">node</span><span class="p">,</span>
                                <span class="s">&#34;and class&#34;</span><span class="p">,</span> <span class="k">case</span><span class="p">.</span><span class="nx">className</span><span class="p">,</span>
                                <span class="s">&#34;expected&#34;</span><span class="p">,</span> <span class="k">case</span><span class="p">.</span><span class="nx">expectedResult</span><span class="p">,</span>
                                <span class="s">&#34;got&#34;</span><span class="p">,</span> <span class="nx">result</span><span class="p">,</span>
                        <span class="p">)</span>
                <span class="p">}</span>
        <span class="p">}</span>
<span class="p">}</span></code></pre></div>
<p>As you can see, the <code>hasClass</code> function will detect if a class name is in the
list of classes on a <code>Node</code>. Now, let&rsquo;s move on to implementing MyDOM&rsquo;s
implementation of finding all <code>Node</code>s by class name:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="kd">func</span> <span class="nf">findAllByClassName</span><span class="p">(</span><span class="nx">root</span> <span class="o">*</span><span class="nx">Node</span><span class="p">,</span> <span class="nx">className</span> <span class="kt">string</span><span class="p">)</span> <span class="p">[]</span><span class="o">*</span><span class="nx">Node</span> <span class="p">{</span>
        <span class="nx">result</span> <span class="o">:=</span> <span class="nb">make</span><span class="p">([]</span><span class="o">*</span><span class="nx">Node</span><span class="p">,</span> <span class="mi">0</span><span class="p">)</span>
        <span class="nx">queue</span> <span class="o">:=</span> <span class="nb">make</span><span class="p">([]</span><span class="o">*</span><span class="nx">Node</span><span class="p">,</span> <span class="mi">0</span><span class="p">)</span>
        <span class="nx">queue</span> <span class="p">=</span> <span class="nb">append</span><span class="p">(</span><span class="nx">queue</span><span class="p">,</span> <span class="nx">root</span><span class="p">)</span>
        <span class="k">for</span> <span class="nb">len</span><span class="p">(</span><span class="nx">queue</span><span class="p">)</span> <span class="p">&gt;</span> <span class="mi">0</span> <span class="p">{</span>
                <span class="nx">nextUp</span> <span class="o">:=</span> <span class="nx">queue</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span>
                <span class="nx">queue</span> <span class="p">=</span> <span class="nx">queue</span><span class="p">[</span><span class="mi">1</span><span class="p">:]</span>
                <span class="k">if</span> <span class="nx">nextUp</span><span class="p">.</span><span class="nf">hasClass</span><span class="p">(</span><span class="nx">className</span><span class="p">)</span> <span class="p">{</span>
                        <span class="nx">result</span> <span class="p">=</span> <span class="nb">append</span><span class="p">(</span><span class="nx">result</span><span class="p">,</span> <span class="nx">nextUp</span><span class="p">)</span>
                <span class="p">}</span>
                <span class="k">if</span> <span class="nb">len</span><span class="p">(</span><span class="nx">nextUp</span><span class="p">.</span><span class="nx">children</span><span class="p">)</span> <span class="p">&gt;</span> <span class="mi">0</span> <span class="p">{</span>
                        <span class="k">for</span> <span class="nx">_</span><span class="p">,</span> <span class="nx">child</span> <span class="o">:=</span> <span class="k">range</span> <span class="nx">nextUp</span><span class="p">.</span><span class="nx">children</span> <span class="p">{</span>
                                <span class="nx">queue</span> <span class="p">=</span> <span class="nb">append</span><span class="p">(</span><span class="nx">queue</span><span class="p">,</span> <span class="nx">child</span><span class="p">)</span>
                        <span class="p">}</span>
                <span class="p">}</span>
        <span class="p">}</span>
        <span class="k">return</span> <span class="nx">result</span>
<span class="p">}</span></code></pre></div>
<p>If the algorithm seems familiar, that&rsquo;s because you&rsquo;re looking at a modified
<code>findById</code> function. <code>findAllByClassName</code> works just like <code>findById</code>, but
instead of <code>return</code>ing the moment it finds a match, it will just append the
matched <code>Node</code> to the <code>result</code> slice. It will continue doing that until all of
the <code>Node</code>s have been visited.</p>
<p>If there are no matches, the <code>result</code> slice will be empty. If there are any
matches, they will be returned as part of the <code>result</code> slice.</p>
<p>Last thing worth mentioning is that to traverse the tree we used a
Breadth-first approach here - the algorithm uses a queue for each of the
<code>Node</code>s and loops over them while appending to the <code>result</code> slice if a match is
found.</p>
<figure class="imagecaption">
<img class="caption" src="/golang-datastructures-trees/cancel.png" caption="" alt="Delete graphic">
<span class="caption-text" style="font-size: 0.75em;display: table;margin: 0 auto 1.5em;"></span>
</figure>
<h2 id="deleting-nodes-" class="relative group">Deleting nodes 🗑 <span class="absolute top-0 w-6 transition-opacity opacity-0 ltr:-left-6 rtl:-right-6 not-prose group-hover:opacity-100"><a class="group-hover:text-primary-300 dark:group-hover:text-neutral-700" style="text-decoration-line: none !important;" href="#deleting-nodes-" aria-label="Anchor">#</a></span></h2>
<p>Another functionality that is often used in the DOM is the ability to remove
nodes. Just like the DOM can do it, also our MyDOM <sup>(TM)</sup> should be
able to handle such operations.</p>
<p>The simplest way to do this operation in JavaScript is:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-javascript" data-lang="javascript"><span class="kd">var</span> <span class="nx">el</span> <span class="o">=</span> <span class="nb">document</span><span class="p">.</span><span class="nx">getElementById</span><span class="p">(</span><span class="s1">&#39;foo&#39;</span><span class="p">);</span>
<span class="nx">el</span><span class="p">.</span><span class="nx">remove</span><span class="p">();</span></code></pre></div>
<p>While our <code>document</code> knows how to handle <code>getElementById</code> (by calling
<code>findById</code> under the hood), our <code>Node</code>s do not know how to handle a <code>remove</code>
function. Removing a <code>Node</code> from the MyDOM <sup>(TM)</sup> tree would be a
two-step process:</p>
<ol>
<li>We have to look up to the <code>parent</code> of the <code>Node</code> and remove it from its
parent&rsquo;s <code>children</code> collection;</li>
<li>If the to-be-removed <code>Node</code> has any children, we have to remove those from
the DOM. This means we have to remove all pointers to each of the children
and its parent (the node to-be-removed) so Go&rsquo;s garbage collector can free
up that memory.</li>
</ol>
<p>And here&rsquo;s a simple way to achieve that:</p>
<div class="highlight"><pre tabindex="0" class="chroma"><code class="language-go" data-lang="go"><span class="kd">func</span> <span class="p">(</span><span class="nx">node</span> <span class="o">*</span><span class="nx">Node</span><span class="p">)</span> <span class="nf">remove</span><span class="p">()</span> <span class="p">{</span>
        <span class="c1">// Remove the node from it&#39;s parents children collection
</span><span class="c1"></span>        <span class="k">for</span> <span class="nx">idx</span><span class="p">,</span> <span class="nx">sibling</span> <span class="o">:=</span> <span class="k">range</span> <span class="nx">n</span><span class="p">.</span><span class="nx">parent</span><span class="p">.</span><span class="nx">children</span> <span class="p">{</span>
                <span class="k">if</span> <span class="nx">sibling</span> <span class="o">==</span> <span class="nx">node</span> <span class="p">{</span>
                        <span class="nx">node</span><span class="p">.</span><span class="nx">parent</span><span class="p">.</span><span class="nx">children</span> <span class="p">=</span> <span class="nb">append</span><span class="p">(</span>
                                <span class="nx">node</span><span class="p">.</span><span class="nx">parent</span><span class="p">.</span><span class="nx">children</span><span class="p">[:</span><span class="nx">idx</span><span class="p">],</span>
                                <span class="nx">node</span><span class="p">.</span><span class="nx">parent</span><span class="p">.</span><span class="nx">children</span><span class="p">[</span><span class="nx">idx</span><span class="o">+</span><span class="mi">1</span><span class="p">:]</span><span class="o">...</span><span class="p">,</span>
                        <span class="p">)</span>
                <span class="p">}</span>
        <span class="p">}</span>

        <span class="c1">// If the node has any children, set their parent to nil and set the node&#39;s children collection to nil
</span><span class="c1"></span>        <span class="k">if</span> <span class="nb">len</span><span class="p">(</span><span class="nx">node</span><span class="p">.</span><span class="nx">children</span><span class="p">)</span> <span class="o">!=</span> <span class="mi">0</span> <span class="p">{</span>
                <span class="k">for</span> <span class="nx">_</span><span class="p">,</span> <span class="nx">child</span> <span class="o">:=</span> <span class="k">range</span> <span class="nx">node</span><span class="p">.</span><span class="nx">children</span> <span class="p">{</span>
                        <span class="nx">child</span><span class="p">.</span><span class="nx">parent</span> <span class="p">=</span> <span class="kc">nil</span>
                <span class="p">}</span>
                <span class="nx">node</span><span class="p">.</span><span class="nx">children</span> <span class="p">=</span> <span class="kc">nil</span>
        <span class="p">}</span>
<span class="p">}</span></code></pre></div>
<p>A <code>*Node</code> would have a <code>remove</code> function, which does the two-step process of
the <code>Node</code>&rsquo;s removal.</p>
<p>In the first step, we take the node out of the <code>parent</code>&rsquo;s children list, by
looping over them and removing the node by appending the elements before the
node in the list, and the elements after the node.</p>
<p>In the second step, after checking for the presence of any children on the
node, we remove the reference to the <code>parent</code> from all the children and then we
set the <code>Node</code>&rsquo;s children to <code>nil</code>.</p>
