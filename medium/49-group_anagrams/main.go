package main

// we need to create a type to be able to hold it as a key in the map
// you cannot use a [] as a key for a map
type Key [26]byte

// helper function for converting a string to a 26 byte array.
// this 26 byte array is used to map each letter in the word to a byte
// in the alphabet and if the word contains a then we bump the value for a with 1.
// if there are two a's then the value would be 2. This way we can create a unique key defining that specific word by
// the number of characters in the word
// example of eyes
// a b c d e f g .. s .. y .. z
// 0 0 0 0 2 0 0 .. 1 .. 1 .. 0

func strKey(str string) (key Key) {
	// loop throught the characters in the word str
	for i := range str {
		// convert each letter in the word to the ascii character representation and subtract with the representation
		// of "a". Then bump with 1
		key[str[i]-'a']++
	}
	// in golang if you have named return types => (key Key) then a "return" statement
	// will return the values of those keys
	return
}

func groupAnagrams(strs []string) [][]string {
	// create a hasmap of Key(byte array) as the key and []string as the value
	groups := make(map[Key][]string)

	// loop through each word in the []string
	for _, v := range strs {
		// conver the word to the []byte representation for the key
		key := strKey(v)
		// append the []byte representation as the key with the value being the word to the hashmap
		groups[key] = append(groups[key], v)
	}

	// then create the result object [][]string with the initial capacity of the lenght of the groups map
	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		// then append each array from the hashmap to the array of arrays, result
		result = append(result, v)
	}

	return result
}
