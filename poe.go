//Package poe let's you generate phonetic keys.
package poe

import (
	"math"
	"math/rand"
	"time"
)

const consonants = "bcdfghjklmnpqrstvwxyz"
const vowels = "aeiou"

//Generates a phonetic key
func KeyGen(keyLength int) string {
	key := ""
	start := Round(rand.Float64())
	for i := 0; i < keyLength; i++ {
		if i%2 == int(start) {
			key += randCon()
		} else {
			key += randVowel()
		}
	}
	return key
}

//Picks a random vowel
func randVowel() string {

	rand.Seed(time.Now().UnixNano())
	return string(vowels[rand.Intn(len(vowels))])
}

//Picks a random consonant
func randCon() string {
	rand.Seed(time.Now().UnixNano())
	return string(consonants[rand.Intn(len(consonants))])
}

//Rounds a float
func Round(f float64) float64 {
	return math.Floor(f + .5)
}
