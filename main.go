package main

// handlers are like controllers in MVC
// router stores the url patterns and their corresponding handlers
// web server listens for incoming requests
import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func pHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Future I Love Me Professional page"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// functionally equivalent to snippetCreate's if statement
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "GET right wit yoself, sucka", 405)
		return
	}

	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// w.Header().Set() must be called before w.WriteHeader() or w.Write()
		// otherwise, the header will not be set
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader() can only be called once per request
		w.WriteHeader(405)

		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func snippetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a list of snippets..."))
}

func main() {
	// create a new servemux, which is an HTTP request multiplexer
	// a multiplexer is responsible for matching incoming requests against a list of
	// registered routes and calling the associated handler for the route when a match is found
	// the servemux is also known as a router
	// fact check Copilot's suggestion: https://pkg.go.dev/net/http#ServeMux :)

	// Always create your own servemux, don't use http.DefaultServeMux
	// it's sharedb with all go programs by default
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/mark", pHome)
	mux.HandleFunc("/snippet/", snippetIndex)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
