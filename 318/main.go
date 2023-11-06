type Trie struct {
    children [26]*Trie
    isEnd bool
}

func New() *Trie {
    return &Trie{}
}

func (trie *Trie) Insert(word string) {
    for _, c := range word {
        if trie.children[c-'a'] == nil {
            trie.children[c-'a'] = New()
        }
        trie = trie.children[c-'a']
    }
    trie.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil || !node.children[ch].isEnd {
			return false
		}
		node = node.children[ch]
	}
	return true
}


func longestWord(words []string) string {
    trie := New()
    trie.isEnd = true
    for _, word := range words {
        trie.Insert(word)
    }
    currentMax := ""
    for _, word := range words {
		if trie.Search(word) && (len(word) > len(currentMax) || len(word) == len(currentMax) && word < currentMax) {
			currentMax = word
		}
	}
    return currentMax
}