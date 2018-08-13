package words

import (
	"container/list"
	"fmt"
)

// Word contains the text string of the word and its sibling words that it can murph into.
type Word struct {
	Text     string
	Siblings []*Word
	Visited  bool
}

// NewWord creates a new word
func NewWord(text string) *Word {
	return &Word{
		Text:     text,
		Siblings: []*Word{},
		Visited:  false,
	}
}

// AllWords contains a map of all words
type AllWords map[string]*Word

// BuildWordSiblingFinder creates a WordSiblingFinder that can be used to connect each sibling for each word.
func (aw *AllWords) BuildWordSiblingFinder() *WordSiblingFinder {
	wsf := WordSiblingFinder{}

	fmt.Println("Building WordMap...")
	fmt.Printf("There are %d words in this dictionary\n", len(*aw)) // debug

	// iterate thru each word
	for wtext, w := range *aw {

		wlen := len(wtext) // length of the current word

		// iterate thru each letter in that word.
		for idx := 0; idx < wlen; idx++ {
			substr := wtext[:idx] + wtext[idx+1:]
			if wsf[substr] == nil {
				wsf[substr] = make([]WordSiblingGroup, wlen, wlen)
			}

			wsf[substr][idx] = append(wsf[substr][idx], w)
		}
	}
	return &wsf
}

// ChainWords find the shortest path from one word to another
func (aw *AllWords) ChainWords(wstr1 string, wstr2 string) {
	w1 := (*aw)[wstr1]

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
		for _, w := range currentTodo.this.Siblings {
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

// WordSiblingGroup is part of WordSiblingFinder. By looking up a substring from WordSiblingFinder,
// an array of WordSiblingGroup is found. The position of the array is the position where
// the letter being substituted for that particular WordSiblingGroup
type WordSiblingGroup []*Word

// WordSiblingFinder maps substrings to an array of WordSiblingGroups. The index of WordSibling array represents
// the position of the letter being substituted for that WordSibling.
type WordSiblingFinder map[string][]WordSiblingGroup

// ConnectSiblings connects all word siblings
func (wsf *WordSiblingFinder) ConnectSiblings() {

	for _, wsgs := range *wsf { // wsgs: WordSiblingGroups
		for _, wsg := range wsgs {
			for widx, w := range wsg {
				w.Siblings = append(w.Siblings, wsg[:widx]...)
				w.Siblings = append(w.Siblings, wsg[widx+1:]...)
			}
		}
	}
}
