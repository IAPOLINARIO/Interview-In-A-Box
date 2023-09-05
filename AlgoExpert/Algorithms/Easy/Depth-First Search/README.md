<div class="html">
<p>
  You're given a <span>Node</span> class that has a <span>name</span> and an
  array of optional <span>children</span> nodes. When put together, nodes form
  an acyclic tree-like structure.
</p>
<p>
  Implement the <span>depthFirstSearch</span> method on the
  <span>Node</span> class, which takes in an empty array, traverses the tree
  using the Depth-first Search approach (specifically navigating the tree from
  left to right), stores all of the nodes' names in the input array, and returns
  it.
</p>
<p>
  If you're unfamiliar with Depth-first Search, we recommend watching the
  Conceptual Overview section of this question's video explanation before
  starting to code.
</p>
<h3>Sample Input</h3>
<pre><span class="CodeEditor-promptParameter">graph</span> = A
     /  |  \
    B   C   D
   / \     / \
  E   F   G   H
     / \   \
    I   J   K
</pre>
<h3>Sample Output</h3>
<pre>["A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"]
</pre>
</div>
