package randomstring

import (
	"crypto/rand"
	"log"
	"strings"
)

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

	if strings.Contains(pattern, "a") {
		dictionary += types.Lower
	}

	if strings.Contains(pattern, "A") {
		dictionary += types.Upper
	}

	if strings.Contains(pattern, "0") {
		dictionary += types.Number
	}

	if strings.Contains(pattern, "!") {
		dictionary += types.Special
	}

	if strings.Contains(pattern, "*") {
		dictionary = types.All
	}

	var bytes = make([]byte, length)
	_, err := rand.Read(bytes)
	failOnError(err, "Could not successfully read bytes.")
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}

	return string(bytes)
}
