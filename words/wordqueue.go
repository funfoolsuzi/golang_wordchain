package words

// Node is ...
type Node struct {
	word       *Word
	next       *Node
	upperchain []*Word
}

// NewNode does ...
func NewNode(w *Word) *Node {
	return &Node{
		word:       w,
		next:       nil,
		upperchain: []*Word{},
	}
}

// NewNodeWithUpper is ...
func NewNodeWithUpper(w *Word, upper *Node) *Node {
	n := NewNode(w)
	n.upperchain = append(upper.upperchain, upper.word)
	return n
}

// GetChain does ...
func (n *Node) GetChain() []string {
	strs := []string{}
	if n == nil {
		return strs
	}

	for _, w := range n.upperchain {
		strs = append(strs, w.Text)
	}
	strs = append(strs, n.word.Text)
	return strs
}

// WordQueue is ...
type WordQueue struct {
	head  *Node
	tail  *Node
	count int
}

// NewWordQueue does ...
func NewWordQueue(w *Word) *WordQueue {
	n := NewNode(w)

	wq := &WordQueue{
		head:  n,
		tail:  n,
		count: 1,
	}

	return wq
}

// AddWord does ...
func (wq *WordQueue) AddWord(w *Word, upper *Node) {
	n := NewNodeWithUpper(w, upper)
	if wq.count == 0 {
		wq.head = n
	}
	if wq.tail != nil {
		wq.tail.next = n
	}
	wq.tail = n
	wq.count++
}

// CheckNAddSiblings does ...
func (wq *WordQueue) CheckNAddSiblings(n *Node, dest *Word) *Node {
	for _, w := range n.word.Siblings {
		wq.AddWord(w, n)
		if w == dest {
			return wq.GetTail()
		}
	}
	return nil
}

// PopHead does ...
func (wq *WordQueue) PopHead() *Node {
	n := wq.head
	wq.head = wq.head.next
	wq.count--
	return n
}

// Count does ...
func (wq *WordQueue) Count() int {
	return wq.count
}

// GetTail returns the tail node
func (wq *WordQueue) GetTail() *Node {
	return wq.tail
}
