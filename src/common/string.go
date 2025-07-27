package common

import (
	CryptoRand "crypto/rand"
	"io"
	"math/rand"
	"strings"
	"unicode"

	"github.com/FathiMohammadDev/car-selling/config"
	"golang.org/x/crypto/bcrypt"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

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

func GenerateOtp() (string, error) {
	cfg := config.GetConfig()
	const digits = "0123456789"

	var otp strings.Builder
	otp.Grow(cfg.Otp.Digits)

	buffer := make([]byte, cfg.Otp.Digits)

	_, err := io.ReadFull(CryptoRand.Reader, buffer)
	if err != nil {
		return "", err
	}

	for _, b := range buffer {
		otp.WriteByte(digits[int(b)%cfg.Otp.Digits])
	}
	return otp.String(), nil
}

func CheckPassword(password string) bool {
	cfg := config.GetConfig()
	if len(password) < cfg.Password.MinLength {
		return false
	}

	if cfg.Password.IncludeChars && !HasLetter(password) {
		return false
	}

	if cfg.Password.IncludeDigits && !HasDigits(password) {
		return false
	}

	if cfg.Password.IncludeLowercase && !HasLower(password) {
		return false
	}

	if cfg.Password.IncludeUppercase && !HasUpper(password) {
		return false
	}

	return true
}

func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasDigits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
