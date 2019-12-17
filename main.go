package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	length := 16

	if len(os.Args) == 2 {
		arg := os.Args[1]
		argLength, err := strconv.Atoi(arg)
		if err != nil {
			panic("length must be a number")
		}

		length = argLength
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < length; i++ {
		letter := string(getRandomLetter())
		print(letter)
	}

	println()
}

func getRandomLetter() rune {
	maxNumber := 25 + 25 + 9
	randomNumber := rand.Intn(maxNumber)

	return getLetterByInt(randomNumber)
}

func getLetterByInt(number int) rune {
	if number <= 25 {
		return rune('A' + number)
	} else if number <= 25+25 {
		return rune('a' + number - 25)
	}

	return rune('0' + number - 25 - 25)
}
