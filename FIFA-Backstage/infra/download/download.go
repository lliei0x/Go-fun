package download

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrDownloader      = errors.New("download html failed")
	ErrSeleniumService = errors.New("selenium service failed")
	ErrWebDriver       = errors.New("web driver failed")
	ErrWebDriverGet    = errors.New("web driver get url failed")
)

// Downloader get document by url
func Downloader(url string) (*goquery.Document, error) {

	request, err := http.NewRequest("Get", url, nil) //提交请求
	if err != nil {
		return nil, ErrDownloader
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36") // 添加 header 信息
	client := http.DefaultClient                                                                                                                       //生成client 参数为默认
	response, err := client.Do(request)                                                                                                                //处理返回结果
	if err != nil {
		return nil, ErrDownloader
	}
	defer response.Body.Close()
	return goquery.NewDocumentFromReader(response.Body)
}
