package quote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

var Quotes map[string][]Quote

type Quote struct {
	From string
	Text string
	URL string
}

var twitter = regexp.MustCompile(`^(@)`)
var twitterURL = "https://twitter.com/"
var locales = [] string {"en_US"}


func newQuote(from string, txt string) Quote {
	href := twitter.ReplaceAllString(from, twitterURL)
	if href == from{
		href = "#"
	}
	return Quote{
		From: from,
		Text: txt,
		URL:  href,
	}
}

type jsonQuote struct {
	Text string `json:"text"`
	From string `json:"from"`
}

type jsonQuotes struct {
	Data []jsonQuote `json:"data"`
}

func Init() error {
	Quotes = make(map[string][]Quote)
	for _, l := range locales {
		jsonFile, err := os.Open(fmt.Sprintf("data/%s.json", l))

		if err!=nil {
			return errors.Wrap(err, fmt.Sprintf("could not find %v.json", l))
		}
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err!=nil {
			return errors.Wrap(err, fmt.Sprintf("could not read %v.json", l))
		}
		err = jsonFile.Close()
		if err!=nil {
			return errors.Wrap(err, fmt.Sprintf("could not close %v.json", l))
		}
		var qs jsonQuotes
		err = json.Unmarshal(bytes.TrimPrefix(byteValue,  []byte("\xef\xbb\xbf")), &qs)
		if err!=nil {
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			return errors.Wrap(err, fmt.Sprintf("could not unmarshal %v.json", l))
		}
		for _, q := range qs.Data{
			Quotes[l] = append(Quotes[l], newQuote(q.From,q.Text))
		}
	}
	return nil
}