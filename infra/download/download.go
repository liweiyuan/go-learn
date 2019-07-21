package download

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"errors"
)

var (
	ErrDownloader      = errors.New("download html failed")
	ErrSeleniumService = errors.New("selenium service failed")
	ErrWebDriver       = errors.New("web driver failed")
	ErrWebDriverGet    = errors.New("web driver get url failed")
)

func Download(url string) (*goquery.Document, error) {

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, ErrDownloader
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")

	client := http.DefaultClient

	response, err := client.Do(request)
	if err != nil {
		return nil, ErrDownloader
	}
	defer response.Body.Close()

	return goquery.NewDocumentFromReader(response.Body)

}
