package main

import "strings"

// finds the index of a given letter in a given array
func IndexOf(arr []string, target byte) int {
	for i, str := range arr {
		for _, char := range str {
			if byte(char) == target {
					return i
			}
	  }
	}
	return -1 // Return -1 if the element is not found
}

// converts a given string to array
func StringToArray(str string) []string {
	var result []string
	for _, char := range str {
		result = append(result, string(char))
	}
	return result
}

// checks if an array contains a letter
func ContainsLetter(arr []string, target byte) bool {
	for _, str := range arr {
		if strings.ContainsRune(str, rune(target)) {
			return true
		}
	}
	return false
}