package preprocess

import (
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func readUtf16LE(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	
	decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	utf8Data, _, err := transform.Bytes(decoder, data)
	if err != nil {
		log.Fatal(err)
	}

	str := string(utf8Data)
	return str
}

func TextRead(path string, minLen int) []string {
	fileText := readUtf16LE(path)

	textList := make([]string, 0)
	totalLen := 0
	tempStr := ""
	for _, str := range strings.Split(string(fileText), "\n") {
		totalLen += len(str)
		tempStr += str + "\n"
		if totalLen >= minLen {
			textList = append(textList, tempStr)
			totalLen = 0
			tempStr = ""
		}
	}
	return textList
}
