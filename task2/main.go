package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)


func wordCount(s string) map[string]int {

	lower := strings.ToLower(s)
	words := strings.Fields(lower)

	filtered := []string{}
	for _, word := range words {
		var builder strings.Builder
		for _, char := range word {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				builder.WriteRune(char)
			}
		}
		cleanWord := builder.String()
		if cleanWord != "" {
			filtered = append(filtered, cleanWord)
		}
	}

	wordCount := make(map[string]int)
	for _, word := range filtered {
		wordCount[word]++
	}

	return wordCount
}


func isPalindrome(s string) bool {

    lower := strings.ToLower(s)
    var builder strings.Builder
    for _, r := range lower {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            builder.WriteString(string(r))
        }
    } 

    s = builder.String()
	str := []rune(s)
	left, right := 0, len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}

		left++;
		right--

	}

	return true

}

func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert a string : ")
	s, _ := reader.ReadString('\n')
	word_count := wordCount(s)
	is_palindrome := isPalindrome(s)

	fmt.Println("Word Count : ", word_count)
	fmt.Println("Is Palindrome : ", is_palindrome)

}
