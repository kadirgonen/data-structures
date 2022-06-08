<p>A <strong>graph</strong> is a representation of a network structure. There are tons of graph real world examples, the Internet and the social graph being the classic ones.</p>

<p>It&rsquo;s basically a set of <strong>nodes</strong> connected by <strong>edges</strong>.</p>

<p>I&rsquo;ll skip the mathematical concepts since you can find them everywhere and jump directly to a Go implementation of a graph.</p>

<h2 id="implementation">Implementation</h2>

<p>A graph data structure will expose those methods:</p>

<ul>
<li><code>AddNode()</code> inserts a node</li>
<li><code>AddEdge()</code> adds an edge between two nodes</li>
</ul>

<p>and <code>String()</code> for inspection purposes.</p>

<p>A Graph is defined as</p>
<div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-go" data-lang="go"><span style="color:#66d9ef">type</span> <span style="color:#a6e22e">ItemGraph</span> <span style="color:#66d9ef">struct</span> {
    <span style="color:#a6e22e">nodes</span> []<span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>
    <span style="color:#a6e22e">edges</span> <span style="color:#66d9ef">map</span>[<span style="color:#a6e22e">Node</span>][]<span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>
    <span style="color:#a6e22e">lock</span>  <span style="color:#a6e22e">sync</span>.<span style="color:#a6e22e">RWMutex</span>
}</code></pre></div>
<p>with Node being</p>
<div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-go" data-lang="go"><span style="color:#66d9ef">type</span> <span style="color:#a6e22e">Node</span> <span style="color:#66d9ef">struct</span> {
    <span style="color:#a6e22e">value</span> <span style="color:#a6e22e">Item</span>
}</code></pre></div>
<p>I&rsquo;ll implement an undirected graph, which means that adding an edge from A to B also adds an edge from B to A.</p>

<h2 id="implementation-1">Implementation</h2>
<div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-go" data-lang="go"><span style="color:#75715e">// Package graph creates a ItemGraph data structure for the Item type
</span><span style="color:#75715e"></span><span style="color:#f92672">package</span> <span style="color:#a6e22e">graph</span>

<span style="color:#f92672">import</span> (
    <span style="color:#e6db74">&#34;fmt&#34;</span>
    <span style="color:#e6db74">&#34;sync&#34;</span>
)

<span style="color:#75715e">// Item the type of the binary search tree
</span><span style="color:#75715e"></span><span style="color:#66d9ef">type</span> <span style="color:#a6e22e">Item</span> <span style="color:#a6e22e">generic</span>.<span style="color:#a6e22e">Type</span>

<span style="color:#75715e">// Node a single node that composes the tree
</span><span style="color:#75715e"></span><span style="color:#66d9ef">type</span> <span style="color:#a6e22e">Node</span> <span style="color:#66d9ef">struct</span> {
    <span style="color:#a6e22e">value</span> <span style="color:#a6e22e">Item</span>
}

<span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">n</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>) <span style="color:#a6e22e">String</span>() <span style="color:#66d9ef">string</span> {
    <span style="color:#66d9ef">return</span> <span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Sprintf</span>(<span style="color:#e6db74">&#34;%v&#34;</span>, <span style="color:#a6e22e">n</span>.<span style="color:#a6e22e">value</span>)
}

<span style="color:#75715e">// ItemGraph the Items graph
</span><span style="color:#75715e"></span><span style="color:#66d9ef">type</span> <span style="color:#a6e22e">ItemGraph</span> <span style="color:#66d9ef">struct</span> {
    <span style="color:#a6e22e">nodes</span> []<span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>
    <span style="color:#a6e22e">edges</span> <span style="color:#66d9ef">map</span>[<span style="color:#a6e22e">Node</span>][]<span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>
    <span style="color:#a6e22e">lock</span>  <span style="color:#a6e22e">sync</span>.<span style="color:#a6e22e">RWMutex</span>
}

