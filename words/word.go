package words

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
