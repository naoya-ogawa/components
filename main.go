package main

import (
	"io/ioutil"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html")

	w.Write(b)
	//w.Write([]byte("This is example"))
	//fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":80", nil)
}
