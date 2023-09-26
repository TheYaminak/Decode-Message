package main

import (
	"fmt"
)

func decodeMessage(encodedMessage string) string {

	reversedMessage := reverseString(encodedMessage)

	decodedMessage := ""
	for _, char := range reversedMessage {
		decodedChar := byte(char) - 5
		decodedMessage += string(decodedChar)
	}

	return decodedMessage
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	encodedMessage := "8757%TLSNRTI%TYSFX%YXJK[JI%YF%^YWFU%JPFP%JMY%SNTO"
	decodedMessage := decodeMessage(encodedMessage)
	fmt.Println("Decoded message:", decodedMessage)
}
