package artGenerator

import (
	"fmt"
	"os"
	"strings"
)

// Takes the input string and path to the banner file to use
func GenerateArt(inputString, bannerPath string) (string, error) {

	if inputString == "" {
		return "", nil
	}

	if inputString == "\\n" || inputString == "\n" {
		return "\n", nil
	}

	standardArts, err := os.ReadFile(bannerPath)

	if err != nil {
		return "", fmt.Errorf("Error reading banner, %v", err)
	}

	convArt := strings.Split(string(standardArts), "\n")

	// handle possible escape new line characters
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")

	words := strings.Split(inputString, "\n")

	var result strings.Builder

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}
		var wordArt strings.Builder

		for charHeight := 1; charHeight <= 8; charHeight++ {
			for _, ch := range word {
				if ch < 32 || ch > 126 {
					continue
				}

				chAsciiIndex := int(ch) - 32

				asciiArtIndex := (chAsciiIndex * 9) + charHeight
				char := convArt[asciiArtIndex]
				wordArt.WriteString(char)
			}
			wordArt.WriteString("\n")
		}
		result.WriteString(wordArt.String())
	}
	return result.String(), nil
}
