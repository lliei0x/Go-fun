package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"leeif.me/Go-fun/spider-training/infra/errors"
)

// GetHttpResponse get response for http request
func GetHttpResponse(url string, ok bool) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorRequest
	}
	// request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.ErrorResponse
	}
	defer response.Body.Close()
	if response.StatusCode <= 200 && response.StatusCode >= 300 {
		return nil, errors.ErrorStatusCode
	}
	if ok {
		utf8ContentReader := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8ContentReader)
	} else {
		return ioutil.ReadAll(response.Body)
	}
}

// PostHttpResponse ...
func PostHttpResponse(url string, body string, ok bool) ([]byte, error) {

	payload := strings.NewReader(body)
	requests, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, errors.ErrorRequest
	}
	requests.Header.Add("X-Requested-With", "XMLHttpRequest")
	requests.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	requests.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(requests)
	if err != nil {
		return nil, errors.ErrorResponse
	}

	fmt.Println(response.StatusCode)

	defer response.Body.Close()
	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}
	return ioutil.ReadAll(response.Body)

}
