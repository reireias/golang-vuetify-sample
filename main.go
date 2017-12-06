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
	router.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./web/public"))))
	http.ListenAndServe(":8080", router)
}
