package handler

import (
	"net/http"

	"github.com/eamirgh/open-quotes/conf"
	"github.com/eamirgh/open-quotes/quote"
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
