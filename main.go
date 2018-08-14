package main

import (
	"encoding/json"
	"fmt"

	"github.com/funfoolsuzi/golang_wordchain/helper"
	"github.com/funfoolsuzi/golang_wordchain/words"
)

func main() {

	fmt.Println("Preparing dictionary")
	dictBytes, err := helper.GetDictionaryBytes()
	if err != nil {
		fmt.Println("err")
	}

	// Parse 'dictionary json file bytes' to a map
	dictMap := &map[string][]string{}
	if err = json.Unmarshal(dictBytes, dictMap); err != nil {
		panic("Failed to unmarshal dictionary:" + err.Error())
	}

	// Transfer all the keys from the map to a WordMap.
	// Because we only need the keys(words)
	allwords := words.CreateAllWords(dictMap)

	wlm := allwords.BuildWordSiblingFinder()
	wlm.ConnectSiblings()

	fmt.Printf("%v\n", allwords.FindChain("cat", "dog"))
}
