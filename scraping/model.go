package scraping

type Site interface {
	GetThumbnailQuery() string
	GetBigImageQuery() string
	Name() string
}

type SiteAttribute struct {
	name string
}

func (s SiteAttribute) Name() string{
	return s.name
}

type Tshundora struct {
	SiteAttribute
}

func (Tshundora) GetThumbnailQuery() string {
	return "div.home_tall_box > div.home-img > a"
}

func (Tshundora) GetBigImageQuery() string {
	return "div.single_inside_content > div.post-img > a"
}

type Wallpaperboys struct {
	SiteAttribute
}

func (Wallpaperboys) GetThumbnailQuery() string {
	return "div#content_inside > div.home_post_box > a"
}

func (Wallpaperboys) GetBigImageQuery() string {
	return "div.main_single_content > div.wallpaper-t-img > a"
}
