package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		keys, ok := r.URL.Query()["key"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}

		key := keys[0]

		w.Write([]byte(`
			{
				"Input Slice": ` + key + `
				"Output Slice":` + fmt.Sprintf("%s", flatten(key)) + `
			}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"Response": "Not Found"}`))

	}
}

func flatten(input string) []string {
	input = strings.ReplaceAll(input, "[", "")
	input = strings.ReplaceAll(input, "]", "")
	splitting := strings.Split(input, ",")
	return splitting
}

func main() {
	s := &server{}
	http.Handle("/flat", s)
	log.Fatal(http.ListenAndServe(":8181", nil))
}
