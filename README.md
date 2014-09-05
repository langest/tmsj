tmsj - Teach me some Japanese
=============================

A simple go program that parses a json for words or phrases you want to learn and outputs them into stdout.
I will primarily use this to learn Japanese by running it periodically somehow. It will therefore output a random phrase/word/kanji or something that I should learn.

It will probably work for most languages since you can just change the content of the json files.

The kanji meanings I use here is not always both on-yomi and kun-yomi. I mearly use the ones that I should learn in my Japanese course.

Install instructions
====================
- Compile with go install.
- Add a file .tmsj in your home dir that looks the following:
```
kanji=/path/to/kanji.json
glossary=/path/to/glossary.json

```
- Run with `tmsj`

TODO
====
Add flags to run different functionalities of the program such as:
- Write conky integration
- Flags that tell how many facts you want and which type
- Write script and settings so that conky can display output from the program
- Add declension to verbs and adjectives in glossary
- Add support for reading glossary and kanji from multiple files
