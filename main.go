package main

// handlers are like controllers in MVC
// router stores the url patterns and their corresponding handlers
// web server listens for incoming requests
import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func pHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Future I Love Me Professional page"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func main() {
	// create a new servemux, which is an HTTP request multiplexer
	// a multiplexer is responsible for matching incoming requests against a list of
	// registered routes and calling the associated handler for the route when a match is found
	// the servemux is also known as a router
	// fact check Copilot's suggestion: https://pkg.go.dev/net/http#ServeMux :)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/mark", pHome)
	mux.HandleFunc("/snippet", showSnippet)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
