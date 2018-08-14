package helper

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/funfoolsuzi/golang_wordchain/words"
)

// DownloadDictionary downloads the dictionary from github.
func DownloadDictionary() (io.ReadCloser, error) {
	// out, err := os.Create("dict.json")
	// if err != nil {
	// 	return err
	// }
	// defer out.Close()

	resp, err := http.Get("https://raw.githubusercontent.com/tmobil/CodingExercise/master/english.json")
	if err != nil {
		return nil, fmt.Errorf("DownloadDictionary() error in calling http.Get(): %v", err)
	}
	return resp.Body, nil
}

// GetDictionaryBytes will return the dicitonary as an io.ReadCloser
func GetDictionaryBytes() ([]byte, error) {

	if _, err := os.Stat("dict.json"); os.IsNotExist(err) {
		dictReader, err := DownloadDictionary()
		defer dictReader.Close()
		if err != nil {
			return nil, fmt.Errorf("GetDictionaryBytes() error in calling DownloadDictionary(): %v", err)
		}

		out, err := os.Create("dict.json")
		defer out.Close()
		if err != nil {
			return nil, fmt.Errorf("GetDictionaryBytes() error in calling os.Create(): %v", err)
		}

		_, err = io.Copy(out, dictReader)
		if err != nil {
			return nil, fmt.Errorf("GetDictionaryBytes() error in calling io.Copy(): %v", err)
		}
	}

	b, err := ioutil.ReadFile("dict.json")
	if err != nil {
		return nil, fmt.Errorf("GetDictionaryBytes() error in calling ioutil.ReadFile(): %v", err)
	}

	return b, nil
}

// RunUI runs the command line user interface
func RunUI(aw *words.AllWords) {

	aw.ResetVisitStatus()

	var w1, w2, res string

	// first word
	for w1 == "" {
		fmt.Println("Plz enter the first word")
		fmt.Scanln(&w1)
		// TODO: check if in dicitonary
	}

	// second word
	for w2 == "" {
		fmt.Println("Plz enter the second word")
		fmt.Scanln(&w2)
		// TODO: check if in dictionary
	}

	if len(w1) == len(w2) {
		aw.FindChain(w1, w2)
	} else {
		fmt.Println("words with variant length are not supported right now")
	}

	// try again?
	fmt.Println("Try again? Press Enter or type \"q\" to quit")
	fmt.Scanln(&res)
	if res != "q" {
		RunUI(aw)
	}
}
