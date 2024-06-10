package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tars47/go-choose-your-own-adventure/adventure"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "path", "story.json", "a path to json file that contains the story")
}

func main() {

	flag.Parse()

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.Handle("GET /{chapter}", adventure.NewStory(fileName))

	fmt.Println("--- Server Running on Port 4747 ---")
	log.Fatal(http.ListenAndServe(":4747", nil))
}
