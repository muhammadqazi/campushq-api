package utils

import (
	"math/rand"
	"time"
)

func PasswordGenerator(length int, retries int) string {
	lowerChars := "abcdefghijklmnopqrstuvwxyz"
	upperChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars := "0123456789"
	specialChars := "!@#$%^&*"

	password := ""
	types := 0

	// Generate at least one character of each type
	password += string(getRandomChar(lowerChars))
	password += string(getRandomChar(upperChars))
	password += string(getRandomChar(numberChars))
	password += string(getRandomChar(specialChars))
	types = 4

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate remaining characters randomly
	for i := 4; i < length; i++ {
		randomType := random.Intn(4) + 1
		switch randomType {
		case 1:
			password += string(getRandomChar(lowerChars))
			types |= 1
		case 2:
			password += string(getRandomChar(upperChars))
			types |= 2
		case 3:
			password += string(getRandomChar(numberChars))
			types |= 4
		case 4:
			password += string(getRandomChar(specialChars))
			types |= 8
		}
	}

	if types != 15 && retries < 20 {
		return PasswordGenerator(length, retries+1)
	}

	return password
}

func getRandomChar(chars string) byte {
	return chars[rand.Intn(len(chars))]
}
