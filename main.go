package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	//Command-line flags
	//port := flag.String("p", "3000", "port to serve on")
	//flag.Parse()

	checkFile("example.js")

}

func testHandle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Ok")
}

func checkFile(filename string) {
	extension := getFileExtension(filename)

	regexpFile, err := ioutil.ReadFile("rules/" + extension + ".rule")
	if err != nil {
		fmt.Print(err)
		return
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return
	}

	re := regexp.MustCompile(string(regexpFile))
	matches := re.FindAllIndex(content, -1)
	for i := 0; i < len(matches); i++ {
		line := findLine(content, matches[i][0])
		log.Printf("Warning! Found rule on line: %d", line)
	}
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
func getFileExtension(filename string) string {
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
