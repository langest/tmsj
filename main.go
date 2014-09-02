package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello")
	parseJson()
}

type kanjiMap map[string]map[string][]string
type glossMap map[string]glossary

type glossary struct {
	Kanji       string `json:kanji`
	Translation string `json:translation`
}

func parseJson() {
	kanji, err := os.Open("kanji.json")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := kanji.Close(); err != nil {
			panic(err)
		}
	}()

	gloss, err := os.Open("glossary.json")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := gloss.Close(); err != nil {
			panic(err)
		}
	}()

	kDec := json.NewDecoder(kanji)
	gDec := json.NewDecoder(gloss)

	var kMap kanjiMap
	var gMap glossMap

	for {
		if err := kDec.Decode(&kMap); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	for {
		if err := gDec.Decode(&gMap); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("kMap:", kMap)
	log.Println("gMap:", gMap)
}