<span style="color:#75715e">// AddNode adds a node to the graph
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">g</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">ItemGraph</span>) <span style="color:#a6e22e">AddNode</span>(<span style="color:#a6e22e">n</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>) {
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Lock</span>()
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">nodes</span> = append(<span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">nodes</span>, <span style="color:#a6e22e">n</span>)
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Unlock</span>()
}

<span style="color:#75715e">// AddEdge adds an edge to the graph
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">g</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">ItemGraph</span>) <span style="color:#a6e22e">AddEdge</span>(<span style="color:#a6e22e">n1</span>, <span style="color:#a6e22e">n2</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>) {
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Lock</span>()
    <span style="color:#66d9ef">if</span> <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">edges</span> <span style="color:#f92672">==</span> <span style="color:#66d9ef">nil</span> {
        <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">edges</span> = make(<span style="color:#66d9ef">map</span>[<span style="color:#a6e22e">Node</span>][]<span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>)
    }
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">edges</span>[<span style="color:#f92672">*</span><span style="color:#a6e22e">n1</span>] = append(<span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">edges</span>[<span style="color:#f92672">*</span><span style="color:#a6e22e">n1</span>], <span style="color:#a6e22e">n2</span>)
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">edges</span>[<span style="color:#f92672">*</span><span style="color:#a6e22e">n2</span>] = append(<span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">edges</span>[<span style="color:#f92672">*</span><span style="color:#a6e22e">n2</span>], <span style="color:#a6e22e">n1</span>)
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Unlock</span>()
}

<span style="color:#75715e">// AddEdge adds an edge to the graph
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">g</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">ItemGraph</span>) <span style="color:#a6e22e">String</span>() {
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RLock</span>()
    <span style="color:#a6e22e">s</span> <span style="color:#f92672">:=</span> <span style="color:#e6db74">&#34;&#34;</span>
    <span style="color:#66d9ef">for</span> <span style="color:#a6e22e">i</span> <span style="color:#f92672">:=</span> <span style="color:#ae81ff">0</span>; <span style="color:#a6e22e">i</span> &lt; len(<span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">nodes</span>); <span style="color:#a6e22e">i</span><span style="color:#f92672">++</span> {
        <span style="color:#a6e22e">s</span> <span style="color:#f92672">+=</span> <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">nodes</span>[<span style="color:#a6e22e">i</span>].<span style="color:#a6e22e">String</span>() <span style="color:#f92672">+</span> <span style="color:#e6db74">&#34; -&gt; &#34;</span>
        <span style="color:#a6e22e">near</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">edges</span>[<span style="color:#f92672">*</span><span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">nodes</span>[<span style="color:#a6e22e">i</span>]]
        <span style="color:#66d9ef">for</span> <span style="color:#a6e22e">j</span> <span style="color:#f92672">:=</span> <span style="color:#ae81ff">0</span>; <span style="color:#a6e22e">j</span> &lt; len(<span style="color:#a6e22e">near</span>); <span style="color:#a6e22e">j</span><span style="color:#f92672">++</span> {
            <span style="color:#a6e22e">s</span> <span style="color:#f92672">+=</span> <span style="color:#a6e22e">near</span>[<span style="color:#a6e22e">j</span>].<span style="color:#a6e22e">String</span>() <span style="color:#f92672">+</span> <span style="color:#e6db74">&#34; &#34;</span>
        }
        <span style="color:#a6e22e">s</span> <span style="color:#f92672">+=</span> <span style="color:#e6db74">&#34;\n&#34;</span>
    }
    <span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Println</span>(<span style="color:#a6e22e">s</span>)
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RUnlock</span>()
}</code></pre></div>
<p>Quite straightforward. Here&rsquo;s a test that when run, will populate the graph and print</p>
<div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-bash" data-lang="bash">$ go test
A -&gt; B C D
B -&gt; A E
C -&gt; A E
D -&gt; A
E -&gt; B C F
F -&gt; E

