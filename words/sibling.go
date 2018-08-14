package words

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
