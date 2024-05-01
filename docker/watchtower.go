package docker

import (
	"fmt"
	"kaki-tech/kt-bot/config"
	"net/http"
	"strings"
)

var watchTowerUrl string
var watchTowerApiToken string

func configureWatchtower() {
	watchTowerUrl = config.GetVariable(config.WatchtowerApiUrlKey)

	if !strings.HasSuffix(watchTowerUrl, "/") {
		watchTowerUrl += "/"
	}

	watchTowerUrl += "v1/update?image="

	watchTowerApiToken = config.GetVariable(config.WatchtowerApiTokenKey)
}

func updateWithWatchtower(image string) error {
	updateUrl := watchTowerUrl + image

	req, _ := http.NewRequest("GET", updateUrl, nil)

	if len(watchTowerApiToken) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", watchTowerApiToken))
	}

	res, err := http.DefaultClient.Do(req)

	if err == nil && res.StatusCode != 200 {
		return fmt.Errorf("unable to update image %v: %v", image, err)
	} else {
		return err
	}
}
