package service

import (
	"github.com/julianshen/og"
	"github.com/makki0205/ogp-cache/model"
)

func PurseOgp(url string) (model.OGP, error) {
	ogp := model.OGP{}
	err := og.GetPageDataFromUrl(url, &ogp)
	return ogp, err
}
