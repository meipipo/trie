package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	tr := MakeTrie()
	tr.AddWord("abc")
	// Trie: root (EOF) -> "a" -> "b" -> "c" (exist)
	assert.Equal(t, tr.has("abc"), true)
	assert.Equal(t, tr.has("ab"), false)
	// Information of root.
	assert.Equal(t, tr.root.c, "")
	assert.Equal(t, tr.root.accept, "")
	assert.Equal(t, len(tr.root.childrenMap), 1)
	// Information of "a".
	firstNode := tr.root.childrenMap["a"]
	assert.Equal(t, firstNode.c, "a")
	assert.Equal(t, firstNode.accept, "")
	assert.Equal(t, len(firstNode.childrenMap), 1)
	// Information of "b".
	secondNode := firstNode.childrenMap["b"]
	assert.Equal(t, secondNode.c, "b")
	assert.Equal(t, secondNode.accept, "")
	assert.Equal(t, len(secondNode.childrenMap), 1)
	// Information of "c".
	thirdNode := secondNode.childrenMap["c"]
	assert.Equal(t, thirdNode.c, "c")
	assert.Equal(t, thirdNode.accept, "abc")
	assert.Equal(t, len(thirdNode.childrenMap), 0)

	tr.AddWord("aaaa")
	// Trie: root (EOF) -> "a" -> "b" -> "c" (exist)
	//                         -> "a" -> "a" -> "a"
	assert.Equal(t, tr.has("abc"), true)
	assert.Equal(t, tr.has("aaaa"), true)
	assert.Equal(t, tr.has("ab"), false)
	// Information of root.
	assert.Equal(t, tr.root.c, "")
	assert.Equal(t, tr.root.accept, "")
	assert.Equal(t, len(tr.root.childrenMap), 1)
	// Information of "a".
	firstNode = tr.root.childrenMap["a"]
	assert.Equal(t, firstNode.c, "a")
	assert.Equal(t, firstNode.accept, "")
	assert.Equal(t, len(firstNode.childrenMap), 2)
	// Information of "b".
	secondNode = firstNode.childrenMap["b"]
	assert.Equal(t, secondNode.c, "b")
	assert.Equal(t, secondNode.accept, "")
	assert.Equal(t, len(secondNode.childrenMap), 1)
	// Information of second "a".
	secondNode2 := firstNode.childrenMap["a"]
	assert.Equal(t, secondNode2.c, "a")
	assert.Equal(t, secondNode2.accept, "")
	assert.Equal(t, len(secondNode2.childrenMap), 1)
	// Information of "c".
	thirdNode = secondNode.childrenMap["c"]
	assert.Equal(t, thirdNode.c, "c")
	assert.Equal(t, thirdNode.accept, "abc")
	assert.Equal(t, len(thirdNode.childrenMap), 0)
	// Information of third "a".
	thirdNode2 := secondNode2.childrenMap["a"]
	assert.Equal(t, thirdNode2.c, "a")
	assert.Equal(t, thirdNode2.accept, "")
	assert.Equal(t, len(thirdNode2.childrenMap), 1)
	// Information of forth "a".
	forthNode := thirdNode2.childrenMap["a"]
	assert.Equal(t, forthNode.c, "a")
	assert.Equal(t, forthNode.accept, "aaaa")
	assert.Equal(t, len(forthNode.childrenMap), 0)

	tr.AddWord("a")
	// Trie: root (EOF) -> "a"(exist) -> "b" -> "c" (exist)
	//                                -> "a" -> "a" -> "a" (exist)
	assert.Equal(t, tr.has("abc"), true)
	assert.Equal(t, tr.has("a"), true)
	assert.Equal(t, tr.has("aa"), false)
	// Information of root.
	assert.Equal(t, tr.root.c, "")
	assert.Equal(t, tr.root.accept, "")
	assert.Equal(t, len(tr.root.childrenMap), 1)
	// Information of "a".
	firstNode = tr.root.childrenMap["a"]
	assert.Equal(t, firstNode.c, "a")
	assert.Equal(t, firstNode.accept, "a")
	assert.Equal(t, len(firstNode.childrenMap), 2)

	tr2 := MakeTrie()
	tr2.AddWord("x")
	// Trie: root (EOF) -> "x" (exist)
	assert.Equal(t, tr2.has("x"), true)
	assert.Equal(t, tr2.has("a"), false)
	// Information of root.
	assert.Equal(t, tr2.root.c, "")
	assert.Equal(t, tr2.root.accept, "")
	assert.Equal(t, len(tr2.root.childrenMap), 1)
	// Information of "x".
	firstNode2 := tr2.root.childrenMap["x"]
	assert.Equal(t, firstNode2.c, "x")
	assert.Equal(t, firstNode2.accept, "x")
	assert.Equal(t, len(firstNode2.childrenMap), 0)
}

func TestPrefix(t *testing.T) {
	tr := MakeTrie()
	tr.AddWord("abc")
	tr.AddWord("aba")
	tr.AddWord("d")

	assert.Contains(t, tr.WordsPrefix(""), "d")
	assert.Contains(t, tr.WordsPrefix(""), "abc")
	assert.Contains(t, tr.WordsPrefix(""), "abc")
	assert.Contains(t, tr.WordsPrefix("a"), "abc")
	assert.Contains(t, tr.WordsPrefix("a"), "aba")
	assert.Contains(t, tr.WordsPrefix("ab"), "abc")
	assert.Contains(t, tr.WordsPrefix("ab"), "aba")
	assert.Contains(t, tr.WordsPrefix("abc"), "abc")
	assert.Nil(t, tr.WordsPrefix("c"))
}
