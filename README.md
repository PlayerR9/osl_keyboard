# osl_keyboard
A small tool that allows to write text in OSL according to its writing system.


## Table of Contents

1. [Table of Contents](#table-of-contents)
2. [Overview](#overview)
3. [Installation](#installation)
   1. [Installing Go](#installing-go)
   2. [Installing the Font](#installing-the-font)
   3. [Building the Project](#building-the-project)
4. [Usage](#usage)
   1. [Executable](#executable)
5. [Configuration](#configuration)
6. [Features](#features)
7. [Testing](#testing)
8. [Contributing](#contributing)
9. [Roadmap](#roadmap)
10. [License](#license)
11. [Acknowledgments](#acknowledgments)
12. [Contact Information](#contact-information)
13. [FAQ](#faq)


## Overview

### Project Purpose

**osl_keyboard** is a small program that allows to write text in OSL according to its writing system which is displayed through the use of the IzlaRjan font.

More specifically: Given an input text file, it reads the content line by line. Through the use of a map, it "tokenize" it and, after applying the necessary transformation, it prints the result in the output file.


### Character Set

For now, the character set is "almost" entirely contained within the Private Use Area (PUA) of the font. This was done for:
- Avoiding conflicts with other fonts (or applications) that may be installed on the system.
- Ensuring the font can be extended easily in the future; such as adding Latin version of the writing system.

However, the actual range of the character set goes from U+E000 up to U+E048; minus the characters that are not in the PUA range.

More specifically:
- The following Unicode codepoints are reserved for vowels in stressed syllables or when they appear at the start of sentences:
   - U+E001, U+E004, U+E006, U+E009, U+E00C, U+E00F, U+E012, U+E015, and U+E017.
- The following Unicode codepoints are reserved for consonants:
   - U+E01A, U+E01C, U+E01D, U+E01F, U+E022, U+E015, U+E028, U+E02E, and U+E032.
- The following Unicode codepoints are reserved for vowels in unstressed syllables and when they don't appear at the start of sentences:
   - U+E01B, U+E01E, U+E020, U+E023, U+E026, U+E029, U+E02C, U+E02F, and U+E031
- The range between U+E035 to U+E048 Unicode codepoints are reserved for diacritics.


**Quick Table**

Here's a quick table of the reserved codepoints:

| **Consonants** | **Uppercase** | **Lowercase** |
| :--------: | :-----------: | :-----------: |
|  `r`       | `U+E000`      | `U+E01A`      |
|  `s`       | `U+E002`      | `U+E01C`      |
|  `v`       | `U+E003`      | `U+E01D`      |
|  `m`       | `U+E005`      | `U+E01F`      |
|  `k`       | `U+E008`      | `U+E022`      |
|  `b`       | `U+E00B`      | `U+E025`      |
|  `n`       | `U+E014`      | `U+E02E`      |
|  `t`       | `U+E016`      | `U+E030`      |
|  `d`       | `U+E018`      | `U+E032`      |
|  `u`       | `U+E001`      | `U+E01B`      |


| **Vowels** | **Uppercase** | **Lowercase** |
| :--------: | :-----------: | :-----------: |
|  `ö`       | `U+E004`      | `U+E01E`      |
|  `y`       | `U+E006`      | `U+E020`      |
|  `e`       | `U+E00C`      | `U+E026`      |
|  `i`       | `U+E00F`      | `U+E029`      |
|  `a`       | `U+E012`      | `U+E02C`      |
|  `o`       | `U+E015`      | `U+E02F`      |
|  `ë`       | `U+E017`      | `U+E031`      |


| **Diacritics** | **Narrow** | **Wide** |
| :------------: | :--------: | :------: |
|   `ṇ`          | `U+E037`   | `U+E03C` |
|   `ṃ`          | `U+E038`   | `U+E03D` |
|   `l`          | `U+E039`   | `U+E03E` |


| **Tones** | **Codepoint** |
| :-------: | :-----------: |
|  `\m`     | `U+E04C`      |
|  `\h`     | `U+E04D`      |
|  `\r`     | `U+E04E`      |
|  `\l`     | `U+E04F`      |


| **Pauses** | **Codepoint** | **Description** |
| :--------: | :-----------: | :-------------: |
|  `.`       | `U+E034`      | separate sentences (long pause) |
|  `\|`      |               | (short pause) |


| **Other** | **Uppercase** | **Lowercase** | **Narrow** | **Wide** |
| :-------: | :-----------: | :-----------: | :--------: | :------: |
|  `ä`      | `U+E009`      | `U+E023`      | `U+E036`   | `U+E03B` |


In all of these tables, the first column contains the characters that are used in the input file. If you don't have a keyboard that can type those characters, you can just copy&paste them from here.


### Rules for the Input File

It is important to note that, in the input file, spaces (` ` and `\t`) are ignored. This means that, for example, the input file always produces the same output if the whitespace is the only difference.

However, the period (`.`) is used to end a sentence.


Here's are the rules of the OSL conlang:
- The first vowel of a sentence is always capitalized.
- If punctuation marks occur before a vowel, then said vowel will be capitalized.
- When words with two syllables have the first once stressed, then the second one will get its first letter capitalized. (Which in OSL is always a vowel.)
- Variadic words (i.e., words with arbitrary length) have the stress on the syllable that is capitalized.

Of course, these rules make sense only if you have learned the OSL conlang. If not, you can just ignore these rules.


Lastly, here are another character distinctions:
- **Narrow:** `a`, `e`, `z`, `ay`, `h`, `ey`, `d`, `s`, `i`, `o`, `b`, `u`, and `g`.
- **Wide:** `n`, `m`, `r`, `w`, and `wy`.




## Installation

As of now, the process of installation is very rudimentary as it requires Go to be installed. However, this may be changed in the future.


### Installing Go

In order to build the source, you need to have Go installed. If you don't have Go, you can download it from [Go's official website](https://go.dev/doc/install) and follow the instructions. *Remember to install the 1.23.1 version of Go.*


### Installing the Font

This program make use of a font that needs to be installed on the system for it to work. You can find the font [here](https://github.com/PlayerR9/osl_keyboard/src/fonts/izlarjan.ttf).


### Building

After the Go installation, you can build the project using the following command:
```bash
# clone the repository
$ git clone https://github.com/PlayerR9/osl_keyboard.git

# locate the project
$ cd osl_keyboard

# compile the project
$ go build cmd/main.go -o osl_keyboard
```



## Usage

### Executable

The executable `./osl_keyboard` is just a transformer that converts the content from an input file and save it to an output file.

To run the project, use the following command:
```bash
$ ./osl_keyboard <input_file> <output_file>
```
Where:
- `<input_file>`: The path of the file containing the text in OSL's romanization system.
- `<output_file>`: The path of the file where the transformed text will be saved.

After running the command, the result is saved in the specified output file; ready to be copy&pasted wherever you want.


NOTES:
> Of course, there is the `help` command that shows the commands that can be used.


### Library

This package can be imported by other Go projects to use the functionality as it was a library. This can be useful for making other kinds of tools.


## Contributing

Aside from helping with the code, for those who are interested in contributing, they can also help by creating new fonts; as long as it adheres to the character set specification, you are free to create any font you'd like.

If you do that, I'll definitely add it to the project; as well as your name in the contributor list.



## Roadmap

- [ ] Make a proper installer in order to avoid building from source.


## Acknowledgments

I'd like to thank the following people for their work/contributions:
- `@BinarySeries` for designing the writing system and for creating the font.
- `@maxx` for designing the writing system.
- `@di_ignoranza` for making the OSL conlang itself and for giving the permission to create this program.

and all members of the Hjoron community for their support.


Special thanks to FontStruct; check out their [website](https://fontstruct.com/) for more information.