<h1 class="mt-0 text-4xl font-extrabold text-neutral-900 dark:text-neutral">
Golang Datastructures: Trees
</h1>
<div class="mt-8 mb-12 text-base text-neutral-500 dark:text-neutral-400">
<div class="flex flex-row flex-wrap items-center">
<time datetime="2019-02-13 00:00:00 &#43;0000 UTC">13 February 2019</time><span class="px-2 text-primary-500">&middot;</span><span title="Reading time">14 mins</span>
</div>
</div>
</header>
<section class="flex flex-col max-w-full mt-0 prose lg:flex-row dark:prose-invert">
<div class="order-first px-0 lg:max-w-xs ltr:lg:pl-8 rtl:lg:pr-8 lg:order-last">
<div class="ltr:pl-5 rtl:pr-5 toc lg:sticky lg:top-10">
<details open class="mt-0 overflow-hidden rounded-lg rtl:pr-5 ltr:pl-5 ltr:-ml-5 rtl:-mr-5 lg:mt-3">
<summary class="block py-1 text-lg font-semibold cursor-pointer rtl:pr-5 ltr:pl-5 ltr:-ml-5 rtl:-mr-5 text-neutral-800 dark:text-neutral-100 lg:hidden bg-neutral-100 dark:bg-neutral-700">
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
<h2 id="where-to-next" class="relative group">Where to next? <span class="absolute top-0 w-6 transition-opacity opacity-0 ltr:-left-6 rtl:-right-6 not-prose group-hover:opacity-100"><a class="group-hover:text-primary-300 dark:group-hover:text-neutral-700" style="text-decoration-line: none !important;" href="#where-to-next" aria-label="Anchor">#</a></span></h2>
<p>Obviously, our MyDOM <sup>(TM)</sup> implementation is never going to become a
replacement for the DOM. But, I believe that it&rsquo;s an interesting example that
can help you learn and it&rsquo;s pretty interesting problem to think about. We
interact with browsers every day, so thinking how they could function under the
hood is an interesting exercise.</p>
<p>If you would like to play with our tree structure and write more functionality,
you can head over to WC3&rsquo;s JavaScript HTML DOM Document
<a href="https://www.w3schools.com/js/js_htmldom_document.asp">documentation</a> and think
about adding more functionality to MyDOM.</p>
<p>Obviously, the idea behind this article was to learn more about trees (graphs)
and learn about the popular searching/traversal algorithms that are used out
there. But, by all means, please keep on exploring and experimenting and drop
me a comment about what improvements you did to your MyDOM implementation.</p>
<div id="revue-embed">
<form action="https://www.getrevue.co/profile/itsilija/add_subscriber" method="post" id="revue-form" name="revue-form" target="_blank">
<p><b>Liked this article?</b>
Subscribe to my newsletter and get future articles in your inbox. It's a
short and sweet read, going out to hundreds of other engineers.</p>
<div class="revue-form-group">
<input class="revue-form-field" placeholder="Your email address..." type="email" name="member[email]" id="member_email">
</div>
<div class="revue-form-actions">
<input type="submit" value="Subscribe" name="member[subscribe]" id="member_submit">
</div>
</form>
</div>
</div>
</section>
<footer class="pt-8 max-w-prose">
<div class="flex">
<img class="w-24 h-24 !mt-0 !mb-0 ltr:mr-4 rtl:ml-4 rounded-full" width="96" height="96" alt="Author" src="/img/avatar_huae39953c6189b6aa60e47103b358d088_112612_192x192_fill_q75_box_smart1.jpg" />
<div class="place-self-center">
<div class="text-[0.6rem] leading-3 text-neutral-500 dark:text-neutral-400 uppercase">
Author
</div>
<div class="font-semibold leading-6 text-neutral-800 dark:text-neutral-300">
Ilija Eftimov
</div>
<div class="text-2xl sm:text-lg">
<div class="flex flex-wrap text-neutral-400 dark:text-neutral-500">
<a class="px-1 hover:text-primary-700 dark:hover:text-primary-400" href="https://twitter.com/efilija" target="_blank" aria-label="Twitter" rel="me noopener noreferrer">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="twitter" class="svg-inline--fa fa-twitter fa-w-16" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M459.37 151.716c.325 4.548.325 9.097.325 13.645 0 138.72-105.583 298.558-298.558 298.558-59.452 0-114.68-17.219-161.137-47.106 8.447.974 16.568 1.299 25.34 1.299 49.055 0 94.213-16.568 130.274-44.832-46.132-.975-84.792-31.188-98.112-72.772 6.498.974 12.995 1.624 19.818 1.624 9.421 0 18.843-1.3 27.614-3.573-48.081-9.747-84.143-51.98-84.143-102.985v-1.299c13.969 7.797 30.214 12.67 47.431 13.319-28.264-18.843-46.781-51.005-46.781-87.391 0-19.492 5.197-37.36 14.294-52.954 51.655 63.675 129.3 105.258 216.365 109.807-1.624-7.797-2.599-15.918-2.599-24.04 0-57.828 46.782-104.934 104.934-104.934 30.213 0 57.502 12.67 76.67 33.137 23.715-4.548 46.456-13.32 66.599-25.34-7.798 24.366-24.366 44.833-46.132 57.827 21.117-2.273 41.584-8.122 60.426-16.243-14.292 20.791-32.161 39.308-52.628 54.253z"></path></svg>
</span>
</a
        >
