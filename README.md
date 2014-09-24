tmsj - Teach me some Japanese
=============================

A simple go program that parses a json for words or phrases you want to learn and outputs them into stdout.
I will primarily use this to learn Japanese by running it periodically somehow. It will therefore output a random phrase/word/kanji or something that I should learn.

It might work for other languages since you can just change the content of the json files. Such as changing the contents glossary.json into the language you want to learn.

The kanji meanings I have filled in in this repository is not always both on-yomi and kun-yomi. I only add the ones that I should learn in my Japanese course.

I have added a flag `-c` for printing to a file that conky can parse. The conky script for this will come up soon.

Install instructions
====================
- Compile with go install.
- Add a file .tmsj in your home dir that looks the following:
```
kanji=/path/to/kanji.json
glossary=/path/to/glossary.json

```
- Run with `tmsj` and use `-h` for help.

TODO
====
Add flags to run different functionalities of the program such as:
- Flags that tell how many facts you want and which type (maybe)
- Add declension to verbs and adjectives in glossary (for advanced learning)
- Add support for reading glossary and kanji from multiple files
