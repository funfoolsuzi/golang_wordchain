package words

import (
	"fmt"
	"strings"
)

// CreateAllWords creates an Allwords instance and return its pointer
func CreateAllWords(dictMap *map[string][]string) *AllWords {
	words := []*Word{}
	for k := range *dictMap {
		lower := strings.ToLower(k)
		words = append(words, NewWord(lower))
		delete(*dictMap, k)
	}
	return &AllWords{
		words:            words,
		siblingConnected: false,
		reset:            false,
	}
}

// AllWords contains a map of all words
type AllWords struct {
	words            []*Word
	siblingConnected bool
	reset            bool
}

// Find does
func (aw *AllWords) Find(wtext string) *Word {
	for _, w := range aw.words {
		if w.Text == wtext {
			return w
		}
	}
	return nil
}

// Count returns the size of AllWords
func (aw *AllWords) Count() int {
	return len(aw.words)
}

// BuildWordSiblingFinder creates a WordSiblingFinder that can be used to connect each sibling for each word.
func (aw *AllWords) BuildWordSiblingFinder() *WordSiblingFinder {
	wsf := WordSiblingFinder{}

	fmt.Println("Building WordMap...")
	fmt.Printf("There are %d words in this dictionary\n", aw.Count()) // debug

	// iterate thru each word
	for _, w := range aw.words {

		wlen := len(w.Text) // length of the current word

		// iterate thru each letter in that word.
		for idx := 0; idx < wlen; idx++ {
			substr := w.Text[:idx] + w.Text[idx+1:]
			if wsf[substr] == nil {
				wsf[substr] = make([]WordSiblingGroup, wlen, wlen)
			}

			wsf[substr][idx] = append(wsf[substr][idx], w)
		}
	}
	return &wsf
}

// FindChain find the shortest path from one word to another
func (aw *AllWords) FindChain(wstr1 string, wstr2 string) []string {
	var endNode *Node
	result := []string{}

	w1 := aw.Find(wstr1)
	w2 := aw.Find(wstr2)
	if w1 == nil || w2 == nil {
		return result
	}

	q := NewWordQueue(w1)

	for q.Count() != 0 {
		head := q.PopHead()
		if head.word.Visited {
			continue
		}

		head.word.Visited = true

		endNode = q.CheckNAddSiblings(head, w2)
		if endNode != nil {
			break
		}
	}

	return endNode.GetChain()
}

// ResetVisitStatus resets 'Visited' in each word to false
func (aw *AllWords) ResetVisitStatus() {
	for _, w := range aw.words {
		w.Visited = false
	}
}