<a class="px-1 hover:text-primary-700 dark:hover:text-primary-400" href="https://www.linkedin.com/in/ieftimov/" target="_blank" aria-label="Linkedin" rel="me noopener noreferrer">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="linkedin" class="svg-inline--fa fa-linkedin fa-w-14" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path fill="currentColor" d="M416 32H31.9C14.3 32 0 46.5 0 64.3v383.4C0 465.5 14.3 480 31.9 480H416c17.6 0 32-14.5 32-32.3V64.3c0-17.8-14.4-32.3-32-32.3zM135.4 416H69V202.2h66.5V416zm-33.2-243c-21.3 0-38.5-17.3-38.5-38.5S80.9 96 102.2 96c21.2 0 38.5 17.3 38.5 38.5 0 21.3-17.2 38.5-38.5 38.5zm282.1 243h-66.4V312c0-24.8-.5-56.7-34.5-56.7-34.6 0-39.9 27-39.9 54.9V416h-66.4V202.2h63.7v29.2h.9c8.9-16.8 30.6-34.5 62.9-34.5 67.2 0 79.7 44.3 79.7 101.9V416z"></path></svg>
</span>
</a
        >
<a class="px-1 hover:text-primary-700 dark:hover:text-primary-400" href="/cdn-cgi/l/email-protection#2d4f41424a6d44484b594440425b034e4240" target="_blank" aria-label="Email" rel="me noopener noreferrer">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="at" class="svg-inline--fa fa-at fa-w-16" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M256 8C118.941 8 8 118.919 8 256c0 137.059 110.919 248 248 248 48.154 0 95.342-14.14 135.408-40.223 12.005-7.815 14.625-24.288 5.552-35.372l-10.177-12.433c-7.671-9.371-21.179-11.667-31.373-5.129C325.92 429.757 291.314 440 256 440c-101.458 0-184-82.542-184-184S154.542 72 256 72c100.139 0 184 57.619 184 160 0 38.786-21.093 79.742-58.17 83.693-17.349-.454-16.91-12.857-13.476-30.024l23.433-121.11C394.653 149.75 383.308 136 368.225 136h-44.981a13.518 13.518 0 0 0-13.432 11.993l-.01.092c-14.697-17.901-40.448-21.775-59.971-21.775-74.58 0-137.831 62.234-137.831 151.46 0 65.303 36.785 105.87 96 105.87 26.984 0 57.369-15.637 74.991-38.333 9.522 34.104 40.613 34.103 70.71 34.103C462.609 379.41 504 307.798 504 232 504 95.653 394.023 8 256 8zm-21.68 304.43c-22.249 0-36.07-15.623-36.07-40.771 0-44.993 30.779-72.729 58.63-72.729 22.292 0 35.601 15.241 35.601 40.77 0 45.061-33.875 72.73-58.161 72.73z"></path></svg>
</span>
</a
        >
