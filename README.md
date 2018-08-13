## WORD CHAIN PROBLEM

#### Strategy

1. Get the dictionary. Map it.

2. Define connection between a pair of words as the ability to change one letter of a certain word to match the other word.

3. Try to construct an [undirected graph](https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)#Undirected_graph) with Word as vertex and connection as edge.

4. Use [BFS](https://en.wikipedia.org/wiki/Breadth-first_search), Breadth-First Search to find the shortest path between two words.


#### Implementation

1. Build a map. Transform all dictionary words as keys. Then make objects(instance of struct Word) that each contains the dictionary word and references to other same-kind objects. Use this object as the value of the map.

2. For each word in the dictionary, go through each letter. When examing each letter, take out that letter and reconcatenate the rest of the string. We put each of this reconcatenated substring as a key to a new map we call WordLinkMap. The counterpart to any of these keys is an array of WordLink(s). The array index matches the position where the letter has been taken out. As of WordLink, behind this type, it is just an array of references to different Words.
Illustration:
```
WordLinkMap {
    "hat": {
        /* '*hat', word at position 0 has been taken out */
        WordLink { &Word{...} /* chat */, &Word{...} /* that */, ...},
        /* 'h*at', word at position 1 has been taken out */
        WordLink { &Word{...} /* heat */, ...},
        /* 'ha*t', word at position 2 has been taken out */
        WordLink { },
        /* 'hat*', word at position 3 has been taken out */
        WordLink { &Word{...} /* hate */, &Word{...} /* hath */, ...},
    },
    ...
    "ue": {
        WordLink { &Word{...} /* hue */, &Word{...} /* vue */, ...},
        ...
    }
}
```

3. Once the WordLinkMap has been filled out completed based on the dictionary. We can connect each word based on each WordLink. We loop thru each substring key in the map. Under each substring key, we then loop thru each position where letter has been taken out. Under each position, We have a WordLink. We make sure all the Words in that WordLink connect with each other. After this is done. The graph is complete.

4. Run BFS on this graph.