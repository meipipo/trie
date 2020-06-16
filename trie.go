package trie

type charNode struct {
	c           string
	accept      string
	childrenMap map[string]*charNode
}

type Trie struct {
	root *charNode
}

// makeCharNode returns charNode of given character.
func makeCharNode(s string) charNode {
	return charNode{
		c:           s,
		accept:      "",
		childrenMap: make(map[string]*charNode),
	}
}

// MakeTrie returns empty trie.
func MakeTrie() Trie {
	// Root node has to be empty.
	r := makeCharNode("")
	return Trie{
		root: &r,
	}
}

// AddWord adds new word to trie.
func (t *Trie) AddWord(word string) {
	parent := t.root
	for i := 0; i < len(word); i++ {
		w := string(word[i])
		if v, ok := parent.childrenMap[w]; ok { // if given character's node already exists
			if v.c == w {
				// If the character is the last of the word.
				if i == len(word)-1 {
					v.accept = word
				}
			}
			// Traverse next.
			parent = v
		} else { // if given character firstly appears
			// Make new node.
			node := makeCharNode(w)
			if i == len(word)-1 {
				node.accept = word
			}
			// Add new node to parent.
			parent.childrenMap[w] = &node
			parent = &node
		}
	}
	return
}

// has returns true if given word is exist in trie.
func (t *Trie) has(word string) bool {
	n := t.root
	for i := 0; i < len(word); i++ {
		w := string(word[i])
		if v, ok := n.childrenMap[w]; !ok {
			return false
		} else if i == len(word)-1 && v.accept == "" {
			return false
		} else {
			n = v
		}
	}
	return true
}

// wordsTraverse returns existing words under the given node with adding prefix.
func wordsTraverse(n *charNode) []string {
	var words []string
	for _, v := range n.childrenMap {
		if v.accept != "" {
			words = append(words, v.accept)
			if len(v.childrenMap) != 0 {
				words = append(words, wordsTraverse(v)...)
			}
		} else {
			words = append(words, wordsTraverse(v)...)
		}
	}
	return words
}

// WordsPrefix returns existing words with given prefix.
func (t *Trie) WordsPrefix(pre string) []string {
	var words []string
	if t.has(pre) {
		words = append(words, pre)
	}
	n := t.root
	for i := 0; i < len(pre); i++ {
		w := string(pre[i])
		if v, ok := n.childrenMap[w]; !ok {
			return words
		} else {
			n = v
			if i == len(pre)-1 {
				break
			}
		}
	}
	words = append(words, wordsTraverse(n)...)
	return words
}
