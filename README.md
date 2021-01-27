# Generating Random Strings in Go
> *Inspired by [randomatic](https://github.com/jonschlinkert/randomatic)*

First package in Go for generating random strings with a given pattern and length. 

## Usage
* Generate a string of length 10 consisting of only numeric characters: `RandomString("0", 10)`
* Generate a string of length 5 consisting of only special characters: `RandomString("!", 5)`
* Generate a string of length 15 consisting of uppercase or lowercase characters only: `RandomString("Aa", 15)`
* Generate a string of length 2 consisting of any characters: `RandomString("*", 2)`
