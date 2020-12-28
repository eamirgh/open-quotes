package quote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/eamirgh/open-quotes/conf"
	"github.com/pkg/errors"
)

var Quotes map[string][]Quote

type Quote struct {
	From string `json:"from"`
	Text string `json:"text"`
	URL  string `json:"-"`
}

var twitter = regexp.MustCompile(`^(@)`)
var twitterURL = "https://twitter.com/"

func newQuote(from string, txt string) Quote {
	href := twitter.ReplaceAllString(from, twitterURL)
	if href == from {
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
	for _, l := range conf.Locales {
		jsonFile, err := os.Open(fmt.Sprintf("data/%s.json", l))

		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("could not find %v.json", l))
		}
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("could not read %v.json", l))
		}
		err = jsonFile.Close()
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("could not close %v.json", l))
		}
		var qs jsonQuotes
		err = json.Unmarshal(bytes.TrimPrefix(byteValue, []byte("\xef\xbb\xbf")), &qs)
		if err != nil {
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			return errors.Wrap(err, fmt.Sprintf("could not unmarshal %v.json", l))
		}
		for _, q := range qs.Data {
			Quotes[l] = append(Quotes[l], newQuote(q.From, q.Text))
		}
	}
	return nil
}

// func generateVerificationToken(length int) string {
// 	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
// 	b := make([]byte, length)
// 	n, err := io.ReadAtLeast(rand.Reader, b, length)
// 	if n != length {
// 		panic(err)
// 	}
// 	for i := 0; i < len(b); i++ {
// 		b[i] = table[int(b[i])%len(table)]
// 	}
// 	return string(b)
// }

func randomizeQuotes(quotes []Quote, count int) []Quote {
	rand.Seed(time.Now().UnixNano() * time.Now().UnixNano())
	rand.Shuffle(len(quotes), func(i, j int) { quotes[i], quotes[j] = quotes[j], quotes[i] })
	return quotes[:count]
}
func RandomQuotes(locale string, count int) ([]Quote, error) {
	if quotes, ok := Quotes[locale]; ok {
		if len(quotes) <= count {
			return quotes, nil
		}
		return randomizeQuotes(quotes, count), nil
	}
	return nil, errors.New("no quotes for your locale")
}

func RandomQuote(locale string) ([]Quote, error) {
	return RandomQuotes(locale, 1)
}