<a class="px-1 hover:text-primary-700 dark:hover:text-primary-400" href="https://dev.to/fteem" target="_blank" aria-label="Dev" rel="me noopener noreferrer">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="dev" class="svg-inline--fa fa-dev fa-w-14" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path fill="currentColor" d="M120.12 208.29c-3.88-2.9-7.77-4.35-11.65-4.35H91.03v104.47h17.45c3.88 0 7.77-1.45 11.65-4.35 3.88-2.9 5.82-7.25 5.82-13.06v-69.65c-.01-5.8-1.96-10.16-5.83-13.06zM404.1 32H43.9C19.7 32 .06 51.59 0 75.8v360.4C.06 460.41 19.7 480 43.9 480h360.2c24.21 0 43.84-19.59 43.9-43.8V75.8c-.06-24.21-19.7-43.8-43.9-43.8zM154.2 291.19c0 18.81-11.61 47.31-48.36 47.25h-46.4V172.98h47.38c35.44 0 47.36 28.46 47.37 47.28l.01 70.93zm100.68-88.66H201.6v38.42h32.57v29.57H201.6v38.41h53.29v29.57h-62.18c-11.16.29-20.44-8.53-20.72-19.69V193.7c-.27-11.15 8.56-20.41 19.71-20.69h63.19l-.01 29.52zm103.64 115.29c-13.2 30.75-36.85 24.63-47.44 0l-38.53-144.8h32.57l29.71 113.72 29.57-113.72h32.58l-38.46 144.8z"></path></svg>
</span>
</a
        >
<a class="px-1 hover:text-primary-700 dark:hover:text-primary-400" href="https://github.com/fteem" target="_blank" aria-label="Github" rel="me noopener noreferrer">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="github" class="svg-inline--fa fa-github fa-w-16" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 496 512"><path fill="currentColor" d="M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3.3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5.3-6.2 2.3zm44.2-1.7c-2.9.7-4.9 2.6-4.6 4.9.3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3.7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3.3 2.9 2.3 3.9 1.6 1 3.6.7 4.3-.7.7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3.7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5 1 1.3-1.3.7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-1zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z"></path></svg>
</span>
</a
        >
<a class="px-1 hover:text-primary-700 dark:hover:text-primary-400" href="https://t.me/ieftimovcom" target="_blank" aria-label="Telegram" rel="me noopener noreferrer">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="telegram-plane" class="svg-inline--fa fa-telegram-plane fa-w-14" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path fill="currentColor" d="M446.7 98.6l-67.6 318.8c-5.1 22.5-18.4 28.1-37.3 17.5l-103-75.9-49.7 47.8c-5.5 5.5-10.1 10.1-20.7 10.1l7.4-104.9 190.9-172.5c8.3-7.4-1.8-11.5-12.9-4.1L117.8 284 16.2 252.2c-22.1-6.9-22.5-22.1 4.6-32.7L418.2 66.4c18.4-6.9 34.5 4.1 28.5 32.2z"></path></svg>
</span>
</a
        >
</div>
</div>
</div>
</div>
<section class="flex flex-row flex-wrap justify-center pt-4 text-xl">
<a class="bg-neutral-300 text-neutral-700 dark:bg-neutral-700 dark:text-neutral-300 dark:hover:bg-primary-400 dark:hover:text-neutral-800 m-1 hover:bg-primary-500 hover:text-neutral rounded min-w-[2.4rem] inline-block text-center p-1" href="https://www.facebook.com/sharer/sharer.php?u=https://ieftimov.com/posts/golang-datastructures-trees/&amp;quote=Golang%20Datastructures:%20Trees" title="Share on Facebook" aria-label="Share on Facebook">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="facebook" class="svg-inline--fa fa-facebook fa-w-16" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M504 256C504 119 393 8 256 8S8 119 8 256c0 123.78 90.69 226.38 209.25 245V327.69h-63V256h63v-54.64c0-62.15 37-96.48 93.67-96.48 27.14 0 55.52 4.84 55.52 4.84v61h-31.28c-30.8 0-40.41 19.12-40.41 38.73V256h68.78l-11 71.69h-57.78V501C413.31 482.38 504 379.78 504 256z"></path></svg>
</span>
</a
        >
