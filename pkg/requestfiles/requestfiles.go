package requestfiles

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const ContentType = "Content-Type"

func GetHandler(path string, position uint8) func(http.ResponseWriter, *http.Request) {

	return func(res http.ResponseWriter, r *http.Request) {

		allpath := path + r.URL.Path[position:]

		fmt.Printf("allpath:%v", allpath)
		b, err := ioutil.ReadFile(allpath)
		if err != nil {
			http.Error(res, allpath+"：はありません。", http.StatusInternalServerError)
			return
		}
		res.Header().Set(ContentType, "text/html")
		res.Write(b)
	}
}
