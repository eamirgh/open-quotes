package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eamirgh/open-quotes/conf"
	"github.com/eamirgh/open-quotes/quote"
	"github.com/gorilla/mux"
)

func httpResp(w http.ResponseWriter, code int, msg string) (int, error) {
	w.WriteHeader(code)
	return w.Write([]byte(msg))
}
func Ping(w http.ResponseWriter, _ *http.Request) {
	httpResp(w, http.StatusOK, "Pong!")
}

func Index(w http.ResponseWriter, _ *http.Request) {
	Q, err := quote.RandomQuote("en_US")
	if err != nil {
		httpResp(w, http.StatusInternalServerError, "something went wrong")
		return
	}
	err = conf.Tpl.ExecuteTemplate(w, "index.gohtml", Q[0])
	if err != nil {
		httpResp(w, http.StatusInternalServerError, "something went wrong")
		return
	}
}
func GetQuotes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var err error
	locale := conf.Locales[0]
	count := 1
	if _, ok := params["locale"]; ok {
		locale = params["locale"]
	}
	if _, ok := params["count"]; ok {
		count, err = strconv.Atoi(params["count"])
		if err!= nil{
			httpResp(w, http.StatusBadRequest, "something went wrong")
			return
		}
	}
	qs, err := quote.RandomQuotes(locale,count)
	if err!= nil{
		httpResp(w, http.StatusBadRequest, "something went wrong")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	err = e.Encode(qs)
	if err != nil {
		httpResp(w, http.StatusInternalServerError, "something went wrong")
		return
	}
}
