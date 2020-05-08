package instagram

import (
	"instagram-downloader/internal/model"
	"instagram-downloader/internal/network"
)

func TakeImages(url string) {
	obj := model.InstagramResponse{}
	url = prepareUrl(url)
	response, err := network.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	err = parse(response, &obj)
	if err != nil {
		panic(err)
	}

	err = chooseImages(obj)
	if err != nil {
		panic(err)
	}
}
