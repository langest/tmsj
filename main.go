package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	config = os.Getenv("HOME") + "/.tmsj"
)

type kanjiMap []kanji
type glossMap []glossary

type kanji struct {
	Kanji         string   `json:kanji`
	Pronunciation []string `json:pronunciation`
	Translation   []string `json:translation`
}

type glossary struct {
	HiraKata    string `json:hirakata`
	Kanji       string `json:kanji`
	Translation string `json:translation`
}

func main() {
	kanjiPath, glossPath, err := loadConf(config)
	if err != nil {
		log.Fatal("reading config:", err)
		return
	}
	kMap, gMap := parseJson(kanjiPath, glossPath)
	printOneRandom(kMap, gMap)
}

func loadConf(path string) (string, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", "", err
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
	kanji, gloss := getPaths(lines)
	return kanji, gloss, scanner.Err()
}

func getPaths(confFile []string) (kanji, gloss string) {
	for _, conf := range confFile {
		c := strings.Split(conf, "=")
		if c[0] == "kanji" {
			kanji = c[1]
		} else if c[0] == "glossary" {
			gloss = c[1]
		} else {
			log.Println("Could not recognize filetype:", c[0])
		}
	}
	return
}
func printOneRandom(kMap kanjiMap, gMap glossMap) {
	//TODO format the output prettier
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("Some Japanese for you to enjoy!")
	fmt.Println("-------------Kanji-------------")
	fmt.Println(kMap[rnd.Intn(len(kMap))])
	fmt.Println("-------and some glossary-------")
	fmt.Println(gMap[rnd.Intn(len(gMap))])

}

func parseJson(kanjiPath, glossPath string) (kMap kanjiMap, gMap glossMap) {
	kanji, err := os.Open(kanjiPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := kanji.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	gloss, err := os.Open(glossPath)
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
	return
}
