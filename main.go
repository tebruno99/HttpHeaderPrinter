package main

import (
	"embed"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", AllHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8999", r)
}

func AllHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	t, err := template.ParseFS(tpls, "*")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if err = t.ExecuteTemplate(w, "index.tmpl", r.Header); err != nil {
		panic(err)
	}
}

//go:embed *.tmpl
var tpls embed.FS
