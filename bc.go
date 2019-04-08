package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	Regex   string
	Level   int
	Message string
}

var logLevel []string

func main() {
	//Setting log level
	logLevel = []string{"Warning"}

	//Command-line flags
	filePath := flag.String("f", "", "file to check")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("No file selected")
		return
	}

	checkFile(*filePath)
}

func ruleParser(filePath string) []rule {
	fOpen, _ := os.Open(filePath)
	defer fOpen.Close()

	regexRule := regexp.MustCompile(`(?:regex= ).*?(?: level=)`)
	levelRule := regexp.MustCompile(`(?:level= ).*?(?: msg=)`)
	msgRule := regexp.MustCompile(`(?:msg= )(.*)`)

	parsedRules := []rule{}

	fscanner := bufio.NewScanner(fOpen)
	for fscanner.Scan() {
		rawText := fscanner.Text()

		regexMatch := strings.TrimSuffix(strings.TrimPrefix(regexRule.FindString(rawText), "regex= "), " level=")
		levelRaw := strings.TrimSuffix(strings.TrimPrefix(levelRule.FindString(rawText), "level= "), " msg=")
		msgMatch := strings.TrimPrefix(msgRule.FindString(rawText), "msg= ")

		levelMatch, err := strconv.Atoi(levelRaw)
		if err != nil {
			log.Fatal("Fail ATOI")
		}
		newRule := rule{
			Regex:   regexMatch,
			Level:   levelMatch,
			Message: msgMatch,
		}
		parsedRules = append(parsedRules, newRule)
	}
	
	return parsedRules
}

func checkFile(filename string) {
	extension := getFileExtension(filename)

	//load rules
	rules := ruleParser("rules/" + extension + ".rule")

	//load file content
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return
	}

	cont := 0
	for i := 0; i < len(rules); i++ {
		re := regexp.MustCompile(string(rules[i].Regex))
		matches := re.FindAllIndex(content, -1)
		for j := 0; j < len(matches); j++ {
			cont++
			line := findLine(content, matches[j][0])
			fmt.Printf("%s:%d: %s: %s\n", filename, line, logLevel[rules[i].Level], rules[i].Message)
		}
	}
	fmt.Printf("total: %d warnings", cont)
}

//Finds the line of a given index in a string
func findLine(str []byte, index int) int {
	cont := 1
	for i := 0; i < index; i++ {
		if str[i] == 10 {
			cont++
		}
	}
	return cont
}

//Gets the file extension from a filename
func getFileExtension(filePath string) string {
	split := strings.Split(filePath, `\`)
	lastIndex := len(split) - 1
	filename := split[lastIndex]
	return strings.Split(filename, ".")[1]
}

//Reads a file content
func readFileContent(filePath string) ([]byte, error) {
	fileRead, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Print(err)
		return []byte(""), err
	}
	return fileRead, nil
}
