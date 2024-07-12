package util

import (
	"fmt"
	"log"
	"regexp"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ScriptValidationCheck(hash string) (string, error) {
	// Regex pattern for the desired hash format
	pattern := `^[a-f0-9]{64} *-$`
	matched, err := regexp.MatchString(pattern, hash)
	if err != nil {
		return "", err
	}
	if !matched {
		return "", fmt.Errorf("hash does not match the required format")
	}
	return hash, nil
}
