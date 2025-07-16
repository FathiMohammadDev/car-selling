package common

import (
	"math/rand"
	"strings"

	"github.com/FathiMohammadDev/car-selling/config"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func GeneratePassword() string {
	var generatedPassword strings.Builder

	cfg := config.GetConfig()
	passwordLength := cfg.Password.MinLength

	minLowerCase := 3
	minUperCase := 3
	minNum := 3
	minSpecialChar := 2

	if !cfg.Password.IncludeChars {
		minSpecialChar = 0
	}

	if !cfg.Password.IncludeDigits {
		minNum = 0
	}

	if !cfg.Password.IncludeUppercase {
		minUperCase = 0
	}

	if !cfg.Password.IncludeLowercase {
		minLowerCase = 0
	}

	for i := 0; i < minSpecialChar; i++ {
		randomIndex := rand.Intn(len(specialCharSet))
		generatedPassword.WriteString(string(specialCharSet[randomIndex]))
	}
	for i := 0; i < minNum; i++ {
		randomIndex := rand.Intn(len(numberSet))
		generatedPassword.WriteString(string(numberSet[randomIndex]))
	}
	for i := 0; i < minLowerCase; i++ {
		randomIndex := rand.Intn(len(lowerCharSet))
		generatedPassword.WriteString(string(lowerCharSet[randomIndex]))
	}
	for i := 0; i < minUperCase; i++ {
		randomIndex := rand.Intn(len(upperCharSet))
		generatedPassword.WriteString(string(upperCharSet[randomIndex]))
	}

	remainingLength := passwordLength - minLowerCase - minNum - minSpecialChar - minUperCase

	for i := 0; i < remainingLength; i++ {
		randomIndex := rand.Intn(len(allCharSet))
		generatedPassword.WriteString(string(allCharSet[randomIndex]))
	}

	inRune := []rune(generatedPassword.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)

}
