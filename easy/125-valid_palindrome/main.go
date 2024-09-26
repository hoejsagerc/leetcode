package main

import (
	"regexp"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// creating regex for removing all non a-z, A-Z and 0-9 characters
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	// handling if s is an empty string
	if s == " " {
		return true
	}

	// cleaning the string with the regex expression
	clean := re.ReplaceAllString(s, "")
	// converting the string to lowercase
	clean = strings.ToLower(clean)
	// initiaing the two pointers
	left, right := 0, len(clean)-1

	// while looping for as long as the left pointer is lower than the right pointer
	for left < right {
		// if the to pointers are the same value then add and subtract 1 and continue to next value
		if clean[left] == clean[right] {
			right--
			left++
			continue
		}
		// if they arent equal then return false because it is not a palindrome then.
		return false
	}
	// if all of the characters where equal then return true
	return true
}

// A more performant way to clean the string instead of using regex
func anotherWay(s string) bool {
	// creating an anonymous function which accepts a rune and returns a rune
	f := func(r rune) rune {
		// using unicode to check if character is a letter or a number
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			// if not then return -1 which in the strings.Map() context will remove
			// the character from the string
			return -1
		}
		// returns the character in lowercase
		return unicode.ToLower(r)
	}
	// Map function will loop over each character in the string and apply function f
	s = strings.Map(f, s)

	// same as the other method
	left, right := 0, len(s)
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
