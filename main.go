package main

import (
	"fmt"
	"strings"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz";

/*
* Encode provided keyword
* Returns an array number representing each alphabet of the given keyword based on the index of each letter
*/
func encodeKeywordWithNum(keyword string) []int {
	var numberRepresentation []int;
	for index:=0; index < len(keyword); index++ {
		var element byte = keyword[index]; 
		var alphabetsAsArray []string = StringToArray(alphabets);
		var elementIndexInAlphabets = IndexOf(alphabetsAsArray, element)
		numberRepresentation = append(numberRepresentation, elementIndexInAlphabets)
	}
	return numberRepresentation
}

func encryptMessage (message string, keyword string) string {
	// To avoid any complexities, the message is converted to lowercase as there is no provision for uppercase letters
	message = strings.ToLower(message);

	var encodedKeyword = encodeKeywordWithNum(keyword)
	var encryptedMessage string = ""

	// This is used to track the current position of the keyword as iterated
	var currentKeywordPostition int = 0

	var alphabetsAsArray []string = StringToArray(alphabets)

	for index:=0; index < len(message); index++ {
		var letter byte = message[index]
		var letterOriginalIndex = IndexOf(alphabetsAsArray, letter)
		var shiftPosition = encodedKeyword[currentKeywordPostition]

		/*
		* We only want to deal with letters not symbols or numbers
		* Any other character will skip any transformation and simply be appended to the encrypted message
		*/
		if (ContainsLetter(alphabetsAsArray, letter)) {
			currentKeywordPostition++
			if (currentKeywordPostition >= len(encodedKeyword)) {
				// Reset the current keyword position when exhausted
				currentKeywordPostition = 0
			}
			// replacementValueIndex is the index of the new value that is meant to replace a letter
			var replacementValueIndex = letterOriginalIndex + shiftPosition

			/* 
			* if - the index of a value is more than the total number of alphabets (26), we find the  
			* remainder (i.e remainder = value % 26) and starting from zero, we count until we reach the
			* alphabet holding the value of that remainder. We then use that alphabet as the replacement value
			* else - we simply find the replacement value and append the encrypted message
			*/
			if(replacementValueIndex >= 26) {
				var remainder int = replacementValueIndex % 26; 
				var replacementValue = alphabetsAsArray[remainder]
				encryptedMessage = encryptedMessage + replacementValue
			} else {
				var replacementValue = alphabetsAsArray[replacementValueIndex]
				encryptedMessage = encryptedMessage + replacementValue
			}
		} else {
			encryptedMessage = encryptedMessage + string(letter)
		}
	}

	fmt.Println("Message successfully decrypted:")
	fmt.Println("Message", "=>", message)
  fmt.Println("Encrypted message", "=>", encryptedMessage)
	return encryptedMessage
}

func decryptMessage(message string, keyword string) string {
  var encodedKeyword = encodeKeywordWithNum(keyword)
	var decryptedMessage string = ""

	// This is used to track the current position of the keyword as iterated
	var currentKeywordPostition int = 0
	var alphabetsAsArray []string = StringToArray(alphabets)

	for index:=0; index < len(message); index++ {
		var letter byte = message[index]
		var letterOriginalIndex = IndexOf(alphabetsAsArray, letter)
		var shiftPosition = encodedKeyword[currentKeywordPostition]

		/*
		* We only want to deal with letters not symbols or numbers
		* Any other character will skip any transformation and simply be appended to the encrypted message
		*/
		if (ContainsLetter(alphabetsAsArray, letter)) {
			currentKeywordPostition++
			if (currentKeywordPostition >= len(encodedKeyword)) {
				// Reset the current keyword position when exhausted
				currentKeywordPostition = 0
			}
			// replacementValueIndex is the index of the new value that is meant to replace a letter
			var replacementValueIndex = letterOriginalIndex - shiftPosition

			/* 
			* if - the index of a value is less than the index of the first alphabet 'a'(0) i.e its a negative number, we find the
			* sum (i.e sumResult = value + 26) and starting from 26, we count backward until we reach the
			* alphabet holding the value of that sum. We then use that alphabet as the replacement value
			* else - we simply find the replacement value and append the encrypted message
			*/
			if(replacementValueIndex < 0) {
				var sumResult int =  replacementValueIndex + 26
				var replacementValue = alphabetsAsArray[sumResult]
				decryptedMessage = decryptedMessage + replacementValue
			} else {
				var replacementValue = alphabetsAsArray[replacementValueIndex]
				decryptedMessage = decryptedMessage + replacementValue
			}
		} else {
			decryptedMessage = decryptedMessage + string(letter)
		}
	}

	fmt.Println("Message successfully decrypted:")
	fmt.Println("Encrypted message", "=>", message)
  fmt.Println("Decrypted message", "=>", decryptedMessage)
	return decryptedMessage
}

func main() {
	// change the message and keyword to whatever you want
	encryptMessage("MESSAGE HERE", "KEYWORDHERE") 
	decryptMessage("MESSAGE HERE", "KEYWORDHERE")
}