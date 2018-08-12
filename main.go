package main

import (
	"encoding/json"
	"fmt"

	"github.com/funfoolsuzi/golang_wordchain/helper"
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

	// Transfer all the keys from the map to a string slice.
	// Because we only need the keys(words)
	words := []string{}
	for k := range *dictMap {
		words = append(words, k)
		delete(*dictMap, k)
	}

	fmt.Println(words[0])
}
