package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
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
	log.Println("kanjiPath:", kanjiPath)
	log.Println("glossPath:", glossPath)
}

func setConf(confFile []string) {
	for _, conf := range confFile {
		c := strings.Split(conf, "=")
		if c[0] == "kanji" {
			kanjiPath = c[1]
		} else if c[0] == "glossary" {
			glossPath = c[1]
		} else {
			log.Println("Could not recognize filetype:", c[0])
		}
	}
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