<a class="bg-neutral-300 text-neutral-700 dark:bg-neutral-700 dark:text-neutral-300 dark:hover:bg-primary-400 dark:hover:text-neutral-800 m-1 hover:bg-primary-500 hover:text-neutral rounded min-w-[2.4rem] inline-block text-center p-1" href="https://twitter.com/intent/tweet/?url=https://ieftimov.com/posts/golang-datastructures-trees/&amp;text=Golang%20Datastructures:%20Trees" title="Tweet on Twitter" aria-label="Tweet on Twitter">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="twitter" class="svg-inline--fa fa-twitter fa-w-16" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M459.37 151.716c.325 4.548.325 9.097.325 13.645 0 138.72-105.583 298.558-298.558 298.558-59.452 0-114.68-17.219-161.137-47.106 8.447.974 16.568 1.299 25.34 1.299 49.055 0 94.213-16.568 130.274-44.832-46.132-.975-84.792-31.188-98.112-72.772 6.498.974 12.995 1.624 19.818 1.624 9.421 0 18.843-1.3 27.614-3.573-48.081-9.747-84.143-51.98-84.143-102.985v-1.299c13.969 7.797 30.214 12.67 47.431 13.319-28.264-18.843-46.781-51.005-46.781-87.391 0-19.492 5.197-37.36 14.294-52.954 51.655 63.675 129.3 105.258 216.365 109.807-1.624-7.797-2.599-15.918-2.599-24.04 0-57.828 46.782-104.934 104.934-104.934 30.213 0 57.502 12.67 76.67 33.137 23.715-4.548 46.456-13.32 66.599-25.34-7.798 24.366-24.366 44.833-46.132 57.827 21.117-2.273 41.584-8.122 60.426-16.243-14.292 20.791-32.161 39.308-52.628 54.253z"></path></svg>
</span>
</a
        >
<a class="bg-neutral-300 text-neutral-700 dark:bg-neutral-700 dark:text-neutral-300 dark:hover:bg-primary-400 dark:hover:text-neutral-800 m-1 hover:bg-primary-500 hover:text-neutral rounded min-w-[2.4rem] inline-block text-center p-1" href="https://reddit.com/submit/?url=https://ieftimov.com/posts/golang-datastructures-trees/&amp;resubmit=true&amp;title=Golang%20Datastructures:%20Trees" title="Submit to Reddit" aria-label="Submit to Reddit">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="reddit" class="svg-inline--fa fa-reddit fa-w-16" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M201.5 305.5c-13.8 0-24.9-11.1-24.9-24.6 0-13.8 11.1-24.9 24.9-24.9 13.6 0 24.6 11.1 24.6 24.9 0 13.6-11.1 24.6-24.6 24.6zM504 256c0 137-111 248-248 248S8 393 8 256 119 8 256 8s248 111 248 248zm-132.3-41.2c-9.4 0-17.7 3.9-23.8 10-22.4-15.5-52.6-25.5-86.1-26.6l17.4-78.3 55.4 12.5c0 13.6 11.1 24.6 24.6 24.6 13.8 0 24.9-11.3 24.9-24.9s-11.1-24.9-24.9-24.9c-9.7 0-18 5.8-22.1 13.8l-61.2-13.6c-3-.8-6.1 1.4-6.9 4.4l-19.1 86.4c-33.2 1.4-63.1 11.3-85.5 26.8-6.1-6.4-14.7-10.2-24.1-10.2-34.9 0-46.3 46.9-14.4 62.8-1.1 5-1.7 10.2-1.7 15.5 0 52.6 59.2 95.2 132 95.2 73.1 0 132.3-42.6 132.3-95.2 0-5.3-.6-10.8-1.9-15.8 31.3-16 19.8-62.5-14.9-62.5zM302.8 331c-18.2 18.2-76.1 17.9-93.6 0-2.2-2.2-6.1-2.2-8.3 0-2.5 2.5-2.5 6.4 0 8.6 22.8 22.8 87.3 22.8 110.2 0 2.5-2.2 2.5-6.1 0-8.6-2.2-2.2-6.1-2.2-8.3 0zm7.7-75c-13.6 0-24.6 11.1-24.6 24.9 0 13.6 11.1 24.6 24.6 24.6 13.8 0 24.9-11.1 24.9-24.6 0-13.8-11-24.9-24.9-24.9z"></path></svg>
</span>
</a
        >
<a class="bg-neutral-300 text-neutral-700 dark:bg-neutral-700 dark:text-neutral-300 dark:hover:bg-primary-400 dark:hover:text-neutral-800 m-1 hover:bg-primary-500 hover:text-neutral rounded min-w-[2.4rem] inline-block text-center p-1" href="https://www.linkedin.com/shareArticle?mini=true&amp;url=https://ieftimov.com/posts/golang-datastructures-trees/&amp;title=Golang%20Datastructures:%20Trees" title="Share on LinkedIn" aria-label="Share on LinkedIn">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fab" data-icon="linkedin" class="svg-inline--fa fa-linkedin fa-w-14" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path fill="currentColor" d="M416 32H31.9C14.3 32 0 46.5 0 64.3v383.4C0 465.5 14.3 480 31.9 480H416c17.6 0 32-14.5 32-32.3V64.3c0-17.8-14.4-32.3-32-32.3zM135.4 416H69V202.2h66.5V416zm-33.2-243c-21.3 0-38.5-17.3-38.5-38.5S80.9 96 102.2 96c21.2 0 38.5 17.3 38.5 38.5 0 21.3-17.2 38.5-38.5 38.5zm282.1 243h-66.4V312c0-24.8-.5-56.7-34.5-56.7-34.6 0-39.9 27-39.9 54.9V416h-66.4V202.2h63.7v29.2h.9c8.9-16.8 30.6-34.5 62.9-34.5 67.2 0 79.7 44.3 79.7 101.9V416z"></path></svg>
</span>
</a
        >
