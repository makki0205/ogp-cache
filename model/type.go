package model

type OGP struct {
	Title        string   `json:"title" meta:"og:title"`
	Type         string   `json:"type" meta:"og:type"`
	Description  string   `json:"description" meta:"og:description"`
	Url          string   `json:"url" meta:"og:url"`
	SiteName     string   `json:"site_name" meta:"og:site_name"`
	Image        string   `json:"image" meta:"og:image"`
	ShortcutIcon string   `json:"-"`
	Twtter       Twtter   `json:"twtter"`
	Facebook     Facebook `json:"facebook"`
}

type Twtter struct {
	Card string `json:"card" meta:"twitter:card"`
	Site string `json:"site" meta:"twitter:site"`
}

type Facebook struct {
	AppID string `json:"app_id" meta:"fb:app_id"`
}
