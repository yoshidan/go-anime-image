package scraping

import "fmt"

type Wallpaperboys struct {
	url string
	name string
}

func (p *Wallpaperboys) Download() {
	scraper := &Scraper{
		thumbnailQuery:"div#content_inside > div.home_post_box > a",
		bigImageQuery: "div.main_single_content > div.wallpaper-t > a",
		name: p.name,
	}
	scraper.Execute(p.url)
}

func NewWallpaperboys(keyword string) *Wallpaperboys {
	name := "wallpaperboys.com"
	url := fmt.Sprintf("https://%s?s=%s", name, keyword)
	return &Wallpaperboys{url, name}
}
