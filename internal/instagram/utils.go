package instagram

import (
	"encoding/json"
	"fmt"
	"instagram-downloader/internal/files"
	"instagram-downloader/internal/model"
	"io/ioutil"
	"net/http"
)

func chooseImages(ig model.InstagramResponse) error {
	if edges := ig.Graphql.ShortcodeMedia.EdgeSidecarToChildren.Edges; len(edges) > 0 {
		for _, edge := range edges {
			err := saveImages(edge.Node, ig.Graphql.ShortcodeMedia.Owner)
			if err != nil {
				return err
			}
		}
	} else {
		image := ig.Graphql.ShortcodeMedia
		err := saveImages(image.Image, image.Owner)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveImages(image model.Image, owner model.Owner) error {
	pics := image.DisplayResources
	largestPic := pics[len(pics)-1]

	err := files.Save(largestPic.Src, fmt.Sprintf("%s_%s/%s.jpg", owner.ID, owner.Username, image.Shortcode))
	if err != nil {
		return err
	}
	return nil
}

func parse(response *http.Response, target interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &target)
	return err
}

func prepareUrl(url string) string {
	return url + "?__a=1"
}
