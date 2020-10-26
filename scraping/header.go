package scraping

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func NewRequest(url string, siteName string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("referrer", fmt.Sprintf("https://%s",siteName))
	req.Header.Add("user-agent", "user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	req.Header.Add("accept-language", "en-US,en;q=0.9,ja;q=0.8")

	client := new(http.Client)
	resp, err := client.Do(req)
	return resp, errors.WithStack(err)
}