PASS</code></pre></div><div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-go" data-lang="go"><span style="color:#f92672">package</span> <span style="color:#a6e22e">graph</span>

<span style="color:#f92672">import</span> (
    <span style="color:#e6db74">&#34;fmt&#34;</span>
    <span style="color:#e6db74">&#34;testing&#34;</span>
)

<span style="color:#66d9ef">var</span> <span style="color:#a6e22e">g</span> <span style="color:#a6e22e">ItemGraph</span>

<span style="color:#66d9ef">func</span> <span style="color:#a6e22e">fillGraph</span>() {
    <span style="color:#a6e22e">nA</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">Node</span>{<span style="color:#e6db74">&#34;A&#34;</span>}
    <span style="color:#a6e22e">nB</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">Node</span>{<span style="color:#e6db74">&#34;B&#34;</span>}
    <span style="color:#a6e22e">nC</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">Node</span>{<span style="color:#e6db74">&#34;C&#34;</span>}
    <span style="color:#a6e22e">nD</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">Node</span>{<span style="color:#e6db74">&#34;D&#34;</span>}
    <span style="color:#a6e22e">nE</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">Node</span>{<span style="color:#e6db74">&#34;E&#34;</span>}
    <span style="color:#a6e22e">nF</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">Node</span>{<span style="color:#e6db74">&#34;F&#34;</span>}
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">AddNode</span>(<span style="color:#f92672">&amp;</span><span style="color:#a6e22e">nA</span>)
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">AddNode</span>(<span style="color:#f92672">&amp;</span><span style="color:#a6e22e">nB</span>)
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">AddNode</span>(<span style="color:#f92672">&amp;</span><span style="color:#a6e22e">nC</span>)
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">AddNode</span>(<span style="color:#f92672">&amp;</span><span style="color:#a6e22e">nD</span>)
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">AddNode</span>(<span style="color:#f92672">&amp;</span><span style="color:#a6e22e">nE</span>)
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">AddNode</span>(<span style="color:#f92672">&amp;</span><span style="color:#a6e22e">nF</span>)

}

<span style="color:#66d9ef">func</span> <span style="color:#a6e22e">TestAdd</span>(<span style="color:#a6e22e">t</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">testing</span>.<span style="color:#a6e22e">T</span>) {
    <span style="color:#a6e22e">fillGraph</span>()
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">String</span>()
}</code></pre></div>
<h2 id="traversing-the-graph-bfs">Traversing the graph: BFS</h2>

<p><strong>BFS</strong> (Breadth-First Search) is one of the most widely known algorithm to traverse a graph. Starting from a node, it first traverses all its directly linked nodes, then processes the nodes linked to those, and so on.</p>

<p>It&rsquo;s implemented using a queue, which is generated from my <a href="/golang-data-structure-queue/">Go implementation of a queue data structure</a> with code generation for the node type:</p>
<div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-go" data-lang="go"><span style="color:#75715e">// This file was automatically generated by genny.
</span><span style="color:#75715e">// Any changes will be lost if this file is regenerated.
</span><span style="color:#75715e">// see https://github.com/cheekybits/genny
</span><span style="color:#75715e"></span><span style="color:#f92672">package</span> <span style="color:#a6e22e">graph</span>

<span style="color:#f92672">import</span> <span style="color:#e6db74">&#34;sync&#34;</span>

<span style="color:#75715e">// NodeQueue the queue of Nodes
</span><span style="color:#75715e"></span><span style="color:#66d9ef">type</span> <span style="color:#a6e22e">NodeQueue</span> <span style="color:#66d9ef">struct</span> {
    <span style="color:#a6e22e">items</span> []<span style="color:#a6e22e">Node</span>
    <span style="color:#a6e22e">lock</span>  <span style="color:#a6e22e">sync</span>.<span style="color:#a6e22e">RWMutex</span>
}

