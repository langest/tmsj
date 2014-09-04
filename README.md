tmsj - Teach me some Japanese
=============================

A simple go program that parses a json for words or phrases you want to learn and outputs them into stdout.
I will primarily use this to learn Japanese by running it periodically somehow. It will therefore output a random phrase/word/kanji or something that I should learn.

It will probably work for most languages since you can just change the content of the json files.

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
- Add support to printing to file at given intervals so that integration with conky scripts are possible
- Flags that tell how many facts you want and which type
- Write script and settings so that conky can display output from the program
