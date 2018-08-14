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

	allwords := words.CreateAllWords(dictMap)

	fmt.Println()
	fmt.Println("Example: \"cat\" to \"dog\"")

	allwords.FindChain("cat", "dog")

	allwords.RunUI()
}
