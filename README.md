## WORD CHAIN PROBLEM

The majority of the app is done before the Aug 14th COB deadline. I only changed documentation and added time elapsed in the result after the deadline. The commit history reflects it. I just found this problem really fun to solve and want to keep on improving this mini app.

### Setup Guide:

1. Make sure your [GOPATH](https://github.com/golang/go/wiki/GOPATH) is set up. If not follow [this link](https://github.com/golang/go/wiki/GOPATH).
2. Install this package by
```
    go get github.com/funfoolsuzi/golang_wordchain
```
3. go to package directory
```
    cd $GOPATH/src/github.com/golang/go/wiki/GOPATH
```
4. run the app
```
    go run main.go
```
The program will download the dictionary JSON file from the originial git repo on Github.


### Strategy

1. Get the dictionary. Map it. List all the words in a data model to keep track.

2. A _Siblingship_ is the ability of a word to change one of its letters to murph into another word.

3. Try to construct an [undirected graph](https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)#Undirected_graph) with __Word__ as _vertex_ and __Siblingship__ as _edge_. (_actually this could be multiple graphs_)

4. Use [BFS](https://en.wikipedia.org/wiki/Breadth-first_search), Breadth-First Search to find the shortest path between two words.


### Implementation

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

3. Once the _WordSiblingFinder_ has been filled out completed based on the dictionary. We can connect each _Word_ within each _WordSiblingGroup_. We loop thru each substring key in _WordSiblingFinder_. Under each substring key, we then loop thru each position where letter has been taken out. Under each position, We have a WordSiblingGroup. We make sure all the Words in that WordSiblingGroup connect with their Siblings. After this is done. The "graph" is complete.

4. Run BFS on this graph. (TODO: add details)

### Optimization Todos:
1. ~~Replace golang build-in list with a custom Queue.~~ _DONE_
2. Words can murph to different length.
3. ~~Add command-line user interface.~~ _DONE~
4. Category this "graph" into multiple graphs. Because this "graph" could actually be multiple graphs. And this will reduce the complexity on finding if two words are actually even reachable to each other.
5. Unit testing.
6. ~~Record search time elapsed.~~ _DONE_
7. Finish BFS implementation doc.