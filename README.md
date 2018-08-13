## WORD CHAIN PROBLEM

#### Strategy

1. Get the dictionary. Map it. List all the words in a data model to keep track.

2. A _Siblingship_ is the ability to change one letter of a word to murph into another word.

3. Try to construct an [undirected graph](https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)#Undirected_graph) with __Word__ as _vertex_ and __Siblingship__ as _edge_.

4. Use [BFS](https://en.wikipedia.org/wiki/Breadth-first_search), Breadth-First Search to find the shortest path between two words.


#### Implementation

1. Push all dictionary key(word as a textstring) to an array/list that has a _Word_ struct wrapping around each item. This _Word_ struct contains meta info and a list of Siblings. The list of Sibings are initiated as empty array/list. This list of _Word_(s) is then wrapped by a data struct called _AllWords_.

2. For each word in _AllWords_, go through each letter. When examing each letter, take out that letter and reconcatenate the rest of the string. We put each of this reconcatenated substring as a key to a new map we call _WordSiblingFinder_. The counterpart to any of these keys is an array of _WordSiblingGroup_(s). The array index matches the position where the letter has been taken out. As of _WordSiblingGroup_, behind this type, it is just an array of references to different _Word_(s).
Illustration:
```
WordSiblingFinder {
    "hat": {
        /* '*hat', word at position 0 has been taken out */
        WordSiblingGroup { &Word{...} /* chat */, &Word{...} /* that */, ...},
        /* 'h*at', word at position 1 has been taken out */
        WordSiblingGroup { &Word{...} /* heat */, ...},
        /* 'ha*t', word at position 2 has been taken out */
        WordSiblingGroup { },
        /* 'hat*', word at position 3 has been taken out */
        WordSiblingGroup { &Word{...} /* hate */, &Word{...} /* hath */, ...},
    },
    ...
    "ue": {
        WordSiblingGroup { &Word{...} /* hue */, &Word{...} /* vue */, ...},
        ...
    }
}
```

3. Once the _WordSiblingFinder_ has been filled out completed based on the dictionary. We can connect each _Word_ within each _WordSiblingGroup_. We loop thru each substring key in _WordSiblingFinder_. Under each substring key, we then loop thru each position where letter has been taken out. Under each position, We have a WordSiblingGroup. We make sure all the Words in that WordSiblingGroup connect with their Siblings. After this is done. The graph is complete.

4. Run BFS on this graph.