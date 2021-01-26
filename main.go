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
	types := characterTypes{
		Lower:   "abcdefghijklmnopqrstuvwxyz",
		Upper:   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Number:  "0123456789",
		Special: `~!@#$%^&()_+-={}[];\',.`,
	}
	types.All = types.Lower + types.Upper + types.Number + types.Special

	dictionary := ""
	patternContains := makePatternContainsFunc(pattern)

	if patternContains("*") {
		dictionary = types.All
	} else {
		if patternContains("a") {
			dictionary += types.Lower
		}

		if patternContains("A") {
			dictionary += types.Upper
		}

		if patternContains("0") {
			dictionary += types.Number
		}

		if patternContains("!") {
			dictionary += types.Special
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
