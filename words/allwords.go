package words

import (
	"fmt"
	"strings"
	"time"
)

// CreateAllWords creates an Allwords instance and return its pointer
func CreateAllWords(dictMap *map[string][]string) *AllWords {
	words := []*Word{}
	for k := range *dictMap {
		lower := strings.ToLower(k)
		words = append(words, NewWord(lower))
		delete(*dictMap, k)
	}
	aw := &AllWords{
		words:            words,
		siblingConnected: false,
	}

	wsf := aw.BuildWordSiblingFinder()
	wsf.ConnectSiblings()

	return aw
}

// AllWords contains a map of all words
type AllWords struct {
	words            []*Word
	siblingConnected bool
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
	t := time.Now()
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
	fmt.Printf("Took %v to build Sibling Finder.", time.Since(t))
	return &wsf
}

// FindChain find the shortest path from one word to another
func (aw *AllWords) FindChain(wstr1 string, wstr2 string) []string {
	t := time.Now()
	fmt.Println()
	var endNode *Node
	result := []string{}

	w1 := aw.Find(wstr1)
	if w1 == nil {
		fmt.Println(wstr1, " is not found.")
		return result
	}
	w2 := aw.Find(wstr2)
	if w2 == nil {
		fmt.Println(wstr2, " is not found.")
		return result
	}

	q := NewWordQueue(w1)

	fmt.Printf("Start to search word chain between %s and %s\n", wstr1, wstr2)
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

	result = endNode.GetChain()
	fmt.Println(result, " time elapsed: ", time.Since(t))
	fmt.Println()

	return result
}

// ResetVisitStatus resets 'Visited' in each word to false
func (aw *AllWords) ResetVisitStatus() {
	for _, w := range aw.words {
		w.Visited = false
	}
}

// RunUI runs the command line user interface
func (aw *AllWords) RunUI() {

	aw.ResetVisitStatus()

	var w1, w2, res string

	// first word
	for w1 == "" {
		fmt.Print("Plz enter the first word\n>>>")
		fmt.Scanln(&w1)
		// TODO: check if in dicitonary
	}

	// second word
	for w2 == "" {
		fmt.Print("Plz enter the second word\n>>>")
		fmt.Scanln(&w2)
		// TODO: check if in dictionary
	}

	if len(w1) == len(w2) {
		aw.FindChain(w1, w2)
	} else {
		fmt.Println("words with variant length are not supported right now")
	}

	// try again?
	fmt.Print("Try again? Press \"Enter\" to continue, or type \"q\" to quit\n>>>")
	fmt.Scanln(&res)
	if res != "q" {
		aw.RunUI()
	}
}
