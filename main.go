package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	//Command-line flags
	port := flag.String("p", "3000", "port to serve on")
	flag.Parse()

	router := mux.NewRouter()

	// Handle API routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", testHandle).Methods("GET")

	// Serve static files from public directory
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))

	checkFile("example.js")

	log.Println("Starting server on port " + *port)
	//log.Fatal(http.ListenAndServe(":"+*port, router))

}

func testHandle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Ok")
}

func checkFile(filename string) {
	extension := strings.Split(filename, ".")[1]

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

	log.Println(findLine(content, 21))
	re := regexp.MustCompile(string(regexpFile))
	log.Println(re.FindAllIndex([]byte(content), -1))
	log.Println(re.FindAllStringIndex(string(content), -1))
}

func findLine(str []byte, index int) int {
	cont := 1
	for i := 0; i < index; i++ {
		if str[i] == 10 {
			cont++
		}
	}
	return cont
}
