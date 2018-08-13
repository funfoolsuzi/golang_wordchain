package words

import (
	"container/list"
	"fmt"
)

// Word contains the text string of the word and its related words that i can murph into.
type Word struct {
	Text        string
	LinkedWords []*Word
	Visited     bool
}

// NewWord does
func NewWord(text string) *Word {
	return &Word{
		Text:        text,
		LinkedWords: []*Word{},
		Visited:     false,
	}
}

// WordMap is
type WordMap map[string]*Word

// WordLink is part of WordLinkMap. By looking up a substring from WordLinkMap,
// an array of WordLinks is found. The position of the array is the position where
// the letter being substituted for that particular WordLink
type WordLink []*Word

// WordLinkMap maps substrings to an array of WordLinks. The index of WordLink array represents
// the position of the letter being substituted for that WordLink.
type WordLinkMap map[string][]WordLink

// BuildWordLinkMapFromWords creates a WordLinkMap out of a WordMap
func (wm *WordMap) BuildWordLinkMapFromWords() *WordLinkMap {
	wlm := WordLinkMap{}

	fmt.Println("Building WordMap...")
	fmt.Printf("There are %d words in this dictionary\n", len(*wm)) // debug

	// iterate thru each word
	for w, wObj := range *wm {

		wlen := len(w) // length of the current word

		// iterate thru each letter in that word.
		for idx := 0; idx < wlen; idx++ {
			substr := w[:idx] + w[idx+1:]
			if wlm[substr] == nil {
				wlm[substr] = make([]WordLink, wlen, wlen)
			}

			wlm[substr][idx] = append(wlm[substr][idx], wObj)
		}
	}
	return &wlm
}

// ChainWords does
func (wm *WordMap) ChainWords(wstr1 string, wstr2 string) {
	w1 := (*wm)[wstr1]

	l := list.New()
	type node struct {
		this    *Word
		parents []*Word
	}
	var res *node

	l.PushBack(node{this: w1, parents: []*Word{}})
	for l.Len() != 0 {
		currentElem := l.Front()
		currentTodo := currentElem.Value.(node)
		if currentTodo.this.Visited {
			l.Remove(currentElem)
			continue
		}
		if currentTodo.this.Text == wstr2 {
			res = &currentTodo
			break
		}
		// if len(currentTodo.parents) == 2 {
		// 	fmt.Println(currentTodo.this.Text)
		// }
		for _, w := range currentTodo.this.LinkedWords {
			l.PushBack(node{this: w, parents: append(currentTodo.parents, currentTodo.this)})
		}
		currentTodo.this.Visited = true
		l.Remove(currentElem)
	}

	if res == nil {
		fmt.Println("no chain")
		return
	}

	fmt.Print("Found chain: ")
	for _, w := range res.parents {
		fmt.Print(w.Text, " ")
	}
	fmt.Println(wstr2)
}

// ConnectWords does
func (wlm *WordLinkMap) ConnectWords() {

	for _, wls := range *wlm { // ss: SubString, wls: WordLinkS
		for _, wl := range wls {
			for widx, w := range wl {

				w.LinkedWords = append(w.LinkedWords, wl[:widx]...)
				w.LinkedWords = append(w.LinkedWords, wl[widx+1:]...)
			}
		}
	}
}
