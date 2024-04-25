package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	inputFile, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Read input line by line, process each line, and write to output file
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	for scanner.Scan() {
		inputLine := scanner.Text()
		outputLine := ReplaceHex(inputLine)
		outputLine = ReplaceBin(outputLine)
		outputLine = ReplaceWithAn(outputLine)
		outputLine = FormatQuotes(outputLine)
		outputLine = ReplaceCase(outputLine)
		outputLine = FormatPunctuation(outputLine)

		writer.WriteString(outputLine + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}
}

// ReplaceHex replaces the word before "(hex)" with its decimal equivalent
func ReplaceHex(input string) string {
	regul := regexp.MustCompile(`(\b[0-9A-Fa-f]+) \((hex)\)`)
	output := regul.ReplaceAllStringFunc(input, func(match string) string {
		parts := regul.FindStringSubmatch(match)
		if len(parts) == 3 {
			hexValue := parts[1]
			decimalValue, err := strconv.ParseInt(hexValue, 16, 64)
			if err == nil {
				return strconv.FormatInt(decimalValue, 10)
			}
		}
		return match // Return the original match if there's an error
	})
	return output
}

// ReplaceBin replaces the word before "(bin)" with its decimal equivalent
func ReplaceBin(input string) string {
	regul := regexp.MustCompile(`(\b[01]+) \((bin)\)`)
	output := regul.ReplaceAllStringFunc(input, func(match string) string {
		parts := regul.FindStringSubmatch(match)
		if len(parts) == 3 {
			binValue := parts[1]
			decimalValue, err := strconv.ParseInt(binValue, 2, 64)
			if err == nil {
				return strconv.FormatInt(decimalValue, 10)
			}
		}
		return match
	})
	return output
}

// ReplaceLowercase converts the word before "(low)" to its lowercase version
/*func ReplaceLowercase(input string) string {
	re := regexp.MustCompile(`\b(\w+) \((low)\)`)
	output := re.ReplaceAllStringFunc(input, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 3 {
			word := strings.ToLower(parts[1])
			return word
		}
		return match
	})
	return output
} */

// ReplaceWithAn replaces "a" with "an" if the next word begins with a vowel (a, e, i, o, u) or "h"
func ReplaceWithAn(input string) string {
	regul := regexp.MustCompile(`\b([aA])\s+(\b[aeiou]|h)`)
	output := regul.ReplaceAllStringFunc(input, func(match string) string {
		parts := regul.FindStringSubmatch(match)
		if len(parts) == 3 {
			if parts[1] == "a" {
				return "an " + parts[2]
			} else if parts[1] == "A" {
				return "An " + parts[2]
			}
		}
		return match
	})
	return output
}

func FormatPunctuation(text string) string {
	// Adjusting for spaces before punctuation and ensuring proper space after
	regul := regexp.MustCompile(`\s*([.,!?;:]+)(\s)`)
	replacement := func(match []byte) []byte {
		trimmed := strings.TrimSpace(string(match))
		if strings.HasSuffix(trimmed, " ") {
			return []byte(trimmed)
		} else {
			return []byte(trimmed + " ")
		}
	}
	formattedText := string(regul.ReplaceAllFunc([]byte(text), replacement))
	formattedText = regexp.MustCompile(`\s+([.!?]+)`).ReplaceAllString(formattedText, "$1")

	// Specifically adjust for the scenario of space before a comma in the middle of a sentence
	// This targets a comma that is either followed by a space (and more text) or not at the end of the text.
	formattedText = regexp.MustCompile(`\s+([,])`).ReplaceAllString(formattedText, "$1 ")

	// Handling groups of punctuation like "...", "!?", "?!"
	regul = regexp.MustCompile(`(\.\.\.|\!\?|\?\!)`)
	formattedText = regul.ReplaceAllStringFunc(formattedText, func(match string) string {
		return strings.ReplaceAll(match, " ", "")
	})

	return formattedText
}

func FormatQuotes(text string) string {
	// This regex looks for a pattern where a word or words are enclosed in single quotes
	// with optional spaces between the quotes and the word(s).
	// It captures the spaces (if any) and the word(s) within the quotes for replacement.
	regul := regexp.MustCompile(`'\s*(.*?)\s*'`)

	// The replacement removes the spaces by reconstructing the quoted word without them.
	return regul.ReplaceAllString(text, "'$1'")
}

/*func ApplyCapitalize(text string) string {
	// Regular expression to match "(cap)" and the word preceding it
	re := regexp.MustCompile(`(\w+)\s*\(cap\)`)

	// Replacement function that capitalizes the matched word
	// and replaces the "(cap)" marker with an empty string
	replacement := func(match string) string {
		word := strings.Title(re.FindStringSubmatch(match)[1])
		return word
	}

	// Use ReplaceAllStringFunc to apply the replacement function
	formattedText := re.ReplaceAllStringFunc(text, replacement)

	return formattedText
} */

func ReplaceCase(input string) string {
	input = RemoveSpace(input)
	// Split the input string into words
	words := strings.Fields(input)
	// Initialize a slice to store the modified words
	var modifiedWords []string
	// Initialize a variable to keep track of the current word index
	wordIndex := 0
	// Iterate over each word in the input string
	for i := 0; i < len(words); i++ {
		word := words[i]
		// Check if the word contains "(up", "(low", or "(cap" indicating an instruction
		if strings.Contains(word, "(up") || strings.Contains(word, "(low") || strings.Contains(word, "(cap") {
			// Default number of words to convert
			num := 1
			// Determine the transformation type based on the instruction
			transformType := ""
			if strings.Contains(word, "(up") {
				transformType = "up"
				// Parse the number of words to convert if specified
				if strings.Contains(word, "(up,") {
					fmt.Sscanf(word, "(up,%d)", &num)
				}
			} else if strings.Contains(word, "(low") {
				transformType = "low"
				// Parse the number of words to convert if specified
				if strings.Contains(word, "(low,") {
					fmt.Sscanf(word, "(low,%d)", &num)
				}
			} else if strings.Contains(word, "(cap") {
				transformType = "cap"
				// Parse the number of words to convert if specified
				if strings.Contains(word, "(cap,") {
					fmt.Sscanf(word, "(cap,%d)", &num)
				}
			}

			// Calculate the start index for transformation
			start := max(wordIndex-num, 0)
			// Apply transformation to the specified number of words
			for j := start; j < wordIndex; j++ {
				switch transformType {
				case "up":
					modifiedWords[j] = strings.ToUpper(modifiedWords[j])
				case "low":
					modifiedWords[j] = strings.ToLower(modifiedWords[j])
				case "cap":
					modifiedWords[j] = strings.Title(strings.ToLower(modifiedWords[j]))
				}
			}
			// Skip the current instruction word
			continue
		}
		// Add the current word to the list of modified words
		modifiedWords = append(modifiedWords, word)
		// Increment the word index
		wordIndex++
	}
	// Join the modified words back into a single string
	return strings.Join(modifiedWords, " ")
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func RemoveSpace(s string) string {
	s = strings.ReplaceAll(s, "(up, ", "(up,")
	s = strings.ReplaceAll(s, "(cap, ", "(cap,")
	s = strings.ReplaceAll(s, "(low, ", "(low,")
	return s
}
