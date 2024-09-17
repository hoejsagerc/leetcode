package main

import (
	"reflect"
	"sort"
	"strings"
)

// -------------------- METHOD 2 --------------------------
// This will have a O(1)
// ====> FUNC FACT: leetcode performance (most performant solution)
// 			runtime: 	17ms 	-> beats 5.50%
//			memory: 	6.64mb	-> beats 5.27%

func isAnagram(s string, t string) bool {
	sSlice := strings.Split(s, "")
	tSlice := strings.Split(t, "")

	sort.Strings(sSlice)
	sort.Strings(tSlice)
	return reflect.DeepEqual(sSlice, tSlice)
}

// -------------------- METHOD 1 --------------------------
// ====> FUNC FACT: leetcode performance (most performant solution)
// 			runtime: 	8ms 	-> beats 41.43%
//			memory: 	6.58mb	-> beats 5.27%

// func isAnagram(s string, t string) bool {
// 	sSlice := strings.Split(s, "") // convert to []string
// 	tSlice := strings.Split(t, "")
// 	// if the lengt of the two strings does not match then it can't be an anagram
// 	if len(sSlice) != len(tSlice) {
// 		return false
// 	}

// 	// get the charMap with occrences of each character for both strings
// 	sMap := getCharMap(sSlice)
// 	tMap := getCharMap(tSlice)

// 	// for each character and occurence count in the sMap
// 	for char, count := range sMap {
// 		// if the character exists in the tMap
// 		if _, ok := tMap[char]; ok {
// 			// if the occurence count of that character in the tMap is not the same
// 			// as the occurence count of that character in the sMap, then return false
// 			if tMap[char] != count {
// 				return false
// 			}
// 		} else {
// 			// if the character does not exists in the tMap then return false
// 			return false
// 		}
// 	}
// 	// else just return true
// 	return true
// }

// // {key => character, value = count of characters}
// func getCharMap(chars []string) map[string]int {
// 	// create a new map with a character from a string as the key, and the count of occurences of the character
// 	// in the string, as the value
// 	var charMap = make(map[string]int)

// 	// loop through all the characters in the string
// 	for _, c := range chars {
// 		// if character already exists in the map
// 		if _, ok := charMap[c]; ok {
// 			// then add 1 to the occrence count
// 			charMap[c] = charMap[c] + 1
// 		} else {
// 			// else (meaning the char does not exists in the map) then add it to the map
// 			charMap[c] = 1
// 		}
// 	}
// 	return charMap
// }
