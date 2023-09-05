# Node Depths

<div class="html">
<p>
  The distance between a node in a Binary Tree and the tree's root is called the
  node's depth.
</p>
<p>
  Write a function that takes in a Binary Tree and returns the sum of its nodes'
  depths.
</p>
<p>
  Each <span>BinaryTree</span> node has an integer <span>value</span>, a
  <span>left</span> child node, and a <span>right</span> child node. Children
  nodes can either be <span>BinaryTree</span> nodes themselves or
  <span>None</span> / <span>null</span>.
</p>
<h3>Sample Input</h3>
<pre><span class="CodeEditor-promptParameter">tree</span> =    1
       /     \
      2       3
    /   \   /   \
   4     5 6     7
 /   \
8     9
</pre>
<h3>Sample Output</h3>
<pre>16
<span class="CodeEditor-promptComment">// The depth of the node with value 2 is 1.</span>
<span class="CodeEditor-promptComment">// The depth of the node with value 3 is 1.</span>
<span class="CodeEditor-promptComment">// The depth of the node with value 4 is 2.</span>
<span class="CodeEditor-promptComment">// The depth of the node with value 5 is 2.</span>
<span class="CodeEditor-promptComment">// Etc..</span>
<span class="CodeEditor-promptComment">// Summing all of these depths yields 16.</span>
</pre>
</div>
