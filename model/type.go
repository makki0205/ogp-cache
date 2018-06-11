package model

type OGP struct {
	Title        string   `json:"title"`
	Type         string   `json:"type"`
	Description  string   `json:"description"`
	Url          string   `json:"url"`
	SiteName     string   `json:"site_name"`
	Image        string   `json:"image"`
	ShortcutIcon string   `json:"shortcut_icon"`
	Twtter       Twtter   `json:"twtter"`
	Facebook     Facebook `json:"facebook"`
}

type Twtter struct {
	Card string `json:"card"`
	Site string `json:"site"`
}

type Facebook struct {
	AppID string `json:"app_id"`
}