<span style="color:#75715e">// New creates a new NodeQueue
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">s</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">NodeQueue</span>) <span style="color:#a6e22e">New</span>() <span style="color:#f92672">*</span><span style="color:#a6e22e">NodeQueue</span> {
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Lock</span>()
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span> = []<span style="color:#a6e22e">Node</span>{}
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Unlock</span>()
    <span style="color:#66d9ef">return</span> <span style="color:#a6e22e">s</span>
}

<span style="color:#75715e">// Enqueue adds an Node to the end of the queue
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">s</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">NodeQueue</span>) <span style="color:#a6e22e">Enqueue</span>(<span style="color:#a6e22e">t</span> <span style="color:#a6e22e">Node</span>) {
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Lock</span>()
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span> = append(<span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span>, <span style="color:#a6e22e">t</span>)
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Unlock</span>()
}

<span style="color:#75715e">// Dequeue removes an Node from the start of the queue
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">s</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">NodeQueue</span>) <span style="color:#a6e22e">Dequeue</span>() <span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span> {
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Lock</span>()
    <span style="color:#a6e22e">item</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span>[<span style="color:#ae81ff">0</span>]
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span> = <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span>[<span style="color:#ae81ff">1</span>:len(<span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span>)]
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">Unlock</span>()
    <span style="color:#66d9ef">return</span> <span style="color:#f92672">&amp;</span><span style="color:#a6e22e">item</span>
}

<span style="color:#75715e">// Front returns the item next in the queue, without removing it
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">s</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">NodeQueue</span>) <span style="color:#a6e22e">Front</span>() <span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span> {
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RLock</span>()
    <span style="color:#a6e22e">item</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span>[<span style="color:#ae81ff">0</span>]
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RUnlock</span>()
    <span style="color:#66d9ef">return</span> <span style="color:#f92672">&amp;</span><span style="color:#a6e22e">item</span>
}

<span style="color:#75715e">// IsEmpty returns true if the queue is empty
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">s</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">NodeQueue</span>) <span style="color:#a6e22e">IsEmpty</span>() <span style="color:#66d9ef">bool</span> {
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RLock</span>()
    <span style="color:#66d9ef">defer</span> <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RUnlock</span>()
    <span style="color:#66d9ef">return</span> len(<span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span>) <span style="color:#f92672">==</span> <span style="color:#ae81ff">0</span>
}

<span style="color:#75715e">// Size returns the number of Nodes in the queue
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">s</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">NodeQueue</span>) <span style="color:#a6e22e">Size</span>() <span style="color:#66d9ef">int</span> {
    <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RLock</span>()
    <span style="color:#66d9ef">defer</span> <span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RUnlock</span>()
    <span style="color:#66d9ef">return</span> len(<span style="color:#a6e22e">s</span>.<span style="color:#a6e22e">items</span>)
}</code></pre></div>
<h3 id="traverse-method">Traverse method</h3>

