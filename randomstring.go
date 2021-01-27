package randomstring

import (
	"crypto/rand"
	"log"
	"strings"
)

// Log and error message and exit the program
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Types of characters
type characterTypes struct {
	Lower   string
	Upper   string
	Number  string
	Special string
	All     string
}

// Make a `contains` function which determines
// if a substring is included in a pattern
func makePatternContainsFunc(pattern string) func(string) bool {
	return func(substring string) bool {
		return strings.Contains(pattern, substring)
	}
}

// Returns a retriver function which returns
// The set of characters
func getCharacterRetriver() func() characterTypes {
	characters := characterTypes{
		Lower:   "abcdefghijklmnopqrstuvwxyz",
		Upper:   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Number:  "0123456789",
		Special: `~!@#$%^&()_+-={}[];\',.`,
	}
	characters.All = characters.Lower + characters.Upper + characters.Number + characters.Special

	return func() characterTypes {
		return characters
	}
}

/*
RandomString generates a random string
specifying a given pattern and length.

Patterns:
	- 0: Numbers allowed
	- A: Uppercase allowed
	- a: Lowercase allowed
	- !: Special characters allowed
	- *: All characters allowed
*/
func RandomString(pattern string, length int) string {
	characters := getCharacterRetriver()()
	dictionary := ""
	patternContains := makePatternContainsFunc(pattern)

	if patternContains("*") {
		dictionary = characters.All
	} else {
		if patternContains("a") {
			dictionary += characters.Lower
		}

		if patternContains("A") {
			dictionary += characters.Upper
		}

		if patternContains("0") {
			dictionary += characters.Number
		}

		if patternContains("!") {
			dictionary += characters.Special
		}
	}

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	failOnError(err, "Could not successfully read bytes.")
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}

	return string(bytes)
}
