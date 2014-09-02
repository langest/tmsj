package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

var (
	config    = os.Getenv("HOME") + "/.lmsj"
	kanjiPath string
	glossPath string
)

type kanjiMap map[string]map[string][]string
type glossMap map[string]glossary

type glossary struct {
	Kanji       string `json:kanji`
	Translation string `json:translation`
}

func main() {
	conf, err := readFile(config)
	if err != nil {
		log.Fatal("reading config:", err)
		return
	}
	setConf(conf)
}

func setConf(confFile []string) {
	log.Println(confFile)
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseJson() {
	kanji, err := os.Open("kanji.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := kanji.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	gloss, err := os.Open("glossary.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := gloss.Close(); err != nil {
			log.Fatal(err)
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