<p>Here&rsquo;s the BFS implementation:</p>
<div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-go" data-lang="go"><span style="color:#75715e">// Traverse implements the BFS traversing algorithm
</span><span style="color:#75715e"></span><span style="color:#66d9ef">func</span> (<span style="color:#a6e22e">g</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">ItemGraph</span>) <span style="color:#a6e22e">Traverse</span>(<span style="color:#a6e22e">f</span> <span style="color:#66d9ef">func</span>(<span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>)) {
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RLock</span>()
    <span style="color:#a6e22e">q</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">NodeQueue</span>{}
    <span style="color:#a6e22e">q</span>.<span style="color:#a6e22e">New</span>()
    <span style="color:#a6e22e">n</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">nodes</span>[<span style="color:#ae81ff">0</span>]
    <span style="color:#a6e22e">q</span>.<span style="color:#a6e22e">Enqueue</span>(<span style="color:#f92672">*</span><span style="color:#a6e22e">n</span>)
    <span style="color:#a6e22e">visited</span> <span style="color:#f92672">:=</span> make(<span style="color:#66d9ef">map</span>[<span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>]<span style="color:#66d9ef">bool</span>)
    <span style="color:#66d9ef">for</span> {
        <span style="color:#66d9ef">if</span> <span style="color:#a6e22e">q</span>.<span style="color:#a6e22e">IsEmpty</span>() {
            <span style="color:#66d9ef">break</span>
        }
        <span style="color:#a6e22e">node</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">q</span>.<span style="color:#a6e22e">Dequeue</span>()
        <span style="color:#a6e22e">visited</span>[<span style="color:#a6e22e">node</span>] = <span style="color:#66d9ef">true</span>
        <span style="color:#a6e22e">near</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">edges</span>[<span style="color:#f92672">*</span><span style="color:#a6e22e">node</span>]

        <span style="color:#66d9ef">for</span> <span style="color:#a6e22e">i</span> <span style="color:#f92672">:=</span> <span style="color:#ae81ff">0</span>; <span style="color:#a6e22e">i</span> &lt; len(<span style="color:#a6e22e">near</span>); <span style="color:#a6e22e">i</span><span style="color:#f92672">++</span> {
            <span style="color:#a6e22e">j</span> <span style="color:#f92672">:=</span> <span style="color:#a6e22e">near</span>[<span style="color:#a6e22e">i</span>]
            <span style="color:#66d9ef">if</span> !<span style="color:#a6e22e">visited</span>[<span style="color:#a6e22e">j</span>] {
                <span style="color:#a6e22e">q</span>.<span style="color:#a6e22e">Enqueue</span>(<span style="color:#f92672">*</span><span style="color:#a6e22e">j</span>)
                <span style="color:#a6e22e">visited</span>[<span style="color:#a6e22e">j</span>] = <span style="color:#66d9ef">true</span>
            }
        }
        <span style="color:#66d9ef">if</span> <span style="color:#a6e22e">f</span> <span style="color:#f92672">!=</span> <span style="color:#66d9ef">nil</span> {
            <span style="color:#a6e22e">f</span>(<span style="color:#a6e22e">node</span>)
        }
    }
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">lock</span>.<span style="color:#a6e22e">RUnlock</span>()
}</code></pre></div>
<p>which can be tested with</p>
<div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-go" data-lang="go"><span style="color:#66d9ef">func</span> <span style="color:#a6e22e">TestTraverse</span>(<span style="color:#a6e22e">t</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">testing</span>.<span style="color:#a6e22e">T</span>) {
    <span style="color:#a6e22e">g</span>.<span style="color:#a6e22e">Traverse</span>(<span style="color:#66d9ef">func</span>(<span style="color:#a6e22e">n</span> <span style="color:#f92672">*</span><span style="color:#a6e22e">Node</span>) {
        <span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Printf</span>(<span style="color:#e6db74">&#34;%v\n&#34;</span>, <span style="color:#a6e22e">n</span>)
    })
}</code></pre></div>
<p>which after being added to our tests, will print the road it took to traverse all our nodes:</p>
<div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-bash" data-lang="bash">A
B
C
D
A
E
F</code></pre></div>
<h2 id="creating-a-concrete-graph-data-structure">Creating a concrete graph data structure</h2>

<p>You can use this generic implemenation to generate type-specific graphs, using</p>
<div class="highlight"><pre tabindex="0" style="color:#f8f8f2;background-color:#272822;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-bash" data-lang="bash">//generate a <span style="color:#e6db74">`</span>IntGraph<span style="color:#e6db74">`</span> graph of <span style="color:#e6db74">`</span>int<span style="color:#e6db74">`</span> values
genny -in graph.go -out graph-int.go gen <span style="color:#e6db74">&#34;Item=int&#34;</span>

//generate a <span style="color:#e6db74">`</span>StringGraph<span style="color:#e6db74">`</span> graph of <span style="color:#e6db74">`</span>string<span style="color:#e6db74">`</span> values
genny -in graph.go -out graph-string.go gen <span style="color:#e6db74">&#34;Item=string&#34;</span></code></pre></div>
