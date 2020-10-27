package scraping

import (
	"com.github.yoshidan/go-anime-image/scraping/internal"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)


type Tsundora struct {
	url string
	name string
}

func NewTsundora(keyword string) *Tsundora {
	name := "tsundora.com"
	url := fmt.Sprintf("https://%s?s=%s", name, keyword)
	return &Tsundora{url, name}
}

func (p *Tsundora) Download() {
	scraper := &Scraper{
		thumbnailQuery: "div.home_tall_box > div.home-img > a"	,
		bigImageQuery: "div.single_inside_content > div.post-img > a",
		name: p.name,
	}
	scraper.Execute(p.url)
}

type Scraper struct {
	thumbnailQuery string
	bigImageQuery string
	name string
}

func (s *Scraper) Execute(url string) {
	fmt.Printf("start download from %s\n", url)
	doc, err := s.getFirstPage(url)
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
	s.downloadWhileHasNext(doc)
}

func (s *Scraper) downloadWhileHasNext(doc *goquery.Document) {
	doc.Find(s.thumbnailQuery).Each(func(index int, selection *goquery.Selection) {
		title := selection.AttrOr("title", "notitle")
		link := s.handleLink(selection)
		if link == nil {
			return
		}
		link.Find(s.bigImageQuery).Each(func(i int, selection *goquery.Selection) {
			imagePage := s.handleLink(selection)
			if imagePage == nil {
				return
			}
			imagePage.Find("img").Each(func(i int, selection *goquery.Selection) {
				s.handleImg(title, selection)
			})
		})
	})

	//seek nextPage if exist
	nextPage := doc.Find("div.pagenavi").Children().Last().AttrOr("href", "")
	if nextPage != "" {
		fmt.Printf("next page %s : sleep 30 second for avoiding server down \n", nextPage)
		time.Sleep(30 * time.Second)
		fmt.Printf("start download at %s\n", nextPage)
		response, err := internal.NewRequest(nextPage, s.name)
		if err != nil {
			fmt.Printf("skip %s because %+v", nextPage , err)
			return
		}
		defer response.Body.Close()
		nextDoc, _ := goquery.NewDocumentFromReader(response.Body)
		s.downloadWhileHasNext(nextDoc)
	}
}

func (s *Scraper) getFirstPage(url string) (*goquery.Document, error ){
	resp, e := internal.NewRequest(url, s.name)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	var data *http.Response
	if resp.StatusCode == 307 {
		targetUrl := resp.Header.Get("Location")
		targetResp, e := internal.NewRequest(targetUrl, s.name)
		if e != nil {
			return nil, e
		}
		defer targetResp.Body.Close()
		data = targetResp
	}else if resp.StatusCode == 200 {
		data = resp
	}else {
		return nil, fmt.Errorf("illegal status code %d",resp.StatusCode)
	}
	return goquery.NewDocumentFromReader(data.Body)
}


func (s *Scraper) handleLink(selection *goquery.Selection) *goquery.Document{
	url := selection.AttrOr("href", "")
	if url == "" {
		return nil
	}
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("skip %s because %+v", url , err)
		return nil
	}
	defer response.Body.Close()
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	return doc
}

func (s *Scraper) handleImg(_ string, selection *goquery.Selection) {
	url := selection.AttrOr("src", "")
	if url == "" {
		return
	}

	// skip if exists
	_, filename := path.Split(url)
	path := fmt.Sprintf("./download/%s", filename)
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return
	}

	fmt.Printf("download %s : sleep 7 second for avoiding server down \n", url)
	time.Sleep(7 * time.Second)
	fmt.Printf("start download from %s\n" , url)

	// download images
	resp, err := internal.NewRequest(url, s.name)
	if err != nil {
		fmt.Printf("skip %s because %+v", url, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("image download error %+v", errors.WithStack(err))
		return
	}

	// create file to download folder
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Printf("file create error %+v", errors.WithStack(err))
		return
	}

	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		fmt.Printf("file write error %+v", errors.WithStack(err))
		return
	}
}

