package go_phrase_scanner

type TrieNode struct {
	char     rune
	children map[rune]*TrieNode
	needles  [][]rune
	is_match bool
	depth    uint
}

func NewTrie(phrases []string) TrieNode {
	n := TrieNode{
		children: make(map[rune]*TrieNode),
		needles:  make([][]rune, len(phrases)),
		depth:    0,
	}
	for i, phrase := range phrases {
		n.needles[i] = []rune(phrase)
	}
	return n
}

func (node *TrieNode) Build() {
	node.children = make(map[rune]*TrieNode)
	for _, n := range node.needles {
		if len(n) == 0 {
			node.is_match = true
			continue
		}
		char := n[0]
		cn := node.children[char]
		if cn == nil {
			cn = &TrieNode{char: char, depth: node.depth + 1, children: make(map[rune]*TrieNode)}
			node.children[char] = cn
		}
		cn.needles = append(node.needles, n[1:])
	}
	node.needles = nil // Free the memory
	for _, cn := range node.children {
		cn.Build()
	}
}

func (node *TrieNode) ScanString(s string, ch chan string) {
	if node.depth != 0 {
		panic("this method may only be called on the root node")
	}
	runes := []rune(s)
	num_runes := len(runes)
	for i := 0; i < num_runes; i++ {
		node.lookup(runes, i, i, ch)
	}
	close(ch)
}

func (node *TrieNode) lookup(runes []rune, i_start int, i_end int, ch chan string) {
	if node.is_match {
		ch <- string(runes[i_start:i_end])
	}
	if i_end == len(runes) {
		return
	}
	char := runes[i_end]
	child := node.children[char]
	if child == nil {
		return
	}
	child.lookup(runes, i_start, i_end+1, ch)
}