<a class="bg-neutral-300 text-neutral-700 dark:bg-neutral-700 dark:text-neutral-300 dark:hover:bg-primary-400 dark:hover:text-neutral-800 m-1 hover:bg-primary-500 hover:text-neutral rounded min-w-[2.4rem] inline-block text-center p-1" href="/cdn-cgi/l/email-protection#142b767b706d297c606064672e3b3b7d7172607d797b623a777b793b647b6760673b737b78757a733970756075676066617760616671673960667171673b327579642f6761767e71776029537b78757a7331262450756075676066617760616671672e3126244066717167" title="Send via email" aria-label="Send via email">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="at" class="svg-inline--fa fa-at fa-w-16" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M256 8C118.941 8 8 118.919 8 256c0 137.059 110.919 248 248 248 48.154 0 95.342-14.14 135.408-40.223 12.005-7.815 14.625-24.288 5.552-35.372l-10.177-12.433c-7.671-9.371-21.179-11.667-31.373-5.129C325.92 429.757 291.314 440 256 440c-101.458 0-184-82.542-184-184S154.542 72 256 72c100.139 0 184 57.619 184 160 0 38.786-21.093 79.742-58.17 83.693-17.349-.454-16.91-12.857-13.476-30.024l23.433-121.11C394.653 149.75 383.308 136 368.225 136h-44.981a13.518 13.518 0 0 0-13.432 11.993l-.01.092c-14.697-17.901-40.448-21.775-59.971-21.775-74.58 0-137.831 62.234-137.831 151.46 0 65.303 36.785 105.87 96 105.87 26.984 0 57.369-15.637 74.991-38.333 9.522 34.104 40.613 34.103 70.71 34.103C462.609 379.41 504 307.798 504 232 504 95.653 394.023 8 256 8zm-21.68 304.43c-22.249 0-36.07-15.623-36.07-40.771 0-44.993 30.779-72.729 58.63-72.729 22.292 0 35.601 15.241 35.601 40.77 0 45.061-33.875 72.73-58.161 72.73z"></path></svg>
</span>
</a
        >
</section>
<div class="pt-8">
<hr class="border-dotted border-neutral-300 dark:border-neutral-600" />
<div class="flex justify-between pt-3">
<span>
<a class="flex group" href="/posts/otp-elixir-genserver-build-own-url-shortener/">
<span class="mr-3 ltr:inline rtl:hidden text-neutral-700 dark:text-neutral group-hover:text-primary-600 dark:group-hover:text-primary-400">&larr;</span
              >
<span class="ml-3 ltr:hidden rtl:inline text-neutral-700 dark:text-neutral group-hover:text-primary-600 dark:group-hover:text-primary-400">&rarr;</span
              >
<span class="flex flex-col">
<span class="mt-[0.1rem] leading-6 group-hover:underline group-hover:decoration-primary-500">OTP in Elixir: Learn GenServer by Building Your Own URL Shortener</span
                >
<span class="mt-[0.1rem] text-xs text-neutral-500 dark:text-neutral-400">
<time datetime="2019-01-26 00:00:00 &#43;0000 UTC">26 January 2019</time>
</span>
</span>
</a>
</span>
<span>
<a class="flex text-right group" href="/posts/when-why-least-frequently-used-cache-implementation-golang/">
<span class="flex flex-col">
<span class="mt-[0.1rem] leading-6 group-hover:underline group-hover:decoration-primary-500">When and Why to use a Least Frequently Used (LFU) cache with an implementation in Golang</span
                >
<span class="mt-[0.1rem] text-xs text-neutral-500 dark:text-neutral-400">
<time datetime="2019-02-27 00:00:00 &#43;0000 UTC">27 February 2019</time>
</span>
</span>
<span class="ml-3 ltr:inline rtl:hidden text-neutral-700 dark:text-neutral group-hover:text-primary-600 dark:group-hover:text-primary-400">&rarr;</span
              >
