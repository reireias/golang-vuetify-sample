package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HelloHandler retrns "hello"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HelloWorld")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":8080", router)
}
