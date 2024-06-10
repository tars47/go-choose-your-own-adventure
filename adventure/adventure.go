package adventure

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("./public/index.html", "./public/invalid.html"))
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story struct {
	data map[string]Chapter
}

func (s *Story) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	chapter := r.PathValue("chapter")
	data, ok := s.data[chapter]
	if !ok {
		tmpl.ExecuteTemplate(w, "invalid.html", s.data)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func (s *Story) parse(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to open %q \n err: %v", fileName, err)
	}
	defer file.Close()

	s.data = make(map[string]Chapter)

	if err := json.NewDecoder(file).Decode(&(s.data)); err != nil {
		log.Fatalf("unable to parse %q \n err: %v", fileName, err)
	}
}

func NewStory(fileName string) *Story {
	s := Story{}
	s.parse(fileName)
	return &s
}
