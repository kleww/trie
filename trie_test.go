package trie

import "testing"

func TestCreateTrie(t *testing.T) {
	tNew := New()
	if (tNew.root.content != 0) ||
		(tNew.root.children != nil) ||
		(tNew.root.wordMarker != false) ||
		(tNew.nbNode != 0) {
		t.Error("Expected", Trie{}, "got", *tNew)
	}
}

func TestAddEmptyWord(t *testing.T) {
	var trie Trie

	var expectedResult = true

	trie.AddWord("")
	result := trie.SearchWord("")
	if result != expectedResult {
		t.Error("Expected", expectedResult, "got", result)
	}
}

func TestAddWord(t *testing.T) {
	var trie Trie

	var expectedResult = true

	trie.AddWord("Hello")
	result := trie.SearchWord("Hello")
	if result != expectedResult {
		t.Error("Expected", expectedResult, "got", result)
	}
}

func TestAddSeveralWordWithCommonRoot(t *testing.T) {
	var trie Trie

	var expectedResult = true

	var words = []string{"", "Air", "Airplane", "Airport", "Hour"}
	for _, w := range words {
		trie.AddWord(w)
	}
	for _, w := range words {
		result := trie.SearchWord(w)
		if result != expectedResult {
			t.Error("Expected", expectedResult, "got", result)
		}
	}
}

func TestSearchNotAddedWord(t *testing.T) {
	var trie Trie

	var expectedResult = false

	result := trie.SearchWord("Hello")
	if result != expectedResult {
		t.Error("Expected", expectedResult, "got", result)
	}
}

func TestSearchNotAddedWordWithEmptyStringAdded(t *testing.T) {
	var trie Trie

	var expectedResult = false

	trie.AddWord("")
	result := trie.SearchWord("Hello")
	if result != expectedResult {
		t.Error("Expected", expectedResult, "got", result)
	}
}
