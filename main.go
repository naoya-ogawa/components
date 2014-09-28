package main

import (
	"./pkg/requestfiles"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte("non data"))
}

func main() {
	http.HandleFunc("/view/", requestfiles.GetHandler("./public/", 6))

	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":80", nil)
}