<span class="mr-3 ltr:hidden rtl:inline text-neutral-700 dark:text-neutral group-hover:text-primary-600 dark:group-hover:text-primary-400">&larr;</span
              >
</a>
</span>
</div>
</div>
</footer>
</article>
<div class="absolute top-[110vh] ltr:right-0 rtl:left-0 w-12 pointer-events-none bottom-[-5.5rem]">
<a href="#the-top" class="w-12 h-12 sticky pointer-events-auto top-[calc(100vh-5rem)] bg-neutral/50 dark:bg-neutral-800/50 backdrop-blur rounded-full text-xl flex items-center justify-center text-neutral-700 dark:text-neutral hover:text-primary-600 dark:hover:text-primary-400" aria-label="Scroll to top" title="Scroll to top">
&uarr;
</a>
</div>
</main>
<div id="search-wrapper" class="fixed inset-0 z-50 flex flex-col p-4 sm:p-6 md:p-[10vh] lg:p-[12vh] w-screen h-screen cursor-default bg-neutral-500/50 backdrop-blur-sm dark:bg-neutral-900/50 invisible" data-url="https://ieftimov.com/">
<div id="search-modal" class="flex flex-col w-full max-w-3xl min-h-0 mx-auto border rounded-md shadow-lg border-neutral-200 top-20 bg-neutral dark:bg-neutral-800 dark:border-neutral-700">
<header class="relative z-10 flex items-center justify-between flex-none px-2">
<form class="flex items-center flex-auto min-w-0">
<div class="flex items-center justify-center w-8 h-8 text-neutral-400">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" class="svg-inline--fa fa-search fa-w-16" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path></svg>
</span>
</div>
<input type="search" id="search-query" class="flex flex-auto h-12 mx-1 bg-transparent appearance-none focus:outline-dotted focus:outline-transparent focus:outline-2" placeholder="Search" tabindex="0" />
</form>
<button id="close-search-button" class="flex items-center justify-center w-8 h-8 text-neutral-700 dark:text-neutral hover:text-primary-600 dark:hover:text-primary-400" title="Close (Esc)">
<span class="relative inline-block align-text-bottom icon">
<svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="times" class="svg-inline--fa fa-times fa-w-11" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 352 512"><path fill="currentColor" d="M242.72 256l100.07-100.07c12.28-12.28 12.28-32.19 0-44.48l-22.24-22.24c-12.28-12.28-32.19-12.28-44.48 0L176 189.28 75.93 89.21c-12.28-12.28-32.19-12.28-44.48 0L9.21 111.45c-12.28 12.28-12.28 32.19 0 44.48L109.28 256 9.21 356.07c-12.28 12.28-12.28 32.19 0 44.48l22.24 22.24c12.28 12.28 32.2 12.28 44.48 0L176 322.72l100.07 100.07c12.28 12.28 32.2 12.28 44.48 0l22.24-22.24c12.28-12.28 12.28-32.19 0-44.48L242.72 256z"></path></svg>
</span>
</button>
</header>
<section class="flex-auto px-2 overflow-auto">
<ul id="search-results">
</ul>
</section>
</div>
</div>
<footer class="py-10">
<nav class="pb-4 text-base font-medium text-neutral-500 dark:text-neutral-400">
<ul class="flex flex-col list-none sm:flex-row">
<li class="mb-1 sm:mb-0 sm:mr-7 sm:last:mr-0">
<a class="decoration-primary-500 hover:underline hover:decoration-2 hover:underline-offset-2" href="/analytics" title="">Analytics</a
            >
</li>
<li class="mb-1 sm:mb-0 sm:mr-7 sm:last:mr-0">
<a class="decoration-primary-500 hover:underline hover:decoration-2 hover:underline-offset-2" href="/index.xml" title="">RSS</a
            >
</li>
<li class="mb-1 sm:mb-0 sm:mr-7 sm:last:mr-0">
<a class="decoration-primary-500 hover:underline hover:decoration-2 hover:underline-offset-2" href="https://forms.gle/orfXK5jh1LRaiFzo8" title="">Suggest a topic</a
            >
</li>
</ul>
</nav>
<div class="flex justify-between">
<div>
<p class="text-sm text-neutral-500 dark:text-neutral-400">