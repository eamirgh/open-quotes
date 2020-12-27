package handler

import (
	"net/http"

	"github.com/eamirgh/open-quotes/conf"
)

func Ping(w http.ResponseWriter, _ *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong!"))
}

func Index(w http.ResponseWriter, _ *http.Request){
	conf.Tpl.ExecuteTemplate(w,"index.gohtml",nil)
}
