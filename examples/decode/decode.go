package main

import (
	"github.com/devplayg/eggcrate"
	"net/http"
)

/*
	/tmp/static/css/app.css
	/tmp/static/css/fa.css
	/tmp/static/js/app.js
*/

var staticMap map[string][]byte

func main() {
	m, err := eggcrate.Decode(assetData)
	if err != nil {
		panic(err)
	}
	staticMap = m

	http.HandleFunc("/assets/js/app.js", jsHandler)
	http.HandleFunc("/assets/css/app.css", cssHandler)
	http.HandleFunc("/assets/css/fa.css", cssHandler)

	http.ListenAndServe(":8000", nil)
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(staticMap[r.RequestURI])
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "text/css")
	w.Write(staticMap[r.RequestURI])
}
