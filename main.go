package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var config = os.Getenv("HOME") + "/.tmsj"
var forConky = flag.Bool("c", false, "Enable this flag if you want output written to file for conky")

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

func init() {
	flag.Parse()
}

func main() {
	kanjiPath, glossPath, conkyRoot, err := loadConf(config)
	if err != nil {
		log.Fatal("reading config:", err)
		return
	}
	kMap, gMap := parseJson(kanjiPath, glossPath)
	if *forConky {
		printForConky(conkyRoot, kMap, gMap)
	} else {
		printOneRandom(kMap, gMap)
	}
}

func loadConf(path string) (string, string, string, error) {
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
	kanji, gloss, conky := getPaths(lines)
	return kanji, gloss, scanner.Err()
}

func getPaths(confFile []string) (kanji, gloss, conky string) {
	for _, conf := range confFile {
		c := strings.Split(conf, "=")
		if c[0] == "kanji" {
			kanji = c[1]
		} else if c[0] == "glossary" {
			gloss = c[1]
		} else if c[0] == "conky" {
			conky = c[1]
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
	fmt.Println("\n-------------Kanji-------------")
	kEntry := kMap[rnd.Intn(len(kMap))]
	if kEntry.Kanji != "" {
		fmt.Println(kEntry.Kanji)

		fmt.Println("Pronunciations")
		if len(kEntry.Pronunciation) > 0 {
			for _, p := range kEntry.Pronunciation {
				fmt.Println(p)
			}
		}
		fmt.Println("Translations")
		if len(kEntry.Translation) > 0 {
			for _, t := range kEntry.Translation {
				fmt.Println(t)
			}
		}
	}
	fmt.Println("\n-------and some glossary-------")
	gEntry := gMap[rnd.Intn(len(gMap))]
	fmt.Println(gEntry)
	fmt.Println()

}

func printForConky(conkyRoot string, kMap kanjiMap, gMap glossMap) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	kEntry := kMap[rnd.Intn(len(kMap))]
	if kEntry.Kanji != "" {
		fmt.Println(kEntry.Kanji)

		if len(kEntry.Pronunciation) > 0 {
			for _, p := range kEntry.Pronunciation {
				fmt.Println(p)
			}
		}
		if len(kEntry.Translation) > 0 {
			for _, t := range kEntry.Translation {
				fmt.Println(t)
			}
		}
	}
	gEntry := gMap[rnd.Intn(len(gMap))]
	fmt.Println(gEntry)
	fmt.Println()

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
